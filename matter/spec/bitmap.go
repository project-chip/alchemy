package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toBitmap(d *Doc, pc *parseContext, parent types.Entity) (bm *matter.Bitmap, err error) {
	name := text.TrimCaseInsensitiveSuffix(s.Name, " Type")

	dt := s.GetDataType()
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeMap8, false)
		slog.Warn("Bitmap does not declare its derived data type; assuming map8", log.Element("source", d.Path, s.Base), slog.String("bitmap", name))
	} else if !dt.IsMap() {
		return nil, newGenericParseError(s.Base, "unknown bitmap data type: %s", dt.Name)
	}

	bm = matter.NewBitmap(s.Base, parent)
	bm.Name = name
	bm.Type = dt
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)

	if err != nil {
		if err == ErrNoTableFound {
			slog.Warn("no table found for bitmap", log.Element("source", s.Doc.Path, s.Base), slog.String("name", bm.Name))
			err = nil
		} else {
			return nil, newGenericParseError(s.Base, "failed reading bitmap %s: %w", name, err)
		}
	} else {
		for row := range ti.Body() {
			var bit, name, summary string
			var conf conformance.Set
			name, err = ti.ReadValue(row, matter.TableColumnName)
			if err != nil {
				return
			}
			name = matter.StripTypeSuffixes(name)
			summary, err = ti.ReadValue(row, matter.TableColumnSummary, matter.TableColumnDescription)
			if err != nil {
				return
			}
			conf = ti.ReadConformance(row, matter.TableColumnConformance)
			if conf == nil {
				conf = conformance.Set{&conformance.Mandatory{}}
			}
			bit, err = ti.ReadString(row, matter.TableColumnBit)
			if err != nil {
				return
			}
			if len(bit) == 0 {
				bit, err = ti.ReadString(row, matter.TableColumnValue)
				if err != nil {
					return
				}
			}
			if len(name) == 0 && len(summary) > 0 {
				name = matter.Case(summary)
			}
			bv := matter.NewBitmapBit(s.Base, bit, CanonicalName(name), summary, conf)
			bm.AddBit(bv)
		}
	}

	bm.Name = CanonicalName(bm.Name)
	pc.addEntity(bm, s.Base)
	return
}

type bitmapFinder struct {
	entityFinderCommon

	bitmap *matter.Bitmap
}

func newBitmapFinder(bm *matter.Bitmap, inner entityFinder) *bitmapFinder {
	return &bitmapFinder{entityFinderCommon: entityFinderCommon{inner: inner}, bitmap: bm}
}

func (bf *bitmapFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	for _, bmv := range bf.bitmap.Bits {
		if bmv.Name() == identifier && bmv != bf.identity {
			return bmv
		}
	}
	if bf.inner != nil {
		return bf.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (bf *bitmapFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for _, bmv := range bf.bitmap.Bits {

			if bmv == bf.identity {
				continue
			}
			if !yield(bmv.Name(), bmv) {
				return
			}

		}
	})
	if bf.inner != nil {
		bf.inner.suggestIdentifiers(identifier, suggestions)
	}
	return
}

func validateBitmaps(spec *Specification) {
	for c := range spec.Clusters {
		for _, bm := range c.Bitmaps {
			validateBitmap(spec, bm)
		}
		if c.Features != nil {
			validateBitmap(spec, &c.Features.Bitmap)
		}
	}
	for _, bm := range types.FilterSet[*matter.Bitmap](spec.GlobalObjects) {
		validateBitmap(spec, bm)
	}
}

func validateBitmap(spec *Specification, en *matter.Bitmap) {
	bits := make(map[uint64]matter.Bit)
	for _, b := range en.Bits {
		from, to, err := b.Bits()
		if err != nil {
			slog.Warn("Unable to determine range of bitmap values", log.Path("source", b), matter.LogEntity("parent", en), slog.String("name", b.Name()), slog.Any("error", err))
			continue
		}
		for i := from; i <= to; i++ {
			existing, ok := bits[i]
			if ok {
				slog.Error("Overlapping bitmap bit range", log.Path("source", b), matter.LogEntity("parent", en), slog.Uint64("bitFrom", from), slog.Uint64("bitTo", to), slog.Uint64("bitOverlap", i), slog.String("bitName", b.Name()), slog.String("previousBitName", existing.Name()))
				spec.addError(&DuplicateEntityIDError{Entity: b, Previous: existing})
			}
			bits[i] = b
		}
	}

}
