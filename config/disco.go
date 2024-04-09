package config

type DiscoOptions struct {
	LinkIndexTables          bool `json:"linkIndexTables,omitempty"`
	AddMissingColumns        bool `json:"addMissingColumns,omitempty"`
	ReorderColumns           bool `json:"reorderColumns,omitempty"`
	RenameTableHeaders       bool `json:"renameTableHeaders,omitempty"`
	FormatAccess             bool `json:"formatAccess,omitempty"`
	PromoteDataTypes         bool `json:"promoteDataTypes,omitempty"`
	ReorderSections          bool `json:"reorderSections,omitempty"`
	NormalizeTableOptions    bool `json:"normalizeTableOptions,omitempty"`
	FixCommandDirection      bool `json:"fixCommandDirection,omitempty"`
	AppendSubsectionTypes    bool `json:"appendSubsectionTypes,omitempty"`
	UppercaseHex             bool `json:"uppercaseHex,omitempty"`
	AddSpaceAfterPunctuation bool `json:"addSpaceAfterPunctuation,omitempty"`
	RemoveExtraSpaces        bool `json:"removeExtraSpaces,omitempty"`
}

type DiscoSettings map[string]*DiscoOptions
