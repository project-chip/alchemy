package text

import (
	"strings"
	"unicode"
)

// ToIDLSnakeCase converts a string, typically CamelCase or PascalCase, to snake_case.
// It includes custom replacements to handle acronyms commonly found in Matter.
func ToIDLSnakeCase(text string) string {
	replacer := strings.NewReplacer(
		"BLE", "Ble",
		"IDs", "Ids",
		"IPv4", "Ipv4",
		"IPv6", "Ipv6",
		"iOS", "Ios",
		"Int8U", "Int8u",
		"KWh", "Kwh",
		"KVAh", "Kvah",
	)
	text = replacer.Replace(text)

	runes := []rune(text)
	if len(runes) == 0 {
		return ""
	}

	var splits []string
	var current strings.Builder
	current.WriteRune(runes[0])

	for i := 1; i < len(runes); i++ {
		shouldSplit := false
		r := runes[i]
		prev := runes[i-1]

		if unicode.IsUpper(r) {
			if !unicode.IsUpper(prev) {
				shouldSplit = true
			} else if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				shouldSplit = true
			}
		}

		if shouldSplit {
			splits = append(splits, current.String())
			current.Reset()
		}
		current.WriteRune(r)
	}
	splits = append(splits, current.String())

	for i, s := range splits {
		splits[i] = strings.ToLower(s)
	}

	return strings.Join(splits, "_")
}
