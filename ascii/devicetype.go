package ascii

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toDeviceTypes(d *Doc) (entities []mattertypes.Entity, err error) {
	var deviceTypes []*matter.DeviceType
	var description string
	p := parse.FindFirst[*elements.Paragraph](s.Elements)
	if p != nil {
		se := parse.FindFirst[elements.String](p.Elements())
		if se != "" {
			description = strings.ReplaceAll(string(se), "\n", " ")
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
		entities = append(entities, c)
	}
	return entities, nil
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
		c.Name, err = readRowASCIIDocString(row, columnMap, matter.TableColumnDeviceName)
		if err != nil {
			return nil, err
		}
		c.Superset, err = readRowASCIIDocString(row, columnMap, matter.TableColumnSuperset)
		if err != nil {
			return nil, err
		}
		c.Class, err = readRowASCIIDocString(row, columnMap, matter.TableColumnClass)
		if err != nil {
			return nil, err
		}
		c.Scope, err = readRowASCIIDocString(row, columnMap, matter.TableColumnScope)
		if err != nil {
			return nil, err
		}
		deviceTypes = append(deviceTypes, c)
	}

	return deviceTypes, nil
}

func (d *Doc) toBaseDeviceType() (baseDeviceType *matter.DeviceType, err error) {
	for _, e := range d.Elements {
		switch e := e.(type) {
		case *Section:
			var baseClusterRequirements, elementRequirements *Section
			parse.Traverse(e, e.Elements, func(sec *Section, parent parse.HasElements, index int) bool {
				switch sec.SecType {
				case matter.SectionClusterRequirements:
					baseClusterRequirements = sec
				case matter.SectionElementRequirements:
					elementRequirements = sec
				}
				return false
			})
			if baseClusterRequirements == nil && elementRequirements == nil {
				continue
			}
			baseDeviceType = &matter.DeviceType{}
			if baseClusterRequirements != nil {
				baseDeviceType.ClusterRequirements, err = baseClusterRequirements.toClusterRequirements(d)
				if err != nil {
					return
				}
			}
			if elementRequirements != nil {
				baseDeviceType.ElementRequirements, err = elementRequirements.toElementRequirements(d)
				if err != nil {
					return
				}
			}
			return
		}
	}
	if baseDeviceType == nil {
		return nil, fmt.Errorf("failed to find base device type")
	}
	return
}
