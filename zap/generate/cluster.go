package generate

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
	"github.com/hasty/alchemy/zap/amend"
	"github.com/hasty/alchemy/zap/render"
)

type concurrentMap[T any] struct {
	Map map[string]T
	sync.Mutex
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

func renderClusterTemplates(cxt context.Context, spec *matter.Spec, docs map[string]*ascii.Doc, targetDocs []*ascii.Doc, zclRoot string, filesOptions files.Options, overwrite bool) (outputs map[string]*render.Result, provisionalZclFiles []string, err error) {
	var lock sync.Mutex
	outputs = make(map[string]*render.Result)
	slog.InfoContext(cxt, "Rendering ZAP templates...")

	dependencies := &concurrentMap[bool]{Map: make(map[string]bool)}

	for len(targetDocs) > 0 {

		err = files.ProcessDocs(cxt, targetDocs, func(cxt context.Context, doc *ascii.Doc, index, total int) error {

			path := doc.Path

			errata, ok := zap.Erratas[filepath.Base(path)]
			if !ok {
				errata = zap.DefaultErrata
			}

			models, err := doc.ToModel()
			if err != nil {
				return err
			}

			destinations := buildDestinations(zclRoot, doc, models, errata)

			dependencies.Lock()
			dependencies.Map[path] = true
			dependencies.Unlock()

			for newPath, models := range destinations {

				if len(models) == 0 {
					continue
				}

				findDependencies(spec, models, dependencies)

				doc.Domain = getDocDomain(doc)

				if doc.Domain == matter.DomainUnknown {
					if errata.Domain != matter.DomainUnknown {
						doc.Domain = errata.Domain
					} else {
						doc.Domain = matter.DomainGeneral
					}
				}

				existing, err := os.ReadFile(newPath)
				if errors.Is(err, os.ErrNotExist) || overwrite {
					if filesOptions.Serial {
						slog.InfoContext(cxt, "Rendering new ZAP template", "from", path, "to", newPath, "index", index, "count", total)
					}
					var result *render.Result
					result, err = render.Render(cxt, spec, doc, models, errata)
					if err != nil {
						err = fmt.Errorf("failed rendering %s: %w", path, err)
						return err
					}
					lock.Lock()
					outputs[newPath] = result
					provisionalZclFiles = append(provisionalZclFiles, filepath.Base(newPath))
					lock.Unlock()
					if filesOptions.Serial {
						slog.InfoContext(cxt, "Rendered new ZAP template", "from", path, "to", newPath, "index", index, "count", total)
					}
				} else if err != nil {
					return err
				} else {
					if filesOptions.Serial {
						slog.InfoContext(cxt, "Rendering existing ZAP template", "from", path, "to", newPath, "index", index, "count", total)
					}
					var buf bytes.Buffer

					var clusters []matter.Model
					for _, m := range models {
						switch m := m.(type) {
						case *matter.Cluster:
							if m.Hierarchy == "Base" && m.Name != "Mode Base" && m.Name != "Scenes" && m.ID.Valid() {
								clusters = append(clusters, m)
							}
						}
					}
					if len(clusters) == 0 {
						slog.DebugContext(cxt, "Skipped spec file with no clusters", "from", path, "to", newPath, "index", index, "count", total)
						return err
					}
					err = amend.Render(spec, doc, bytes.NewReader(existing), &buf, clusters, errata)
					if err != nil {
						return fmt.Errorf("failed rendering %s: %w", path, err)
					}
					out := selfClosingTags.ReplaceAllString(buf.String(), "/>") // Lame limitation of Go's XML encoder
					lock.Lock()
					outputs[newPath] = &render.Result{ZCL: out, Doc: doc, Models: models}
					lock.Unlock()
					if filesOptions.Serial {
						slog.InfoContext(cxt, "Rendered existing ZAP template", "from", path, "to", newPath, "index", index, "count", total)
					}
				}
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

func buildDestinations(zclRoot string, doc *ascii.Doc, models []matter.Model, errata *zap.Errata) (destinations map[string][]matter.Model) {
	destinations = make(map[string][]matter.Model)
	if len(errata.ClusterSplit) == 0 {
		newFile := zap.ZAPName(doc, errata, models)
		newPath := getZapPath(zclRoot, newFile)
		destinations[newPath] = models
		return
	}
	for clusterID, path := range errata.ClusterSplit {
		cid := matter.ParseID(clusterID)
		if !cid.Valid() {
			slog.Warn("invalid cluster split ID", "id", clusterID)
			continue
		}
		var clusters []matter.Model
		for _, m := range models {
			switch m := m.(type) {
			case *matter.Cluster:
				if m.ID.Equals(cid) {
					clusters = append(clusters, m)
				}
			}
		}
		destinations[getZapPath(zclRoot, path)] = clusters
	}
	return
}
