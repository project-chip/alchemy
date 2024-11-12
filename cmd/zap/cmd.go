package zap

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/zap"
	"github.com/project-chip/alchemy/zap/generate"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "zap",
	Short: "transmute the Matter spec into ZAP templates",
	RunE:  zapTemplates,
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
	Command.Flags().Bool("featureXML", true, "write new style feature XML")
	Command.Flags().Bool("conformanceXML", true, "write new style conformance XML")
	Command.Flags().Bool("specOrder", false, "write ZAP template XML in spec order")
}

func zapTemplates(cmd *cobra.Command, args []string) (err error) {

	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")

	errata.LoadErrataConfig(specRoot)

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	specFiles, err := pipeline.Start(cxt, spec.Targeter(specRoot))
	if err != nil {
		return err
	}

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return err
	}
	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder()
	specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	patchSpec(specBuilder.Spec)

	var appClusterIndexes spec.DocSet
	appClusterIndexes, err = pipeline.Collective(cxt, pipelineOptions, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return err
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

	_, err = pipeline.Parallel(cxt, pipelineOptions, pipeline.ParallelFunc("Assigning index domains", domainIndexer), appClusterIndexes)
	if err != nil {
		return err
	}

	if len(args) > 0 { // Filter the spec by whatever extra args were passed
		filter := files.NewPathFilter[*spec.Doc](args)
		specDocs, err = pipeline.Collective(cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	var clusters spec.DocSet
	var deviceTypes pipeline.Map[string, *pipeline.Data[[]*matter.DeviceType]]
	var namespaces pipeline.Map[string, *pipeline.Data[[]*matter.Namespace]]
	clusters, deviceTypes, namespaces, err = generate.SplitZAPDocs(cxt, specDocs)
	if err != nil {
		return err
	}

	var templateOptions []generate.TemplateOption
	featureXML, _ := cmd.Flags().GetBool("featureXML")
	templateOptions = append(templateOptions, generate.GenerateFeatureXML(featureXML))
	conformanceXML, _ := cmd.Flags().GetBool("conformanceXML")
	specOrder, _ := cmd.Flags().GetBool("specOrder")
	templateOptions = append(templateOptions, generate.GenerateConformanceXML(conformanceXML))
	templateOptions = append(templateOptions, generate.SpecOrder(specOrder))
	templateOptions = append(templateOptions, generate.AsciiAttributes(asciiSettings))
	templateOptions = append(templateOptions, generate.SpecRoot(specRoot))

	var zapTemplateDocs pipeline.StringSet
	var globalObjectFiles pipeline.StringSet
	var provisionalZclFiles pipeline.Paths
	var clusterAliases pipeline.Map[string, []string]
	if clusters.Size() > 0 {
		templateGenerator := generate.NewTemplateGenerator(specBuilder.Spec, fileOptions, pipelineOptions, sdkRoot, templateOptions...)
		zapTemplateDocs, err = pipeline.Parallel(cxt, pipelineOptions, templateGenerator, clusters)
		if err != nil {
			return err
		}
		provisionalZclFiles = templateGenerator.ProvisionalZclFiles
		clusterAliases = templateGenerator.ClusterAliases

		globalObjectFiles, err = templateGenerator.RenderGlobalObjecs(cxt)
		if err != nil {
			return err
		}
	} else {
		clusterAliases = pipeline.NewConcurrentMap[string, []string]()
	}

	var patchedDeviceTypes pipeline.FileSet
	if deviceTypes.Size() > 0 {
		deviceTypePatcher := generate.NewDeviceTypesPatcher(sdkRoot, specBuilder.Spec, clusterAliases, generate.DeviceTypePatcherGenerateFeatureXML(featureXML))
		patchedDeviceTypes, err = pipeline.Collective(cxt, pipelineOptions, deviceTypePatcher, deviceTypes)
		if err != nil {
			return err
		}
	}

	var patchedNamespaces pipeline.FileSet
	if namespaces.Size() > 0 {
		namespacePatcher := generate.NewNamespacePatcher(sdkRoot, specBuilder.Spec)
		patchedNamespaces, err = pipeline.Collective(cxt, pipelineOptions, namespacePatcher, namespaces)
		if err != nil {
			return err
		}
	}

	var clusterList pipeline.FileSet
	clusterListPatcher := generate.NewClusterListPatcher(sdkRoot)
	if clusters.Size() > 0 {
		clusterList, err = pipeline.Collective(cxt, pipelineOptions, clusterListPatcher, clusters)
		if err != nil {
			return
		}
	}

	var provisionalDocs pipeline.FileSet
	if provisionalZclFiles != nil && provisionalZclFiles.Size() > 0 {
		provisionalSaver := generate.NewProvisionalPatcher(sdkRoot, specBuilder.Spec)
		provisionalDocs, err = pipeline.Collective(cxt, pipelineOptions, provisionalSaver, provisionalZclFiles)
		if err != nil {
			return err
		}

	}

	stringWriter := files.NewWriter[string]("", fileOptions)
	if zapTemplateDocs != nil && zapTemplateDocs.Size() > 0 {
		stringWriter.SetName("Writing ZAP templates")
		err = stringWriter.Write(cxt, zapTemplateDocs, pipelineOptions)
		if err != nil {
			return err
		}
	}

	byteWriter := files.NewWriter[[]byte]("", fileOptions)
	if provisionalDocs != nil && provisionalDocs.Size() > 0 {
		byteWriter.SetName("Writing provisional docs")
		err = byteWriter.Write(cxt, provisionalDocs, pipelineOptions)
		if err != nil {
			return err
		}
	}

	if patchedDeviceTypes != nil && patchedDeviceTypes.Size() > 0 {
		byteWriter.SetName("Writing deviceTypes")
		err = byteWriter.Write(cxt, patchedDeviceTypes, pipelineOptions)
		if err != nil {
			return err
		}
	}

	if patchedNamespaces != nil && patchedNamespaces.Size() > 0 {
		byteWriter.SetName("Writing namespaces")
		err = byteWriter.Write(cxt, patchedNamespaces, pipelineOptions)
		if err != nil {
			return err
		}
	}

	if globalObjectFiles != nil && globalObjectFiles.Size() > 0 {
		stringWriter.SetName("Writing global objects")
		err = stringWriter.Write(cxt, globalObjectFiles, pipelineOptions)
		if err != nil {
			return err
		}
	}

	if clusterList != nil && clusterList.Size() > 0 {
		byteWriter.SetName("Writing cluster list")
		err = byteWriter.Write(cxt, clusterList, pipelineOptions)
	}
	return

}

func patchSpec(spec *spec.Specification) {
	/* This is a hacky workaround for a spec problem: SemanticTagStruct is defined twice, in two different ways.
	The first is a global struct that's used by the Descriptor cluster
	The second is a cluster-level struct on Mode Select
	Inserting one as a global object, and the other as a struct on Mode Select breaks zap
	*/
	desc, ok := spec.ClustersByName["Descriptor"]
	if !ok {
		slog.Warn("Could not find Descriptor cluster")
		return
	}
	for o := range spec.GlobalObjects {
		s, ok := o.(*matter.Struct)
		if !ok {
			continue
		}

		if s.Name == "SemanticTagStruct" {
			desc.AddStructs(s)
			delete(spec.GlobalObjects, s)
			break
		}
	}
	/*
		Another hacky workaround: the spec defines LabelStruct under a base cluster called Label Cluster, but the
		ZAP XML has this struct under Fixed Label
	*/
	fixedLabelCluster, ok := spec.ClustersByName["Fixed Label"]
	if !ok {
		slog.Warn("Could not find Fixed Label cluster")
		return
	}
	labelCluster, ok := spec.ClustersByName["Label"]
	if !ok {
		slog.Warn("Could not find Label cluster")
		return
	}
	for _, s := range labelCluster.Structs {
		if s.Name == "LabelStruct" {
			fixedLabelCluster.MoveStruct(s)
			break
		}
	}
}
