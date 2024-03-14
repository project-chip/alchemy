package generate

import (
	"fmt"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
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

func patchNumberAttributeFormat(e *etree.Element, n *matter.Number, name string, valFormat string) {
	if !n.Valid() {
		return
	}
	ex := e.SelectAttr(name)
	if ex == nil {
		e.CreateAttr(name, fmt.Sprintf(valFormat, n.Value()))
		return
	}
	exn := matter.ParseNumber(ex.Value)
	if exn.Valid() && exn.Equals(n) {
		return
	}
	e.CreateAttr(name, fmt.Sprintf(valFormat, n.Value()))
}

func patchNumberAttribute(e *etree.Element, n *matter.Number, name string) {
	if !n.Valid() {
		return
	}
	ex := e.SelectAttr(name)
	if ex == nil {
		e.CreateAttr(name, n.HexString())
		return
	}
	exn := matter.ParseNumber(ex.Value)
	if exn.Valid() && exn.Equals(n) {
		return
	}
	e.CreateAttr(name, n.HexString())
}

func patchNumberElement(e *etree.Element, n *matter.Number) {
	if !n.Valid() {
		return
	}
	exn := matter.ParseNumber(e.Text())
	if exn.Valid() && exn.Equals(n) {
		return
	}
	e.SetText(n.HexString())
}
