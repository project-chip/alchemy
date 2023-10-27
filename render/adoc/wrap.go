package adoc

import "strings"

var newlineStripper = strings.NewReplacer(
	"\n", " ",
	"\r", " ",
)

func wrapText(s string, width int) string {
	s = newlineStripper.Replace(s)
	s = strings.TrimSpace(s)
	if len(s) <= width {
		return s
	}
	var b strings.Builder
	for {
		if len(s) <= width {
			b.WriteString(s)
			break
		}
		i := strings.LastIndex(" ", s[:width])
		if i == -1 {
			b.WriteString(s[:width])
			b.WriteRune('\n')
			s = s[width:]
			continue
		}
		b.WriteString(strings.TrimRight(s[:i], " "))
		s = strings.TrimLeft(s[i:], " ")
	}
	return b.String()
}
