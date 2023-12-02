package conformance

import (
	"strings"

	"github.com/hasty/alchemy/matter"
)

func ParseConformance(conformance string, options ...interface{}) matter.Conformance {
	c, err := tryParseConformance(conformance, options...)
	if err != nil {
		return &GenericConformance{raw: conformance}
	}
	return c
}

func tryParseConformance(conformance string, options ...interface{}) (matter.Conformance, error) {
	conformance = strings.ReplaceAll(conformance, "\\|", "|")

	c, err := ParseReader("", strings.NewReader(conformance))
	if err != nil {
		return nil, err
	}
	return c.(matter.Conformance), err
}

func IsMandatory(conformance matter.Conformance) bool {
	if conformance == nil {
		return false
	}
	switch conformance := conformance.(type) {
	case *MandatoryConformance:
		return conformance.Expression == nil
	case ConformanceSet:
		if len(conformance) > 0 {
			mc, ok := conformance[0].(*MandatoryConformance)
			if ok {
				return mc.Expression == nil
			}
		}
	}
	return false
}

func IsDeprecated(conformance matter.Conformance) bool {
	if conformance == nil {
		return false
	}
	switch conformance := conformance.(type) {
	case *DeprecatedConformance:
		return true
	case ConformanceSet:
		if len(conformance) > 0 {
			_, ok := conformance[0].(*DeprecatedConformance)
			return ok
		}
	}
	return false
}

func IsZigbee(store matter.ConformanceValueStore, conformance matter.Conformance) bool {
	if conformance == nil {
		return false
	}
	var err error
	var withZigbee, withoutZigbee matter.ConformanceState
	cxt := matter.ConformanceContext{Store: store, Values: map[string]any{"Zigbee": true}}
	withZigbee, err = conformance.Eval(cxt)
	if err != nil {
		return false
	}
	delete(cxt.Values, "Zigbee")
	withoutZigbee, err = conformance.Eval(cxt)
	if err != nil {
		return false
	}
	// If the absence of Zigbee makes this no longer allowed, then we call it a Zigbee only feature
	if withoutZigbee == matter.ConformanceStateDisallowed && withZigbee != matter.ConformanceStateDisallowed {
		return true
	}
	return false
}
