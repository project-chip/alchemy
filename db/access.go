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
	readAccess, writeAccess, fabricScoped, timed := extractAccessValues(access)
	return []interface{}{readAccess, writeAccess, fabricScoped, timed}
}

func extractAccessValues(access interface{}) (readAccess string, writeAccess string, fabricScoped int8, timed int8) {
	s, ok := access.(string)
	if !ok {
		return
	}
	am := parse.ParseAccess(s)
	if am == nil {
		return
	}
	if am[matter.AccessCategoryFabric] == "F" {
		fabricScoped = 1
	}
	if am[matter.AccessCategoryTimed] == "T" {
		timed = 1
	}
	rw := am[matter.AccessCategoryReadWrite]
	var hasRead, hasWrite bool
	switch rw {
	case "RW", "R[W]":
		hasRead = true
		hasWrite = true
	case "R":
		hasRead = true
	case "W":
		hasWrite = true
	}
	ps, ok := am[matter.AccessCategoryPrivileges]
	if !ok {
		return
	}
	for _, r := range []rune(ps) {
		if hasRead {
			readAccess = string(r)
			hasRead = false
		} else if hasWrite {
			writeAccess = string(r)
			break
		}
	}

	return
}
