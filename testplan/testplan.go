package testplan

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/iancoleman/strcase"
)

func renderTestPlans(cxt context.Context, spec *matter.Spec, docs map[string]*ascii.Doc, targetDocs []*ascii.Doc, testPlanRoot string, filesOptions files.Options, overwrite bool) (outputs map[string]string, err error) {
	var lock sync.Mutex
	outputs = make(map[string]string)
	slog.InfoContext(cxt, "Rendering test plans...")

	for len(targetDocs) > 0 {

		err = files.ProcessDocs(cxt, targetDocs, func(cxt context.Context, doc *ascii.Doc, index, total int) error {

			path := doc.Path

			entities, err := doc.Entities()
			if err != nil {
				return err
			}

			destinations := buildDestinations(testPlanRoot, doc, entities)

			for newPath, cluster := range destinations {

				_, err := os.ReadFile(newPath)
				if (err == nil || !errors.Is(err, os.ErrNotExist)) && !overwrite {
					slog.InfoContext(cxt, "Skipping existing test plan", slog.String("path", newPath))
					continue
				}
				if filesOptions.Serial {
					slog.InfoContext(cxt, "Rendering test plan", "from", path, "to", newPath, "index", index, "count", total)
				}

				var result string
				result, err = renderClusterTestPlan(doc, cluster)
				if err != nil {
					return fmt.Errorf("failed rendering %s: %w", path, err)
				}

				lock.Lock()
				outputs[newPath] = result
				lock.Unlock()
			}
			return nil
		}, filesOptions)
		if err != nil {
			return
		}

		targetDocs = nil

	}
	return
}

func buildDestinations(sdkRoot string, doc *ascii.Doc, entities []types.Entity) (destinations map[string]*matter.Cluster) {
	destinations = make(map[string]*matter.Cluster)

	for _, e := range entities {
		switch e := e.(type) {
		case *matter.Cluster:
			fileName := strings.ToLower(strcase.ToSnake(e.Name))
			newPath := getTestPlanPath(sdkRoot, fileName)
			destinations[newPath] = e
		}
	}
	return

}
