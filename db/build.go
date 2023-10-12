package db

import (
	"log/slog"

	"github.com/dolthub/go-mysql-server/sql"
)

func (h *Host) Build(sc *sql.Context) error {
	h.base = &sectionInfo{children: make(map[string][]*sectionInfo)}
	var sis []*sectionInfo
	for _, d := range h.docs {
		slog.InfoContext(sc, "Indexing", "path", d.Path)
		si, err := h.indexDoc(sc, d)
		if err != nil {
			return err
		}
		sis = append(sis, si)

	}
	h.base.children[documentTable] = sis
	return h.createTables(sc, h.base)
}

func (h *Host) createTables(sc *sql.Context, bs *sectionInfo) error {
	for _, tableName := range h.tableNames {
		ts, ok := tableSchema[tableName]
		if !ok {
			slog.Error("Table missing", "name", tableName)
			continue
		}
		sis := findSectionInfos(bs, tableName)
		err := h.createTable(sc, tableName, ts.parent, sis, ts.columns)
		if err != nil {
			return err
		}
	}
	return nil
}

func findSectionInfos(base *sectionInfo, name string) []*sectionInfo {
	var list []*sectionInfo
	si, ok := base.children[name]
	if ok {
		list = append(list, si...)
	}
	for _, c := range base.children {
		for _, s := range c {
			list = append(list, findSectionInfos(s, name)...)
		}
	}
	return list
}
