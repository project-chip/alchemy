package disco

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func (b *Ball) organizeAttributesSection(cxt *discoContext, doc *ascii.Doc, top *ascii.Section, attributes *ascii.Section) error {
	attributesTable := ascii.FindFirstTable(attributes)
	if attributesTable == nil {
		return fmt.Errorf("no attributes table found")
	}
	return b.organizeAttributesTable(cxt, doc, top, attributes, attributesTable)
}

func (b *Ball) organizeAttributesTable(cxt *discoContext, doc *ascii.Doc, top *ascii.Section, attributes *ascii.Section, attributesTable *types.Table) error {
	rows := ascii.TableRows(attributesTable)

	_, columnMap, extraColumns, err := ascii.MapTableColumns(rows)
	if err != nil {
		return err
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange attributes table without header row")
	}

	if len(columnMap) < 5 {
		return fmt.Errorf("can't rearrange attributes table with so few matches")
	}

	err = b.fixAccessCells(doc, rows, columnMap)
	if err != nil {
		return err
	}

	err = fixConstraintCells(rows, columnMap)
	if err != nil {
		return err
	}

	if b.ShouldLinkAttributes {
		err = b.linkAttributes(cxt, attributes, rows, columnMap)
		if err != nil {
			return err
		}
	}

	err = getPotentialDataTypes(cxt, attributes, rows, columnMap)
	if err != nil {
		return err
	}

	reorderColumns(doc, attributes, rows, matter.AttributesTableColumnOrder[:], columnMap, extraColumns)

	return nil
}

func (b *Ball) fixAccessCells(doc *ascii.Doc, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (err error) {
	if len(rows) < 2 {
		return
	}
	accessIndex, ok := columnMap[matter.TableColumnAccess]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[accessIndex]
		vc, e := ascii.GetTableCellValue(cell)
		if e != nil {
			continue
		}
		access := matter.ParseAccess(vc)
		if len(access) > 0 {
			var out strings.Builder
			for _, ac := range matter.AccessCategoryOrder {
				a, ok := access[ac]
				if !ok {
					continue
				}
				if ac == matter.AccessCategoryReadWrite && a == "R*W" { // fix deprecated form
					a = "R[W]"
				}
				if out.Len() > 0 {
					out.WriteRune(' ')
				}
				out.WriteString(a)
			}
			err = setCellString(cell, out.String())
			if err != nil {
				return
			}
		}
	}
	return
}

var minMaxPattern = regexp.MustCompile(`^(Max|Min)( .*)$`)

func fixConstraintCells(rows []*types.TableRow, columnMap map[matter.TableColumn]int) (err error) {
	if len(rows) < 2 {
		return
	}
	constraintIndex, ok := columnMap[matter.TableColumnConstraint]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[constraintIndex]
		vc, e := ascii.GetTableCellValue(cell)
		if e != nil {
			continue
		}
		fixed := minMaxPattern.ReplaceAllStringFunc(vc, func(s string) string {
			r := []rune(s)
			r[0] = unicode.ToLower(r[0])
			return string(r)
		})
		if vc != fixed {
			setCellString(cell, fixed)
		}
	}
	return
}

func (b *Ball) linkAttributes(cxt *discoContext, section *ascii.Section, rows []*types.TableRow, columnMap map[matter.TableColumn]int) error {
	attributeCells := make(map[string]*types.Paragraph)
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}

	for _, row := range rows {
		cell := row.Cells[nameIndex]
		cv, err := ascii.GetTableCellValue(cell)
		if err != nil {
			continue
		}

		if len(cell.Elements) == 0 {
			continue
		}
		p, ok := cell.Elements[0].(*types.Paragraph)
		if !ok {
			continue
		}
		if len(p.Elements) == 0 {
			continue
		}
		_, ok = p.Elements[0].(*types.StringElement)
		if !ok {
			continue
		}

		name := strings.TrimSpace(cv)
		nameKey := strings.ToLower(name)

		if _, ok := attributeCells[nameKey]; !ok {
			attributeCells[nameKey] = p
		}
	}
	for _, el := range section.Elements {
		if s, ok := el.(*ascii.Section); ok {
			attributeName := matter.StripReferenceSuffixes(s.Name)
			name := strings.TrimSpace(attributeName)

			p, ok := attributeCells[strings.ToLower(name)]
			if !ok {
				continue
			}
			var id string
			ide, ok := s.Base.Attributes[types.AttrID]
			if ok {
				id, ok = ide.(string)
			}
			if !ok {
				var label string
				id, label = normalizeAnchorID(s.Name, nil)
				setAnchorID(s.Base, id, label)
			}
			icr, _ := types.NewInternalCrossReference(id, nil)
			p.SetElements([]interface{}{icr})
		}
	}

	return nil
}
