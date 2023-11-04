package conformance

import "strings"

func ParseConformance(conformance string, options ...interface{}) (Conformance, error) {
	c, err := ParseReader("", strings.NewReader(conformance))
	if err != nil {
		return nil, err
	}
	return c.(Conformance), nil
}
