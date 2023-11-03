package zap

import (
	"encoding/xml"
	"fmt"
	"os"
)

func Read(path string) (models []interface{}, err error) {
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
	fmt.Printf("zap file len %d\n", len(zap.Clusters))
	return
}
