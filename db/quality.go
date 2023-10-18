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
		var val int8 = 1
		for _, r := range rs {
			switch r {
			case 'X':
				nullable = val
			case 'N':
				nonVolatile = val
			case 'F':
				fixed = val
			case 'S':
				scene = val
			case 'P':
				reportable = val
			case 'C':
				changesOmitted = val
			case 'I':
				singleton = val
			case '!':
				val = -1
				continue
			}
			if val == -1 {
				val = 1
			}
		}
	}
	return []interface{}{nullable, nonVolatile, fixed, scene, reportable, changesOmitted, singleton}
}
