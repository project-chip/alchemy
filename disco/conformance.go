package disco

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

func (b *Ball) fixConformanceCells(docParse *docParse, section *subSection, rows []*asciidoc.TableRow, columnMap spec.ColumnIndex) (err error) {
	if len(rows) < 2 {
		return
	}
	if b.errata.IgnoreSection(section.section.Name, errata.DiscoPurposeTableConformance) {
		return nil
	}
	conformanceIndex, ok := columnMap[matter.TableColumnConformance]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cell(conformanceIndex)
		vc, e := spec.RenderTableCell(cell)
		if e != nil {
			continue
		}

		conf := conformance.ParseConformance(vc)

		docParse.conformanceCache[cell] = conf

		cs := conf.ASCIIDocString()

		if cs != vc {
			setCellString(cell, cs)
		}

	}
	return
}

func disambiguateConformance(docParse *docParse) (err error) {
	globalChoices := make(map[string]string)
	parse.Traverse(docParse.doc, docParse.doc.Elements(), func(table *asciidoc.Table, parent parse.HasElements, index int) parse.SearchShould {
		ti, ok := docParse.tableCache[table]
		if !ok {

			var err error
			ti, err = spec.ReadTable(docParse.doc, table)
			if err != nil {
				return parse.SearchShouldContinue
			}
		}
		conformanceIndex, ok := ti.ColumnMap[matter.TableColumnConformance]
		if !ok {
			return parse.SearchShouldContinue
		}
		conformanceCounter := 1
		localChoices := make(map[string]string)
		for i, row := range ti.Rows {
			if i == ti.HeaderRowIndex {
				continue
			}
			cell := row.Cell(conformanceIndex)
			conf, ok := docParse.conformanceCache[cell]
			if !ok {
				vc, e := spec.RenderTableCell(cell)
				if e != nil {
					continue
				}
				conf = conformance.ParseConformance(vc)
			}
			var modified bool
			for _, c := range conf {
				switch c := c.(type) {
				case *conformance.Optional:
					if c.Choice != nil {
						localChoice, ok := localChoices[c.Choice.Set]
						if ok {
							if c.Choice.Set != localChoice {
								c.Choice.Set = localChoice
								modified = true
							}
							break
						}

						_, ok = globalChoices[c.Choice.Set]
						if ok {
							for {
								var nextChoice string
								nextChoice, conformanceCounter = nextConformanceChoice(conformanceCounter)
								_, ok = globalChoices[nextChoice]
								if !ok {
									localChoices[c.Choice.Set] = nextChoice
									c.Choice.Set = nextChoice
									modified = true
									break
								}
							}
						}
						localChoices[c.Choice.Set] = c.Choice.Set
						globalChoices[c.Choice.Set] = c.Choice.Set
					}
				}
			}
			if modified {
				setCellString(cell, conf.ASCIIDocString())
			}
		}
		return parse.SearchShouldContinue
	})
	return nil
}

func nextConformanceChoice(current int) (set string, next int) {
	next = current + 1
	var result strings.Builder
	for {
		current--
		remainder := current % 26
		result.WriteByte('a' + byte(remainder))
		current /= 26
		if current == 0 {
			break
		}
	}
	set = result.String()
	r := []rune(set)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	set = string(r)
	return
}
