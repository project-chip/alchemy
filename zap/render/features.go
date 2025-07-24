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
	"github.com/project-chip/alchemy/zap"
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
	var features []*matter.Feature

	if cluster.Features != nil {
		for feature := range cluster.Features.FeatureBits() {
			if conformance.IsZigbee(feature.Conformance()) || zap.IsDisallowed(feature, feature.Conformance()) {
				continue
			}
			bit := matter.ParseNumber(feature.Bit())
			if !bit.Valid() {
				continue
			}
			features = append(features, feature)
		}
	}

	needFeatures := len(features) > 0

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
	err = cr.renderFeatureElements(cluster, features, fse, cr.configurator.Errata)
	return
}

func (cr *configuratorRenderer) renderFeatureElements(cluster *matter.Cluster, features []*matter.Feature, featuresElement *etree.Element, errata *errata.SDK) (err error) {
	for _, feature := range features {

		bit := matter.ParseNumber(feature.Bit())
		if !bit.Valid() {
			err = fmt.Errorf("invalid bit generating feature: %s", feature.Bit())
			return
		}

		featureElement := featuresElement.CreateElement("feature")
		featureElement.CreateAttr("bit", bit.IntString())
		featureElement.CreateAttr("code", feature.Code)
		name := errata.OverrideName(feature, feature.Name())
		featureElement.CreateAttr("name", name)
		if len(feature.Summary()) > 0 {
			featureElement.CreateAttr("summary", scrubDescription(feature.Summary()))
		}
		cr.setProvisional(featureElement, feature)

		var conformanceElement *etree.Element
		conformanceElement, err = dm.CreateConformanceElement(feature.Conformance(), nil)
		if err != nil {
			return err
		}
		if conformanceElement != nil {
			featureElement.AddChild(conformanceElement)
		}
	}
	return
}

func scrubDescription(description string) string {
	return strings.Join(strings.Fields(description), " ")
}
