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

		err = b.fixAccessCells(dp, attributes, mattertypes.EntityTypeAttribute)
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

		err = b.linkIndexTables(cxt, attributes)
		if err != nil {
			return err
		}

		err = b.reorderColumns(dp.doc, attributes.section, attributesTable, matter.TableTypeAttributes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Ball) linkIndexTables(cxt *discoContext, section *subSection) error {
	if !b.options.linkIndexTables {
		return nil
	}
	if section.table.element == nil {
		return nil
	}
	attributeCells := make(map[string]*types.Paragraph)
	nameIndex, ok := section.table.columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}

	for _, row := range section.table.rows {
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
	for _, ss := range section.children {
		s := ss.section
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
			if ok && strings.HasPrefix(id, "_") {
				ok = false
			}
		}
		if !ok {
			var label string
			id, label = normalizeAnchorID(s.Name, nil, nil)
			setAnchorID(s.Base, id, label)
		}
		icr, _ := types.NewInternalCrossReference(id, nil)
		err := p.SetElements([]any{icr})
		if err != nil {
			return err
		}

	}

	return nil
}
