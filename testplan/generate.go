package testplan

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
	"github.com/iancoleman/strcase"
)

type Options struct {
	Files     files.Options
	Ascii     []configuration.Setting
	Overwrite bool
}

func Generate(cxt context.Context, specRoot string, testPlanRoot string, paths []string, options Options) error {

	slog.InfoContext(cxt, "Loading spec...")
	spec, docs, err := files.LoadSpec(cxt, specRoot, options.Files, options.Ascii)
	if err != nil {
		return err
	}

	slog.InfoContext(cxt, "Splitting spec...")
	docsByType, err := files.SplitSpec(docs)
	if err != nil {
		return err
	}
	appClusterIndexes := docsByType[matter.DocTypeAppClusterIndex]
	//deviceTypes := docsByType[matter.DocTypeDeviceType]

	docsByPath := make(map[string]*ascii.Doc)
	for _, doc := range docs {
		docsByPath[doc.Path] = doc
	}

	slog.InfoContext(cxt, "Assigning index domains...")

	err = files.ProcessDocs(cxt, appClusterIndexes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}
		doc.Domain = zap.StringToDomain(top.Name)
		slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		return nil
	}, options.Files)
	if err != nil {
		return err
	}

	slog.InfoContext(cxt, "Extracting clusters...")
	var clusters []*ascii.Doc
	for _, d := range docs {

		entities, err := d.Entities()
		if err != nil {
			slog.Warn("error parsing doc", "path", d.Path, "error", err)
			continue
		}

		for _, m := range entities {
			switch m.(type) {
			case *matter.Cluster:
				clusters = append(clusters, d)

			}
		}
	}

	if len(paths) > 0 {
		filteredDocs := make([]*ascii.Doc, 0, len(paths))
		pathMap := make(map[string]struct{})
		for _, p := range paths {
			pathMap[filepath.Base(p)] = struct{}{}
		}
		for _, ac := range clusters {
			p := filepath.Base(ac.Path)
			if _, ok := pathMap[p]; ok {
				filteredDocs = append(filteredDocs, ac)
				delete(pathMap, p)
			}
		}
		clusters = filteredDocs
	}

	outputs, err := renderTestPlans(cxt, spec, docsByPath, clusters, testPlanRoot, options.Files, options.Overwrite)
	if err != nil {
		return err
	}

	if !options.Files.DryRun {

		for path, result := range outputs {
			if len(result) == 0 {
				continue
			}

			err = os.WriteFile(path, []byte(result), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func getTestPlanPath(testplanRoot string, name string) string {
	return filepath.Join(testplanRoot, "src/cluster/", name+".adoc")
}

func testPlanName(path string, entities []types.Entity) string {

	path = filepath.Base(path)
	name := strings.TrimSuffix(path, filepath.Ext(path))

	var suffix string
	for _, m := range entities {
		switch m.(type) {
		case *matter.Cluster:
			suffix = "Cluster"
		}
	}
	if !strings.HasSuffix(name, suffix) {
		name += " " + suffix
	}
	return strcase.ToKebab(name)
}
