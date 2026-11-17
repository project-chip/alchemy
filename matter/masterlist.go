package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type MasterList struct {
	source asciidoc.Element

	Type    types.EntityType
	Entries []*MasterListEntry
}

func (ml *MasterList) EntityType() types.EntityType {
	return types.EntityTypeMasterList
}

func (ml *MasterList) Source() asciidoc.Element {
	return ml.source
}

func (ml *MasterList) Origin() (path string, line int) {
	switch s := ml.source.(type) {
	case Source:
		return s.Origin()
	default:
		return "", -1
	}
}

func (ml *MasterList) Parent() types.Entity {
	return nil
}

func (ml *MasterList) Equals(o types.Entity) bool {
	oml, ok := o.(*MasterList)
	if !ok {
		return false
	}
	return oml.Type == ml.Type
}

func NewMasterList(entityType types.EntityType, source asciidoc.Element) *MasterList {
	return &MasterList{
		Type:   entityType,
		source: source,
	}
}

type MasterListEntry struct {
	parent *MasterList
	source asciidoc.Element

	ID                 *Number
	Name               string
	Scope              Scope
	PICS               string
	Base               string
	CertificationState CertificationState
	Classification     Role
	Category           string
	Notes              string
}

func (ml *MasterList) NewMasterListEntry(source asciidoc.Element) *MasterListEntry {
	return &MasterListEntry{
		parent: ml,
		source: source,
	}
}

func (mle *MasterListEntry) EntityType() types.EntityType {
	return types.EntityTypeMasterListEntry
}

func (mle *MasterListEntry) Source() asciidoc.Element {
	return mle.source
}

func (mle *MasterListEntry) Origin() (path string, line int) {
	switch s := mle.source.(type) {
	case Source:
		return s.Origin()
	default:
		return "", -1
	}
}

func (mle *MasterListEntry) Parent() types.Entity {
	return mle.parent
}

func (mle *MasterListEntry) Equals(o types.Entity) bool {
	omle, ok := o.(*MasterListEntry)
	if !ok {
		return false
	}
	if !omle.parent.Equals(omle.parent) {
		return false
	}
	return omle.ID.Equals(mle.ID)
}

type CertificationState uint8

const (
	CertificationStateUnknown CertificationState = iota
	CertificationStateAllocated
	CertificationStateProvisional
	CertificationStateCertifiable
	CertificationStateDeprecated
	CertificationStateNotCertifiable
	CertificationStateDoNotUse
)

var CertificationStateNames = map[CertificationState]string{
	CertificationStateUnknown:        "Unknown",
	CertificationStateAllocated:      "Allocated",
	CertificationStateProvisional:    "Provisional",
	CertificationStateCertifiable:    "Certifiable",
	CertificationStateDeprecated:     "Depreacted",
	CertificationStateNotCertifiable: "NotCertifiable",
	CertificationStateDoNotUse:       "DoNotUse",
}

func (cs CertificationState) String() string {
	return CertificationStateNames[cs]
}
