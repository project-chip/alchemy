package zap

import (
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/zap/render"
)

type Command struct {
	FeatureXML             bool `default:"true" aliases:"featureXML" help:"write new style feature XML" group:"ZAP:"`
	ConformanceXML         bool `default:"true" aliases:"conformanceXML" help:"write new style conformance XML" group:"ZAP:"`
	EndpointCompositionXML bool `default:"false" aliases:"endpointCompositionXML" help:"write new style endpoint composition XML" group:"ZAP:"`
	SpecOrder              bool `default:"false" aliases:"specOrder" help:"write ZAP template XML in spec order" group:"ZAP:"`
	ExtendedQuality        bool `default:"false" aliases:"extendedQuality" help:"write quality element with all qualities, suppressing redundant attributes" group:"ZAP:"`

	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	sdk.SDKOptions             `embed:""`

	Paths []string `arg:"" optional:"" help:"Paths of AsciiDoc files to generate ZAP templates for"`
}

func (z *Command) Run(alchemy *cli.Alchemy) (err error) {

	var options render.Options

	options.Parser = z.ParserOptions.ToOptions()
	options.AsciiSettings = z.ASCIIDocAttributes.ToList()
	options.Pipeline = z.ProcessingOptions

	options.Template = append(options.Template, render.GenerateConformanceXML(z.ConformanceXML))
	options.Template = append(options.Template, render.ExtendedQuality(z.ExtendedQuality))
	options.Template = append(options.Template, render.SpecOrder(z.SpecOrder))
	options.Template = append(options.Template, render.AsciiAttributes(options.AsciiSettings))

	options.DeviceTypes = append(options.DeviceTypes, render.DeviceTypePatcherGenerateFeatureXML(z.FeatureXML))
	options.DeviceTypes = append(options.DeviceTypes, render.DeviceTypePatcherFullEndpointComposition(z.EndpointCompositionXML))

	var output render.Output
	output, err = render.Pipeline(alchemy, z.SdkRoot, z.Paths, options)
	if err != nil {
		return
	}

	stringWriter := files.NewWriter[string]("", z.OutputOptions)
	if output.ZapTemplateDocs != nil && output.ZapTemplateDocs.Size() > 0 {
		stringWriter.SetName("Writing ZAP templates")
		err = stringWriter.Write(alchemy, output.ZapTemplateDocs, options.Pipeline)
		if err != nil {
			return err
		}
	}

	byteWriter := files.NewWriter[[]byte]("", z.OutputOptions)
	if output.ProvisionalDocs != nil && output.ProvisionalDocs.Size() > 0 {
		byteWriter.SetName("Writing provisional docs")
		err = byteWriter.Write(alchemy, output.ProvisionalDocs, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.PatchedDeviceTypes != nil && output.PatchedDeviceTypes.Size() > 0 {
		byteWriter.SetName("Writing deviceTypes")
		err = byteWriter.Write(alchemy, output.PatchedDeviceTypes, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.PatchedNamespaces != nil && output.PatchedNamespaces.Size() > 0 {
		byteWriter.SetName("Writing namespaces")
		err = byteWriter.Write(alchemy, output.PatchedNamespaces, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.GlobalObjectFiles != nil && output.GlobalObjectFiles.Size() > 0 {
		stringWriter.SetName("Writing global objects")
		err = stringWriter.Write(alchemy, output.GlobalObjectFiles, options.Pipeline)
		if err != nil {
			return err
		}
	}

	if output.ClusterList != nil && output.ClusterList.Size() > 0 {
		byteWriter.SetName("Writing cluster list")
		err = byteWriter.Write(alchemy, output.ClusterList, options.Pipeline)
	}

	if output.ZclJson != nil && output.ZclJson.Size() > 0 {
		byteWriter.SetName("Writing ZCL JSON")
		err = byteWriter.Write(alchemy, output.ZclJson, options.Pipeline)
	}

	return
}
