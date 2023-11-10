package matter

import "github.com/bytesparadise/libasciidoc/pkg/types"

type TableColumn uint8

const (
	TableColumnUnknown TableColumn = iota
	TableColumnID                  // Special section type for everything that comes before any known sections
	TableColumnName
	TableColumnType
	TableColumnConstraint
	TableColumnQuality
	TableColumnDefault
	TableColumnAccess
	TableColumnConformance
	TableColumnPriority
	TableColumnHierarchy
	TableColumnRole
	TableColumnContext
	TableColumnScope
	TableColumnPICS
	TableColumnValue
	TableColumnBit
	TableColumnCode
	TableColumnFeature
	TableColumnDeviceName
	TableColumnSuperset
	TableColumnClass
	TableColumnDirection
	TableColumnDescription
	TableColumnResponse
	TableColumnSummary
	TableColumnCluster
	TableColumnElement
	TableColumnClientServer
)

var TableColumnNames = map[TableColumn]string{
	TableColumnUnknown:      "Unknown",
	TableColumnID:           "ID",
	TableColumnName:         "Name",
	TableColumnType:         "Type",
	TableColumnConstraint:   "Constraint",
	TableColumnQuality:      "Quality",
	TableColumnDefault:      "Default",
	TableColumnAccess:       "Access",
	TableColumnConformance:  "Conformance",
	TableColumnPriority:     "Priority",
	TableColumnHierarchy:    "Hierarchy",
	TableColumnRole:         "Role",
	TableColumnContext:      "Context",
	TableColumnScope:        "Scope",
	TableColumnPICS:         "PICS",
	TableColumnValue:        "Value",
	TableColumnBit:          "Bit",
	TableColumnCode:         "Code",
	TableColumnFeature:      "Feature",
	TableColumnDeviceName:   "DeviceName",
	TableColumnSuperset:     "Superset",
	TableColumnClass:        "Class",
	TableColumnDirection:    "Direction",
	TableColumnDescription:  "Description",
	TableColumnResponse:     "Response",
	TableColumnSummary:      "Summary",
	TableColumnCluster:      "Cluster",
	TableColumnElement:      "Element",
	TableColumnClientServer: "Client/Server",
}

var BannedTableAttributes = [...]string{"cols", "frame", "width"}

var AttributesTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnType,
	TableColumnConstraint,
	TableColumnQuality,
	TableColumnDefault,
	TableColumnAccess,
	TableColumnConformance,
}

var AppClusterClassificationTableColumnOrder = [...]TableColumn{
	TableColumnHierarchy,
	TableColumnRole,
	TableColumnScope,
	TableColumnContext, // This will get renamed to Scope
	TableColumnPICS,
}

var DeviceTypeClassificationTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnDeviceName,
	TableColumnSuperset,
	TableColumnClass, // This will get renamed to Scope
	TableColumnScope,
}

var ClassificationTableColumnNames = map[TableColumn]string{
	TableColumnHierarchy:  "Hierarchy",
	TableColumnRole:       "Role",
	TableColumnScope:      "Scope",
	TableColumnContext:    "Scope", // Rename Context to Scope
	TableColumnPICS:       "PICS Code",
	TableColumnID:         "ID",
	TableColumnDeviceName: "Device Name",
	TableColumnSuperset:   "Superset",
	TableColumnClass:      "Class",
}

var ClusterIDSectionName = "Cluster ID"

var ClusterIDTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
}

var ClusterIDTableColumnNames = map[TableColumn]string{
	TableColumnID:   "ID",
	TableColumnName: "Name",
}

var CommandsSectionName = "Commands"

var CommandsTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnDirection,
	TableColumnResponse,
	TableColumnAccess,
	TableColumnConformance,
}

var CommandsTableColumnNames = map[TableColumn]string{
	TableColumnID:          "ID",
	TableColumnName:        "Name",
	TableColumnDirection:   "Direction",
	TableColumnResponse:    "Response",
	TableColumnAccess:      "Access",
	TableColumnConformance: "Conformance",
}

var AllowedTableAttributes = types.Attributes{
	"id":      nil,
	"title":   nil,
	"valign":  "middle",
	"options": types.Options{"header"},
}

var StructTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnType,
	TableColumnConstraint,
	TableColumnQuality,
	TableColumnDefault,
	TableColumnAccess,
	TableColumnConformance,
}

var StructTableColumnNames = map[TableColumn]string{
	TableColumnID:          "ID",
	TableColumnName:        "Name",
	TableColumnType:        "Type",
	TableColumnConstraint:  "Constraint",
	TableColumnQuality:     "Quality",
	TableColumnDefault:     "Default",
	TableColumnAccess:      "Access",
	TableColumnConformance: "Conformance",
}

var EnumTableColumnOrder = [...]TableColumn{
	TableColumnValue,
	TableColumnName,
	TableColumnSummary,
	TableColumnConformance,
}

var EnumTableColumnNames = map[TableColumn]string{
	TableColumnValue:       "Value",
	TableColumnName:        "Name",
	TableColumnSummary:     "Summary",
	TableColumnConformance: "Conformance",
}

var BitmapTableColumnOrder = [...]TableColumn{
	TableColumnBit,
	TableColumnName,
	TableColumnSummary,
	TableColumnConformance,
}

var BitmapTableColumnNames = map[TableColumn]string{
	TableColumnBit:         "ID",
	TableColumnName:        "Name",
	TableColumnSummary:     "Summary",
	TableColumnConformance: "Conformance",
}

var EventsTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnPriority,
	TableColumnAccess,
	TableColumnConformance,
}

var EventsTableColumnNames = map[TableColumn]string{
	TableColumnID:          "ID",
	TableColumnName:        "Name",
	TableColumnPriority:    "Priority",
	TableColumnAccess:      "Access",
	TableColumnConformance: "Conformance",
}

var EventTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnType,
	TableColumnConstraint,
	TableColumnQuality,
	TableColumnDefault,
	TableColumnConformance,
}

var EventTableColumnNames = map[TableColumn]string{
	TableColumnID:          "ID",
	TableColumnName:        "Name",
	TableColumnType:        "Type",
	TableColumnConstraint:  "Constraint",
	TableColumnQuality:     "Quality",
	TableColumnDefault:     "Default",
	TableColumnConformance: "Conformance",
}
