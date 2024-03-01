package dm

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/matter"
)

func Render(cxt context.Context, specRoot string, sdkRoot string, filesOptions files.Options, paths []string, asciiSettings []configuration.Setting) error {

	slog.InfoContext(cxt, "Loading spec...")
	_, docs, err := files.LoadSpec(cxt, specRoot, filesOptions, asciiSettings)
	if err != nil {
		return fmt.Errorf("error loading spec from %s: %w", specRoot, err)
	}

	if len(paths) > 0 {
		filteredDocs := make([]*ascii.Doc, 0, len(paths))
		pathMap := make(map[string]struct{})
		for _, p := range paths {
			pathMap[filepath.Base(p)] = struct{}{}
		}
		for _, d := range docs {
			p := filepath.Base(d.Path)
			if _, ok := pathMap[p]; ok {
				filteredDocs = append(filteredDocs, d)
				delete(pathMap, p)
			}

		}
		docs = filteredDocs
	}

	var appClusters, deviceTypes []*ascii.Doc
	for _, doc := range docs {
		entites, err := doc.Entities()
		if err != nil {
			slog.ErrorContext(cxt, "error converting doc to entities", "doc", doc.Path, "error", err)
			continue
		}
		for _, e := range entites {
			switch e.(type) {
			case *matter.Cluster:
				appClusters = append(appClusters, doc)
			case *matter.DeviceType:
				deviceTypes = append(deviceTypes, doc)
			}
		}
	}

	err = renderAppClusters(cxt, sdkRoot, appClusters, filesOptions)
	if err != nil {
		return err
	}
	return renderDeviceTypes(cxt, sdkRoot, deviceTypes, filesOptions)
}
