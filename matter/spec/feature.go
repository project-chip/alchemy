package spec

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toFeatures(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (features *matter.Features, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading features: %w", err)

	}
	features = &matter.Features{
		Bitmap: matter.Bitmap{
			Name: "Feature",
			Type: types.NewDataType(types.BaseDataTypeMap32, false),
		},
	}
	featureMap := make(map[string]*matter.Feature)
	for row := range ti.Body() {
		var bit, code, name, summary string
		var conf conformance.Set
		bit, err = ti.ReadString(row, matter.TableColumnBit)
		if err != nil {
			return
		}
		if len(bit) == 0 {
			bit, err = ti.ReadString(row, matter.TableColumnID)
			if err != nil {
				return
			}
		}

		name, err = ti.ReadValue(row, matter.TableColumnFeature, matter.TableColumnName)
		if err != nil {
			return
		}
		name = matter.StripTypeSuffixes(name)

		code, err = ti.ReadString(row, matter.TableColumnCode)
		if err != nil {
			return
		}
		summary, err = ti.ReadValue(row, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conf = ti.ReadConformance(row, matter.TableColumnConformance)
		if conf == nil {
			conf = conformance.Set{&conformance.Optional{}}
		}
		f := matter.NewFeature(bit, name, code, summary, conf)
		features.Bits = append(features.Bits, f)
		featureMap[name] = f
	}

	for _, s := range parse.Skim[*Section](s.Elements()) {
		switch s.SecType {
		case matter.SectionFeature:

			name := text.TrimCaseInsensitiveSuffix(s.Name, " Feature")
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
