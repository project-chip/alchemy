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
	TableColumnAttributeID
	TableColumnClusterID
	TableColumnEventID
	TableColumnCommandID
	TableColumnFieldID
	TableColumnDevice
	TableColumnDeviceID
	TableColumnDeviceName
	TableColumnSupersetOf
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
	TableColumnConditionID
	TableColumnModeTagValue
	TableColumnNamespace
	TableColumnLocation
	TableColumnField
)

var TableColumnNames = map[TableColumn]string{
	TableColumnUnknown:      "Unknown",
	TableColumnID:           "ID",
	TableColumnName:         "Name",
	TableColumnType:         "Type",
	TableColumnConstraint:   "Constraint",
	TableColumnQuality:      "Quality",
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
	TableColumnEventID:      "Event ID",
	TableColumnCommandID:    "Command ID",
	TableColumnDevice:       "Device",
	TableColumnDeviceID:     "Device Type ID",
	TableColumnDeviceName:   "Device Name",
	TableColumnSupersetOf:   "Superset Of",
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
	TableColumnConditionID:  "Condition ID",
	TableColumnModeTagValue: "Mode Tag Value",
	TableColumnNamespace:    "Namespace",
	TableColumnLocation:     "Location",
	TableColumnField:        "Field",
	TableColumnFieldID:      "Field ID",
}

func (tc TableColumn) String() string {
	name, ok := TableColumnNames[tc]
	if ok {
		return name
	}
	return fmt.Sprintf("unknown table column name: %d", tc)
}

var AllowedTableAttributes = map[asciidoc.AttributeName]asciidoc.Elements{
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
	TableTypeDeviceTypeRequirements
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
			TableColumnDeviceID,
			TableColumnDeviceName,
			TableColumnSupersetOf,
			TableColumnClass, // This will get renamed to Scope
			TableColumnScope,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnID: TableColumnDeviceID,
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
			TableColumnClusterID,
			TableColumnName,
			TableColumnConformance,
		},
		RequiredColumns: []TableColumn{
			TableColumnID,
			TableColumnName,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnID: TableColumnClusterID,
		},
	},
	TableTypeCommands: {
		ColumnOrder: []TableColumn{
			TableColumnCommandID,
			TableColumnName,
			TableColumnDirection,
			TableColumnResponse,
			TableColumnAccess,
			TableColumnQuality,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnID: TableColumnCommandID,
		},
	},
	TableTypeCommandFields: {
		ColumnOrder: []TableColumn{
			TableColumnFieldID,
			TableColumnName,
			TableColumnDirection,
			TableColumnResponse,
			TableColumnAccess,
			TableColumnQuality,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnID: TableColumnFieldID,
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
			TableColumnEventID,
			TableColumnName,
			TableColumnPriority,
			TableColumnQuality,
			TableColumnAccess,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnID: TableColumnEventID,
		},
	},
	TableTypeEventFields: {
		ColumnOrder: []TableColumn{
			TableColumnFieldID,
			TableColumnName,
			TableColumnType,
			TableColumnConstraint,
			TableColumnQuality,
			TableColumnFallback,
			TableColumnConformance,
		},
		ColumnRenames: map[TableColumn]TableColumn{
			TableColumnID: TableColumnFieldID,
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
	TableTypeDeviceTypeRequirements: {
		ColumnOrder: []TableColumn{
			TableColumnDeviceID,
			TableColumnName,
			TableColumnConstraint,
			TableColumnConformance,
			TableColumnLocation,
		},
	},
}

var ClusterIDSectionName = "Cluster ID"
var ClusterIDsSectionName = "Cluster IDs"

var CommandsSectionName = "Commands"
