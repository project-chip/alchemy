package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/hasty/matterfmt/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	logrus.SetLevel(logrus.ErrorLevel)

	cxt := context.Background()

	var dryRun bool
	var serial bool

	var linkAttributes bool

	app := &cli.App{
		Name:  "matterfmt",
		Usage: "builds stuff",
		Action: func(c *cli.Context) error {
			return cmd.Format(cxt, c.Args().Slice(), dryRun, serial)
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "dryrun",
				Aliases:     []string{"dry"},
				Usage:       "whether or not to actually output files",
				Destination: &dryRun,
			},
			&cli.BoolFlag{
				Name:        "serial",
				Usage:       "process files one-by-one",
				Destination: &serial,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "disco",
				Aliases: []string{"c"},
				Usage:   "Discoball documents",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "linkAttributes",
						Usage:       "whether or not to actually output files",
						Destination: &linkAttributes,
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := cmd.DiscoBall(cxt, cCtx.Args().Slice(), dryRun, serial, linkAttributes)
					if err != nil {
						return cli.Exit(err, -1)
					}
					return nil
				},
			},
			{
				Name:    "format",
				Aliases: []string{"a"},
				Usage:   "just format Matter documents",
				Action: func(cCtx *cli.Context) error {
					return cmd.Format(cxt, cCtx.Args().Slice(), dryRun, serial)
				},
			},
			{
				Name:    "dump",
				Aliases: []string{"c"},
				Usage:   "dump the parse tree of Matter documents",
				Action: func(cCtx *cli.Context) error {
					return cmd.Dump(cxt, cCtx.Args().Slice())
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("failed running", "error", err)
	}
}
