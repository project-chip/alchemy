package generate

import (
	"slices"
)

func getVariables(t *test) []string {
	var variables []string
	variableNames := make(map[string]struct{})
	for _, s := range t.Groups {
		for _, s := range s.Steps {
			if s.Response.SaveAs != "" {
				if _, ok := variableNames[s.Response.SaveAs]; !ok {
					variableNames[s.Response.SaveAs] = struct{}{}
					variables = append(variables, s.Response.SaveAs)
				}
			}
		}
	}

	slices.Sort(variables)
	return variables
}
