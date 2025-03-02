package python

import (
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/testplan/pics"
)

func picsHelper(pics []pics.Expression) raymond.SafeString {
	ps := make([]string, 0, len(pics))
	for _, r := range pics {
		ps = append(ps, r.PythonString())
	}
	return raymond.SafeString(strings.Join(ps, ","))
}

func picsGuardHelper(exp pics.Expression, aliases map[string]string) raymond.SafeString {
	var sb strings.Builder
	exp.PythonBuilder(aliases, &sb)
	return raymond.SafeString(sb.String())
}
