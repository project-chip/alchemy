package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/render/zcl"
	"github.com/hasty/alchemy/zap"
	zaprender "github.com/hasty/alchemy/zap/render"
	"github.com/iancoleman/strcase"
)

type zclRenderer struct {
	processor
	asciiParser

	serial bool
	dryRun bool
}

var selfClosingTags = regexp.MustCompile("></[^>]+>")

func ZCL(cxt context.Context, specRoot string, zclRoot string, options ...Option) error {
	z := &zclRenderer{}
	for _, o := range options {
		err := o(z)
		if err != nil {
			return err
		}
	}
	return z.run(cxt, specRoot, zclRoot)
}

func (z *zclRenderer) run(cxt context.Context, specRoot string, zclRoot string) error {
	var appClusterPaths []string
	var appClusterIndexPaths []string
	filepath.WalkDir(specRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" {
			docType, e := ascii.GetDocType(path)
			if e != nil {
				return nil
			}
			switch docType {
			case matter.DocTypeAppCluster:
				appClusterPaths = append(appClusterPaths, path)
			case matter.DocTypeAppClusterIndex:
				appClusterIndexPaths = append(appClusterIndexPaths, path)
			}
		}
		return nil
	})

	domains := make(map[string]matter.Domain)

	err := z.processFiles(cxt, appClusterIndexPaths, func(cxt context.Context, file string, index, total int) error {
		fmt.Fprintf(os.Stderr, "ZCLing index %s (%d of %d)...\n", file, index, total)
		doc, err := ascii.Open(file, z.settings...)
		if err != nil {
			return err
		}

		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}

		domain := zap.StringToDomain(top.Name)
		fmt.Printf("Index Name: %s domain: %v\n", top.Name, domain)

		parse.Search[*types.Section](top.Base.Elements, func(t *types.Section) bool {
			link := parse.FindFirst[*types.InlineLink](t.Title)
			if link != nil {
				linkPath, ok := link.Location.Path.(string)
				if ok {
					linkPath = filepath.Join(filepath.Dir(file), linkPath)
					fmt.Printf("file link %s -> %s\n", file, linkPath)
					domains[linkPath] = domain
				}
			}
			return false
		})
		return nil
	})

	if err != nil {
		return err
	}

	outputs := make(map[string]*zcl.Result)
	var lock sync.Mutex
	err = z.processFiles(cxt, appClusterPaths, func(cxt context.Context, file string, index int, total int) error {
		doc, err := ascii.Open(file, z.settings...)
		if err != nil {
			return err
		}
		if domain, ok := domains[file]; ok {
			doc.Domain = domain
		} else {
			doc.Domain = matter.DomainCHIP
		}

		newFile := filepath.Base(file)
		newFile = zap.ZAPName(strings.TrimSuffix(newFile, filepath.Ext(file)))
		newFile = strcase.ToKebab(newFile)
		newPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip", newFile+".xml")

		existing, err := os.ReadFile(newPath)
		if errors.Is(err, os.ErrNotExist) {
			var result *zcl.Result
			result, err = zcl.Render(cxt, doc)
			if err != nil {
				err = fmt.Errorf("failed rendering %s: %w", file, err)
				return err
			}
			lock.Lock()
			outputs[newPath] = result
			lock.Unlock()
			fmt.Fprintf(os.Stderr, "ZCL'd %s to %s (%d of %d)...\n", file, newPath, index, total)

		} else if err != nil {
			return err
		} else {
			fmt.Fprintf(os.Stderr, "ZCLing existing %s to %s (%d of %d)...\n", file, newPath, index, total)
			var buf bytes.Buffer
			models, err := doc.ToModel()
			if err != nil {
				return err
			}
			err = zaprender.Render(doc, bytes.NewReader(existing), &buf, models)
			if err != nil {
				return err
			}
			out := selfClosingTags.ReplaceAllString(buf.String(), "/>") // Lame limitation of Go's XML encoder
			lock.Lock()
			outputs[newPath] = &zcl.Result{ZCL: out, Doc: doc, Models: models}
			lock.Unlock()
			fmt.Fprintf(os.Stderr, "ZCL'd existing %s to %s (%d of %d)...\n", file, newPath, index, total)
		}

		return nil
	})
	if err != nil {
		return err
	}
	for path, result := range outputs {
		if len(result.Models) == 0 {
			continue
		}

		if !z.dryRun {
			err = os.WriteFile(path, []byte(result.ZCL), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
