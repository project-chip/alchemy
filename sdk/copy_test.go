package sdk

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type copyTestSpec struct {
	spec *spec.Specification
	p    *patch

	owner     *matter.Cluster
	user      *matter.Cluster
	userDoc   *asciidoc.Document
	shared    *matter.Struct
	nested    *matter.Enum
	attribute *matter.Field
	command   *matter.Command
}

// newCopyTestSpec mirrors Channel / Content Launcher: the owner cluster defines a struct (with a field
// using an enum it also owns), and the user cluster references that struct from an attribute (as a list)
// and from a command field
func newCopyTestSpec() *copyTestSpec {
	ts := &copyTestSpec{userDoc: &asciidoc.Document{}}
	ts.spec = &spec.Specification{
		DocRefs:     make(map[types.Entity]*asciidoc.Document),
		LibraryRefs: make(map[types.Entity]*spec.Library),
	}
	ts.p = &patch{spec: ts.spec}

	ts.owner = matter.NewCluster(nil)
	ts.owner.Name = "Content Launcher"
	ts.nested = matter.NewEnum(nil, ts.owner)
	ts.nested.Name = "MetricTypeEnum"
	ts.owner.Enums = append(ts.owner.Enums, ts.nested)
	ts.shared = matter.NewStruct(nil, ts.owner)
	ts.shared.Name = "AdditionalInfoStruct"
	name := matter.NewField(nil, ts.shared, types.EntityTypeStructField)
	name.ID = matter.NewNumber(0)
	name.Name = "Name"
	name.Type = types.NewDataType(types.BaseDataTypeString, types.DataTypeRankScalar)
	metric := matter.NewField(nil, ts.shared, types.EntityTypeStructField)
	metric.ID = matter.NewNumber(1)
	metric.Name = "Metric"
	metric.Type = &types.DataType{Name: "MetricTypeEnum", BaseType: types.BaseDataTypeCustom, Entity: ts.nested}
	ts.shared.Fields = append(ts.shared.Fields, name, metric)
	ts.owner.Structs = append(ts.owner.Structs, ts.shared)

	ts.user = matter.NewCluster(nil)
	ts.user.Name = "Channel"
	ts.spec.DocRefs[ts.user] = ts.userDoc
	ts.attribute = matter.NewAttribute(nil, ts.user)
	ts.attribute.ID = matter.NewNumber(0x12)
	ts.attribute.Name = "ExternalIDList"
	ts.attribute.Type = &types.DataType{BaseType: types.BaseDataTypeList, EntryType: &types.DataType{Name: "AdditionalInfoStruct", BaseType: types.BaseDataTypeCustom, Entity: ts.shared}}
	ts.user.Attributes = append(ts.user.Attributes, ts.attribute)
	ts.command = matter.NewCommand(nil, ts.user)
	ts.command.Name = "RecordProgram"
	commandField := matter.NewField(nil, ts.command, types.EntityTypeCommandField)
	commandField.ID = matter.NewNumber(2)
	commandField.Name = "ExternalIDList"
	commandField.Type = &types.DataType{Name: "AdditionalInfoStruct", BaseType: types.BaseDataTypeCustom, Entity: ts.shared}
	ts.command.Fields = append(ts.command.Fields, commandField)
	ts.user.Commands = append(ts.user.Commands, ts.command)
	return ts
}

func forceLocal(names ...string) map[string]*errata.SDKType {
	m := make(map[string]*errata.SDKType)
	for _, n := range names {
		m[n] = &errata.SDKType{ForceLocal: true}
	}
	return m
}

func fieldEntity(f *matter.Field) types.Entity {
	if f.Type.IsArray() {
		return f.Type.EntryType.Entity
	}
	return f.Type.Entity
}

