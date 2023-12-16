package db

import (
	"context"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (h *Host) indexClusterModel(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	clusterRow := newDBRow()
	clusterRow.values[matter.TableColumnID] = cluster.ID.IntString()
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
		featureRow.values[matter.TableColumnSummary] = f.Summary
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

func (h *Host) indexCluster(cxt context.Context, doc *ascii.Doc, ds *sectionInfo, top *ascii.Section) error {
	ci := &sectionInfo{id: h.nextId(clusterTable), parent: ds, values: &dbRow{}}
	for _, s := range parse.Skim[*ascii.Section](top.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionClusterID:
			err = appendSectionToRow(cxt, doc, s, ci.values)
		case matter.SectionClassification:
			err = appendSectionToRow(cxt, doc, s, ci.values)
		case matter.SectionFeatures:
			err = h.readTableSection(cxt, doc, ci, s, featureTable)
		case matter.SectionDataTypes:
			err = h.indexDataTypes(cxt, doc, ci, s)
		case matter.SectionEvents:
			err = h.indexEvents(cxt, doc, ci, s)
		case matter.SectionCommands:
			err = h.indexCommands(cxt, doc, ci, s)
		}
		if err != nil {
			return err
		}
	}
	for _, s := range parse.Skim[*ascii.Section](top.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionAttributes:
			err = h.readTableSection(cxt, doc, ci, s, attributeTable)
		}
		if err != nil {
			return err
		}
	}
	ds.children[clusterTable] = append(ds.children[clusterTable], ci)
	return nil
}
