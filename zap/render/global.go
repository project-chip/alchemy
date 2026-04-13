package render

import (
	"context"
	"fmt"
	"log/slog"
	"slices"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/find"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

type GlobalObjectsRenderer struct {
	spec      *spec.Specification
	sdkRoot   string
	generator *TemplateGenerator
}

func NewGlobalObjectsRenderer(spec *spec.Specification, sdkRoot string, generator *TemplateGenerator) *GlobalObjectsRenderer {
	dt := &GlobalObjectsRenderer{
		spec:      spec,
		sdkRoot:   sdkRoot,
		generator: generator,
	}

	return dt
}

func (p GlobalObjectsRenderer) Name() string {
	return "Renering ZAP global objects"
}

func (p GlobalObjectsRenderer) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[string], err error) {

	globalEntities := make(map[types.EntityType][]types.Entity)
	for _, input := range inputs {
		entities := p.spec.EntitiesForDocument(input.Content)
		for _, entity := range entities {
			et := entity.EntityType()
			switch et {
			case types.EntityTypeStruct, types.EntityTypeBitmap, types.EntityTypeEnum:
				globalEntities[entity.EntityType()] = append(globalEntities[entity.EntityType()], entity)
			default:
				slog.Warn("Skipping global entity type currently unsupported in ZAP", matter.LogEntity("entity", entity))
			}
		}
	}
	for entityType := range globalEntities {
		var outPath string
		outPath, err = p.getGlobalPath(entityType)
		if err != nil {
			return
		}
		allEntities := slices.Collect(find.Filter(find.Keys(p.spec.GlobalObjects), func(e types.Entity) bool { return e.EntityType() == entityType }))
		docs := make(map[*asciidoc.Document]struct{})
		for _, e := range allEntities {
			doc, ok := p.spec.DocRefs[e]
			if !ok {
				slog.Warn("missing doc ref for global entity", slog.String("entityType", entityType.String()))
			} else {
				docs[doc] = struct{}{}
			}
		}
		allEntities = append(allEntities, getGlobalTestEntites(entityType)...)
		var configurator *zap.Configurator
		configurator, err = zap.NewConfigurator(p.spec, find.Keys(docs), allEntities, outPath, &errata.DefaultErrata.SDK, true)
		if err != nil {
			return
		}
		configurator.Domain = "CHIP"

		var doc *etree.Document
		doc, err = openConfigurator(configurator, p.generator.pipeline)
		if err != nil {
			return
		}

		cr := newConfiguratorRenderer(p.generator, configurator)
		var out string
		out, err = cr.render(doc, nil)
		outputs = append(outputs, pipeline.NewData(configurator.OutPath, out))
	}
	if len(globalEntities) == 0 {
		return
	}
	return
}

func (tg *GlobalObjectsRenderer) getGlobalPath(entityType types.EntityType) (path string, err error) {
	switch entityType {
	case types.EntityTypeBitmap:
		path = "global-bitmaps"
	case types.EntityTypeEnum:
		path = "global-enums"
	case types.EntityTypeStruct:
		path = "global-structs"
	case types.EntityTypeCommand:
		path = "global-commands"
	case types.EntityTypeEvent:
		path = "global-events"
	case types.EntityTypeDef:
		path = "global-typedefs"
	default:
		err = fmt.Errorf("unexpected global entity type: %s", entityType.String())
		return
	}
	path = getZapPath(tg.sdkRoot, path)
	return
}

func getGlobalTestEntites(entityType types.EntityType) (testEntities []types.Entity) {
	switch entityType {
	case types.EntityTypeBitmap:
		testGlobalBitmap := matter.NewBitmap(nil, nil)
		firstBit := matter.NewBitmapBit(nil, testGlobalBitmap, "0x01", "FirstBit", "", nil)
		secondBit := matter.NewBitmapBit(nil, testGlobalBitmap, "0x02", "SecondBit", "", nil)
		testGlobalBitmap.Bits = append(testGlobalBitmap.Bits, firstBit, secondBit)
		testEntities = append(testEntities, testGlobalBitmap)
	case types.EntityTypeEnum:
		testGlobalEnum := matter.NewEnum(nil, nil)
		testGlobalEnum.Name = "TestGlobalEnum"
		testGlobalEnum.Type = types.NewDataType(types.BaseDataTypeEnum8, types.DataTypeRankScalar)
		someValue := matter.NewEnumValue(nil, testGlobalEnum)
		someValue.Name = "SomeValue"
		someValue.Value = matter.NewNumber(0x0)
		someOtherValue := matter.NewEnumValue(nil, testGlobalEnum)
		someOtherValue.Name = "SomeOtherValue"
		someOtherValue.Value = matter.NewNumber(0x1)
		finalValue := matter.NewEnumValue(nil, testGlobalEnum)
		finalValue.Name = "FinalValue"
		finalValue.Value = matter.NewNumber(0x2)
		testGlobalEnum.Values = matter.EnumValueSet{someValue, someOtherValue, finalValue}
		testEntities = append(testEntities, testGlobalEnum)
	case types.EntityTypeStruct:
		testGlobalStruct := matter.NewStruct(nil, nil)
		testGlobalStruct.Name = "TestGlobalStruct"
		nameField := matter.NewField(nil, testGlobalStruct, types.EntityTypeStructField)
		nameField.ID = matter.NewNumber(0)
		nameField.Name = "Name"
		nameField.Type = types.NewDataType(types.BaseDataTypeString, types.DataTypeRankScalar)
		nameField.Constraint = &constraint.MaxConstraint{
			Maximum: &constraint.IntLimit{Value: 128},
		}
		nameField.Conformance = conformance.Set{&conformance.Mandatory{}}
		myBitmapField := matter.NewField(nil, testGlobalStruct, types.EntityTypeStructField)
		myBitmapField.ID = matter.NewNumber(1)
		myBitmapField.Name = "MyBitmap"
		myBitmapField.Type = types.NewCustomDataType("TestGlobalBitmap", types.DataTypeRankScalar)
		myBitmapField.Quality = matter.QualityNullable
		myBitmapField.Conformance = conformance.Set{&conformance.Mandatory{}}
		myEnumField := matter.NewField(nil, testGlobalStruct, types.EntityTypeStructField)
		myEnumField.ID = matter.NewNumber(2)
		myEnumField.Name = "MyEnum"
		myEnumField.Type = types.NewCustomDataType("TestGlobalEnum", types.DataTypeRankScalar)
		myEnumField.Quality = matter.QualityNullable
		myEnumField.Conformance = conformance.Set{&conformance.Mandatory{}}
		testGlobalStruct.Fields = append(testGlobalStruct.Fields, nameField, myBitmapField, myEnumField)
		testEntities = append(testEntities, testGlobalStruct)
	}
	return
}
