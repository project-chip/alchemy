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
		attributesTable := &attributes.table
		if attributesTable.element == nil {
			return
		}

		if attributesTable.columnMap == nil {
			return fmt.Errorf("can't rearrange attributes table without header row")
		}

		if len(attributesTable.columnMap) < 3 {
			return fmt.Errorf("can't rearrange attributes table with so few matches: %d", len(attributesTable.columnMap))
		}

		err = b.fixAccessCells(dp.doc, attributesTable, mattertypes.EntityTypeAttribute)
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

func (b *Ball) linkAttributes(cxt *discoContext, section *ascii.Section, rows []*types.TableRow, columnMap ascii.ColumnIndex) error {
	attributeCells := make(map[string]*types.Paragraph)
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}

	for _, row := range rows {
		cell := row.Cells[nameIndex]
		cv, err := ascii.RenderTableCell(cell)
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
