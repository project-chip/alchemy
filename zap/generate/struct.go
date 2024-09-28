package generate

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func generateStructs(structs map[*matter.Struct][]*matter.Number, docPath string, configuratorElement *etree.Element, errata *errata.ZAP) (err error) {

	for _, se := range configuratorElement.SelectElements("struct") {

		nameAttr := se.SelectAttr("name")
		if nameAttr == nil {
			slog.Warn("missing name attribute in struct", slog.String("path", docPath))
			continue
		}
		name := nameAttr.Value

		var matchingStruct *matter.Struct
		var clusterIds []*matter.Number
		var skip bool
		for s, handled := range structs {
			if errata.TypeName(s.Name) == name || errata.TypeName(strings.TrimSuffix(s.Name, "Struct")) == name {
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
			slog.Warn("unknown struct name", slog.String("path", docPath), slog.String("structName", name))
			configuratorElement.RemoveChild(se)
			continue
		}
		if errata != nil && errata.SeparateStructs != nil {
			if _, ok := errata.SeparateStructs[name]; ok {

				amendedClusterCodes, remainingClusterIds := amendExistingClusterCodes(se, matchingStruct, clusterIds)

				populateStruct(se, matchingStruct, amendedClusterCodes, false, errata)
				structs[matchingStruct] = remainingClusterIds
				continue
			}
		}
		populateStruct(se, matchingStruct, clusterIds, false, errata)
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
					populateStruct(bme, s, []*matter.Number{clusterID}, false, errata)
					xml.AppendElement(configuratorElement, bme, "enum", "bitmap")
				}
				continue
			}
		}
		bme := etree.NewElement("struct")
		populateStruct(bme, s, clusterIds, true, errata)
		xml.InsertElementByAttribute(configuratorElement, bme, "name", "enum", "bitmap", "domain")
	}

	return
}

func populateStruct(ee *etree.Element, s *matter.Struct, clusterIDs []*matter.Number, provisional bool, errata *errata.ZAP) (remainingClusterIDs []*matter.Number) {

	ee.CreateAttr("name", errata.TypeName(s.Name))
	if provisional {
		ee.CreateAttr("apiMaturity", "provisional")
	}
	if s.FabricScoping == matter.FabricScopingScoped {
		ee.CreateAttr("isFabricScoped", "true")
	} else {
		ee.RemoveAttr("isFabricScoped")
	}

	_, remainingClusterIDs = amendExistingClusterCodes(ee, s, clusterIDs)
	flushClusterCodes(ee, remainingClusterIDs)

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
			if conformance.IsZigbee(s.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
				continue
			}
			if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeStructField) {
				continue
			}
			setStructFieldAttributes(fe, s, f, errata)
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
		setStructFieldAttributes(fe, s, field, errata)
		xml.AppendElement(ee, fe, "cluster")
	}

	return
}

func setStructFieldAttributes(e *etree.Element, s *matter.Struct, v *matter.Field, errata *errata.ZAP) {
	// Remove incorrect attributes from legacy XML
	e.RemoveAttr("code")
	e.RemoveAttr("id")
	xml.PrependAttribute(e, "fieldId", v.ID.IntString())
	e.CreateAttr("name", v.Name)
	writeDataType(e, s.Fields, v, errata)
	if v.Quality.Has(matter.QualityNullable) {
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
	setFieldDefault(e, v, s.Fields)
	renderConstraint(e, s.Fields, v)
}
