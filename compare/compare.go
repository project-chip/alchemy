package compare

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/constraint"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
	zparse "github.com/hasty/alchemy/zap/parse"
	"github.com/iancoleman/strcase"
)

func Compare(cxt context.Context, specRoot string, zclRoot string, settings []configuration.Setting) error {

	slog.InfoContext(cxt, "Loading spec...")
	spec, _, err := files.LoadSpec(cxt, specRoot, files.Options{}, settings)
	if err != nil {
		return err
	}

	zapEntities, err := loadZAPEntities(zclRoot)

	if err != nil {
		return err
	}

	specEntities, err := loadSpecClusterPaths(spec, zclRoot)
	if err != nil {
		return err
	}
	diffs, err := compareEntities(specEntities, zapEntities)
	if err != nil {
		return err
	}
	jm := json.NewEncoder(os.Stdout)
	jm.SetIndent("", "\t")
	return jm.Encode(diffs)
}

func loadSpecClusterPaths(spec *matter.Spec, zclRoot string) (map[string][]mattertypes.Entity, error) {
	specEntities := make(map[string][]mattertypes.Entity)
	specPaths := make(map[string][]mattertypes.Entity)
	for e, path := range spec.DocRefs {
		switch e.(type) {
		case *matter.Cluster:
			specPaths[path] = append(specPaths[path], e)
		}
	}
	for path, entities := range specPaths {
		newFile := filepath.Base(path)
		errata, ok := zap.Erratas[newFile]
		if !ok {
			errata = zap.DefaultErrata
		}

		newFile = zap.ZAPClusterName(path, errata, entities)
		newFile = strcase.ToKebab(newFile)
		newPath := filepath.Join(zclRoot, "src/app/zap-templates/zcl/data-model/chip", newFile+".xml")
		specEntities[newPath] = entities
	}
	return specEntities, nil
}

func loadSpecEntities(appClusterPaths []string, settings []configuration.Setting, domains map[string]matter.Domain, zclRoot string) (map[string][]mattertypes.Entity, error) {
	specEntities := make(map[string][]mattertypes.Entity)
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

		entities, err := doc.Entities()
		if err != nil {
			return nil, err
		}

		errata, ok := zap.Erratas[filepath.Base(file)]
		if !ok {
			errata = zap.DefaultErrata
		}

		fmt.Fprintf(os.Stderr, "ZCL'd %s (%d of %d)...\n", file, i+1, len(appClusterPaths))

		newFile := zap.ZAPClusterName(doc.Path, errata, entities)
		newFile = strcase.ToKebab(newFile)

		newPath := filepath.Join(zclRoot, "src/app/zap-templates/zcl/data-model/chip", newFile+".xml")

		specEntities[newPath] = entities
	}
	return specEntities, nil
}

func getAppDomains(specRoot string, settings []configuration.Setting) ([]string, map[string]matter.Domain, error) {
	var appClusterPaths []string
	var appClusterIndexPaths []string
	err := filepath.WalkDir(specRoot, func(path string, d fs.DirEntry, err error) error {
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
	if err != nil {
		return nil, nil, err
	}

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

func loadZAPEntities(zclRoot string) (map[string][]mattertypes.Entity, error) {
	zapPath := filepath.Join(zclRoot, "src/app/zap-templates/zcl/data-model/chip/*.xml")
	xmlFiles, err := filepath.Glob(zapPath)
	if err != nil {
		return nil, err
	}
	zapEntities := make(map[string][]mattertypes.Entity)
	for _, f := range xmlFiles {
		slog.Debug("ZAP file", slog.String("path", f))

		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		var entities []mattertypes.Entity
		entities, err = zparse.ZAP(file)
		if err != nil {
			return nil, err
		}

		zapEntities[f] = entities
	}
	return zapEntities, nil
}

func compareConformance(spec conformance.Set, zap conformance.Set) (diffs []any) {
	if len(spec) == 0 {
		if len(zap) > 0 {
			diffs = append(diffs, newMissingDiff("", DiffPropertyConformance, SourceSpec))
		}
		return
	} else if len(zap) == 0 {
		diffs = append(diffs, newMissingDiff("", DiffPropertyConformance, SourceZAP))
		return
	}
	specMandatory := conformance.IsMandatory(spec)
	zapMandatory := conformance.IsMandatory(zap)
	if specMandatory != zapMandatory {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyConformance, Spec: spec.String(), ZAP: zap.String()})
	}

	return
}
func compareConstraint(spec constraint.Constraint, zap constraint.Constraint) (diffs []any) {
	if spec == nil {
		if zap == nil {
			return
		}
	} else if zap != nil {
		return
	}
	_, ok := spec.(*constraint.AllConstraint)
	if ok {
		return
	}
	return
}

func namesEqual(specName string, zapName string) bool {
	if strings.EqualFold(specName, zapName) {
		return true
	}
	specName = strcase.ToCamel(specName)
	zapName = strcase.ToCamel(zapName)
	return strings.EqualFold(specName, zapName)
}
