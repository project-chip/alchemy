package ascii

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

type Section struct {
	Doc  *Doc
	Name string

	Parent any
	Base   *elements.Section

	SecType matter.Section

	Elements []elements.Element
}

func NewSection(doc *Doc, parent any, s *elements.Section) (*Section, error) {
	ss := &Section{Doc: doc, Parent: parent, Base: s}

	var name strings.Builder
	err := buildSectionTitle(doc, &name, s.Title)
	if err != nil {
		return nil, err
	}
	ss.Name = name.String()
	for _, e := range s.Elements {
		switch el := e.(type) {
		case *elements.AttributeDeclaration:
			doc.attributes[el.Name] = el.Value
			ss.Elements = append(ss.Elements, NewElement(ss, e))
		case *elements.Section:
			s, err := NewSection(doc, ss, el)
			if err != nil {
				return nil, err
			}
			ss.Elements = append(ss.Elements, s)
		default:
			ss.Elements = append(ss.Elements, NewElement(ss, e))
		}
	}
	return ss, nil
}

func buildSectionTitle(doc *Doc, title *strings.Builder, els ...any) (err error) {
	for _, e := range els {
		switch e := e.(type) {
		case []any:
			err = buildSectionTitle(doc, title, e...)
		case string:
			title.WriteString(e)
		case elements.String:
			title.WriteString(string(e))
		case *elements.SpecialCharacter:
			title.WriteString(e.Character)
		case *elements.AttributeReference:
			attr, ok := doc.attributes[e.Name]
			if !ok {
				return fmt.Errorf("unknown section title attribute: %s", e.Name)
			}
			err = buildSectionTitle(doc, title, attr)
		case *elements.Link:
		case *elements.Bold:

		default:
			err = fmt.Errorf("unknown section title component type: %T", e)
		}
		if err != nil {
			return
		}
	}
	return
}

func (s *Section) AppendSection(ns *Section) error {
	s.Elements = append(s.Elements, ns)
	return nil
}

func (s *Section) GetElements() []elements.Element {
	return s.Elements
}

func (s *Section) SetElements(elements []elements.Element) error {
	s.Elements = elements
	return s.Base.SetElements(elements)
}

func (s Section) Type() elements.ElementType {
	return elements.ElementTypeBlock
}

func (s *Section) GetASCIISection() *elements.Section {
	return s.Base
}

func AssignSectionTypes(doc *Doc, top *Section) error {
	docType, err := doc.DocType()
	if err != nil {
		return err
	}
	switch docType {
	case matter.DocTypeCluster:
		top.SecType = matter.SectionCluster
	case matter.DocTypeDeviceType:
		top.SecType = matter.SectionDeviceType
	default:
		top.SecType = matter.SectionTop
		if strings.HasSuffix(top.Name, " Cluster") {
			top.SecType = matter.SectionCluster
		}
	}

	parse.Traverse(top, top.Elements, func(el any, parent parse.HasElements, index int) bool {
		section, ok := el.(*Section)
		if !ok {
			return false
		}
		ps, ok := parent.(*Section)
		if !ok {
			return false
		}

		section.SecType = getSectionType(ps, section)
		switch section.SecType {
		case matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct:
			if section.Base.Level > 2 {
				slog.Warn("Unusual depth for section type", slog.String("name", section.Name), slog.String("type", section.SecType.String()), slog.String("path", doc.Path))
			}
		}
		slog.Debug("sec type", "name", section.Name, "type", section.SecType, "parent", ps.SecType)
		return false
	})
	return nil
}

