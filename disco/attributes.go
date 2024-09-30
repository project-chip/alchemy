package disco

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) organizeAttributesSection(cxt *discoContext, dp *docParse) (err error) {

	for _, attributes := range dp.attributes {
		attributesTable := attributes.table
		if attributesTable == nil || attributesTable.Element == nil {
			return
		}

		if attributesTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange attributes table without header row")
		}

		if len(attributesTable.ColumnMap) < 3 {
			return fmt.Errorf("can't rearrange attributes table with so few matches: %d", len(attributesTable.ColumnMap))
		}

		err = b.fixAccessCells(dp, attributes, types.EntityTypeAttribute)
		if err != nil {
			return err
		}

		err = b.fixConstraintCells(attributes.section, attributesTable)
		if err != nil {
			return err
		}

		err = b.fixConformanceCells(dp, attributes, attributesTable.Rows, attributesTable.ColumnMap)
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
	if b.errata.IgnoreSection(section.section.Name, errata.DiscoPurposeTableLinkIndexes) {
		return nil
	}
	if section.table == nil || section.table.Element == nil {
		return nil
	}
	attributeCells := make(map[string]*asciidoc.TableCell)
	nameIndex, ok := section.table.ColumnMap[matter.TableColumnName]
	if !ok {
		return nil
	}

	for _, row := range section.table.Rows {
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
			id = ide.AsciiDocString()
			if strings.HasPrefix(id, "_") {
				ok = false
			}
		}
		if !ok {
			label := normalizeAnchorLabel(name, nil)
			id = normalizeAnchorID(name, nil)
			spec.NewAnchor(b.doc, id, s.Base, section.section, label...).SyncToDoc(id)
		}
		icr := asciidoc.NewCrossReference(id)
		err := cell.SetElements(asciidoc.Set{icr})
		if err != nil {
			return err
		}

	}

	return nil
}
