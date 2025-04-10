package generate

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) generateBitmaps(bitmaps map[*matter.Bitmap][]*matter.Number, parent *etree.Element) (err error) {

	for _, eve := range parent.SelectElements("bitmap") {

		nameAttr := eve.SelectAttr("name")
		if nameAttr == nil {
			slog.Warn("missing name attribute in bitmap", slog.String("path", cr.configurator.OutPath))
			continue
		}
		name := nameAttr.Value

		if name == "Feature" {
			// Features are handled separately by configurator
			continue
		}

		var matchingBitmap *matter.Bitmap
		var clusterIds []*matter.Number
		var skip bool
		for bm, handled := range bitmaps {
			typeName := cr.configurator.Errata.OverrideName(bm, bm.Name)
			if typeName == name || strings.TrimSuffix(typeName, "Bitmap") == name {
				matchingBitmap = bm
				clusterIds = handled
				skip = len(handled) == 0
				bitmaps[bm] = nil
				break
			}
		}

		if skip {
			continue
		}

		if matchingBitmap == nil {
			slog.Warn("unknown bitmap name", slog.String("path", cr.configurator.OutPath), slog.String("bitmapName", name))
			parent.RemoveChild(eve)
			continue
		}
		err = cr.populateBitmap(eve, matchingBitmap, clusterIds)
		if err != nil {
			return
		}
	}

	for bm, clusterIds := range bitmaps {
		if len(clusterIds) == 0 {
			continue
		}
		bme := etree.NewElement("bitmap")
		err = cr.populateBitmap(bme, bm, clusterIds)
		if err != nil {
			return
		}
		xml.InsertElementByAttribute(parent, bme, "name", "domain")
	}
	return
}

func (cr *configuratorRenderer) populateBitmap(ee *etree.Element, bm *matter.Bitmap, clusterIds []*matter.Number) (err error) {
	cr.elementMap[ee] = bm
	var valFormat string
	if bm.Name == "Feature" {
		valFormat = "0x%02X"
	} else {
		switch bm.Type.BaseType {
		case types.BaseDataTypeMap64:
			valFormat = "0x%016X"
		case types.BaseDataTypeMap32:
			valFormat = "0x%08X"
		case types.BaseDataTypeMap16:
			valFormat = "0x%04X"
		default:
			valFormat = "0x%02X"
		}

	}

	ee.CreateAttr("name", cr.configurator.Errata.OverrideName(bm, bm.Name))
	var typeName string
	if bm.Type != nil {
		typeName = cr.configurator.Errata.OverrideType(bm, zap.DataTypeName(bm.Type))
	} else {
		typeName = cr.configurator.Errata.OverrideType(bm, "bitmap8")
	}
	ee.CreateAttr("type", typeName)

	if !cr.configurator.Global {
		_, remainingClusterIds := amendExistingClusterCodes(ee, bm, clusterIds)
		flushClusterCodes(ee, remainingClusterIds)
	}

	bitIndex := 0
	bitElements := ee.SelectElements("field")
	for _, be := range bitElements {
		for {
			if bitIndex >= len(bm.Bits) {
				ee.RemoveChild(be)
				break
			}
			bit := bm.Bits[bitIndex]
			bitIndex++
			if conformance.IsZigbee(bm.Bits, bit.Conformance()) || conformance.IsDisallowed(bit.Conformance()) {
				continue
			}
			err = cr.setBitmapFieldAttributes(be, bit, valFormat)
			if err != nil {
				return
			}
			break
		}
	}
	for bitIndex < len(bm.Bits) {
		bit := bm.Bits[bitIndex]
		bitIndex++
		if conformance.IsZigbee(bm.Bits, bit.Conformance()) || conformance.IsDisallowed(bit.Conformance()) {
			continue
		}
		fe := etree.NewElement("field")
		err = cr.setBitmapFieldAttributes(fe, bit, valFormat)
		if err != nil {
			return
		}
		xml.AppendElement(ee, fe, "cluster")
	}

	return
}

func (cr *configuratorRenderer) setBitmapFieldAttributes(e *etree.Element, b matter.Bit, valFormat string) error {

	mask, err := b.Mask()
	if err != nil {
		return err
	}

	name := b.Name()
	name = zap.CleanName(name)
	name = cr.configurator.Errata.OverrideName(b, name)
	e.CreateAttr("name", name)
	ma := e.SelectAttr("mask")
	if ma != nil {
		ev := matter.ParseNumber(ma.Value)
		if ev.Valid() && ev.Value() == mask {
			return nil
		}
	}
	e.CreateAttr("mask", fmt.Sprintf(valFormat, mask))
	return nil

}
