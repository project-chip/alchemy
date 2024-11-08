package matter

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
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
	TableColumnFallback
	TableColumnAccess
	TableColumnConformance
	TableColumnPriority
	TableColumnHierarchy
	TableColumnRole
	TableColumnContext
	TableColumnScope
	TableColumnPICS
	TableColumnPICSCode
	TableColumnValue
	TableColumnBit
	TableColumnCode
	TableColumnStatusCode
	TableColumnFeature
	TableColumnClusterID
	TableColumnDeviceID
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
	TableColumnNamespace
)

var TableColumnNames = map[TableColumn]string{
	TableColumnUnknown:      "Unknown",
	TableColumnID:           "ID",
	TableColumnName:         "Name",
	TableColumnType:         "Type",
	TableColumnConstraint:   "Constraint",
	TableColumnQuality:      "Quality",
	TableColumnDefault:      "Default",
	TableColumnFallback:     "Fallback",
	TableColumnAccess:       "Access",
	TableColumnConformance:  "Conformance",
	TableColumnPriority:     "Priority",
	TableColumnHierarchy:    "Hierarchy",
	TableColumnRole:         "Role",
	TableColumnContext:      "Context",
	TableColumnScope:        "Scope",
	TableColumnPICS:         "PICS",
	TableColumnPICSCode:     "PICS Code",
	TableColumnValue:        "Value",
	TableColumnBit:          "Bit",
	TableColumnCode:         "Code",
	TableColumnStatusCode:   "Status Code",
	TableColumnFeature:      "Feature",
	TableColumnClusterID:    "Cluster ID",
	TableColumnDeviceID:     "Device ID",
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
	TableColumnNamespace:    "Namespace",
}

func (tc TableColumn) String() string {
	name, ok := TableColumnNames[tc]
	if ok {
		return name
	}
	return fmt.Sprintf("unknown table column name: %d", tc)
}

var AllowedTableAttributes = map[asciidoc.AttributeName]asciidoc.Set{
	"id":      nil,
	"title":   nil,
	"valign":  {asciidoc.NewString("middle")},
	"options": {asciidoc.NewString("header")},
}
var BannedTableAttributes = [...]string{"cols", "frame", "width"}

func GetColumnName(column TableColumn, overrides map[TableColumn]TableColumn) (name string, ok bool) {
	if overrides != nil {
		if overrideColumn, hasOverride := overrides[column]; hasOverride {
			name, ok = TableColumnNames[overrideColumn]
			return
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
	TableTypeCommandFields
	TableTypeStruct
	TableTypeEnum
	TableTypeBitmap
	TableTypeEvents
	TableTypeEventFields
	TableTypeFeatures
)

type Table struct {
	AllowedColumns  []TableColumn
	RequiredColumns []TableColumn
	BannedColumns   []TableColumn
	ColumnOrder     []TableColumn
	ColumnRenames   map[TableColumn]TableColumn
}

var Tables = map[TableType]Table{
	TableTypeAttributes: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnType,
			TableColumnConstraint,
			TableColumnQuality,
			TableColumnFallback,
			TableColumnAccess,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnDefault: TableColumnFallback,
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
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnContext: TableColumnScope, // Rename Context to Scope
			TableColumnPICS:    TableColumnPICSCode,
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
			TableColumnQuality,
			TableColumnConformance,
		},
	},
	TableTypeCommandFields: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnDirection,
			TableColumnResponse,
			TableColumnAccess,
			TableColumnQuality,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnDefault: TableColumnFallback,
		},
	},
	TableTypeStruct: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnType,
			TableColumnConstraint,
			TableColumnQuality,
			TableColumnFallback,
			TableColumnAccess,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnDefault: TableColumnFallback,
		},
	},
	TableTypeEnum: {
		ColumnOrder: []TableColumn{
			TableColumnValue,
			TableColumnName,
			TableColumnSummary,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnDescription: TableColumnSummary, // Rename Description to Summary
			TableColumnStatusCode:  TableColumnValue,   // Rename Status Code in enums to Value
			TableColumnID:          TableColumnValue,   // Rename Status Code in enums to Value
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
	TableTypeEventFields: {
		ColumnOrder: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnType,
			TableColumnConstraint,
			TableColumnQuality,
			TableColumnFallback,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnDefault: TableColumnFallback,
		},
	},
	TableTypeFeatures: {
		ColumnOrder: []TableColumn{
			TableColumnBit,
			TableColumnCode,
			TableColumnFeature,
			TableColumnConformance,
			TableColumnSummary,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnID: TableColumnBit, // Rename ID to Bit
		},
	},
}

var ClusterIDSectionName = "Cluster ID"
var ClusterIDsSectionName = "Cluster IDs"

var CommandsSectionName = "Commands"
