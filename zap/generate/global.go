package generate

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (tg *TemplateGenerator) RenderGlobalObjecs(cxt context.Context) (globalFiles pipeline.Map[string, *pipeline.Data[string]], err error) {
	globalFiles = pipeline.NewMap[string, *pipeline.Data[string]]()
	var globalEntities []types.Entity
	tg.globalObjectDependencies.Range(func(entity types.Entity, _ struct{}) bool {
		globalEntities = append(globalEntities, entity)
		return true
	})
	if len(globalEntities) == 0 {
		return
	}

	var hasBitmaps, hasEnums, hasCommands, hasStructs bool

	for _, e := range globalEntities {
		switch e.(type) {
		case *matter.Bitmap:
			hasBitmaps = true
		case *matter.Command:
			hasCommands = true
		case *matter.Enum:
			hasEnums = true
		case *matter.Struct:
			hasStructs = true
		}
	}

	if hasBitmaps {
		globalBitmaps := getGlobalEntities[*matter.Bitmap](tg.spec)
		testBitmap := &matter.Bitmap{
			Name: "TestGlobalBitmap",
			Type: types.NewDataType(types.BaseDataTypeMap32, false),
			Bits: matter.BitSet{
				matter.NewBitmapBit("0x01", "FirstBit", "", nil),
				matter.NewBitmapBit("0x02", "SecondBit", "", nil),
			},
		}
		globalBitmaps[testBitmap] = []*matter.Number{matter.InvalidID}
		err = saveGlobalEntities(cxt, tg, "global-bitmaps", globalBitmaps, generateBitmaps, globalFiles)
		if err != nil {
			return
		}
	}

	if hasEnums {
		globalEnums := getGlobalEntities[*matter.Enum](tg.spec)
		testEnum := &matter.Enum{
			Name: "TestGlobalEnum",
			Type: types.NewDataType(types.BaseDataTypeEnum8, false),
			Values: matter.EnumValueSet{
				&matter.EnumValue{
					Value: matter.NewNumber(0x0),
					Name:  "SomeValue",
				},
				&matter.EnumValue{
					Value: matter.NewNumber(0x1),
					Name:  "SomeOtherValue",
				},
				&matter.EnumValue{
					Value: matter.NewNumber(0x2),
					Name:  "FinalValue",
				},
			},
		}
		globalEnums[testEnum] = []*matter.Number{matter.InvalidID}
		err = saveGlobalEntities(cxt, tg, "global-enums", globalEnums, generateEnums, globalFiles)
		if err != nil {
			return
		}
	}

	if hasCommands {
		err = saveGlobalEntities(cxt, tg, "global-commands", getGlobalEntities[*matter.Command](tg.spec), generateCommands, globalFiles)
		if err != nil {
			return
		}
	}

	if hasStructs {
		globalStructs := getGlobalEntities[*matter.Struct](tg.spec)
		testStruct := &matter.Struct{
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
		}
		globalStructs[testStruct] = []*matter.Number{matter.InvalidID}
		err = saveGlobalEntities(cxt, tg, "global-structs", globalStructs, generateStructs, globalFiles)
		if err != nil {
			return
		}
	}

	return
}

func getGlobalEntities[T comparable](spec *spec.Specification) map[T][]*matter.Number {
	ge := make(map[T][]*matter.Number)
	for globalObject := range spec.GlobalObjects {
		if globalObject, ok := globalObject.(T); ok {
			ge[globalObject] = []*matter.Number{matter.InvalidID}
		}
	}
	return ge
}

type globalEntityGenerator[T comparable] func(entities map[T][]*matter.Number, sourcePath string, parent *etree.Element, errata *zap.Errata) (err error)

func saveGlobalEntities[T comparable](cxt context.Context, tg *TemplateGenerator, path string, entities map[T][]*matter.Number, generator globalEntityGenerator[T], globalFiles pipeline.Map[string, *pipeline.Data[string]]) (err error) {
	var bitmapDoc *etree.Document
	var bitmapRoot *etree.Element
	bitmapPath := getZapPath(tg.sdkRoot, path)
	bitmapDoc, bitmapRoot, err = tg.createGlobalXMLFile(cxt, bitmapPath)
	if err != nil {
		return
	}

	err = generator(entities, bitmapPath, bitmapRoot, nil)
	if err != nil {
		return
	}
	var s string
	s, err = xmlToString(bitmapDoc)
	if err != nil {
		return
	}
	globalFiles.Store(filepath.Base(bitmapPath), pipeline.NewData[string](bitmapPath, s))
	return
}

func (tg *TemplateGenerator) createGlobalXMLFile(cxt context.Context, path string) (doc *etree.Document, root *etree.Element, err error) {

	var existing []byte
	existing, err = os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) || tg.file.DryRun {
		if tg.pipeline.Serial {
			slog.InfoContext(cxt, "Rendering new Global ZAP template", "to", path)
		}
		doc = newZapTemplate()
	} else if err != nil {
		return
	} else {
		if tg.pipeline.Serial {
			slog.InfoContext(cxt, "Rendering existing Global ZAP template", "path", path)
		}
		doc = etree.NewDocument()
		err = doc.ReadFromBytes(existing)
		if err != nil {
			err = fmt.Errorf("failed reading ZAP template %s: %w", path, err)
			return
		}

	}

	root = doc.SelectElement("configurator")
	if root == nil {
		root = doc.CreateElement("configurator")
	}

	de := root.SelectElement("domain")
	if de == nil {
		de = etree.NewElement("domain")
		xml.AppendElement(root, de)
		de.CreateAttr("name", "CHIP")
	}
	return
}
