package disco

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
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
	attributeCells := make(map[string]*asciidoc.TableCell)
	nameIndex, ok := section.table.columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}

	for _, row := range section.table.rows {
		cell := row.Cell(nameIndex)
		cv, err := spec.RenderTableCell(cell)
		if err != nil {
			continue
		}

		name := strings.TrimSpace(cv)
		nameKey := strings.ToLower(name)

		if _, ok := attributeCells[nameKey]; !ok {
			attributeCells[nameKey] = cell
		}
	}
	for _, ss := range section.children {
		s := ss.section
		attributeName := matter.StripReferenceSuffixes(s.Name)
		name := strings.TrimSpace(attributeName)

		cell, ok := attributeCells[strings.ToLower(name)]
		if !ok {
			continue
		}
		var id string
		ide := s.Base.GetAttributeByName(asciidoc.AttributeNameID)
		if ide != nil {
			idv, ok := ide.Value().(*asciidoc.String)
			if ok {
				id = idv.Value
				if strings.HasPrefix(id, "_") {
					ok = false
				}
			}
		}
		if !ok {
			var label asciidoc.Set
			id, label = normalizeAnchorID(s.Name, nil, nil)
			setAnchorID(s.Base, id, label)
		}
		icr := asciidoc.NewCrossReference(id)
		err := cell.SetElements(asciidoc.Set{icr})
		if err != nil {
			return err
		}

	}

	return nil
}
