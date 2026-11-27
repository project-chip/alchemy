package dm

import (
	"testing"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func TestInheritedAttributeConformance(t *testing.T) {
	base := matter.NewCluster(nil)
	base.Name = "Mode Base"

	coreModesFeature := matter.NewFeature(nil, "1", "CoreModes", "COREMODES", "", nil)

	baseSupportedModes := matter.NewAttribute(nil, base)
	baseSupportedModes.ID = matter.NewNumber(0x0000)
	baseSupportedModes.Name = "SupportedModes"
	baseSupportedModes.Conformance = conformance.Set{&conformance.Mandatory{}}

	baseStartUpMode := matter.NewAttribute(nil, base)
	baseStartUpMode.ID = matter.NewNumber(0x0002)
	baseStartUpMode.Name = "StartUpMode"
	baseStartUpMode.Conformance = conformance.Set{&conformance.Optional{}}

	baseCoreModeTags := matter.NewAttribute(nil, base)
	baseCoreModeTags.ID = matter.NewNumber(0x0004)
	baseCoreModeTags.Name = "CoreModeTags"
	baseCoreModeTags.Conformance = conformance.Set{
		&conformance.Provisional{},
		&conformance.Mandatory{Expression: &conformance.IdentifierExpression{ID: "COREMODES", Entity: coreModesFeature}},
	}
	base.Attributes = matter.FieldSet{baseSupportedModes, baseStartUpMode, baseCoreModeTags}

	mid := matter.NewCluster(nil)
	mid.Name = "Mode Mid"
	mid.ParentCluster = base

	derived := matter.NewCluster(nil)
	derived.Name = "Mode Refrigerator"
	derived.ParentCluster = mid

	// Blank conformance inherits mandatory from grandparent base.
	derivedSupportedModes := matter.NewAttribute(nil, derived)
	derivedSupportedModes.ID = matter.NewNumber(0x0000)
	derivedSupportedModes.Name = "SupportedModes"
	derivedSupportedModes.Conformance = conformance.Set{}

	// Explicit override must not be replaced by base conformance.
	derivedStartUpMode := matter.NewAttribute(nil, derived)
	derivedStartUpMode.ID = matter.NewNumber(0x0002)
	derivedStartUpMode.Name = "StartUpMode"
	derivedStartUpMode.Conformance = conformance.Set{&conformance.Disallowed{}}

	// Blank conformance inherits compound conformance (P, COREMODES).
	derivedCoreModeTags := matter.NewAttribute(nil, derived)
	derivedCoreModeTags.ID = matter.NewNumber(0x0004)
	derivedCoreModeTags.Name = "CoreModeTags"

	// Derived-only attribute with blank conformance emits no conformance tag.
	derivedOnly := matter.NewAttribute(nil, derived)
	derivedOnly.ID = matter.NewNumber(0x0080)
	derivedOnly.Name = "CustomAttr"

	derived.Attributes = matter.FieldSet{derivedSupportedModes, derivedStartUpMode, derivedCoreModeTags, derivedOnly}

	doc := etree.NewDocument()
	clusterEl := doc.CreateElement("cluster")
	if err := renderAttributes(derived, clusterEl); err != nil {
		t.Fatalf("renderAttributes failed: %v", err)
	}

	// Ensure overlay entities are not mutated.
	if len(derivedSupportedModes.Conformance) != 0 {
		t.Fatalf("derived attribute conformance was mutated: %v", derivedSupportedModes.Conformance)
	}

	attrs := map[string]*etree.Element{}
	for _, el := range clusterEl.FindElements("attributes/attribute") {
		attrs[el.SelectAttrValue("name", "")] = el
	}

	if attrs["SupportedModes"].FindElement("mandatoryConform") == nil {
		t.Errorf("SupportedModes: expected inherited <mandatoryConform/>")
	}
	if attrs["StartUpMode"].FindElement("disallowConform") == nil {
		t.Errorf("StartUpMode: expected explicit <disallowConform/>")
	}
	if attrs["StartUpMode"].FindElement("optionalConform") != nil {
		t.Errorf("StartUpMode: unexpected <optionalConform/>")
	}
	otherwise := attrs["CoreModeTags"].FindElement("otherwiseConform")
	if otherwise == nil || otherwise.FindElement("provisionalConform") == nil || otherwise.FindElement("mandatoryConform/feature[@name='COREMODES']") == nil {
		t.Errorf("CoreModeTags: expected inherited <otherwiseConform> with provisionalConform and COREMODES feature")
	}
	if len(attrs["CustomAttr"].ChildElements()) != 0 {
		t.Errorf("CustomAttr: expected no conformance element, got %d children", len(attrs["CustomAttr"].ChildElements()))
	}
}

func TestInheritedStructFieldConformance(t *testing.T) {
	base := matter.NewCluster(nil)
	base.Name = "Mode Base"

	baseStruct := matter.NewStruct(nil, base)
	baseStruct.Name = "ModeOptionStruct"

	baseLabel := matter.NewField(nil, baseStruct, types.EntityTypeStructField)
	baseLabel.ID = matter.NewNumber(0)
	baseLabel.Name = "Label"
	baseLabel.Conformance = conformance.Set{&conformance.Mandatory{}}

	baseModeTags := matter.NewField(nil, baseStruct, types.EntityTypeStructField)
	baseModeTags.ID = matter.NewNumber(2)
	baseModeTags.Name = "ModeTags"
	// Sibling field reference in conformance should render <field name="Label"/>, not <struct .../>.
	baseModeTags.Conformance = conformance.Set{
		&conformance.Mandatory{Expression: &conformance.IdentifierExpression{ID: "Label", Entity: baseLabel}},
	}
	baseStruct.Fields = matter.FieldSet{baseLabel, baseModeTags}
	base.Structs = matter.StructSet{baseStruct}

	derived := matter.NewCluster(nil)
	derived.Name = "Mode Refrigerator"
	derived.ParentCluster = base

	derivedStruct := matter.NewStruct(nil, derived)
	derivedStruct.Name = "ModeOptionStruct"

	derivedLabel := matter.NewField(nil, derivedStruct, types.EntityTypeStructField)
	derivedLabel.ID = matter.NewNumber(0)
	derivedLabel.Name = "Label"

	derivedModeTags := matter.NewField(nil, derivedStruct, types.EntityTypeStructField)
	derivedModeTags.ID = matter.NewNumber(2)
	derivedModeTags.Name = "ModeTags"

	derivedStruct.Fields = matter.FieldSet{derivedLabel, derivedModeTags}
	derived.Structs = matter.StructSet{derivedStruct}

	doc := etree.NewDocument()
	dt := doc.CreateElement("dataTypes")
	if err := renderStructs(derived.Structs, dt); err != nil {
		t.Fatalf("renderStructs failed: %v", err)
	}

	fields := map[string]*etree.Element{}
	for _, el := range dt.FindElements("struct[@name='ModeOptionStruct']/field") {
		fields[el.SelectAttrValue("name", "")] = el
	}
	if fields["Label"].FindElement("mandatoryConform") == nil {
		t.Errorf("Label: expected inherited <mandatoryConform/>")
	}
	if fields["ModeTags"].FindElement("mandatoryConform/field[@name='Label']") == nil {
		t.Errorf("ModeTags: expected inherited <mandatoryConform><field name='Label'/></mandatoryConform>")
	}
	if fields["ModeTags"].FindElement("mandatoryConform/struct") != nil {
		t.Errorf("ModeTags: sibling field reference should not emit <struct> wrapper")
	}
}

func TestInheritedCommandEventFeatureEnumBitmapConformance(t *testing.T) {
	base := matter.NewCluster(nil)
	base.Name = "Base"

	// Feature
	baseFeat := matter.NewFeature(nil, "0", "OnOff", "DEPONOFF", "", conformance.Set{&conformance.Optional{}})
	base.Features = matter.NewFeatures(nil, base)
	base.Features.Bits = matter.BitSet{baseFeat}

	// Command + field
	baseCmd := matter.NewCommand(nil, base)
	baseCmd.ID = matter.NewNumber(0)
	baseCmd.Name = "ChangeToMode"
	baseCmd.Direction = matter.InterfaceServer
	baseCmd.Conformance = conformance.Set{&conformance.Mandatory{}}
	baseCmdField := matter.NewField(nil, baseCmd, types.EntityTypeCommandField)
	baseCmdField.ID = matter.NewNumber(0)
	baseCmdField.Name = "NewMode"
	baseCmdField.Conformance = conformance.Set{&conformance.Mandatory{}}
	baseCmd.Fields = matter.FieldSet{baseCmdField}
	base.Commands = matter.CommandSet{baseCmd}

	// Event + field
	baseEv := matter.NewEvent(nil, base)
	baseEv.ID = matter.NewNumber(0)
	baseEv.Name = "StateChanged"
	baseEv.Conformance = conformance.Set{&conformance.Optional{}}
	baseEvField := matter.NewField(nil, baseEv, types.EntityTypeEventField)
	baseEvField.ID = matter.NewNumber(0)
	baseEvField.Name = "NewState"
	baseEvField.Conformance = conformance.Set{&conformance.Mandatory{}}
	baseEv.Fields = matter.FieldSet{baseEvField}
	base.Events = matter.EventSet{baseEv}

	// Enum + value
	baseEnum := matter.NewEnum(nil, base)
	baseEnum.Name = "StatusEnum"
	baseEnumVal := matter.NewEnumValue(nil, baseEnum)
	baseEnumVal.Value = matter.NewNumber(0)
	baseEnumVal.Name = "Success"
	baseEnumVal.Conformance = conformance.Set{&conformance.Mandatory{}}
	baseEnum.Values = matter.EnumValueSet{baseEnumVal}
	base.Enums = matter.EnumSet{baseEnum}

	// Bitmap + bit
	baseBm := matter.NewBitmap(nil, base)
	baseBm.Name = "FlagsBitmap"
	baseBm.Type = types.NewDataType(types.BaseDataTypeMap8, types.DataTypeRankScalar)
	baseBit := matter.NewBitmapBit(nil, baseBm, "0", "Active", "", conformance.Set{&conformance.Mandatory{}})
	baseBm.Bits = matter.BitSet{baseBit}
	base.Bitmaps = matter.BitmapSet{baseBm}

	// Derived cluster with blank conformances
	derived := matter.NewCluster(nil)
	derived.Name = "Derived"
	derived.ParentCluster = base

	derivedFeat := matter.NewFeature(nil, "0", "OnOff", "DEPONOFF", "", nil)
	derived.Features = matter.NewFeatures(nil, derived)
	derived.Features.Bits = matter.BitSet{derivedFeat}

	derivedCmd := matter.NewCommand(nil, derived)
	derivedCmd.ID = matter.NewNumber(0)
	derivedCmd.Name = "ChangeToMode"
	derivedCmdField := matter.NewField(nil, derivedCmd, types.EntityTypeCommandField)
	derivedCmdField.ID = matter.NewNumber(0)
	derivedCmdField.Name = "NewMode"
	derivedCmd.Fields = matter.FieldSet{derivedCmdField}
	derived.Commands = matter.CommandSet{derivedCmd}

	derivedEv := matter.NewEvent(nil, derived)
	derivedEv.ID = matter.NewNumber(0)
	derivedEv.Name = "StateChanged"
	derivedEvField := matter.NewField(nil, derivedEv, types.EntityTypeEventField)
	derivedEvField.ID = matter.NewNumber(0)
	derivedEvField.Name = "NewState"
	derivedEv.Fields = matter.FieldSet{derivedEvField}
	derived.Events = matter.EventSet{derivedEv}

	derivedEnum := matter.NewEnum(nil, derived)
	derivedEnum.Name = "StatusEnum"
	derivedEnumVal := matter.NewEnumValue(nil, derivedEnum)
	derivedEnumVal.Value = matter.NewNumber(0)
	derivedEnumVal.Name = "Success"
	derivedEnum.Values = matter.EnumValueSet{derivedEnumVal}
	derived.Enums = matter.EnumSet{derivedEnum}

	derivedBm := matter.NewBitmap(nil, derived)
	derivedBm.Name = "FlagsBitmap"
	derivedBm.Type = types.NewDataType(types.BaseDataTypeMap8, types.DataTypeRankScalar)
	derivedBit := matter.NewBitmapBit(nil, derivedBm, "0", "Active", "", nil)
	derivedBm.Bits = matter.BitSet{derivedBit}
	derived.Bitmaps = matter.BitmapSet{derivedBm}

	doc := etree.NewDocument()
	clusterEl := doc.CreateElement("cluster")
	if err := renderFeatures(nil, derived, clusterEl); err != nil {
		t.Fatalf("renderFeatures failed: %v", err)
	}
	if err := renderCommands(derived.Commands, clusterEl); err != nil {
		t.Fatalf("renderCommands failed: %v", err)
	}
	if err := renderEvents(derived.Events, clusterEl); err != nil {
		t.Fatalf("renderEvents failed: %v", err)
	}
	dtEl := clusterEl.CreateElement("dataTypes")
	if err := renderEnums(derived.Enums, dtEl); err != nil {
		t.Fatalf("renderEnums failed: %v", err)
	}
	if err := renderBitmaps(derived.Bitmaps, dtEl); err != nil {
		t.Fatalf("renderBitmaps failed: %v", err)
	}

	if clusterEl.FindElement("features/feature[@code='DEPONOFF']/optionalConform") == nil {
		t.Errorf("feature DEPONOFF: expected inherited <optionalConform/>")
	}
	if clusterEl.FindElement("commands/command[@name='ChangeToMode']/mandatoryConform") == nil {
		t.Errorf("command ChangeToMode: expected inherited <mandatoryConform/>")
	}
	if clusterEl.FindElement("commands/command[@name='ChangeToMode']/field[@name='NewMode']/mandatoryConform") == nil {
		t.Errorf("command field NewMode: expected inherited <mandatoryConform/>")
	}
	if clusterEl.FindElement("events/event[@name='StateChanged']/optionalConform") == nil {
		t.Errorf("event StateChanged: expected inherited <optionalConform/>")
	}
	if clusterEl.FindElement("events/event[@name='StateChanged']/field[@name='NewState']/mandatoryConform") == nil {
		t.Errorf("event field NewState: expected inherited <mandatoryConform/>")
	}
	if dtEl.FindElement("enum[@name='StatusEnum']/item[@name='Success']/mandatoryConform") == nil {
		t.Errorf("enum item Success: expected inherited <mandatoryConform/>")
	}
	if dtEl.FindElement("bitmap[@name='FlagsBitmap']/bitfield[@name='Active']/mandatoryConform") == nil {
		t.Errorf("bitmap bit Active: expected inherited <mandatoryConform/>")
	}
}
