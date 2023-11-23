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

func IsZigbee(conformance matter.Conformance) bool {
	if conformance == nil {
		return false
	}
	switch conformance := conformance.(type) {
	case *MandatoryConformance:
		if ex, ok := conformance.Expression.(*IdentifierExpression); ok && ex.ID == "Zigbee" {
			return true
		}
	case ConformanceSet:
		if len(conformance) > 0 {
			if mc, ok := conformance[0].(*MandatoryConformance); ok {
				if ex, ok := mc.Expression.(*IdentifierExpression); ok && ex.ID == "Zigbee" {
					return true
				}
			}
			return false
		}
	}
	return false
}
