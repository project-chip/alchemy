package cli

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/zap"
	"github.com/project-chip/alchemy/zap/render"
)

type ZAP struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	spec.FilterOptions         `embed:""`
	sdk.SDKOptions             `embed:""`
	render.TemplateOptions     `embed:""`
}

func (z *ZAP) Run(cc *Context) (err error) {

	err = sdk.CheckAlchemyVersion(z.SdkRoot)
	if err != nil {
		return
	}

	var specDocs spec.DocSet
	var specification *spec.Specification
	specification, specDocs, err = spec.Parse(cc, z.ParserOptions, z.ProcessingOptions, z.ASCIIDocAttributes.ToList())
	if err != nil {
		return
	}

	err = spec.PatchSpecForSdk(specification)
	if err != nil {
		return
	}

	var appClusterIndexes spec.DocSet
	appClusterIndexes, err = pipeline.Collective(cc, z.ProcessingOptions, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return
	}

	domainIndexer := func(cxt context.Context, input *pipeline.Data[*spec.Doc], index, total int32) (outputs []*pipeline.Data[*spec.Doc], extra []*pipeline.Data[*spec.Doc], err error) {
		doc := input.Content
		e := errata.GetErrata(input.Content.Path.Relative)
		if e != nil && e.Spec.Domain != "" {
			doc.Domain = zap.StringToDomain(e.Spec.Domain)
			if doc.Domain != matter.DomainUnknown {
				slog.DebugContext(cxt, "Assigned domain from errata", "file", input.Content.Path.Relative, "domain", doc.Domain)
				return
			}
		}
		top := parse.FindFirst[*asciidoc.Section](doc.Iterator(), doc)
		if top != nil {
			doc.Domain = zap.StringToDomain(doc.SectionName(top))
			slog.DebugContext(cxt, "Assigned domain", "file", input.Content.Path.Relative, "domain", doc.Domain)
		}
		return
	}

	_, err = pipeline.Parallel(cc, z.ProcessingOptions, pipeline.ParallelFunc("Assigning index domains", domainIndexer), appClusterIndexes)
	if err != nil {
		return
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specification, z.FilterOptions, z.ProcessingOptions)
	if err != nil {
		return
	}

	var clusters, deviceTypes, namespaces, globalObjectDependencies spec.DocSet
	clusters, deviceTypes, namespaces, err = render.SplitZAPDocs(cc, specDocs)
	if err != nil {
		return
	}

	if clusters.Size() > 0 {
		dependencyTracer := render.NewDependencyTracer(specification)

		clusters, err = pipeline.Collective(cc, z.ProcessingOptions, dependencyTracer, clusters)
		if err != nil {
			return
		}

		clusters, err = filterSpecErrors(cc, clusters, specification, z.FilterOptions, z.ProcessingOptions)
		if err != nil {
			return
		}

		globalObjectDependencies = dependencyTracer.GlobalObjectDependencies
	}

	if deviceTypes.Size() > 0 {
		deviceTypes, err = filterSpecErrors(cc, deviceTypes, specification, z.FilterOptions, z.ProcessingOptions)
		if err != nil {
			return
		}
	}

	if namespaces.Size() > 0 {
		namespaces, err = filterSpecErrors(cc, namespaces, specification, z.FilterOptions, z.ProcessingOptions)
		if err != nil {
			return
		}
	}

	err = checkSpecErrors(cc, specification, z.FilterOptions, clusters, deviceTypes, namespaces)
	if err != nil {
		return
	}

	var zapTemplateDocs, globalObjectFiles pipeline.StringSet
	var patchedDeviceTypes, patchedNamespaces, clusterList, indexDocs, zclJson pipeline.FileSet

	var clusterAliases pipeline.Map[string, []string]
	if clusters.Size() > 0 {

		var templateGenerator *render.TemplateGenerator
		templateGenerator, err = render.NewTemplateGenerator(specification, z.ProcessingOptions, z.SdkRoot, z.TemplateOptions)
		if err != nil {
			return
		}
		zapTemplateDocs, err = pipeline.Parallel(cc, z.ProcessingOptions, templateGenerator, clusters)
		if err != nil {
			return
		}
		clusterAliases = templateGenerator.ClusterAliases

		globalObjectRenderer := render.NewGlobalObjectsRenderer(specification, z.SdkRoot, templateGenerator)
		globalObjectFiles, err = pipeline.Collective(cc, z.ProcessingOptions, globalObjectRenderer, globalObjectDependencies)
		if err != nil {
			return
		}

	} else {
		clusterAliases = pipeline.NewConcurrentMap[string, []string]()
	}

	if deviceTypes.Size() > 0 {

		deviceTypePatcher := render.NewDeviceTypesPatcher(z.SdkRoot, specification, clusterAliases, z.TemplateOptions)
		patchedDeviceTypes, err = pipeline.Collective(cc, z.ProcessingOptions, deviceTypePatcher, deviceTypes)
		if err != nil {
			return
		}
	}

	if namespaces.Size() > 0 {

		namespacePatcher := render.NewNamespacePatcher(z.SdkRoot, specification)
		patchedNamespaces, err = pipeline.Collective(cc, z.ProcessingOptions, namespacePatcher, namespaces)
		if err != nil {
			return
		}
	}

	if clusters.Size() > 0 {

		clusterListPatcher := render.NewClusterListPatcher(z.SdkRoot)
		clusterList, err = pipeline.Collective(cc, z.ProcessingOptions, clusterListPatcher, clusters)
		if err != nil {
			return
		}

		zclPatcher := render.NewZclPatcher(z.SdkRoot, specification, zapTemplateDocs)
		zclJson, err = pipeline.Collective(cc, z.ProcessingOptions, zclPatcher, clusters)
		if err != nil {
			return
		}

		provisionalPatcher := render.NewIndexFilesPatcher(z.SdkRoot, specification)
		indexDocs, err = pipeline.Collective(cc, z.ProcessingOptions, provisionalPatcher, zapTemplateDocs)
		if err != nil {
			return
		}
	}

	stringWriter := files.NewWriter[string]("", z.OutputOptions)
	if zapTemplateDocs != nil && zapTemplateDocs.Size() > 0 {
		stringWriter.SetName("Writing ZAP templates")
		err = stringWriter.Write(cc, zapTemplateDocs, z.ProcessingOptions)
		if err != nil {
			return err
		}
	}

	byteWriter := files.NewWriter[[]byte]("", z.OutputOptions)
	if indexDocs != nil && indexDocs.Size() > 0 {
		byteWriter.SetName("Writing provisional docs")
		err = byteWriter.Write(cc, indexDocs, z.ProcessingOptions)
		if err != nil {
			return err
		}
	}

	if patchedDeviceTypes != nil && patchedDeviceTypes.Size() > 0 {
		byteWriter.SetName("Writing deviceTypes")
		err = byteWriter.Write(cc, patchedDeviceTypes, z.ProcessingOptions)
		if err != nil {
			return err
		}
	}

	if patchedNamespaces != nil && patchedNamespaces.Size() > 0 {
		byteWriter.SetName("Writing namespaces")
		err = byteWriter.Write(cc, patchedNamespaces, z.ProcessingOptions)
		if err != nil {
			return err
		}
	}

	if globalObjectFiles != nil && globalObjectFiles.Size() > 0 {
		stringWriter.SetName("Writing global objects")
		err = stringWriter.Write(cc, globalObjectFiles, z.ProcessingOptions)
		if err != nil {
			return err
		}
	}

	if clusterList != nil && clusterList.Size() > 0 {
		byteWriter.SetName("Writing cluster list")
		err = byteWriter.Write(cc, clusterList, z.ProcessingOptions)
	}

	if zclJson != nil && zclJson.Size() > 0 {
		byteWriter.SetName("Writing ZCL JSON")
		err = byteWriter.Write(cc, zclJson, z.ProcessingOptions)
	}

	return
}
