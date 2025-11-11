package mle

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter/spec"
)

type deviceTypeInfo struct {
	DeviceTypeID string
}

func deviceTypeIdTaken(id string, masterDeviceTypeMap map[string]deviceTypeInfo) (taken bool, name string) {
	var dti deviceTypeInfo
	for name, dti = range masterDeviceTypeMap {
		if dti.DeviceTypeID == id {
			taken = true
			return
		}
	}
	taken = false
	return
}

func deviceTypeIdReserved(id string, reserveds []string) (reserved bool) {
	for _, r := range reserveds {
		if r == id {
			reserved = true
			return
		}
	}
	reserved = false
	return
}

func parseMasterDeviceTypeList(filePath string) (dti map[string]deviceTypeInfo, reserveds []string, violations map[string][]spec.Violation, err error) {
	dti = make(map[string]deviceTypeInfo)
	reserveds = make([]string, 0)
	violations = make(map[string][]spec.Violation)

	requiredColumns := []string{"Device Type ID", "Device Type Name"}

	doc, err := parse.File(filePath)
	if err != nil {
		return
	}

	t := findTableWithColumns(doc.Elements, requiredColumns)
	if t == nil {
		slog.Error("No device type master list table found.")
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

		id := getCellTextAtCol(row, colIndices["Device Type ID"])
		name := getCellTextAtCol(row, colIndices["Device Type Name"])

		if id == "" || name == "" {
			continue
		}
		if name == "Reserved" {
			reserveds = append(reserveds, id)
			continue
		}
		if taken, _ := deviceTypeIdTaken(id, dti); taken {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type ID is duplicated on Master List. ID='%s'", id)}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if _, ok := dti[name]; ok {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type Name is duplicated on Master List. name='%s'", name)}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}

		dti[name] = deviceTypeInfo{
			DeviceTypeID: id,
		}
	}

	return
}
