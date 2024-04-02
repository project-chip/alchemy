package dump

import "github.com/bytesparadise/libasciidoc/pkg/configuration"

type Options struct {
	ASCII bool
	JSON  bool

	ASCIISettings []configuration.Setting
}
