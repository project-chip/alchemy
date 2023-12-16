package generate

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

var selfClosingTags = regexp.MustCompile("></[^>]+>")

type Options struct {
	Files     files.Options
	Ascii     []configuration.Setting
	Overwrite bool
}

func Migrate(cxt context.Context, specRoot string, zclRoot string, paths []string, options Options) error {

	slog.InfoContext(cxt, "Loading spec...")
	docs, err := files.LoadSpec(cxt, specRoot, options.Files, options.Ascii)
	if err != nil {
		return err
	}

	slog.InfoContext(cxt, "Building spec tree...")
	spec, err := ascii.BuildSpec(docs)
	if err != nil {
		return err
	}

	slog.InfoContext(cxt, "Splitting spec...")
	docsByType, err := files.SplitSpec(docs)
	if err != nil {
		return err
	}
	appClusterIndexes := docsByType[matter.DocTypeAppClusterIndex]
	deviceTypes := docsByType[matter.DocTypeDeviceType]

	docsByPath := make(map[string]*ascii.Doc)
	for _, doc := range docs {
		docsByPath[doc.Path] = doc
	}

	slog.InfoContext(cxt, "Assigning index domains...")

	files.ProcessDocs(cxt, appClusterIndexes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}
		doc.Domain = zap.StringToDomain(top.Name)
		slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		return nil
	}, options.Files)

	slog.InfoContext(cxt, "Extracting clusters...")
	var clusters []*ascii.Doc
	for _, d := range docs {

		models, err := d.ToModel()
		if err != nil {
			slog.Warn("error parsing doc", "path", d.Path, "error", err)
			continue
		}

		for _, m := range models {
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

	outputs, provisionalZclFiles, err := renderClusterTemplates(cxt, spec, docsByPath, clusters, zclRoot, options.Files, options.Overwrite)
	if err != nil {
		return err
	}

	files.ProcessDocs(cxt, deviceTypes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		slog.Debug("Device type doc", "name", doc.Path)

		models, err := doc.ToModel()
		if err != nil {
			return err
		}
		for _, m := range models {
			slog.Debug("model", "type", m)
		}
		return nil
	}, options.Files)

	if !options.Files.DryRun {

		for path, result := range outputs {
			if len(result.Models) == 0 {
				continue
			}

			err = os.WriteFile(path, []byte(result.ZCL), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}

		if len(provisionalZclFiles) > 0 {
			slog.Info("Patching ZAP JSON...")
			err = patchZapJson(zclRoot, provisionalZclFiles)

			if err != nil {
				return err
			}

			slog.Info("Patching workflow tests YAML...")
			err = patchTestsYaml(zclRoot, provisionalZclFiles)
			if err != nil {
				return err
			}

			slog.Info("Patching scripts/matter.lint...")
			err = patchLint(zclRoot, provisionalZclFiles)
			if err != nil {
				return err
			}
		}

		err = patchBuildGN(zclRoot, clusters)
		if err != nil {
			return err
		}

		slog.Info("Patching src/app/zap_cluster_list.json...")
		err = patchClusterList(zclRoot, clusters)
		if err != nil {
			return err
		}

	}

	return nil
}

func getZapPath(zclRoot string, name string) string {
	newPath := filepath.Join(zclRoot, "src/app/zap-templates/zcl/data-model/chip", name+".xml")
	return newPath
}
