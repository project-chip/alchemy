package spec

import (
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func buildSectionTitle(pps *preparseFileState, section *asciidoc.Section, reader asciidoc.Reader, title *strings.Builder, els ...asciidoc.Element) (err error) {
	for e := range reader.Iterate(section, els) {
		switch e := e.(type) {
		case *asciidoc.String:
			title.WriteString(e.Value)
		case *asciidoc.SpecialCharacter:
			title.WriteString(e.Character)
		case *asciidoc.UserAttributeReference:

			val := pps.Get(string(e.Name()))
			if val != nil {
				switch val := val.(type) {
				case asciidoc.Elements:
					err = buildSectionTitle(pps, section, reader, title, val...)
				case asciidoc.Element:
					err = buildSectionTitle(pps, section, reader, title, val)
				default:
					err = newGenericParseError(e, "unexpected section title attribute value type: %T", val)
					return
				}
			} else {
				title.WriteRune('{')
				title.WriteString(e.Name())
				title.WriteRune('}')
				slog.Warn("unknown section title user attribute", slog.String("attributeName", e.Name()), log.Path("source", e))
				return
			}
		case *asciidoc.Link, *asciidoc.LinkMacro:
		case *asciidoc.Bold:
			err = buildSectionTitle(pps, section, reader, title, e.Children()...)
		default:
			err = newGenericParseError(section, "unknown section title component type: %T", e)
		}
		if err != nil {
			return
		}
		if he, ok := e.(asciidoc.ParentElement); ok {
			err = buildSectionTitle(pps, section, reader, title, he.Children()...)
		}
		if err != nil {
			return
		}
	}
	return
}

func AssignSectionTypes(doc *Doc, top *asciidoc.Section) error {
	if _, ok := doc.sectionType(top); ok {
		return nil
	}
	docType, err := doc.DocType()
	if err != nil {
		return err
	}
	var secType matter.Section
	switch docType {
	case matter.DocTypeCluster:
		secType = matter.SectionCluster
	case matter.DocTypeDeviceType:
		secType = matter.SectionDeviceType
	case matter.DocTypeNamespace:
		secType = matter.SectionNamespace
	default:
		secType = matter.SectionTop
		if strings.HasSuffix(doc.SectionName(top), " Cluster") {
			secType = matter.SectionCluster
		}
	}
	assignSectionType(doc, top, secType)

	parse.Search(doc.Reader(), top, top.Children(), func(el any, parent asciidoc.Parent, index int) parse.SearchShould {
		section, ok := el.(*asciidoc.Section)
		if !ok {
			return parse.SearchShouldContinue
		}
		ps, ok := parent.(*asciidoc.Section)
		if !ok {
			return parse.SearchShouldContinue
		}

		assignSectionType(doc, section, getSectionType(doc, ps, section))
		switch doc.SectionType(section) {
		case matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct, matter.SectionDataTypeDef:
			if section.Level > 2 {
				slog.Debug("Unusual depth for section type", slog.String("name", doc.SectionName(section)), slog.String("type", doc.SectionType(section).String()), slog.String("path", doc.Path.String()))
			}
		}
		return parse.SearchShouldContinue
	})
	return nil
}

func assignSectionType(doc *Doc, s *asciidoc.Section, sectionType matter.Section) {
	name := doc.SectionName(s)
	var ignore bool
	switch sectionType {
	case matter.SectionDataTypeBitmap:
		ignore = doc.errata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesBitmap)
	case matter.SectionDataTypeEnum:
		ignore = doc.errata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesEnum)
	case matter.SectionDataTypeStruct:
		ignore = doc.errata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesStruct)
	case matter.SectionDataTypeDef:
		ignore = doc.errata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesDef)
	case matter.SectionCluster:
		ignore = doc.errata.Spec.IgnoreSection(name, errata.SpecPurposeCluster)
	case matter.SectionDeviceType:
		ignore = doc.errata.Spec.IgnoreSection(name, errata.SpecPurposeDeviceType)
	case matter.SectionFeatures:
		ignore = doc.errata.Spec.IgnoreSection(name, errata.SpecPurposeFeatures)
	}
	if ignore {
		return
	}
	doc.SetSectionType(s, sectionType)
}

func FindSectionByType(doc *Doc, top *asciidoc.Section, sectionType matter.Section) *asciidoc.Section {
	var found *asciidoc.Section
	parse.Search(doc.Reader(), top, top.Children(), func(el *asciidoc.Section, parent asciidoc.Parent, index int) parse.SearchShould {
		if doc.SectionType(el) == sectionType {
			found = el
			return parse.SearchShouldStop
		}
		return parse.SearchShouldContinue
	})
	return found
}

