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

type deviceTypeInfo struct {
	ID   *matter.Number
	Name string
}

func parseMasterDeviceTypeList(filePath asciidoc.Path) (dti map[uint64]deviceTypeInfo, reserveds map[uint64]log.Source, violations map[string][]spec.Violation, err error) {
	dti = make(map[uint64]deviceTypeInfo)
	reserveds = make(map[uint64]log.Source)
	uniqueNames := make(map[string]*matter.Number)
	violations = make(map[string][]spec.Violation)

	requiredColumns := []matter.TableColumn{matter.TableColumnDeviceID, matter.TableColumnDeviceName}

	doc, err := parse.File(filePath)
	if err != nil {
		return
	}

	var deviceTypeListTable *spec.TableInfo
	deviceTypeListTable, err = findTableWithColumns(doc, requiredColumns)
	if err != nil {
		return
	}
	if deviceTypeListTable == nil {
		slog.Error("No device type master list table found.")
		return
	}

	for row := range deviceTypeListTable.Body() {
		var id *matter.Number
		var name string
		id, err = deviceTypeListTable.ReadID(asciidoc.RawReader, row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		if !id.Valid() {
			continue
		}
		name, _, err = deviceTypeListTable.ReadName(asciidoc.RawReader, row, matter.TableColumnDeviceName)
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
			if _, taken := dti[id.Value()]; taken {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is both reserved and in use in Master List. ID='%s'", id.HexString())}
				v.Path, v.Line = row.Origin()
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			reserveds[id.Value()] = row
			continue
		}
		if err != nil {
			return
		}
		if _, reserved := reserveds[id.Value()]; reserved {
			v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is both reserved and in use in Master List. ID='%s'", id.HexString())}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if _, taken := dti[id.Value()]; taken {
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
		dti[id.Value()] = deviceTypeInfo{
			ID:   id,
			Name: name,
		}
		uniqueNames[name] = id
	}

	return
}
