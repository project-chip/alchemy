package matter

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
)

func ParseQuality(s string) Quality {
	var q Quality
	for _, r := range s {
		switch r {
		case 'X':
			q |= QualityNullable
		case 'N':
			q |= QualityNonVolatile
		case 'F':
			q |= QualityFixed
		case 'S':
			q |= QualityScene
		case 'P':
			q |= QualityReportable
		case 'C':
			q |= QualityChangedOmitted
		case 'K':
			q |= QualityDiagnostics
		case 'I':
			q |= QualitySingleton
		case 'L':
			q |= QualityLargeMessage
		case 'A':
			q |= QualitySourceAttribution
		}
	}
	return q
}

func (q Quality) Has(o Quality) bool {
	return (q & o) == o
}
