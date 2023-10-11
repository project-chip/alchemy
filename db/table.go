package db

import (
	"log/slog"
	"strings"

	"github.com/dolthub/go-mysql-server/memory"
	mms "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
	"github.com/iancoleman/strcase"
)

var (
	documentTable = "document"
	clusterTable  = "cluster"
	featureTable  = "feature"
)

func (h *Host) createTable(cxt *mms.Context, tableName string, parentName string, sections []*sectionInfo, columns []matter.TableColumn) error {
	extraColumns := make(map[string]*parse.ExtraColumn)
	for _, si := range sections {
		for e := range si.values.extras {
			extraColumns[strings.ToLower(e)] = &parse.ExtraColumn{Name: e}
		}
	}
	schema := mms.Schema{
		{Name: tableName + "_id", Type: types.Int32, Nullable: false, Source: tableName, PrimaryKey: true},
	}
	if len(parentName) > 0 {
		schema = append(schema, &mms.Column{Name: parentName + "_id", Type: types.Int32, Nullable: false, Source: tableName, PrimaryKey: true})
	}

	for _, col := range columns {
		name, ok := matter.TableColumnNames[col]
		if !ok {
			continue
		}
		var colType mms.Type = types.Text
		switch col {
		case matter.TableColumnType, matter.TableColumnID:
			colType = types.Int64
		}
		schema = append(schema, &mms.Column{Name: strcase.ToSnake(name), Type: colType, Nullable: true, Source: tableName, PrimaryKey: false})
	}

	offset := len(schema)
	var extra []parse.ExtraColumn
	for _, e := range extraColumns {
		schema = append(schema, &mms.Column{Name: strcase.ToSnake(e.Name), Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: true})
		extra = append(extra, parse.ExtraColumn{Name: e.Name})
		e.Offset = offset
		offset++
	}
	for _, s := range schema {
		slog.Info("adding column", "table", s.Source, "column", s)
	}
	t := memory.NewTable(tableName, mms.NewPrimaryKeySchema(schema), h.db.GetForeignKeyCollection())
	h.tables[tableName] = t
	h.db.AddTable(tableName, t)
	for _, si := range sections {
		row := mms.NewRow(si.id)
		offset := 1
		if len(parentName) > 0 {
			row = append(row, si.parent.id)
			offset++
		}
		for _, col := range columns {
			sr := schema[offset]
			v, ok := si.values.values[col]
			if ok {
				var err error
				switch sr.Type {
				case types.Int64:
					switch val := v.(type) {
					case matter.DocType:
						v = int64(val)
					case string:
						v, err = parseNumber(val)
					}
				}
				if err != nil {
					v = nil
				}
				row = append(row, v)
			} else {
				slog.Info("missing col", "name", col)
				row = append(row, nil)
			}
			offset++
		}
		for _, e := range extra {
			v, ok := si.values.extras[e.Name]
			if ok {
				row = append(row, v)
			} else {
				slog.Info("missing val", "name", e.Name)
				row = append(row, nil)
			}
			offset++
		}
		slog.Info("inserting row", "table", tableName, "row", row)
		err := t.Insert(cxt, row)
		if err != nil {
			return err
		}

	}
	return nil
}
