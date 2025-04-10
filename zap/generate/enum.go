package generate

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) generateEnums(enums map[*matter.Enum][]*matter.Number, ce *etree.Element) (err error) {

	for _, eve := range ce.SelectElements("enum") {

		nameAttr := eve.SelectAttr("name")
		if nameAttr == nil {
			slog.Warn("ZAP: missing name attribute in enum", slog.String("path", cr.configurator.OutPath))
			continue
		}
		name := nameAttr.Value

		var matchingEnum *matter.Enum
		var clusterIds []*matter.Number
		var skip bool
		for bm, handled := range enums {
			typeName := cr.configurator.Errata.OverrideName(bm, bm.Name)
			if typeName == name || strings.TrimSuffix(typeName, "Enum") == name {
				matchingEnum = bm
				skip = len(handled) == 0
				clusterIds = handled
				enums[bm] = nil
				break
			}
		}

		if skip {
			continue
		}

		if matchingEnum == nil {
			slog.Warn("ZAP: unknown enum name", slog.String("path", cr.configurator.OutPath), slog.String("enumName", name))
			ce.RemoveChild(eve)
			continue
		}
		err = cr.populateEnum(eve, matchingEnum, clusterIds)
		if err != nil {
			return
		}
	}

	var remainingEnums []*matter.Enum
	for en, clusterIds := range enums {
		if len(clusterIds) == 0 {
			continue
		}
		remainingEnums = append(remainingEnums, en)
	}

	if len(remainingEnums) == 0 {
		return
	}

	slices.SortStableFunc(remainingEnums, func(a, b *matter.Enum) int { return strings.Compare(a.Name, b.Name) })

	for _, en := range remainingEnums {
		bme := etree.NewElement("enum")
		clusterIds := enums[en]
		err = cr.populateEnum(bme, en, clusterIds)
		if err != nil {
			return
		}
		xml.InsertElementByAttribute(ce, bme, "name", "bitmap", "domain")
	}

	return
}

func (cr *configuratorRenderer) populateEnum(ee *etree.Element, en *matter.Enum, clusterIds []*matter.Number) (err error) {
	cr.elementMap[ee] = en
	var valFormat string
	switch en.Type.BaseType {
	case types.BaseDataTypeEnum16:
		valFormat = "0x%04X"
	default:
		valFormat = "0x%02X"
	}

	ee.CreateAttr("name", cr.configurator.Errata.OverrideName(en, en.Name))
	ee.CreateAttr("type", cr.configurator.Errata.OverrideType(en, en.Type.Name))

	if !cr.configurator.Global {
		_, remainingClusterIds := amendExistingClusterCodes(ee, en, clusterIds)
		flushClusterCodes(ee, remainingClusterIds)
	}

	itemIndex := 0
	itemElements := ee.SelectElements("item")
	for _, be := range itemElements {
		for {
			if itemIndex >= len(en.Values) {
				ee.RemoveChild(be)
				break
			}
			value := en.Values[itemIndex]
			itemIndex++
			if conformance.IsZigbee(en.Values, value.Conformance) || conformance.IsDisallowed(value.Conformance) {
				continue
			}
			cr.setEnumItemAttributes(be, value, valFormat)
			break
		}
	}
	for itemIndex < len(en.Values) {
		value := en.Values[itemIndex]
		itemIndex++
		if conformance.IsZigbee(en.Values, value.Conformance) || conformance.IsDisallowed(value.Conformance) {
			continue
		}
		ie := etree.NewElement("item")
		cr.setEnumItemAttributes(ie, value, valFormat)
		xml.AppendElement(ee, ie, "cluster")
	}

	return
}

func (cr *configuratorRenderer) setEnumItemAttributes(e *etree.Element, v *matter.EnumValue, valFormat string) {
	name := zap.CleanName(v.Name)
	name = cr.configurator.Errata.OverrideName(v, name)
	e.CreateAttr("name", name)
	patchNumberAttributeFormat(e, v.Value, "value", valFormat)
}
