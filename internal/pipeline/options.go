package pipeline

type ProcessingOptions struct {
	Serial     bool `default:"false" help:"Process files one-by-one" group:"Processing:"`
	NoProgress bool `default:"false" name:"hide-progress-bar" help:"Hide the progress bar" group:"Processing:"`
}
