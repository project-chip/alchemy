package ascii

import "github.com/bytesparadise/libasciidoc/pkg/types"

type Section struct {
	Name string

	Base *types.Section

	SecType DocSectionType

	Elements []interface{}
}

type DocSectionType uint8

const (
	DocSectionTypePreface DocSectionType = 0
	DocSectionTypeAttributes
	DocSectionTypeFeatures
	DocSectionTypeDataTypes
	DocSectionTypeCommands
)

func GetSectionTitle(s *types.Section) string {
	for _, te := range s.GetTitle() {
		switch tel := te.(type) {
		case *types.StringElement:
			return tel.Content
		}
	}
	return ""
}
