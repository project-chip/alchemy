package generate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/iancoleman/orderedmap"
)

func patchZapJson(sdkRoot string, files []string) error {
	err := patchZapJsonFile(sdkRoot, files, "src/app/zap-templates/zcl/zcl.json")
	if err != nil {
		return err
	}
	return patchZapJsonFile(sdkRoot, files, "src/app/zap-templates/zcl/zcl-with-test-extensions.json")
}

func patchZapJsonFile(sdkRoot string, files []string, file string) error {
	zclJSONPath := path.Join(sdkRoot, file)
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
	fileMap := make(map[string]struct{})
	for _, file := range files {
		fileMap[file] = struct{}{}
	}
	for _, i := range is {
		if s, ok := i.(string); ok {
			xmls = append(xmls, s)
			delete(fileMap, s)
		}
	}

	xmls = mergeLines(xmls, fileMap, 2)

	slices.Compact(xmls)
	o.Set("xmlFile", xmls)

	zclJSONBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(zclJSONPath, []byte(zclJSONBytes), os.ModeAppend|0644)
}
