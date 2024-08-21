package conformance

import (
	"encoding/json"
	"strings"
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
		s.WriteString(c.ASCIIDocString())
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
	return trimUnnecessaryParens(s.String())
}

func trimUnnecessaryParens(as string) string {
	if len(as) > 2 {
		if as[0] == '(' && as[len(as)-1] == ')' {
			as = as[1 : len(as)-1]
		}
	}
	return as
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

func (cs Set) Eval(context Context) (State, error) {
	for _, c := range cs {
		cs, err := c.Eval(context)
		if err != nil {
			return StateUnknown, err
		}
		if cs == StateUnknown {
			continue
		}
		return cs, nil
	}
	return StateDisallowed, nil
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
