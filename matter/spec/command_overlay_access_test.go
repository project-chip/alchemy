package spec

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter"
)

func TestCommandOverlayTableAccess(t *testing.T) {
	for _, direction := range []struct {
		name    string
		heading string
		cell    string
		want    matter.Interface
	}{
		{"missing", "", "", matter.InterfaceUnknown},
		{"blank", " |Direction", " |", matter.InterfaceUnknown},
		{"incoming", " |Direction", " |client => server", matter.InterfaceServer},
		{"response", " |Direction", " |server => client", matter.InterfaceClient},
	} {
		t.Run(direction.name, func(t *testing.T) {
			for _, access := range []struct {
				name string
				cell string
				want matter.Privilege
			}{
				{"manage", "M", matter.PrivilegeManage},
				{"operate", "O", matter.PrivilegeOperate},
				{"blank", "", matter.PrivilegeUnknown},
			} {
				t.Run(access.name, func(t *testing.T) {
					input := "|===\n|ID |Name |Access |Conformance" + direction.heading +
						"\n|9 |SyntheticCommand |" + access.cell + " |M" + direction.cell + "\n|===\n"
					document, err := parse.String(asciidoc.Path{}, input)
					if err != nil {
						t.Fatal(err)
					}
					var table *asciidoc.Table
					parse.SkimFunc(asciidoc.RawReader, document, document.Children(), func(found *asciidoc.Table) bool {
						table = found
						return true
					})
					if table == nil {
						t.Fatal("synthetic command table missing")
					}
					info, err := ReadTable(document, asciidoc.RawReader, table)
					if err != nil {
						t.Fatal(err)
					}
					if len(info.Rows) != 2 {
						t.Fatalf("rows = %d, want 2", len(info.Rows))
					}
					library := &Library{Reader: asciidoc.RawReader}
					command, err := (&commandFactory{}).New(nil, library, library, document, nil, info, info.Rows[1], "SyntheticCommand", nil)
					if err != nil {
						t.Fatal(err)
					}
					if command.Direction != direction.want {
						t.Fatalf("direction = %v, want %v", command.Direction, direction.want)
					}
					if command.Access.Invoke != access.want {
						t.Fatalf("invoke privilege = %v, want %v", command.Access.Invoke, access.want)
					}
				})
			}
		})
	}
}
