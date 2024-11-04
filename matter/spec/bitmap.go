package spec

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toBitmap(d *Doc, pc *parseContext) (bm *matter.Bitmap, err error) {
	name := text.TrimCaseInsensitiveSuffix(s.Name, " Type")

	dt := s.GetDataType()
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeMap8, false)
		slog.Warn("Bitmap does not declare its derived data type; assuming map8", log.Element("path", d.Path, s.Base), slog.String("bitmap", name))
	} else if !dt.IsMap() {
		return nil, fmt.Errorf("unknown bitmap data type: %s", dt.Name)
	}

	bm = matter.NewBitmap(s.Base)
	bm.Name = name
	bm.Type = dt
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)

	if err != nil {
		if err == ErrNoTableFound {
			slog.Warn("no table found for bitmap", log.Element("path", s.Doc.Path, s.Base), slog.String("name", bm.Name))
			return bm, nil
		}
		return nil, fmt.Errorf("failed reading bitmap %s: %w", name, err)
	}

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
		bm.Bits = append(bm.Bits, bv)
	}
	pc.orderedEntities = append(pc.orderedEntities, bm)
	pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], bm)
	bm.Name = CanonicalName(bm.Name)
	return
}
