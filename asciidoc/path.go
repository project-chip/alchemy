package asciidoc

import "path/filepath"

type Path struct {
	Absolute string
	Relative string
}

func (p Path) String() string {
	return p.Relative
}

func (p Path) Base() string {
	return filepath.Base(p.Absolute)
}

func (p Path) Ext() string {
	return filepath.Ext(p.Absolute)
}

func (p Path) Dir() string {
	return filepath.Dir(p.Absolute)
}

func (p Path) Origin() (path string, line int) {
	return p.Relative, -1
}

func NewPath(path string, rootPath string) (p Path, err error) {
	if !filepath.IsAbs(path) {
		p.Absolute, err = filepath.Abs(path)
		if err != nil {
			return
		}
	} else {
		p.Absolute = path
	}
	var r string
	r, err = filepath.Abs(rootPath)
	if err != nil {
		return
	}
	p.Relative, err = filepath.Rel(r, p.Absolute)
	return
}
