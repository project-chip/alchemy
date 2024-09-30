package spec

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Section struct {
	Doc  *Doc
	Name string

	Parent any
	Base   *asciidoc.Section

	SecType              matter.Section
	sectionTypesAssigned bool

	asciidoc.Set
}

func NewSection(doc *Doc, parent any, s *asciidoc.Section) (*Section, error) {
	ss := &Section{Doc: doc, Parent: parent, Base: s}

	var name strings.Builder
	err := buildSectionTitle(doc, &name, s.Title...)
	if err != nil {
		return nil, err
	}
	ss.Name = name.String()
	for _, e := range s.Elements() {
		switch el := e.(type) {
		case *asciidoc.AttributeEntry:
			doc.attributes[el.Name] = el.Elements()
			ss.Append(NewElement(ss, e))
		case *asciidoc.Section:
			s, err := NewSection(doc, ss, el)
			if err != nil {
				return nil, err
			}
			ss.Append(s)
		default:
			ss.Append(NewElement(ss, e))
		}
	}
	return ss, nil
}

func buildSectionTitle(doc *Doc, title *strings.Builder, els ...asciidoc.Element) (err error) {
	for _, e := range els {
		switch e := e.(type) {
		case *asciidoc.String:
			title.WriteString(e.Value)
		case *asciidoc.SpecialCharacter:
			title.WriteString(e.Character)
		case *asciidoc.UserAttributeReference:

			attr, ok := doc.attributes[asciidoc.AttributeName(e.Name())]
			if !ok {
				title.WriteRune('{')
				title.WriteString(e.Name())
				title.WriteRune('}')
				slog.Warn("unknown section title attribute", "name", e.Name())
				return
			}
			switch val := attr.(type) {
			case asciidoc.Set:
				err = buildSectionTitle(doc, title, val...)
			case asciidoc.Element:
				err = buildSectionTitle(doc, title, val)
			default:
				err = fmt.Errorf("unexpected section title attribute value type: %T", attr)
			}
		case *asciidoc.Link, *asciidoc.LinkMacro:
		case *asciidoc.Bold:
			err = buildSectionTitle(doc, title, e.Elements()...)
		default:
			err = fmt.Errorf("unknown section title component type: %T", e)
		}
		if err != nil {
			return
		}
		if he, ok := e.(asciidoc.HasElements); ok {
			err = buildSectionTitle(doc, title, he.Elements()...)
		}
		if err != nil {
			return
		}
	}
	return
}

func (s *Section) AppendSection(ns *Section) error {
	return s.Set.Append(ns)
}

func (s *Section) SetElements(elements asciidoc.Set) error {
	s.Set.SetElements(elements)
	return s.Base.SetElements(elements)
}

func (s Section) Type() asciidoc.ElementType {
	return asciidoc.ElementTypeBlock
}

func (e *Section) Equals(o asciidoc.Element) bool {
	return e.Base.Equals(o)
}

func (s *Section) GetASCIISection() *asciidoc.Section {
	return s.Base
}

func AssignSectionTypes(doc *Doc, top *Section) error {
	if top.sectionTypesAssigned {
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
		if strings.HasSuffix(top.Name, " Cluster") {
			secType = matter.SectionCluster
		}
	}
	assignSectionType(doc, top, secType)

	parse.Traverse(top, top.Elements(), func(el any, parent parse.HasElements, index int) parse.SearchShould {
		section, ok := el.(*Section)
		if !ok {
			return parse.SearchShouldContinue
		}
		ps, ok := parent.(*Section)
		if !ok {
			return parse.SearchShouldContinue
		}

		assignSectionType(doc, section, getSectionType(ps, section))
		switch section.SecType {
		case matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct, matter.SectionDataTypeDef:
			if section.Base.Level > 2 {
				slog.Debug("Unusual depth for section type", slog.String("name", section.Name), slog.String("type", section.SecType.String()), slog.String("path", doc.Path.String()))
			}
		}
		slog.Debug("sec type", "name", section.Name, "type", section.SecType, "parent", ps.Name, "parentType", ps.SecType)
		return parse.SearchShouldContinue
	})
	top.sectionTypesAssigned = true
	return nil
}

