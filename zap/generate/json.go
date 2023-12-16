package generate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"slices"
	"sort"

	"github.com/iancoleman/orderedmap"
)

func patchZapJson(zclRoot string, files []string) error {
	err := patchZapJsonFile(zclRoot, files, "src/app/zap-templates/zcl/zcl.json")
	if err != nil {
		return err
	}
	return patchZapJsonFile(zclRoot, files, "src/app/zap-templates/zcl/zcl-with-test-extensions.json")
}

func patchZapJsonFile(zclRoot string, files []string, file string) error {
	zclJSONPath := path.Join(zclRoot, file)
	zclJSONBytes, err := os.ReadFile(zclJSONPath)
	if err != nil {
		return err
	}

	o := orderedmap.New()
	err = json.Unmarshal(zclJSONBytes, &o)
	if err != nil {
		return err
	}
	val, ok := o.Get("xmlFile")
	if !ok {
		return fmt.Errorf("no xmlField field in zcl.json")
	}
	is, ok := val.([]interface{})
	if !ok {
		return fmt.Errorf("xmlField not a string array in zcl.json; %T", val)
	}
	xmls := make([]string, 0, len(is)+len(files))
	for _, i := range is {
		if s, ok := i.(string); ok {
			xmls = append(xmls, s)
		}
	}
	xmls = append(xmls, files...)
	sort.Slice(xmls, func(i, j int) bool {
		a := xmls[i]
		b := xmls[j]
		if a == "access-control-definitions.xml" {
			return true
		}
		if b == "access-control-definitions.xml" {
			return false
		}
		if a < b {
			return true
		}
		return false
	})
	slices.Compact(xmls)
	o.Set("xmlFile", xmls)

	zclJSONBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(zclJSONPath, []byte(zclJSONBytes), os.ModeAppend|0644)
}
