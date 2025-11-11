package mle

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter/spec"
)

type clusterInfo struct {
	ClusterID string
	PICScode  string
}

func clusterIdTaken(id string, masterClusterMap map[string]clusterInfo) (taken bool, name string) {
	var ci clusterInfo
	if id == "n/a" {
		taken = false
		return
	}
	for name, ci = range masterClusterMap {
		if ci.ClusterID == id {
			taken = true
			return
		}
	}
	taken = false
	return
}

func clusterPICSTaken(p string, masterClusterMap map[string]clusterInfo) (taken bool, name string) {
	var ci clusterInfo
	if p == "" {
		taken = false
		return
	}
	for name, ci = range masterClusterMap {
		if ci.PICScode == p {
			taken = true
			return
		}
	}
	taken = false
	return
}

func clusterIdReserved(id string, reserveds []string) (reserved bool) {
	for _, r := range reserveds {
		if r == id {
			reserved = true
			return
		}
	}
	reserved = false
	return
}

func parseMasterClusterList(filePath string) (ci map[string]clusterInfo, reserveds []string, violations map[string][]spec.Violation, err error) {
	ci = make(map[string]clusterInfo)
	reserveds = make([]string, 0)
	violations = make(map[string][]spec.Violation)

	requiredColumns := []string{"Cluster ID", "Cluster Name", "PICS Code"}

	doc, err := parse.File(filePath)
	if err != nil {
		return
	}

	t := findTableWithColumns(doc.Elements, requiredColumns)
	if t == nil {
		slog.Error("No cluster master list table found.")
		return
	}

	colIndices := make(map[string]int)
	for _, colName := range requiredColumns {
		colIndices[colName], err = getColumnIndex(t, colName)
		if err != nil {
			return
		}
	}

	for i := 1; i < len(t.Elements); i++ {
		row, ok := t.Elements[i].(*asciidoc.TableRow)
		if !ok {
			continue
		}

		id := getCellTextAtCol(row, colIndices["Cluster ID"])
		name := getCellTextAtCol(row, colIndices["Cluster Name"])
		pics := getCellTextAtCol(row, colIndices["PICS Code"])

		if id == "" || name == "" {
			continue
		}
		if name == "Reserved" {
			reserveds = append(reserveds, id)
			continue
		}
		if taken, _ := clusterIdTaken(id, ci); taken {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is duplicated on Master List. ID='%s'", id)}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if taken, _ := clusterPICSTaken(pics, ci); taken {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster PICS is duplicated on Master List. PICS='%s'", pics)}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if _, ok := ci[name]; ok {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster Name is duplicated on Master List. name='%s'", name)}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}

		ci[name] = clusterInfo{
			ClusterID: id,
			PICScode:  pics,
		}
	}

	return
}
