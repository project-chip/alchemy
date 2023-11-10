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

func Migrate(cxt context.Context, specRoot string, zclRoot string, filesOptions files.Options, asciiSettings []configuration.Setting) error {

	slog.InfoContext(cxt, "Loading spec...")
	docs, err := loadSpec(cxt, specRoot, filesOptions, asciiSettings)
	if err != nil {
		return err
	}

	slog.InfoContext(cxt, "Building spec tree...")
	ascii.BuildTree(docs)

	slog.InfoContext(cxt, "Splitting spec...")
	docsByType, err := splitSpec(docs)
	if err != nil {
		return err
	}
	appClusters := docsByType[matter.DocTypeAppCluster]
	appClusterIndexes := docsByType[matter.DocTypeAppClusterIndex]

	slog.InfoContext(cxt, "Assigning index domains...")
	files.ProcessDocs(cxt, appClusterIndexes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {

		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}
		doc.Domain = zap.StringToDomain(top.Name)
		slog.InfoContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		return nil
	}, filesOptions)

	outputs, provisionalZclFiles, err := renderTemplates(cxt, appClusters, zclRoot, filesOptions)
	if err != nil {
		return err
	}

	if !filesOptions.DryRun {

		for path, result := range outputs {
			if len(result.Models) == 0 {
				continue
			}

			if !filesOptions.DryRun {
				err = os.WriteFile(path, []byte(result.ZCL), os.ModeAppend|0644)
				if err != nil {
					return err
				}
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
