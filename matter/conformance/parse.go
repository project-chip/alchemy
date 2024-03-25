package conformance

import (
	"strings"
)

func ParseConformance(conformance string) Set {
	c, err := tryParseConformance(conformance)
	if err != nil {
		return Set{&Generic{raw: conformance}}
	}
	return c
}

func tryParseConformance(conformance string) (Set, error) {
	conformance = strings.ReplaceAll(conformance, "\\|", "|")

	c, err := ParseReader("", strings.NewReader(conformance))
	if err != nil {
		return nil, err
	}
	return c.(Set), err
}

func IsMandatory(conformance Conformance) bool {
	if conformance == nil {
		return false
	}
	switch conformance := conformance.(type) {
	case *Mandatory:
		return conformance.Expression == nil
	case Set:
		if len(conformance) > 0 {
			mc, ok := conformance[0].(*Mandatory)
			if ok {
				return mc.Expression == nil
			}
		}
	}
	return false
}

func IsDeprecated(conformance Conformance) bool {
	if conformance == nil {
		return false
	}
	switch conformance := conformance.(type) {
	case *Deprecated:
		return true
	case Set:
		if len(conformance) > 0 {
			_, ok := conformance[0].(*Deprecated)
			return ok
		}
	}
	return false
}

func IsDisallowed(conformance Conformance) bool {
	if conformance == nil {
		return false
	}
	switch conformance := conformance.(type) {
	case *Disallowed:
		return true
	case Set:
		if len(conformance) > 0 {
			_, ok := conformance[0].(*Disallowed)
			return ok
		}
	}
	return false
}

func IsZigbee(store IdentifierStore, conformance Conformance) bool {
	if conformance == nil {
		return false
	}
	var err error
	var withZigbee, withoutZigbee State
	cxt := Context{Identifiers: store, Values: map[string]any{"Zigbee": true}}
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
	if withoutZigbee == StateDisallowed && withZigbee != StateDisallowed {
		return true
	}
	return false
}
