package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
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
				var extraClusterRequirements []*matter.ClusterRequirement
				dt.ElementRequirements, extraClusterRequirements, err = s.toElementRequirements(d)
				dt.ClusterRequirements = append(dt.ClusterRequirements, extraClusterRequirements...)
			case matter.SectionComposedDeviceTypeClusterRequirements:
				dt.ComposedDeviceTypeClusterRequirements, err = s.toComposedDeviceTypeClusterRequirements(d)
			case matter.SectionComposedDeviceTypeElementRequirements:
				var extraComposedDeviceClusterRequirements []*matter.DeviceTypeClusterRequirement
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
		c.SupersetOf, err = ti.ReadString(row, matter.TableColumnSupersetOf)
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
			var extraClusterRequirements []*matter.ClusterRequirement
			baseDeviceType.ElementRequirements, extraClusterRequirements, err = elementRequirements.toElementRequirements(d)
			baseDeviceType.ClusterRequirements = append(baseDeviceType.ClusterRequirements, extraClusterRequirements...)
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

func (spec *Specification) associateDeviceTypeRequirements() (err error) {
	if spec.BaseDeviceType != nil {
		err = spec.associateDeviceTypeRequirement(spec.BaseDeviceType)
		if err != nil {
			return
		}
	}
	for _, dt := range spec.DeviceTypes {
		err = spec.associateDeviceTypeRequirement(dt)
		if err != nil {
			return
		}
	}
	for _, dt := range spec.DeviceTypes {
		err = spec.associateComposedDeviceTypeRequirement(dt)
		if err != nil {
			return
		}
	}
	return
}

func (spec *Specification) associateDeviceTypeRequirement(dt *matter.DeviceType) (err error) {
	switch dt.SupersetOf {
	case "":
		if dt != spec.BaseDeviceType {
			dt.SubsetDeviceType = spec.BaseDeviceType
		}
	default:
		superset, ok := spec.DeviceTypesByName[dt.SupersetOf]
		if !ok {
			spec.addError(&UnknownSupersetError{DeviceType: dt})
			slog.Error("Device type superset not found", slog.String("deviceType", dt.Name), slog.String("superset", dt.SupersetOf), log.Path("source", dt))
			break
		}
		if superset == dt {
			slog.Error("Device type superset is the same as the device type", slog.String("deviceType", dt.Name), log.Path("source", dt))
			break
		}
		dt.SubsetDeviceType = superset
	}
	for _, cr := range dt.ClusterRequirements {
		if cr.Cluster != nil {
			continue
		}
		cr.Cluster = findDeviceTypeRequirementCluster(spec, cr.ClusterID, cr.ClusterName, cr)
		if cr.Cluster == nil {
			slog.Error("unknown cluster ID for cluster requirement on device type",
				slog.String("clusterId", cr.ClusterID.HexString()),
				slog.String("clusterName", cr.ClusterName),
				slog.String("deviceType", dt.Name),
				log.Path("source", cr))
			spec.addError(&UnknownClusterRequirementError{Requirement: cr})
		}
	}
	for _, er := range dt.ElementRequirements {
		if er.Cluster != nil {
			continue
		}
		er.Cluster = findDeviceTypeRequirementCluster(spec, er.ClusterID, er.ClusterName, er)
		if er.Cluster == nil {
			slog.Error("unknown cluster ID for element requirement on device type",
				slog.String("clusterId", er.ClusterID.HexString()),
				slog.String("clusterName", er.ClusterName),
				slog.String("deviceType", dt.Name),
				log.Path("source", er))
			spec.addError(&UnknownElementRequirementClusterError{Requirement: er})
		}
	}
	referencedClusters := make(map[*matter.Cluster]struct{})
	buildReferencedClusters(dt, referencedClusters)
	for _, er := range dt.ElementRequirements {
		err = associateElementRequirement(spec, dt, er, referencedClusters)
		if err != nil {
			return
		}
	}

	return
}

