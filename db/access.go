package db

import (
	mms "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func getAccessSchemaColumns(tableName string) []*mms.Column {
	return []*mms.Column{
		{Name: "read_access", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "write_access", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "fabric_scoped", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "timed", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
	}
}

func getAccessSchemaColumnValues(tableName string, access interface{}) []interface{} {
	var readAccess, writeAccess string
	var fabricScoped, timed int8
	if s, ok := access.(string); ok {
		am := parse.ParseAccess(s)
		if am != nil {
			fab := am[matter.AccessCategoryFabric]
			if fab == "F" {
				fabricScoped = 1
			}
		}
	}
	return []interface{}{readAccess, writeAccess, fabricScoped, timed}
}
