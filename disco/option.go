package disco

type Option func(b *Ball)

type options struct {
	linkAttributes           bool
	addMissingColumns        bool
	reorderColumns           bool
	renameTableHeaders       bool
	formatAccess             bool
	promoteDataTypes         bool
	reorderSections          bool
	normalizeTableOptions    bool
	fixCommandDirection      bool
	appendSubsectionTypes    bool
	uppercaseHex             bool
	addSpaceAfterPunctuation bool
	removeExtraSpaces        bool
}

var defaultOptions = options{
	linkAttributes:           false,
	addMissingColumns:        true,
	reorderColumns:           true,
	renameTableHeaders:       true,
	formatAccess:             true,
	promoteDataTypes:         true,
	reorderSections:          true,
	normalizeTableOptions:    true,
	fixCommandDirection:      true,
	appendSubsectionTypes:    true,
	uppercaseHex:             true,
	addSpaceAfterPunctuation: true,
	removeExtraSpaces:        true,
}

func LinkAttributes(link bool) Option {
	return func(b *Ball) {
		b.options.linkAttributes = link
	}
}

func AddMissingColumns(add bool) Option {
	return func(b *Ball) {
		b.options.addMissingColumns = add
	}
}

func ReorderColumns(reorder bool) Option {
	return func(b *Ball) {
		b.options.reorderColumns = reorder
	}
}

func RenameTableHeaders(rename bool) Option {
	return func(b *Ball) {
		b.options.renameTableHeaders = rename
	}
}

func FormatAccess(format bool) Option {
	return func(b *Ball) {
		b.options.formatAccess = format
	}
}

func PromoteDataTypes(promote bool) Option {
	return func(b *Ball) {
		b.options.promoteDataTypes = promote
	}
}

func ReorderSections(reorder bool) Option {
	return func(b *Ball) {
		b.options.reorderSections = reorder
	}
}

func FixCommandDirection(add bool) Option {
	return func(b *Ball) {
		b.options.fixCommandDirection = add
	}
}

func AppendSubsectionTypes(add bool) Option {
	return func(b *Ball) {
		b.options.appendSubsectionTypes = add
	}
}

func UppercaseHex(add bool) Option {
	return func(b *Ball) {
		b.options.uppercaseHex = add
	}
}

func AddSpaceAfterPunctuation(add bool) Option {
	return func(b *Ball) {
		b.options.addSpaceAfterPunctuation = add
	}
}

func RemoveExtraSpaces(add bool) Option {
	return func(b *Ball) {
		b.options.removeExtraSpaces = add
	}
}

func NormalizeTableOptions(add bool) Option {
	return func(b *Ball) {
		b.options.normalizeTableOptions = add
	}
}
