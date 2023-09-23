package cmd

import (
	"context"
	"fmt"
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
		fmt.Printf("Formatting %s (%d of %d)...", f, (i + 1), len(files))
		out, err := getOutputContext(cxt, f)
		if err != nil {
			return err
		}
		os.WriteFile(f, []byte(render.Render(cxt, out.Doc)), os.ModeAppend)
	}
	return nil
}
