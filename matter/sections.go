package matter

type Section uint8

const (
	SectionUnknown Section = iota
	SectionPrefix          // Special section type for everything that comes before any known sections
	SectionTop
	SectionIntroduction
	SectionRevisionHistory
	SectionClassification
	SectionCluster
	SectionClusterID
	SectionFeatures
	SectionDependencies
	SectionDataTypes
	SectionDataTypeBitmap
	SectionDataTypeEnum
	SectionDataTypeStruct
	SectionDeviceType
	SectionStatusCodes
	SectionAttributes
	SectionAttribute
	SectionCommands
	SectionCommand
	SectionEvents
	SectionEvent
	SectionConditions
	SectionClusterRequirements
	SectionClusterRestrictions
	SectionElementRequirements
	SectionEndpointComposition
	SectionField
)

var TopLevelSectionOrders map[DocType][]Section

func init() {
	TopLevelSectionOrders = make(map[DocType][]Section)
	TopLevelSectionOrders[DocTypeAppCluster] = []Section{
		SectionPrefix,
		SectionRevisionHistory,
		SectionClassification,
		SectionClusterID,
		SectionFeatures,
		SectionDependencies,
		SectionDataTypes,
		SectionStatusCodes,
		SectionAttributes,
		SectionCommands,
		SectionEvents,
	}
	TopLevelSectionOrders[DocTypeDeviceType] = []Section{
		SectionPrefix,
		SectionRevisionHistory,
		SectionClassification,
		SectionConditions,
		SectionClusterRequirements,
		SectionClusterRestrictions,
		SectionElementRequirements,
		SectionEndpointComposition,
	}
}

var sectionTypeStrings = map[Section]string{
	SectionPrefix:              "Prefix",
	SectionTop:                 "Top",
	SectionUnknown:             "Unknown",
	SectionIntroduction:        "Introduction",
	SectionRevisionHistory:     "RevisionHistory",
	SectionClassification:      "Classification",
	SectionCluster:             "Cluster",
	SectionClusterID:           "ClusterID",
	SectionFeatures:            "Features",
	SectionDependencies:        "Dependencies",
	SectionDataTypes:           "DataTypes",
	SectionDataTypeBitmap:      "Bitmap",
	SectionDataTypeEnum:        "Enum",
	SectionDataTypeStruct:      "Struct",
	SectionDeviceType:          "DeviceType",
	SectionStatusCodes:         "StatusCodes",
	SectionAttributes:          "Attributes",
	SectionAttribute:           "Attribute",
	SectionCommands:            "Commands",
	SectionEvents:              "Events",
	SectionEvent:               "Event",
	SectionConditions:          "Conditions",
	SectionClusterRequirements: "ClusterRequirements",
	SectionClusterRestrictions: "ClusterRestrictions",
	SectionElementRequirements: "ElementRequirements",
	SectionEndpointComposition: "EndpointComposition",
	SectionField:               "Field",
}

func SectionTypeString(st Section) string {
	if s, ok := sectionTypeStrings[st]; ok {
		return s
	}
	return "invalid"

}

var sectionTypeNames = map[Section]string{
	SectionPrefix:              "Prefix",
	SectionTop:                 "Top",
	SectionUnknown:             "Unknown",
	SectionIntroduction:        "Introduction",
	SectionRevisionHistory:     "Revision History",
	SectionClassification:      "Classification",
	SectionCluster:             "Cluster",
	SectionClusterID:           "Cluster ID",
	SectionFeatures:            "Features",
	SectionDependencies:        "Dependencies",
	SectionDataTypes:           "Data Types",
	SectionDataTypeBitmap:      "Bitmap",
	SectionDataTypeEnum:        "Enum",
	SectionDataTypeStruct:      "Struct",
	SectionDeviceType:          "Device Type",
	SectionStatusCodes:         "Status Codes",
	SectionAttributes:          "Attributes",
	SectionAttribute:           "Attribute",
	SectionCommands:            "Commands",
	SectionCommand:             "Command",
	SectionEvents:              "Events",
	SectionConditions:          "Conditions",
	SectionClusterRequirements: "Cluster Requirements",
	SectionClusterRestrictions: "Cluster Restrictions",
	SectionElementRequirements: "Element Requirements",
	SectionEndpointComposition: "Endpoint Composition",
	SectionField:               "Field",
}

func SectionTypeName(st Section) string {
	if s, ok := sectionTypeNames[st]; ok {
		return s
	}
	return ""
}
