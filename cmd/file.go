package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
	"golang.org/x/sync/errgroup"
)

var bannedPaths map[string]string = map[string]string{
	"namespaces/Namespace-Common-Position.adoc":         "parser does not support nested tables",
	"service_device_management/PowerSourceCluster.adoc": "parser gets stuck parsing",
	"secure_channel/Discovery.adoc":                     "parser gets stuck parsing",
}

func getFilePaths(filepaths []string) ([]string, error) {
	filtered := make([]string, 0, len(filepaths))
	for _, filepath := range filepaths {
		paths, err := doublestar.FilepathGlob(filepath)
		if err != nil {
			return nil, err
		}
		for _, p := range paths {
			var banned bool
			for ban, reason := range bannedPaths {
				if strings.HasSuffix(p, ban) {
					fmt.Printf("Skipping excluded file %s; %s...\n", p, reason)
					banned = true
				}
			}
			if banned {
				continue
			}
			filtered = append(filtered, p)
		}

	}
	return filtered, nil
}

func getOutputContext(cxt context.Context, path string) (*output.Context, error) {
	doc, err := ascii.Open(path)
	if err != nil {
		return nil, err
	}

	return output.NewContext(cxt, doc), nil
}

func processFiles(cxt context.Context, filepaths []string, serial bool, dryRun bool, processor func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error)) error {
	files, err := getFilePaths(filepaths)
	if err != nil {
		return err
	}
	if serial {
		for i, file := range files {
			result, outPath, err := processor(cxt, file, i, len(files))
			if err != nil {
				return err
			}
			if !dryRun {
				err = os.WriteFile(outPath, []byte(result), os.ModeAppend)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	var complete int32
	g, errCxt := errgroup.WithContext(cxt)
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				done := atomic.AddInt32(&complete, 1)

				result, outPath, err := processor(errCxt, file, int(done), len(files))
				if err != nil {
					return err
				}
				if !dryRun {
					err = os.WriteFile(outPath, []byte(result), os.ModeAppend)
					if err != nil {
						return err
					}
				}
				return nil
			})
		}(f, i)

	}
	return g.Wait()
}
