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
		if err != nil {
			return
		}
		for _, entity := range entities {
			et := entity.EntityType()
			switch et {
			case types.EntityTypeStruct, types.EntityTypeBitmap, types.EntityTypeEnum:
				globalEntities[entity.EntityType()] = append(globalEntities[entity.EntityType()], entity)
			default:
				slog.Warn("Skipping currently unsupported global entity type", matter.LogEntity("entity", entity))
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
		testEntities = append(testEntities, &matter.Bitmap{
			Name: "TestGlobalBitmap",
			Type: types.NewDataType(types.BaseDataTypeMap32, false),
			Bits: matter.BitSet{
				matter.NewBitmapBit(nil, "0x01", "FirstBit", "", nil),
				matter.NewBitmapBit(nil, "0x02", "SecondBit", "", nil),
			},
		})
	case types.EntityTypeEnum:
		testGlobalEnum := matter.NewEnum(nil, nil)
		testGlobalEnum.Name = "TestGlobalEnum"
		testGlobalEnum.Type = types.NewDataType(types.BaseDataTypeEnum8, false)
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
		testEntities = append(testEntities, &matter.Struct{
			Name: "TestGlobalStruct",
			Fields: matter.FieldSet{
				&matter.Field{
					ID:   matter.NewNumber(0),
					Name: "Name",
					Type: types.NewDataType(types.BaseDataTypeString, false),
					Constraint: &constraint.MaxConstraint{
						Maximum: &constraint.IntLimit{Value: 128},
					},
					Conformance: conformance.Set{&conformance.Mandatory{}},
				},
				&matter.Field{
					ID:          matter.NewNumber(1),
					Name:        "MyBitmap",
					Type:        types.NewCustomDataType("TestGlobalBitmap", false),
					Quality:     matter.QualityNullable,
					Conformance: conformance.Set{&conformance.Mandatory{}},
				},
				&matter.Field{
					ID:          matter.NewNumber(2),
					Name:        "MyEnum",
					Type:        types.NewCustomDataType("TestGlobalEnum", false),
					Quality:     matter.QualityNullable,
					Conformance: conformance.Set{&conformance.Optional{}},
				},
			},
		})
	}
	return
}
