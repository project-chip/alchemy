package dump

import "github.com/bytesparadise/libasciidoc/pkg/configuration"

type Options struct {
	Ascii bool
	Json  bool

	AsciiSettings []configuration.Setting
}
