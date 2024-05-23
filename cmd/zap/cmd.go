package zap

import (
	"context"
	"log/slog"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/zap"
	"github.com/hasty/alchemy/zap/generate"
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
	Command.Flags().Bool("featureXML", false, "write new style feature XML")
}

func zapTemplates(cmd *cobra.Command, args []string) (err error) {

	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	asciiSettings = append(spec.GithubSettings(), asciiSettings...)

	specFiles, err := pipeline.Start[struct{}](cxt, files.SpecTargeter(specRoot))
	if err != nil {
		return err
	}

	docParser := spec.NewParser(asciiSettings)
	specDocs, err := pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	var specParser files.SpecParser
	specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, &specParser, specDocs)
	if err != nil {
		return err
	}

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
	clusters, deviceTypes, err = generate.SplitZAPDocs(cxt, specDocs)
	if err != nil {
		return err
	}

	var templateOptions []generate.TemplateOption
	featureXML, _ := cmd.Flags().GetBool("featureXML")
	if featureXML {
		templateOptions = append(templateOptions, generate.GenerateFeatureXML(true))
	}

	templateGenerator := generate.NewTemplateGenerator(specParser.Spec, fileOptions, pipelineOptions, sdkRoot, templateOptions...)
	zapTemplateDocs, err := pipeline.Process[*spec.Doc, string](cxt, pipelineOptions, templateGenerator, clusters)
	if err != nil {
		return err
	}

	deviceTypePatcher := generate.NewDeviceTypesPatcher(sdkRoot, specParser.Spec)
	var patchedDeviceTypes pipeline.Map[string, *pipeline.Data[[]byte]]
	patchedDeviceTypes, err = pipeline.Process[[]*matter.DeviceType, []byte](cxt, pipelineOptions, deviceTypePatcher, deviceTypes)
	if err != nil {
		return err
	}

	clusterListPatcher := generate.NewClusterListPatcher(sdkRoot)
	var clusterList pipeline.Map[string, *pipeline.Data[[]byte]]
	clusterList, err = pipeline.Process[*spec.Doc, []byte](cxt, pipelineOptions, clusterListPatcher, clusters)
	if err != nil {
		return
	}

	provisionalSaver := generate.NewProvisionalPatcher(sdkRoot)
	var provisionalDocs pipeline.Map[string, *pipeline.Data[[]byte]]
	provisionalDocs, err = pipeline.Process[struct{}, []byte](cxt, pipelineOptions, provisionalSaver, templateGenerator.ProvisionalZclFiles)
	if err != nil {
		return err
	}

	stringWriter := files.NewWriter[string]("Writing ZAP templates", fileOptions)

	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, stringWriter, zapTemplateDocs)
	if err != nil {
		return err
	}

	byteWriter := files.NewWriter[[]byte]("Writing provisional docs", fileOptions)
	_, err = pipeline.Process[[]byte, struct{}](cxt, pipelineOptions, byteWriter, provisionalDocs)
	if err != nil {
		return err
	}

	byteWriter.SetName("Writing deviceTypes")
	_, err = pipeline.Process[[]byte, struct{}](cxt, pipelineOptions, byteWriter, patchedDeviceTypes)
	if err != nil {
		return err
	}

	byteWriter.SetName("Writing cluster list")
	_, err = pipeline.Process[[]byte, struct{}](cxt, pipelineOptions, byteWriter, clusterList)
	if err != nil {
		return err
	}
	return

}
