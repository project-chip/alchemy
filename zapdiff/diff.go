package zapdiff

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

)

type xpathSegment struct {
	tag    string
	attr   string
	value  string
	isAttr bool // true for [@attr='val'], false for [tag='val']
}

var segmentRegex = regexp.MustCompile(`^([a-zA-Z0-9_-]+)(?:\[(@?)([a-zA-Z0-9_-]+)='([^']*)'\])?$`)

func parseEntityUniqueIdentifier(id string) ([]xpathSegment, error) {
	parts := strings.Split(id, "/")
	var segments []xpathSegment
	for _, part := range parts {
		if part == "" {
			continue
		}
		matches := segmentRegex.FindStringSubmatch(part)
		if matches == nil {
			return nil, fmt.Errorf("invalid segment: %s", part)
		}
		seg := xpathSegment{tag: matches[1]}
		if len(matches) > 3 && matches[3] != "" {
			seg.isAttr = matches[2] == "@"
			seg.attr = matches[3]
			seg.value = matches[4]
		}
		segments = append(segments, seg)
	}
	return segments, nil
}

// findElementLines searches for the element identified by id in the file at path and returns its lines.
func findElementLines(path string, id string) ([]string, int, error) {
	segments, err := parseEntityUniqueIdentifier(id)
	if err != nil {
		return nil, 0, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var allLines []string
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text()+"\n")
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	segIdx := 0
	foundParent := false
	parentStartLine := -1
	startLine := -1

	for i := 0; i < len(allLines); i++ {
		line := allLines[i]
		if segIdx >= len(segments) {
			break
		}
		seg := segments[segIdx]

		// Boundary check: if we are looking for a child segment, and we see the parent close tag, we fail!
		if segIdx > 0 {
			parentTag := segments[segIdx-1].tag
			if strings.Contains(line, "</"+parentTag+">") {
				return nil, 0, nil // Not found within parent boundary
			}
		}

		if seg.attr != "" && !seg.isAttr {
			// Child element filter
			if !foundParent {
				if strings.Contains(line, "<"+seg.tag) {
					foundParent = true
					parentStartLine = i
				}
			}
			if foundParent {
				childStr := fmt.Sprintf("<%s>%s</%s>", seg.attr, seg.value, seg.attr)
				if strings.Contains(line, childStr) {
					segIdx++
					foundParent = false
					if segIdx >= len(segments) {
						startLine = parentStartLine
					}
				}
			}
		} else {
			// Normal or attr filter
			if strings.Contains(line, "<"+seg.tag) {
				if seg.attr == "" {
					segIdx++
					if segIdx >= len(segments) {
						startLine = i
					}
				} else {
					attrStr1 := fmt.Sprintf(`%s="%s"`, seg.attr, seg.value)
					attrStr2 := fmt.Sprintf(`%s='%s'`, seg.attr, seg.value)
					if strings.Contains(line, attrStr1) || strings.Contains(line, attrStr2) {
						segIdx++
						if segIdx >= len(segments) {
							startLine = i
						}
					}
				}
			}
		}
	}

	if segIdx < len(segments) || startLine == -1 {
		return nil, 0, nil // Not found
	}

	targetTag := segments[len(segments)-1].tag
	
	if strings.Contains(allLines[startLine], "/>") {
		return []string{allLines[startLine]}, startLine + 1, nil
	}

	endLine := startLine
	closeTag := "</" + targetTag + ">"
	for i := startLine; i < len(allLines); i++ {
		if strings.Contains(allLines[i], closeTag) {
			endLine = i
			break
		}
	}

	return allLines[startLine : endLine+1], startLine + 1, nil
}
