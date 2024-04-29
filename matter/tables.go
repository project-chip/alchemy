package matter

import (
	"fmt"

	"github.com/hasty/adoc/elements"
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

var AllowedTableAttributes = map[elements.AttributeName]any{
	"id":      nil,
	"title":   nil,
	"valign":  "middle",
	"options": []string{"header"},
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
	TableTypeAppClusterClassification
	TableTypeDeviceTypeClassification
	TableTypeClassification
	TableTypeClusterID
	TableTypeCommands
	TableTypeStruct
	TableTypeEnum
	TableTypeBitmap
	TableTypeEvents
	TableTypeEvent
)

type Table struct {
	AllowedColumns  []TableColumn
	RequiredColumns []TableColumn
	BannedColumns   []TableColumn
	ColumnOrder     []TableColumn
	ColumnNames     map[TableColumn]string
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
	TableTypeAppClusterClassification: {
		ColumnOrder: []TableColumn{
			TableColumnHierarchy,
			TableColumnRole,
			TableColumnScope,
			TableColumnContext, // This will get renamed to Scope
			TableColumnPICS,
		},
	},
	TableTypeDeviceTypeClassification: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnDeviceName,
			TableColumnSuperset,
			TableColumnClass, // This will get renamed to Scope
			TableColumnScope,
		},
	},
	TableTypeClassification: {
		ColumnNames: map[TableColumn]string{
			TableColumnContext: "Scope", // Rename Context to Scope
			TableColumnPICS:    "PICS Code",
		},
	},
	TableTypeClusterID: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnConformance,
		},
		RequiredColumns: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnConformance,
		},
	},
	TableTypeCommands: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnDirection,
			TableColumnResponse,
			TableColumnAccess,
			TableColumnConformance,
		},
	},
	TableTypeStruct: {
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
	TableTypeEnum: {
		ColumnOrder: []TableColumn{
			TableColumnValue,
			TableColumnName,
			TableColumnSummary,
			TableColumnConformance,
		},
	},
	TableTypeBitmap: {
		ColumnOrder: []TableColumn{
			TableColumnBit,
			TableColumnName,
			TableColumnSummary,
			TableColumnConformance,
		},
	},
	TableTypeEvents: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnPriority,
			TableColumnQuality,
			TableColumnAccess,
			TableColumnConformance,
		},
	},
	TableTypeEvent: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnType,
			TableColumnConstraint,
			TableColumnQuality,
			TableColumnDefault,
			TableColumnConformance,
		},
	},
}

var ClusterIDSectionName = "Cluster ID"
var ClusterIDsSectionName = "Cluster IDs"

var CommandsSectionName = "Commands"
