package matter

type Section uint8

const (
	SectionUnknown Section = iota
	SectionPrefix          // Special section type for everything that comes before any known sections
	SectionIntroduction
	SectionRevisionHistory
	SectionClassification
	SectionClusterID
	SectionFeatures
	SectionDependencies
	SectionDataTypes
	SectionStatusCodes
	SectionAttributes
	SectionCommands
	SectionEvents
	SectionConditions
	SectionClusterRequirements
	SectionClusterRestrictions
	SectionElementRequirements
	SectionEndpointComposition
)

func SectionTypeString(st Section) string {
	switch st {
	case SectionPrefix:
		return "Prefix"
	case SectionUnknown:
		return "Unknown"
	case SectionIntroduction:
		return "Introduction"
	case SectionRevisionHistory:
		return "RevisionHistory"
	case SectionClassification:
		return "Classification"
	case SectionClusterID:
		return "ClusterID"
	case SectionFeatures:
		return "Features"
	case SectionDependencies:
		return "Dependencies"
	case SectionDataTypes:
		return "DataTypes"
	case SectionStatusCodes:
		return "StatusCodes"
	case SectionAttributes:
		return "Attributes"
	case SectionCommands:
		return "Commands"
	case SectionEvents:
		return "Events"
	case SectionConditions:
		return "Conditions"
	case SectionClusterRequirements:
		return "Cluster Requirements"
	case SectionClusterRestrictions:
		return "Cluster Restrictions"
	case SectionElementRequirements:
		return "Element Requirements"
	case SectionEndpointComposition:
		return "Endpoint Composition"
	default:
		return "invalid"
	}
}
