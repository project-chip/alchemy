package compare

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
	zparse "github.com/hasty/alchemy/zap/parse"
	"github.com/iancoleman/strcase"
)

func Compare(cxt context.Context, specRoot string, zclRoot string, settings []configuration.Setting) error {

	zapModels, err := loadZAPModels(zclRoot)

	if err != nil {
		return err
	}

	appClusterPaths, domains, err := getAppDomains(specRoot, settings)

	if err != nil {
		return err
	}

	specModels, err := loadSpecModels(appClusterPaths, settings, domains, zclRoot)
	if err != nil {
		return err
	}
	diffs, err := compareModels(specModels, zapModels)
	if err != nil {
		return err
	}
	jm := json.NewEncoder(os.Stdout)
	jm.SetIndent("", "\t")
	jm.Encode(diffs)
	return nil
}

func loadSpecModels(appClusterPaths []string, settings []configuration.Setting, domains map[string]matter.Domain, zclRoot string) (map[string][]matter.Model, error) {
	specModels := make(map[string][]matter.Model)
	for i, file := range appClusterPaths {
		doc, err := ascii.OpenFile(file, append(ascii.GithubSettings(), settings...)...)
		if err != nil {
			return nil, err
		}
		if domain, ok := domains[file]; ok {
			doc.Domain = domain
		} else {
			doc.Domain = matter.DomainCHIP
		}

		models, err := doc.ToModel()
		if err != nil {
			return nil, err
		}

		errata, ok := zap.Erratas[filepath.Base(file)]
		if !ok {
			errata = zap.DefaultErrata
		}

		fmt.Fprintf(os.Stderr, "ZCL'd %s (%d of %d)...\n", file, i+1, len(appClusterPaths))

		newFile := filepath.Base(file)
		newFile = zap.ZAPName(doc, errata, models)
		newFile = strcase.ToKebab(newFile)

		newPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip", newFile+".xml")

		specModels[newPath] = models
	}
	return specModels, nil
}

func getAppDomains(specRoot string, settings []configuration.Setting) ([]string, map[string]matter.Domain, error) {
	var appClusterPaths []string
	var appClusterIndexPaths []string
	filepath.WalkDir(specRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" {
			docType, e := ascii.GetDocType(path)
			if e != nil {
				return e
			}
			switch docType {
			case matter.DocTypeCluster:
				appClusterPaths = append(appClusterPaths, path)
			case matter.DocTypeAppClusterIndex:
				appClusterIndexPaths = append(appClusterIndexPaths, path)
			}
		}
		return nil
	})

	domains := make(map[string]matter.Domain)

	for i, file := range appClusterIndexPaths {
		fmt.Fprintf(os.Stderr, "ZCLing index %s (%d of %d)...\n", file, i+1, len(appClusterIndexPaths))
		doc, err := ascii.OpenFile(file, settings...)
		if err != nil {
			return nil, nil, err
		}

		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil, nil, err
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
	}
	return appClusterPaths, domains, nil
}

func loadZAPModels(zclRoot string) (map[string][]matter.Model, error) {
	zapPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip/*.xml")
	xmlFiles, err := filepath.Glob(zapPath)
	if err != nil {
		return nil, err
	}
	zapModels := make(map[string][]matter.Model)
	for _, f := range xmlFiles {
		fmt.Printf("ZAP file: %s\n", f)

		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		var models []matter.Model
		models, err = zparse.ZAP(file)
		if err != nil {
			return nil, err
		}

		zapModels[f] = models
	}
	return zapModels, nil
}
