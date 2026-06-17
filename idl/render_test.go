package idl

import (
	"context"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func TestEntityShouldBeIncluded(t *testing.T) {
	// Create a dummy specification
	specification := &spec.Specification{}

	// Let's create an entity that has provisional conformance
	provField := &matter.Field{
		Name:        "ProvField",
		ID:          matter.NewNumber(1),
		Conformance: conformance.Set{&conformance.Provisional{}},
	}

	nonProvField := &matter.Field{
		Name:        "NonProvField",
		ID:          matter.NewNumber(2),
		Conformance: conformance.Set{&conformance.Mandatory{}},
	}

	// 1. Test Mode = "none"
	filterNone := ProvisionalFilter{Mode: "none"}
	if !entityShouldBeIncluded(specification, filterNone, provField) {
		t.Errorf("expected provisional field to be included with Mode=none")
	}
	if !entityShouldBeIncluded(specification, filterNone, nonProvField) {
		t.Errorf("expected non-provisional field to be included with Mode=none")
	}

	// 2. Test Mode = "all"
	filterAll := ProvisionalFilter{Mode: "all"}
	if entityShouldBeIncluded(specification, filterAll, provField) {
		t.Errorf("expected provisional field to be excluded with Mode=all")
	}
	if !entityShouldBeIncluded(specification, filterAll, nonProvField) {
		t.Errorf("expected non-provisional field to be included with Mode=all")
	}

}

func TestSuppressProvisionalIntegration(t *testing.T) {
	// 1. Mock specification
	specification := &spec.Specification{
		ClustersByID:    make(map[uint64]*matter.Cluster),
		ClustersByName:  make(map[string]*matter.Cluster),
		DeviceTypesByID: make(map[uint64]*matter.DeviceType),
		DataTypeRefs:    spec.NewEntityRefs[types.Entity](),
		ClusterRefs:     spec.NewEntityRefs[*matter.Cluster](),
	}

	specification.DeviceTypesByID[1] = &matter.DeviceType{
		Name: "Test Device Type",
		ID:   matter.NewNumber(1),
	}

	cluster := matter.NewCluster(nil)
	cluster.Name = "MyCluster"
	cluster.ID = matter.NewNumber(1)
	cluster.Conformance = conformance.Set{&conformance.Mandatory{}} // Non-provisional

	specification.ClustersByID[1] = cluster
	specification.ClustersByName["MyCluster"] = cluster

	provCluster := matter.NewCluster(nil)
	provCluster.Name = "ProvCluster"
	provCluster.ID = matter.NewNumber(2)
	provCluster.Conformance = conformance.Set{&conformance.Provisional{}}

	specification.ClustersByID[2] = provCluster
	specification.ClustersByName["ProvCluster"] = provCluster

	provCluster2 := matter.NewCluster(nil)
	provCluster2.Name = "ProvCluster2"
	provCluster2.ID = matter.NewNumber(3)
	provCluster2.Conformance = conformance.Set{&conformance.Provisional{}}

	specification.ClustersByID[3] = provCluster2
	specification.ClustersByName["ProvCluster2"] = provCluster2

	// Enums
	provEnum := &matter.Enum{Name: "ProvEnum", Type: types.NewDataType(types.BaseDataTypeEnum8, types.DataTypeRankScalar)}
	provEnum.SetParent(cluster)
	cluster.AddEnums(provEnum)

	provEnumVal := matter.NewEnumValue(nil, provEnum)
	provEnumVal.Name = "kValue"
	provEnumVal.Value = matter.NewNumber(0)
	provEnumVal.Conformance = conformance.Set{&conformance.Mandatory{}}
	provEnum.Values = append(provEnum.Values, provEnumVal)

	nonProvEnum := &matter.Enum{Name: "NonProvEnum", Type: types.NewDataType(types.BaseDataTypeEnum8, types.DataTypeRankScalar)}
	nonProvEnum.SetParent(cluster)
	cluster.AddEnums(nonProvEnum)

	nonProvEnumVal := matter.NewEnumValue(nil, nonProvEnum)
	nonProvEnumVal.Name = "kValue"
	nonProvEnumVal.Value = matter.NewNumber(0)
	nonProvEnumVal.Conformance = conformance.Set{&conformance.Mandatory{}}
	nonProvEnum.Values = append(nonProvEnum.Values, nonProvEnumVal)

	// Structs
	provStruct := &matter.Struct{Name: "ProvStruct"}
	provStruct.SetParent(cluster)
	cluster.AddStructs(provStruct)

	provStructField := &matter.Field{
		Name:        "a",
		ID:          matter.NewNumber(1),
		Conformance: conformance.Set{&conformance.Mandatory{}},
		Type:        types.NewCustomDataType("int8u", types.DataTypeRankScalar),
	}
	provStructField.SetParent(provStruct)
	provStruct.Fields = append(provStruct.Fields, provStructField)

	nonProvStruct := &matter.Struct{Name: "NonProvStruct"}
	nonProvStruct.SetParent(cluster)
	cluster.AddStructs(nonProvStruct)

	nonProvStructField := &matter.Field{
		Name:        "a",
		ID:          matter.NewNumber(1),
		Conformance: conformance.Set{&conformance.Mandatory{}},
		Type:        types.NewCustomDataType("int8u", types.DataTypeRankScalar),
	}
	nonProvStructField.SetParent(nonProvStruct)
	nonProvStruct.Fields = append(nonProvStruct.Fields, nonProvStructField)

	// Bitmaps
	provBitmap := &matter.Bitmap{Name: "ProvBitmap", Type: types.NewDataType(types.BaseDataTypeMap8, types.DataTypeRankScalar)}
	provBitmap.SetParent(cluster)
	cluster.AddBitmaps(provBitmap)

	provBitmapBit := matter.NewBitmapBit(nil, provBitmap, "0", "MyBit", "Summary", conformance.Set{&conformance.Mandatory{}})
	provBitmap.AddBit(provBitmapBit)

	nonProvBitmap := &matter.Bitmap{Name: "NonProvBitmap", Type: types.NewDataType(types.BaseDataTypeMap8, types.DataTypeRankScalar)}
	nonProvBitmap.SetParent(cluster)
	cluster.AddBitmaps(nonProvBitmap)

	nonProvBitmapBit := matter.NewBitmapBit(nil, nonProvBitmap, "0", "MyBit", "Summary", conformance.Set{&conformance.Mandatory{}})
	nonProvBitmap.AddBit(nonProvBitmapBit)

	// Attributes
	provAttrEnum := &matter.Field{
		Name:        "ProvAttrEnum",
		ID:          matter.NewNumber(1),
		Conformance: conformance.Set{&conformance.Provisional{}},
		Type:        types.NewCustomDataType("ProvEnum", types.DataTypeRankScalar),
	}
	provAttrEnum.SetParent(cluster)
	provAttrEnum.Type.Entity = provEnum

	provAttrStruct := &matter.Field{
		Name:        "ProvAttrStruct",
		ID:          matter.NewNumber(2),
		Conformance: conformance.Set{&conformance.Provisional{}},
		Type:        types.NewCustomDataType("ProvStruct", types.DataTypeRankScalar),
	}
	provAttrStruct.SetParent(cluster)
	provAttrStruct.Type.Entity = provStruct

	provAttrBitmap := &matter.Field{
		Name:        "ProvAttrBitmap",
		ID:          matter.NewNumber(3),
		Conformance: conformance.Set{&conformance.Provisional{}},
		Type:        types.NewCustomDataType("ProvBitmap", types.DataTypeRankScalar),
	}
	provAttrBitmap.SetParent(cluster)
	provAttrBitmap.Type.Entity = provBitmap

	nonProvAttrEnum := &matter.Field{
		Name:        "NonProvAttrEnum",
		ID:          matter.NewNumber(4),
		Conformance: conformance.Set{&conformance.Mandatory{}},
		Type:        types.NewCustomDataType("NonProvEnum", types.DataTypeRankScalar),
	}
	nonProvAttrEnum.SetParent(cluster)
	nonProvAttrEnum.Type.Entity = nonProvEnum

	nonProvAttrStruct := &matter.Field{
		Name:        "NonProvAttrStruct",
		ID:          matter.NewNumber(5),
		Conformance: conformance.Set{&conformance.Mandatory{}},
		Type:        types.NewCustomDataType("NonProvStruct", types.DataTypeRankScalar),
	}
	nonProvAttrStruct.SetParent(cluster)
	nonProvAttrStruct.Type.Entity = nonProvStruct

	nonProvAttrBitmap := &matter.Field{
		Name:        "NonProvAttrBitmap",
		ID:          matter.NewNumber(6),
		Conformance: conformance.Set{&conformance.Mandatory{}},
		Type:        types.NewCustomDataType("NonProvBitmap", types.DataTypeRankScalar),
	}
	nonProvAttrBitmap.SetParent(cluster)
	nonProvAttrBitmap.Type.Entity = nonProvBitmap

	cluster.Attributes = matter.FieldSet{
		provAttrEnum, provAttrStruct, provAttrBitmap,
		nonProvAttrEnum, nonProvAttrStruct, nonProvAttrBitmap,
	}

	// Commands
	provCmd := &matter.Command{
		Name:        "ProvCmd",
		ID:          matter.NewNumber(1),
		Conformance: conformance.Set{&conformance.Provisional{}},
	}
	provCmd.SetParent(cluster)

	nonProvCmd := &matter.Command{
		Name:        "NonProvCmd",
		ID:          matter.NewNumber(2),
		Conformance: conformance.Set{&conformance.Mandatory{}},
	}
	nonProvCmd.SetParent(cluster)

	cluster.Commands = matter.CommandSet{provCmd, nonProvCmd}

	// Events
	provEvt := &matter.Event{
		Name:        "ProvEvt",
		ID:          matter.NewNumber(1),
		Conformance: conformance.Set{&conformance.Provisional{}},
	}
	provEvt.SetParent(cluster)

	nonProvEvt := &matter.Event{
		Name:        "NonProvEvt",
		ID:          matter.NewNumber(2),
		Conformance: conformance.Set{&conformance.Mandatory{}},
	}
	nonProvEvt.SetParent(cluster)

	cluster.Events = matter.EventSet{provEvt, nonProvEvt}

	// Register references
	specification.DataTypeRefs.Add(provAttrEnum, provEnum)
	specification.DataTypeRefs.Add(nonProvAttrEnum, nonProvEnum)
	specification.DataTypeRefs.Add(provAttrStruct, provStruct)
	specification.DataTypeRefs.Add(nonProvAttrStruct, nonProvStruct)
	specification.DataTypeRefs.Add(provAttrBitmap, provBitmap)
	specification.DataTypeRefs.Add(nonProvAttrBitmap, nonProvBitmap)

	specification.ClusterRefs.Add(cluster, provCmd)
	specification.ClusterRefs.Add(cluster, nonProvCmd)
	specification.ClusterRefs.Add(cluster, provEvt)
	specification.ClusterRefs.Add(cluster, nonProvEvt)

	// Create synthetic zap file data
	syntheticFile := &File{
		EndpointTypes: []EndpointType{
			{
				ID:             0,
				Name:           "Test Endpoint",
				DeviceTypeCode: 1,
				Clusters: []ClusterRef{
					{
						Code: 1,
						Name: "MyCluster",
						Side: "server",
					},
					{
						Code: 2,
						Name: "ProvCluster",
						Side: "server",
					},
					{
						Code: 3,
						Name: "ProvCluster2",
						Side: "server",
					},
				},
			},
		},
		Endpoints: []JSONEndpoint{
			{
				EndpointId:        0,
				EndpointTypeIndex: 0,
			},
		},
	}

	input := pipeline.NewData("test.matter", syntheticFile)
	ctx := context.Background()

	// 2. Test Case A: --suppress-provisional none
	rendererNone, err := NewIdlRenderer(specification)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}
	rendererNone.SuppressProvisional = "none"

	outputsNone, _, err := rendererNone.Process(ctx, input, 0, 1)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}
	if len(outputsNone) == 0 {
		t.Fatalf("expected output, got none")
	}
	contentNone := outputsNone[0].Content
	// Check that BOTH provisional and non-provisional elements are present
	if !strings.Contains(contentNone, "enum ProvEnum") || !strings.Contains(contentNone, "enum NonProvEnum") {
		t.Errorf("expected enums in none output: %s", contentNone)
	}
	if !strings.Contains(contentNone, "struct ProvStruct") || !strings.Contains(contentNone, "struct NonProvStruct") {
		t.Errorf("expected structs in none output: %s", contentNone)
	}
	if !strings.Contains(contentNone, "bitmap ProvBitmap") || !strings.Contains(contentNone, "bitmap NonProvBitmap") {
		t.Errorf("expected bitmaps in none output: %s", contentNone)
	}
	if !strings.Contains(contentNone, "provAttrEnum") || !strings.Contains(contentNone, "nonProvAttrEnum") {
		t.Errorf("expected attributes in none output: %s", contentNone)
	}
	if !strings.Contains(contentNone, "struct ProvCmd") || !strings.Contains(contentNone, "struct NonProvCmd") {
		t.Errorf("expected commands in none output: %s", contentNone)
	}
	if !strings.Contains(contentNone, "optional event ProvEvt") || !strings.Contains(contentNone, "event NonProvEvt") {
		t.Errorf("expected events in none output: %s", contentNone)
	}

	// 3. Test Case B: --suppress-provisional all
	rendererAll, err := NewIdlRenderer(specification)
	if err != nil {
		t.Fatalf("failed to create renderer: %v", err)
	}
	rendererAll.SuppressProvisional = "all"

	outputsAll, _, err := rendererAll.Process(ctx, input, 0, 1)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}
	contentAll := outputsAll[0].Content
	// Check that ONLY non-provisional elements are present, and ALL provisional elements are suppressed
	if strings.Contains(contentAll, "enum ProvEnum") || !strings.Contains(contentAll, "enum NonProvEnum") {
		t.Errorf("expected enums in all output (ProvEnum suppressed, NonProvEnum kept): %s", contentAll)
	}
	if strings.Contains(contentAll, "struct ProvStruct") || !strings.Contains(contentAll, "struct NonProvStruct") {
		t.Errorf("expected structs in all output (ProvStruct suppressed, NonProvStruct kept): %s", contentAll)
	}
	if strings.Contains(contentAll, "bitmap ProvBitmap") || !strings.Contains(contentAll, "bitmap NonProvBitmap") {
		t.Errorf("expected bitmaps in all output (ProvBitmap suppressed, NonProvBitmap kept): %s", contentAll)
	}
	if strings.Contains(contentAll, "provAttrEnum") || !strings.Contains(contentAll, "nonProvAttrEnum") {
		t.Errorf("expected attributes in all output (ProvAttrEnum suppressed, NonProvAttrEnum kept): %s", contentAll)
	}
	if strings.Contains(contentAll, "struct ProvCmd") || !strings.Contains(contentAll, "struct NonProvCmd") {
		t.Errorf("expected commands in all output (ProvCmd suppressed, NonProvCmd kept): %s", contentAll)
	}
	if strings.Contains(contentAll, "optional event ProvEvt") || !strings.Contains(contentAll, "event NonProvEvt") {
		t.Errorf("expected events in all output (ProvEvt suppressed, NonProvEvt kept): %s", contentAll)
	}
}

func TestGetClusterFileName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"On/Off", "on_off.matter"},
		{"PM10 Concentration Measurement", "pm10_concentration_measurement.matter"},
		{"PM1 Concentration Measurement", "pm1_concentration_measurement.matter"},
		{"PM2.5 Concentration Measurement", "pm25_concentration_measurement.matter"},
		{"Door Lock", "door_lock.matter"},
		{"Thermostat", "thermostat.matter"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := getClusterFileName(tt.input)
			if got != tt.want {
				t.Errorf("getClusterFileName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
