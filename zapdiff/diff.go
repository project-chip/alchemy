package zapdiff

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type xpathSegment struct {
	tag    string
	attr   string
	value  string
	isAttr bool
}

func parseEntityUniqueIdentifier(id string) ([]xpathSegment, error) {
	parts := strings.Split(id, "/")
	var segments []xpathSegment
	for _, p := range parts {
		seg := xpathSegment{}
		bracketIdx := strings.Index(p, "[")
		if bracketIdx == -1 {
			seg.tag = p
			segments = append(segments, seg)
			continue
		}
		seg.tag = p[:bracketIdx]
		rest := p[bracketIdx+1 : len(p)-1] // remove matching ]
		equalsIdx := strings.Index(rest, "=")
		if equalsIdx == -1 {
			return nil, fmt.Errorf("invalid segment: %s", p)
		}
		seg.attr = rest[:equalsIdx]
		seg.value = strings.Trim(rest[equalsIdx+1:], "'")
		if strings.HasPrefix(seg.attr, "@") {
			seg.attr = seg.attr[1:]
			seg.isAttr = true
		}
		segments = append(segments, seg)
	}
	return segments, nil
}

// getCustomDiffLines reads a file and returns lines around the target element.
// It returns the lines, the start index of the target in the returned lines,
// the end index of the target, and an error.
func getCustomDiffLines(path string, targetID string) ([]string, int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	segments, err := parseEntityUniqueIdentifier(targetID)
	if err != nil {
		return nil, 0, 0, err
	}

	currentLine := 0
	lastFoundLine := -1
	foundLastSegment := false

	for i, seg := range segments {
		found := false
		for j := currentLine; j < len(lines); j++ {
			line := lines[j]
			
			// Guard against crossing into a sibling of the parent element
			if i > 0 {
				parentSeg := segments[i-1]
				// If we see another instance of the parent tag, we've likely left its scope
				if strings.Contains(line, "<"+parentSeg.tag) {
					break
				}
				// If we see the closing tag of the parent, we've definitely left its scope
				if strings.Contains(line, "</"+parentSeg.tag+">") {
					break
				}
			}

			if strings.Contains(line, seg.tag) {
				if seg.isAttr {
					attrPattern1 := fmt.Sprintf(`%s="%s"`, seg.attr, seg.value)
					attrPattern2 := fmt.Sprintf(`%s='%s'`, seg.attr, seg.value)
					if strings.Contains(line, attrPattern1) || strings.Contains(line, attrPattern2) {
						lastFoundLine = j
						currentLine = j + 1
						found = true
						if i == len(segments)-1 {
							foundLastSegment = true
						}
						break
					}
					continue // Attributes must be on the same line as the tag
				}

				if seg.value != "" {
					if strings.Contains(line, seg.value) {
						lastFoundLine = j
						currentLine = j + 1 // Search next segment after this line
						found = true
						if i == len(segments)-1 {
							foundLastSegment = true
						}
						break
					}
					// Search up to 5 lines ahead for value
					foundValue := false
					for k := 1; k <= 5 && j+k < len(lines); k++ {
						if strings.Contains(lines[j+k], seg.value) {
							lastFoundLine = j // Keep tag line as the found line for context
							currentLine = j + k + 1
							found = true
							foundValue = true
							if i == len(segments)-1 {
								foundLastSegment = true
							}
							break
						}
					}
					if foundValue {
						break
					}
				}
				if seg.value == "" {
					lastFoundLine = j
					currentLine = j + 1
					found = true
					if i == len(segments)-1 {
						foundLastSegment = true
					}
					break
				}
			}
		}
		if found {
			// Check for self-closing tag if there are more segments to find
			if i < len(segments)-1 {
				trimmedLine := strings.TrimSpace(lines[lastFoundLine])
				if strings.HasSuffix(trimmedLine, "/>") {
					// Parent self-closes, cannot have children!
					found = false
				}
			}
		}
		if !found {
			// If a middle segment is not found, we might fail to find the target.
			break
		}
	}

	var targetIdx int
	highlight := true

	if foundLastSegment {
		targetIdx = lastFoundLine
	} else if lastFoundLine != -1 {
		// Fallback to parent
		targetIdx = lastFoundLine
		highlight = false
	} else {
		// Element and all parents not found
		return nil, 0, 0, nil
	}

	start := targetIdx - 8
	if start < 0 {
		start = 0
	}
	end := targetIdx + 8
	if end >= len(lines) {
		end = len(lines) - 1
	}

	resultLines := lines[start : end+1]
	relStart := -1
	relEnd := -1
	if highlight {
		relStart = targetIdx - start
		relEnd = targetIdx - start
	}

	return resultLines, relStart, relEnd, nil
}
