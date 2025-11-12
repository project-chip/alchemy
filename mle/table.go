package mle

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func findTableWithColumns(doc *asciidoc.Document, requiredColumns []matter.TableColumn) (matchingTable *spec.TableInfo, err error) {
	for table := range parse.FindAll[*asciidoc.Table](doc, asciidoc.RawReader, doc) {
		var ti *spec.TableInfo
		ti, err = spec.ReadTable(doc, asciidoc.RawReader, table)
		if err != nil {
			return
		}
		if ti.ColumnMap.HasAll(requiredColumns...) {
			matchingTable = ti
			break
		}
	}
	return
}
