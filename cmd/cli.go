//go:build !db && !github

package cmd

import (
	"os"

	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/dump"
)

var commands struct {
	Format      cli.Format      `cmd:"" help:"disco ball Matter spec documents specified by the filename_pattern" group:"Spec Commands:"`
	Disco       cli.Disco       `cmd:"" help:"disco ball Matter spec documents specified by the filename_pattern" group:"Spec Commands:"`
	ZAP         cli.ZAP         `cmd:"" help:"transmute the Matter spec into ZAP templates, optionally filtered to the files specified by filename_pattern" group:"SDK Commands:"`
	ZAPDiff     cli.ZAPDiff     `cmd:"" name:"zap-diff" help:"Compares two set of ZAP XMLs for any incosistency." group:"SDK Commands:"`
	Conformance cli.Conformance `cmd:"" help:"test conformance values"  group:"Spec Commands:"`
	Dump        dump.Command    `cmd:"" hidden:"" help:"dump the parse tree of Matter documents specified by filename_pattern"`
	DM          cli.DataModel   `cmd:"" help:"transmute the Matter spec into data model XML; optionally filtered to the files specified in filename_pattern" group:"SDK Commands:"`
	TestPlan    cli.TestPlan    `cmd:"" name:"test-plan" aliases:"testplan" help:"create an initial test plan from the spec, optionally filtered to the files specified in filename_pattern" group:"Testing Commands:"`
	TestScript  cli.TestScript  `cmd:"" name:"test-script" aliases:"testscript" help:"create shell python scripts from the spec, optionally filtered to the files specified by filename_pattern" group:"Testing Commands:"`
	Validate    cli.Validate    `cmd:"" help:"validate the Matter specification object model" group:"Spec Commands:"`
	Yaml2Python cli.Yaml2Python `cmd:"" name:"yaml-2-python" aliases:"yaml2python" help:"create a shell python script from a test YAML, optionally filtered to the files specified by filename_pattern"  group:"Testing Commands:"`
	Wordlist    cli.Wordlist    `cmd:"" hidden:"" name:"wordlist" help:"add words to wordlist.txt"`
	ErrDiff     cli.ErrDiff     `cmd:"" name:"err-diff" hidden:"" help:"Checks for new errors caused by a PR in review."`
	Version     Version         `cmd:"" hidden:"" name:"version" help:"display version number"`

	globalFlags `embed:""`
}

func init() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}
}
