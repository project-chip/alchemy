package config

type Config struct {
	MinimumVersion string    `yaml:"minimum-version"`
	Libraries      []Library `yaml:"libraries"`
}

type Library struct {
	Name string `yaml:"name"`
	Root string `yaml:"root-document"`
}
