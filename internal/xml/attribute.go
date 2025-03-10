package xml

import "github.com/beevik/etree"

// Sets an attribute value, as long as it's not already set
func SetNonexistentAttr(el *etree.Element, name string, value string) *etree.Attr {
	a := el.SelectAttr(name)
	if a == nil {
		a = el.CreateAttr(name, value)
	}
	return a
}

func PrependAttribute(el *etree.Element, name string, value string, after ...string) {
	el.CreateAttr(name, value)
	a := el.RemoveAttr(name)
	if len(after) > 0 {
		var lastAfterIndex int = -1
		for index, at := range el.Attr {
			for _, af := range after {
				if af == at.Key {
					lastAfterIndex = index
					break
				}
			}
		}
		if lastAfterIndex >= 0 && lastAfterIndex != len(el.Attr)-1 {
			el.Attr = append(el.Attr[:lastAfterIndex+1], append([]etree.Attr{*a}, el.Attr[lastAfterIndex+1:]...)...)
			return
		}
	}
	el.Attr = append([]etree.Attr{*a}, el.Attr...)
}
