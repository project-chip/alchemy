package main

import (
	"context"
	"fmt"
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
			err := cmd.DiscoBall(cxt, c)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			return err
		},
		Commands: []*cli.Command{
			{
				Name:    "disco",
				Aliases: []string{"c"},
				Usage:   "Discoball documents",
				Action: func(cCtx *cli.Context) error {
					return cmd.DiscoBall(cxt, cCtx)
				},
			},
			{
				Name:    "format",
				Aliases: []string{"a"},
				Usage:   "just format Matter documents",
				Action: func(cCtx *cli.Context) error {
					return cmd.Format(cxt, cCtx)
				},
			},
			{
				Name:    "dump",
				Aliases: []string{"c"},
				Usage:   "dump the parse tree of Matter documents",
				Action: func(cCtx *cli.Context) error {
					return cmd.Dump(cxt, cCtx)
				},
			},
		},
	}

	(app).Run(os.Args)

}
