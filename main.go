package main

import (
	"context"
	"os"

	"github.com/hasty/matterfmt/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	logrus.SetLevel(logrus.ErrorLevel)

	cxt := context.Background()

	app := &cli.App{
		Name:  "matterfmt",
		Usage: "builds stuff",
		Action: func(c *cli.Context) error {
			return cmd.DiscoBall(cxt, c)
		},
		Commands: []*cli.Command{
			{
				Name:    "disco",
				Aliases: []string{"c"},
				Usage:   "discoball documents",
				Action: func(cCtx *cli.Context) error {
					return cmd.DiscoBall(cxt, cCtx)
				},
			},
			{
				Name:    "format",
				Aliases: []string{"a"},
				Usage:   "just format a Matter document",
				Action: func(cCtx *cli.Context) error {
					return cmd.Format(cxt, cCtx)
				},
			},
			{
				Name:    "dump",
				Aliases: []string{"c"},
				Usage:   "dump the parse tree of a document",
				Action: func(cCtx *cli.Context) error {
					return cmd.Dump(cxt, cCtx)
				},
			},
		},
	}

	(app).Run(os.Args)

}
