package conformance

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type IdentifierExpression struct {
	ID  string
	Not bool
}

func (ie *IdentifierExpression) String() string {
	if ie.Not {
		return fmt.Sprintf("not %s", ie.ID)
	}
	return ie.ID
}

func (ie *IdentifierExpression) Eval(context matter.ConformanceContext) (bool, error) {
	return evalIdentifier(context, ie.ID, ie.Not)
}

func evalIdentifier(context matter.ConformanceContext, id string, not bool) (bool, error) {
	if context.Values != nil {
		v, ok := context.Values[id]
		if ok {
			if b, ok := v.(bool); ok {
				return b != not, nil
			}
		}
	}
	if context.Store != nil {
		if context.VisitedReferences == nil {
			context.VisitedReferences = make(map[string]struct{})
		} else if _, ok := context.VisitedReferences[id]; ok {
			return false, nil
		}
		ref := context.Store.ConformanceReference(id)
		context.VisitedReferences[id] = struct{}{}
		if ref != nil {
			conf := ref.GetConformance()
			if conf != nil {
				cs, err := conf.Eval(context)
				if err != nil {
					return false, err
				}
				return (cs == matter.ConformanceStateMandatory || cs == matter.ConformanceStateOptional || cs == matter.ConformanceStateProvisional || cs == matter.ConformanceStateDeprecated) != not, nil
			}
		}
	}
	return not, nil
}
