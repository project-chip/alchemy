package spec

import (
	"log/slog"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (d *Doc) readFields(ti *TableInfo, entityType types.EntityType, parent types.Entity) (fields []*matter.Field, err error) {
	ids := make(map[uint64]*matter.Field)
	for row := range ti.Body() {
		f := matter.NewField(row, parent)
		f.Name, err = ti.ReadValue(row, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Name = matter.StripTypeSuffixes(f.Name)
		f.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		f.Type, err = ti.ReadDataType(row, matter.TableColumnType)
		if err != nil {
			slog.Debug("error reading field data type", slog.String("path", d.Path.String()), slog.String("name", f.Name), slog.Any("error", err))
			err = nil
		}

		f.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		f.Quality, err = ti.ReadQuality(row, entityType, matter.TableColumnQuality)
		if err != nil {
			return
		}
		f.Default, err = ti.ReadString(row, matter.TableColumnDefault)
		if err != nil {
			return
		}

		var a string
		a, err = ti.ReadString(row, matter.TableColumnAccess)
		if err != nil {
			return
		}
		f.Access, _ = ParseAccess(a, entityType)
		f.ID, err = ti.ReadID(row, matter.TableColumnID)
		if err != nil {
			return
		}
		if f.ID.Valid() {
			id := f.ID.Value()
			existing, ok := ids[id]
			if ok {
				slog.Error("duplicate field ID", log.Path("source", f), slog.String("name", f.Name), slog.Uint64("id", id), log.Path("original", existing))
				continue
			}
			ids[id] = f
		}

		if f.Type != nil {
			var cs constraint.Set
			switch f.Type.BaseType {
			case types.BaseDataTypeMessageID:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case types.BaseDataTypeIPAddress:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 4}}, &constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case types.BaseDataTypeIPv4Address:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 4}}}
			case types.BaseDataTypeIPv6Address:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case types.BaseDataTypeIPv6Prefix:
				cs = []constraint.Constraint{&constraint.RangeConstraint{Minimum: &constraint.IntLimit{Value: 1}, Maximum: &constraint.IntLimit{Value: 17}}}
			case types.BaseDataTypeHardwareAddress:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 6}}, &constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 8}}}
			}
			if cs != nil {
				if f.Type.IsArray() {
					lc, ok := f.Constraint.(*constraint.ListConstraint)
					if ok {
						lc.EntryConstraint = constraint.AppendConstraint(lc.EntryConstraint, cs...)
					}
				} else {
					f.Constraint = constraint.AppendConstraint(f.Constraint, cs...)
				}

			}
		}
		f.Name = CanonicalName(f.Name)
		fields = append(fields, f)
	}
	return
}

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
		if a.Type != nil && a.Type.BaseType == types.BaseDataTypeTag {
			findTagNamespace(s, a)
		}
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

func findTagNamespace(s *Section, field *matter.Field) error {
	var found bool
	parse.Traverse(s, s.Set, func(ref *asciidoc.CrossReference, parent parse.HasElements, index int) parse.SearchShould {
		if ref.ID == "ref_StandardNamespaces" {
			label := buildReferenceName(ref.Set)
			name := strings.TrimSpace(text.TrimCaseInsensitiveSuffix(label, " Namespace"))
			if len(name) > 0 {
				field.Type.Name = name
				found = true
				return parse.SearchShouldStop
			}
		}
		return parse.SearchShouldContinue
	})
	if !found {
		slog.Warn("Tag field does not specify namespace", slog.String("field", field.Name), log.Path("origin", field))
	}
	return nil
}
