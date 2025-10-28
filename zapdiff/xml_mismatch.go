package zapdiff

import "fmt"

type XmlMismatchLevel uint8

const (
	MismatchLevel1 XmlMismatchLevel = iota
	MismatchLevel2
	MismatchLevel3
	MismatchLevel4
	MismatchLevel5
)

func (l XmlMismatchLevel) String() string {
	switch l {
	case MismatchLevel1:
		return "L1"
	case MismatchLevel2:
		return "L2"
	case MismatchLevel3:
		return "L3"
	case MismatchLevel4:
		return "L4"
	case MismatchLevel5:
		return "L5"

	default:
		return "UNKNOWN"
	}
}

type XmlMismatchType uint8

const (
	XmlMismatchNone XmlMismatchType = iota

	// File level
	XmlMismatchNewFile

	// Generic Tag/Attr Mismatches
	XmlMismatchMissingTag
	XmlMismatchMissingAttr
	XmlMismatchAttrValue

	// Enums
	XmlMismatchMissingEnum
	XmlMismatchMissingEnumItem
	XmlMismatchEnumItemMissingAttr
	XmlMismatchEnumItemAttrValue

	// Structs
	XmlMismatchMissingStruct
	XmlMismatchMissingStructItem
	XmlMismatchStructItemMissingAttr
	XmlMismatchStructItemAttrValue

	// Bitmaps
	XmlMismatchMissingBitmap
	XmlMismatchMissingBitmapField
	XmlMismatchBitmapMissingAttr
	XmlMismatchBitmapAttrValue
	XmlMismatchBitmapFieldMissingAttr
	XmlMismatchBitmapFieldAttrValue

	// Clusters (Top Level)
	XmlMismatchMissingCluster
	XmlMismatchClusterMissingAttr
	XmlMismatchClusterAttrValue

	// Clusters
	XmlMismatchMissingClusterCommand
	XmlMismatchClusterCommandMissingAttr
	XmlMismatchClusterCommandAttrValue
	XmlMismatchMissingClusterAttribute
	XmlMismatchClusterAttributeMissingAttr
	XmlMismatchClusterAttributeAttrValue
	XmlMismatchMissingClusterEvent
	XmlMismatchClusterEventMissingAttr
	XmlMismatchClusterEventAttrValue
	XmlMismatchMissingClusterFeature

	XmlMismatchClusterDetails
)

func (t XmlMismatchType) String() string {
	switch t {
	case XmlMismatchNone:
		return "None"

	// File
	case XmlMismatchNewFile:
		return "FileNotFound"

	// Generic
	case XmlMismatchMissingTag:
		return "MissingTag"
	case XmlMismatchMissingAttr:
		return "MissingAttr"
	case XmlMismatchAttrValue:
		return "AttrValue"

	// Enums
	case XmlMismatchMissingEnum:
		return "MissingEnum"
	case XmlMismatchMissingEnumItem:
		return "MissingEnumItem"
	case XmlMismatchEnumItemMissingAttr:
		return "EnumItemMissingAttr"
	case XmlMismatchEnumItemAttrValue:
		return "EnumItemAttrValue"

	// Structs
	case XmlMismatchMissingStruct:
		return "MissingStruct"
	case XmlMismatchMissingStructItem:
		return "MissingStructItem"
	case XmlMismatchStructItemMissingAttr:
		return "StructItemMissingAttr"
	case XmlMismatchStructItemAttrValue:
		return "StructItemAttrValue"

	// Bitmaps
	case XmlMismatchMissingBitmap:
		return "MissingBitmap"
	case XmlMismatchMissingBitmapField:
		return "MissingBitmapField"
	case XmlMismatchBitmapMissingAttr:
		return "BitmapMissingAttr"
	case XmlMismatchBitmapAttrValue:
		return "BitmapAttrValue"
	case XmlMismatchBitmapFieldMissingAttr:
		return "BitmapFieldMissingAttr"
	case XmlMismatchBitmapFieldAttrValue:
		return "BitmapFieldAttrValue"

	// Clusters (Top Level)
	case XmlMismatchMissingCluster:
		return "MissingCluster"
	case XmlMismatchClusterMissingAttr:
		return "ClusterMissingAttr"
	case XmlMismatchClusterAttrValue:
		return "ClusterAttrValue"

	// Clusters
	case XmlMismatchMissingClusterCommand:
		return "MissingClusterCommand"
	case XmlMismatchClusterCommandMissingAttr:
		return "ClusterCommandMissingAttr"
	case XmlMismatchClusterCommandAttrValue:
		return "ClusterCommandAttrValue"
	case XmlMismatchMissingClusterAttribute:
		return "MissingClusterAttribute"
	case XmlMismatchClusterAttributeMissingAttr:
		return "ClusterAttributeMissingAttr"
	case XmlMismatchClusterAttributeAttrValue:
		return "ClusterAttributeAttrValue"
	case XmlMismatchMissingClusterEvent:
		return "MissingClusterEvent"
	case XmlMismatchClusterEventMissingAttr:
		return "ClusterEventMissingAttr"
	case XmlMismatchClusterEventAttrValue:
		return "ClusterEventAttrValue"
	case XmlMismatchMissingClusterFeature:
		return "MissingClusterFeature"

	case XmlMismatchClusterDetails:
		return "ClusterDetails"

	default:
		return "Unknown Mismatch"
	}
}

