package zapdiff

import "fmt"

type XmlMismatchLevel uint8

const (
	MismatchLevel1 XmlMismatchLevel = iota
	MismatchLevel2
	MismatchLevel3
)

func (l XmlMismatchLevel) String() string {
	switch l {
	case MismatchLevel1:
		return "L1"
	case MismatchLevel2:
		return "L2"
	case MismatchLevel3:
		return "L3"

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
		return "File: Not Found"

	// Generic
	case XmlMismatchMissingTag:
		return "Tag: Missing"
	case XmlMismatchMissingAttr:
		return "Attribute: Missing"
	case XmlMismatchAttrValue:
		return "Attribute: Value Mismatch"

	// Enums
	case XmlMismatchMissingEnum:
		return "Enum: Missing"
	case XmlMismatchMissingEnumItem:
		return "Enum Item: Missing"
	case XmlMismatchEnumItemMissingAttr:
		return "Enum Item: Missing Attribute"
	case XmlMismatchEnumItemAttrValue:
		return "Enum Item: Attribute Value Mismatch"

	// Structs
	case XmlMismatchMissingStruct:
		return "Struct: Missing"
	case XmlMismatchMissingStructItem:
		return "Struct Item: Missing"
	case XmlMismatchStructItemMissingAttr:
		return "Struct Item: Missing Attribute"
	case XmlMismatchStructItemAttrValue:
		return "Struct Item: Attribute Value Mismatch"

	// Bitmaps
	case XmlMismatchMissingBitmap:
		return "Bitmap: Missing"
	case XmlMismatchMissingBitmapField:
		return "Bitmap Field: Missing"
	case XmlMismatchBitmapMissingAttr:
		return "Bitmap: Missing Attribute"
	case XmlMismatchBitmapAttrValue:
		return "Bitmap: Attribute Value Mismatch"
	case XmlMismatchBitmapFieldMissingAttr:
		return "Bitmap Field: Missing Attribute"
	case XmlMismatchBitmapFieldAttrValue:
		return "Bitmap Field: Attribute Value Mismatch"

	// Clusters (Top Level)
	case XmlMismatchMissingCluster:
		return "Cluster: Missing"
	case XmlMismatchClusterMissingAttr:
		return "Cluster: Missing Attribute"
	case XmlMismatchClusterAttrValue:
		return "Cluster: Attribute Value Mismatch"

	// Clusters
	case XmlMismatchMissingClusterCommand:
		return "Cluster Command: Missing"
	case XmlMismatchClusterCommandMissingAttr:
		return "Cluster Command: Missing Attribute"
	case XmlMismatchClusterCommandAttrValue:
		return "Cluster Command: Attribute Value Mismatch"
	case XmlMismatchMissingClusterAttribute:
		return "Cluster Attribute: Missing"
	case XmlMismatchClusterAttributeMissingAttr:
		return "Cluster Attribute: Missing Attribute"
	case XmlMismatchClusterAttributeAttrValue:
		return "Cluster Attribute: Attribute Value Mismatch"
	case XmlMismatchMissingClusterEvent:
		return "Cluster Event: Missing"
	case XmlMismatchClusterEventMissingAttr:
		return "Cluster Event: Missing Attribute"
	case XmlMismatchClusterEventAttrValue:
		return "Cluster Event: Attribute Value Mismatch"
	case XmlMismatchMissingClusterFeature:
		return "Cluster Feature: Missing"

	case XmlMismatchClusterDetails:
		return "Cluster: Details Mismatch"

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
		return MismatchLevel3
	case XmlMismatchMissingEnumItem:
		return MismatchLevel3
	case XmlMismatchEnumItemMissingAttr:
		return MismatchLevel1
	case XmlMismatchEnumItemAttrValue:
		return MismatchLevel3

	// Structs
	case XmlMismatchMissingStruct:
		return MismatchLevel3
	case XmlMismatchMissingStructItem:
		return MismatchLevel3
	case XmlMismatchStructItemMissingAttr:
		return MismatchLevel1
	case XmlMismatchStructItemAttrValue:
		return MismatchLevel3

	// Bitmaps
	case XmlMismatchMissingBitmap:
		return MismatchLevel3
	case XmlMismatchMissingBitmapField:
		return MismatchLevel3
	case XmlMismatchBitmapMissingAttr:
		return MismatchLevel1
	case XmlMismatchBitmapAttrValue:
		return MismatchLevel3
	case XmlMismatchBitmapFieldMissingAttr:
		return MismatchLevel1
	case XmlMismatchBitmapFieldAttrValue:
		return MismatchLevel3

	// Clusters (Top Level)
	case XmlMismatchMissingCluster:
		return MismatchLevel3
	case XmlMismatchClusterMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterAttrValue:
		return MismatchLevel3

	// Clusters
	case XmlMismatchMissingClusterCommand:
		return MismatchLevel3
	case XmlMismatchClusterCommandMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterCommandAttrValue:
		return MismatchLevel3
	case XmlMismatchMissingClusterAttribute:
		return MismatchLevel3
	case XmlMismatchClusterAttributeMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterAttributeAttrValue:
		return MismatchLevel3
	case XmlMismatchMissingClusterEvent:
		return MismatchLevel3
	case XmlMismatchClusterEventMissingAttr:
		return MismatchLevel1
	case XmlMismatchClusterEventAttrValue:
		return MismatchLevel3
	case XmlMismatchMissingClusterFeature:
		return MismatchLevel3

	case XmlMismatchClusterDetails:
		return MismatchLevel3

	default:
		return MismatchLevel1
	}
}

type XmlMismatch struct {
	Path                   string
	Details                string
	Type                   XmlMismatchType
	EntityUniqueIdentifier string
}

func (m XmlMismatch) Level() XmlMismatchLevel {
	return m.Type.Level()
}

func (m *XmlMismatch) Error() string {
	return fmt.Sprintf("[%s] %s - in %s: %s", m.Level().String(), m.Type.String(), m.Path, m.Details)
}
