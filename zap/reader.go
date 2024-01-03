package zap

import (
	"encoding/xml"
	"os"
)

func Read(path string) (entities []interface{}, err error) {
	var file []byte
	file, err = os.ReadFile(path)
	if err != nil {
		return
	}
	var zap XMLConfigurator
	err = xml.Unmarshal(file, &zap)
	if err != nil {
		return
	}
	return
}
