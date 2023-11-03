package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hasty/matterfmt/zap"
)

type zclConparer struct {
	processor
	asciiParser

	serial bool
	dryRun bool
}

func Compare(cxt context.Context, specRoot string, zclRoot string, options ...Option) error {
	z := &zclConparer{}
	for _, o := range options {
		err := o(z)
		if err != nil {
			return err
		}
	}
	return z.run(cxt, specRoot, zclRoot)
}

func (z *zclConparer) run(cxt context.Context, specRoot string, zclRoot string) error {

	zapPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip/*.xml")
	xmlFiles, err := filepath.Glob(zapPath)
	if err != nil {
		return err
	}
	for _, f := range xmlFiles {
		fmt.Printf("ZAP file: %s\n", f)
		/*_, err = zap.Read(f)
		if err != nil {
			return err
		}*/
		file, err := os.Open(f)
		if err != nil {
			return err
		}
		defer file.Close()
		var models []any
		models, err = zap.Parse(file)
		if err != nil {
			return err
		}
		encoder := json.NewEncoder(os.Stdout)
		//encoder.SetIndent("", "\t")
		err = encoder.Encode(models)
		if err != nil {
			return err
		}
	}
	/*var appClusterPaths []string
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

		domain := getDomain(top.Name)
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
		var result *zcl.Result
		result, err = zcl.Render(cxt, doc)
		if err != nil {
			err = fmt.Errorf("failed rendering %s: %w", file, err)
			return err
		}

		fmt.Fprintf(os.Stderr, "ZCL'd %s (%d of %d)...\n", file, index, total)
		lock.Lock()
		outputs[file] = result
		lock.Unlock()
		return nil
	})
	if err != nil {
		return err
	}
	for path, result := range outputs {
		if len(result.Models) == 0 {
			continue
		}

		newFile := filepath.Base(path)
		newFile = getZCLName(strings.TrimSuffix(newFile, filepath.Ext(path)))
		newFile = strcase.ToKebab(newFile)

		newPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip", newFile+".xml")
		if !z.dryRun {
			err = os.WriteFile(newPath, []byte(result.ZCL), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}
	}*/
	return nil
}
