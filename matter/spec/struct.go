package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toStruct(spec *Specification, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (ms *matter.Struct, err error) {
	name := text.TrimCaseInsensitiveSuffix(library.SectionName(s), " Type")
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading struct \"%s\": %w", name, err)
	}
	ms = matter.NewStruct(s, parent)
	ms.Name = name

	if ti.HeaderRowIndex > 0 {
		firstRow := ti.Rows[0]
		tableCells := firstRow.TableCells()
		if len(tableCells) > 0 {
			cv, rowErr := RenderTableCell(reader, tableCells[0])
			if rowErr == nil {
				cv = strings.ToLower(cv)
				if strings.Contains(cv, "fabric scoped") || strings.Contains(cv, "fabric-scoped") {
					ms.FabricScoping = matter.FabricScopingScoped
				}
			}
		}
	}
	var fieldMap map[string]*matter.Field
	ms.Fields, fieldMap, err = library.readFields(spec, reader, ti, types.EntityTypeStructField, ms)
	if err != nil {
		return
	}
	library.addEntity(s, ms)
	err = library.mapFields(reader, d, s, fieldMap)
	if err != nil {
		return
	}
	ms.Name = CanonicalName(ms.Name)
	return
}

func validateStructs(spec *Specification) {
	for c := range spec.Clusters {
		for _, s := range c.Structs {
			validateFields(spec, s, s.Fields)
		}
	}
	for obj := range spec.GlobalObjects {
		switch obj := obj.(type) {
		case *matter.Struct:
			validateFields(spec, obj, obj.Fields)
		}
	}
}
