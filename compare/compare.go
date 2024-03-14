package compare

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/zap"
)

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
		doc, err := ascii.ParseFile(file, settings...)
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
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyConformance, Spec: spec.AsciiDocString(), ZAP: zap.AsciiDocString()})
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
	specName = matter.Case(specName)
	zapName = matter.Case(zapName)
	return strings.EqualFold(specName, zapName)
}
