package testplan

import (
	"unicode"
	"unicode/utf8"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func ifHasQualityHelper(q matter.Quality, quality string, options *raymond.Options) raymond.SafeString {
	var hasQuality bool
	switch quality {
	case "nullable":
		hasQuality = q.Has(matter.QualityNullable)
	case "nonVolatile":
		hasQuality = q.Has(matter.QualityNonVolatile)
	case "fixed":
		hasQuality = q.Has(matter.QualityFixed)
	case "scene":
		hasQuality = q.Has(matter.QualityScene)
	case "reportable":
		hasQuality = q.Has(matter.QualityReportable)
	case "changedOmitted":
		hasQuality = q.Has(matter.QualityChangedOmitted)
	case "diagnostics":
		hasQuality = q.Has(matter.QualityDiagnostics)
	case "singleton":
		hasQuality = q.Has(matter.QualitySingleton)
	case "largeMessage":
		hasQuality = q.Has(matter.QualityLargeMessage)
	case "sourceAttribution":
		hasQuality = q.Has(matter.QualitySourceAttribution)
	case "atomicWrite":
		hasQuality = q.Has(matter.QualityAtomicWrite)
	case "quieterReporting":
		hasQuality = q.Has(matter.QualityQuieterReporting)
	}
	if hasQuality {
		return raymond.SafeString(options.Fn())
	}
	return raymond.SafeString(options.Inverse())
}

func dataTypeHelper(dt types.DataType) raymond.SafeString {
	switch dt.BaseType {
	case types.BaseDataTypeCustom:
		if dt.Entity != nil {
			switch entity := dt.Entity.(type) {
			case *matter.Enum:
				return raymond.SafeString(entity.Type.Name)
			case *matter.Bitmap:
				return raymond.SafeString(entity.Type.Name)
			}
		}
		return raymond.SafeString(dt.Name)
	case types.BaseDataTypeVoltage, types.BaseDataTypePower, types.BaseDataTypeEnergy, types.BaseDataTypeAmperage:
		return "int64"
	default:
		return raymond.SafeString(dt.Name)
	}
}

func dataTypeArticleHelper(dt types.DataType) raymond.SafeString {
	firstLetter, _ := utf8.DecodeRuneInString(string(dataTypeHelper(dt)))
	switch unicode.ToLower(firstLetter) {
	case 'a', 'e', 'i', 'o', 'u': // Not perfect, but not importing a dictionary for this
		return raymond.SafeString("an")
	default:
		return raymond.SafeString("a")
	}
}

func ifDataTypeIsArrayHelper(dt types.DataType, options *raymond.Options) raymond.SafeString {
	if dt.IsArray() {
		return raymond.SafeString(options.Fn())
	}
	return raymond.SafeString(options.Inverse())
}
