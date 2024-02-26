package config

type DiscoSettings struct {
	LinkIndexTables          bool
	AddMissingColumns        bool
	ReorderColumns           bool
	RenameTableHeaders       bool
	FormatAccess             bool
	PromoteDataTypes         bool
	ReorderSections          bool
	NormalizeTableOptions    bool
	FixCommandDirection      bool
	AppendSubsectionTypes    bool
	UppercaseHex             bool
	AddSpaceAfterPunctuation bool
	RemoveExtraSpaces        bool
}
