package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type LengthLimit struct {
	Value string `json:"value"`
}

func (ll *LengthLimit) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("len(%s)", ll.Value)
}

func (ll *LengthLimit) DataModelString(dataType *types.DataType) string {
	return ll.Value
}

func (ll *LengthLimit) Equal(o Limit) bool {
	if oc, ok := o.(*LengthLimit); ok {
		return oc.Value == ll.Value
	}
	return false
}

func (ll *LengthLimit) Min(cc Context) (min types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(ll.Value)
	if rc == nil {
		return
	}
	return rc.Min(cc)
}

func (ll *LengthLimit) Max(cc Context) (max types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(ll.Value)
	if rc == nil {
		return
	}
	return rc.Max(cc)
}

func (ll *LengthLimit) Default(cc Context) (def types.DataTypeExtreme) {
	return cc.Default(ll.Value)
}

func (ll *LengthLimit) Clone() Limit {
	return &LengthLimit{Value: ll.Value}
}

func (ll *LengthLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "length",
		"value": ll.Value,
	}
	return json.Marshal(js)
}
