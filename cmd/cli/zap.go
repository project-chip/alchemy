package cli

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/zap"
	"github.com/project-chip/alchemy/zap/regen"
	"github.com/project-chip/alchemy/zap/render"
)

type ZAP struct {
	ZAPXML   ZAPXML   `cmd:"" name:"xml" default:"withargs"`
	ZAPRegen ZAPRegen `cmd:"" name:"regen"`
}

type ZAPXML struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	spec.FilterOptions         `embed:""`
	sdk.SDKOptions             `embed:""`
	render.TemplateOptions     `embed:""`
}

func (z *ZAPXML) Run(cc *Context) (err error) {

	err = sdk.CheckAlchemyVersion(z.SdkRoot)
	if err != nil {
		return
	}

	var specDocs spec.DocSet
	var specification *spec.Specification
	specification, specDocs, err = spec.Parse(cc, z.ParserOptions, z.ProcessingOptions, []spec.BuilderOption{spec.PatchForSdk(true)}, z.ASCIIDocAttributes.ToList())
	if err != nil {
		return
	}

	err = sdk.ApplyErrata(specification)
	if err != nil {
		return
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specification, z.FilterOptions, z.ProcessingOptions)
	if err != nil {
		return
	}

	var clusters, deviceTypes, namespaces, globalObjectDependencies spec.DocSet
	clusters, deviceTypes, namespaces, err = render.SplitZAPDocs(cc, specification, specDocs)
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

		clusterListPatcher := render.NewClusterListPatcher(specification, z.SdkRoot)
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

type ZAPRegen struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	spec.FilterOptions         `embed:""`
	sdk.SDKOptions             `embed:""`
	render.TemplateOptions     `embed:""`
}

func (z *ZAPRegen) Run(cc *Context) (err error) {
	zapTargeter := regen.Targeter(z.SdkRoot)

	var specification *spec.Specification
	specification, _, err = spec.Parse(cc, z.ParserOptions, z.ProcessingOptions, []spec.BuilderOption{spec.PatchForSdk(true)}, z.ASCIIDocAttributes.ToList())
	if err != nil {
		return
	}

	err = sdk.ApplyErrata(specification)
	if err != nil {
		return
	}

	var zapPaths pipeline.Paths
	zapPaths, err = pipeline.Start(cc, zapTargeter)
	if err != nil {
		return
	}

	var reader regen.Reader
	reader, err = regen.NewReader()
	if err != nil {
		return
	}

	var zapFiles pipeline.Map[string, *pipeline.Data[*zap.File]]
	zapFiles, err = pipeline.Parallel(cc, z.ProcessingOptions, reader, zapPaths)
	if err != nil {
		return
	}

	var renderer regen.IdlRenderer
	renderer, err = regen.NewIdlRenderer(specification)
	if err != nil {
		return
	}

	var matterFiles pipeline.StringSet

	matterFiles, err = pipeline.Parallel(cc, z.ProcessingOptions, renderer, zapFiles)
	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing test scripts", z.OutputOptions)
	err = writer.Write(cc, matterFiles, z.ProcessingOptions)

	return nil
}
