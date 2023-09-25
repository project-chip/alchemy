package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/hasty/matterfmt/render"
	"github.com/urfave/cli/v2"
)

func Format(cxt context.Context, cCtx *cli.Context) error {
	files, err := getFilePaths(cCtx)
	if err != nil {
		return err
	}
	for i, f := range files {
		slog.Info("Formatting", "file", f, "index", (i + 1), "count", len(files))
		out, err := getOutputContext(cxt, f)
		if err != nil {
			return err
		}
		result, err := render.Render(cxt, out.Doc)
		if err != nil {
			return err
		}
		os.WriteFile(f, []byte(result), os.ModeAppend)
	}
	return nil
}
