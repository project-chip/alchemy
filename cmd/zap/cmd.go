package zap

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/zap/render"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "zap [filename_pattern]",
	Short: "transmute the Matter spec into ZAP templates, optionally filtered to the files specified by filename_pattern",
	RunE:  zapTemplates,
}

func init() {
	flags := Command.Flags()
	flags.String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	flags.String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
	flags.Bool("featureXML", true, "write new style feature XML")
	flags.Bool("conformanceXML", true, "write new style conformance XML")
	flags.Bool("endpointCompositionXML", false, "write new style endpoint composition XML")
	flags.Bool("specOrder", false, "write ZAP template XML in spec order")
	flags.Bool("extendedQuality", false, "write quality element with all qualities, suppressing redundant attributes")
	flags.Bool("inline", false, "use inline parser")
}

func zapTemplates(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()
	flags := cmd.Flags()

	sdkRoot, _ := flags.GetString("sdkRoot")

	var options render.Options

	fileOptions := files.OutputOptions(flags)
	options.Parser = spec.ParserOptions(flags)

	options.AsciiSettings = common.ASCIIDocAttributes(flags)
	options.Pipeline = pipeline.PipelineOptions(flags)

	featureXML, _ := flags.GetBool("featureXML")
	options.Template = append(options.Template, render.GenerateFeatureXML(featureXML))
	conformanceXML, _ := flags.GetBool("conformanceXML")
	extendedQuality, _ := flags.GetBool("extendedQuality")
	endpointCompositionXML, _ := flags.GetBool("endpointCompositionXML")
	specOrder, _ := flags.GetBool("specOrder")
	options.Template = append(options.Template, render.GenerateConformanceXML(conformanceXML))
	options.Template = append(options.Template, render.ExtendedQuality(extendedQuality))
	options.Template = append(options.Template, render.SpecOrder(specOrder))
	options.Template = append(options.Template, render.AsciiAttributes(options.AsciiSettings))

	options.DeviceTypes = append(options.DeviceTypes, render.DeviceTypePatcherGenerateFeatureXML(featureXML))
	options.DeviceTypes = append(options.DeviceTypes, render.DeviceTypePatcherFullEndpointComposition(endpointCompositionXML))

	var output render.Output
	output, err = render.Pipeline(cxt, sdkRoot, args, options)
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

	if output.ZclJson != nil && output.ZclJson.Size() > 0 {
		byteWriter.SetName("Writing ZCL JSON")
		err = byteWriter.Write(cxt, output.ZclJson, options.Pipeline)
	}

	return

}
