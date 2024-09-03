package errata

import (
	"fmt"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"
)

type yamlBitmap[BitmapType ~uint32 | ~uint64] interface {
	Has(bt BitmapType) bool
}

func marshalYamlBitmap[BitmapType ~uint32 | ~uint64](valueMap map[string]BitmapType, value yamlBitmap[BitmapType], allValue yamlBitmap[BitmapType]) ([]byte, error) {
	if value == allValue {
		return []byte("all"), nil
	}
	var purposes []string
	for k, v := range valueMap {
		if value.Has(v) {
			purposes = append(purposes, k)
		}
	}
	slices.Sort(purposes)
	return yaml.Marshal(purposes)
}

func unmarshalYamlBitmap[BitmapType ~uint32 | ~uint64](valueMap map[string]BitmapType, value *BitmapType, b []byte) error {
	var v []string
	if err := yaml.Unmarshal(b, &v); err != nil {
		return err
	}
	for _, p := range v {
		dp, ok := valueMap[strings.ToLower(strings.TrimSpace(p))]
		if !ok {
			return fmt.Errorf("unknown purpose: %s", strings.TrimSpace(p))
		}
		*value |= dp
	}
	return nil
}
