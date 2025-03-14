package generate

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
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

type dataExtremeType int8

const (
	dataExtremeTypeMinimum  dataExtremeType = 0
	dataExtremeTypeMaximum                  = iota
	dataExtremeTypeFallback                 = iota
)

func patchDataExtremeAttribute(e *etree.Element, attribute string, de *types.DataTypeExtreme, field *matter.Field, dataExtremeType dataExtremeType) {
	if !de.Defined() || de.IsNull() {
		e.RemoveAttr(attribute)
		return
	}
	if de.IsNumeric() {
		switch dataExtremeType {
		case dataExtremeTypeMinimum:
			fieldMinimum := types.Min(zap.ToUnderlyingType(field.Type.BaseType), field.Quality.Has(matter.QualityNullable))
			if cmp, ok := de.Compare(fieldMinimum); ok && cmp == -1 {
				slog.Warn("Field has minimum lower than the range of its data type; overriding", slog.String("name", field.Name), log.Path("source", field), slog.String("specifiedMinimum", de.ZapString(field.Type)), slog.String("fieldMinimum", fieldMinimum.ZapString(field.Type)))
				de = &fieldMinimum
			}
			if types.Min(zap.ToUnderlyingType(field.Type.BaseType), false).ValueEquals(*de) {
				e.RemoveAttr(attribute)
				return
			}
		case dataExtremeTypeMaximum:
			fieldMaximum := types.Max(zap.ToUnderlyingType(field.Type.BaseType), field.Quality.Has(matter.QualityNullable))
			if cmp, ok := de.Compare(fieldMaximum); ok && cmp == 1 {
				slog.Warn("Field has maximum greater than the range of its data type; overriding", slog.String("name", field.Name), log.Path("source", field), slog.String("specifiedMaximum", de.ZapString(field.Type)), slog.String("fieldMaximum", fieldMaximum.ZapString(field.Type)))
				de = &fieldMaximum
			}
			if types.Max(zap.ToUnderlyingType(field.Type.BaseType), false).ValueEquals(*de) {
				e.RemoveAttr(attribute)
				return
			}
		}

		n := matter.NumberFromExtreme(de)
		ex := e.SelectAttr(attribute)
		if ex == nil {
			e.CreateAttr(attribute, de.ZapString(field.Type))
			return
		}
		exn := matter.ParseNumber(ex.Value)
		if exn.Valid() && exn.Equals(n) {
			return
		}
		e.CreateAttr(attribute, de.ZapString(field.Type))
	} else {
		def := de.ZapString(field.Type)
		if def != "" {
			e.CreateAttr(attribute, def)
		} else {
			e.RemoveAttr(attribute)
		}
	}
}
