package render

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
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
		if needFeatures && !cr.generator.options.FeatureXML {

			err = cr.populateBitmap(bm, &features.Bitmap, clusterIds)
			needFeatures = false
		} else {
			configuratorElement.RemoveChild(bm)
		}
	}
	if cr.generator.options.FeatureXML {
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
	err = cr.renderFeatureElements(cluster, fse, true, cr.configurator.Errata)
	return
}

func (cr *configuratorRenderer) renderFeatureElements(cluster *matter.Cluster, features *etree.Element, excludeDisallowed bool, errata *errata.SDK) (err error) {
	for _, b := range cluster.Features.Bits {
		f, ok := b.(*matter.Feature)
		if !ok {
			err = fmt.Errorf("feature bits contains non-feature bit %s on cluster %s ", b.Name(), cluster.Name)
			return
		}
		if excludeDisallowed && conformance.IsDisallowed(f.Conformance()) {
			continue
		}
		bit := matter.ParseNumber(f.Bit())
		if !bit.Valid() {
			continue
		}
		feature := features.CreateElement("feature")
		feature.CreateAttr("bit", bit.IntString())
		feature.CreateAttr("code", f.Code)
		name := errata.OverrideName(b, f.Name())
		feature.CreateAttr("name", name)
		if len(f.Summary()) > 0 {
			feature.CreateAttr("summary", scrubDescription(f.Summary()))
		}
		cr.setProvisional(feature, f)

		var conformanceElement *etree.Element
		conformanceElement, err = dm.CreateConformanceElement(f.Conformance(), nil)
		if err != nil {
			return err
		}
		if conformanceElement != nil {
			feature.AddChild(conformanceElement)
		}
	}
	return
}

func scrubDescription(description string) string {
	return strings.Join(strings.Fields(description), " ")
}
