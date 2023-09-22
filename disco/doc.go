package disco

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func getDocType(doc *ascii.Doc) (matter.Doc, error) {
	if len(doc.Path) == 0 {
		return matter.DocUnknown, fmt.Errorf("missing path")
	}

	path, err := filepath.Abs(doc.Path)
	if err != nil {
		return matter.DocUnknown, err
	}
	dir, file := filepath.Split(path)
	pathParts := strings.Split(dir, string(os.PathSeparator))

	for i := len(pathParts) - 1; i >= 0; i-- {
		part := pathParts[i]
		switch part {
		case "device_types":
			if firstLetterIsLower(file) {
				return matter.DocDeviceTypeIndex, nil
			}
			return matter.DocDeviceType, nil
		case "app_clusters":
			if firstLetterIsLower(file) {
				return matter.DocAppClusterIndex, nil
			}
			return matter.DocAppCluster, nil
		}
	}
	return matter.DocUnknown, fmt.Errorf("could not determine doc type from path %s", doc.Path)
}

func firstLetterIsLower(s string) bool {
	firstLetter, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLower(firstLetter)
}
