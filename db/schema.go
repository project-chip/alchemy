package db

import "github.com/project-chip/alchemy/matter"

var (
	documentTable                     = "document"
	clusterTable                      = "cluster"
	clusterRevisionTable              = "cluster_revision"
	featureTable                      = "feature"
	structTable                       = "struct"
	structField                       = "struct_field"
	bitmapTable                       = "bitmap"
	bitmapValue                       = "bitmap_value"
	enumTable                         = "enum"
	enumValue                         = "enum_value"
	attributeTable                    = "attribute"
	eventTable                        = "event"
	eventFieldTable                   = "event_field"
	commandTable                      = "command"
	commandFieldTable                 = "command_field"
	namespaceTable                    = "namespace"
	tagTable                          = "tag"
	deviceTypeTable                   = "device_type"
	deviceTypeRevisionTable           = "device_type_revision"
	deviceTypeConditionTable          = "device_type_condition"
	deviceTypeClusterRequirementTable = "device_type_cluster_requirement"
	deviceTypeCompositionTable        = "device_type_composition"
	typedefTable                      = "typedef"
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
	clusterRevisionTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnDescription,
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
	bitmapTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnName,
			matter.TableColumnType,
		},
	},
	bitmapValue: {
		parent: bitmapTable,
		columns: []matter.TableColumn{
			matter.TableColumnBit,
			matter.TableColumnName,
			matter.TableColumnDescription,
			matter.TableColumnSummary,
			matter.TableColumnConformance,
		},
	},
	enumTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnName,
			matter.TableColumnDescription,
			matter.TableColumnType,
		},
	},
	enumValue: {
		parent: enumTable,
		columns: []matter.TableColumn{
			matter.TableColumnValue,
			matter.TableColumnName,
			matter.TableColumnSummary,
			matter.TableColumnConformance,
		},
	},
	structTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnName,
			matter.TableColumnDescription,
			matter.TableColumnScope,
		},
	},
	structField: {
		parent: structTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnType,
			matter.TableColumnConstraint,
			matter.TableColumnQuality,
			matter.TableColumnFallback,
			matter.TableColumnAccess,
			matter.TableColumnConformance,
		},
	},
	typedefTable: {
		parent: clusterTable,
		columns: []matter.TableColumn{
			matter.TableColumnName,
			matter.TableColumnDescription,
			matter.TableColumnType,
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
			matter.TableColumnFallback,
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
			matter.TableColumnFallback,
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
			matter.TableColumnQuality,
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
			matter.TableColumnFallback,
			matter.TableColumnConformance,
		},
	},
	deviceTypeTable: {
		parent: documentTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnSupersetOf,
			matter.TableColumnClass,
			matter.TableColumnScope,
		},
	},
	deviceTypeRevisionTable: {
		parent: deviceTypeTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnDescription,
		}},
	deviceTypeConditionTable: {
		parent: deviceTypeTable,
		columns: []matter.TableColumn{
			matter.TableColumnFeature,
			matter.TableColumnDescription,
		}},
	deviceTypeClusterRequirementTable: {
		parent: deviceTypeTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnQuality,
			matter.TableColumnConformance,
			matter.TableColumnDirection,
		},
	},
	deviceTypeCompositionTable: {
		parent: deviceTypeTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
			matter.TableColumnConstraint,
			matter.TableColumnConformance,
		},
	},
	namespaceTable: {
		parent: documentTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
		},
	},
	tagTable: {
		parent: namespaceTable,
		columns: []matter.TableColumn{
			matter.TableColumnID,
			matter.TableColumnName,
		},
	},
}
