package dm

import (
	"testing"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
)

func TestCommandOverlayInvokePrivilege(t *testing.T) {
	for _, direction := range []struct {
		name  string
		value matter.Interface
		want  string
	}{
		{"inherited", matter.InterfaceUnknown, ""},
		{"incoming", matter.InterfaceServer, "commandToServer"},
		{"response", matter.InterfaceClient, "responseFromServer"},
	} {
		t.Run(direction.name, func(t *testing.T) {
			for _, privilege := range []struct {
				name  string
				value matter.Privilege
				want  string
			}{
				{"manage", matter.PrivilegeManage, "manage"},
				{"operate", matter.PrivilegeOperate, "operate"},
				{"blank", matter.PrivilegeUnknown, ""},
			} {
				t.Run(privilege.name, func(t *testing.T) {
					command := matter.NewCommand(nil, nil)
					command.ID = matter.NewNumber(9)
					command.Name = "SyntheticCommand"
					command.Direction = direction.value
					command.Access.Invoke = privilege.value
					command.Conformance = conformance.Set{}
					document := etree.NewDocument()
					if err := renderCommands(matter.CommandSet{command}, document.CreateElement("cluster")); err != nil {
						t.Fatal(err)
					}
					serialized, err := document.WriteToString()
					if err != nil {
						t.Fatal(err)
					}
					reparsed := etree.NewDocument()
					if err := reparsed.ReadFromString(serialized); err != nil {
						t.Fatal(err)
					}
					rendered := reparsed.FindElement("cluster/commands/command")
					if rendered == nil {
						t.Fatal("command missing after XML round trip")
					}
					if got := rendered.SelectAttrValue("direction", ""); got != direction.want {
						t.Fatalf("direction = %q, want %q", got, direction.want)
					}
					if direction.value == matter.InterfaceUnknown && rendered.SelectAttr("direction") != nil {
						t.Fatal("unknown direction must remain absent")
					}
					want := privilege.want
					if direction.value == matter.InterfaceClient {
						want = ""
					}
					access := rendered.FindElements("access")
					if want == "" {
						if len(access) != 0 {
							t.Fatal("blank access or response acquired access")
						}
					} else if len(access) != 1 || access[0].SelectAttrValue("invokePrivilege", "") != want {
						t.Fatalf("expected exactly one access element with invokePrivilege=%q", want)
					}
				})
			}
		})
	}
}
