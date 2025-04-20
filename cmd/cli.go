//go:build !db && !github

package cmd

import (
	"github.com/project-chip/alchemy/cmd/compare"
	"github.com/project-chip/alchemy/cmd/disco"
	"github.com/project-chip/alchemy/cmd/dm"
	"github.com/project-chip/alchemy/cmd/dump"
	"github.com/project-chip/alchemy/cmd/format"
	"github.com/project-chip/alchemy/cmd/testplan"
	"github.com/project-chip/alchemy/cmd/testscript"
	"github.com/project-chip/alchemy/cmd/validate"
	"github.com/project-chip/alchemy/cmd/yaml2python"
	"github.com/project-chip/alchemy/cmd/zap"
)

var commands struct {
	Format      format.Command      `cmd:"" help:"disco ball Matter spec documents specified by the filename_pattern" group:"Spec Commands:"`
	Disco       disco.Command       `cmd:"" help:"disco ball Matter spec documents specified by the filename_pattern" group:"Spec Commands:"`
	ZAP         zap.Command         `cmd:"" help:"transmute the Matter spec into ZAP templates, optionally filtered to the files specified by filename_pattern" group:"SDK Commands:"`
	Compare     compare.Command     `cmd:"" help:"compare the spec to zap-templates and output a JSON diff"  group:"SDK Commands:"`
	Conformance Conformance         `cmd:"" help:"test conformance values"  group:"Spec Commands:"`
	Dump        dump.Command        `cmd:"" hidden:"" help:"dump the parse tree of Matter documents specified by filename_pattern"`
	DM          dm.Command          `cmd:"" help:"transmute the Matter spec into data model XML; optionally filtered to the files specified in filename_pattern" group:"SDK Commands:"`
	TestPlan    testplan.Command    `cmd:"" name:"test-plan" aliases:"testplan" help:"create an initial test plan from the spec, optionally filtered to the files specified in filename_pattern" group:"Testing Commands:"`
	TestScript  testscript.Command  `cmd:"" name:"test-script" aliases:"testscript" help:"create shell python scripts from the spec, optionally filtered to the files specified by filename_pattern" group:"Testing Commands:"`
	Validate    validate.Command    `cmd:"" help:"validate the Matter specification object model" group:"Spec Commands:"`
	Yaml2Python yaml2python.Command `cmd:"" name:"yaml-2-python" aliases:"yaml2python" help:"create a shell python script from a test YAML, optionally filtered to the files specified by filename_pattern"  group:"Testing Commands:"`
	Version     Version             `cmd:"" hidden:"" name:"version" help:"display version number"`

	globalFlags `embed:""`
}
