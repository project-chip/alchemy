package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/render/zcl"
)

func ZCL(cxt context.Context, filepaths []string, dryRun bool, serial bool) error {
	return processFiles(cxt, filepaths, serial, dryRun, func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error) {
		outPath = path.Base(file) + ".xml"
		doc, err := ascii.Open(file)
		if err != nil {
			return
		}
		result, err = zcl.Render(cxt, doc)
		if err != nil {
			return
		}

		fmt.Fprintf(os.Stderr, "ZCL'd %s (%d of %d)...\n", file, index, total)
		fmt.Print(result)
		return
	})

}
