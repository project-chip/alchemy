package zcl

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func renderAppCluster(cxt context.Context, models []interface{}, w *etree.Element, errata *errata) (err error) {

	for _, top := range errata.topOrder {
		switch top {
		case matter.SectionCluster:
			for _, m := range models {
				switch v := m.(type) {
				case *matter.Cluster:
					err = renderCluster(cxt, v, w, errata)
				}
				if err != nil {
					return err
				}
			}
		case matter.SectionDataTypes:
			for _, m := range models {
				switch v := m.(type) {
				case *matter.Cluster:
					renderDataTypes(v, w, errata)
				}
			}
		case matter.SectionFeatures:
			for _, m := range models {
				switch v := m.(type) {
				case *matter.Cluster:
					renderFeatures(cxt, v, w, errata)
				}
			}
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
