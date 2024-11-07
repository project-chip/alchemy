package disco

type Option func(options *options)

type options struct {
	linkIndexTables               bool
	addMissingColumns             bool
	reorderColumns                bool
	renameTableHeaders            bool
	formatAccess                  bool
	promoteDataTypes              bool
	reorderSections               bool
	normalizeTableOptions         bool
	fixCommandDirection           bool
	appendSubsectionTypes         bool
	uppercaseHex                  bool
	addSpaceAfterPunctuation      bool
	removeExtraSpaces             bool
	normalizeFeatureNames         bool
	disambiguateConformanceChoice bool
	normalizeAnchors              bool
	removeMandatoryDefaults       bool
}

var defaultOptions = options{
	linkIndexTables:               false,
	addMissingColumns:             true,
	reorderColumns:                true,
	renameTableHeaders:            true,
	formatAccess:                  true,
	promoteDataTypes:              true,
	reorderSections:               true,
	normalizeTableOptions:         true,
	fixCommandDirection:           true,
	appendSubsectionTypes:         true,
	uppercaseHex:                  true,
	addSpaceAfterPunctuation:      true,
	removeExtraSpaces:             true,
	normalizeFeatureNames:         true,
	disambiguateConformanceChoice: false,
	normalizeAnchors:              false,
	removeMandatoryDefaults:       false,
}

func LinkIndexTables(link bool) Option {
	return func(options *options) {
		options.linkIndexTables = link
	}
}

func AddMissingColumns(add bool) Option {
	return func(options *options) {
		options.addMissingColumns = add
	}
}

func ReorderColumns(reorder bool) Option {
	return func(options *options) {
		options.reorderColumns = reorder
	}
}

func RenameTableHeaders(rename bool) Option {
	return func(options *options) {
		options.renameTableHeaders = rename
	}
}

func FormatAccess(format bool) Option {
	return func(options *options) {
		options.formatAccess = format
	}
}

func PromoteDataTypes(promote bool) Option {
	return func(options *options) {
		options.promoteDataTypes = promote
	}
}

func ReorderSections(reorder bool) Option {
	return func(options *options) {
		options.reorderSections = reorder
	}
}

func FixCommandDirection(add bool) Option {
	return func(options *options) {
		options.fixCommandDirection = add
	}
}

func AppendSubsectionTypes(add bool) Option {
	return func(options *options) {
		options.appendSubsectionTypes = add
	}
}

func UppercaseHex(add bool) Option {
	return func(options *options) {
		options.uppercaseHex = add
	}
}

func AddSpaceAfterPunctuation(add bool) Option {
	return func(options *options) {
		options.addSpaceAfterPunctuation = add
	}
}

func RemoveExtraSpaces(add bool) Option {
	return func(options *options) {
		options.removeExtraSpaces = add
	}
}

func NormalizeTableOptions(add bool) Option {
	return func(options *options) {
		options.normalizeTableOptions = add
	}
}

func NormalizeFeatureNames(add bool) Option {
	return func(options *options) {
		options.normalizeFeatureNames = add
	}
}

func DisambiguateConformanceChoice(add bool) Option {
	return func(options *options) {
		options.disambiguateConformanceChoice = add
	}
}

func NormalizeAnchors(add bool) Option {
	return func(options *options) {
		options.normalizeAnchors = add
	}
}

func RemoveMandatoryDefaults(add bool) Option {
	return func(options *options) {
		options.removeMandatoryDefaults = add
	}
}
