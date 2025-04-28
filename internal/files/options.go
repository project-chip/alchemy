package files

type OutputOptions struct {
	DryRun bool `default:"false" short:"d" aliases:"dryrun" help:"whether or not to actually output files" group:"Output:"`
	Patch  bool `default:"false" short:"p" help:"generate patch file" group:"Output:"`
}
