package dm

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func Render(cxt context.Context, specRoot string, sdkRoot string, filesOptions files.Options, paths []string, asciiSettings []configuration.Setting) error {

	slog.InfoContext(cxt, "Loading spec...")
	_, docs, err := files.LoadSpec(cxt, specRoot, filesOptions, asciiSettings)
	if err != nil {
		return fmt.Errorf("error loading spec from %s: %w", specRoot, err)
	}

	slog.InfoContext(cxt, "Splitting spec...")
	docsByType, err := files.SplitSpec(docs)
	if err != nil {
		return fmt.Errorf("error splitting spec: %w", err)
	}
	appClusters := docsByType[matter.DocTypeCluster]
	appClusterIndexes := docsByType[matter.DocTypeAppClusterIndex]
	deviceTypes := docsByType[matter.DocTypeDeviceType]

	slog.InfoContext(cxt, "Assigning index domains...")

	err = files.ProcessDocs(cxt, appClusterIndexes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}
		doc.Domain = zap.StringToDomain(top.Name)
		slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		return nil
	}, filesOptions)
	if err != nil {
		return fmt.Errorf("error assigning domains from %s: %w", specRoot, err)
	}

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

	err = renderAppClusters(cxt, sdkRoot, appClusters, filesOptions)
	if err != nil {
		return err
	}
	return renderDeviceTypes(cxt, sdkRoot, deviceTypes, filesOptions)
}
