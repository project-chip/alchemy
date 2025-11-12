package mle

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type clusterInfo struct {
	ClusterID *matter.Number
	Name      string
	PICScode  string
}

func parseMasterClusterList(filePath asciidoc.Path) (ci map[uint64]clusterInfo, reserveds map[uint64]log.Source, violations map[string][]spec.Violation, err error) {
	ci = make(map[uint64]clusterInfo)
	uniqueNames := make(map[string]*matter.Number)
	uniquePics := make(map[string]*matter.Number)
	violations = make(map[string][]spec.Violation)
	reserveds = make(map[uint64]log.Source)

	requiredColumns := []matter.TableColumn{matter.TableColumnClusterID, matter.TableColumnClusterName, matter.TableColumnPICSCode}

	doc, err := parse.File(filePath)
	if err != nil {
		return
	}

	var clusterListTable *spec.TableInfo
	clusterListTable, err = findTableWithColumns(doc, requiredColumns)
	if err != nil {
		return
	}

	if clusterListTable == nil {
		slog.Error("No cluster master list table found.")
		return
	}

	for row := range clusterListTable.Body() {
		var id *matter.Number
		var name, pics string
		id, err = clusterListTable.ReadID(asciidoc.RawReader, row, matter.TableColumnClusterID)
		if err != nil {
			return
		}
		if !id.Valid() {
			continue
		}
		name, _, err = clusterListTable.ReadName(asciidoc.RawReader, row, matter.TableColumnClusterName)
		if err != nil {
			return
		}
		switch name {
		case "":
			continue
		case "Reserved":
			if _, taken := reserveds[id.Value()]; taken {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Duplicate reserved Cluster ID in Master List. ID='%s'", id.HexString())}
				v.Path, v.Line = row.Origin()
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			if _, taken := ci[id.Value()]; taken {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is both reserved and in use in Master List. ID='%s'", id.HexString())}
				v.Path, v.Line = row.Origin()
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			reserveds[id.Value()] = row
			continue
		}
		pics, err = clusterListTable.ReadString(asciidoc.RawReader, row, matter.TableColumnPICSCode)
		if err != nil {
			return
		}
		if _, reserved := reserveds[id.Value()]; reserved {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is both reserved and in use in Master List. ID='%s'", id.HexString())}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if _, taken := ci[id.Value()]; taken {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is duplicated on Master List. ID='%s'", id.HexString())}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if _, taken := uniqueNames[name]; taken {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster Name is duplicated on Master List. name='%s'", name)}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if pics != "" {
			if _, taken := uniquePics[pics]; taken {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster PICS is duplicated on Master List. PICS='%s'", pics)}
				v.Path, v.Line = row.Origin()
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			uniquePics[pics] = id
		}
		ci[id.Value()] = clusterInfo{
			ClusterID: id,
			Name:      name,
			PICScode:  pics,
		}
		uniqueNames[name] = id
	}

	return
}
