package generate

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

func generateBitmaps(configurator *zap.Configurator, ce *etree.Element, cluster *matter.Cluster, errata *zap.Errata) (err error) {

	for _, eve := range ce.SelectElements("bitmap") {

		nameAttr := eve.SelectAttr("name")
		if nameAttr == nil {
			slog.Warn("missing name attribute in bitmap", slog.String("path", configurator.Doc.Path))
			continue
		}
		name := nameAttr.Value

		if name == "Feature" {
			// Features are handled separately by configurator
			continue
		}

		var matchingBitmap *matter.Bitmap
		var clusterIds []string
		var skip bool
		for bm, handled := range configurator.Bitmaps {
			if bm.Name == name || strings.TrimSuffix(bm.Name, "Bitmap") == name {
				matchingBitmap = bm
				clusterIds = handled
				skip = len(handled) == 0
				configurator.Bitmaps[bm] = nil
				break
			}
		}

		if skip {
			continue
		}

		if matchingBitmap == nil {
			slog.Warn("unknown bitmap name", slog.String("path", configurator.Doc.Path), slog.String("bitmapName", name))
			continue
		}
		populateBitmap(configurator, eve, matchingBitmap, clusterIds, errata)
	}

	for bm, clusterIds := range configurator.Bitmaps {
		if len(clusterIds) == 0 {
			continue
		}
		bme := etree.NewElement("bitmap")
		populateBitmap(configurator, bme, bm, clusterIds, errata)
		insertElementByName(ce, bme, "name", "domain")
	}
	return
}

func populateBitmap(configurator *zap.Configurator, ee *etree.Element, bm *matter.Bitmap, clusterIds []string, errata *zap.Errata) (err error) {

	var valFormat string
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

	ee.CreateAttr("name", bm.Name)
	if bm.Type != nil {
		ee.CreateAttr("type", zap.ConvertDataTypeNameToZap(bm.Type.Name))
	} else {
		ee.CreateAttr("type", "bitmap8")
	}

	_, remainingClusterIds := amendExistingClusterCodes(ee, bm, clusterIds)
	flushClusterCodes(ee, remainingClusterIds)

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
			if conformance.IsZigbee(bm.Bits, bit.Conformance) {
				continue
			}
			err = setBitmapFieldAttributes(be, bit, valFormat)
			if err != nil {
				return
			}
			break
		}
	}
	for bitIndex < len(bm.Bits) {
		bit := bm.Bits[bitIndex]
		bitIndex++
		if conformance.IsZigbee(bm.Bits, bit.Conformance) {
			continue
		}
		fe := etree.NewElement("field")
		err = setBitmapFieldAttributes(fe, bit, valFormat)
		if err != nil {
			return
		}
		appendElement(ee, fe, "cluster")
	}

	return
}

func setBitmapFieldAttributes(e *etree.Element, b *matter.Bit, valFormat string) error {

	mask, err := b.Mask()
	if err != nil {
		return err
	}

	name := zap.CleanName(b.Name)
	e.CreateAttr("name", name)
	e.CreateAttr("mask", fmt.Sprintf(valFormat, mask))
	return nil

}