package zap

import (
	"encoding/xml"

	"github.com/hasty/matterfmt/matter"
)

type XMLAccess struct {
	XMLName   xml.Name         `xml:"access"`
	OP        string           `xml:"op,attr"`
	Privilege matter.Privilege `xml:"privilege,attr"`
	Role      string           `xml:"role,attr"`
	Modifier  string           `xml:"modifier,attr"`
}

func ToAccessModel(xas []XMLAccess) matter.Access {
	a := matter.Access{}
	for _, xa := range xas {
		if xa.Privilege == matter.PrivilegeUnknown {
			continue
		}
		switch xa.OP {
		case "invoke":
			a.Invoke = xa.Privilege
		case "read":
			a.Read = xa.Privilege
		case "write":
			a.Write = xa.Privilege
		}
	}
	return a
}
