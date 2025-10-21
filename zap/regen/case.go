package regen

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/project-chip/alchemy/internal/text"
)

var (
	wordPattern         = regexp.MustCompile(`[ _\-/]`)
	invalidCharsPattern = regexp.MustCompile(`[+()&]`)
	cleanupPattern      = regexp.MustCompile(`[^A-Za-z0-9_]`)
)

func caseify(s string, camelCase bool, preserveAcronyms bool) string {
	cleanedLabel := invalidCharsPattern.ReplaceAllString(s, "")
	tokens := wordPattern.Split(cleanedLabel, -1)

	var result strings.Builder
	for index, token := range tokens {
		runes := []rune(token)

		if len(runes) == 0 {
			continue
		}

		isAllUpperCase := text.IsUpperCase(token)

		if isAllUpperCase && preserveAcronyms {
			result.WriteString(token)
			continue
		}

		firstRune := runes[0]
		var processedFirstRune rune

		if index == 0 && camelCase {
			processedFirstRune = unicode.ToLower(firstRune)
		} else {
			processedFirstRune = unicode.ToUpper(firstRune)
		}
		result.WriteRune(processedFirstRune)

		if len(runes) > 1 {
			result.WriteString(string(runes[1:]))
		} else {
			result.WriteString(strings.ToLower(string(runes[1:])))
		}
	}

	str := result.String()

	if camelCase {
		originalRunes := []rune(s)
		if !wordPattern.MatchString(s) &&
			len(originalRunes) > 1 &&
			string(originalRunes[0:2]) == strings.ToUpper(string(originalRunes[0:2])) &&
			s != strings.ToUpper(s) {

			// JS: str = str[0].toUpperCase() + str.substring(1)
			// Re-uppercase the first letter that was just lowercased.
			strRunes := []rune(str)
			if len(strRunes) > 0 {
				strRunes[0] = unicode.ToUpper(strRunes[0])
				str = string(strRunes)
			}
		}
	}

	return cleanupPattern.ReplaceAllString(str, "")
}
