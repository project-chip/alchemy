package db

import (
	mms "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/hasty/matterfmt/matter"
)

func getAccessSchemaColumns(tableName string) []*mms.Column {
	return []*mms.Column{
		{Name: "read_access", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "write_access", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "fabric_scoped", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "fabric_sensitive", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "timed", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
	}
}

func getAccessSchemaColumnValues(tableName string, access interface{}) []interface{} {
	readAccess, writeAccess, fabricScoped, fabricSensitive, timed := matter.ExtractAccessValues(access)
	return []interface{}{readAccess, writeAccess, fabricScoped, fabricSensitive, timed}
}
