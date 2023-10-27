package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/hasty/matterfmt/cmd"
	"github.com/hasty/matterfmt/disco"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	logrus.SetLevel(logrus.ErrorLevel)

	cxt := context.Background()

	var dryRun bool
	var serial bool

	var linkAttributes bool
	var dumpAscii bool

	var specRoot string
	var zclRoot string

	var attributes cli.StringSlice

	formatAction := func(cCtx *cli.Context) error {
		options := []cmd.Option{
			cmd.Serial(serial),
			cmd.DryRun(dryRun),
		}
		return cmd.Format(cxt, cCtx.Args().Slice(), options...)
	}

	app := &cli.App{
		Name:   "alchemy",
		Usage:  "builds stuff",
		Action: formatAction,
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
			&cli.StringSliceFlag{
				Name:        "attribute",
				Usage:       "attribute for preprocessing",
				Destination: &attributes,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "format",
				Aliases: []string{"fmt"},
				Usage:   "just format Matter documents",
				Action:  formatAction,
			},
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
					options := []cmd.Option{
						cmd.Serial(serial),
						cmd.DryRun(dryRun),
						cmd.Disco(disco.LinkAttributes(linkAttributes)),
					}

					err := cmd.DiscoBall(cxt, cCtx.Args().Slice(), options...)
					if err != nil {
						return cli.Exit(err, -1)
					}
					return nil
				},
			},

			{
				Name:  "zcl",
				Usage: "translate Matter spec to ZCL",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "specRoot",
						Usage:       "the src root of the spec repo",
						Destination: &specRoot,
					},
					&cli.StringFlag{
						Name:        "zclRoot",
						Usage:       "the zcl root of the connected-ip repo",
						Destination: &zclRoot,
					},
				},
				Action: func(cCtx *cli.Context) error {
					options := []cmd.Option{
						cmd.Serial(serial),
						cmd.DryRun(dryRun),
						cmd.AsciiAttributes(attributes.Value()),
					}
					return cmd.ZCL(cxt, specRoot, zclRoot, options...)
				},
			},
			{
				Name:  "db",
				Usage: "just format Matter documents",
				Action: func(cCtx *cli.Context) error {
					options := []cmd.Option{
						cmd.Serial(serial),
						cmd.AsciiAttributes(attributes.Value()),
					}
					return cmd.Database(cxt, cCtx.Args().Slice(), options...)
				},
			},
			{
				Name:    "dump",
				Aliases: []string{"c"},
				Usage:   "dump the parse tree of Matter documents",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "ascii",
						Usage:       "dump asciidoc object model",
						Destination: &dumpAscii,
					},
				},
				Action: func(cCtx *cli.Context) error {
					options := []cmd.Option{
						cmd.DumpAscii(dumpAscii),
						cmd.AsciiAttributes(attributes.Value()),
					}
					return cmd.Dump(cxt, cCtx.Args().Slice(), options...)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("failed running", "error", err)
	}
}
