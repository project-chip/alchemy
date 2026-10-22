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

func TestCommandOverlayResponseAccessFlags(t *testing.T) {
	for _, direction := range []struct {
		name     string
		overlay  matter.Interface
		base     matter.Interface
		response bool
	}{
		{"direct_response", matter.InterfaceClient, matter.InterfaceClient, true},
		{"inherited_response", matter.InterfaceUnknown, matter.InterfaceClient, true},
		{"inherited_request", matter.InterfaceUnknown, matter.InterfaceServer, false},
	} {
		t.Run(direction.name, func(t *testing.T) {
			for _, flags := range []struct {
				name         string
				fabricScoped bool
				timed        bool
			}{
				{"none", false, false},
				{"fabric_scoped", true, false},
				{"timed", false, true},
				{"fabric_scoped_timed", true, true},
			} {
				t.Run(flags.name, func(t *testing.T) {
					base := matter.NewCluster(nil)
					baseCommand := matter.NewCommand(nil, base)
					baseCommand.ID = matter.NewNumber(9)
					baseCommand.Name = "SyntheticCommand"
					baseCommand.Direction = direction.base
					base.Commands = matter.CommandSet{baseCommand}
					derived := matter.NewCluster(nil)
					derived.ParentCluster = base
					command := matter.NewCommand(nil, derived)
					command.ID = matter.NewNumber(9)
					command.Name = "SyntheticCommand"
					command.Direction = direction.overlay
					command.Access.Invoke = matter.PrivilegeManage
					if flags.fabricScoped {
						command.Access.FabricScoping = matter.FabricScopingScoped
					}
					if flags.timed {
						command.Access.Timing = matter.TimingTimed
					}
					command.Conformance = conformance.Set{}
					document := etree.NewDocument()
					if err := renderCommands(matter.CommandSet{command}, document.CreateElement("cluster")); err != nil {
						t.Fatal(err)
					}
					rendered := document.FindElement("cluster/commands/command")
					if rendered == nil {
						t.Fatal("command missing")
					}
					if command.Direction != direction.overlay {
						t.Fatal("direction lookup mutated the command")
					}
					if direction.overlay == matter.InterfaceUnknown && rendered.SelectAttr("direction") != nil {
						t.Fatal("inherited direction must remain absent")
					}
					access := rendered.FindElements("access")
					if direction.response && !flags.fabricScoped {
						if len(access) != 0 {
							t.Fatal("non-fabric-scoped response acquired access")
						}
						return
					}
					if len(access) != 1 {
						t.Fatalf("access elements = %d, want 1", len(access))
					}
					wantInvoke := "manage"
					if direction.response {
						wantInvoke = ""
					}
					if got := access[0].SelectAttrValue("invokePrivilege", ""); got != wantInvoke {
						t.Fatalf("invokePrivilege = %q, want %q", got, wantInvoke)
					}
					if got := access[0].SelectAttrValue("fabricScoped", "") == "true"; got != flags.fabricScoped {
						t.Fatalf("fabricScoped = %t, want %t", got, flags.fabricScoped)
					}
					if got := access[0].SelectAttrValue("timed", "") == "true"; got != flags.timed {
						t.Fatalf("timed = %t, want %t", got, flags.timed)
					}
				})
			}
		})
	}
}
