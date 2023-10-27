package cmd

import "github.com/bytesparadise/libasciidoc/pkg/configuration"

type asciiParser struct {
	settings []configuration.Setting
}

func (ap *asciiParser) addSetting(s configuration.Setting) {
	ap.settings = append(ap.settings, s)
}

type asciiSettings interface {
	addSetting(s configuration.Setting)
}
