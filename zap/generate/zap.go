package generate

import (
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/internal/files"
)

type Options struct {
	Files     files.Options
	Ascii     []configuration.Setting
	Overwrite bool
}

func getZapPath(sdkRoot string, name string) string {
	newPath := filepath.Join(sdkRoot, "src/app/zap-templates/zcl/data-model/chip", name+".xml")
	return newPath
}
