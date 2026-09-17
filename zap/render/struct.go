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
	"github.com/project-chip/alchemy/provisional"
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
			if s.Name == name || strings.TrimSuffix(s.Name, "Struct") == name {
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
			slog.Warn("Removing unrecognized struct from ZAP XML", slog.String("path", cr.configurator.OutPath), slog.String("structName", name))
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
		if cr.isProvisionalViolation(s) {
			err = fmt.Errorf("new struct added without provisional conformance: %s", s.Name)
			return
		}
		clusterIds := structs[s]
		if errata != nil && errata.SeparateStructs != nil {
			if _, ok := errata.SeparateStructs[s.Name]; ok {

				for _, clusterID := range clusterIds {
					bme := etree.NewElement("struct")
					_, err = cr.populateStruct(bme, s, []*matter.Number{clusterID}, false)
					if err != nil {
						return
					}
					xml.AppendElement(configuratorElement, bme, "enum", "bitmap")
				}
				continue
			}
		}

		bme := etree.NewElement("struct")
		_, err = cr.populateStruct(bme, s, clusterIds, true)
		if err != nil {
			return
		}
		xml.InsertElementByAttribute(configuratorElement, bme, "name", "enum", "bitmap", "domain")
	}

	return
}

func (cr *configuratorRenderer) populateStruct(ee *etree.Element, s *matter.Struct, clusterIDs []*matter.Number, newlyAdded bool) (remainingClusterIDs []*matter.Number, err error) {
	cr.elementMap[ee] = s
	ee.CreateAttr("name", s.Name)
	switch provisional.Policy(cr.generator.options.ProvisionalPolicy) {
	case provisional.PolicyNone:
		if newlyAdded {
			ee.CreateAttr("apiMaturity", "provisional")
		}
	default:
		cr.setProvisional(ee, s)
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
			if conformance.IsZigbee(f.Conformance) || zap.IsDisallowed(f, f.Conformance) || conformance.IsDeprecated(f.Conformance) {
				continue
			}
			if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeStructField) {
				continue
			}
			cr.setFieldAttributes(fe, types.EntityTypeStruct, f, s.Fields, newlyAdded)
			break
		}
	}
	for fieldIndex < len(s.Fields) {
		field := s.Fields[fieldIndex]
		fieldIndex++
		if conformance.IsZigbee(field.Conformance) || zap.IsDisallowed(field, field.Conformance) {
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(field.ID, types.EntityTypeStructField) {
			continue
		}
		if cr.isProvisionalViolation(field) {
			err = fmt.Errorf("new struct field added without provisional conformance: %s.%s", s.Name, field.Name)
			return
		}
		fe := etree.NewElement("item")
		cr.setFieldAttributes(fe, types.EntityTypeStruct, field, s.Fields, newlyAdded)
		xml.AppendElement(ee, fe, "cluster")
	}

	return
}
