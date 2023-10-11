package db

import (
	"context"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

type sectionInfo struct {
	id     int32
	values *dbRow

	parent *sectionInfo

	children map[string][]*sectionInfo
}

func (h *Host) indexCluster(cxt context.Context, ds *sectionInfo, top *ascii.Section) error {
	ci := &sectionInfo{id: h.nextId(clusterTable), parent: ds, values: &dbRow{}}
	for _, s := range ascii.FindAll[*ascii.Section](top.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionClusterID:
			readSimpleSection(cxt, s, ci.values)
		case matter.SectionClassification:
			readSimpleSection(cxt, s, ci.values)
		case matter.SectionFeatures:
			h.readTableSection(cxt, ci, s, featureTable)
		}
		if err != nil {
			return err
		}
	}
	ds.children[clusterTable] = append(ds.children[clusterTable], ci)
	return nil
}
