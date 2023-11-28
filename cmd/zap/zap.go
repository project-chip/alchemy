package zap

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
	"github.com/iancoleman/strcase"
)

var selfClosingTags = regexp.MustCompile("></[^>]+>")

type migrateOptions struct {
	filesOptions  files.Options
	asciiSettings []configuration.Setting
}

func Migrate(cxt context.Context, specRoot string, zclRoot string, paths []string, options migrateOptions) error {

	slog.InfoContext(cxt, "Loading spec...")
	docs, err := files.LoadSpec(cxt, specRoot, options.filesOptions, options.asciiSettings)
	if err != nil {
		return err
	}

	slog.InfoContext(cxt, "Building spec tree...")
	ascii.BuildTree(docs)

	slog.InfoContext(cxt, "Splitting spec...")
	docsByType, err := files.SplitSpec(docs)
	if err != nil {
		return err
	}
	appClusters := docsByType[matter.DocTypeAppCluster]
	appClusterIndexes := docsByType[matter.DocTypeAppClusterIndex]
	deviceTypes := docsByType[matter.DocTypeDeviceType]

	slog.InfoContext(cxt, "Assigning index domains...")

	files.ProcessDocs(cxt, appClusterIndexes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}
		doc.Domain = zap.StringToDomain(top.Name)
		slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		return nil
	}, options.filesOptions)

	if len(paths) > 0 {
		filteredDocs := make([]*ascii.Doc, 0, len(paths))
		pathMap := make(map[string]struct{})
		for _, p := range paths {
			pathMap[filepath.Base(p)] = struct{}{}
		}
		for _, ac := range appClusters {
			p := filepath.Base(ac.Path)
			if _, ok := pathMap[p]; ok {
				filteredDocs = append(filteredDocs, ac)
				delete(pathMap, p)
			}
		}
		appClusters = filteredDocs
	}

	outputs, provisionalZclFiles, err := renderAppClusterTemplates(cxt, appClusters, zclRoot, options.filesOptions)
	if err != nil {
		return err
	}

	files.ProcessDocs(cxt, deviceTypes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		slog.Info("Device type doc", "name", doc.Path)

		models, err := doc.ToModel()
		if err != nil {
			return err
		}
		for _, m := range models {
			slog.Info("model", "type", m)
		}
		return nil
	}, options.filesOptions)

	if !options.filesOptions.DryRun {

		for path, result := range outputs {
			if len(result.Models) == 0 {
				continue
			}

			err = os.WriteFile(path, []byte(result.ZCL), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}

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

		slog.Info("Patching src/controller/data_model/BUILD.gn...")
		err = patchBuildGN(zclRoot, appClusters)
		if err != nil {
			return err
		}

		slog.Info("Patching src/app/zap_cluster_list.json...")
		err = patchClusterList(zclRoot, appClusters)
		if err != nil {
			return err
		}

	}

	return nil
}

func getZapPath(zclRoot string, path string) string {
	newFile := filepath.Base(path)
	newFile = zap.ZAPName(strings.TrimSuffix(newFile, filepath.Ext(path)))
	newFile = strcase.ToKebab(newFile)
	newPath := filepath.Join(zclRoot, "src/app/zap-templates/zcl/data-model/chip", newFile+".xml")
	return newPath
}
