package ascii

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

type Section struct {
	Name string

	Parent interface{}
	Base   *types.Section

	SecType matter.Section

	Elements []interface{}
}

func NewSection(parent interface{}, s *types.Section) (*Section, error) {
	ss := &Section{Parent: parent, Base: s}

	switch name := types.Reduce(s.Title).(type) {
	case string:
		ss.Name = name
	case []interface{}:
		var complexName strings.Builder
		for _, e := range name {
			switch v := e.(type) {
			case *types.StringElement:
				complexName.WriteString(v.Content)
			case string:
				complexName.WriteString(v)
			case *types.Symbol:
				complexName.WriteString(v.Name)
			case *types.SpecialCharacter:
				complexName.WriteString(v.Name)
			case *types.InlineLink:
			case *types.QuotedText:

			default:
				return nil, fmt.Errorf("unknown section title component type: %T", e)
			}
		}
		ss.Name = complexName.String()
	default:
		return nil, fmt.Errorf("unknown section title type: %T", name)
	}
	for _, e := range s.Elements {
		switch el := e.(type) {
		case *types.Section:
			s, err := NewSection(ss, el)
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

func (s *Section) AppendSection(ns *Section) error {
	s.Elements = append(s.Elements, ns)
	return nil
}

func (s *Section) GetElements() []interface{} {
	return s.Elements
}

func (s *Section) SetElements(elements []interface{}) error {
	s.Elements = elements
	return nil
}

func (s *Section) GetAsciiSection() *types.Section {
	return s.Base
}

func AssignSectionTypes(docType matter.DocType, top *Section) {
	switch docType {
	case matter.DocTypeAppCluster:
		top.SecType = matter.SectionCluster
	case matter.DocTypeDeviceType:
		top.SecType = matter.SectionDeviceType
	default:
		top.SecType = matter.SectionTop
	}

	parse.Traverse(top, top.Elements, func(el interface{}, parent parse.HasElements, index int) bool {
		section, ok := el.(*Section)
		if !ok {
			return false
		}
		ps, ok := parent.(*Section)
		if !ok {
			return false
		}

		section.SecType = getSectionType(ps, section)
		return false
	})
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
	switch parent.SecType {
	case matter.SectionTop, matter.SectionCluster, matter.SectionDeviceType:
		switch name {
		case "introduction":
			return matter.SectionIntroduction
		case "revision history":
			return matter.SectionRevisionHistory
		case "classification":
			return matter.SectionClassification
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
		case "cluster requirements":
			return matter.SectionClusterRequirements
		case "cluster restrictions":
			return matter.SectionClusterRestrictions
		case "element requirements":
			return matter.SectionElementRequirements
		case "endpoint composition":
			return matter.SectionEndpointComposition
		default:
			if strings.HasSuffix(name, " attribute set") {
				return matter.SectionAttributes
			}
			return matter.SectionUnknown
		}
	case matter.SectionAttributes:
		if strings.HasSuffix(name, " attribute") {
			return matter.SectionAttribute
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
	case matter.SectionCommand:
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
	default:
		//slog.Info("unknown section type", "name", name)
		// Ugh, some heuristics now
		name = strings.TrimSpace(section.Name)
		if strings.HasSuffix(name, "Bitmap type") || strings.HasSuffix(name, "Bitmap") {
			return matter.SectionDataTypeBitmap
		}
		if strings.HasSuffix(name, "Enum type") || strings.HasSuffix(name, "Enum") {
			return matter.SectionDataTypeEnum
		}
		if strings.HasSuffix(name, "Struct type") || strings.HasSuffix(name, "Struct") {
			return matter.SectionDataTypeStruct
		}
	}
	return matter.SectionUnknown
}

func (s *Section) ToModels(d *Doc) ([]interface{}, error) {
	var models []interface{}
	switch s.SecType {
	case matter.SectionCluster:
		clusters, err := s.toClusters(d)
		if err != nil {
			return nil, err
		}
		models = append(models, clusters...)
	case matter.SectionDeviceType:
		deviceTypes, err := s.toDeviceTypes(d)
		if err != nil {
			return nil, err
		}
		models = append(models, deviceTypes...)
	default:
		//slog.Info("unknown section type", "secType", s.SecType)
	}
	return models, nil
}

var dataTypeDefinitionPattern = regexp.MustCompile(`is\s+derived\s+from\s+(?:<<enum-def\s*,\s*)?(enum8|enum16|enum32|map8|map16|map32)(?:\s*>>)?`)

func (s *Section) GetDataType() string {
	var dataType string
	se := parse.FindFirst[*types.StringElement](s.Elements)
	if se != nil {
		match := dataTypeDefinitionPattern.FindStringSubmatch(se.Content)
		if match != nil {
			dataType = match[1]
		}
	}
	return dataType
}
