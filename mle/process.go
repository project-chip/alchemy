package mle

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func Process(root string, s *spec.Specification) (violations map[string][]spec.Violation, err error) {
	var masterClusterListPath, masterDeviceTypeListPath asciidoc.Path
	masterClusterListPath, err = asciidoc.NewPath(filepath.Join(root, "master_lists/MasterClusterList.adoc"), root)
	if err != nil {
		return
	}
	masterDeviceTypeListPath, err = asciidoc.NewPath(filepath.Join(root, "master_lists/MasterDeviceTypeList.adoc"), root)
	if err != nil {
		return
	}

	masterClusterMap, reservedClusterIDs, vc, err := parseMasterList(masterClusterListPath, matter.TableColumnClusterID, matter.TableColumnClusterName, matter.TableColumnPICSCode)
	if err != nil {
		return
	}
	masterDeviceTypeMap, reservedDeviceTypeIDs, vd, err := parseMasterList(masterDeviceTypeListPath, matter.TableColumnDeviceID, matter.TableColumnDeviceName, matter.TableColumnUnknown)
	if err != nil {
		return
	}

	violations = spec.MergeViolations(vc, vd)

	for c := range s.Clusters {
		if !c.ID.Valid() {
			continue
		}

		if source, isReserved := reservedClusterIDs[c.ID.Value()]; isReserved {
			v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is reserved. ID='%s' Violating cluster name='%s' ", c.ID.HexString(), c.Name)}
			v.Path, v.Line = source.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		ci, ok := masterClusterMap[c.ID.Value()]
		if ok {
			if ci.Name != c.Name {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Cluster ID is present on Master List with another name. ID='%s' cluster name='%s' Name on master list='%s'", c.ID.HexString(), c.Name, ci.Name)}
				v.Path, v.Line = ci.row.Origin()
				violations[v.Path] = append(violations[v.Path], v)

			}
			if ci.PICS != c.PICS {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster PICS mismatch. Cluster='%s' PICS in data model='%s' PICS on master list='%s'", c.Name, c.PICS, ci.PICS)}
				v.Path, v.Line = ci.row.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			}
		} else {
			v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Cluster ID is missing from the Master List of clusters. ID='%s' cluster name='%s'", c.ID.HexString(), c.Name)}
			v.Path, v.Line = c.Origin()
			violations[v.Path] = append(violations[v.Path], v)
		}
	}

	for _, dt := range s.DeviceTypes {
		if !dt.ID.Valid() {
			continue
		}
		if source, isReserved := reservedDeviceTypeIDs[dt.ID.Value()]; isReserved {
			v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type ID is reserved. ID='%s' Violating device type name='%s' ", dt.ID.HexString(), dt.Name)}
			v.Path, v.Line = source.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		di, ok := masterDeviceTypeMap[dt.ID.Value()]
		if ok {
			if di.Name != dt.Name {
				v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Device Type ID is present on Master List with another name. ID='%s' device type name='%s' Name on master list='%s'", dt.ID.HexString(), dt.Name, di.Name)}
				v.Path, v.Line = di.row.Origin()
				violations[v.Path] = append(violations[v.Path], v)

			}
		} else {
			v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This device type is missing from the Master List of device types. ID='%s' device type name='%s'", dt.ID.HexString(), dt.Name)}
			v.Path, v.Line = dt.Origin()
			violations[v.Path] = append(violations[v.Path], v)
		}
	}
	return
}

type masterListInfo struct {
	ID   *matter.Number
	Name string
	PICS string
	row  *asciidoc.TableRow
}

func parseMasterList(filePath asciidoc.Path, idColumn matter.TableColumn, nameColumn matter.TableColumn, picsColumn matter.TableColumn) (masterList map[uint64]masterListInfo, reserveds map[uint64]log.Source, violations map[string][]spec.Violation, err error) {
	masterList = make(map[uint64]masterListInfo)
	uniqueNames := make(map[string]*matter.Number)
	uniquePics := make(map[string]*matter.Number)
	violations = make(map[string][]spec.Violation)
	reserveds = make(map[uint64]log.Source)

	requiredColumns := []matter.TableColumn{idColumn, nameColumn}
	if picsColumn != matter.TableColumnUnknown {
		requiredColumns = append(requiredColumns, picsColumn)
	}

	doc, err := parse.File(filePath)
	if err != nil {
		return
	}

	var listTable *spec.TableInfo
	listTable, err = findTableWithColumns(doc, requiredColumns)
	if err != nil {
		return
	}

	if listTable == nil {
		slog.Error("No master list table found.")
		return
	}

	for row := range listTable.Body() {
		var id *matter.Number
		var name, pics string
		var entity types.Entity

		id, err = listTable.ReadID(asciidoc.RawReader, row, idColumn)
		if err != nil {
			return
		}
		if !id.Valid() {
			continue
		}
		name, _, err = listTable.ReadName(asciidoc.RawReader, row, nameColumn)
		if err != nil {
			return
		}

		switch idColumn {
		case matter.TableColumnClusterID:
			entity = &matter.Cluster{
				ID:   id,
				Name: name,
			}
		case matter.TableColumnDeviceID:
			entity = &matter.DeviceType{
				ID:   id,
				Name: name,
			}
		}

		switch name {
		case "":
			continue
		case "Reserved":
			if _, taken := reserveds[id.Value()]; taken {
				v := spec.Violation{Entity: entity, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Duplicate reserved %s in Master List. ID='%s'", idColumn.String(), id.HexString())}
				v.Path, v.Line = row.Origin()
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			if _, taken := masterList[id.Value()]; taken {
				v := spec.Violation{Entity: entity, Type: spec.ViolationMasterList, Text: fmt.Sprintf("%s is both reserved and in use in Master List. ID='%s'", idColumn.String(), id.HexString())}
				v.Path, v.Line = row.Origin()
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			reserveds[id.Value()] = row
			continue
		}
		if _, reserved := reserveds[id.Value()]; reserved {
			v := spec.Violation{Entity: entity, Type: spec.ViolationMasterList, Text: fmt.Sprintf("%s is both reserved and in use in Master List. ID='%s'", idColumn.String(), id.HexString())}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if _, taken := masterList[id.Value()]; taken {
			v := spec.Violation{Entity: entity, Type: spec.ViolationMasterList, Text: fmt.Sprintf("%s is duplicated on Master List. ID='%s'", idColumn.String(), id.HexString())}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if _, taken := uniqueNames[name]; taken {
			v := spec.Violation{Entity: entity, Type: spec.ViolationMasterList, Text: fmt.Sprintf("%s is duplicated on Master List. name='%s'", nameColumn.String(), name)}
			v.Path, v.Line = row.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		if listTable.ColumnMap.HasAny(picsColumn) {
			pics, err = listTable.ReadString(asciidoc.RawReader, row, picsColumn)
			if err != nil {
				return
			}
			if pics != "" {
				if _, taken := uniquePics[pics]; taken {
					v := spec.Violation{Entity: entity, Type: spec.ViolationMasterList, Text: fmt.Sprintf("%s is duplicated on Master List. PICS='%s'", picsColumn.String(), pics)}
					v.Path, v.Line = row.Origin()
					violations[v.Path] = append(violations[v.Path], v)
					continue
				}
				uniquePics[pics] = id
			}
		}

		masterList[id.Value()] = masterListInfo{
			ID:   id,
			Name: name,
			PICS: pics,
			row:  row,
		}
		uniqueNames[name] = id
	}

	return
}

func findTableWithColumns(doc *asciidoc.Document, requiredColumns []matter.TableColumn) (matchingTable *spec.TableInfo, err error) {
	for table := range parse.FindAll[*asciidoc.Table](doc, asciidoc.RawReader, doc) {
		var ti *spec.TableInfo
		ti, err = spec.ReadTable(doc, asciidoc.RawReader, table)
		if err != nil {
			return
		}
		if ti.ColumnMap.HasAll(requiredColumns...) {
			matchingTable = ti
			break
		}
	}
	return
}
