package spec

import (
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toConditions(d *Doc, dt *matter.DeviceType) (conditions []*matter.Condition, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s.Base, "error reading conditions table: %w", err)
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
				err = newGenericParseError(ti.Element, "failed to find tag column in section %s", s.Name)
				return
			}
		}
	}
	for row := range ti.Body() {
		c := matter.NewCondition(s.Base, dt)
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

func (s *Section) toBaseDeviceTypeConditions(d *Doc, dt *matter.DeviceType) (conditions []*matter.Condition, err error) {
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
			c := matter.NewCondition(row, dt)
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
		err = readRowCellValueElements(d, row, row.Set, &sb)
		if err != nil {
			continue
		}
		c := matter.NewCondition(row, dt)
		c.Feature = sb.String()
		conditions = append(conditions, c)
	}
	return
}

type conditionFinder struct {
	entityFinderCommon

	deviceType     *matter.DeviceType
	baseDeviceType *matter.DeviceType
}

func newConditionFinder(deviceType *matter.DeviceType, baseDeviceType *matter.DeviceType, inner entityFinder) *conditionFinder {
	return &conditionFinder{
		entityFinderCommon: entityFinderCommon{inner: inner},
		deviceType:         deviceType,
		baseDeviceType:     baseDeviceType,
	}
}

func (cf *conditionFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	for _, con := range cf.deviceType.Conditions {
		if con.Feature == identifier && con != cf.identity {
			return con
		}
	}
	if cf.baseDeviceType != nil {
		for _, con := range cf.baseDeviceType.Conditions {
			if con.Feature == identifier && con != cf.identity {
				return con
			}
		}
	}
	if cf.inner != nil {
		return cf.inner.findEntityByIdentifier(identifier, source)
	}
	return nil

}

func (cf *conditionFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {

	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for _, con := range cf.deviceType.Conditions {
			if con == cf.identity {
				continue
			}
			if !yield(con.Feature, con) {
				return
			}
		}
		if cf.baseDeviceType != nil {
			for _, con := range cf.baseDeviceType.Conditions {
				if con == cf.identity {
					continue
				}
				if !yield(con.Feature, con) {
					return
				}
			}
		}
	})

	if cf.inner != nil {
		cf.inner.suggestIdentifiers(identifier, suggestions)
	}
	return
}
