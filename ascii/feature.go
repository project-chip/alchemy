package ascii

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
)

func (s *Section) toFeatures(d *Doc, entityMap map[types.WithAttributes][]mattertypes.Entity) (features *matter.Features, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading features: %w", err)

	}
	features = &matter.Features{
		Bitmap: matter.Bitmap{
			Name: "Feature",
			Type: mattertypes.NewDataType("map32", false),
		},
	}
	featureMap := make(map[string]*matter.Feature)
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		var bit, code, name, summary string
		var conf conformance.Set
		bit, err = readRowAsciiDocString(row, columnMap, matter.TableColumnBit)
		if err != nil {
			return
		}
		if len(bit) == 0 {
			bit, err = readRowAsciiDocString(row, columnMap, matter.TableColumnID)
			if err != nil {
				return
			}
		}

		name, err = readRowValue(d, row, columnMap, matter.TableColumnFeature, matter.TableColumnName)
		if err != nil {
			return
		}
		name = StripTypeSuffixes(name)

		code, err = readRowAsciiDocString(row, columnMap, matter.TableColumnCode)
		if err != nil {
			return
		}
		summary, err = readRowAsciiDocString(row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conf = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if conf == nil {
			conf = conformance.Set{&conformance.Optional{}}
		}
		f := matter.NewFeature(bit, name, code, summary, conf)
		features.Bits = append(features.Bits, f)
		featureMap[name] = f
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionFeature:

			name := strings.TrimSuffix(s.Name, " Attribute")
			a, ok := featureMap[name]
			if !ok {
				slog.Debug("unknown feature", "feature", name)
				continue
			}

			entityMap[s.Base] = append(entityMap[s.Base], a)
		}
	}
	return
}
