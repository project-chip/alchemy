package xml

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
)

func ReadSimpleElement(parent *etree.Element, name string) (value string, ok bool) {
	e := parent.SelectElement(name)
	if e == nil {
		return "", false
	}
	return e.Text(), true
}

func SetNonexistentAttr(el *etree.Element, name string, value string) *etree.Attr {
	a := el.SelectAttr(name)
	if a == nil {
		a = el.CreateAttr(name, value)
	}
	return a
}

func CreateSimpleElementIfNotExists(parent *etree.Element, name string, value string, afterElements ...string) *etree.Element {
	el := parent.SelectElement(name)
	if el != nil {
		return el
	}
	if len(afterElements) == 0 {
		el = parent.CreateElement(name)
	} else {
		el = etree.NewElement(name)
		AppendElement(parent, el, afterElements...)
	}
	if len(value) > 0 {
		el.SetText(value)
	}
	return el
}

func SetOrCreateSimpleElement(parent *etree.Element, name string, value string, afterElements ...string) *etree.Element {
	el := parent.SelectElement(name)
	if el == nil {
		if len(afterElements) == 0 {
			el = parent.CreateElement(name)
		} else {
			el = etree.NewElement(name)
			AppendElement(parent, el, afterElements...)
		}

	}
	if len(value) > 0 {
		el.SetText(value)
	}
	return el
}

func SetOrCreateSimpleNumber(parent *etree.Element, name string, value string, afterElements ...string) *etree.Element {
	el := parent.SelectElement(name)
	if el == nil {
		if len(afterElements) == 0 {
			el = parent.CreateElement(name)
		} else {
			el = etree.NewElement(name)
			AppendElement(parent, el, afterElements...)
		}

	}
	if len(value) > 0 {
		el.SetText(value)
	}
	return el
}

func AppendElement(parent *etree.Element, el *etree.Element, alternatives ...string) {
	tags := append([]string{el.Tag}, alternatives...)

	var lastSimilarElementIndex int = -1
	for _, tag := range tags {
		for i, e := range parent.Child {
			el, ok := e.(*etree.Element)
			if ok && el.Tag == tag {
				lastSimilarElementIndex = i
			}
		}
		if lastSimilarElementIndex >= 0 {
			break
		}
	}
	lastSimilarElementIndex++
	parent.InsertChildAt(lastSimilarElementIndex, el)
}

func InsertElementByAttribute(parent *etree.Element, el *etree.Element, attribute string, alternatives ...string) {
	name := el.SelectAttrValue(attribute, "")
	tag := el.Tag
	var insertIndex int = -1
	for i, e := range parent.Child {
		el, ok := e.(*etree.Element)
		if ok && el.Tag == tag {
			elName := el.SelectAttrValue(attribute, "")
			cmp := strings.Compare(elName, name)
			if cmp > 0 {
				insertIndex = i
				break
			}
		}
	}
	if insertIndex == -1 {
		AppendElement(parent, el, alternatives...)
		return
	}
	parent.InsertChildAt(insertIndex, el)
}

func InsertElementByAttributeNumber(parent *etree.Element, el *etree.Element, attribute string, number *matter.Number, alternatives ...string) {
	tag := el.Tag
	var insertIndex int = -1
	for i, e := range parent.Child {
		el, ok := e.(*etree.Element)
		if ok && el.Tag == tag {
			elName := el.SelectAttrValue(attribute, "")
			elNumber := matter.ParseNumber(elName)
			if elNumber.Valid() {
				cmp := elNumber.Compare(number)
				if cmp > 0 {
					insertIndex = i
					break
				}
			}
		}
	}
	if insertIndex == -1 {
		AppendElement(parent, el, alternatives...)
		return
	}
	parent.InsertChildAt(insertIndex, el)
}

func InsertElementByName(parent *etree.Element, el *etree.Element, alternatives ...string) {
	text := el.Text()
	tag := el.Tag
	var insertIndex int = -1
	for i, e := range parent.Child {
		el, ok := e.(*etree.Element)
		if ok && el.Tag == tag {
			elText := el.Text()
			cmp := strings.Compare(elText, text)
			if cmp > 0 {
				insertIndex = i
				break
			}
		}
	}
	if insertIndex == -1 {
		AppendElement(parent, el, alternatives...)
		return
	}
	parent.InsertChildAt(insertIndex, el)
}

func PrependAttribute(el *etree.Element, name string, value string) {
	el.CreateAttr(name, value)
	a := el.RemoveAttr(name)
	el.Attr = append([]etree.Attr{*a}, el.Attr...)
}
