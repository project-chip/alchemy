package render

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/zap"
)

type Options struct {
	Pipeline      pipeline.Options
	AsciiSettings []asciidoc.AttributeName
	Template      []TemplateOption
	DeviceTypes   []DeviceTypePatcherOption
	Parser        []spec.ParserOption
}

type Output struct {
	ZapTemplateDocs    pipeline.StringSet
	GlobalObjectFiles  pipeline.StringSet
	PatchedDeviceTypes pipeline.FileSet
	PatchedNamespaces  pipeline.FileSet
	ClusterList        pipeline.FileSet
	ProvisionalDocs    pipeline.FileSet
	ZclJson            pipeline.FileSet
}

func Pipeline(cxt context.Context, sdkRoot string, docPaths []string, options Options) (output Output, err error) {

	err = sdk.CheckAlchemyVersion(sdkRoot)
	if err != nil {
		return
	}

	var specParser spec.Parser
	specParser, err = spec.NewParser(options.AsciiSettings, options.Parser...)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	var specFiles pipeline.Paths
	specFiles, err = pipeline.Start(cxt, specParser.Targets)
	if err != nil {
		return
	}

	var specDocs spec.DocSet
	specDocs, err = pipeline.Parallel(cxt, options.Pipeline, specParser, specFiles)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(specParser.Root)
	specDocs, err = pipeline.Collective(cxt, options.Pipeline, &specBuilder, specDocs)
	if err != nil {
		return
	}

	err = spec.PatchSpecForSdk(specBuilder.Spec)
	if err != nil {
		return
	}

	var appClusterIndexes spec.DocSet
	appClusterIndexes, err = pipeline.Collective(cxt, options.Pipeline, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return
	}

	domainIndexer := func(cxt context.Context, input *pipeline.Data[*spec.Doc], index, total int32) (outputs []*pipeline.Data[*spec.Doc], extra []*pipeline.Data[*spec.Doc], err error) {
		doc := input.Content
		top := parse.FindFirst[*spec.Section](doc.Elements())
		if top != nil {
			doc.Domain = zap.StringToDomain(top.Name)
			slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		}
		return
	}

	_, err = pipeline.Parallel(cxt, options.Pipeline, pipeline.ParallelFunc("Assigning index domains", domainIndexer), appClusterIndexes)
	if err != nil {
		return
	}

	if len(docPaths) > 0 { // Filter the spec by whatever extra args were passed
		filter := paths.NewFilter[*spec.Doc](specParser.Root, docPaths)
		specDocs, err = pipeline.Collective(cxt, options.Pipeline, filter, specDocs)
		if err != nil {
			return
		}
	}

	var clusters spec.DocSet
	var deviceTypes spec.DocSet
	var namespaces pipeline.Map[string, *pipeline.Data[[]*matter.Namespace]]
	clusters, deviceTypes, namespaces, err = SplitZAPDocs(cxt, specDocs)
	if err != nil {
		return
	}

	var provisionalZclFiles pipeline.Paths
	var clusterAliases pipeline.Map[string, []string]
	if clusters.Size() > 0 {
		templateGenerator := NewTemplateGenerator(specBuilder.Spec, options.Pipeline, sdkRoot, options.Template...)
		output.ZapTemplateDocs, err = pipeline.Parallel(cxt, options.Pipeline, templateGenerator, clusters)
		if err != nil {
			return
		}
		provisionalZclFiles = templateGenerator.ProvisionalZclFiles
		clusterAliases = templateGenerator.ClusterAliases

		output.GlobalObjectFiles, err = templateGenerator.RenderGlobalObjecs(cxt)
		if err != nil {
			return
		}
	} else {
		clusterAliases = pipeline.NewConcurrentMap[string, []string]()
		provisionalZclFiles = pipeline.NewConcurrentMap[string, *pipeline.Data[struct{}]]()
	}

	if deviceTypes.Size() > 0 {
		deviceTypePatcher := NewDeviceTypesPatcher(sdkRoot, specBuilder.Spec, clusterAliases, options.DeviceTypes...)
		output.PatchedDeviceTypes, err = pipeline.Collective(cxt, options.Pipeline, deviceTypePatcher, deviceTypes)
		if err != nil {
			return
		}
	}

	if namespaces.Size() > 0 {
		namespacePatcher := NewNamespacePatcher(sdkRoot, specBuilder.Spec)
		output.PatchedNamespaces, err = pipeline.Collective(cxt, options.Pipeline, namespacePatcher, namespaces)
		if err != nil {
			return
		}
	}

	if clusters.Size() > 0 {
		clusterListPatcher := NewClusterListPatcher(sdkRoot)
		output.ClusterList, err = pipeline.Collective(cxt, options.Pipeline, clusterListPatcher, clusters)
		if err != nil {
			return
		}

		zclPatcher := NewZclPatcher(sdkRoot, specBuilder.Spec, provisionalZclFiles)
		output.ZclJson, err = pipeline.Collective(cxt, options.Pipeline, zclPatcher, clusters)
		if err != nil {
			return
		}
	}

	if provisionalZclFiles.Size() > 0 {
		provisionalPatcher := NewProvisionalPatcher(sdkRoot, specBuilder.Spec)
		output.ProvisionalDocs, err = pipeline.Collective(cxt, options.Pipeline, provisionalPatcher, provisionalZclFiles)
		if err != nil {
			return
		}
	}
	return
}