func assignSectionType(doc *Doc, s *Section, sectionType matter.Section) {
	var ignore bool
	switch sectionType {
	case matter.SectionDataTypeBitmap:
		ignore = doc.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeDataTypesBitmap)
	case matter.SectionDataTypeEnum:
		ignore = doc.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeDataTypesEnum)
	case matter.SectionDataTypeStruct:
		ignore = doc.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeDataTypesStruct)
	case matter.SectionDataTypeDef:
		ignore = doc.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeDataTypesDef)
	case matter.SectionCluster:
		ignore = doc.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeCluster)
	case matter.SectionDeviceType:
		ignore = doc.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeDeviceType)
	}
	if ignore {
		return
	}
	s.SecType = sectionType
}

func FindSectionByType(top *Section, sectionType matter.Section) *Section {
	var found *Section
	parse.Search(top.Elements(), func(s *Section) parse.SearchShould {
		if s.SecType == sectionType {
			found = s
			return parse.SearchShouldStop
		}
		return parse.SearchShouldContinue
	})
	return found
}

func getSectionType(parent *Section, section *Section) matter.Section {
	name := strings.ToLower(strings.TrimSpace(section.Name))
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
	}
	switch parent.SecType {
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
			return deriveSectionType(section, parent)
		}
	case matter.SectionDerivedClusterNamespace:
		switch name {
		case "mode tags":
			return matter.SectionModeTags
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

		idAttribute := section.Base.GetAttributeByName(asciidoc.AttributeNameID)
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
		return deriveSectionType(section, parent)
	case matter.SectionCommand, matter.SectionDataTypeStruct:
		if strings.HasSuffix(name, " field") {
			return matter.SectionField
		}
		return deriveSectionType(section, parent)
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
		return deriveSectionType(section, parent)
	}
	return matter.SectionUnknown
}

