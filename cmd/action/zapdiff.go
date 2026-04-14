package action

import (
	"log/slog"

	"github.com/project-chip/alchemy/cmd/cli"
)

type ZAPDiffConfig struct {
	ComparisonSets []ComparisonSet `yaml:"comparison_sets"`
}

type ComparisonSet struct {
	SDKRef                  string `yaml:"sdk_ref"`
	SDKLabel                string `yaml:"sdk_label"`
	SpecRef                 string `yaml:"spec_ref"`
	SpecLabel               string `yaml:"spec_label"`
	ZapGenerationAttributes string `yaml:"zap_generation_attributes"`
	AlchemyRef              string `yaml:"alchemy_ref"`
}

type ZAPDiff struct {
	BaselineXMLDir   string `name:"baseline-xml" help:"Path to baseline XML file or directory" required:"true"`
	GeneratedXMLDir  string `name:"generated-xml" help:"Path to generated XML file or directory" required:"true"`
	GeneratedSDKRoot string `name:"sdk-root" help:"Path to SDK root directory (for ZAP generation)" required:"true"`
	SDKLabel         string `name:"sdk-label" help:"Label for SDK" default:"SDK"`
	SpecLabel        string `name:"spec-label" help:"Label for Spec" default:"Spec"`
	GenAttributes    string `name:"gen-attributes" help:"Zap generation attributes"`
}

func (z *ZAPDiff) Run(cc *cli.Context) (err error) {
	if z.GenAttributes != "" {
		slog.Info("Running ZAP generation", "attributes", z.GenAttributes)
		zapCmd := &cli.ZAP{}
		zapCmd.Root = "."
		zapCmd.SdkRoot = z.GeneratedSDKRoot
		zapCmd.Attribute = []string{z.GenAttributes}
		zapCmd.FeatureXML = true
		zapCmd.ConformanceXML = true
		err = zapCmd.Run(cc)
		if err != nil {
			slog.Error("ZAP generation failed", "error", err)
			return err
		}
	}

	slog.Info("Running ZAPDiff", "xml1", z.BaselineXMLDir, "xml2", z.GeneratedXMLDir)
	diffCmd := &cli.ZAPDiff{}
	diffCmd.XmlRoot1 = z.BaselineXMLDir
	diffCmd.XmlRoot2 = z.GeneratedXMLDir
	diffCmd.Label1 = z.SDKLabel
	diffCmd.Label2 = z.SpecLabel
	diffCmd.MismatchLevel = 1

	err = diffCmd.Run(cc)
	if err != nil {
		slog.Error("ZAPDiff failed", "error", err)
		return err
	}

	slog.Info("zapdiff Action completed successfully")
	return nil
}
