package db

import (
	"context"

	"github.com/project-chip/alchemy/matter"
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

	ci := h.newSectionInfo(clusterTable, parent, clusterRow, cluster)

	if cluster.Features != nil {
		for _, f := range cluster.Features.Bits {
			featureRow := newDBRow()
			featureRow.values[matter.TableColumnBit] = f.Bit()
			featureRow.values[matter.TableColumnName] = f.Name()
			featureRow.values[matter.TableColumnFeature] = f.Name()
			featureRow.values[matter.TableColumnSummary] = f.Summary()
			fci := h.newSectionInfo(featureTable, ci, featureRow, f)

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
		bi := h.newSectionInfo(clusterRevisionTable, parent, row, rev)
		parent.children[clusterRevisionTable] = append(parent.children[clusterRevisionTable], bi)
	}
}
