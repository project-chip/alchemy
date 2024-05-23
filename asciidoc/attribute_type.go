package asciidoc

import "strings"

type AttributeType uint8

const (
	AttributeTypeNone AttributeType = iota
	AttributeTypeID
	AttributeTypeStyle
	AttributeTypeRole
	AttributeTypeTitle
	AttributeTypeAlternateText
	AttributeTypeColumns
)

type AttributeName string

const (
	AttributeNameNone          AttributeName = ""
	AttributeNameID            AttributeName = "id"
	AttributeNameColumns       AttributeName = "cols"
	AttributeNameStyle         AttributeName = "style"
	AttributeNameTitle         AttributeName = "title"
	AttributeNameReferenceText AttributeName = "reftext"
	AttributeNameAlternateText AttributeName = "alt"
	AttributeNameHeight        AttributeName = "height"
	AttributeNameWidth         AttributeName = "width"
	AttributeNamePDFWidth      AttributeName = "pdfwidth"
	AttributeNameAlign         AttributeName = "align"
	AttributeNameFloat         AttributeName = "float"
)

type AttributeNames []AttributeName

func (an AttributeNames) Join() string {
	var sb strings.Builder
	for _, a := range an {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(string(a))
	}
	return sb.String()
}

func (an AttributeNames) Equals(oan AttributeNames) bool {
	if len(an) != len(oan) {
		return false
	}
	for i, n := range an {
		on := oan[i]
		if n != on {
			return false
		}
	}
	return true
}

func attributeNameToType(name AttributeName) AttributeType {
	switch name {
	case AttributeNameAlternateText:
		return AttributeTypeAlternateText
	case AttributeNameColumns:
		return AttributeTypeColumns
	default:
		return AttributeTypeNone
	}
}

type AttributeQuoteType uint8

const (
	AttributeQuoteTypeNone AttributeQuoteType = iota
	AttributeQuoteTypeSingle
	AttributeQuoteTypeDouble
)
