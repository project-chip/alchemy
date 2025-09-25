package spec

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func toFeatures(d *Doc, s *asciidoc.Section, pc *parseContext) (features *matter.Features, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
			return
		}
		return nil, fmt.Errorf("failed reading features: %w", err)

	}
	features = matter.NewFeatures(s, nil)
	featureMap := make(map[string]*matter.Feature)
	for row := range ti.ContentRows() {
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
		f := matter.NewFeature(row, bit, name, code, summary, conf)
		features.AddFeatureBit(f)
		featureMap[name] = f
	}

	for s := range parse.Skim[*asciidoc.Section](d.Reader(), s, s.Children()) {
		switch d.SectionType(s) {
		case matter.SectionFeature:

			name := text.TrimCaseInsensitiveSuffix(d.SectionName(s), " Feature")
			a, ok := featureMap[name]
			if !ok {
				slog.Debug("unknown feature", "feature", name)
				continue
			}

			pc.entitiesByElement[s] = append(pc.entitiesByElement[s], a)
		}
	}
	return
}

type featureFinder struct {
	entityFinderCommon

	features *matter.Features
}

func newFeatureFinder(features *matter.Features, inner entityFinder) *featureFinder {
	return &featureFinder{entityFinderCommon: entityFinderCommon{inner: inner}, features: features}
}

func (ff *featureFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {

	if ff.features != nil {
		for f := range ff.features.FeatureBits() {
			if f.Code == identifier {
				return f
			}
		}
	}
	if ff.inner != nil {
		return ff.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (ff *featureFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	if ff.features != nil {
		suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
			for f := range ff.features.FeatureBits() {
				if f == ff.identity {
					continue
				}
				if !yield(f.Code, f) {
					return
				}

			}
		})
	}
	if ff.inner != nil {
		ff.inner.suggestIdentifiers(identifier, suggestions)
	}
}
