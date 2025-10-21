package errata

import "github.com/project-chip/alchemy/internal/yaml"

type DiscoSection struct {
	Skip DiscoPurpose `yaml:"skip,omitempty"`
}

type Disco struct {
	Sections map[string]DiscoSection `yaml:"sections,omitempty"`
}

func (d *Disco) IgnoreSection(sectionName string, purpose DiscoPurpose) bool {
	if d == nil {
		return false
	}
	if d.Sections == nil {
		return false
	}
	if p, ok := d.Sections[sectionName]; ok {
		return (p.Skip & purpose) != DiscoPurposeNone
	}
	return false
}

type DiscoPurpose uint64

const (
	DiscoPurposeNone                        DiscoPurpose = 0
	DiscoPurposeTableAccess                              = 1 << (iota - 1)
	DiscoPurposeTableConformance                         = 1 << (iota - 1)
	DiscoPurposeTableConstraint                          = 1 << (iota - 1)
	DiscoPurposeTableLinkIndexes                         = 1 << (iota - 1)
	DiscoPurposeTableRenameHeaders                       = 1 << (iota - 1)
	DiscoPurposeTableAddMissingColumns                   = 1 << (iota - 1)
	DiscoPurposeTableReorderColumns                      = 1 << (iota - 1)
	DiscoPurposeDataTypeRename                           = 1 << (iota - 1)
	DiscoPurposeDataTypeAppendSuffix                     = 1 << (iota - 1)
	DiscoPurposeDataTypeBitmapFixRange                   = 1 << (iota - 1)
	DiscoPurposeDataTypeCommandFixDirection              = 1 << (iota - 1)
	DiscoPurposeDataTypePromoteInline                    = 1 << (iota - 1)
	DiscoPurposeNormalizeAnchor                          = 1 << (iota - 1)
	DiscoPurposeTableQuality                             = 1 << (iota - 1)

	DiscoPurposeAll DiscoPurpose = DiscoPurposeTableAccess | DiscoPurposeTableConformance | DiscoPurposeTableConstraint | DiscoPurposeTableLinkIndexes | DiscoPurposeTableQuality | DiscoPurposeTableRenameHeaders | DiscoPurposeTableAddMissingColumns | DiscoPurposeTableReorderColumns | DiscoPurposeDataTypeAppendSuffix | DiscoPurposeDataTypeRename
)

var discoPurposes = map[string]DiscoPurpose{
	"table-access":                    DiscoPurposeTableAccess,
	"table-conformance":               DiscoPurposeTableConformance,
	"table-constraint":                DiscoPurposeTableConstraint,
	"table-quality":                   DiscoPurposeTableQuality,
	"table-link-indexes":              DiscoPurposeTableLinkIndexes,
	"table-rename-headers":            DiscoPurposeTableRenameHeaders,
	"table-add-missing-columns":       DiscoPurposeTableRenameHeaders,
	"table-reorder-columns":           DiscoPurposeTableReorderColumns,
	"data-type-rename":                DiscoPurposeDataTypeRename,
	"data-type-append-suffix":         DiscoPurposeDataTypeAppendSuffix,
	"data-type-bitmap-fix-range":      DiscoPurposeDataTypeBitmapFixRange,
	"data-type-command-fix-direction": DiscoPurposeDataTypeCommandFixDirection,
	"data-type-promote-inline":        DiscoPurposeDataTypePromoteInline,
	"normalize-anchor":                DiscoPurposeNormalizeAnchor,
	"all":                             DiscoPurposeAll,
}

func (i DiscoPurpose) Has(o DiscoPurpose) bool {
	return (i & o) == o
}

func (i DiscoPurpose) HasAny(o DiscoPurpose) bool {
	return (i & o) != 0
}

func (i DiscoPurpose) MarshalYAML() ([]byte, error) {
	return yaml.MarshalBitmap(discoPurposes, i, DiscoPurposeAll)
}

func (i *DiscoPurpose) UnmarshalYAML(b []byte) error {
	return yaml.UnmarshalBitmap(discoPurposes, i, b)
}
