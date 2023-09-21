package disco

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hasty/matterfmt/ascii"
)

func getDocType(doc *ascii.Doc) (MatterDoc, error) {
	if len(doc.Path) == 0 {
		return MatterDocUnknown, fmt.Errorf("missing path")
	}

	path, err := filepath.Abs(doc.Path)
	if err != nil {
		return MatterDocUnknown, err
	}
	dir, file := filepath.Split(path)
	pathParts := strings.Split(dir, string(os.PathSeparator))

	for i := len(pathParts) - 1; i >= 0; i-- {
		part := pathParts[i]
		switch part {
		case "device_types":
			if firstLetterIsLower(file) {
				return MatterDocDeviceTypeIndex, nil
			}
			return MatterDocDeviceType, nil
		case "app_clusters":
			if firstLetterIsLower(file) {
				return MatterDocAppClusterIndex, nil
			}
			return MatterDocAppCluster, nil
		}
	}
	return MatterDocUnknown, fmt.Errorf("could not determine doc type from path %s", doc.Path)
}

func firstLetterIsLower(s string) bool {
	firstLetter, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLower(firstLetter)
}
