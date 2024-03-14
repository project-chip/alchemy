package zap

import (
	"context"
	"log/slog"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
	"github.com/hasty/alchemy/zap/generate"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "zap",
	Short: "transmute the Matter spec into ZAP templates",
	RunE:  zapTemplates,
}

func init() {
	Command.Flags().String("specRoot", "", "the root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "", "the root of your clone of project-chip/connectedhomeip")
	Command.Flags().Bool("overwrite", false, "overwrite existing ZAP templates")
	_ = Command.MarkFlagRequired("specRoot")
	_ = Command.MarkFlagRequired("sdkRoot")
}

func zapTemplates(cmd *cobra.Command, args []string) (err error) {

	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")

	asciiSettings := common.AsciiDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	asciiSettings = append(ascii.GithubSettings(), asciiSettings...)

	specFiles, err := pipeline.Start[struct{}](cxt, files.SpecTargeter(specRoot))
	if err != nil {
		return err
	}

	docParser := ascii.NewParser(asciiSettings)
	specDocs, err := pipeline.Process[struct{}, *ascii.Doc](cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	var specParser files.SpecParser
	specDocs, err = pipeline.Process[*ascii.Doc, *ascii.Doc](cxt, pipelineOptions, &specParser, specDocs)
	if err != nil {
		return err
	}

	var appClusterIndexes *xsync.MapOf[string, *pipeline.Data[*ascii.Doc]]
	appClusterIndexes, err = pipeline.Process[*ascii.Doc, *ascii.Doc](cxt, pipelineOptions, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return err
	}

	pipeline.ProcessSerialFunc[*ascii.Doc, *ascii.Doc](cxt, pipelineOptions, appClusterIndexes, "Assigning index domains", func(cxt context.Context, input *pipeline.Data[*ascii.Doc], index, total int32) (outputs []*pipeline.Data[*ascii.Doc], extra []*pipeline.Data[*ascii.Doc], err error) {
		doc := input.Content
		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top != nil {
			doc.Domain = zap.StringToDomain(top.Name)
			slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		}
		return
	})

	if len(args) > 0 { // Filter the spec by whatever extra args were passed
		filter := files.NewPathFilter[*ascii.Doc](args)
		specDocs, err = pipeline.Process[*ascii.Doc, *ascii.Doc](cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	var clusters *xsync.MapOf[string, *pipeline.Data[*ascii.Doc]]
	var deviceTypes *xsync.MapOf[string, *pipeline.Data[[]*matter.DeviceType]]
	clusters, deviceTypes, err = generate.SplitZAPDocs(cxt, specDocs)
	if err != nil {
		return err
	}

	templateGenerator := generate.NewTemplateGenerator(specParser.Spec, fileOptions, pipelineOptions, sdkRoot)
	zapTemplateDocs, err := pipeline.Process[*ascii.Doc, string](cxt, pipelineOptions, templateGenerator, clusters)
	if err != nil {
		return err
	}

	deviceTypePatcher := generate.NewDeviceTypesPatcher(sdkRoot, specParser.Spec)
	var patchedDeviceTypes *xsync.MapOf[string, *pipeline.Data[[]byte]]
	patchedDeviceTypes, err = pipeline.Process[[]*matter.DeviceType, []byte](cxt, pipelineOptions, deviceTypePatcher, deviceTypes)
	if err != nil {
		return err
	}

	clusterListPatcher := generate.NewClusterListPatcher(sdkRoot)
	var clusterList *xsync.MapOf[string, *pipeline.Data[[]byte]]
	clusterList, err = pipeline.Process[*ascii.Doc, []byte](cxt, pipelineOptions, clusterListPatcher, clusters)
	if err != nil {
		return
	}

	provisionalSaver := generate.NewProvisionalPatcher(sdkRoot)
	var provisionalDocs *xsync.MapOf[string, *pipeline.Data[[]byte]]
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