func getSectionType(doc *Doc, parent *asciidoc.Section, section *asciidoc.Section) matter.Section {
	name := strings.ToLower(strings.TrimSpace(doc.SectionName(section)))
	// Names that are unambiguous, so parent section doesn't matter
	switch name {
	case "introduction":
		return matter.SectionIntroduction
	case "revision history":
		return matter.SectionRevisionHistory
	case "classification":
		return matter.SectionClassification
	case "cluster requirements":
		return matter.SectionClusterRequirements
	case "cluster restrictions":
		return matter.SectionClusterRestrictions
	case "element requirements":
		return matter.SectionElementRequirements
	case "endpoint composition":
		return matter.SectionEndpointComposition
	case "derived cluster namespace":
		return matter.SectionDerivedClusterNamespace
	case "global elements":
		return matter.SectionGlobalElements
	case "device type requirements":
		return matter.SectionDeviceTypeRequirements
	case "cluster requirements on composing device types":
		// This is for backwards compatibility; we should have named this
		// section "Element Requirements on Composing Device Types", so
		// we check if it has an Element field in the table, and if it does,
		// it's actually element requirements
		st := guessDataTypeFromTable(doc, section)
		if st == matter.SectionComposedDeviceTypeElementRequirements {
			return st
		}
		return matter.SectionComposedDeviceTypeClusterRequirements
	case "condition requirements on composing device types":
		return matter.SectionComposedDeviceTypeConditionRequirements
	case "element requirements on composing device types":
		return matter.SectionComposedDeviceTypeElementRequirements
	case "semantic tag requirements on composing device types":
		return matter.SectionSemanticTagRequirements
	}
	switch doc.SectionType(parent) {
	case matter.SectionTop, matter.SectionCluster, matter.SectionDeviceType:
		switch name {
		case "cluster identifiers", "cluster id", "cluster ids":
			return matter.SectionClusterID
		case "features":
			return matter.SectionFeatures
		case "dependencies":
			return matter.SectionDependencies
		case "data types":
			return matter.SectionDataTypes
		case "status codes":
			return matter.SectionStatusCodes
		case "attributes":
			return matter.SectionAttributes
		case "commands":
			return matter.SectionCommands
		case "events":
			return matter.SectionEvents
		case "conditions":
			return matter.SectionConditions
		default:
			if strings.HasSuffix(name, " attribute set") {
				return matter.SectionAttributes
			}
			return deriveSectionType(doc, section, parent)
		}
	case matter.SectionDerivedClusterNamespace:
		switch name {
		case "mode tags":
			return matter.SectionModeTags
		default:
			return deriveSectionType(doc, section, parent)
		}
	case matter.SectionAttributes:
		if strings.HasSuffix(name, " attribute") {
			return matter.SectionAttribute
		}
	case matter.SectionFeatures:
		if strings.HasSuffix(name, " feature") {
			return matter.SectionFeature
		}
	case matter.SectionDataTypes:
		if strings.HasSuffix(name, "bitmap type") || strings.HasSuffix(name, "bitmap") {
			return matter.SectionDataTypeBitmap
		}
		if strings.HasSuffix(name, "enum type") || strings.HasSuffix(name, "enum") {
			return matter.SectionDataTypeEnum
		}
		if strings.HasSuffix(name, "struct type") || strings.HasSuffix(name, "struct") || strings.HasSuffix(name, "structure") {
			return matter.SectionDataTypeStruct
		}

		idAttribute := section.GetAttributeByName(asciidoc.AttributeNameID)
		if idAttribute != nil {
			name = strings.ToLower(idAttribute.AsciiDocString())
		}
		if strings.HasSuffix(name, "bitmap") {
			return matter.SectionDataTypeBitmap
		}
		if strings.HasSuffix(name, "enum") {
			return matter.SectionDataTypeEnum
		}
		if strings.HasSuffix(name, "struct") {
			return matter.SectionDataTypeStruct
		}
		return deriveSectionType(doc, section, parent)
	case matter.SectionCommand, matter.SectionDataTypeStruct:
		if strings.HasSuffix(name, " field") {
			return matter.SectionField
		}
		return deriveSectionType(doc, section, parent)
	case matter.SectionCommands:
		if strings.HasSuffix(name, " command") {
			return matter.SectionCommand
		}
	case matter.SectionEvents:
		if strings.HasSuffix(name, " event") {
			return matter.SectionEvent
		}
	case matter.SectionClusterRequirements:
		switch name {
		case "element requirements":
			return matter.SectionElementRequirements
		}
	default:
		return deriveSectionType(doc, section, parent)
	}
	return matter.SectionUnknown
}

