package config

import (
	_ "embed"
)

//go:embed default.yaml
var defaultConfig []byte
