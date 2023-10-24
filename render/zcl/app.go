package zcl

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func renderAppCluster(cxt context.Context, doc *ascii.Doc, w *etree.Document) error {
	//domain, err := getDomain(cxt, doc)
	w.CreateComment(license)
	topLevelSection := parse.FindFirst[*ascii.Section](doc.Elements)
	if topLevelSection == nil {
		return fmt.Errorf("missing top level section")
	}
	c := w.CreateElement("configurator")
	dom := c.CreateElement("domain")
	dom.CreateAttr("name", "CHIP")
	ms, err := doc.ToModel()
	if err != nil {
		return err
	}
	for _, m := range ms {
		switch v := m.(type) {
		case *matter.Cluster:
			err = renderCluster(cxt, v, c)
		}
		if err != nil {
			return err
		}
	}
	for _, m := range ms {
		switch v := m.(type) {
		case *matter.Cluster:
			renderDataTypes(v, c)
		}
	}
	return nil
}

func getDomain(cxt context.Context, doc *ascii.Doc) (string, error) {
	path := filepath.Dir(doc.Path)
	docs, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, f := range docs {
		if f.IsDir() {
			continue
		}
		docType, _ := ascii.GetDocType(f.Name())
		if docType == matter.DocTypeAppClusterIndex {

		}
	}
	return "", fmt.Errorf("could not determine domain from path %s", doc.Path)
}
