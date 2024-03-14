package matter

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

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
	TableColumnRevision
	TableColumnResponse
	TableColumnSummary
	TableColumnCluster
	TableColumnElement
	TableColumnClientServer
	TableColumnCondition
	TableColumnModeTagValue
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
	TableColumnDeviceName:   "Device Name",
	TableColumnSuperset:     "Superset",
	TableColumnClass:        "Class",
	TableColumnDirection:    "Direction",
	TableColumnDescription:  "Description",
	TableColumnRevision:     "Revision",
	TableColumnResponse:     "Response",
	TableColumnSummary:      "Summary",
	TableColumnCluster:      "Cluster",
	TableColumnElement:      "Element",
	TableColumnClientServer: "Client/Server",
	TableColumnCondition:    "Condition",
	TableColumnModeTagValue: "Mode Tag Value",
}

func (tc TableColumn) String() string {
	name, ok := TableColumnNames[tc]
	if ok {
		return name
	}
	return fmt.Sprintf("unknown table column name: %d", tc)
}

var AllowedTableAttributes = types.Attributes{
	"id":      nil,
	"title":   nil,
	"valign":  "middle",
	"options": types.Options{"header"},
}
var BannedTableAttributes = [...]string{"cols", "frame", "width"}

func GetColumnName(column TableColumn, overrides map[TableColumn]string) (name string, ok bool) {
	if overrides != nil {
		if name, ok = overrides[column]; ok {
			return name, true
		}
	}
	name, ok = TableColumnNames[column]
	return
}

type TableType uint8

const (
	TableTypeUnknown TableType = iota
	TableTypeAttributes
)

type Table struct {
	AllowedColumns  []TableColumn
	RequiredColumns []TableColumn
	ColumnOrder     []TableColumn
}

var Tables = map[TableType]Table{
	TableTypeAttributes: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnType,
			TableColumnConstraint,
			TableColumnQuality,
			TableColumnDefault,
			TableColumnAccess,
			TableColumnConformance,
		},
	},
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
	TableColumnContext: "Scope", // Rename Context to Scope
	TableColumnPICS:    "PICS Code",
}

var ClusterIDSectionName = "Cluster ID"

var ClusterIDTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnConformance,
}

var ClusterIDTableRequiredColumns = [...]TableColumn{
	TableColumnID,
	TableColumnName,
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

var EnumTableColumnOrder = [...]TableColumn{
	TableColumnValue,
	TableColumnName,
	TableColumnSummary,
	TableColumnConformance,
}

var BitmapTableColumnOrder = [...]TableColumn{
	TableColumnBit,
	TableColumnName,
	TableColumnSummary,
	TableColumnConformance,
}

var EventsTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnPriority,
	TableColumnQuality,
	TableColumnAccess,
	TableColumnConformance,
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
