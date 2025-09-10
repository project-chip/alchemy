package conformance

import (
	"encoding/json"
	"strings"

	"github.com/project-chip/alchemy/internal/text"
)

type Set []Conformance

func (cs Set) Type() Type {
	return TypeSet
}

func (cs Set) ASCIIDocString() string {
	var s strings.Builder
	hitNonQualified := false
	for _, c := range cs {
		if hitNonQualified {
			// Once we've hit an unqualified optional or mandatory conformance, only deprecated conformance is allowed
			if _, ok := c.(*Deprecated); !ok {
				continue
			}
		}
		if s.Len() > 0 {
			s.WriteString(", ")
		}
		s.WriteString(text.TrimUnnecessaryParens(c.ASCIIDocString()))
		switch c := c.(type) {
		case *Disallowed, *Deprecated:
			return s.String()
		case *Optional:
			if c.Expression == nil {
				hitNonQualified = true
			}
		case *Mandatory:
			if c.Expression == nil {
				hitNonQualified = true
			}
		}
	}
	return s.String()
}

func (cs Set) Description() string {
	var s strings.Builder
	for _, c := range cs {
		if s.Len() > 0 {
			s.WriteString(", otherwise ")
		}
		s.WriteString(c.Description())
	}
	if len(cs) > 1 {
		s.WriteString(", otherwise disallowed")
	}
	return s.String()
}

func (cs Set) Eval(context Context) (ConformanceState, error) {
	var state ConformanceState
	for _, c := range cs {
		cs, err := c.Eval(context)
		if err != nil {
			return ConformanceState{State: StateUnknown, Confidence: ConfidenceDefinite}, nil
		}
		if cs.State == StateUnknown {
			continue
		}
		switch state.State {
		case StateUnknown:
			// We don't have state yet, so use this state
			state = cs
		case StateMandatory:
			switch cs.State {
			case StateMandatory:
				// If the existing state is mandatory, override only if the confidence is higher
				if cs.Confidence > state.Confidence {
					state.Confidence = cs.Confidence
				}
			}
		case StateOptional:
			switch cs.State {
			case StateOptional:
				// If the existing state is optional, override only if the confidence is higher
				if cs.Confidence > state.Confidence {
					state.Confidence = cs.Confidence
					state.State = cs.State
				}
			}
		case StateDisallowed:
		default:
			if cs.Confidence > state.Confidence {
				state.Confidence = cs.Confidence
				state.State = cs.State
			}
		}
	}
	if state.State != StateUnknown {
		return state, nil
	}
	return ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}, nil
}

func (cs Set) Equal(c Conformance) bool {
	ocs, ok := c.(Set)
	if !ok {
		return false
	}
	if len(cs) != len(ocs) {
		return false
	}
	for i, c := range cs {
		oc := cs[i]
		if !oc.Equal(c) {
			return false
		}
	}
	return true
}

func (cs Set) Clone() Conformance {
	return cs.CloneSet()
}

func (cs Set) CloneSet() Set {
	ncs := make(Set, 0, len(cs))
	for _, c := range cs {
		ncs = append(ncs, c.Clone())
	}
	return ncs
}

func (cs Set) MarshalJSON() ([]byte, error) {
	var js []*jsonConformance
	for _, c := range cs {
		js = append(js, &jsonConformance{Conformance: c})
	}
	return json.Marshal(js)
}

type jsonConformance struct {
	Conformance
}

func (jc jsonConformance) MarshalJSON() ([]byte, error) {

	type jscType struct {
		Type Type `json:"type"`
	}

	js, err := json.Marshal(jscType{Type: jc.Conformance.Type()})
	if err != nil {
		return nil, err
	}
	cjs, err := json.Marshal(jc.Conformance)
	if err != nil {
		return nil, err
	}
	if cjs[0] == '{' && cjs[1] == '}' {
		return js, nil
	}
	cjs[0] = ','
	val := append(js[:len(js)-1], cjs...)
	return val, nil
}
