package conformance

import (
	"strings"

	"github.com/hasty/alchemy/matter"
)

type OtherwiseConformance struct {
	Conformances []matter.Conformance
}

func (oc *OtherwiseConformance) String() string {
	var s strings.Builder
	for _, c := range oc.Conformances {
		if s.Len() > 0 {
			s.WriteString(", otherwise ")
		}
		s.WriteString(c.String())
	}
	s.WriteString(", otherwise disallowed")
	return s.String()
}

func (oc *OtherwiseConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	for _, c := range oc.Conformances {
		cs, err := c.Eval(context)
		if err != nil {
			return matter.ConformanceStateUnknown, err
		}
		if cs == matter.ConformanceStateUnknown {
			continue
		}
		return cs, nil
	}
	return matter.ConformanceStateDisallowed, nil
}
