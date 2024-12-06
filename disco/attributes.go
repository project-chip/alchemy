package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) organizeAttributesSection(cxt *discoContext) (err error) {

	for _, attributes := range cxt.parsed.attributes {
		attributesTable := attributes.table
		if attributesTable == nil || attributesTable.Element == nil {
			slog.Warn("Could not organize Attributes section, as no table of attributes was found", log.Path("source", attributes.section.Base))
			return
		}

		if attributesTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange attributes table without header row")
		}

		if len(attributesTable.ColumnMap) < 3 {
			return fmt.Errorf("can't rearrange attributes table with so few matches: %d", len(attributesTable.ColumnMap))
		}

		err = b.renameTableHeaderCells(cxt, attributes.section, attributesTable, matter.Tables[matter.TableTypeAttributes].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", attributes.section.Name, cxt.doc.Path, err)
		}

		err = b.fixAccessCells(cxt, attributes, types.EntityTypeAttribute)
		if err != nil {
			return err
		}

		err = b.fixConstraintCells(cxt, attributes.section, attributesTable)
		if err != nil {
			return err
		}

		err = b.fixConformanceCells(cxt, attributes, attributesTable.Rows, attributesTable.ColumnMap)
		if err != nil {
			return err
		}

		err = b.linkIndexTables(cxt, attributes)
		if err != nil {
			return err
		}

		err = b.reorderColumns(cxt, attributes.section, attributesTable, matter.TableTypeAttributes)
		if err != nil {
			return err
		}

		b.removeMandatoryFallbacks(attributesTable)
	}
	return nil
}

func (b *Baller) linkIndexTables(cxt *discoContext, section *subSection) error {
	if !b.options.linkIndexTables {
		return nil
	}
	if cxt.errata.IgnoreSection(section.section.Name, errata.DiscoPurposeTableLinkIndexes) {
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
			spec.NewAnchor(cxt.doc, id, s.Base, section.section, label...).SyncToDoc(id)
		}
		icr := asciidoc.NewCrossReference(id)
		cell.SetElements(asciidoc.Set{icr})
	}

	return nil
}
