package errata

import (
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
)

type Purpose uint32

const (
	PurposeNone            Purpose = 0
	PurposeDataTypesBitmap         = 1 << (iota - 1)
	PurposeDataTypesEnum           = 1 << (iota - 1)
	PurposeDataTypesStruct         = 1 << (iota - 1)

	PurposeDataTypes = PurposeDataTypesBitmap | PurposeDataTypesEnum | PurposeDataTypesStruct
)

func (i Purpose) Has(o Purpose) bool {
	return (i & o) == o
}

func (i Purpose) HasAny(o Purpose) bool {
	return (i & o) != 0
}

func (i Purpose) MarshalYAML() ([]byte, error) {
	var purposes []string
	if i.Has(PurposeDataTypes) {
		purposes = append(purposes, "data-types")
	} else {
		if i.Has(PurposeDataTypesBitmap) {
			purposes = append(purposes, "data-types-bitmap")
		}
		if i.Has(PurposeDataTypesEnum) {
			purposes = append(purposes, "data-types-enum")
		}
		if i.Has(PurposeDataTypesStruct) {
			purposes = append(purposes, "data-types-struct")
		}
	}

	return []byte(strings.Join(purposes, ", ")), nil
}

func (i *Purpose) UnmarshalYAML(b []byte) error {
	var v string
	if err := yaml.Unmarshal(b, &v); err != nil {
		return err
	}
	parts := strings.Split(v, ",")
	for _, p := range parts {
		switch strings.ToLower(strings.TrimSpace(p)) {
		case "data-types":
			*i |= PurposeDataTypes
		case "data-types-bitmap":
			*i |= PurposeDataTypesBitmap
		case "data-types-enum":
			*i |= PurposeDataTypesEnum
		case "data-types-struct":
			*i |= PurposeDataTypesStruct
		default:
			return fmt.Errorf("unknown errata purpose: %s", p)
		}
	}
	return nil
}
