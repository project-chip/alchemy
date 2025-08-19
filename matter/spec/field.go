package spec

import (
	"log/slog"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (d *Doc) readFields(spec *Specification, ti *TableInfo, entityType types.EntityType, parent types.Entity) (fields []*matter.Field, fieldMap map[string]*matter.Field, err error) {
	fieldMap = make(map[string]*matter.Field)
	for row := range ti.ContentRows() {
		f := matter.NewField(row, parent, entityType)
		var name string
		name, err = ti.ReadValue(row, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Name = matter.StripTypeSuffixes(name)
		f.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		f.Type, err = ti.ReadDataType(ti.Doc.Reader(), row, matter.TableColumnType)
		if err != nil {
			if !conformance.IsDeprecated(f.Conformance) && !conformance.IsDisallowed(f.Conformance) {
				// Clusters inheriting from other clusters don't supply type information, nor do attributes that are deprecated or disallowed
				slog.Debug("error reading field data type", slog.String("path", d.Path.String()), slog.String("name", name), slog.Any("error", err))
				return
			} else {
				err = nil
			}
		}

		f.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		f.Quality, err = ti.ReadQuality(row, entityType, matter.TableColumnQuality)
		if err != nil {
			return
		}
		f.Fallback = ti.ReadFallback(row, matter.TableColumnFallback)

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
						if constraint.IsAllOrEmpty(lc.EntryConstraint) {
							lc.EntryConstraint = cs
						} else {
							lc.EntryConstraint = constraint.AppendConstraint(lc.EntryConstraint, cs...)
						}
					}
				} else {
					if constraint.IsAllOrEmpty(f.Constraint) {
						f.Constraint = cs
					} else {
						f.Constraint = constraint.AppendConstraint(f.Constraint, cs...)
					}
				}

			}
		}
		f.Name = CanonicalName(f.Name)

		fields = append(fields, f)
		if existing, ok := fieldMap[f.Name]; ok {
			slog.Error("duplicate field name", slog.String("name", name), matter.LogEntity("parent", parent), log.Path("source", f), log.Path("previous", existing))
			spec.addError(&DuplicateEntityIDError{Entity: f, Previous: existing})
		}
		fieldMap[f.Name] = f
	}
	return
}

func mapFields(d *Doc, section *asciidoc.Section, fieldMap map[string]*matter.Field, pc *parseContext) error {
	for s := range parse.Skim[*asciidoc.Section](d.Reader(), section, section.Children()) {

		var name string
		switch d.SectionType(s) {
		case matter.SectionAttribute:
			name = text.TrimCaseInsensitiveSuffix(d.SectionName(s), " Attribute")
		case matter.SectionField:
			name = text.TrimCaseInsensitiveSuffix(d.SectionName(s), " Field")
		}
		if len(name) == 0 {
			continue
		}
		a, ok := fieldMap[name]
		if !ok {
			continue
		}
		err := findAnonymousType(d, s, a)
		if err != nil {
			return err
		}
		if a.Type != nil && a.Type.BaseType == types.BaseDataTypeTag {
			err = findTagNamespace(d, s, a)
			if err != nil {
				return err
			}
		}
		pc.entitiesByElement[s] = append(pc.entitiesByElement[s], a)
	}
	return nil
}

func findAnonymousType(doc *Doc, s *asciidoc.Section, field *matter.Field) error {
	if field.Type == nil {
		return nil
	}
	if field.Type.IsEnum() {
		return findAnonymousEnum(doc, s, field)
	}
	if field.Type.IsMap() {
		return findAnonymousBitmap(doc, s, field)
	}
	return nil
}

func findAnonymousEnum(doc *Doc, s *asciidoc.Section, field *matter.Field) error {
	slog.Debug("possible anonymous enum", "name", field.Name, "type", field.Type)
	ti, err := parseFirstTable(doc, s)
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
	ae := matter.NewAnonymousEnum(s, field)
	ae.Type = field.Type
	for row := range ti.ContentRows() {
		ev := matter.NewEnumValue(s, ae)
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
		ae.Values = append(ae.Values, ev)
	}
	if len(ae.Values) > 0 {
		field.AnonymousType = ae
	}
	return nil
}

func findAnonymousBitmap(doc *Doc, s *asciidoc.Section, field *matter.Field) error {
	slog.Debug("possible anonymous enum", "name", field.Name, "type", field.Type)
	ti, err := parseFirstTable(doc, s)
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
	bm := matter.NewAnonymousBitmap(s, field)
	bm.Type = field.Type
	for row := range ti.ContentRows() {
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

		bv := matter.NewBitmapBit(s, bit, name, summary, conf)
		bm.Bits = append(bm.Bits, bv)
	}
	if len(bm.Bits) > 0 {
		field.AnonymousType = bm
	}
	return nil
}

func findTagNamespace(doc *Doc, s *asciidoc.Section, field *matter.Field) error {
	var found bool
	parse.Search(doc.Reader(), s, s.Elements, func(ref *asciidoc.CrossReference, parent asciidoc.Parent, index int) parse.SearchShould {
		if doc.anchorId(doc.Reader(), ref, ref, ref.ID) == "ref_StandardNamespaces" {
			label := buildReferenceName(doc.Reader(), ref, ref.Elements)
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
		slog.Warn("Tag field does not specify namespace", slog.String("field", field.Name), log.Element("source", doc.Path, field.Source()))
	}
	return nil
}

type fieldFinder struct {
	entityFinderCommon

	fields matter.FieldSet
}

func newFieldFinder(fields matter.FieldSet, inner entityFinder) *fieldFinder {
	return &fieldFinder{entityFinderCommon: entityFinderCommon{inner: inner}, fields: fields}
}

func (ff *fieldFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	for _, c := range ff.fields {
		if c.Name == identifier && c != ff.identity {
			return c
		}
	}
	if ff.inner != nil {
		return ff.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (ff *fieldFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for _, f := range ff.fields {

			if f == ff.identity {
				continue
			}
			if !yield(f.Name, f) {
				return
			}

		}
	})
	if ff.inner != nil {
		ff.inner.suggestIdentifiers(identifier, suggestions)
	}
}

func validateFields(spec *Specification, parent types.Entity, fields matter.FieldSet) {
	idu := make(idUniqueness[*matter.Field])
	nu := make(nameUniqueness[*matter.Field])
	for _, f := range fields {
		if !f.ID.Valid() {
			slog.Warn("Field has invalid ID", log.Path("source", f), matter.LogEntity("parent", parent), slog.String("fieldName", f.Name))
		}
		fieldId := f.ID.Value()
		idu.check(spec, f.ID, f)
		nu.check(spec, f)
		if fieldId >= 0xFE && fieldId <= 0x100 {
			slog.Warn("Struct is using global field ID", log.Path("source", f), matter.LogEntity("parent", parent), slog.String("fieldName", f.Name), slog.String("fieldId", f.ID.HexString()))
		}

		validateAccess(spec, f, f.Access)
	}
}
