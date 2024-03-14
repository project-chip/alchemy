package generate

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/zap"
)

func generateStructs(configurator *zap.Configurator, configuratorElement *etree.Element, cluster *matter.Cluster, errata *zap.Errata) (err error) {

	for _, se := range configuratorElement.SelectElements("struct") {

		nameAttr := se.SelectAttr("name")
		if nameAttr == nil {
			slog.Warn("missing name attribute in struct", slog.String("path", configurator.Doc.Path))
			continue
		}
		name := nameAttr.Value

		var matchingStruct *matter.Struct
		var clusterIds []*matter.Number
		var skip bool
		for s, handled := range configurator.Structs {
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
			slog.Warn("unknown struct name", slog.String("path", configurator.Doc.Path), slog.String("structName", name))
			configuratorElement.RemoveChild(se)
			continue
		}
		if errata.SeparateStructs != nil {
			if _, ok := errata.SeparateStructs[name]; ok {

				amendedClusterCodes, remainingClusterIds := amendExistingClusterCodes(se, matchingStruct, clusterIds)

				populateStruct(configurator, se, matchingStruct, cluster, errata, amendedClusterCodes, false)
				configurator.Structs[matchingStruct] = remainingClusterIds
				continue
			}
		}
		populateStruct(configurator, se, matchingStruct, cluster, errata, clusterIds, false)
		configurator.Structs[matchingStruct] = nil
	}

	var remainingStructs []*matter.Struct
	for s, clusterIds := range configurator.Structs {
		if len(clusterIds) == 0 {
			continue
		}
		remainingStructs = append(remainingStructs, s)
	}

	slices.SortFunc(remainingStructs, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, s := range remainingStructs {
		clusterIds := configurator.Structs[s]
		if errata.SeparateStructs != nil {
			if _, ok := errata.SeparateStructs[s.Name]; ok {

				for _, clusterID := range clusterIds {
					bme := etree.NewElement("struct")
					populateStruct(configurator, bme, s, cluster, errata, []*matter.Number{clusterID}, false)
					appendElement(configuratorElement, bme, "enum", "bitmap")
				}
				continue
			}
		}
		bme := etree.NewElement("struct")
		populateStruct(configurator, bme, s, cluster, errata, clusterIds, true)
		insertElementByAttribute(configuratorElement, bme, "name", "enum", "bitmap", "domain")
	}

	return
}

func populateStruct(configurator *zap.Configurator, ee *etree.Element, s *matter.Struct, cluster *matter.Cluster, errata *zap.Errata, clusterIDs []*matter.Number, provisional bool) (remainingClusterIDs []*matter.Number) {

	ee.CreateAttr("name", s.Name)
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
			if conformance.IsZigbee(s.Fields, f.Conformance) {
				continue
			}
			setStructFieldAttributes(fe, s, f)
			break
		}
	}
	for fieldIndex < len(s.Fields) {
		field := s.Fields[fieldIndex]
		fieldIndex++
		if conformance.IsZigbee(s.Fields, field.Conformance) {
			continue
		}
		fe := etree.NewElement("item")
		setStructFieldAttributes(fe, s, field)
		appendElement(ee, fe, "cluster")
	}

	return
}

func setStructFieldAttributes(e *etree.Element, s *matter.Struct, v *matter.Field) {

	e.CreateAttr("fieldId", v.ID.IntString())
	e.CreateAttr("name", v.Name)
	writeDataType(e, s.Fields, v)
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
