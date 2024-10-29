package handlebars

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/mailgun/raymond/v2"
)

func LoadTemplate(baseTemplate string, templateDir fs.FS) (*raymond.Template, error) {

	t := raymond.MustParse(baseTemplate)
	err := fs.WalkDir(templateDir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if !strings.EqualFold(ext, ".handlebars") {
			return nil
		}
		val, err := fs.ReadFile(templateDir, path)
		if err != nil {
			return fmt.Errorf("error reading template file %s: %w", path, err)
		}
		p, err := raymond.Parse(string(val))
		if err != nil {
			return fmt.Errorf("error parsing template file %s: %w", path, err)
		}
		t.RegisterPartialTemplate(strings.TrimSuffix(path, filepath.Ext(path)), p)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}
