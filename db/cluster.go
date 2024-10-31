package db

import (
	"context"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (h *Host) indexClusterModel(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	if !cluster.ID.Valid() {
		return nil
	}
	clusterRow := newDBRow()
	clusterRow.values[matter.TableColumnID] = cluster.ID.IntString()
	clusterRow.values[matter.TableColumnName] = cluster.Name
	clusterRow.values[matter.TableColumnHierarchy] = cluster.Hierarchy
	clusterRow.values[matter.TableColumnRole] = cluster.Role
	clusterRow.values[matter.TableColumnScope] = cluster.Scope
	clusterRow.values[matter.TableColumnPICS] = cluster.PICS

	ci := &sectionInfo{id: h.nextID(clusterTable), parent: parent, values: clusterRow, children: make(map[string][]*sectionInfo)}

	if cluster.Features != nil {
		for _, f := range cluster.Features.Bits {
			featureRow := newDBRow()
			featureRow.values[matter.TableColumnBit] = f.Bit()
			featureRow.values[matter.TableColumnName] = f.Name()
			featureRow.values[matter.TableColumnFeature] = f.Name()
			featureRow.values[matter.TableColumnSummary] = f.Summary()
			fci := &sectionInfo{id: h.nextID(featureTable), parent: ci, values: featureRow}
			ci.children[featureTable] = append(ci.children[featureTable], fci)
		}
	}

	for _, a := range cluster.Attributes {
		h.readField(a, ci, attributeTable, types.EntityTypeAttribute)
	}

	h.indexClusterRevisionsModel(cluster, ci)

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

func (h *Host) indexClusterRevisionsModel(cluster *matter.Cluster, parent *sectionInfo) {
	for _, rev := range cluster.Revisions {
		row := newDBRow()
		row.values[matter.TableColumnID] = rev.Number
		row.values[matter.TableColumnDescription] = rev.Description
		bi := &sectionInfo{id: h.nextID(clusterRevisionTable), parent: parent, values: row}
		parent.children[clusterRevisionTable] = append(parent.children[clusterRevisionTable], bi)
	}
}

func (h *Host) indexCluster(cxt context.Context, doc *spec.Doc, ds *sectionInfo, top *spec.Section) error {
	ci := &sectionInfo{id: h.nextID(clusterTable), parent: ds, values: &dbRow{}}
	for _, s := range parse.Skim[*spec.Section](top.Elements()) {
		var err error
		switch s.SecType {
		case matter.SectionClusterID:
			err = appendSectionToRow(cxt, doc, s, ci.values)
		case matter.SectionRevisionHistory:
			err = h.readTableSection(cxt, doc, ci, s, clusterRevisionTable)
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
	for _, s := range parse.Skim[*spec.Section](top.Elements()) {
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
