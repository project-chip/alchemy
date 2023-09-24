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
	TableColumnHierarchy
	TableColumnRole
	TableColumnContext
	TableColumnScope
	TableColumnPICS
	TableColumnValue
	TableColumnBit
	TableColumnDeviceName
	TableColumnSuperset
	TableColumnClass
	TableColumnDirection
	TableColumnResponse
)

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
	TableColumnValue,
	TableColumnName,
}

var EnumTableColumnOrder = [...]TableColumn{
	TableColumnValue,
	TableColumnName,
}

var BitmapTableColumnOrder = [...]TableColumn{
	TableColumnBit,
	TableColumnName,
}
