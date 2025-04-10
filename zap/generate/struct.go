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

func (cr *configuratorRenderer) generateStructs(structs map[*matter.Struct][]*matter.Number, configuratorElement *etree.Element) (err error) {
	errata := cr.configurator.Errata
	for _, se := range configuratorElement.SelectElements("struct") {

		nameAttr := se.SelectAttr("name")
		if nameAttr == nil {
			slog.Warn("missing name attribute in struct", slog.String("path", cr.configurator.OutPath))
			continue
		}
		name := nameAttr.Value

		var matchingStruct *matter.Struct
		var clusterIds []*matter.Number
		var skip bool
		for s, handled := range structs {
			typeName := errata.OverrideName(s, s.Name)
			if typeName == name || strings.TrimSuffix(typeName, "Struct") == name {
				matchingStruct = s
				skip = len(handled) == 0
				clusterIds = handled
				break
			}
		}

		if skip {
			continue
		}

		if matchingStruct == nil {
			slog.Warn("unknown struct name", slog.String("path", cr.configurator.OutPath), slog.String("structName", name))
			configuratorElement.RemoveChild(se)
			continue
		}
		if errata != nil && errata.SeparateStructs != nil {
			if _, ok := errata.SeparateStructs[name]; ok {

				amendedClusterCodes, remainingClusterIds := amendExistingClusterCodes(se, matchingStruct, clusterIds)

				cr.populateStruct(se, matchingStruct, amendedClusterCodes, false)
				structs[matchingStruct] = remainingClusterIds
				continue
			}
		}
		cr.populateStruct(se, matchingStruct, clusterIds, false)
		structs[matchingStruct] = nil
	}

	var remainingStructs []*matter.Struct
	for s, clusterIds := range structs {
		if len(clusterIds) == 0 {
			continue
		}
		remainingStructs = append(remainingStructs, s)
	}

	slices.SortStableFunc(remainingStructs, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, s := range remainingStructs {
		clusterIds := structs[s]
		if errata != nil && errata.SeparateStructs != nil {
			if _, ok := errata.SeparateStructs[s.Name]; ok {

				for _, clusterID := range clusterIds {
					bme := etree.NewElement("struct")
					cr.populateStruct(bme, s, []*matter.Number{clusterID}, false)
					xml.AppendElement(configuratorElement, bme, "enum", "bitmap")
				}
				continue
			}
		}
		bme := etree.NewElement("struct")
		cr.populateStruct(bme, s, clusterIds, true)
		xml.InsertElementByAttribute(configuratorElement, bme, "name", "enum", "bitmap", "domain")
	}

	return
}

func (cr *configuratorRenderer) populateStruct(ee *etree.Element, s *matter.Struct, clusterIDs []*matter.Number, provisional bool) (remainingClusterIDs []*matter.Number) {
	cr.elementMap[ee] = s
	ee.CreateAttr("name", cr.configurator.Errata.OverrideName(s, s.Name))
	if provisional {
		ee.CreateAttr("apiMaturity", "provisional")
	}
	if s.FabricScoping == matter.FabricScopingScoped {
		ee.CreateAttr("isFabricScoped", "true")
	} else {
		ee.RemoveAttr("isFabricScoped")
	}

	if !cr.configurator.Global {
		_, remainingClusterIDs = amendExistingClusterCodes(ee, s, clusterIDs)
		flushClusterCodes(ee, remainingClusterIDs)
	}

	fieldIndex := 0
	fieldElements := ee.SelectElements("item")
	for _, fe := range fieldElements {
		for {
			if fieldIndex >= len(s.Fields) {
				ee.RemoveChild(fe)
				break
			}
			f := s.Fields[fieldIndex]
			fieldIndex++
			if conformance.IsZigbee(s.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) || conformance.IsDeprecated(f.Conformance) {
				continue
			}
			if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeStructField) {
				continue
			}
			cr.setStructFieldAttributes(fe, s, f)
			break
		}
	}
	for fieldIndex < len(s.Fields) {
		field := s.Fields[fieldIndex]
		fieldIndex++
		if conformance.IsZigbee(s.Fields, field.Conformance) || conformance.IsDisallowed(field.Conformance) {
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(field.ID, types.EntityTypeStructField) {
			continue
		}
		fe := etree.NewElement("item")
		cr.setStructFieldAttributes(fe, s, field)
		xml.AppendElement(ee, fe, "cluster")
	}

	return
}

func (cr *configuratorRenderer) setStructFieldAttributes(e *etree.Element, s *matter.Struct, v *matter.Field) {
	// Remove incorrect attributes from legacy XML
	e.RemoveAttr("code")
	e.RemoveAttr("id")
	xml.PrependAttribute(e, "fieldId", v.ID.IntString())
	name := zap.CleanName(v.Name)
	name = cr.configurator.Errata.OverrideName(s, name)
	e.CreateAttr("name", name)
	cr.writeDataType(e, types.EntityTypeStruct, s.Name, s.Fields, v)
	if v.Quality.Has(matter.QualityNullable) && !cr.generator.generateExtendedQualityElement {
		e.CreateAttr("isNullable", "true")
	} else {
		e.RemoveAttr("isNullable")
	}
	if !conformance.IsMandatory(v.Conformance) {
		e.CreateAttr("optional", "true")
	} else {
		e.RemoveAttr("optional")
	}
	if v.Access.IsFabricSensitive() {
		e.CreateAttr("isFabricSensitive", "true")
	} else {
		e.RemoveAttr("isFabricSensitive")
	}
	cr.setFieldFallback(e, v, s.Fields)
	cr.setQuality(e, v.EntityType(), v.Quality)
	cr.renderConstraint(e, s.Fields, v)
}
