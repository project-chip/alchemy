package matter

import (
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type Quality uint32

const (
	QualityNone              Quality = 0
	QualityNullable                  = 1 << (iota - 1)
	QualityNonVolatile               = 1 << (iota - 1)
	QualityFixed                     = 1 << (iota - 1)
	QualityScene                     = 1 << (iota - 1)
	QualityReportable                = 1 << (iota - 1)
	QualityChangedOmitted            = 1 << (iota - 1)
	QualityDiagnostics               = 1 << (iota - 1)
	QualitySingleton                 = 1 << (iota - 1)
	QualityLargeMessage              = 1 << (iota - 1)
	QualitySourceAttribution         = 1 << (iota - 1)
	QualityAtomicWrite               = 1 << (iota - 1)
	QualityQuieterReporting          = 1 << (iota - 1)

	QualityAll = QualityNullable | QualityNonVolatile | QualityFixed | QualityScene | QualityReportable | QualityChangedOmitted | QualityDiagnostics | QualitySingleton | QualityLargeMessage | QualitySourceAttribution | QualityAtomicWrite | QualityQuieterReporting
)

var identifierQualities = map[rune]Quality{
	'X': QualityNullable,
	'N': QualityNonVolatile,
	'F': QualityFixed,
	'S': QualityScene,
	'P': QualityReportable,
	'C': QualityChangedOmitted,
	'K': QualityDiagnostics,
	'I': QualitySingleton,
	'L': QualityLargeMessage,
	'A': QualitySourceAttribution,
	'T': QualityAtomicWrite,
	'Q': QualityQuieterReporting,
}

var AllowedQualities = map[types.EntityType]Quality{
	types.EntityTypeAttribute:    QualitySourceAttribution | QualityChangedOmitted | QualityFixed | QualityNonVolatile | QualityQuieterReporting | QualityScene | QualityAtomicWrite | QualityNullable,
	types.EntityTypeCommand:      QualityLargeMessage,
	types.EntityTypeCommandField: QualitySourceAttribution | QualityChangedOmitted | QualityFixed | QualityNonVolatile | QualityQuieterReporting | QualityScene | QualityAtomicWrite | QualityNullable,
	types.EntityTypeStructField:  QualitySourceAttribution | QualityChangedOmitted | QualityFixed | QualityNonVolatile | QualityQuieterReporting | QualityScene | QualityAtomicWrite | QualityNullable,
	types.EntityTypeEvent:        QualitySourceAttribution,
	types.EntityTypeEventField:   QualitySourceAttribution | QualityChangedOmitted | QualityFixed | QualityNonVolatile | QualityQuieterReporting | QualityScene | QualityAtomicWrite | QualityNullable,
	types.EntityTypeCluster:      QualitySingleton | QualityDiagnostics,
}

var qualityIdentifiers map[Quality]rune

func init() {
	qualityIdentifiers = make(map[Quality]rune, len(identifierQualities))
	for i, q := range identifierQualities {
		qualityIdentifiers[q] = i
	}
}

func ParseQuality(s string) Quality {
	var q Quality
	for _, r := range s {
		if qi, ok := identifierQualities[r]; ok {
			q |= qi
		}
	}
	return q
}

func (q Quality) Has(o Quality) bool {
	return (q & o) == o
}

func (q Quality) Any(o Quality) bool {
	return (q & o) != QualityNone
}

func (q Quality) String() string {
	var s strings.Builder
	for tq, i := range qualityIdentifiers {
		if (q & tq) == tq {
			s.WriteRune(i)
		}
	}
	return s.String()
}

func (q *Quality) Inherit(oq Quality) {
	*q = (*q | oq)
}
