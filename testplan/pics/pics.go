package pics

import (
	"fmt"
	"strconv"
	"strings"
)

type PICSExpression struct {
	PICS string `json:"id"`
	Not  bool   `json:"not,omitempty"`
}

func (pe PICSExpression) String() string {
	if pe.Not {
		return "!" + pe.PICS
	}
	return pe.PICS
}

func (pe PICSExpression) PythonString() string {
	var sb strings.Builder
	if pe.Not {
		sb.WriteString("not ")
	}
	sb.WriteString(`self.check_pics(`)
	sb.WriteString(strconv.Quote(pe.PICS))
	sb.WriteString(")")
	return sb.String()
}

func (pe PICSExpression) PythonBuilder(aliases map[string]string, sb *strings.Builder) {
	alias, ok := aliases[pe.PICS]
	if pe.Not {
		sb.WriteString("not ")
	}
	if ok && alias != "" {
		sb.WriteString(alias)
		return
	}
	sb.WriteString(`self.check_pics(`)
	sb.WriteString(strconv.Quote(pe.PICS))
	sb.WriteString(")")
}

func ParseString(pics string) (Expression, error) {
	pics = strings.TrimSpace(pics)
	if pics == "" {
		return nil, nil
	}
	c, err := Parse("", []byte(pics))
	if err != nil {
		err = fmt.Errorf("error parsing PICS \"%s\": %w", pics, err)
		return nil, err
	}
	switch c := c.(type) {
	case Expression:
		return c, nil
	case []any:
		for _, cs := range c {
			fmt.Printf("type: %T\n", cs)
		}
	}
	return c.(Expression), nil
}