func findField(fs matter.FieldSet, name string) *matter.Field {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func TestCopyStructIntoReferencingCluster(t *testing.T) {
	ts := newCopyTestSpec()
	sdkErrata := &errata.SDK{Types: &errata.SDKTypes{Structs: forceLocal("AdditionalInfoStruct")}}
	copied := make(map[types.Entity]bool)

	if err := ts.p.copyTypesToCluster(ts.user, sdkErrata, copied); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ts.user.Structs) != 1 {
		t.Fatalf("user cluster structs = %d, want 1", len(ts.user.Structs))
	}
	clone := ts.user.Structs[0]
	if clone == ts.shared {
		t.Fatalf("user cluster holds the original struct, want a copy")
	}
	if clone.Parent() != ts.user {
		t.Fatalf("copy parent = %v, want the user cluster", clone.Parent())
	}
	if fieldEntity(ts.attribute) != clone || fieldEntity(ts.command.Fields[0]) != clone {
		t.Fatalf("user cluster fields should both point to the single copy")
	}
	if !copied[clone] {
		t.Fatalf("copy should be recorded for the uncopied-entry warning")
	}
	if ts.spec.DocRefs[clone] != ts.userDoc || ts.spec.DocRefs[clone.Fields[0]] != ts.userDoc {
		t.Fatalf("copy and its fields should belong to the user cluster's document")
	}

	if len(ts.owner.Structs) != 1 || ts.owner.Structs[0] != ts.shared || ts.shared.Parent() != ts.owner {
		t.Fatalf("owner cluster must keep its original struct")
	}
	// The copy is not listed in copy-enums, so it still uses the owner's enum, and the original is untouched
	if fieldEntity(findField(clone.Fields, "Metric")) != ts.nested || fieldEntity(findField(ts.shared.Fields, "Metric")) != ts.nested {
		t.Fatalf("unlisted nested enum should not be copied")
	}
	if len(ts.user.Enums) != 0 {
		t.Fatalf("user cluster enums = %d, want 0", len(ts.user.Enums))
	}
}

func TestCopyNestedTypesWithoutModifyingOriginal(t *testing.T) {
	ts := newCopyTestSpec()
	sdkErrata := &errata.SDK{Types: &errata.SDKTypes{
		Structs: forceLocal("AdditionalInfoStruct"),
		Enums:   forceLocal("MetricTypeEnum"),
	}}

	if err := ts.p.copyTypesToCluster(ts.user, sdkErrata, make(map[types.Entity]bool)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(ts.user.Structs) != 1 || len(ts.user.Enums) != 1 {
		t.Fatalf("user cluster structs = %d, enums = %d, want 1 and 1", len(ts.user.Structs), len(ts.user.Enums))
	}
	clonedEnum := ts.user.Enums[0]
	if clonedEnum == ts.nested || clonedEnum.Parent() != ts.user {
		t.Fatalf("nested enum should be copied into the user cluster")
	}
	if fieldEntity(findField(ts.user.Structs[0].Fields, "Metric")) != clonedEnum {
		t.Fatalf("copied struct should use the copied enum")
	}
	if fieldEntity(findField(ts.shared.Fields, "Metric")) != ts.nested {
		t.Fatalf("the owner's struct must still use the owner's enum")
	}
}

func TestCopyAppliesReferencingDocumentErrata(t *testing.T) {
	ts := newCopyTestSpec()
	// A single types entry both marks the struct force-local and overrides the copy's fields
	sdkErrata := &errata.SDK{Types: &errata.SDKTypes{Structs: map[string]*errata.SDKType{
		"AdditionalInfoStruct": {ForceLocal: true, Fields: []*errata.SDKType{{Name: "Name", Constraint: "max 64"}}},
	}}}

	if err := ts.p.copyTypesToCluster(ts.user, sdkErrata, make(map[types.Entity]bool)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c := findField(ts.user.Structs[0].Fields, "Name").Constraint; c == nil || c.ASCIIDocString(nil) != "max 64" {
		t.Fatalf("copy constraint = %v, want max 64", c)
	}
	if findField(ts.shared.Fields, "Name").Constraint != nil {
		t.Fatalf("the owner's struct must not get the referencing document's errata")
	}
}

func TestCopyIgnoresUnlistedAndLocalTypes(t *testing.T) {
	ts := newCopyTestSpec()

	// Not marked force-local (an override or keep entry alone doesn't copy): references are left alone
	notForceLocal := &errata.SDK{Types: &errata.SDKTypes{Structs: map[string]*errata.SDKType{
		"AdditionalInfoStruct": {Keep: true, Fields: []*errata.SDKType{{Name: "Name", Constraint: "max 64"}}},
		"OtherStruct":          {ForceLocal: true},
	}}}
	if err := ts.p.copyTypesToCluster(ts.user, notForceLocal, make(map[types.Entity]bool)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fieldEntity(ts.attribute) != ts.shared || len(ts.user.Structs) != 0 {
		t.Fatalf("unlisted struct should not be copied")
	}

	// Listed, but owned by the cluster being processed: nothing to copy
	copied := make(map[types.Entity]bool)
	if err := ts.p.copyTypesToCluster(ts.owner, &errata.SDK{Types: &errata.SDKTypes{Enums: forceLocal("MetricTypeEnum")}}, copied); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ts.owner.Enums) != 1 || len(copied) != 0 {
		t.Fatalf("a cluster's own enum should not be copied")
	}
}
