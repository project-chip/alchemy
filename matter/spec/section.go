package spec

import (
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/overlay"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type SectionInfoCache interface {
	SectionName(section *asciidoc.Section) string
	SetSectionName(section *asciidoc.Section, name string)
	SectionType(section *asciidoc.Section) matter.Section
	SetSectionType(section *asciidoc.Section, sectionType matter.Section)

	ErrataForPath(docPath string) *errata.Errata
}

func buildSectionTitle(variables overlay.Variables, section *asciidoc.Section, reader asciidoc.Reader, title *strings.Builder, els ...asciidoc.Element) (err error) {
	for e := range reader.Iterate(section, els) {
		switch e := e.(type) {
		case *asciidoc.String:
			title.WriteString(e.Value)
		case *asciidoc.SpecialCharacter:
			title.WriteString(e.Character)
		case *asciidoc.UserAttributeReference:

			val := variables.Get(string(e.Name()))
			if val != nil {
				switch val := val.(type) {
				case asciidoc.Elements:
					err = buildSectionTitle(variables, section, reader, title, val...)
				case asciidoc.Element:
					err = buildSectionTitle(variables, section, reader, title, val)
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
			err = buildSectionTitle(variables, section, reader, title, reader.Children(e)...)
		default:
			err = newGenericParseError(section, "unknown section title component type: %T", e)
		}
		if err != nil {
			return
		}
		if he, ok := e.(asciidoc.ParentElement); ok {
			err = buildSectionTitle(variables, section, reader, title, reader.Children(he)...)
		}
		if err != nil {
			return
		}
	}
	return
}

type emptyVariableStore struct {
}

func (emptyVariableStore) Get(string) any {
	return nil
}

func (emptyVariableStore) IsSet(name string) bool {
	return false
}
func (emptyVariableStore) Set(name string, value any) {}
func (emptyVariableStore) Unset(name string)          {}

func AssignSectionNames(sectionInfoCache SectionInfoCache, reader asciidoc.Reader, doc *asciidoc.Document) error {
	vs := &emptyVariableStore{}
	parse.Search(doc, reader, doc, reader.Children(doc), func(doc *asciidoc.Document, section *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		var sb strings.Builder
		buildSectionTitle(vs, section, reader, &sb, section.Title...)
		slog.Info("assigning name", "name", sb.String())
		sectionInfoCache.SetSectionName(section, sb.String())
		return parse.SearchShouldContinue
	})
	return nil
}

func AssignSectionTypes(sectionInfoCache SectionInfoCache, docInfoCache DocumentInfoCache, reader asciidoc.Reader, doc *asciidoc.Document, top *asciidoc.Section) (err error) {
	//slog.Info("Assigning section type", "path", doc.Path.Relative, log.Path("source", top))
	var lastDoc *asciidoc.Document
	parse.Search(doc, reader, doc, reader.Children(doc), func(doc *asciidoc.Document, section *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		//slog.Info("Searching section type", "section", sectionInfoCache.SectionName(section), "path", doc.Path.Relative, log.Path("source", top))

		parentSection, _ := parent.(*asciidoc.Section)
		if doc != lastDoc {
			// This is the top section of a document
			var docType matter.DocType
			docType, err = docInfoCache.DocType(doc)
			if err != nil {
				return parse.SearchShouldStop
			}
			var parentSectionType matter.Section
			if parentSection != nil {
				parentSectionType = sectionInfoCache.SectionType(parentSection)
			}

			lastDoc = doc

			var secType matter.Section
			switch parentSectionType {
			case matter.SectionTop, matter.SectionUnknown:
				switch docType {
				case matter.DocTypeCluster:
					secType = matter.SectionCluster
				case matter.DocTypeDeviceType, matter.DocTypeBaseDeviceType:
					secType = matter.SectionDeviceType
				case matter.DocTypeNamespace:
					secType = matter.SectionNamespace

				default:
					secType = matter.SectionTop
					if strings.HasSuffix(sectionInfoCache.SectionName(section), " Cluster") {
						secType = matter.SectionCluster
					}
				}
			}

			if secType != matter.SectionUnknown {
				assignSectionType(sectionInfoCache, doc, section, secType)
				return parse.SearchShouldContinue
			}
		}
		assignSectionType(sectionInfoCache, doc, section, getSectionType(sectionInfoCache, reader, doc, parentSection, section))
		switch sectionInfoCache.SectionType(section) {
		case matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct, matter.SectionDataTypeDef:
			if section.Level > 2 {
				slog.Debug("Unusual depth for section type", slog.String("name", sectionInfoCache.SectionName(section)), slog.String("type", sectionInfoCache.SectionType(section).String()), slog.String("path", doc.Path.String()))
			}
		}
		return parse.SearchShouldContinue
	})
	return
}

func assignSectionType(sectionInfoCache SectionInfoCache, doc *asciidoc.Document, s *asciidoc.Section, sectionType matter.Section) {
	name := sectionInfoCache.SectionName(s)
	docErrata := sectionInfoCache.ErrataForPath(doc.Path.Relative)
	var ignore bool
	switch sectionType {
	case matter.SectionDataTypeBitmap:
		ignore = docErrata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesBitmap)
	case matter.SectionDataTypeEnum:
		ignore = docErrata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesEnum)
	case matter.SectionDataTypeStruct:
		ignore = docErrata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesStruct)
	case matter.SectionDataTypeDef:
		ignore = docErrata.Spec.IgnoreSection(name, errata.SpecPurposeDataTypesDef)
	case matter.SectionCluster:
		ignore = docErrata.Spec.IgnoreSection(name, errata.SpecPurposeCluster)
	case matter.SectionDeviceType:
		ignore = docErrata.Spec.IgnoreSection(name, errata.SpecPurposeDeviceType)
	case matter.SectionFeatures:
		ignore = docErrata.Spec.IgnoreSection(name, errata.SpecPurposeFeatures)
	}
	if ignore {
		return
	}
	sectionInfoCache.SetSectionType(s, sectionType)
}