func (spec *Specification) associateComposedDeviceTypeRequirement(dt *matter.DeviceType) (err error) {
	deviceTypes := make(map[*matter.DeviceType]*matter.DeviceTypeRequirement)
	for _, dr := range dt.DeviceTypeRequirements {
		if dr.DeviceType == nil {
			dr.DeviceType = findDeviceTypeRequirementDeviceType(spec, dr.DeviceTypeID, dr.DeviceTypeName, dr)
		}
		if dr.DeviceType == nil {
			slog.Error("unknown device type ID for cluster requirement on composing device type",
				slog.String("deviceTypeId", dr.DeviceTypeID.HexString()),
				slog.String("deviceTypeName", dr.DeviceTypeName),
				slog.String("deviceType", dt.Name),
				log.Path("source", dr))
			spec.addError(&UnknownComposingDeviceTypeRequirementDeviceTypeError{Requirement: dr})
		} else {
			deviceTypes[dr.DeviceType] = dr
			if dr.Location == matter.DeviceTypeRequirementLocationUnknown {
				switch dr.DeviceType.Class {
				case "Simple":
					dr.Location = matter.DeviceTypeRequirementLocationChildEndpoint
				case "Utility":
					dr.Location = matter.DeviceTypeRequirementLocationDeviceEndpoint
				default:
					slog.Error("Unable to determine location for device type requirement", slog.String("deviceTypeClass", dr.DeviceType.Class), slog.String("deviceTypeName", dr.DeviceType.Name), log.Path("source", dr))
				}
			}
		}
	}
	for _, cr := range dt.ComposedDeviceTypeClusterRequirements {
		if cr.ClusterRequirement.Cluster == nil {
			cr.ClusterRequirement.Cluster = findDeviceTypeRequirementCluster(spec, cr.ClusterRequirement.ClusterID, cr.ClusterRequirement.ClusterName, cr.ClusterRequirement)
			if cr.ClusterRequirement.Cluster == nil {
				slog.Error("unknown cluster ID for cluster requirement on composing device type",
					slog.String("clusterId", cr.ClusterRequirement.ClusterID.HexString()),
					slog.String("clusterName", cr.ClusterRequirement.ClusterName),
					slog.String("deviceType", dt.Name),
					log.Path("source", cr.ClusterRequirement))
				spec.addError(&UnknownComposingDeviceTypeRequirementClusterError{Requirement: cr})
			}
		}
		if cr.DeviceType == nil {
			referencedDeviceType := findDeviceTypeRequirementDeviceType(spec, cr.DeviceTypeID, cr.DeviceTypeName, cr.ClusterRequirement)
			if referencedDeviceType == nil {
				slog.Error("unknown device type ID for cluster requirement on composing device type",
					slog.String("deviceTypeId", cr.DeviceTypeID.HexString()),
					slog.String("deviceTypeName", cr.DeviceTypeName),
					slog.String("deviceType", dt.Name),
					log.Path("source", cr.ClusterRequirement))
				spec.addError(&UnknownComposingDeviceTypeClusterRequirementDeviceTypeError{Requirement: cr})
			} else {
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok {
					slog.Error("Cluster requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", cr.DeviceTypeID.HexString()),
						slog.String("deviceTypeName", cr.DeviceTypeName),
						slog.String("deviceType", dt.Name),
						log.Path("source", cr.ClusterRequirement))
					spec.addError(&UnreferencedComposingDeviceTypeClusterRequirementDeviceTypeError{Requirement: cr})
				} else {
					cr.DeviceType = referencedDeviceType
					cr.DeviceTypeRequirement = dtr
				}
			}
		}
	}
	for _, er := range dt.ComposedDeviceTypeElementRequirements {
		if er.ElementRequirement.Cluster == nil {
			er.ElementRequirement.Cluster = findDeviceTypeRequirementCluster(spec, er.ElementRequirement.ClusterID, er.ElementRequirement.ClusterName, er.ElementRequirement)
			if er.ElementRequirement.Cluster == nil {
				slog.Error("unknown cluster ID for element requirement on composing device type",
					slog.String("clusterId", er.ElementRequirement.ClusterID.HexString()),
					slog.String("clusterName", er.ElementRequirement.ClusterName),
					slog.String("deviceType", dt.Name),
					log.Path("source", er.ElementRequirement))
				spec.addError(&UnknownComposingElementRequirementClusterError{Requirement: er})
			}
		}
		if er.DeviceType == nil {
			referencedDeviceType := findDeviceTypeRequirementDeviceType(spec, er.DeviceTypeID, er.DeviceTypeName, er.ElementRequirement)
			if referencedDeviceType == nil {
				slog.Error("unknown device type ID for cluster requirement on composing device type",
					slog.String("deviceTypeId", er.DeviceTypeID.HexString()),
					slog.String("deviceTypeName", er.DeviceTypeName),
					slog.String("deviceType", dt.Name),
					log.Path("source", er.ElementRequirement))
				spec.addError(&UnknownComposingDeviceTypeElementRequirementDeviceTypeError{Requirement: er})
			} else {
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok {
					slog.Error("Element requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", er.DeviceTypeID.HexString()),
						slog.String("deviceTypeName", er.DeviceTypeName),
						slog.String("deviceType", dt.Name),
						log.Path("source", er.ElementRequirement))
					spec.addError(&UnreferencedComposingDeviceTypeElementRequirementDeviceTypeError{Requirement: er})
				} else {
					er.DeviceType = referencedDeviceType
					er.DeviceTypeRequirement = dtr
				}
			}
		}
	}
	for _, er := range dt.ComposedDeviceTypeElementRequirements {
		referencedClusters := make(map[*matter.Cluster]struct{})
		buildReferencedClusters(er.DeviceType, referencedClusters)
		err = associateElementRequirement(spec, er.DeviceType, er.ElementRequirement, referencedClusters)
		if err != nil {
			return
		}
	}
	return
}

