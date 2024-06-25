package db

import (
	mms "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/project-chip/alchemy/matter"
)

func getBitmapSchemaColumns(tableName string) []*mms.Column {
	return []*mms.Column{
		{Name: "fromBit", Type: types.Uint64, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "toBit", Type: types.Uint64, Nullable: true, Source: tableName, PrimaryKey: false},
	}
}

func getBitmapSchemaColumnValues(b any) []any {

	bit, ok := b.(matter.Bit)
	if !ok {
		return []any{nil, nil}
	}
	from, to, err := bit.Bits()
	if err != nil {
		return []any{nil, nil}
	}
	return []any{from, to}
}
