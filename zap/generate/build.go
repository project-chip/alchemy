package generate

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

type concurrentMap[K comparable, T any] struct {
	Map map[K]T
	sync.Mutex
}

func newConcurrentMap[K comparable, T any]() *concurrentMap[K, T] {
	return &concurrentMap[K, T]{Map: make(map[K]T)}
}

func getDocDomain(doc *ascii.Doc) matter.Domain {
	if doc.Domain != matter.DomainUnknown {
		return doc.Domain
	}
	for _, p := range doc.Parents() {
		d := getDocDomain(p)
		if d != matter.DomainUnknown {
			return d
		}
	}
	return matter.DomainUnknown
}

func renderClusterTemplates(cxt context.Context, spec *matter.Spec, docs map[string]*ascii.Doc, targetDocs []*ascii.Doc, sdkRoot string, filesOptions files.Options, overwrite bool) (outputs map[string]string, provisionalZclFiles []string, err error) {
	var lock sync.Mutex
	outputs = make(map[string]string)
	slog.InfoContext(cxt, "Rendering ZAP templates...")

	dependencies := newConcurrentMap[string, bool]()

	for len(targetDocs) > 0 {

		err = files.ProcessDocs(cxt, targetDocs, func(cxt context.Context, doc *ascii.Doc, index, total int) error {

			path := doc.Path

			errata, ok := zap.Erratas[filepath.Base(path)]
			if !ok {
				errata = zap.DefaultErrata
			}

			entities, err := doc.Entities()
			if err != nil {
				return err
			}

			destinations := buildDestinations(sdkRoot, doc, entities, errata)

			dependencies.Lock()
			dependencies.Map[path] = true
			dependencies.Unlock()

			for newPath, entities := range destinations {

				if len(entities) == 0 {
					continue
				}

				findDependencies(spec, entities, dependencies)

				doc.Domain = getDocDomain(doc)

				if doc.Domain == matter.DomainUnknown {
					if errata.Domain != matter.DomainUnknown {
						doc.Domain = errata.Domain
					} else {
						doc.Domain = matter.DomainGeneral
					}
				}

				if len(entities) == 0 {
					slog.WarnContext(cxt, "Skipped spec file with no entities", "from", path, "to", newPath, "index", index, "count", total)
					return err
				}

				configurator, err := zap.NewConfigurator(spec, doc, entities)
				if err != nil {
					return err
				}

				var result string

				var doc *etree.Document
				var provisional bool

				existing, err := os.ReadFile(newPath)
				if errors.Is(err, os.ErrNotExist) || overwrite {
					if filesOptions.Serial {
						slog.InfoContext(cxt, "Rendering new ZAP template", "from", path, "to", newPath, "index", index, "count", total)
					}
					provisional = true
					doc = newZapTemplate()
				} else if err != nil {
					return err
				} else {
					if filesOptions.Serial {
						slog.InfoContext(cxt, "Rendering existing ZAP template", "from", path, "to", newPath, "index", index, "count", total)
					}
					doc = etree.NewDocument()
					err = doc.ReadFromBytes(existing)
					if err != nil {
						return fmt.Errorf("failed reading ZAP template %s: %w", path, err)
					}

				}
				result, err = renderZapTemplate(configurator, doc, errata)
				if err != nil {
					return fmt.Errorf("failed rendering %s: %w", path, err)
				}
				lock.Lock()
				outputs[newPath] = result
				if provisional {
					provisionalZclFiles = append(provisionalZclFiles, filepath.Base(newPath))

				}
				lock.Unlock()
			}
			return nil
		}, filesOptions)
		if err != nil {
			return
		}

		targetDocs = nil

		if len(dependencies.Map) > 0 {
			for dep, handled := range dependencies.Map {
				if handled {
					continue
				}
				targetDoc, ok := docs[dep]
				if ok {
					slog.Info("Adding dependent doc to render list", "path", dep)
					targetDocs = append(targetDocs, targetDoc)
				} else {
					slog.Warn("unknown dependency path", "path", dep)
				}
			}
		}
	}
	return
}

func buildDestinations(sdkRoot string, doc *ascii.Doc, entities []types.Entity, errata *zap.Errata) (destinations map[string][]types.Entity) {
	destinations = make(map[string][]types.Entity)
	if len(errata.ClusterSplit) == 0 {
		newFile := zap.ZAPClusterName(doc.Path, errata, entities)
		newPath := getZapPath(sdkRoot, newFile)
		destinations[newPath] = entities
		return
	}
	for clusterID, path := range errata.ClusterSplit {
		cid := matter.ParseNumber(clusterID)
		if !cid.Valid() {
			slog.Warn("invalid cluster split ID", "id", clusterID)
			continue
		}
		var clusters []types.Entity
		for _, m := range entities {
			switch m := m.(type) {
			case *matter.Cluster:
				if m.ID.Equals(cid) {
					clusters = append(clusters, m)
				}
			}
		}
		destinations[getZapPath(sdkRoot, path)] = clusters
	}
	return
}
