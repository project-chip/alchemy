package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (s *Section) toDeviceTypes(d *Doc) (models []matter.Model, err error) {
	var deviceTypes []*matter.DeviceType
	var description string
	p := parse.FindFirst[*types.Paragraph](s.Elements)
	if p != nil {
		se := parse.FindFirst[*types.StringElement](p.Elements)
		if se != nil {
			description = strings.ReplaceAll(se.Content, "\n", " ")
		}
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionClassification:
			deviceTypes, err = readDeviceTypeIDs(d, s)
		}
		if err != nil {
			return nil, err
		}
	}

	for _, c := range deviceTypes {
		c.Description = description

		elements := parse.Skim[*Section](s.Elements)
		for _, s := range elements {
			switch s.SecType {
			case matter.SectionClusterRequirements:
				var crs []*matter.ClusterRequirement
				crs, err = s.toClusterRequirements(d)
				if err == nil {
					c.ClusterRequirements = append(c.ClusterRequirements, crs...)
				}
			case matter.SectionElementRequirements:
				c.ElementRequirements, err = s.toElementRequirements(d)
			case matter.SectionConditions:
				c.Conditions, err = s.toConditions(d)
			case matter.SectionRevisionHistory:
				c.Revisions, err = readRevisionHistory(d, s)
			default:
			}
			if err != nil {
				return nil, fmt.Errorf("error reading section in %s: %w", d.Path, err)
			}
		}
	}
	for _, c := range deviceTypes {
		models = append(models, c)
	}
	return models, nil
}

func readDeviceTypeIDs(doc *Doc, s *Section) ([]*matter.DeviceType, error) {
	rows, headerRowIndex, columnMap, _, err := parseFirstTable(doc, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading device type ID: %w", err)
	}
	var deviceTypes []*matter.DeviceType
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		c := &matter.DeviceType{}
		c.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return nil, err
		}
		c.Name, err = readRowValue(row, columnMap, matter.TableColumnDeviceName)
		if err != nil {
			return nil, err
		}
		c.Superset, err = readRowValue(row, columnMap, matter.TableColumnSuperset)
		if err != nil {
			return nil, err
		}
		c.Class, err = readRowValue(row, columnMap, matter.TableColumnClass)
		if err != nil {
			return nil, err
		}
		c.Scope, err = readRowValue(row, columnMap, matter.TableColumnScope)
		if err != nil {
			return nil, err
		}
		deviceTypes = append(deviceTypes, c)
	}

	return deviceTypes, nil
}
