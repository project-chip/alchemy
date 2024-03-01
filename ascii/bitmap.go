package ascii

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toBitmap(d *Doc, entityMap map[types.WithAttributes][]mattertypes.Entity) (bm *matter.Bitmap, err error) {
	name := strings.TrimSuffix(s.Name, " Type")

	dt := s.GetDataType()
	if dt == nil {
		dt = mattertypes.NewDataType("map8", false)
	}

	if !dt.IsMap() {
		return nil, fmt.Errorf("unknown bitmap data type: %s", dt.Name)
	}

	bm = &matter.Bitmap{
		Name: name,
		Type: dt,
	}
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)

	if err != nil {
		if err == ErrNoTableFound {
			slog.Warn("no table found for bitmap", slog.String("name", bm.Name))
			return bm, nil
		}
		return nil, fmt.Errorf("failed reading bitmap %s: %w", name, err)
	}

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		var bit, name, summary string
		var conf conformance.Set
		name, err = readRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		name = StripTypeSuffixes(name)
		summary, err = readRowAsciiDocString(row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conf = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if conf == nil {
			conf = conformance.Set{&conformance.Mandatory{}}
		}
		bit, err = readRowAsciiDocString(row, columnMap, matter.TableColumnBit)
		if err != nil {
			return
		}
		if len(bit) == 0 {
			bit, err = readRowAsciiDocString(row, columnMap, matter.TableColumnValue)
			if err != nil {
				return
			}
		}
		if len(name) == 0 && len(summary) > 0 {
			name = matter.Case(summary)
		}
		bv := matter.NewBitmapBit(bit, name, summary, conf)
		bm.Bits = append(bm.Bits, bv)
	}
	entityMap[s.Base] = append(entityMap[s.Base], bm)
	return
}
