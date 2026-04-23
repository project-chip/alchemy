package zapdiff

import "github.com/beevik/etree"

func getMismatchMissingType(e *etree.Element) XmlMismatchType {
	p := e.GetPath()
	switch p {
	case "/configurator/enum":
		return XmlMismatchMissingEnum
	case "/configurator/enum/item":
		return XmlMismatchMissingEnumItem
	case "/configurator/struct":
		return XmlMismatchMissingStruct
	case "/configurator/struct/item":
		return XmlMismatchMissingStructItem
	case "/configurator/bitmap":
		return XmlMismatchMissingBitmap
	case "/configurator/bitmap/field":
		return XmlMismatchMissingBitmapField
	case "/configurator/cluster":
		return XmlMismatchMissingCluster
	case "/configurator/cluster/command":
		return XmlMismatchMissingClusterCommand
	case "/configurator/cluster/attribute":
		return XmlMismatchMissingClusterAttribute
	case "/configurator/cluster/event":
		return XmlMismatchMissingClusterEvent

	case "/configurator/cluster/name",
		"/configurator/cluster/domain",
		"/configurator/cluster/description",
		"/configurator/cluster/code",
		"/configurator/cluster/define",
		"/configurator/cluster/client",
		"/configurator/cluster/server":
		return XmlMismatchClusterDetails

	case "/configurator/cluster/features/feature":
		return XmlMismatchMissingClusterFeature

	default:
		return XmlMismatchMissingTag
	}
}

func getMismatchMissingAttrType(e *etree.Element) XmlMismatchType {
	p := e.GetPath()
	switch p {
	case "/configurator/struct/item":
		return XmlMismatchStructItemMissingAttr
	case "/configurator/enum/item":
		return XmlMismatchEnumItemMissingAttr
	case "/configurator/bitmap":
		return XmlMismatchBitmapMissingAttr
	case "/configurator/bitmap/field":
		return XmlMismatchBitmapFieldMissingAttr
	case "/configurator/cluster":
		return XmlMismatchClusterMissingAttr
	case "/configurator/cluster/command":
		return XmlMismatchClusterCommandMissingAttr
	case "/configurator/cluster/attribute":
		return XmlMismatchClusterAttributeMissingAttr
	case "/configurator/cluster/event":
		return XmlMismatchClusterEventMissingAttr
	default:
		return XmlMismatchMissingAttr
	}
}

func getMismatchAttrValueType(e *etree.Element) XmlMismatchType {
	p := e.GetPath()
	switch p {
	case "/configurator/struct/item":
		return XmlMismatchStructItemAttrValue
	case "/configurator/enum/item":
		return XmlMismatchEnumItemAttrValue
	case "/configurator/bitmap":
		return XmlMismatchBitmapAttrValue
	case "/configurator/bitmap/field":
		return XmlMismatchBitmapFieldAttrValue
	case "/configurator/cluster":
		return XmlMismatchClusterAttrValue
	case "/configurator/cluster/command":
		return XmlMismatchClusterCommandAttrValue
	case "/configurator/cluster/attribute":
		return XmlMismatchClusterAttributeAttrValue
	case "/configurator/cluster/event":
		return XmlMismatchClusterEventAttrValue
	default:
		return XmlMismatchAttrValue
	}
}
