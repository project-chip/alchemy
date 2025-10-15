package config

import "github.com/project-chip/alchemy/asciidoc"

type Library struct {
	Name string `yaml:"name"`
	Root string `yaml:"root-document"`

	Documents map[string]LibraryDocument `yaml:"documents,omitempty"`
}

type Config struct {
	MinimumVersion string    `yaml:"minimum-version"`
	root           string    `yaml:"-"`
	Libraries      []Library `yaml:"libraries"`

	libraryRoots map[string]struct{}
}

func (c *Config) Root() string {
	return c.root
}

func (c *Config) IsLibraryRootPath(path asciidoc.Path) bool {
	_, ok := c.libraryRoots[path.Relative]
	return ok
}

type LibraryDocument struct {
	Domain string `yaml:"domain"`
}
