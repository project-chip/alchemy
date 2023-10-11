package db

import "github.com/hasty/matterfmt/matter"

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
}
