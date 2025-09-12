package errata

import "github.com/project-chip/alchemy/internal/yaml"

type SpecSection struct {
	Skip SpecPurpose `yaml:"skip,omitempty"`
}

type Spec struct {
	UtilityInclude bool                   `yaml:"utility-include,omitempty"`
	Sections       map[string]SpecSection `yaml:"sections,omitempty"`
	DocRoot        bool                   `yaml:"doc-root,omitempty"`
	Domain         string                 `yaml:"domain,omitempty"`
}

func (spec *Spec) IgnoreSection(sectionName string, purpose SpecPurpose) bool {
	if spec == nil {
		return false
	}
	if spec.Sections == nil {
		return false
	}
	if p, ok := spec.Sections[sectionName]; ok {
		return (p.Skip & purpose) != SpecPurposeNone
	}
	return false
}

type SpecPurpose uint64

const (
	SpecPurposeNone             SpecPurpose = 0
	SpecPurposeDataTypesBitmap              = 1 << (iota - 1)
	SpecPurposeDataTypesEnum                = 1 << (iota - 1)
	SpecPurposeDataTypesStruct              = 1 << (iota - 1)
	SpecPurposeCluster                      = 1 << (iota - 1)
	SpecPurposeDeviceType                   = 1 << (iota - 1)
	SpecPurposeCommandArguments             = 1 << (iota - 1)
	SpecPurposeDataTypesDef                 = 1 << (iota - 1)
	SpecPurposeFeatures                     = 1 << (iota - 1)

	SpecPurposeDataTypes SpecPurpose = SpecPurposeDataTypesBitmap | SpecPurposeDataTypesEnum | SpecPurposeDataTypesStruct | SpecPurposeDataTypesDef
	SpecPurposeAll       SpecPurpose = SpecPurposeDataTypes | SpecPurposeCluster | SpecPurposeDeviceType | SpecPurposeCommandArguments | SpecPurposeFeatures
)

var specPurposes = map[string]SpecPurpose{
	"data-types":        SpecPurposeDataTypes,
	"data-types-bitmap": SpecPurposeDataTypesBitmap,
	"data-types-enum":   SpecPurposeDataTypesEnum,
	"data-types-struct": SpecPurposeDataTypesStruct,
	"cluster":           SpecPurposeCluster,
	"device-type":       SpecPurposeDeviceType,
	"command-arguments": SpecPurposeCommandArguments,
	"features":          SpecPurposeFeatures,
	"all":               SpecPurposeAll,
}

func (i SpecPurpose) Has(o SpecPurpose) bool {
	return (i & o) == o
}

func (i SpecPurpose) HasAny(o SpecPurpose) bool {
	return (i & o) != 0
}

func (i SpecPurpose) MarshalYAML() ([]byte, error) {
	return yaml.MarshalBitmap(specPurposes, i, SpecPurposeAll)
}

func (i *SpecPurpose) UnmarshalYAML(b []byte) error {
	return yaml.UnmarshalBitmap(specPurposes, i, b)
}