func (library *Library) FindSectionByType(reader asciidoc.Reader, doc *asciidoc.Document, top *asciidoc.Section, sectionType matter.Section) *asciidoc.Section {
	var found *asciidoc.Section
	parse.Search(doc, reader, top, reader.Children(top), func(doc *asciidoc.Document, el *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		if library.SectionType(el) == sectionType {
			found = el
			return parse.SearchShouldStop
		}
		return parse.SearchShouldContinue
	})
	return found
}

func getSectionType(sectionInfoCache SectionInfoCache, reader asciidoc.Reader, doc *asciidoc.Document, parent *asciidoc.Section, section *asciidoc.Section) matter.Section {
	name := strings.ToLower(strings.TrimSpace(sectionInfoCache.SectionName(section)))
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
	case "cluster requirements on composing device types", "cluster requirements on component device types":
		// This is for backwards compatibility; we should have named this
		// section "Element Requirements on Composing Device Types", so
		// we check if it has an Element field in the table, and if it does,
		// it's actually element requirements
		st := guessDataTypeFromTable(reader, doc, section)
		if st == matter.SectionComposedDeviceTypeElementRequirements {
			return st
		}
		return matter.SectionComposedDeviceTypeClusterRequirements
	case "semantic tag requirements":
		return matter.SectionSemanticTagRequirements
	case "condition requirements":
		return matter.SectionConditionRequirements
	case "element requirements on composing device types", "element requirements on component device types":
		return matter.SectionComposedDeviceTypeElementRequirements
	case "semantic tag requirements on composing device types", "semantic tag requirements on component device types":
		return matter.SectionComposedDeviceTypeSemanticTagRequirements
	}
	switch sectionInfoCache.SectionType(parent) {
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
			return deriveSectionType(sectionInfoCache, reader, doc, section, parent)
		}
	case matter.SectionDerivedClusterNamespace:
		switch name {
		case "mode tags":
			return matter.SectionModeTags
		default:
			return deriveSectionType(sectionInfoCache, reader, doc, section, parent)
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
		return deriveSectionType(sectionInfoCache, reader, doc, section, parent)
	case matter.SectionCommand, matter.SectionDataTypeStruct:
		if strings.HasSuffix(name, " field") {
			return matter.SectionField
		}
		return deriveSectionType(sectionInfoCache, reader, doc, section, parent)
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
		return deriveSectionType(sectionInfoCache, reader, doc, section, parent)
	}
	return matter.SectionUnknown
}

func deriveSectionType(sectionInfoCache SectionInfoCache, reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, parent *asciidoc.Section) matter.Section {

	// Ugh, some heuristics now
	name := strings.TrimSpace(sectionInfoCache.SectionName(section))
	if text.HasCaseInsensitiveSuffix(name, "Bitmap Type") || text.HasCaseInsensitiveSuffix(name, "Bitmap") {
		return matter.SectionDataTypeBitmap
	}
	if text.HasCaseInsensitiveSuffix(name, "Enum Type") || text.HasCaseInsensitiveSuffix(name, "Enum") {
		return matter.SectionDataTypeEnum
	}
	if text.HasCaseInsensitiveSuffix(name, " Command") {
		return matter.SectionCommand
	}
	if text.HasCaseInsensitiveSuffix(name, "Struct Type") || text.HasCaseInsensitiveSuffix(name, "Struct") {
		return matter.SectionDataTypeStruct
	}
	if text.HasCaseInsensitiveSuffix(name, " Constant Type") {
		return matter.SectionDataTypeConstant
	}
	if text.HasCaseInsensitiveSuffix(name, " Conditions") {
		return matter.SectionConditions
	}
	if parent != nil {
		switch sectionInfoCache.SectionType(parent) {
		case matter.SectionDataTypes:
			guessedType := guessDataTypeFromTable(reader, doc, section)
			if guessedType != matter.SectionUnknown {
				return guessedType
			}
		case matter.SectionDataTypeBitmap:
			if text.HasCaseInsensitiveSuffix(name, " Bit") || text.HasCaseInsensitiveSuffix(name, " Bits") {
				return matter.SectionBit
			}
		case matter.SectionDataTypeEnum:
			if text.HasCaseInsensitiveSuffix(name, " Value") {
				return matter.SectionValue
			}
		case matter.SectionDataTypeStruct, matter.SectionCommand, matter.SectionEvent:
			if text.HasCaseInsensitiveSuffix(name, " Field") {
				return matter.SectionField
			}
		}
	}
	dataType, _ := GetDataType(sectionInfoCache, reader, doc, section)
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
	return guessDataTypeFromTable(reader, doc, section)
}

