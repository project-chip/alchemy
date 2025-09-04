package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
)

func (b *Baller) organizeTable(cxt *discoContext, name string, sections []*subSection, tableType matter.TableType) (err error) {
	for _, section := range sections {
		sectionTable := section.table
		if sectionTable == nil || sectionTable.Element == nil {
			slog.Warn("Could not organize section, as no table was found", slog.String("section", name), log.Path("source", section.section))
			return
		}

		if sectionTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange %s table without header row in %s", name, cxt.doc.Path.Relative)
		}

		if len(sectionTable.ColumnMap) < 2 {
			return fmt.Errorf("can't rearrange %s table with so few matches in %s", name, cxt.doc.Path.Relative)
		}

		err = b.renameTableHeaderCells(cxt, section.section, sectionTable, matter.Tables[tableType].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in %s table in %s: %w", name, cxt.doc.Path, err)
		}

		err = b.reorderColumns(cxt, section.section, sectionTable, tableType)
		if err != nil {
			return err
		}
	}
	return
}

func (b *Baller) organizeDeviceIDSection(cxt *discoContext) (err error) {
	return b.organizeTable(cxt, "device ID", cxt.parsed.deviceIDs, matter.TableTypeDeviceID)
}

func (b *Baller) organizeClusterRequirementsSection(cxt *discoContext) (err error) {
	return b.organizeTable(cxt, "cluster requirement", cxt.parsed.clusterRequirements, matter.TableTypeClusterRequirements)
}

func (b *Baller) organizeElementRequirementsSection(cxt *discoContext) (err error) {
	return b.organizeTable(cxt, "element requirement", cxt.parsed.elementRequirements, matter.TableTypeElementRequirements)
}

func (b *Baller) organizeComposedDeviceClusterRequirementsSection(cxt *discoContext) (err error) {
	return b.organizeTable(cxt, "composed device cluster requirement", cxt.parsed.composedClusterRequirements, matter.TableTypeComposedDeviceTypeClusterRequirements)
}

func (b *Baller) organizeComposedDeviceElementRequirementsSection(cxt *discoContext) (err error) {
	return b.organizeTable(cxt, "composed device element requirement", cxt.parsed.composedElementRequirements, matter.TableTypeComposedDeviceTypeElementRequirements)
}

func (b *Baller) organizeConditionRequirementsSection(cxt *discoContext) (err error) {
	return b.organizeTable(cxt, "condition requirement", cxt.parsed.conditionRequirements, matter.TableTypeConditionRequirements)
}
