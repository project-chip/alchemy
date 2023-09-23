package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/matterfmt/disco"
	"github.com/hasty/matterfmt/render"
	"github.com/urfave/cli/v2"
)

func DiscoBall(cxt context.Context, cCtx *cli.Context) error {
	files, err := getFilePaths(cCtx)
	if err != nil {
		return err
	}
	for i, f := range files {
		fmt.Fprintf(os.Stderr, "Disco-balling %s (%d of %d)...\n", f, (i + 1), len(files))
		out, err := getOutputContext(cxt, f)
		if err != nil {
			return err
		}
		err = disco.Ball(out.Doc)
		if err != nil {
			return err
		}
		os.WriteFile(f, []byte(render.Render(cxt, out.Doc)), os.ModeAppend)
	}
	return nil
}
