package generate

import (
	"slices"
	"strings"
)

func writeVariables(t *test, sb *strings.Builder) {
	variableNames := make(map[string]struct{})
	var variables []string
	for _, s := range t.steps {
		if s.Response.SaveAs != "" {
			if _, ok := variableNames[s.Response.SaveAs]; !ok {
				variableNames[s.Response.SaveAs] = struct{}{}
				variables = append(variables, s.Response.SaveAs)
			}
		}
	}
	if len(variables) == 0 {
		return
	}
	slices.Sort(variables)
	sb.WriteRune('\n')
	for _, v := range variables {
		sb.WriteString("        ")
		sb.WriteString(v)
		sb.WriteString(" = None\n")
	}
	sb.WriteRune('\n')
}
