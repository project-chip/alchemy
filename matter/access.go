package matter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/hasty/alchemy/matter/types"
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

func init() {
	privilegeNameMap = make(map[string]Privilege, len(PrivilegeNames))
	for p, n := range PrivilegeNames {
		privilegeNameMap[n] = p
	}
}

func (p Privilege) String() string {
	return PrivilegeNames[p]
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

func (fs FabricScoping) MarshalJSON() ([]byte, error) {
	return json.Marshal(fs.String())
}

func (fs *FabricScoping) UnmarshalJSON(data []byte) error {
	var scoping string
	if err := json.Unmarshal(data, &scoping); err != nil {
		return fmt.Errorf("error parsing fabric scoping %s: %w", string(data), err)
	}
	switch scoping {
	case "scoped":
		*fs = FabricScopingScoped
	case "unscoped":
		*fs = FabricScopingUnscoped
	case "unknown":
		*fs = FabricScopingUnknown
	default:
		return fmt.Errorf("unknown fabric scoping: %s", scoping)
	}
	return nil
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

func (fs FabricSensitivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(fs.String())
}

func (fs *FabricSensitivity) UnmarshalJSON(data []byte) error {
	var sensitivity string
	if err := json.Unmarshal(data, &sensitivity); err != nil {
		return fmt.Errorf("error parsing fabric sensitivity %s: %w", string(data), err)
	}
	switch sensitivity {
	case "sensitive":
		*fs = FabricSensitivitySensitive
	case "insensitive":
		*fs = FabricSensitivityInsensitive
	case "unknown":
		*fs = FabricSensitivityUnknown
	default:
		return fmt.Errorf("unknown fabric sensitivity: %s", sensitivity)
	}
	return nil
}

type Timing uint8

const (
	TimingUnknown Timing = iota
	TimingTimed
	TimingUntimed
)

func (t Timing) String() string {
	switch t {
	case TimingTimed:
		return "timed"
	case TimingUntimed:
		return "untimed"
	default:
		return "unknown"
	}
}

func (fs Timing) MarshalJSON() ([]byte, error) {
	return json.Marshal(fs.String())
}

func (t *Timing) UnmarshalJSON(data []byte) error {
	var timing string
	if err := json.Unmarshal(data, &timing); err != nil {
		return fmt.Errorf("error parsing timing %s: %w", string(data), err)
	}
	switch timing {
	case "timed":
		*t = TimingTimed
	case "untimed":
		*t = TimingUntimed
	case "unknown":
		*t = TimingUnknown
	default:
		return fmt.Errorf("unknown timing: %s", timing)
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

func DefaultAccess(entityType types.EntityType) Access {
	a := Access{FabricSensitivity: FabricSensitivityInsensitive, FabricScoping: FabricScopingUnscoped, Timing: TimingUntimed}
	switch entityType {
	case types.EntityTypeCommand:
		a.Invoke = PrivilegeOperate
	case types.EntityTypeStruct: // Structs don't get R/W access
	default:
		a.Read = PrivilegeView
	}
	return a
}
