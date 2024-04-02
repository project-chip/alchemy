package matter

import (
	"strings"
)

var DisallowedReferenceSuffixes = []string{"Command", "Feature", "Attribute", "Field", "Event"}

func StripReferenceSuffixes(newID string) string {
	for _, suffix := range DisallowedReferenceSuffixes {
		if strings.HasSuffix(newID, suffix) {
			newID = newID[0 : len(newID)-len(suffix)]
			break
		}
	}
	return newID
}