func deriveSectionType(section *Section, parent *Section) matter.Section {

	// Ugh, some heuristics now
	name := strings.TrimSpace(section.Name)
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
	if strings.HasSuffix(name, " Conditions") {
		return matter.SectionConditions
	}
	if parent != nil {
		switch parent.SecType {
		case matter.SectionDataTypes:
			guessedType := guessDataTypeFromTable(section)
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
	dataType := section.GetDataType()
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
	slog.Debug("unknown section type", "path", section.Doc.Path, "name", name)
	return guessDataTypeFromTable(section)
}

func guessDataTypeFromTable(section *Section) (sectionType matter.Section) {
	firstTable := FindFirstTable(section)
	if firstTable == nil {
		return
	}
	ti, err := ReadTable(section.Doc, firstTable)
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
		sectionType = matter.SectionComposedDeviceTypeRequirements
		return
	}
	if hasID && hasType && !hasBit && !hasValue {
		sectionType = matter.SectionDataTypeStruct
		return
	}
	if !hasSummary && !hasDescription {
		return
	}
	if hasBit {
		sectionType = matter.SectionDataTypeBitmap
	}
	if hasValue || hasStatusCode {
		sectionType = matter.SectionDataTypeEnum
	}
	return
}

func (s *Section) toEntities(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) ([]types.Entity, error) {
	var entities []types.Entity
	switch s.SecType {
	case matter.SectionCluster:
		clusters, err := s.toClusters(d, entityMap)
		if err != nil {
			return nil, err
		}
		entities = append(entities, clusters...)
	case matter.SectionDeviceType:
		deviceTypes, err := s.toDeviceTypes(d)
		if err != nil {
			return nil, err
		}
		entities = append(entities, deviceTypes...)
	case matter.SectionNamespace:
		ns, err := s.toNamespace(d, entityMap)
		if err != nil {
			return nil, err
		}
		entities = append(entities, ns...)
	}
	return entities, nil
}

func (s *Section) toGlobalObjects(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) ([]types.Entity, error) {
	var entities []types.Entity
	switch s.SecType {
	case matter.SectionCluster, matter.SectionDeviceType:
		return nil, nil
	default:
		var err error
		var looseEntities []types.Entity
		looseEntities, err = findLooseEntities(d, s, entityMap)
		if err != nil {
			return nil, fmt.Errorf("error reading section %s: %w", s.Name, err)
		}
		if len(looseEntities) > 0 {
			entities = append(entities, looseEntities...)
		}
	}
	return entities, nil
}

var dataTypeDefinitionPattern = regexp.MustCompile(`is\s+derived\s+from\s+(?:<<enum-def\s*,\s*)?(enum8|enum16|enum32|map8|map16|map32|uint8|uint16|uint24|uint32|uint40|uint48|uint56|uint64|int8|int16|int24|int32|int40|int48|int56|int64|string)(?:\s*>>)?`)

func (s *Section) GetDataType() *types.DataType {
	var dts string
	for _, el := range s.Elements() {
		if se, ok := el.(*Element); ok {
			el = se.Base
		}
		switch el := el.(type) {
		case asciidoc.EmptyLine:
		case *asciidoc.String:
			match := dataTypeDefinitionPattern.FindStringSubmatch(el.Value)
			if match != nil {
				dts = match[1]
				break
			}
			if strings.HasPrefix(el.Value, "This struct") {
				dts = text.TrimCaseInsensitiveSuffix(s.Name, " Type")
			}
		case *asciidoc.CrossReference:
			switch el.ID {
			case "ref_DataTypeBitmap", "ref_DataTypeEnum":
				label := asciidoc.ValueToString(el.Elements())
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

func findLooseEntities(doc *Doc, section *Section, entityMap map[asciidoc.Attributable][]types.Entity) (entities []types.Entity, err error) {
	traverse(doc, section, errata.SpecPurposeDataTypes, func(section *Section, parent parse.HasElements, index int) parse.SearchShould {
		switch section.SecType {
		case matter.SectionDataTypeBitmap:
			var bm *matter.Bitmap
			bm, err = section.toBitmap(doc, entityMap)
			if err != nil {
				slog.Warn("Error converting loose section to bitmap", log.Element("path", doc.Path, section.Base), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, bm)
			}
		case matter.SectionDataTypeEnum:
			var e *matter.Enum
			e, err = section.toEnum(doc, entityMap)
			if err != nil {
				slog.Warn("Error converting loose section to enum", log.Element("path", doc.Path, section.Base), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, e)
			}
		case matter.SectionDataTypeStruct:
			var s *matter.Struct
			s, err = section.toStruct(doc, entityMap)
			if err != nil {
				slog.Warn("Error converting loose section to struct", log.Element("path", doc.Path, section.Base), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, s)
			}
		case matter.SectionDataTypeDef:
			var t *matter.TypeDef
			t, err = section.toTypeDef(doc, entityMap)
			if err != nil {
				slog.Warn("Error converting loose section to typedef", log.Element("path", doc.Path, section.Base), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, t)
			}
		case matter.SectionGlobalElements:
			var ges []types.Entity
			ges, err = section.toGlobalElements(doc, entityMap)
			if err != nil {
				slog.Warn("Error converting loose section to global entities", log.Element("path", doc.Path, section.Base), slog.Any("error", err))
				err = nil
			} else {
				entities = append(entities, ges...)
			}
		}
		if err != nil {
			return parse.SearchShouldStop
		}
		return parse.SearchShouldContinue
	})
	return
}

func skim(doc *Doc, parent asciidoc.HasElements, purpose errata.SpecPurpose) (sections []*Section) {
	for _, s := range parse.Skim[*Section](parent.Elements()) {
		if doc.errata.Spec.IgnoreSection(s.Name, purpose) {
			continue
		}
		sections = append(sections, s)
	}
	return
}

func traverse(doc *Doc, parent asciidoc.HasElements, purpose errata.SpecPurpose, callback parse.TraverseCallback[*Section]) (sections []*Section) {
	parse.Traverse(doc, parent.Elements(), func(s *Section, parent parse.HasElements, index int) parse.SearchShould {
		if doc.errata.Spec.IgnoreSection(s.Name, purpose) {
			return parse.SearchShouldContinue
		}
		return callback(s, parent, index)
	})
	return
}
