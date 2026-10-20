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

func TestCommandOverlayInheritedAccessDirection(t *testing.T) {
	for _, test := range []struct {
		name       string
		directions []matter.Interface
		baseID     uint64
		want       string
	}{
		{"response", []matter.Interface{matter.InterfaceClient}, 9, ""},
		{"request", []matter.Interface{matter.InterfaceServer}, 9, "manage"},
		{"ancestor_response", []matter.Interface{matter.InterfaceUnknown, matter.InterfaceClient}, 9, ""},
		{"ancestor_request", []matter.Interface{matter.InterfaceUnknown, matter.InterfaceServer}, 9, "manage"},
		{"unresolved", []matter.Interface{matter.InterfaceUnknown}, 9, "manage"},
		{"no_base", nil, 9, "manage"},
		{"different_id", []matter.Interface{matter.InterfaceClient}, 10, "manage"},
	} {
		t.Run(test.name, func(t *testing.T) {
			derived := matter.NewCluster(nil)
			cluster := derived
			for _, direction := range test.directions {
				base := matter.NewCluster(nil)
				command := matter.NewCommand(nil, base)
				command.ID = matter.NewNumber(test.baseID)
				command.Name = "SyntheticCommand"
				command.Direction = direction
				sibling := matter.NewCommand(nil, base)
				sibling.ID = matter.NewNumber(test.baseID)
				sibling.Name = "DifferentCommand"
				sibling.Direction = matter.InterfaceServer
				base.Commands = matter.CommandSet{sibling, command}
				cluster.ParentCluster = base
				cluster = base
			}
			overlay := matter.NewCommand(nil, derived)
			overlay.ID = matter.NewNumber(9)
			overlay.Name = "SyntheticCommand"
			overlay.Access.Invoke = matter.PrivilegeManage
			overlay.Conformance = conformance.Set{}
			document := etree.NewDocument()
			if err := renderCommands(matter.CommandSet{overlay}, document.CreateElement("cluster")); err != nil {
				t.Fatal(err)
			}
			rendered := document.FindElement("cluster/commands/command")
			if rendered == nil {
				t.Fatal("command missing")
			}
			if overlay.Direction != matter.InterfaceUnknown || rendered.SelectAttr("direction") != nil {
				t.Fatal("direction lookup must not mutate the overlay or emit a direction")
			}
			access := rendered.FindElements("access")
			if test.want == "" {
				if len(access) != 0 {
					t.Fatal("inherited response acquired invoke access")
				}
			} else if len(access) != 1 || access[0].SelectAttrValue("invokePrivilege", "") != test.want {
				t.Fatalf("expected exactly one access element with invokePrivilege=%q", test.want)
			}
		})
	}
}
