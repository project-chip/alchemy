package db

import (
	mms "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
)

func getQualitySchemaColumns(tableName string) []*mms.Column {
	return []*mms.Column{
		{Name: "nullable", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "non_volatile", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "fixed", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "scene", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "reportable", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "changes_omitted", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "singleton", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
	}
}

func getQualitySchemaColumnValues(access interface{}) []interface{} {
	var nullable, nonVolatile, fixed, scene, reportable, changesOmitted, singleton int8
	if s, ok := access.(string); ok {
		rs := []rune(s)
		for _, r := range rs {
			switch r {
			case 'X':
				nullable = 1
			case 'N':
				nonVolatile = 1
			case 'F':
				fixed = 1
			case 'S':
				scene = 1
			case 'P':
				reportable = 1
			case 'C':
				changesOmitted = 1
			case 'I':
				singleton = 1
			}
		}
	}
	return []interface{}{nullable, nonVolatile, fixed, scene, reportable, changesOmitted, singleton}
}
