package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toBitmap(reader asciidoc.Reader, d *asciidoc.Document, section *asciidoc.Section, parent types.Entity) (bm *matter.Bitmap, err error) {
	name := text.TrimCaseInsensitiveSuffix(library.SectionName(section), " Type")

	var dt *types.DataType
	dt, err = GetDataType(library, reader, d, section)
	if err != nil {
		return nil, newGenericParseError(section, "error parsing bitmap data type: %v", err)
	}
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeMap8, types.DataTypeRankScalar)
		slog.Warn("Bitmap does not declare its derived data type; assuming map8", log.Element("source", d.Path, section), slog.String("bitmap", name))
	} else if !dt.IsMap() {
		return nil, newGenericParseError(section, "unknown bitmap data type: %s", dt.Name)
	}

	bm = matter.NewBitmap(section, parent)
	bm.Name = name
	bm.Type = dt
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, section)

	if err != nil {
		if err == ErrNoTableFound {
			if !isBaseOrDerivedCluster(parent) {
				slog.Warn("no table found for bitmap", log.Element("source", d.Path, section), slog.String("name", bm.Name))
			}
			err = nil
		} else {
			return nil, newGenericParseError(section, "failed reading bitmap %s: %w", name, err)
		}
	} else {
		for row := range ti.ContentRows() {
			var bit, name, summary string
			var conf conformance.Set
			name, err = ti.ReadValue(reader, row, matter.TableColumnName)
			if err != nil {
				return
			}
			name = matter.StripTypeSuffixes(name)
			summary, err = ti.ReadValue(reader, row, matter.TableColumnSummary, matter.TableColumnDescription)
			if err != nil {
				return
			}
			conf = ti.ReadConformance(reader, row, matter.TableColumnConformance)
			if conf == nil {
				conf = conformance.Set{&conformance.Mandatory{}}
			}
			bit, err = ti.ReadString(reader, row, matter.TableColumnBit)
			if err != nil {
				return
			}
			if len(bit) == 0 {
				bit, err = ti.ReadString(reader, row, matter.TableColumnValue)
				if err != nil {
					return
				}
			}
			if len(name) == 0 && len(summary) > 0 {
				name = matter.Case(summary)
			}
			bv := matter.NewBitmapBit(section, bm, bit, CanonicalName(name), summary, conf)
			bm.AddBit(bv)
		}
	}
	library.addEntity(section, bm)
	bm.Name = CanonicalName(bm.Name)
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
	nu := make(nameUniqueness[matter.Bit])
	cv := make(conformanceValidation)
	for _, b := range en.Bits {
		nu.check(spec, b)
		cv.add(b, b.Conformance())
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
	cv.check(spec)
}
