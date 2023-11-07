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

func renderAppCluster(cxt context.Context, doc *ascii.Doc, models []interface{}, w *etree.Element, errata *Errata) (err error) {

	var clusters []*matter.Cluster
	for _, m := range models {
		switch v := m.(type) {
		case *matter.Cluster:
			clusters = append(clusters, v)
		}
	}

	for _, top := range errata.topOrder {
		switch top {
		case matter.SectionCluster:
			for _, c := range clusters {
				err = renderCluster(cxt, doc, c, w, errata)
				if err != nil {
					return
				}
			}
		case matter.SectionDataTypes:
			if len(clusters) > 0 {
				renderDataTypes(clusters[0], clusters, w, errata)
			}

		case matter.SectionFeatures:
			if len(clusters) > 0 {
				renderFeatures(cxt, clusters[0].Features, clusters, w, errata)
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
