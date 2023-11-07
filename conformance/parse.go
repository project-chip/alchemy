package conformance

import (
	"strings"

	"github.com/hasty/alchemy/matter"
)

func ParseConformance(conformance string, options ...interface{}) (matter.Conformance, error) {
	c, err := ParseReader("", strings.NewReader(conformance))
	if err != nil {
		return nil, err
	}
	return c.(matter.Conformance), nil
}
