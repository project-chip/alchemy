package constraint

import (
	"encoding/json"
	"fmt"
)

func UnmarshalConstraint(raw json.RawMessage) (c Constraint, err error) {
	var list []json.RawMessage
	err = json.Unmarshal(raw, &list)
	if err == nil {
		c, err = UnmarshalConstraintSetJSON(list)
		return
	}
	var base constraintJSONBase
	err = json.Unmarshal(raw, &base)
	if err != nil {
		return
	}
	switch base.Type {
	case "generic":
		c = &GenericConstraint{}
		err = json.Unmarshal(raw, c)
		if err != nil {
			return
		}
	default:
		err = fmt.Errorf("unknown constraint type: \"%s\"", base.Type)
	}
	return
}

type constraintJSONBase struct {
	Type string `json:"type"`
}

func UnmarshalLimit(raw json.RawMessage) (cl Limit, err error) {
	var base constraintJSONBase
	err = json.Unmarshal(raw, &base)
	if err != nil {
		return
	}
	switch base.Type {
	case "boolean":
		cl = &BooleanLimit{}
		err = json.Unmarshal(raw, cl)
	case "string":
		cl = &StringLimit{}
		err = json.Unmarshal(raw, cl)
	case "temperature":
		cl = &TemperatureLimit{}
		err = json.Unmarshal(raw, cl)
	case "unspecified":
		cl = &UnspecifiedLimit{}
	default:
		err = fmt.Errorf("unknown constraint limit type: \"%s\"", base.Type)
	}
	return
}
