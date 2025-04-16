package files

type OutputOptions struct {
	DryRun bool `default:"false" aliases:"d,dryrun" help:"whether or not to actually output files" group:"Output:"`
	Patch  bool `default:"false" aliases:"p" help:"generate patch file" group:"Output:"`
}
