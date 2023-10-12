package db

import "github.com/hasty/matterfmt/matter"

var (
	documentTable     = "document"
	clusterTable      = "cluster"
	featureTable      = "feature"
	dataTypeTable     = "data_type"
	structField       = "struct_field"
	bitmapValue       = "bitmap_value"
	enumValue         = "enum_value"
	attributeTable    = "attribute"
	eventTable        = "event"
	eventFieldTable   = "event_field"
	commandTable      = "command"
	commandFieldTable = "command_field"
)

type tableSchemaDef struct {
	parent  string
	columns []matter.TableColumn
}

var tableSchema = map[string]tableSchemaDef{
	documentTable: {
		columns: []matter.TableColumn{
			matter.TableColumnName,
			matter.TableColumnType,
		}},
	clusterTable: {
		parent: documentTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnHierarchy,
			matter.TableColumnRole,
			matter.TableColumnScope,
			matter.TableColumnPICS,
		}},
	featureTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnBit,
			matter.TableColumnCode,
			matter.TableColumnFeature,
			matter.TableColumnSummary,
		},
	},
	dataTypeTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnName,
			matter.TableColumnType,
		},
	},
	bitmapValue: {
		parent: dataTypeTable,
		columns: []matter.TableColumn{
			matter.TableColumnBit,
			matter.TableColumnName,
			matter.TableColumnSummary,
			matter.TableColumnConformance,
		},
	},
	enumValue: {
		parent: dataTypeTable,
		columns: []matter.TableColumn{
			matter.TableColumnValue,
			matter.TableColumnName,
			matter.TableColumnSummary,
			matter.TableColumnConformance,
		},
	},
	structField: {
		parent: dataTypeTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnType,
			matter.TableColumnConstraint,
			matter.TableColumnQuality,
			matter.TableColumnDefault,
			matter.TableColumnAccess,
			matter.TableColumnConformance,
		},
	},
	attributeTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnType,
			matter.TableColumnConstraint,
			matter.TableColumnQuality,
			matter.TableColumnDefault,
			matter.TableColumnAccess,
			matter.TableColumnConformance,
		},
	},
	eventTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnPriority,
			matter.TableColumnAccess,
			matter.TableColumnConformance,
		},
	},
	eventFieldTable: {
		parent: eventTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnType,
			matter.TableColumnConstraint,
			matter.TableColumnQuality,
			matter.TableColumnDefault,
			matter.TableColumnConformance,
		},
	},
	commandTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnDirection,
			matter.TableColumnResponse,
			matter.TableColumnAccess,
			matter.TableColumnConformance,
		},
	},
	commandFieldTable: {
		parent: commandTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnType,
			matter.TableColumnConstraint,
			matter.TableColumnQuality,
			matter.TableColumnDefault,
			matter.TableColumnConformance,
		},
	},
}
