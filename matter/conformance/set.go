package conformance

import (
	"strings"
)

type ConformanceSet []Conformance

func (cs ConformanceSet) String() string {
	var s strings.Builder
	for _, c := range cs {
		if s.Len() > 0 {
			s.WriteString(", otherwise ")
		}
		s.WriteString(c.String())
		switch c := c.(type) {
		case *OptionalConformance, *ProvisionalConformance, *DisallowedConformance, *DeprecatedConformance:
			return s.String()
		case *MandatoryConformance:
			if c.Expression == nil {
				return s.String()
			}
		}
	}
	if len(cs) > 1 {
		s.WriteString(", otherwise disallowed")
	}
	return s.String()
}

func (cs ConformanceSet) Eval(context ConformanceContext) (ConformanceState, error) {
	for _, c := range cs {
		cs, err := c.Eval(context)
		if err != nil {
			return ConformanceStateUnknown, err
		}
		if cs == ConformanceStateUnknown {
			continue
		}
		return cs, nil
	}
	return ConformanceStateDisallowed, nil
}
