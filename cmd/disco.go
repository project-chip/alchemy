package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/matterfmt/disco"
	"github.com/hasty/matterfmt/render"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func DiscoBall(cxt context.Context, cCtx *cli.Context) error {

	files, err := getFilePaths(cCtx)
	if err != nil {
		return err
	}
	g, errCxt := errgroup.WithContext(cxt)
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				fmt.Fprintf(os.Stderr, "Disco-balling %s (%d of %d)...\n", file, (index + 1), len(files))
				out, err := getOutputContext(errCxt, file)
				if err != nil {
					return err
				}
				err = disco.Ball(disco.NewContext(errCxt), out.Doc)
				if err != nil {
					return fmt.Errorf("error disco balling %s: %w", file, err)
				}
				result, err := render.Render(errCxt, out.Doc)
				if err != nil {
					return err
				}
				return os.WriteFile(file, []byte(result), os.ModeAppend)
			})
		}(f, i)

	}
	return g.Wait()
}
