package zap

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
	"github.com/hasty/alchemy/zap/amend"
	"github.com/hasty/alchemy/zap/render"
)

type concurrentMap struct {
	Map map[string]any
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

func renderAppClusterTemplates(cxt context.Context, appClusters []*ascii.Doc, zclRoot string, filesOptions files.Options, overwrite bool) (outputs map[string]*render.Result, provisionalZclFiles []string, err error) {
	var lock sync.Mutex
	outputs = make(map[string]*render.Result)
	slog.InfoContext(cxt, "Rendering ZAP templates...")

	missingStructs := &concurrentMap{Map: make(map[string]any)}
	err = files.ProcessDocs(cxt, appClusters, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		path := doc.Path
		newPath := getZapPath(zclRoot, path)

		models, err := doc.ToModel()
		if err != nil {
			return err
		}

		patchMissingTypes(doc, models, missingStructs)

		doc.Domain = getDocDomain(doc)
		existing, err := os.ReadFile(newPath)
		if errors.Is(err, os.ErrNotExist) || overwrite {
			if filesOptions.Serial {
				slog.InfoContext(cxt, "Rendering new ZAP template", "from", path, "to", newPath, "index", index, "count", total)
			}
			var result *render.Result
			result, err = render.Render(cxt, doc, models)
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

			var clusters []any
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
			err = amend.Render(doc, bytes.NewReader(existing), &buf, clusters)
			if err != nil {
				return err
			}
			out := selfClosingTags.ReplaceAllString(buf.String(), "/>") // Lame limitation of Go's XML encoder
			lock.Lock()
			outputs[newPath] = &render.Result{ZCL: out, Doc: doc, Models: models}
			lock.Unlock()
			if filesOptions.Serial {
				slog.InfoContext(cxt, "Rendered existing ZAP template", "from", path, "to", newPath, "index", index, "count", total)
			}
		}
		return nil
	}, filesOptions)
	return
}
