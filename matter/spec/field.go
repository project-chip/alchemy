package spec

import (
	"log/slog"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) mapFields(fieldMap map[string]*matter.Field, entityMap map[asciidoc.Attributable][]types.Entity) error {
	for _, s := range parse.Skim[*Section](s.Elements()) {
		var name string
		switch s.SecType {
		case matter.SectionAttribute:
			name = text.TrimCaseInsensitiveSuffix(s.Name, " Attribute")
		case matter.SectionField:
			name = text.TrimCaseInsensitiveSuffix(s.Name, " Field")
		}
		if len(name) == 0 {
			continue
		}
		a, ok := fieldMap[name]
		if !ok {
			continue
		}
		findAnonymousType(s, a)
		entityMap[s.Base] = append(entityMap[s.Base], a)
	}
	return nil
}

func findAnonymousType(s *Section, field *matter.Field) error {
	if field.Type == nil {
		return nil
	}
	if field.Type.IsEnum() {
		return findAnonymousEnum(s, field)
	}
	if field.Type.IsMap() {
		return findAnonymousBitmap(s, field)
	}
	return nil
}

func findAnonymousEnum(s *Section, field *matter.Field) error {
	slog.Debug("possible anonymous enum", "name", field.Name, "type", field.Type)
	ti, err := parseFirstTable(s.Doc, s)
	if err != nil {
		if err == ErrNoTableFound {
			return nil
		}
	}
	valueIndex, ok := ti.ColumnMap[matter.TableColumnValue]
	if !ok {
		slog.Debug("no value", "name", field.Name, "type", field.Type)
		return nil
	}
	var evs matter.EnumValueSet
	for row := range ti.Body() {
		ev := matter.NewEnumValue(s.Base)
		ev.Conformance = conformance.Set{&conformance.Mandatory{}}
		ev.Value, err = ti.ReadID(row, matter.TableColumnValue)
		if err != nil {
			return err
		}
		ev.Summary, err = ti.ReadValue(row, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return err
		}
		if len(ev.Summary) == 0 {
			if len(row.TableCells()) > valueIndex+1 {
				ev.Summary, err = ti.ReadValueByIndex(row, valueIndex+1)
				if err != nil {
					return err
				}
			}
		}
		if strings.Contains(ev.Summary, " ") {
			ev.Name = strcase.ToCamel(ev.Summary)
		} else {
			ev.Name = ev.Summary
		}
		evs = append(evs, ev)
	}
	if len(evs) > 0 {
		field.AnonymousType = &matter.AnonymousEnum{
			Type:   field.Type,
			Values: evs,
		}
	}
	return nil
}

func findAnonymousBitmap(s *Section, field *matter.Field) error {
	slog.Debug("possible anonymous enum", "name", field.Name, "type", field.Type)
	ti, err := parseFirstTable(s.Doc, s)
	if err != nil {
		if err == ErrNoTableFound {
			return nil
		}
	}
	_, ok := ti.ColumnMap[matter.TableColumnBit]
	if !ok {
		slog.Debug("no bit", "name", field.Name, "type", field.Type)
		return nil
	}
	var bvs matter.BitSet
	for row := range ti.Body() {
		var bit, name, summary string
		conf := conformance.Set{&conformance.Mandatory{}}
		name, err = ti.ReadValue(row, matter.TableColumnName)
		if err != nil {
			return err
		}
		name = matter.StripTypeSuffixes(name)
		summary, err = ti.ReadValue(row, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return err
		}

		bit, err = ti.ReadString(row, matter.TableColumnBit)
		if err != nil {
			return err
		}
		if len(bit) == 0 {
			bit, err = ti.ReadString(row, matter.TableColumnValue)
			if err != nil {
				return err
			}
		}
		if len(summary) == 0 && len(name) > 0 {
			summary = name
			if strings.Contains(name, " ") {
				name = matter.Case(name)
			}
		} else if len(name) == 0 && len(summary) > 0 {
			name = matter.Case(summary)
		}

		bv := matter.NewBitmapBit(s.Base, bit, name, summary, conf)

		bvs = append(bvs, bv)
	}
	if len(bvs) > 0 {
		field.AnonymousType = &matter.AnonymousBitmap{
			Type: field.Type,
			Bits: bvs,
		}
	}
	return nil
}