func findDeviceTypeRequirementCluster(spec *Specification, id *matter.Number, name string, entity types.Entity) (cluster *matter.Cluster) {
	var ok bool
	if cluster, ok = spec.ClustersByID[id.Value()]; ok {
		if name != cluster.Name {
			slog.Warn("Mismatch between cluster requirement ID and cluster name", slog.String("clusterId", id.HexString()), slog.String("clusterName", cluster.Name), slog.String("requirementName", name), log.Path("source", entity))
			spec.addError(&ClusterReferenceNameMismatch{Cluster: cluster, Name: name, Source: entity})
		}
		return
	}
	if cluster, ok = spec.ClustersByName[name]; ok {
		slog.Warn("linking cluster requirement by name on device type since cluster ID was not recognized",
			slog.String("clusterId", id.HexString()),
			slog.String("clusterName", name),
			matter.LogEntity("deviceTypeRequirement", entity),
			log.Path("source", entity))
		return
	}
	return
}

func findDeviceTypeRequirementDeviceType(spec *Specification, id *matter.Number, name string, entity types.Entity) (deviceType *matter.DeviceType) {
	var ok bool
	if deviceType, ok = spec.DeviceTypesByID[id.Value()]; ok {
		if name != deviceType.Name {
			slog.Warn("Mismatch between device type ID and device type name", slog.String("deviceTypeId", id.HexString()), slog.String("deviceTypeName", deviceType.Name), slog.String("requirementName", name), log.Path("source", entity))
			spec.addError(&DeviceTypeReferenceNameMismatch{DeviceType: deviceType, Name: name, Source: entity})
		}
		return
	}
	if deviceType, ok = spec.DeviceTypesByName[name]; ok {
		slog.Warn("linking device type requirement by name since device type ID was not recognized",
			slog.String("deviceTypeId", id.HexString()),
			slog.String("deviceTypeName", name),
			matter.LogEntity("deviceTypeRequirement", entity),
			log.Path("source", entity))
		return
	}
	return
}

func buildReferencedClusters(deviceType *matter.DeviceType, referencedClusters map[*matter.Cluster]struct{}) {
	parent := deviceType.SubsetDeviceType
	if parent != nil {
		buildReferencedClusters(parent, referencedClusters)
	}
	for _, rc := range deviceType.ClusterRequirements {
		if rc.Cluster != nil {
			referencedClusters[rc.Cluster] = struct{}{}
		}
	}
	for _, dtr := range deviceType.DeviceTypeRequirements {
		if dtr.DeviceType != nil {
			buildReferencedClusters(dtr.DeviceType, referencedClusters)
		}
	}
}

func validateDeviceTypes(spec *Specification) {
	for _, dt := range spec.DeviceTypes {

		referencedClusters := make(map[*matter.Cluster]struct{})
		buildReferencedClusters(spec.BaseDeviceType, referencedClusters)
		buildReferencedClusters(dt, referencedClusters)
		for _, cr := range dt.ClusterRequirements {
			if cr.Cluster == nil {
				continue
			}
			name := stripNonAlphabeticalCharacters(cr.ClusterName)
			clusterName := stripNonAlphabeticalCharacters(cr.Cluster.Name)
			if !strings.EqualFold(name, clusterName) {
				slog.Warn("Cluster Requirement name mismatch", slog.String("deviceType", dt.Name), slog.String("clusterName", cr.ClusterName), slog.String("referencedName", cr.Cluster.Name))
				continue
			}
		}
		for _, er := range dt.ElementRequirements {
			validateAccess(spec, er, er.Access)
			validateElementRequirement(spec, dt, er, referencedClusters)
		}
		for _, der := range dt.ComposedDeviceTypeElementRequirements {
			validateAccess(spec, der.ElementRequirement, der.ElementRequirement.Access)
			validateElementRequirement(spec, der.DeviceType, der.ElementRequirement, referencedClusters)
		}
	}
}

