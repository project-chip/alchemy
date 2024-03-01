package text

import "strings"

type commentState int

const (
	commentStateNone = iota
	commentStateForwardSlash
	commentStateInSingleLineComment
	commentStateInMultiLineComment
	commentStateAsteriskInMultilineComment
	commentStateInString
	commentStateInStringEscaped
)

func RemoveComments(s string) string {
	var out strings.Builder
	out.Grow(len(s))
	var state commentState
	var last rune
	for _, r := range s {
		switch state {
		case commentStateNone:
			switch r {
			case '/':
				state = commentStateForwardSlash
				last = r
				continue
			case '"':
				state = commentStateInString
			default:
				if last != 0 {
					out.WriteRune(last)
					last = 0
				}
			}
		case commentStateForwardSlash:
			switch r {
			case '/':
				state = commentStateInSingleLineComment
				last = 0
				continue
			case '*':
				state = commentStateInMultiLineComment
				last = 0
				continue
			default:
				if last != 0 {
					out.WriteRune(last)
					last = 0
				}
				state = commentStateNone
			}
		case commentStateInSingleLineComment:
			switch r {
			case '\n':
				state = commentStateNone
			default:
				continue
			}
		case commentStateInMultiLineComment:
			if r == '*' {
				state = commentStateAsteriskInMultilineComment
			}
			continue
		case commentStateAsteriskInMultilineComment:
			switch r {
			case '/':
				state = commentStateNone
			default:
				state = commentStateInMultiLineComment
			}
			continue
		case commentStateInString:
			switch r {
			case '"':
				state = commentStateNone
			case '\\':
				state = commentStateInStringEscaped
			}
		case commentStateInStringEscaped:
			state = commentStateInString
		}

		out.WriteRune(r)
	}
	return out.String()
}