func deriveSectionType(doc *Doc, section *asciidoc.Section, parent *asciidoc.Section) matter.Section {

	// Ugh, some heuristics now
	name := strings.TrimSpace(doc.SectionName(section))
	if strings.HasSuffix(name, "Bitmap Type") || strings.HasSuffix(name, "Bitmap") {
		return matter.SectionDataTypeBitmap
	}
	if strings.HasSuffix(name, "Enum Type") || strings.HasSuffix(name, "Enum") {
		return matter.SectionDataTypeEnum
	}
	if strings.HasSuffix(name, " Command") {
		return matter.SectionCommand
	}
	if strings.HasSuffix(name, "Struct Type") || strings.HasSuffix(name, "Struct") {
		return matter.SectionDataTypeStruct
	}
	if text.HasCaseInsensitiveSuffix(name, " Constant Type") {
		return matter.SectionDataTypeConstant
	}
	if strings.HasSuffix(name, " Conditions") {
		return matter.SectionConditions
	}
	if parent != nil {
		switch doc.SectionType(parent) {
		case matter.SectionDataTypes:
			guessedType := guessDataTypeFromTable(doc, section)
			if guessedType != matter.SectionUnknown {
				return guessedType
			}
		case matter.SectionDataTypeBitmap:
			if strings.HasSuffix(name, " Bit") || strings.HasSuffix(name, " Bits") {
				return matter.SectionBit
			}
		case matter.SectionDataTypeEnum:
			if strings.HasSuffix(name, " Value") {
				return matter.SectionValue
			}
		case matter.SectionDataTypeStruct, matter.SectionCommand, matter.SectionEvent:
			if strings.HasSuffix(name, " Field") {
				return matter.SectionField
			}
		}
	}
	dataType := GetDataType(doc, section)
	if dataType != nil {
		if dataType.IsEnum() {
			return matter.SectionDataTypeEnum
		} else if dataType.IsMap() {
			return matter.SectionDataTypeBitmap
		} else if dataType.BaseType == types.BaseDataTypeCustom {
			return matter.SectionDataTypeStruct
		} else if dataType.BaseType.IsSimple() {
			return matter.SectionDataTypeDef
		}
	}
	slog.Debug("unknown section type", "path", doc.Path, "name", name)
	return guessDataTypeFromTable(doc, section)
}

func guessDataTypeFromTable(doc *Doc, section *asciidoc.Section) (sectionType matter.Section) {
	firstTable := FindFirstTable(doc.Reader(), section)
	if firstTable == nil {
		return
	}
	ti, err := ReadTable(doc, doc.Reader(), firstTable)
	if err != nil {
		return
	}
	_, hasName := ti.ColumnIndex(matter.TableColumnName)
	if !hasName {
		return
	}
	_, hasID := ti.ColumnIndex(matter.TableColumnID)
	_, hasType := ti.ColumnIndex(matter.TableColumnType)
	_, hasBit := ti.ColumnIndex(matter.TableColumnBit)
	_, hasValue := ti.ColumnIndex(matter.TableColumnValue)
	_, hasSummary := ti.ColumnIndex(matter.TableColumnSummary)
	_, hasDescription := ti.ColumnIndex(matter.TableColumnDescription)
	_, hasStatusCode := ti.ColumnIndex(matter.TableColumnStatusCode)
	_, hasDeviceId := ti.ColumnIndex(matter.TableColumnDeviceID)
	_, hasClusterId := ti.ColumnIndex(matter.TableColumnClusterID)
	_, hasElement := ti.ColumnIndex(matter.TableColumnElement)
	if hasDeviceId && hasClusterId && hasElement {
		sectionType = matter.SectionComposedDeviceTypeElementRequirements
		return
	}
	if hasID && hasType && !hasBit && !hasValue {
		sectionType = matter.SectionDataTypeStruct
		return
	}
	if !hasSummary && !hasDescription {
		if !hasValue && hasStatusCode {
			sectionType = matter.SectionStatusCodes
			return
		}
		return
	}
	if hasBit {
		sectionType = matter.SectionDataTypeBitmap
		return
	}
	if hasValue || hasStatusCode {
		sectionType = matter.SectionDataTypeEnum
	}
	return
}

