package matter

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
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

type EntityReference struct {
	ID   *Number                  `json:"id,omitempty"`
	Name string                   `json:"name,omitempty"`
	XRef *asciidoc.CrossReference `json:"-"`
}

func (er EntityReference) Clone() EntityReference {
	return EntityReference{
		ID:   er.ID.Clone(),
		Name: er.Name,
		XRef: er.XRef,
	}
}

func (er EntityReference) Equals(oer EntityReference) bool {
	if er.ID != nil && er.ID.Valid() && oer.ID != nil && oer.ID.Valid() {
		return er.ID.Equals(oer.ID)
	}
	if er.Name != "" && oer.Name != "" {
		return er.Name == oer.Name
	}
	if er.XRef != nil {
		return er.XRef.Equals(oer.XRef)
	}
	return false
}

type ElementReference struct {
	Name  string                   `json:"name,omitempty"`
	Field string                   `json:"field,omitempty"`
	XRef  *asciidoc.CrossReference `json:"-"`
}

func (er ElementReference) Equals(oer ElementReference) bool {
	if er.Name != "" {
		if er.Name != oer.Name {
			return false
		}
	}
	if er.Field != oer.Field {
		return false
	}
	if er.XRef != nil {
		return er.XRef.Equals(oer.XRef)
	}
	return true
}
