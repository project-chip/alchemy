package disco

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (b *Ball) organizeAttributesSection(cxt *discoContext, dp *docParse) (err error) {

	for _, attributes := range dp.attributes {
		attributesTable := attributes.table
		if attributesTable.element == nil {
			return
		}

		if attributesTable.columnMap == nil {
			return fmt.Errorf("can't rearrange attributes table without header row")
		}

		if len(attributesTable.columnMap) < 3 {
			return fmt.Errorf("can't rearrange attributes table with so few matches: %d", len(attributesTable.columnMap))
		}

		err = b.fixAccessCells(dp.doc, attributesTable.rows, attributesTable.columnMap, mattertypes.EntityTypeAttribute)
		if err != nil {
			return err
		}

		err = fixConstraintCells(dp.doc, attributesTable.rows, attributesTable.columnMap)
		if err != nil {
			return err
		}

		err = fixConformanceCells(dp.doc, attributesTable.rows, attributesTable.columnMap)
		if err != nil {
			return err
		}

		if b.options.linkAttributes {
			err = b.linkAttributes(cxt, attributes.section, attributesTable.rows, attributesTable.columnMap)
			if err != nil {
				return err
			}
		}

		b.reorderColumns(dp.doc, attributes.section, attributesTable.rows, matter.AttributesTableColumnOrder[:], attributesTable.columnMap, attributesTable.extraColumns)
	}
	return nil
}

/*func (b *Ball) organizeAttributesSection(cxt *discoContext, doc *ascii.Doc, top *ascii.Section, attributes *ascii.Section) error {
	attributesTable := ascii.FindFirstTable(attributes)
	if attributesTable == nil {
		slog.Debug("no attributes table found", "sectionName", top.Name)
		return nil
	}
	return b.organizeAttributesTable(cxt, doc, top, attributes, attributesTable)
}

func (b *Ball) organizeAttributesTable(cxt *discoContext, doc *ascii.Doc, top *ascii.Section, attributes *ascii.Section, attributesTable *types.Table) error {
	rows := ascii.TableRows(attributesTable)

	_, columnMap, extraColumns, err := ascii.MapTableColumns(doc, rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for attributes table in section %s: %w", top.Name, err)
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange attributes table without header row")
	}

	if len(columnMap) < 3 {
		return fmt.Errorf("can't rearrange attributes table with so few matches: %d", len(columnMap))
	}

	err = b.fixAccessCells(doc, rows, columnMap, mattertypes.EntityTypeAttribute)
	if err != nil {
		return err
	}

	err = fixConstraintCells(doc, rows, columnMap)
	if err != nil {
		return err
	}

	err = fixConformanceCells(doc, rows, columnMap)
	if err != nil {
		return err
	}

	if b.options.linkAttributes {
		err = b.linkAttributes(cxt, attributes, rows, columnMap)
		if err != nil {
			return err
		}
	}

	b.reorderColumns(doc, attributes, rows, matter.AttributesTableColumnOrder[:], columnMap, extraColumns)

	return nil
}*/

func (b *Ball) linkAttributes(cxt *discoContext, section *ascii.Section, rows []*types.TableRow, columnMap ascii.ColumnIndex) error {
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
				id, label = normalizeAnchorID(s.Name, nil, nil)
				setAnchorID(s.Base, id, label)
			}
			icr, _ := types.NewInternalCrossReference(id, nil)
			err := p.SetElements([]interface{}{icr})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
