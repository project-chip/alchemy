package zap

import (
	"context"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
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
	Command.Flags().Bool("endpointCompositionXML", false, "write new style endpoint composition XML")
	Command.Flags().Bool("specOrder", false, "write ZAP template XML in spec order")
}

func zapTemplates(cmd *cobra.Command, args []string) (err error) {

	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")

	var options generate.Options

	fileOptions := files.Flags(cmd)

	options.AsciiSettings = common.ASCIIDocAttributes(cmd)
	options.Pipeline = pipeline.Flags(cmd)

	featureXML, _ := cmd.Flags().GetBool("featureXML")
	options.Template = append(options.Template, generate.GenerateFeatureXML(featureXML))
	conformanceXML, _ := cmd.Flags().GetBool("conformanceXML")
	endpointCompositionXML, _ := cmd.Flags().GetBool("endpointCompositionXML")
	specOrder, _ := cmd.Flags().GetBool("specOrder")
	options.Template = append(options.Template, generate.GenerateConformanceXML(conformanceXML))
	options.Template = append(options.Template, generate.SpecOrder(specOrder))
	options.Template = append(options.Template, generate.AsciiAttributes(options.AsciiSettings))
	options.Template = append(options.Template, generate.SpecRoot(specRoot))

	options.DeviceTypes = append(options.DeviceTypes, generate.DeviceTypePatcherGenerateFeatureXML(featureXML))
	options.DeviceTypes = append(options.DeviceTypes, generate.DeviceTypePatcherFullEndpointComposition(endpointCompositionXML))

	var output generate.Output
	output, err = generate.Pipeline(cxt, specRoot, sdkRoot, args, options)
	if err != nil {
		return
	}

	stringWriter := files.NewWriter[string]("", fileOptions)
	if output.ZapTemplateDocs != nil && output.ZapTemplateDocs.Size() > 0 {
		stringWriter.SetName("Writing ZAP templates")
		err = stringWriter.Write(cxt, output.ZapTemplateDocs, options.Pipeline)
		if err != nil {
			return err
		}
	}

	byteWriter := files.NewWriter[[]byte]("", fileOptions)
	if output.ProvisionalDocs != nil && output.ProvisionalDocs.Size() > 0 {
		byteWriter.SetName("Writing provisional docs")
		err = byteWriter.Write(cxt, output.ProvisionalDocs, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.PatchedDeviceTypes != nil && output.PatchedDeviceTypes.Size() > 0 {
		byteWriter.SetName("Writing deviceTypes")
		err = byteWriter.Write(cxt, output.PatchedDeviceTypes, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.PatchedNamespaces != nil && output.PatchedNamespaces.Size() > 0 {
		byteWriter.SetName("Writing namespaces")
		err = byteWriter.Write(cxt, output.PatchedNamespaces, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.GlobalObjectFiles != nil && output.GlobalObjectFiles.Size() > 0 {
		stringWriter.SetName("Writing global objects")
		err = stringWriter.Write(cxt, output.GlobalObjectFiles, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.ClusterList != nil && output.ClusterList.Size() > 0 {
		byteWriter.SetName("Writing cluster list")
		err = byteWriter.Write(cxt, output.ClusterList, options.Pipeline)
	}

	return

}