func associateElementRequirement(spec *Specification, dt *matter.DeviceType, er *matter.ElementRequirement, referencedClusters map[*matter.Cluster]struct{}) (err error) {
	if er.Cluster == nil {
		return
	}
	_, ok := referencedClusters[er.Cluster]
	if !ok {
		slog.Error("Element Requirement references non-required cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), log.Path("source", er))
		return
	}
	switch er.Element {
	case types.EntityTypeAttribute:

		for _, a := range er.Cluster.Attributes {
			if strings.EqualFold(a.Name, er.Name) {
				er.Entity = a
				break
			}
		}
	case types.EntityTypeFeature:
		for _, fb := range er.Cluster.Features.Bits {
			f, ok := fb.(*matter.Feature)
			if !ok {
				continue
			}
			if f.Code == er.Name || strings.EqualFold(f.Name(), er.Name) {
				er.Entity = f
				break
			}
		}
	case types.EntityTypeCommand:
		for _, cmd := range er.Cluster.Commands {
			if strings.EqualFold(cmd.Name, er.Name) {
				er.Entity = cmd
				break
			}
		}
	case types.EntityTypeCommandField:
		var command *matter.Command
		for _, cmd := range er.Cluster.Commands {
			if strings.EqualFold(cmd.Name, er.Name) {
				command = cmd
				break
			}
		}
		if command == nil {
			break
		}
		for _, f := range command.Fields {
			if strings.EqualFold(f.Name, er.Field) {
				er.Entity = f
				break
			}
		}
	case types.EntityTypeEvent:
		for _, e := range er.Cluster.Events {
			if strings.EqualFold(e.Name, er.Name) {
				er.Entity = e
				break
			}
		}

	default:
		slog.Error("Unexpected entity type", slog.String("entityType", er.Element.String()), log.Path("source", er))
		err = fmt.Errorf("unexpected element type: %s", er.Element.String())
	}
	return
}

func validateElementRequirement(spec *Specification, dt *matter.DeviceType, er *matter.ElementRequirement, referencedClusters map[*matter.Cluster]struct{}) {
	if er.Cluster == nil {
		return
	}
	_, ok := referencedClusters[er.Cluster]
	if !ok {
		slog.Error("Element Requirement references non-required cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), log.Path("source", er))
		spec.addError(ElementRequirementUnreferencedClusterError{Requirement: er})
		return
	}
	switch er.Element {
	case types.EntityTypeAttribute:
		if er.Entity == nil {
			slog.Error("Element Requirement references unknown attribute", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("attributeName", er.Name), log.Path("source", er))
			spec.addError(ElementRequirementUnknownElementError{Requirement: er})
		}
	case types.EntityTypeFeature:
		if er.Entity == nil {
			slog.Error("Element Requirement references unknown feature", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("featureName", er.Name), log.Path("source", er))
			spec.addError(ElementRequirementUnknownElementError{Requirement: er})
		}
	case types.EntityTypeCommand:
		if er.Entity == nil {
			slog.Error("Element Requirement references unknown command", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("commandName", er.Name), log.Path("source", er))
			spec.addError(ElementRequirementUnknownElementError{Requirement: er})
		}
	case types.EntityTypeCommandField:
		if er.Entity == nil {
			slog.Error("Element Requirement references unknown command field", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("commandName", er.Name), slog.String("commandField", er.Field), log.Path("source", er))
			spec.addError(ElementRequirementUnknownElementError{Requirement: er})
		}
	case types.EntityTypeEvent:
		if er.Entity == nil {
			slog.Error("Element Requirement references unknown event", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("commandName", er.Name), log.Path("source", er))
			spec.addError(ElementRequirementUnknownElementError{Requirement: er})
		}
	default:
		slog.Error("Unknown entity type", slog.String("entityType", er.Element.String()))
	}
}
