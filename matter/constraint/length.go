package constraint

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter/types"
)

type LengthLimit struct {
	Reference Limit `json:"ref"`
}

func (ll *LengthLimit) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("len(%s)", ll.Reference.ASCIIDocString(dataType))
}

func (ll *LengthLimit) DataModelString(dataType *types.DataType) string {
	return ll.Reference.DataModelString(dataType)
}

func (ll *LengthLimit) Equal(o Limit) bool {
	if oc, ok := o.(*LengthLimit); ok {
		return oc.Reference.Equal(ll.Reference)
	}
	return false
}

func (ll *LengthLimit) Min(cc Context) (min types.DataTypeExtreme) {
	switch ref := ll.Reference.(type) {
	case *ReferenceLimit:
		min = cc.MinEntityValue(ref.Entity, ref.Field)
	case *IdentifierLimit:
		min = cc.MinEntityValue(ref.Entity, ref.Field)
	default:
		slog.Warn("Unknown limit type on length limit", log.Type("type", ref))
	}
	return
}

func (ll *LengthLimit) Max(cc Context) (max types.DataTypeExtreme) {
	switch ref := ll.Reference.(type) {
	case *ReferenceLimit:
		max = ref.Max(cc)
	case *IdentifierLimit:
		max = ref.Max(cc)
	default:
		slog.Warn("Unknown limit type on length limit", log.Type("type", ref))
	}
	return
}

func (ll *LengthLimit) Fallback(cc Context) (def types.DataTypeExtreme) {
	switch ref := ll.Reference.(type) {
	case *ReferenceLimit:
		def = cc.Fallback(ref.Entity, ref.Field)
	case *IdentifierLimit:
		def = cc.Fallback(ref.Entity, ref.Field)
	default:
		slog.Warn("Unknown limit type on length limit", log.Type("type", ref))
	}
	return
}

func (ll *LengthLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (ll *LengthLimit) Clone() Limit {
	return &LengthLimit{Reference: ll.Reference.Clone()}
}

func (ll *LengthLimit) MarshalJSON() ([]byte, error) {
	js := struct {
		Type string `json:"type"`
		LengthLimit
	}{
		Type:        "length",
		LengthLimit: *ll,
	}
	return json.Marshal(js)
}
