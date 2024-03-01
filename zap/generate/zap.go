package generate

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

type Options struct {
	Files     files.Options
	Ascii     []configuration.Setting
	Overwrite bool
}

func Generate(cxt context.Context, specRoot string, sdkRoot string, paths []string, options Options) error {

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
	var clusterDocs []*ascii.Doc
	var deviceTypeDocs []*ascii.Doc
	for _, d := range docs {

		entities, err := d.Entities()
		if err != nil {
			slog.Warn("error parsing doc", "path", d.Path, "error", err)
			continue
		}

		for _, m := range entities {
			switch m.(type) {
			case *matter.Cluster:
				clusterDocs = append(clusterDocs, d)
			case *matter.DeviceType:
				deviceTypeDocs = append(deviceTypeDocs, d)
			}
		}
	}

	clusterDocs = filterDocs(clusterDocs, paths)

	outputs, provisionalZclFiles, err := renderClusterTemplates(cxt, spec, docsByPath, clusterDocs, sdkRoot, options.Files, options.Overwrite)
	if err != nil {
		return err
	}

	err = renderDeviceTypes(cxt, spec, filterDocs(deviceTypeDocs, paths), sdkRoot, options.Files)
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

		if len(provisionalZclFiles) > 0 {
			slog.Info("Patching ZAP JSON...")
			err = patchZapJson(sdkRoot, provisionalZclFiles)

			if err != nil {
				return err
			}

			slog.Info("Patching workflow tests YAML...")
			err = patchTestsYaml(sdkRoot, provisionalZclFiles)
			if err != nil {
				return err
			}

			slog.Info("Patching scripts/matter.lint...")
			err = patchLint(sdkRoot, provisionalZclFiles)
			if err != nil {
				return err
			}
		}

		slog.Info("Patching src/app/zap_cluster_list.json...")
		err = patchClusterList(sdkRoot, clusterDocs)
		if err != nil {
			return err
		}

	}

	return nil
}

func filterDocs(docs []*ascii.Doc, paths []string) []*ascii.Doc {
	if len(docs) == 0 {
		return docs
	}
	filteredDocs := make([]*ascii.Doc, 0, len(paths))
	pathMap := make(map[string]struct{})
	for _, p := range paths {
		pathMap[filepath.Base(p)] = struct{}{}
	}
	for _, ac := range docs {
		p := filepath.Base(ac.Path)
		if _, ok := pathMap[p]; ok {
			filteredDocs = append(filteredDocs, ac)
			delete(pathMap, p)
		}
	}
	return filteredDocs
}

func getZapPath(sdkRoot string, name string) string {
	newPath := filepath.Join(sdkRoot, "src/app/zap-templates/zcl/data-model/chip", name+".xml")
	return newPath
}
