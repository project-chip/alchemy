package dm

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func renderGlobalDataType[T types.Entity](entities []T, entityType types.EntityType, dmRoot string, globalFiles pipeline.StringSet) (err error) {

	if len(entities) == 0 {
		return
	}
	var outPath string
	outPath, err = getGlobalPath(dmRoot, entityType)

	if err != nil {
		return
	}

	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())

	root := &x.Element
	switch entities := any(entities).(type) {
	case []*matter.Bitmap:
		err = renderBitmaps(entities, root.CreateElement("bitmaps"))
	case []*matter.Enum:
		err = renderEnums(entities, root.CreateElement("enums"))
	case []*matter.Struct:
		err = renderStructs(entities, root.CreateElement("structs"))
	case []*matter.TypeDef:
		err = renderTypeDefs(entities, root.CreateElement("typeDefs"))
	case []*matter.Command:
		err = renderCommands(entities, root)
	case []*matter.Event:
		err = renderEvents(entities, root)
	default:
		err = fmt.Errorf("unexpected global data type list: %T", entities)
	}

	if err != nil {
		return
	}

	x.Indent(2)

	var b bytes.Buffer
	_, err = x.WriteTo(&b)
	if err != nil {
		return
	}
	globalFiles.Store(outPath, pipeline.NewData(outPath, b.String()))
	return
}

func (p *Renderer) GenerateGlobalObjects() (globalFiles pipeline.StringSet, err error) {

	globalFiles = pipeline.NewMap[string, *pipeline.Data[string]]()

	var bitmaps []*matter.Bitmap
	var enums []*matter.Enum
	var structs []*matter.Struct
	var typeDefs []*matter.TypeDef
	var commands []*matter.Command
	var events []*matter.Event

	for e := range p.spec.GlobalObjects {
		switch e := e.(type) {
		case *matter.Bitmap:
			bitmaps = append(bitmaps, e)
		case *matter.Enum:
			enums = append(enums, e)
		case *matter.Struct:
			structs = append(structs, e)
		case *matter.TypeDef:
			typeDefs = append(typeDefs, e)
		case *matter.Command:
			commands = append(commands, e)
		case *matter.Event:
			events = append(events, e)
		default:
			err = fmt.Errorf("unexpected global data type: %T", e)
			return
		}
	}

	err = renderGlobalDataType(bitmaps, types.EntityTypeBitmap, p.dmRoot, globalFiles)
	if err != nil {
		return
	}

	err = renderGlobalDataType(enums, types.EntityTypeEnum, p.dmRoot, globalFiles)
	if err != nil {
		return
	}
	err = renderGlobalDataType(structs, types.EntityTypeStruct, p.dmRoot, globalFiles)
	if err != nil {
		return
	}

	err = renderGlobalDataType(typeDefs, types.EntityTypeDef, p.dmRoot, globalFiles)
	if err != nil {
		return
	}

	err = renderGlobalDataType(commands, types.EntityTypeCommand, p.dmRoot, globalFiles)
	if err != nil {
		return
	}
	err = renderGlobalDataType(events, types.EntityTypeEvent, p.dmRoot, globalFiles)
	if err != nil {
		return
	}

	return
}

func getGlobalPath(dmRoot string, entityType types.EntityType) (path string, err error) {
	switch entityType {
	case types.EntityTypeBitmap:
		path = "Bitmaps"
	case types.EntityTypeEnum:
		path = "Enums"
	case types.EntityTypeStruct:
		path = "Structs"
	case types.EntityTypeCommand:
		path = "Commands"
	case types.EntityTypeEvent:
		path = "Events"
	case types.EntityTypeDef:
		path = "TypeDefs"
	default:
		err = fmt.Errorf("unexpected global entity type: %s", entityType.String())
		return
	}
	path = filepath.Join(dmRoot, "/globals", path+".xml")
	return
}
