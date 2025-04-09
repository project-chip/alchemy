package generate

import (
	"log/slog"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
)

func (cr *configuratorRenderer) generateFeatures(configuratorElement *etree.Element, features *matter.Features) (err error) {

	needFeatures := features != nil && len(features.Bits) > 0

	bitmaps := configuratorElement.SelectElements("bitmap")

	clusterIds := cr.configurator.Features
	for _, bm := range bitmaps {
		nameAttr := bm.SelectAttr("name")
		if nameAttr == nil || nameAttr.Value != "Feature" {
			continue
		}
		if needFeatures && !cr.generator.generateFeaturesXML {

			err = cr.populateBitmap(bm, &features.Bitmap, clusterIds)
			needFeatures = false
		} else {
			configuratorElement.RemoveChild(bm)
		}
	}
	if cr.generator.generateFeaturesXML {
		return
	}
	if needFeatures {
		fe := etree.NewElement("bitmap")
		err = cr.populateBitmap(fe, &features.Bitmap, clusterIds)
		if err != nil {
			return
		}
		xml.AppendElement(configuratorElement, fe, "domain")
	}
	return
}

func (cr *configuratorRenderer) generateFeaturesXML(configuratorElement *etree.Element, cluster *matter.Cluster) (err error) {
	features := cluster.Features
	needFeatures := features != nil && len(features.Bits) > 0

	bitmaps := configuratorElement.SelectElements("bitmap")

	for _, bm := range bitmaps {
		nameAttr := bm.SelectAttr("name")
		if nameAttr == nil || nameAttr.Value != "Feature" {
			continue
		}
		configuratorElement.RemoveChild(bm)
	}
	fse := configuratorElement.SelectElement("features")
	if !needFeatures {
		if fse != nil {
			configuratorElement.RemoveChild(fse)
		}
		return
	}
	if fse == nil {
		fse = etree.NewElement("features")
		xml.AppendElement(configuratorElement, fse, "client", "server", "code", "define", "description", "domain", "name")
	} else {
		fse.Child = nil
	}
	doc, ok := cr.generator.spec.DocRefs[cluster]
	if !ok {
		slog.Warn("missing doc ref for cluster", slog.String("clusterName", cluster.Name))
	}
	err = dm.RenderFeatureElements(doc, cluster, fse, true, cr.configurator.Errata)
	return
}
