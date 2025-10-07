package disco

type Option func(options *DiscoOptions)

type DiscoOptions struct {
	LinkIndexTables               bool `default:"false" aliases:"linkIndexTables" help:"link index tables to child sections" group:"Discoballing:"`
	AddMissingColumns             bool `default:"true" aliases:"addMissingColumns" help:"add standard columns missing from tables" group:"Discoballing:"`
	ReorderColumns                bool `default:"true" aliases:"reorderColumns" help:"rearrange table columns into disco-ball order" group:"Discoballing:"`
	RenameTableHeaders            bool `default:"true" aliases:"renameTableHeaders" help:"rename table headers to disco-ball standard names" group:"Discoballing:"`
	FormatAccess                  bool `default:"true" aliases:"formatAccess" help:"reformat access columns in disco-ball order" group:"Discoballing:"`
	FormatQuality                 bool `default:"true" help:"reformat quality columns in disco-ball order" group:"Discoballing:"`
	PromoteDataTypes              bool `default:"true" aliases:"promoteDataTypes" help:"promote inline data types to Data Types section" group:"Discoballing:"`
	ReorderSections               bool `default:"true" aliases:"reorderSections" help:"reorder sections in disco-ball order" group:"Discoballing:"`
	NormalizeTableOptions         bool `default:"true" aliases:"normalizeTableOptions" help:"remove existing table options and replace with standard disco-ball options" group:"Discoballing:"`
	FixCommandDirection           bool `default:"true" aliases:"fixCommandDirection" help:"normalize command directions" group:"Discoballing:"`
	AppendSubsectionTypes         bool `default:"true" aliases:"appendSubsectionTypes" help:"add missing suffixes to data type sections (e.g. \"Bit\", \"Value\", \"Field\", etc.)" group:"Discoballing:"`
	UppercaseHex                  bool `default:"true" aliases:"uppercaseHex" help:"uppercase hex values" group:"Discoballing:"`
	AddSpaceAfterPunctuation      bool `default:"true" aliases:"addSpaceAfterPunctuation" help:"add missing space after punctuation" group:"Discoballing:"`
	RemoveExtraSpaces             bool `default:"true" aliases:"removeExtraSpaces" help:"remove extraneous spaces" group:"Discoballing:"`
	NormalizeFeatureNames         bool `default:"true" aliases:"normalizeFeatureNames" help:"correct invalid feature names" group:"Discoballing:"`
	DisambiguateConformanceChoice bool `default:"true" aliases:"disambiguateConformanceChoice" help:"ensure conformance choices are only used once per document" group:"Discoballing:"`
	NormalizeAnchors              bool `default:"false" aliases:"normalizeAnchors" help:"rewrite anchors and references without labels" group:"Discoballing:"`
	RemoveMandatoryFallbacks      bool `default:"true" aliases:"removeMandatoryFallbacks" help:"remove fallback values for mandatory fields" group:"Discoballing:"`
	RenameSections                bool `default:"false" help:"rename sections to disco-ball standard names" group:"Discoballing:"`
}

var DefaultOptions = DiscoOptions{
	LinkIndexTables:               false,
	AddMissingColumns:             true,
	ReorderColumns:                true,
	RenameTableHeaders:            true,
	FormatAccess:                  true,
	FormatQuality:                 true,
	PromoteDataTypes:              true,
	ReorderSections:               true,
	NormalizeTableOptions:         true,
	FixCommandDirection:           true,
	AppendSubsectionTypes:         true,
	UppercaseHex:                  true,
	AddSpaceAfterPunctuation:      true,
	RemoveExtraSpaces:             true,
	NormalizeFeatureNames:         true,
	DisambiguateConformanceChoice: true,
	NormalizeAnchors:              false,
	RemoveMandatoryFallbacks:      true,
	RenameSections:                false,
}
