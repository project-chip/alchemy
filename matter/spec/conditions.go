package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
)

func (s *Section) toConditions(d *Doc) (conditions []*matter.Condition, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading conditions table: %w", err)
		}
		return
	}
	featureIndex, ok := ti.ColumnMap[matter.TableColumnFeature]
	if !ok {
		featureIndex, ok = ti.ColumnMap[matter.TableColumnCondition]
		if !ok {
			featureIndex = -1
			for _, col := range ti.ExtraColumns {
				if strings.HasSuffix(col.Name, "Tag") {
					featureIndex = col.Offset
					break
				}
			}
			if featureIndex == -1 {
				err = fmt.Errorf("failed to find tag column in section %s", s.Name)
				return
			}
		}
	}
	for row := range ti.Body() {
		c := matter.NewCondition(s.Base)
		c.Feature, err = ti.ReadStringAtOffset(row, featureIndex)
		if err != nil {
			return
		}
		c.Description, err = ti.ReadString(row, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conditions = append(conditions, c)
	}
	return
}

func (s *Section) toBaseDeviceTypeConditions(d *Doc) (conditions []*matter.Condition, err error) {
	if !text.HasCaseInsensitiveSuffix(s.Name, " Conditions") {
		return
	}

	var ti *TableInfo
	t := FindFirstTable(s)
	if t == nil {
		return
	}
	ti, err = parseTable(d, s, t)
	if err == nil {
		tagOffset := -1
		for _, col := range ti.ExtraColumns {
			if text.HasCaseInsensitiveSuffix(col.Name, "Tag") {
				tagOffset = col.Offset
				break
			}
		}
		if tagOffset == -1 {
			return
		}
		for row := range ti.Body() {
			c := matter.NewCondition(row)
			c.Feature, err = ti.ReadStringAtOffset(row, tagOffset)
			if err != nil {
				return
			}
			c.Description, err = ti.ReadString(row, matter.TableColumnDescription, matter.TableColumnSummary)
			if err != nil {
				return
			}
			conditions = append(conditions, c)
		}
		return
	}

	if t.ColumnCount != 1 {
		return
	}
	// There are some condition tables with no valid Matter columns, so we handle them manually

	for i, row := range t.TableRows() {
		if i == 0 {
			// Skip the first row, as it's a header
			continue
		}
		var sb strings.Builder
		err = readRowCellValueElements(d, row.Set, &sb)
		if err != nil {
			continue
		}
		c := matter.NewCondition(row)
		c.Feature = sb.String()
		conditions = append(conditions, c)
	}
	return
}
