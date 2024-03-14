package generate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/iancoleman/orderedmap"
)

func patchZapJsonFile(sdkRoot string, file string, files []string) (zclJSONPath string, zclJSONBytes []byte, err error) {
	zclJSONPath = path.Join(sdkRoot, file)
	zclJSONBytes, err = os.ReadFile(zclJSONPath)
	if err != nil {
		return
	}

	o := orderedmap.New()
	err = json.Unmarshal(zclJSONBytes, &o)
	if err != nil {
		return
	}
	val, ok := o.Get("xmlFile")
	if !ok {
		err = fmt.Errorf("missing xmlFile element in %s", zclJSONPath)
		return
	}
	is, ok := val.([]interface{})
	if !ok {
		err = fmt.Errorf("xmlFile element in %s is not array", zclJSONPath)
		return
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
		err = fmt.Errorf("error marshaling %s: %w", zclJSONPath, err)
		return
	}
	return
}