func toEntities(spec *Specification, d *Doc, s *asciidoc.Section, pc *parseContext) (err error) {
	switch d.SectionType(s) {
	case matter.SectionCluster:
		err = toClusters(spec, d, s, pc)
		if err != nil {
			return
		}
	case matter.SectionDeviceType:
		err = toDeviceTypes(spec, d, s, pc)
		if err != nil {
			return
		}
	case matter.SectionNamespace:
		err = toNamespace(spec, d, s, pc)
		if err != nil {
			return
		}
	default:
		var looseEntities []types.Entity
		looseEntities, err = findLooseEntities(spec, d, s, pc, nil)
		if err != nil {
			err = newGenericParseError(s, "error reading section \"%s\": %w", d.SectionName(s), err)
			return
		}
		if len(looseEntities) > 0 {
			pc.entities = append(pc.entities, looseEntities...)
		}
	}
	return
}

var dataTypeDefinitionPattern = regexp.MustCompile(`(?:(?:This\s+data\s+type\s+SHALL\s+be\s+a)|(?:is\s+derived\s+from))\s+(?:<<enum-def\s*,\s*)?(enum8|enum16|enum32|map8|map16|map32|uint8|uint16|uint24|uint32|uint40|uint48|uint56|uint64|int8|int16|int24|int32|int40|int48|int56|int64|string)(?:\s*>>)?`)

func GetDataType(doc *Doc, s *asciidoc.Section) *types.DataType {
	var dts string
	for el := range doc.Reader().Iterate(s, s.Children()) {
		switch el := el.(type) {
		case *asciidoc.EmptyLine:
		case *asciidoc.String:
			match := dataTypeDefinitionPattern.FindStringSubmatch(el.Value)
			if match != nil {
				dts = match[1]
				break
			}
			if strings.HasPrefix(el.Value, "This struct") {
				dts = text.TrimCaseInsensitiveSuffix(doc.SectionName(s), " Type")
			}
		case *asciidoc.CrossReference:
			crID := doc.anchorId(doc.Reader(), el, el, el.ID)
			switch crID {
			case "ref_DataTypeBitmap", "ref_DataTypeEnum":
				label := asciidoc.ValueToString(el.Children())
				if len(label) == 0 {
					continue
				}
				label = strings.TrimSpace(label)
				dts = label
			}
		default:
			if el.Type() == asciidoc.ElementTypeBlock {
				break
			}
		}
		if len(dts) > 0 {
			break
		}
	}
	if len(dts) > 0 {
		return types.ParseDataType(dts, false)
	}
	return nil
}

func findLooseEntities(spec *Specification, doc *Doc, section *asciidoc.Section, pc *parseContext, parentEntity types.Entity) (entities []types.Entity, err error) {
	traverseSections(doc, section, errata.SpecPurposeDataTypes, func(section *asciidoc.Section, parent asciidoc.Parent, index int) parse.SearchShould {
		switch doc.SectionType(section) {
		case matter.SectionDataTypeBitmap:
			var bm *matter.Bitmap
			bm, err = toBitmap(doc, section, pc, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to bitmap", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, bm)
			}
		case matter.SectionDataTypeEnum:
			var e *matter.Enum
			e, err = toEnum(doc, section, pc, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to enum", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, e)
			}
		case matter.SectionDataTypeStruct:
			var s *matter.Struct
			s, err = toStruct(spec, doc, section, pc, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to struct", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, s)
			}
		case matter.SectionDataTypeDef:
			var t *matter.TypeDef
			t, err = toTypeDef(doc, section, pc, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to typedef", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, t)
			}
		case matter.SectionGlobalElements:
			var ges []types.Entity
			ges, err = toGlobalElements(spec, doc, section, pc, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to global entities", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, ges...)
			}
		case matter.SectionStatusCodes:
			var me *matter.Enum
			me, err = toEnum(doc, section, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || doc.parsed {
					slog.Warn("Error converting section to status code", log.Element("source", doc.Path, section), slog.Any("error", err))
				}
				err = nil
			} else {
				entities = append(entities, me)
			}
		}
		if err != nil {
			return parse.SearchShouldStop
		}
		return parse.SearchShouldContinue
	})
	return
}

func traverseSections(doc *Doc, parent asciidoc.ParentElement, purpose errata.SpecPurpose, callback parse.ElementSearchCallback[*asciidoc.Section]) (sections []*asciidoc.Section) {
	parse.Search(doc.Reader(), doc, parent.Children(), func(s *asciidoc.Section, parent asciidoc.Parent, index int) parse.SearchShould {
		if doc.errata.Spec.IgnoreSection(doc.SectionName(s), purpose) {
			return parse.SearchShouldContinue
		}
		return callback(s, parent, index)
	})
	return
}
