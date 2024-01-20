package generate

import (
	"slices"
	"strings"
)

func mergeLines(lines []string, newLineMap map[string]struct{}, skip int) []string {
	for _, l := range lines {
		delete(newLineMap, l)
	}
	if len(newLineMap) == 0 {
		return lines
	}
	insertedLines := make([]string, 0, len(newLineMap))
	for newLine := range newLineMap {
		lines = append(lines, newLine)
		insertedLines = append(insertedLines, newLine)
	}
	reorderLinesSemiAlphabetically(lines, insertedLines, skip)
	return lines
}

func reorderLinesSemiAlphabetically(list []string, newLines []string, skip int) {
	for _, insertedName := range newLines {
		currentIndex := slices.Index(list, insertedName)
		if currentIndex >= 0 {
			for i, key := range list {
				if i < skip {
					continue
				}
				if strings.Compare(insertedName, key) < 0 {
					if i < currentIndex {
						for j := currentIndex; j > i; j-- {
							list[j] = list[j-1]
						}
						list[i] = insertedName
					}
					break
				}
			}
		}
	}
}
