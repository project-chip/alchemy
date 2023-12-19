package conformance

import (
	"encoding/json"
	"fmt"
)

type FeatureExpression struct {
	ID  string `json:"id"`
	Not bool   `json:"not"`
}

func (fe *FeatureExpression) String() string {
	if fe.Not {
		return fmt.Sprintf("not %s", fe.ID)
	}
	return fe.ID
}

func (fe *FeatureExpression) Eval(context Context) (bool, error) {
	return evalIdentifier(context, fe.ID, fe.Not)
}

func (fe *FeatureExpression) Equal(e Expression) bool {
	if fe == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	ofe, ok := e.(*FeatureExpression)
	if !ok {
		return false
	}
	if fe.Not != ofe.Not {
		return false
	}
	if fe.ID != ofe.ID {
		return false
	}
	return true
}

func (fe *FeatureExpression) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "feature",
		"id":   fe.ID,
	}
	if fe.Not {
		js["not"] = true
	}
	return json.Marshal(js)
}
