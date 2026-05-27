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
	SectionConditionRequirements
	SectionDeviceTypeRequirements
	SectionClusterRequirements
	SectionClusterRestrictions
	SectionElementRequirements
	SectionSemanticTagRequirements
	SectionComposedDeviceTypeClusterRequirements
	SectionComposedDeviceTypeElementRequirements
	SectionComposedDeviceTypeSemanticTagRequirements
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
		SectionConditionRequirements,
		SectionDeviceTypeRequirements,
		SectionClusterRequirements,
		SectionElementRequirements,
		SectionSemanticTagRequirements,
		SectionDeviceTypeRequirements,
		SectionComposedDeviceTypeClusterRequirements,
		SectionComposedDeviceTypeElementRequirements,
		SectionComposedDeviceTypeSemanticTagRequirements,
	},
}

var DataTypeSectionOrder = []Section{SectionPrefix, SectionDataTypeConstant, SectionDataTypeDef, SectionDataTypeBitmap, SectionDataTypeEnum, SectionDataTypeStruct, SectionDataTypeDef}
var DeviceRequirementsSectionOrder = []Section{SectionClusterRestrictions,
	SectionElementRequirements,
	SectionEndpointComposition,
	SectionConditionRequirements,
	SectionComposedDeviceTypeClusterRequirements,
	SectionComposedDeviceTypeElementRequirements,
	SectionSemanticTagRequirements}

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
	SectionComposedDeviceTypeClusterRequirements: "ComposedDeviceTypeClusterRequirements",
	SectionConditionRequirements:                 "ComposedDeviceTypeConditionRequirements",
	SectionComposedDeviceTypeElementRequirements: "ComposedDeviceTypeElementRequirements",
	SectionEndpointComposition:                   "EndpointComposition",
	SectionField:                                 "Field",
	SectionValue:                                 "Value",
	SectionBit:                                   "Bit",
	SectionDerivedClusterNamespace:               "DerivedClusterNamespace",
	SectionModeTags:                              "ModeTags",
	SectionGlobalElements:                        "GlobalElements",
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
	SectionComposedDeviceTypeClusterRequirements: "Cluster Requirements on Component Device Types",
	SectionConditionRequirements:                 "Condition Requirements on Component Device Types",
	SectionComposedDeviceTypeElementRequirements: "Element Requirements on Component Device Types",
	SectionEndpointComposition:                   "Endpoint Composition",
	SectionField:                                 "Field",
	SectionDerivedClusterNamespace:               "Derived Cluster Namespace",
	SectionModeTags:                              "Mode Tags",
}

var canonicalSectionTypeNames = map[Section]string{
	SectionIntroduction:                          "Introduction",
	SectionRevisionHistory:                       "Revision History",
	SectionClassification:                        "Classification",
	SectionFeatures:                              "Features",
	SectionDependencies:                          "Dependencies",
	SectionDataTypes:                             "Data Types",
	SectionAttributes:                            "Attributes",
	SectionCommands:                              "Commands",
	SectionEvents:                                "Events",
	SectionDeviceTypeRequirements:                "Device Type Requirements",
	SectionClusterRequirements:                   "Cluster Requirements",
	SectionClusterRestrictions:                   "Cluster Restrictions",
	SectionElementRequirements:                   "Element Requirements",
	SectionComposedDeviceTypeClusterRequirements: "Cluster Requirements on Component Device Types",
	SectionConditionRequirements:                 "Condition Requirements",
	SectionComposedDeviceTypeElementRequirements: "Element Requirements on Component Device Types",
	SectionEndpointComposition:                   "Endpoint Composition",
	SectionDerivedClusterNamespace:               "Derived Cluster Namespace",
	SectionModeTags:                              "Mode Tags",
}

func CanonicalSectionTypeName(st Section) string {
	if s, ok := canonicalSectionTypeNames[st]; ok {
		return s
	}
	return ""
}