func (t XmlMismatchType) Level() XmlMismatchLevel {
	switch t {
	// File
	case XmlMismatchNewFile:
		return MismatchLevel1

	// Generic
	case XmlMismatchMissingTag:
		return MismatchLevel2
	case XmlMismatchMissingAttr:
		return MismatchLevel1
	case XmlMismatchAttrValue:
		return MismatchLevel2

	// Enums
	case XmlMismatchMissingEnum:
		return MismatchLevel4
	case XmlMismatchMissingEnumItem:
		return MismatchLevel4
	case XmlMismatchEnumItemMissingAttr:
		return MismatchLevel1
	case XmlMismatchEnumItemAttrValue:
		return MismatchLevel4

	// Structs
	case XmlMismatchMissingStruct:
		return MismatchLevel4
	case XmlMismatchMissingStructItem:
		return MismatchLevel4
	case XmlMismatchStructItemMissingAttr:
		return MismatchLevel1
	case XmlMismatchStructItemAttrValue:
		return MismatchLevel4

	// Bitmaps
	case XmlMismatchMissingBitmap:
		return MismatchLevel4
	case XmlMismatchMissingBitmapField:
		return MismatchLevel4
	case XmlMismatchBitmapMissingAttr:
		return MismatchLevel1
	case XmlMismatchBitmapAttrValue:
		return MismatchLevel4
	case XmlMismatchBitmapFieldMissingAttr:
		return MismatchLevel1
	case XmlMismatchBitmapFieldAttrValue:
		return MismatchLevel4

	// Clusters (Top Level)
	case XmlMismatchMissingCluster:
		return MismatchLevel4
	case XmlMismatchClusterMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterAttrValue:
		return MismatchLevel4

	// Clusters
	case XmlMismatchMissingClusterCommand:
		return MismatchLevel4
	case XmlMismatchClusterCommandMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterCommandAttrValue:
		return MismatchLevel4
	case XmlMismatchMissingClusterAttribute:
		return MismatchLevel4
	case XmlMismatchClusterAttributeMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterAttributeAttrValue:
		return MismatchLevel4
	case XmlMismatchMissingClusterEvent:
		return MismatchLevel4
	case XmlMismatchClusterEventMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterEventAttrValue:
		return MismatchLevel4
	case XmlMismatchMissingClusterFeature:
		return MismatchLevel4

	case XmlMismatchClusterDetails:
		return MismatchLevel4

	default:
		return MismatchLevel1
	}
}

type XmlMismatch struct {
	Path      string
	Details   string
	Type      XmlMismatchType
	ElementID string
}

func (m XmlMismatch) Level() XmlMismatchLevel {
	return m.Type.Level()
}

func (m *XmlMismatch) Error() string {
	return fmt.Sprintf("[%s] %s - in %s: %s", m.Level().String(), m.Type.String(), m.Path, m.Details)
}
