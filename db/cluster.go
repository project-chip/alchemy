package db

import (
	"context"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (h *Host) indexClusterModel(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	clusterRow := newDBRow()
	clusterRow.values[matter.TableColumnID] = cluster.ID
	clusterRow.values[matter.TableColumnName] = cluster.Name
	clusterRow.values[matter.TableColumnHierarchy] = cluster.Hierarchy
	clusterRow.values[matter.TableColumnRole] = cluster.Role
	clusterRow.values[matter.TableColumnScope] = cluster.Scope
	clusterRow.values[matter.TableColumnPICS] = cluster.PICS

	ci := &sectionInfo{id: h.nextId(clusterTable), parent: parent, values: clusterRow, children: make(map[string][]*sectionInfo)}

	for _, f := range cluster.Features {
		featureRow := newDBRow()
		featureRow.values[matter.TableColumnBit] = f.Bit
		featureRow.values[matter.TableColumnCode] = f.Code
		featureRow.values[matter.TableColumnFeature] = f.Name
		featureRow.values[matter.TableColumnSummary] = f.Description
		fci := &sectionInfo{id: h.nextId(featureTable), parent: parent, values: featureRow}
		parent.children[featureTable] = append(parent.children[featureTable], fci)
	}

	for _, a := range cluster.Attributes {
		h.readField(a, ci, attributeTable)
	}

	err := h.indexDataTypeModels(cxt, ci, cluster)
	if err != nil {
		return err
	}
	err = h.indexEventModels(cxt, ci, cluster)
	if err != nil {
		return err
	}

	err = h.indexCommandModels(cxt, ci, cluster)
	if err != nil {
		return err
	}

	parent.children[clusterTable] = append(parent.children[clusterTable], ci)
	return nil
}

func (h *Host) indexCluster(cxt context.Context, ds *sectionInfo, top *ascii.Section) error {
	ci := &sectionInfo{id: h.nextId(clusterTable), parent: ds, values: &dbRow{}}
	for _, s := range parse.Skim[*ascii.Section](top.Elements) {
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
	for _, s := range parse.Skim[*ascii.Section](top.Elements) {
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
