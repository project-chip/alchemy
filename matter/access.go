package matter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log/slog"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter/types"
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

func (fs FabricSensitivity) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(fs.String())
}

func (fs *FabricSensitivity) UnmarshalYAML(b []byte) error {
	var sensitivity string
	if err := yaml.Unmarshal(b, &sensitivity); err != nil {
		return fmt.Errorf("error parsing fabric sensitivity %s: %w", string(b), err)
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

func (t Timing) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
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
	Read   Privilege `json:"read,omitempty" yaml:"read,omitempty"`
	Write  Privilege `json:"write,omitempty"  yaml:"write,omitempty"`
	Invoke Privilege `json:"invoke,omitempty"  yaml:"invoke,omitempty"`

	OptionalWrite     bool              `json:"optionalWrite,omitempty"  yaml:"optionalWrite,omitempty"`
	FabricScoping     FabricScoping     `json:"fabricScoped,omitempty"  yaml:"fabricScoping,omitempty"`
	FabricSensitivity FabricSensitivity `json:"fabricSensitive,omitempty"  yaml:"fabricSensitivity,omitempty"`

	Timing Timing `json:"timed,omitempty"`
}

func (a Access) IsEmpty() bool {
	if a.Read != PrivilegeUnknown || a.Write != PrivilegeUnknown || a.Invoke != PrivilegeUnknown {
		return false
	}
	if a.OptionalWrite || a.FabricScoping != FabricScopingUnknown || a.FabricSensitivity != FabricSensitivityUnknown {
		return false
	}
	return a.Timing == TimingUnknown
}

func IsFabricScopingAllowed(e types.Entity) bool {
	switch e := e.(type) {
	case *Field:
		switch e.EntityType() {
		case types.EntityTypeAttribute:
			return true
		default:
			return false
		}
	case *Struct, *Command:
		return true
	default:
		return false
	}
}

func IsFabricSensitivityAllowed(e types.Entity) bool {
	switch e := e.(type) {
	case *Field:
		switch e.EntityType() {
		case types.EntityTypeAttribute:
			if e.Type == nil {
				return false
			}
			if !e.Type.IsArray() {
				return false
			}
			if e.Type.EntryType == nil || e.Type.EntryType.Entity == nil {
				return false
			}
			switch te := e.Type.EntryType.Entity.(type) {
			case *Struct:
				return te.FabricScoping == FabricScopingScoped
			default:
				return false
			}
		case types.EntityTypeStructField:
			parent := e.Parent()
			switch parent := parent.(type) {
			case *Struct:
				return parent.FabricScoping == FabricScopingScoped
			default:
				slog.Warn("Unexpected struct field parent type", log.Type("parentType", parent), log.Path("source", e))
				return false
			}
		default:
			return false
		}
	case *Event:
		return true
	default:
		return false
	}
}

func (a Access) String() string {
	var sb strings.Builder
	if a.Read != PrivilegeUnknown {
		sb.WriteString("read: ")
		sb.WriteString(a.Read.String())
	}
	if a.Write != PrivilegeUnknown {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("write: ")
		sb.WriteString(a.Write.String())
		if a.OptionalWrite {
			sb.WriteString(" (optional)")
		}
	}
	if a.Invoke != PrivilegeUnknown {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("invoke: ")
		sb.WriteString(a.Invoke.String())
	}
	if a.FabricScoping != FabricScopingUnknown {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(a.FabricScoping.String())
	}
	if a.FabricSensitivity != FabricSensitivityUnknown {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(a.FabricSensitivity.String())
	}
	if a.Timing != TimingUnknown {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(a.Timing.String())
	}
	return sb.String()
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
	if a.Read == PrivilegeUnknown {
		a.Read = parent.Read
	}
	if a.Write == PrivilegeUnknown {
		a.Write = parent.Write
		a.OptionalWrite = parent.OptionalWrite
	}
	if a.Invoke == PrivilegeUnknown {
		a.Invoke = parent.Invoke
	}
	if a.FabricScoping == FabricScopingUnknown {
		a.FabricScoping = parent.FabricScoping
	}
	if a.Timing == TimingUnknown {
		a.Timing = parent.Timing
	}
}

func DefaultAccess(entityType types.EntityType) Access {
	return Access{
		Read:              DefaultReadPrivilege(entityType),
		Invoke:            DefaultInvokePrivilege(entityType),
		FabricSensitivity: FabricSensitivityInsensitive,
		FabricScoping:     FabricScopingUnscoped,
		Timing:            TimingUntimed,
	}
}

func DefaultReadPrivilege(entityType types.EntityType) Privilege {
	switch entityType {
	case types.EntityTypeAttribute, types.EntityTypeEventField: // Structs don't get R/W access
		return PrivilegeView
	default:
		return PrivilegeUnknown
	}
}

func DefaultWritePrivilege(entityType types.EntityType) Privilege {
	switch entityType {
	case types.EntityTypeAttribute, types.EntityTypeEventField: // Structs don't get R/W access
		return PrivilegeOperate
	default:
		return PrivilegeUnknown
	}
}

func DefaultInvokePrivilege(entityType types.EntityType) Privilege {
	switch entityType {
	case types.EntityTypeCommand:
		return PrivilegeOperate
	default:
		return PrivilegeUnknown
	}
}
