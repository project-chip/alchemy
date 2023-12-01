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

	OptionalWrite   bool `json:"optionalWrite,omitempty"`
	FabricScoped    bool `json:"fabricScoped,omitempty"`
	FabricSensitive bool `json:"fabricSensitive,omitempty"`

	Timed bool `json:"timed,omitempty"`
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
	if a.FabricScoped != oa.FabricScoped {
		return false
	}
	if a.FabricSensitive != oa.FabricSensitive {
		return false
	}
	if a.Timed != oa.Timed {
		return false
	}
	return true
}
