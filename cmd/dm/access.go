package dm

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderAccess(ax *etree.Element, a *matter.Field) {
	acx := ax.CreateElement("access")
	if a.Access.Read != matter.PrivilegeUnknown {
		if a.Access.Read == matter.PrivilegeView {
			acx.CreateAttr("read", "true")
		} else {
			acx.CreateAttr("read", "optional")
		}
	}
	if a.Access.Write != matter.PrivilegeUnknown {
		if a.Access.Write == matter.PrivilegeOperate {
			acx.CreateAttr("write", "true")
		} else {
			acx.CreateAttr("write", "optional")
		}
	}
	if a.Access.Read != matter.PrivilegeUnknown {
		acx.CreateAttr("readPrivilege", strings.ToLower(matter.PrivilegeNamesShort[a.Access.Read]))
	}
	if a.Access.Write != matter.PrivilegeUnknown {
		acx.CreateAttr("writePrivilege", strings.ToLower(matter.PrivilegeNamesShort[a.Access.Write]))
	}
}
