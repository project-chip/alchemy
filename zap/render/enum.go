package render

import (
	"fmt"
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
	errata := cr.configurator.Errata
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
		for en, handled := range enums {
			if en.Name == name || strings.TrimSuffix(en.Name, "Enum") == name {
				matchingEnum = en
				skip = len(handled) == 0
				clusterIds = handled
				enums[en] = nil
				break
			}
		}

		if skip {
			continue
		}

		if matchingEnum == nil {
			slog.Warn("Removing unrecognized enum from ZAP XML", slog.String("path", cr.configurator.OutPath), slog.String("enumName", name))
			ce.RemoveChild(eve)
			continue
		}
		if errata != nil && errata.SeparateEnums != nil {
			if _, ok := errata.SeparateEnums[name]; ok {
				amendedClusterCodes, remainingClusterIds := amendExistingClusterCodes(eve, matchingEnum, clusterIds)
				cr.populateEnum(eve, matchingEnum, amendedClusterCodes)
				enums[matchingEnum] = remainingClusterIds
				continue
			}
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
		if cr.isProvisionalViolation(en) {
			err = fmt.Errorf("new enum added without provisional conformance: %s", en.Name)
			return
		}
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

	ee.CreateAttr("name", en.Name)
	ee.CreateAttr("type", en.Type.Name)

	cr.setProvisional(ee, en)

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
			if conformance.IsZigbee(value.Conformance) || zap.IsDisallowed(value, value.Conformance) {
				continue
			}
			cr.setEnumItemAttributes(be, value, valFormat)
			break
		}
	}
	for itemIndex < len(en.Values) {
		value := en.Values[itemIndex]
		itemIndex++
		if conformance.IsZigbee(value.Conformance) || zap.IsDisallowed(value, value.Conformance) {
			continue
		}
		if cr.isProvisionalViolation(value) {
			err = fmt.Errorf("new enum value added without provisional conformance: %s.%s", en.Name, value.Name)
			return
		}
		ie := etree.NewElement("item")
		cr.setEnumItemAttributes(ie, value, valFormat)
		xml.AppendElement(ee, ie, "cluster")
	}

	return
}

func (cr *configuratorRenderer) setEnumItemAttributes(e *etree.Element, v *matter.EnumValue, valFormat string) {
	name := zap.CleanName(v.Name)
	e.CreateAttr("name", name)
	patchNumberAttributeFormat(e, v.Value, "value", valFormat)
}