func guessDataTypeFromTable(reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section) (sectionType matter.Section) {
	firstTable := FindFirstTable(reader, section)
	if firstTable == nil {
		return
	}
	ti, err := ReadTable(doc, reader, firstTable)
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

var dataTypeDefinitionPattern = regexp.MustCompile(`(?:(?:This\s+data\s+type\s+SHALL\s+be\s+a)|(?:is\s+derived\s+from))\s+(?:<<enum-def\s*,\s*)?(enum8|enum16|enum32|map8|map16|map32|map64|uint8|uint16|uint24|uint32|uint40|uint48|uint56|uint64|int8|int16|int24|int32|int40|int48|int56|int64|string)(?:\s*>>)?`)

func GetDataType(sectionInfoCache SectionInfoCache, reader asciidoc.Reader, doc *asciidoc.Document, s *asciidoc.Section) (*types.DataType, error) {
	var dts string
	for el := range reader.Iterate(s, reader.Children(s)) {
		switch el := el.(type) {
		case *asciidoc.EmptyLine:
		case *asciidoc.String:
			match := dataTypeDefinitionPattern.FindStringSubmatch(el.Value)
			if match != nil {
				dts = match[1]
				break
			}
			if strings.HasPrefix(el.Value, "This struct") {
				dts = text.TrimCaseInsensitiveSuffix(sectionInfoCache.SectionName(s), " Type")
			}
		case *asciidoc.CrossReference:
			crID, err := reader.StringValue(el, el.ID)
			if err != nil {
				return nil, err
			}
			switch crID {
			case "ref_DataTypeBitmap", "ref_DataTypeEnum":
				label := asciidoc.ValueToString(reader.Children(el))
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
		return types.ParseDataType(dts, types.DataTypeRankScalar), nil
	}
	return nil, nil
}

func (library *Library) findLooseEntities(spec *Specification, reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, parentEntity types.Entity) (entities []types.Entity, err error) {
	library.traverseSections(reader, doc, section, errata.SpecPurposeDataTypes, func(doc *asciidoc.Document, section *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		switch library.SectionType(section) {
		case matter.SectionDataTypeBitmap:
			var bm *matter.Bitmap
			bm, err = library.toBitmap(reader, doc, section, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to bitmap", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				library.addEntity(section, bm)
				entities = append(entities, bm)

			}
		case matter.SectionDataTypeEnum:
			var e *matter.Enum
			e, err = library.toEnum(reader, doc, section, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to enum", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				library.addEntity(section, e)
				entities = append(entities, e)

			}
		case matter.SectionDataTypeStruct:
			var s *matter.Struct
			s, err = library.toStruct(spec, reader, doc, section, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to struct", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				library.addEntity(section, s)
				entities = append(entities, s)

			}
		case matter.SectionDataTypeDef:
			var t *matter.TypeDef
			t, err = library.toTypeDef(reader, doc, section, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to typedef", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				library.addEntity(section, t)
				entities = append(entities, t)

			}
		case matter.SectionGlobalElements:
			var ges []types.Entity
			ges, err = library.toGlobalElements(spec, reader, doc, section, parentEntity)
			if err != nil {
				slog.Warn("Error converting loose section to global entities", log.Element("source", doc.Path, section), slog.Any("error", err))
				err = nil
			} else {
				for _, e := range ges {
					library.addEntity(section, e)

				}
				entities = append(entities, ges...)

			}
		case matter.SectionStatusCodes:
			var me *matter.Enum
			me, err = library.toEnum(reader, doc, section, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable {
					slog.Warn("Error converting section to status code", log.Element("source", doc.Path, section), slog.Any("error", err))
				}
				err = nil
			} else {
				library.addEntity(section, me)
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

func (library *Library) traverseSections(reader asciidoc.Reader, doc *asciidoc.Document, parent asciidoc.ParentElement, purpose errata.SpecPurpose, callback parse.ElementSearchCallback[*asciidoc.Section]) (sections []*asciidoc.Section) {
	parse.Search(doc, reader, doc, reader.Children(parent), func(doc *asciidoc.Document, s *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		de := library.SpecErrata(doc.Path.Relative)
		if de.IgnoreSection(library.SectionName(s), purpose) {
			return parse.SearchShouldContinue
		}
		return callback(doc, s, parent, index)
	})
	return
}
