package zap

import (
	"encoding/xml"

	"github.com/hasty/matterfmt/matter"
)

type XMLAccess struct {
	XMLName   xml.Name `xml:"access"`
	OP        string   `xml:"op,attr"`
	Privilege string   `xml:"privilege,attr"`
	Role      string   `xml:"role,attr"`
	Modifier  string   `xml:"modifier,attr"`
}

func ToAccessModel(xas []XMLAccess) *matter.Access {
	a := &matter.Access{}
	for _, xa := range xas {
		p := parsePrivilege(xa.Privilege)
		if p == matter.PrivilegeUnknown {
			p = parsePrivilege(xa.Role)
		}
		if p == matter.PrivilegeUnknown {
			continue
		}
		switch xa.OP {
		case "invoke":
			a.Invoke = p
		case "read":
			a.Read = p
		case "write":
			a.Write = p
		}
	}
	return a
}

func parsePrivilege(a string) matter.Privilege {
	switch a {
	case "view":
		return matter.PrivilegeView
	case "manage":
		return matter.PrivilegeManage
	case "administer":
		return matter.PrivilegeAdminister
	case "operate":
		return matter.PrivilegeOperate
	default:
		return matter.PrivilegeUnknown
	}
}
