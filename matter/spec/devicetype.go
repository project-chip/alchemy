package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toDeviceTypes(d *Doc) (entities []types.Entity, err error) {
	var deviceTypes []*matter.DeviceType
	var description string
	p := parse.FindFirst[*asciidoc.Paragraph](s.Elements())
	if p != nil {
		se := parse.FindFirst[*asciidoc.String](p.Elements())
		if se != nil {
			description = strings.ReplaceAll(se.Value, "\n", " ")
		}
	}

	for _, s := range parse.Skim[*Section](s.Elements()) {
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

		elements := parse.Skim[*Section](s.Elements())
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
			case matter.SectionComposedDeviceTypeRequirements:
				c.ComposedDeviceTypeRequirements, err = s.toComposedDeviceTypeRequirements(d)
			case matter.SectionConditions:
				c.Conditions, err = s.toConditions(d)
			case matter.SectionDeviceTypeRequirements:
				c.DeviceTypeRequirements, err = s.toDeviceTypeRequirements(d)
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
	ti, err := parseFirstTable(doc, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading device type ID: %w", err)
	}
	var deviceTypes []*matter.DeviceType
	for row := range ti.Body() {
		c := matter.NewDeviceType(row)
		c.ID, err = ti.ReadID(row, matter.TableColumnID)
		if err != nil {
			return nil, err
		}
		c.Name, err = ti.ReadString(row, matter.TableColumnDeviceName)
		if err != nil {
			return nil, err
		}
		c.Superset, err = ti.ReadString(row, matter.TableColumnSuperset)
		if err != nil {
			return nil, err
		}
		c.Class, err = ti.ReadString(row, matter.TableColumnClass)
		if err != nil {
			return nil, err
		}
		c.Scope, err = ti.ReadString(row, matter.TableColumnScope)
		if err != nil {
			return nil, err
		}
		deviceTypes = append(deviceTypes, c)
	}

	return deviceTypes, nil
}

func (d *Doc) toBaseDeviceType() (baseDeviceType *matter.DeviceType, err error) {
	for _, top := range parse.Skim[*Section](d.Elements()) {
		err = AssignSectionTypes(d, top)
		if err != nil {
			return
		}
		var baseClusterRequirements, elementRequirements *Section
		parse.Traverse(top, top.Elements(), func(sec *Section, parent parse.HasElements, index int) parse.SearchShould {
			switch sec.SecType {
			case matter.SectionClusterRequirements:
				baseClusterRequirements = sec
			case matter.SectionElementRequirements:
				elementRequirements = sec
			}
			return parse.SearchShouldContinue
		})
		if baseClusterRequirements == nil && elementRequirements == nil {
			continue
		}
		baseDeviceType = matter.NewDeviceType(top.Base)
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
	return nil, fmt.Errorf("failed to find base device type")
}
