package db

import (
	"context"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func (h *Host) indexCluster(cxt context.Context, ds *sectionInfo, top *ascii.Section) error {
	ci := &sectionInfo{id: h.nextId(clusterTable), parent: ds, values: &dbRow{}}
	for _, s := range ascii.Skim[*ascii.Section](top.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionClusterID:
			appendSectionToRow(cxt, s, ci.values)
		case matter.SectionClassification:
			appendSectionToRow(cxt, s, ci.values)
		case matter.SectionFeatures:
			h.readTableSection(cxt, ci, s, featureTable)
		case matter.SectionDataTypes:
			h.indexDataTypes(cxt, ci, s)
		case matter.SectionEvents:
			h.indexEvents(cxt, ci, s)
		case matter.SectionCommands:
			h.indexCommands(cxt, ci, s)
		}
		if err != nil {
			return err
		}
	}
	for _, s := range ascii.Skim[*ascii.Section](top.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionAttributes:
			err = h.readTableSection(cxt, ci, s, attributeTable)
		}
		if err != nil {
			return err
		}
	}
	ds.children[clusterTable] = append(ds.children[clusterTable], ci)
	return nil
}



