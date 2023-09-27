package disco

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func organizeAttributesSection(doc *ascii.Doc, top *ascii.Section, attributes *ascii.Section) error {
	attributesTable := findFirstTable(attributes)
	if attributesTable != nil {
		err := organizeAttributesTable(doc, top, attributes, attributesTable)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no attributes table found")
	}
	return nil
}

func organizeAttributesTable(doc *ascii.Doc, top *ascii.Section, attributes *ascii.Section, attributesTable *types.Table) error {
	rows := combineRows(attributesTable)

	_, columnMap, extraColumns := findColumns(rows)

	if columnMap == nil {
		return fmt.Errorf("can't rearrange attributes table without header row")
	}

	if len(columnMap) < 5 {
		return fmt.Errorf("can't rearrange attributes table with so few matches")
	}

	err := fixAccessCells(doc, rows, columnMap)
	if err != nil {
		return err
	}

	sectionDataMap := getPotentialDataTypes(attributes, rows, columnMap)

	promoteDataTypes(top, attributes, sectionDataMap)

	reorderColumns(doc, attributes, rows, matter.AttributesTableColumnOrder[:], columnMap, extraColumns)

	return nil
}

var accessPattern = regexp.MustCompile(`^(?:(?:^|\s+)(?:(?P<ReadWrite>(?:R\*W)|(?:R\[W\])|(?:[RW]+))|(?P<Fabric>[FS]+)|(?P<Privileges>[VOMA]+)|(?P<Timed>T)))+\s*$`)
var accessPatternMatchMap map[int]matter.AccessCategory

func init() {
	accessPatternMatchMap = make(map[int]matter.AccessCategory)
	for i, name := range accessPattern.SubexpNames() {
		switch name {
		case "ReadWrite":
			accessPatternMatchMap[i] = matter.AccessCategoryReadWrite
		case "Fabric":
			accessPatternMatchMap[i] = matter.AccessCategoryFabric
		case "Privileges":
			accessPatternMatchMap[i] = matter.AccessCategoryPrivileges
		case "Timed":
			accessPatternMatchMap[i] = matter.AccessCategoryTimed
		}
	}
}

func fixAccessCells(doc *ascii.Doc, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (err error) {
	if len(rows) < 2 {
		return
	}
	accessIndex, ok := columnMap[matter.TableColumnAccess]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[accessIndex]
		vc, e := getCellValue(cell)
		if e != nil {
			continue
		}
		matches := accessPattern.FindStringSubmatch(vc)
		if matches == nil {
			continue
		}
		access := make(map[matter.AccessCategory]string)
		for i, s := range matches {
			if s == "" {
				continue
			}
			category, ok := accessPatternMatchMap[i]
			if !ok {
				continue
			}
			access[category] = s
		}
		if len(access) > 0 {
			var out strings.Builder
			for _, ac := range matter.AccessCategoryOrder {
				a, ok := access[ac]
				if !ok {
					continue
				}
				if ac == matter.AccessCategoryReadWrite && a == "R*W" { // fix deprecated form
					a = "R[W]"
				}
				if out.Len() > 0 {
					out.WriteRune(' ')
				}
				out.WriteString(a)
			}
			err = setCellString(cell, out.String())
			if err != nil {
				return
			}
		}
	}
	return
}