func FindSectionByType(top *Section, sectionType matter.Section) *Section {
	var found *Section
	parse.Search(top.Elements, func(s *Section) bool {
		if s.SecType == sectionType {
			found = s
			return true
		}
		return false
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
			return matter.SectionUnknown
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
		name = strings.ToLower(section.Base.GetID())
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
	if strings.HasSuffix(name, "Struct Type") || strings.HasSuffix(name, "Struct") {
		return matter.SectionDataTypeStruct
	}
	if strings.HasSuffix(name, " Conditions") {
		return matter.SectionConditions
	}
	if parent != nil {
		switch parent.SecType {
		case matter.SectionDataTypes:
			guessedType := guessDataTypeFromTable(section, parent)
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
		} else if dataType.BaseType == mattertypes.BaseDataTypeCustom {
			return matter.SectionDataTypeStruct
		}
	}
	slog.Debug("unknown section type", "path", section.Doc.Path, "name", name)
	return matter.SectionUnknown
}

func guessDataTypeFromTable(section *Section, parent *Section) (sectionType matter.Section) {
	firstTable := FindFirstTable(section)
	if firstTable == nil {
		return
	}
	rows := TableRows(firstTable)
	_, columnMap, _, err := MapTableColumns(section.Doc, rows)
	if err != nil {
		return
	}
	_, hasName := columnMap[matter.TableColumnName]
	if !hasName {
		return
	}
	_, hasID := columnMap[matter.TableColumnID]
	_, hasType := columnMap[matter.TableColumnType]
	if hasID && hasType {
		sectionType = matter.SectionDataTypeStruct
		return
	}
	_, hasSummary := columnMap[matter.TableColumnSummary]
	if !hasSummary {
		return
	}
	_, hasBit := columnMap[matter.TableColumnBit]
	_, hasValue := columnMap[matter.TableColumnValue]
	if hasBit {
		sectionType = matter.SectionDataTypeBitmap
	}
	if hasValue {
		sectionType = matter.SectionDataTypeEnum
	}
	return
}

func (s *Section) toEntities(d *Doc, entityMap map[elements.Attributable][]mattertypes.Entity) ([]mattertypes.Entity, error) {
	var entities []mattertypes.Entity
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
	default:
		var err error
		var looseEntities []mattertypes.Entity
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

var dataTypeDefinitionPattern = regexp.MustCompile(`is\s+derived\s+from\s+(?:<<enum-def\s*,\s*)?(enum8|enum16|enum32|map8|map16|map32)(?:\s*>>)?`)

//var dataTypeDefinitionPattern = regexp.MustCompile(`is\s+derived\s+from\s+(?:(?:<<enum-def\s*,\s*)|(?:<<ref_DataTypeBitmap,))?(enum8|enum16|enum32|map8|map16|map32)(?:\s*>>)?`)

func (s *Section) GetDataType() *mattertypes.DataType {
	var para *elements.Paragraph
	var ok bool
	for _, e := range s.Base.Elements {
		para, ok = e.(*elements.Paragraph)
		if ok {
			break
		}
	}
	if !ok {
		return nil
	}
	var dts string
	for _, el := range para.Elements {
		switch el := el.(type) {
		case *elements.String:
			match := dataTypeDefinitionPattern.FindStringSubmatch(el.Content)
			if match != nil {
				dts = match[1]
				break
			}
			if strings.HasPrefix(el.Content, "This struct") {
				dts = strings.TrimSuffix(s.Name, " Type")
			}
		case *elements.InternalCrossReference:
			id, ok := el.ID.(string)
			if !ok {
				continue
			}
			switch id {
			case "ref_DataTypeBitmap", "ref_DataTypeEnum":
				label, ok := el.Label.(string)
				if !ok {
					continue
				}
				label = strings.TrimSpace(label)
				dts = label
			}
		}
		if len(dts) > 0 {
			break
		}
	}
	if len(dts) > 0 {
		return mattertypes.ParseDataType(dts, false)
	}
	return nil
}

func findLooseEntities(doc *Doc, section *Section, entityMap map[elements.Attributable][]mattertypes.Entity) (entities []mattertypes.Entity, err error) {
	parse.Traverse(doc, section.Elements, func(section *Section, parent parse.HasElements, index int) bool {
		switch section.SecType {
		case matter.SectionDataTypeBitmap:
			var bm *matter.Bitmap
			bm, err = section.toBitmap(doc, entityMap)
			if err == nil {
				entities = append(entities, bm)
			}
		case matter.SectionDataTypeEnum:
			var e *matter.Enum
			e, err = section.toEnum(doc, entityMap)
			if err == nil {
				entities = append(entities, e)
			}
		case matter.SectionDataTypeStruct:
			var s *matter.Struct
			s, err = section.toStruct(doc, entityMap)
			if err == nil {
				entities = append(entities, s)
			}
		}
		return err != nil
	})
	return
}
