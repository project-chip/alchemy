package spec

import (
	"fmt"
	"path/filepath"

	"github.com/project-chip/alchemy/internal/files"
)

type BuilderOption func(tg *Builder)

func IgnoreHierarchy(ignore bool) BuilderOption {
	return func(b *Builder) {
		b.ignoreHierarchy = ignore
	}
}

type BuilderOptions struct {
	IgnoreHierarchy bool `default:"false" help:"ignore hierarchy" group:"Spec:"`
}

type ParserOptions struct {
	Root   string `name:"spec-root" default:"connectedhomeip-spec" aliases:"specRoot" help:"the src root of your clone of CHIP-Specifications/connectedhomeip-spec"  group:"Spec:"`
	Inline bool   `default:"true" help:"use inline parser"  group:"Spec:"`
}

func (po *ParserOptions) AfterApply() error {
	if !filepath.IsAbs(po.Root) {
		var err error
		po.Root, err = filepath.Abs(po.Root)
		if err != nil {
			return err
		}
	}
	exists, err := files.Exists(po.Root)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("spec root %s does not exist", po.Root)
	}
	return nil
}
