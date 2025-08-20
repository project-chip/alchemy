package asciidoc

import (
	"fmt"
	"slices"
)

type AttributeList []Attribute

func (a AttributeList) Attributes() []Attribute {
	return a
}

func (al AttributeList) Equals(oal AttributeList) bool {
	if len(al) != len(oal) {
		return false
	}
	for i, a := range al {
		oa := oal[i]
		if !a.Equals(oa) {
			return false
		}
	}
	return true
}

func (al AttributeList) SetAttribute(name AttributeName, value Elements) {
	for _, a := range al {
		switch a := a.(type) {
		case *NamedAttribute:
			if a.Name == name {
				a.Val = value
				return
			}
		}
	}
	al.AppendAttribute(NewNamedAttribute(string(name), value, AttributeQuoteTypeDouble))
}

func (al *AttributeList) DeleteAttribute(name AttributeName) {
	index := -1
	for i, a := range *al {
		if na, ok := a.(*NamedAttribute); ok && na.Name == name {
			index = i
			break
		}
	}
	if index >= 0 {
		*al = slices.Delete(*al, index, index)
	}
}

func (a *AttributeList) SetAttributes(as ...Attribute) {
	*a = as
}

func (a *AttributeList) AppendAttribute(as ...Attribute) {
	*a = append(*a, as...)
}

func (al *AttributeList) ReadAttributes(el Element, attributes ...Attribute) (err error) {

	count := len(attributes)
	positionIndex := 0
	for i := 0; i < count; i++ {
		a := attributes[i]
		switch a := a.(type) {
		case *PositionalAttribute:
			a.Offset = positionIndex
			switch positionIndex {
			case 0:
				switch el.(type) {
				case *InlineImage, *BlockImage, *Link, *LinkMacro, *DocumentCrossReference:
					a.ImpliedName = AttributeNameAlternateText
				default:
					sa := parseShorthandAttribute(a)
					if sa != nil {
						attributes[i] = sa
					}
				}
			case 1:
				switch el.(type) {
				case *InlineImage, *BlockImage:
					a.ImpliedName = AttributeNameWidth
				}
			case 2:
				switch el.(type) {
				case *InlineImage, *BlockImage:
					a.ImpliedName = AttributeNameHeight
				}
			}
			positionIndex++
		case *NamedAttribute:
			switch a.Name {
			case AttributeNameColumns:
				attributes[i], err = parseColumnAttribute(a)
				if err != nil {
					return
				}
			}
		case *TitleAttribute:
		case *AnchorAttribute:
		default:
			err = fmt.Errorf("unexpected attribute type: %T", a)
			return
		}
	}
	*al = append(*al, attributes...)
	return
}

func (a *AttributeList) GetAttributeByName(name AttributeName) *NamedAttribute {
	for _, attr := range *a {
		switch attr := attr.(type) {
		case *NamedAttribute:
			if attr.Name == name {
				return attr
			}
		}
	}
	return nil
}

func (a AttributeList) traverse(parent Parent, yield func(Parent, Parent) bool) bool {
	for _, attr := range a.Attributes() {
		switch attr := attr.(type) {
		case *AnchorAttribute:
			if !yield(parent, &attr.Label) {
				return false
			}

		case *NamedAttribute:
			if !yield(parent, attr) {
				return false
			}
		case *PositionalAttribute:
			if !yield(parent, attr) {
				return false
			}
		case *TitleAttribute:
			if !yield(parent, attr) {
				return false
			}
		}
	}
	return true
}

func (a AttributeList) Clone() (oa AttributeList) {
	oa = make(AttributeList, len(a))
	for i, attr := range a {
		oa[i] = attr.Clone()
	}
	return
}
