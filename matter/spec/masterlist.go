package spec

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func ParseCertificationState(s string) matter.CertificationState {
	switch s {
	case "A":
		return matter.CertificationStateAllocated
	case "P":
		return matter.CertificationStateProvisional
	case "C":
		return matter.CertificationStateCertifiable
	case "D":
		return matter.CertificationStateDeprecated
	case "N":
		return matter.CertificationStateNotCertifiable
	case "X":
		return matter.CertificationStateDoNotUse
	default:
		return matter.CertificationStateUnknown
	}
}

func (library *Library) toMasterList(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section) (ns *matter.MasterList, err error) {
	var deviceList, clusterList *TableInfo
	var namespaceLists []*TableInfo
	for t := range parse.FindAll[*asciidoc.Table](d, reader, s) {
		ti, err := ReadTable(d, reader, t)
		if err != nil {
			continue
		}
		if ti.ColumnMap.HasAll(matter.TableColumnDeviceID, matter.TableColumnDeviceName) {
			deviceList = ti
			break
		}
		if ti.ColumnMap.HasAll(matter.TableColumnClusterID, matter.TableColumnClusterName) {
			clusterList = ti
			break
		}
		if ti.ColumnMap.HasAll(matter.TableColumnNamespaceID, matter.TableColumnNamespace) {
			namespaceLists = append(namespaceLists, ti)
		}
	}
	if clusterList != nil {
		ns, err = library.toClusterMasterList(reader, s, clusterList)
		if err != nil {
			return
		}
	}
	if deviceList != nil {
		ns, err = library.toDeviceMasterList(reader, s, deviceList)
		if err != nil {
			return
		}
	}
	if len(namespaceLists) > 0 {
		ns, err = library.toNamepsaceMasterList(reader, s, namespaceLists)
		if err != nil {
			return
		}
	}
	return
}

func (library *Library) toClusterMasterList(reader asciidoc.Reader, s *asciidoc.Section, ti *TableInfo) (ml *matter.MasterList, err error) {
	ml = matter.NewMasterList(types.EntityTypeCluster, s)
	for row := range ti.ContentRows() {
		mle := ml.NewMasterListEntry(row)
		mle.ID, err = ti.ReadID(reader, row, matter.TableColumnClusterID)
		if err != nil {
			return
		}
		mle.Name, _, err = ti.ReadName(reader, row, matter.TableColumnClusterName)
		if err != nil {
			return
		}
		mle.PICS, err = ti.ReadValue(reader, row, matter.TableColumnPICSCode)
		if err != nil {
			return
		}
		mle.Base, err = ti.ReadValue(reader, row, matter.TableColumnBaseCluster)
		if err != nil {
			return
		}
		mle.Classification, err = ti.ReadRole(reader, row, matter.TableColumnClassification)
		if err != nil {
			return
		}
		mle.Scope, err = ti.ReadScope(reader, row, matter.TableColumnScope)
		if err != nil {
			return
		}
		mle.Category, err = ti.ReadValue(reader, row, matter.TableColumnCategory)
		if err != nil {
			return
		}
		mle.CertificationState, err = ti.ReadCertificationState(reader, row, matter.TableColumnCertificationState)
		if err != nil {
			return
		}
		mle.Notes, err = ti.ReadValue(reader, row, matter.TableColumnNotes)
		if err != nil {
			return
		}
		ml.Entries = append(ml.Entries, mle)
	}
	return
}

func (library *Library) toDeviceMasterList(reader asciidoc.Reader, s *asciidoc.Section, ti *TableInfo) (ml *matter.MasterList, err error) {
	ml = matter.NewMasterList(types.EntityTypeDeviceType, s)
	for row := range ti.ContentRows() {
		mle := ml.NewMasterListEntry(row)
		mle.ID, err = ti.ReadID(reader, row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		mle.Name, _, err = ti.ReadName(reader, row, matter.TableColumnDeviceName)
		if err != nil {
			return
		}
		mle.Scope, err = ti.ReadScope(reader, row, matter.TableColumnScope)
		if err != nil {
			return
		}
		mle.Category, err = ti.ReadValue(reader, row, matter.TableColumnCategory)
		if err != nil {
			return
		}
		mle.CertificationState, err = ti.ReadCertificationState(reader, row, matter.TableColumnCertificationState)
		if err != nil {
			return
		}
		mle.Notes, err = ti.ReadValue(reader, row, matter.TableColumnNotes)
		if err != nil {
			return
		}
		ml.Entries = append(ml.Entries, mle)
	}
	return
}

func (library *Library) toNamepsaceMasterList(reader asciidoc.Reader, s *asciidoc.Section, tis []*TableInfo) (ml *matter.MasterList, err error) {
	ml = matter.NewMasterList(types.EntityTypeNamespace, s)
	for _, ti := range tis {
		for row := range ti.ContentRows() {
			mle := ml.NewMasterListEntry(row)
			mle.ID, err = ti.ReadID(reader, row, matter.TableColumnNamespaceID)
			if err != nil {
				return
			}
			mle.Name, _, err = ti.ReadName(reader, row, matter.TableColumnNamespace)
			if err != nil {
				return
			}
			mle.Notes, err = ti.ReadValue(reader, row, matter.TableColumnNotes)
			if err != nil {
				return
			}
			ml.Entries = append(ml.Entries, mle)
		}
	}
	return
}
