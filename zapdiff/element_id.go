package zapdiff

import (
	"fmt"

	"github.com/beevik/etree"
)

func parentAndSelfAttr(e *etree.Element, attr string) string {
	parentID := getElementID(e.Parent())
	return fmt.Sprintf("%s/%s[@%s='%s']", parentID, e.Tag, attr, e.SelectAttrValue(attr, ""))
}

func parentAndSelfText(e *etree.Element) string {
	parentID := getElementID(e.Parent())
	return fmt.Sprintf("%s[%s='%s']/%s", parentID, e.Tag, e.Text(), e.Tag)
}

func getElementID(e *etree.Element) string {
	if e == nil {
		return ""
	}
	p := e.GetPath()

	switch p {
	case "/configurator":
		return "configurator"
	case "/configurator/global/attribute",
		"/configurator/enum",
		"/configurator/enum/item",
		"/configurator/struct",
		"/configurator/struct/item",
		"/configurator/bitmap",
		"/configurator/bitmap/field",
		"/configurator/cluster/command",
		"/configurator/cluster/command/arg",
		"/configurator/cluster/attribute",
		"/configurator/cluster/event",
		"/configurator/cluster/event/field",
		"/configurator/cluster/features/feature":
		return parentAndSelfAttr(e, "name")
	case "/configurator/enum/cluster",
		"/configurator/struct/cluster":
		return parentAndSelfAttr(e, "code")
	case "/configurator/cluster":
		parentID := getElementID(e.Parent())
		code := e.SelectAttrValue("code", "")
		if code != "" {
			return fmt.Sprintf("%s/%s[@code='%s']", parentID, e.Tag, code)
		}
		nameEl := e.SelectElement("name")
		if nameEl != nil {
			nameText := nameEl.Text()
			return fmt.Sprintf("%s/%s[name='%s']", parentID, e.Tag, nameText)
		} else {
			return getElementXPathSegment(e)
		}
	case "/configurator/cluster/name",
		"/configurator/cluster/domain",
		"/configurator/cluster/description",
		"/configurator/cluster/code",
		"/configurator/cluster/define",
		"/configurator/cluster/client",
		"/configurator/cluster/server":
		return parentAndSelfText(e)

	default:
		parentID := getElementID(e.Parent())
		selfSegment := getElementXPathSegment(e)
		return fmt.Sprintf("%s/%s", parentID, selfSegment)
	}
}
