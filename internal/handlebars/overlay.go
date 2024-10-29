package handlebars

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/project-chip/alchemy/internal/files"
)

type Overlay struct {
	diskRoot     string
	embedded     fs.FS
	embeddedRoot string
}

func NewOverlay(diskRoot string, embedded fs.FS, embeddedRoot string) *Overlay {
	embedded, _ = fs.Sub(embedded, embeddedRoot)
	return &Overlay{
		diskRoot:     diskRoot,
		embedded:     embedded,
		embeddedRoot: embeddedRoot,
	}
}

func (ov *Overlay) Open(path string) (fs.File, error) {
	if ov.diskRoot != "" {
		var err error
		path, err = filepath.Rel(ov.embeddedRoot, path)
		if err != nil {
			return nil, err
		}
		path = filepath.Join(ov.diskRoot, path)
		exists, err := files.Exists(path)
		if err != nil {
			return nil, err
		}
		if exists {
			return os.Open(path)
		}
	}
	return ov.embedded.Open(path)
}

func (ov *Overlay) ReadDir(path string) ([]fs.DirEntry, error) {
	if ov.diskRoot != "" {
		return os.ReadDir(filepath.Join(ov.diskRoot, path))
	}
	return fs.ReadDir(ov.embedded, path)
}

func (ov *Overlay) ReadFile(path string) ([]byte, error) {
	if ov.diskRoot != "" {
		path = filepath.Join(ov.diskRoot, path)
		exists, err := files.Exists(path)
		if err != nil {
			return nil, err
		}
		if exists {
			return os.ReadFile(path)
		}
	}
	return fs.ReadFile(ov.embedded, path)
}

func (ov *Overlay) Flush() error {
	if ov.embedded == nil {
		return nil
	}
	if ov.diskRoot == "" {
		return nil
	}
	exists, err := files.Exists(ov.diskRoot)
	if err != nil {
		return err
	}
	if !exists {
		err = os.MkdirAll(ov.diskRoot, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return fs.WalkDir(ov.embedded, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		outPath := filepath.Join(ov.diskRoot, path)
		if d.IsDir() {
			err := os.MkdirAll(outPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			exists, err := files.Exists(outPath)
			if err != nil {
				return err
			}
			if !exists {
				b, err := fs.ReadFile(ov.embedded, path)
				if err != nil {
					return err
				}
				err = os.WriteFile(outPath, b, os.ModePerm)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
