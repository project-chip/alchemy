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
	SectionFeature
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
	SectionNamespace
	SectionConditions
	SectionDeviceTypeRequirements
	SectionClusterRequirements
	SectionClusterRestrictions
	SectionElementRequirements
	SectionComposedDeviceTypeClusterRequirements
	SectionComposedDeviceTypeConditionRequirements
	SectionComposedDeviceTypeElementRequirements
	SectionEndpointComposition
	SectionField
	SectionValue
	SectionBit
	SectionDerivedClusterNamespace
	SectionModeTags
	SectionGlobalElements
	SectionDataTypeDef
	SectionDataTypeConstant
)

var TopLevelSectionOrders = map[DocType][]Section{
	DocTypeCluster: {
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
	},
	DocTypeDeviceType: {
		SectionPrefix,
		SectionRevisionHistory,
		SectionClassification,
		SectionConditions,
		SectionDeviceTypeRequirements,
		SectionClusterRequirements,
		SectionClusterRestrictions,
		SectionElementRequirements,
		SectionEndpointComposition,
	},
}

var DataTypeSectionOrder = []Section{SectionPrefix, SectionDataTypeConstant, SectionDataTypeBitmap, SectionDataTypeEnum, SectionDataTypeStruct, SectionDataTypeDef}

var sectionTypeStrings = map[Section]string{
	SectionPrefix:                 "Prefix",
	SectionTop:                    "Top",
	SectionUnknown:                "Unknown",
	SectionIntroduction:           "Introduction",
	SectionRevisionHistory:        "RevisionHistory",
	SectionClassification:         "Classification",
	SectionCluster:                "Cluster",
	SectionClusterID:              "ClusterID",
	SectionFeatures:               "Features",
	SectionFeature:                "Feature",
	SectionDependencies:           "Dependencies",
	SectionDataTypes:              "DataTypes",
	SectionDataTypeBitmap:         "Bitmap",
	SectionDataTypeEnum:           "Enum",
	SectionDataTypeStruct:         "Struct",
	SectionDataTypeDef:            "TypeDef",
	SectionDataTypeConstant:       "Constant",
	SectionDeviceType:             "DeviceType",
	SectionStatusCodes:            "StatusCodes",
	SectionAttributes:             "Attributes",
	SectionAttribute:              "Attribute",
	SectionCommands:               "Commands",
	SectionCommand:                "Command",
	SectionEvents:                 "Events",
	SectionEvent:                  "Event",
	SectionNamespace:              "Namespace",
	SectionConditions:             "Conditions",
	SectionDeviceTypeRequirements: "DeviceTypeRequirements",
	SectionClusterRequirements:    "ClusterRequirements",
	SectionClusterRestrictions:    "ClusterRestrictions",
	SectionElementRequirements:    "ElementRequirements",
	SectionComposedDeviceTypeClusterRequirements:   "ComposedDeviceTypeClusterRequirements",
	SectionComposedDeviceTypeConditionRequirements: "ComposedDeviceTypeConditionRequirements",
	SectionComposedDeviceTypeElementRequirements:   "ComposedDeviceTypeElementRequirements",
	SectionEndpointComposition:                     "EndpointComposition",
	SectionField:                                   "Field",
	SectionValue:                                   "Value",
	SectionBit:                                     "Bit",
	SectionDerivedClusterNamespace:                 "DerivedClusterNamespace",
	SectionModeTags:                                "ModeTags",
	SectionGlobalElements:                          "GlobalElements",
}

func (st Section) String() string {
	if s, ok := sectionTypeStrings[st]; ok {
		return s
	}
	return "invalid"
}

var sectionTypeNames = map[Section]string{
	SectionPrefix:                                "Prefix",
	SectionTop:                                   "Top",
	SectionUnknown:                               "Unknown",
	SectionIntroduction:                          "Introduction",
	SectionRevisionHistory:                       "Revision History",
	SectionClassification:                        "Classification",
	SectionCluster:                               "Cluster",
	SectionClusterID:                             "Cluster ID",
	SectionFeatures:                              "Features",
	SectionDependencies:                          "Dependencies",
	SectionDataTypes:                             "Data Types",
	SectionDataTypeBitmap:                        "Bitmap",
	SectionDataTypeEnum:                          "Enum",
	SectionDataTypeStruct:                        "Struct",
	SectionDataTypeDef:                           "Type Definition",
	SectionDataTypeConstant:                      "Constant",
	SectionDeviceType:                            "Device Type",
	SectionStatusCodes:                           "Status Codes",
	SectionAttributes:                            "Attributes",
	SectionAttribute:                             "Attribute",
	SectionCommands:                              "Commands",
	SectionCommand:                               "Command",
	SectionEvents:                                "Events",
	SectionNamespace:                             "Namespace",
	SectionConditions:                            "Conditions",
	SectionDeviceTypeRequirements:                "Device Type Requirements",
	SectionClusterRequirements:                   "Cluster Requirements",
	SectionClusterRestrictions:                   "Cluster Restrictions",
	SectionElementRequirements:                   "Element Requirements",
	SectionComposedDeviceTypeClusterRequirements: "Cluster Requirements on Composing Device Types",
	SectionComposedDeviceTypeConditionRequirements: "Condition Requirements on Composing Device Types",
	SectionComposedDeviceTypeElementRequirements:   "Element Requirements on Composing Device Types",
	SectionEndpointComposition:                     "Endpoint Composition",
	SectionField:                                   "Field",
	SectionDerivedClusterNamespace:                 "Derived Cluster Namespace",
	SectionModeTags:                                "Mode Tags",
}

func SectionTypeName(st Section) string {
	if s, ok := sectionTypeNames[st]; ok {
		return s
	}
	return ""
}
