package errata

type SpecSection struct {
	Skip SpecPurpose `yaml:"skip,omitempty"`
}

type Spec struct {
	UtilityInclude bool                   `yaml:"utility-include,omitempty"`
	Sections       map[string]SpecSection `yaml:"sections,omitempty"`
}

func GetSpec(path string) *Spec {
	e := GetErrata(path)
	return &e.Spec
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

type SpecPurpose uint32

const (
	SpecPurposeNone             SpecPurpose = 0
	SpecPurposeDataTypesBitmap              = 1 << (iota - 1)
	SpecPurposeDataTypesEnum                = 1 << (iota - 1)
	SpecPurposeDataTypesStruct              = 1 << (iota - 1)
	SpecPurposeCluster                      = 1 << (iota - 1)
	SpecPurposeDeviceType                   = 1 << (iota - 1)
	SpecPurposeCommandArguments             = 1 << (iota - 1)

	SpecPurposeDataTypes SpecPurpose = SpecPurposeDataTypesBitmap | SpecPurposeDataTypesEnum | SpecPurposeDataTypesStruct
	SpecPurposeAll       SpecPurpose = SpecPurposeDataTypesBitmap | SpecPurposeDataTypesEnum | SpecPurposeDataTypesStruct | SpecPurposeCluster | SpecPurposeDeviceType | SpecPurposeCommandArguments
)

var specPurposes = map[string]SpecPurpose{
	"data-types":        SpecPurposeDataTypes,
	"data-types-bitmap": SpecPurposeDataTypesBitmap,
	"data-types-enum":   SpecPurposeDataTypesEnum,
	"data-types-struct": SpecPurposeDataTypesStruct,
	"cluster":           SpecPurposeCluster,
	"device-type":       SpecPurposeDeviceType,
	"command-arguments": SpecPurposeCommandArguments,
	"all":               SpecPurposeAll,
}

func (i SpecPurpose) Has(o SpecPurpose) bool {
	return (i & o) == o
}

func (i SpecPurpose) HasAny(o SpecPurpose) bool {
	return (i & o) != 0
}

func (i SpecPurpose) MarshalYAML() ([]byte, error) {
	return marshalYamlBitmap(specPurposes, i, SpecPurposeAll)
}

func (i *SpecPurpose) UnmarshalYAML(b []byte) error {
	return unmarshalYamlBitmap(specPurposes, i, b)
}