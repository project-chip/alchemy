package ascii

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

type Section struct {
	Name string

	Base *types.Section

	SecType MatterSection

	Elements []interface{}
}

type MatterSection uint8

const (
	MatterSectionUnknown MatterSection = iota
	MatterSectionPrefix                // Special section type for everything that comes before any known sections
	MatterSectionIntroduction
	MatterSectionRevisionHistory
	MatterSectionClassification
	MatterSectionClusterID
	MatterSectionFeatures
	MatterSectionDependencies
	MatterSectionDataTypes
	MatterSectionStatusCodes
	MatterSectionAttributes
	MatterSectionCommands
	MatterSectionEvents
	MatterSectionConditions
	MatterSectionClusterRequirements
	MatterSectionClusterRestrictions
	MatterSectionElementRequirements
	MatterSectionEndpointComposition
)

func NewSection(s *types.Section) *Section {
	ss := &Section{Base: s}
	for _, te := range s.Title {
		switch tel := te.(type) {
		case *types.StringElement:
			ss.Name = tel.Content
		case *types.InlineLink:

		default:
			//fmt.Printf("unknown section title element type: %T\n", te)
			//ss.Elements = append(ss.Elements, te)
		}
	}
	for _, e := range s.Elements {
		switch el := e.(type) {
		case *types.Section:
			ss.Elements = append(ss.Elements, NewSection(el))
		default:
			ss.Elements = append(ss.Elements, &Element{Base: e})
		}
	}
	return ss
}

func SectionTypeString(st MatterSection) string {
	switch st {
	case MatterSectionPrefix:
		return "Prefix"
	case MatterSectionUnknown:
		return "Unknown"
	case MatterSectionIntroduction:
		return "Introduction"
	case MatterSectionRevisionHistory:
		return "RevisionHistory"
	case MatterSectionClassification:
		return "Classification"
	case MatterSectionClusterID:
		return "ClusterID"
	case MatterSectionFeatures:
		return "Features"
	case MatterSectionDependencies:
		return "Dependencies"
	case MatterSectionDataTypes:
		return "DataTypes"
	case MatterSectionStatusCodes:
		return "StatusCodes"
	case MatterSectionAttributes:
		return "Attributes"
	case MatterSectionCommands:
		return "Commands"
	case MatterSectionEvents:
		return "Events"
	case MatterSectionConditions:
		return "Conditions"
	case MatterSectionClusterRequirements:
		return "Cluster Requirements"
	case MatterSectionClusterRestrictions:
		return "Cluster Restrictions"
	case MatterSectionElementRequirements:
		return "Element Requirements"
	case MatterSectionEndpointComposition:
		return "Endpoint Composition"
	default:
		return "invalid"
	}
}
