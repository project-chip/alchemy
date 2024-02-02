package db

import (
	mms "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func getAccessSchemaColumns(tableName string) []*mms.Column {
	return []*mms.Column{
		{Name: "read_access", Type: types.Int8, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "write_access", Type: types.Int8, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "invoke_access", Type: types.Int8, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "fabric_scoped", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "fabric_sensitive", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "timed", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
	}
}

func getAccessSchemaColumnValues(tableName string, access interface{}) []interface{} {
	var readAccess, writeAccess, invokeAccess, fabricScoped, fabricSensitive, timed int8
	s, ok := access.(string)
	if ok {
		var a matter.Access
		switch tableName {
		case commandTable:
			a = ascii.ParseAccess(s, mattertypes.EntityTypeCommand)
		case structField:
			a = ascii.ParseAccess(s, mattertypes.EntityTypeStruct)
		default:
			a = ascii.ParseAccess(s, mattertypes.EntityTypeUnknown)
		}
		readAccess = int8(a.Read)
		writeAccess = int8(a.Write)
		invokeAccess = int8(a.Invoke)
		if a.IsFabricScoped() {
			fabricScoped = 1
		}
		if a.IsFabricSensitive() {
			fabricSensitive = 1
		}

	}
	return []interface{}{readAccess, writeAccess, invokeAccess, fabricScoped, fabricSensitive, timed}
}
