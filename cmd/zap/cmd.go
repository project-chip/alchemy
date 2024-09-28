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
}

func zapTemplates(cmd *cobra.Command, args []string) (err error) {

	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")

	errata.LoadErrataConfig(specRoot)

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	specFiles, err := pipeline.Start[struct{}](cxt, spec.Targeter(specRoot))
	if err != nil {
		return err
	}

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return err
	}
	specDocs, err := pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder()
	specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	patchSpec(specBuilder.Spec)

	var appClusterIndexes pipeline.Map[string, *pipeline.Data[*spec.Doc]]
	appClusterIndexes, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return err
	}

	pipeline.ProcessSerialFunc[*spec.Doc, *spec.Doc](cxt, pipelineOptions, appClusterIndexes, "Assigning index domains", func(cxt context.Context, input *pipeline.Data[*spec.Doc], index, total int32) (outputs []*pipeline.Data[*spec.Doc], extra []*pipeline.Data[*spec.Doc], err error) {
		doc := input.Content
		top := parse.FindFirst[*spec.Section](doc.Elements())
		if top != nil {
			doc.Domain = zap.StringToDomain(top.Name)
			slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		}
		return
	})

	if len(args) > 0 { // Filter the spec by whatever extra args were passed
		filter := files.NewPathFilter[*spec.Doc](args)
		specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	var clusters pipeline.Map[string, *pipeline.Data[*spec.Doc]]
	var deviceTypes pipeline.Map[string, *pipeline.Data[[]*matter.DeviceType]]
	var namespaces pipeline.Map[string, *pipeline.Data[[]*matter.Namespace]]
	clusters, deviceTypes, namespaces, err = generate.SplitZAPDocs(cxt, specDocs)
	if err != nil {
		return err
	}

	var templateOptions []generate.TemplateOption
	featureXML, _ := cmd.Flags().GetBool("featureXML")
	if featureXML {
		templateOptions = append(templateOptions, generate.GenerateFeatureXML(true))
	}
	templateOptions = append(templateOptions, generate.AsciiAttributes(asciiSettings))
	templateOptions = append(templateOptions, generate.SpecRoot(specRoot))

	var zapTemplateDocs pipeline.Map[string, *pipeline.Data[string]]
	var globalObjectFiles pipeline.Map[string, *pipeline.Data[string]]
	var provisionalZclFiles pipeline.Map[string, *pipeline.Data[struct{}]]
	var clusterAliases pipeline.Map[string, []string]
	if clusters.Size() > 0 {
		templateGenerator := generate.NewTemplateGenerator(specBuilder.Spec, fileOptions, pipelineOptions, sdkRoot, templateOptions...)
		zapTemplateDocs, err = pipeline.Process[*spec.Doc, string](cxt, pipelineOptions, templateGenerator, clusters)
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

	var patchedDeviceTypes pipeline.Map[string, *pipeline.Data[[]byte]]
	if deviceTypes.Size() > 0 {
		deviceTypePatcher := generate.NewDeviceTypesPatcher(sdkRoot, specBuilder.Spec, clusterAliases)
		patchedDeviceTypes, err = pipeline.Process[[]*matter.DeviceType, []byte](cxt, pipelineOptions, deviceTypePatcher, deviceTypes)
		if err != nil {
			return err
		}
	}

	var patchedNamespaces pipeline.Map[string, *pipeline.Data[[]byte]]
	if namespaces.Size() > 0 {
		namespacePatcher := generate.NewNamespacePatcher(sdkRoot, specBuilder.Spec)
		patchedNamespaces, err = pipeline.Process[[]*matter.Namespace, []byte](cxt, pipelineOptions, namespacePatcher, namespaces)
		if err != nil {
			return err
		}
	}

	var clusterList pipeline.Map[string, *pipeline.Data[[]byte]]
	clusterListPatcher := generate.NewClusterListPatcher(sdkRoot)
	if clusters.Size() > 0 {
		clusterList, err = pipeline.Process[*spec.Doc, []byte](cxt, pipelineOptions, clusterListPatcher, clusters)
		if err != nil {
			return
		}
	}

	var provisionalDocs pipeline.Map[string, *pipeline.Data[[]byte]]
	if provisionalZclFiles != nil && provisionalZclFiles.Size() > 0 {
		provisionalSaver := generate.NewProvisionalPatcher(sdkRoot)
		provisionalDocs, err = pipeline.Process[struct{}, []byte](cxt, pipelineOptions, provisionalSaver, provisionalZclFiles)
		if err != nil {
			return err
		}

	}

	stringWriter := files.NewWriter[string]("", fileOptions)
	if zapTemplateDocs != nil && zapTemplateDocs.Size() > 0 {
		stringWriter.SetName("Writing ZAP templates")
		_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, stringWriter, zapTemplateDocs)
		if err != nil {
			return err
		}

	}

	byteWriter := files.NewWriter[[]byte]("", fileOptions)
	if provisionalDocs != nil && provisionalDocs.Size() > 0 {
		byteWriter.SetName("Writing provisional docs")
		_, err = pipeline.Process[[]byte, struct{}](cxt, pipelineOptions, byteWriter, provisionalDocs)
		if err != nil {
			return err
		}
	}

	if patchedDeviceTypes != nil && patchedDeviceTypes.Size() > 0 {
		byteWriter.SetName("Writing deviceTypes")
		_, err = pipeline.Process[[]byte, struct{}](cxt, pipelineOptions, byteWriter, patchedDeviceTypes)
		if err != nil {
			return err
		}
	}

	if patchedNamespaces != nil && patchedNamespaces.Size() > 0 {
		byteWriter.SetName("Writing namespaces")
		_, err = pipeline.Process[[]byte, struct{}](cxt, pipelineOptions, byteWriter, patchedNamespaces)
		if err != nil {
			return err
		}
	}

	if globalObjectFiles != nil && globalObjectFiles.Size() > 0 {
		stringWriter.SetName("Writing global objects")
		_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, stringWriter, globalObjectFiles)
		if err != nil {
			return err
		}
	}

	if clusterList != nil && clusterList.Size() > 0 {
		byteWriter.SetName("Writing cluster list")
		_, err = pipeline.Process[[]byte, struct{}](cxt, pipelineOptions, byteWriter, clusterList)
		if err != nil {
			return err
		}
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
	fixedLabel, ok := spec.ClustersByName["Fixed Label"]
	if !ok {
		slog.Warn("Could not find Fixed Label cluster")
		return
	}
	label, ok := spec.ClustersByName["Label"]
	if !ok {
		slog.Warn("Could not find Label cluster")
		return
	}
	for _, s := range label.Structs {
		if s.Name == "LabelStruct" {
			s.ParentEntity = fixedLabel
			fixedLabel.Structs = append(fixedLabel.Structs, s)
			break
		}
	}
}
