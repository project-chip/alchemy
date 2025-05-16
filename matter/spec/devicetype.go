package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type DeviceTypeSet pipeline.Map[string, *pipeline.Data[[]*matter.DeviceType]]

func (s *Section) toDeviceTypes(spec *Specification, d *Doc, pc *parseContext) (err error) {
	var deviceTypes []*matter.DeviceType

	for s := range parse.Skim[*Section](s.Elements()) {
		switch s.SecType {
		case matter.SectionClassification:
			deviceTypes, err = readDeviceTypeIDs(d, s)
		}
		if err != nil {
			return
		}
	}

	if len(deviceTypes) == 0 {
		return
	}

	description := getDescription(d, deviceTypes[0], s.Elements())

	for _, dt := range deviceTypes {
		dt.Description = description

		elements := parse.FindAll[*Section](s)
		for s := range elements {
			switch s.SecType {
			case matter.SectionClusterRequirements:
				var crs []*matter.ClusterRequirement
				crs, err = s.toClusterRequirements(d)
				if err == nil {
					dt.ClusterRequirements = append(dt.ClusterRequirements, crs...)
				}
			case matter.SectionElementRequirements:
				dt.ElementRequirements, err = s.toElementRequirements(d)
			case matter.SectionComposedDeviceTypeClusterRequirements:
				dt.ComposedDeviceTypeClusterRequirements, err = s.toComposedDeviceTypeClusterRequirements(d)
			case matter.SectionComposedDeviceTypeElementRequirements:
				var extraComposedDeviceClusterRequirements []*matter.ComposedDeviceTypeClusterRequirement
				dt.ComposedDeviceTypeElementRequirements, extraComposedDeviceClusterRequirements, err = s.toComposedDeviceTypeElementRequirements(d)
				dt.ComposedDeviceTypeClusterRequirements = append(dt.ComposedDeviceTypeClusterRequirements, extraComposedDeviceClusterRequirements...)
			case matter.SectionConditions:
				dt.Conditions, err = s.toConditions(d, dt)
			case matter.SectionDeviceTypeRequirements:
				dt.DeviceTypeRequirements, err = s.toDeviceTypeRequirements(d)
			case matter.SectionRevisionHistory:
				dt.Revisions, err = readRevisionHistory(d, s)
			default:
			}
			if err != nil {
				err = fmt.Errorf("error reading section in %s: %w", d.Path, err)
				return
			}
		}
	}
	for _, c := range deviceTypes {
		pc.entities = append(pc.entities, c)
		pc.orderedEntities = append(pc.orderedEntities, c)
		pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], c)
	}
	return
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
	for top := range parse.Skim[*Section](d.Elements()) {
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
		baseDeviceType.Name = "Base Device Type"
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
		parse.Traverse(top, top.Elements(), func(sec *Section, parent parse.HasElements, index int) parse.SearchShould {
			switch sec.SecType {

			case matter.SectionConditions:
				var conditions []*matter.Condition
				conditions, err = sec.toBaseDeviceTypeConditions(d, baseDeviceType)

				baseDeviceType.Conditions = append(baseDeviceType.Conditions, conditions...)
			case matter.SectionRevisionHistory:
				baseDeviceType.Revisions, err = readRevisionHistory(d, sec)
			}
			if err != nil {
				return parse.SearchShouldStop
			}
			return parse.SearchShouldContinue
		})

		return
	}
	return nil, fmt.Errorf("failed to find base device type")
}

func validateDeviceTypes(spec *Specification) {
	for _, dt := range spec.DeviceTypes {

		requiredClusterIDs := make(map[uint64]*matter.Cluster)
		for _, cr := range spec.BaseDeviceType.ClusterRequirements {
			if !cr.ClusterID.Valid() {
				continue
			}
			clusterID := cr.ClusterID.Value()
			c, ok := spec.ClustersByID[clusterID]
			if !ok {
				slog.Error("Base Device Cluster Requirement references unknown cluster ID", slog.String("deviceType", dt.Name), slog.String("clusterId", cr.ClusterID.HexString()))
				requiredClusterIDs[clusterID] = nil
				continue
			}
			requiredClusterIDs[clusterID] = c
		}
		for _, cr := range dt.ClusterRequirements {
			if !cr.ClusterID.Valid() {
				continue
			}
			clusterID := cr.ClusterID.Value()
			c, ok := spec.ClustersByID[clusterID]
			if !ok {
				slog.Error("Cluster Requirement references unknown cluster ID", slog.String("deviceType", dt.Name), slog.String("clusterId", cr.ClusterID.HexString()))
				requiredClusterIDs[clusterID] = nil
				continue
			}
			requiredClusterIDs[clusterID] = c
			name := stripNonAlphabeticalCharacters(cr.ClusterName)
			clusterName := stripNonAlphabeticalCharacters(c.Name)
			if !strings.EqualFold(name, clusterName) {
				slog.Warn("Cluster Requirement name mismatch", slog.String("deviceType", dt.Name), slog.String("clusterName", cr.ClusterName), slog.String("referencedName", c.Name))
				continue
			}
		}
		for _, er := range dt.ElementRequirements {
			validateAccess(spec, er, er.Access)
			if !er.ClusterID.Valid() {
				continue
			}
			c, ok := requiredClusterIDs[er.ClusterID.Value()]
			if !ok {
				slog.Error("Element Requirement references non-required cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName))
				continue
			}
			if c == nil {
				slog.Error("Element Requirement references unknown cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName))
				continue
			}
			switch er.Element {
			case types.EntityTypeAttribute:
				found := false
				for _, a := range c.Attributes {
					if strings.EqualFold(a.Name, er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown attribute", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("attributeName", er.Name))
				}
			case types.EntityTypeFeature:
				found := false
				for _, fb := range c.Features.Bits {
					f, ok := fb.(*matter.Feature)
					if !ok {
						continue
					}
					if f.Code == er.Name || strings.EqualFold(f.Name(), er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown feature", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("featureName", er.Name))
				}
			case types.EntityTypeCommand:
				found := false
				for _, cmd := range c.Commands {
					if strings.EqualFold(cmd.Name, er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown command", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("commandName", er.Name))
				}
			case types.EntityTypeEvent:
				found := false
				for _, e := range c.Events {
					if strings.EqualFold(e.Name, er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown event", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("commandName", er.Name))
				}

			default:
				slog.Error("Unknown entity type", slog.String("entityType", er.Element.String()))
			}
		}
	}
}
