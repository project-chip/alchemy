package matter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Privilege uint8

const (
	PrivilegeUnknown Privilege = iota
	PrivilegeView
	PrivilegeOperate
	PrivilegeManage
	PrivilegeAdminister
)

var PrivilegeNames = map[Privilege]string{
	PrivilegeUnknown:    "Unknown",
	PrivilegeView:       "View",
	PrivilegeOperate:    "Operate",
	PrivilegeManage:     "Manage",
	PrivilegeAdminister: "Administer",
}

var PrivilegeNamesShort = map[Privilege]string{
	PrivilegeUnknown:    "Unknown",
	PrivilegeView:       "View",
	PrivilegeOperate:    "Operate",
	PrivilegeManage:     "Manage",
	PrivilegeAdminister: "Admin",
}

var privilegeNameMap map[string]Privilege

type FabricScoping uint8

const (
	FabricScopingUnknown FabricScoping = iota
	FabricScopingScoped
	FabricScopingUnscoped
)

func (fs FabricScoping) String() string {
	switch fs {
	case FabricScopingScoped:
		return "scoped"
	case FabricScopingUnscoped:
		return "unscoped"
	default:
		return "unknown"
	}
}

type FabricSensitivity uint8

const (
	FabricSensitivityUnknown FabricSensitivity = iota
	FabricSensitivitySensitive
	FabricSensitivityInsensitive
)

func (fs FabricSensitivity) String() string {
	switch fs {
	case FabricSensitivitySensitive:
		return "sensitive"
	case FabricSensitivityInsensitive:
		return "insensitive"
	default:
		return "unknown"
	}
}

type Timing uint8

const (
	TimingUnknown Timing = iota
	TimingTimed
	TimingUntimed
)

func init() {
	privilegeNameMap = make(map[string]Privilege, len(PrivilegeNames))
	for p, n := range PrivilegeNames {
		privilegeNameMap[n] = p
	}
}

func (p Privilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(PrivilegeNames[p])
}

func (p *Privilege) UnmarshalJSON(data []byte) error {
	var privilege string
	if err := json.Unmarshal(data, &privilege); err != nil {
		return fmt.Errorf("error parsing privilege %s: %w", string(data), err)
	}
	var ok bool
	*p, ok = privilegeNameMap[privilege]
	if !ok {
		return fmt.Errorf("unknown privilege: %s", privilege)
	}
	return nil
}

func (p Privilege) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	switch p {
	case PrivilegeView:
		return xml.Attr{Name: name, Value: "view"}, nil
	case PrivilegeManage:
		return xml.Attr{Name: name, Value: "manage"}, nil
	case PrivilegeAdminister:
		return xml.Attr{Name: name, Value: "administer"}, nil
	case PrivilegeOperate:
		return xml.Attr{Name: name, Value: "operate"}, nil
	default:
		return xml.Attr{}, fmt.Errorf("unknown privilege value: %v", p)
	}
}

func (p *Privilege) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "view":
		*p = PrivilegeView
	case "manage":
		*p = PrivilegeManage
	case "administer":
		*p = PrivilegeAdminister
	case "operate":
		*p = PrivilegeOperate
	default:
		return fmt.Errorf("unknown privilege value: %s", attr.Value)
	}
	return nil
}

type Access struct {
	Read   Privilege `json:"read,omitempty"`
	Write  Privilege `json:"write,omitempty"`
	Invoke Privilege `json:"invoke,omitempty"`

	OptionalWrite     bool              `json:"optionalWrite,omitempty"`
	FabricScoping     FabricScoping     `json:"fabricScoped,omitempty"`
	FabricSensitivity FabricSensitivity `json:"fabricSensitive,omitempty"`

	Timing Timing `json:"timed,omitempty"`
}

func (a Access) IsFabricScoped() bool {
	return a.FabricScoping == FabricScopingScoped
}

func (a Access) IsFabricSensitive() bool {
	return a.FabricSensitivity == FabricSensitivitySensitive
}

func (a Access) IsTimed() bool {
	return a.Timing == TimingTimed
}

func (a Access) Equal(oa Access) bool {
	if a.Read != oa.Read {
		return false
	}
	if a.Write != oa.Write {
		return false
	}
	if a.Invoke != oa.Invoke {
		return false
	}
	if a.OptionalWrite != oa.OptionalWrite {
		return false
	}
	if a.FabricScoping != oa.FabricScoping {
		return false
	}
	if a.FabricSensitivity != oa.FabricSensitivity {
		return false
	}
	if a.Timing != oa.Timing {
		return false
	}
	return true
}

func (a *Access) Inherit(parent Access) {
	if a.Read == PrivilegeUnknown && parent.Read != PrivilegeUnknown {
		a.Read = parent.Read
	}
	if a.Write == PrivilegeUnknown && parent.Write != PrivilegeUnknown {
		a.Write = parent.Write
		a.OptionalWrite = parent.OptionalWrite
	}
	if a.Invoke == PrivilegeUnknown && parent.Invoke != PrivilegeUnknown {
		a.Invoke = parent.Invoke
	}
	if a.FabricScoping == FabricScopingUnknown && parent.FabricScoping != FabricScopingUnknown {
		a.FabricScoping = parent.FabricScoping
	}
	if a.Timing == TimingUnknown && parent.Timing != TimingUnknown {
		a.Timing = parent.Timing
	}
}

func DefaultAccess(forInvoke bool) Access {
	a := Access{FabricSensitivity: FabricSensitivityInsensitive, FabricScoping: FabricScopingUnscoped, Timing: TimingUntimed}
	if forInvoke {
		a.Invoke = PrivilegeOperate
	} else {
		a.Read = PrivilegeView
	}
	return a
}
