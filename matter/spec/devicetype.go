package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type DeviceTypeSet pipeline.Map[string, *pipeline.Data[[]*matter.DeviceType]]

func (library *Library) toDeviceTypes(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section) (deviceTypes []*matter.DeviceType, err error) {

	for s := range parse.Skim[*asciidoc.Section](reader, s, reader.Children(s)) {
		switch library.SectionType(s) {
		case matter.SectionClassification:
			deviceTypes, err = readDeviceTypeIDs(reader, d, s)
		}
		if err != nil {
			return
		}
	}

	if len(deviceTypes) == 0 {
		return
	}

	description := library.getDescription(reader, d, deviceTypes[0], s, reader.Children(s))

	for _, dt := range deviceTypes {
		dt.Description = description

		elements := parse.FindAll[*asciidoc.Section](d, reader, s)
		for s := range elements {
			switch library.SectionType(s) {
			case matter.SectionClusterRequirements:
				var crs []*matter.ClusterRequirement
				crs, err = library.toClusterRequirements(reader, d, s, dt)
				if err == nil {
					dt.ClusterRequirements = append(dt.ClusterRequirements, crs...)
				}
			case matter.SectionElementRequirements:
				var extraClusterRequirements []*matter.ClusterRequirement
				dt.ElementRequirements, extraClusterRequirements, err = library.toElementRequirements(reader, d, s, dt)
				dt.ClusterRequirements = append(dt.ClusterRequirements, extraClusterRequirements...)
			case matter.SectionConditions:
				dt.Conditions, err = library.toConditions(reader, d, s, dt)
			case matter.SectionSemanticTagRequirements:
				dt.TagRequirements, err = library.toTagRequirements(reader, d, s, dt)
			case matter.SectionDeviceTypeRequirements:
				dt.DeviceTypeRequirements, err = library.toDeviceTypeRequirements(reader, d, s, dt)
			case matter.SectionComposedDeviceTypeClusterRequirements:
				dt.ComposedDeviceTypeClusterRequirements, err = library.toComposedDeviceTypeClusterRequirements(reader, d, s, dt)
			case matter.SectionComposedDeviceTypeElementRequirements:
				var extraComposedDeviceClusterRequirements []*matter.DeviceTypeClusterRequirement
				dt.ComposedDeviceTypeElementRequirements, extraComposedDeviceClusterRequirements, err = library.toComposedDeviceTypeElementRequirements(reader, d, s, dt)
				dt.ComposedDeviceTypeClusterRequirements = append(dt.ComposedDeviceTypeClusterRequirements, extraComposedDeviceClusterRequirements...)
			case matter.SectionComposedDeviceTypeSemanticTagRequirements:
				dt.ComposedDeviceTagRequirements, err = library.toDeviceTypeTagRequirements(reader, d, s, dt)
			case matter.SectionConditionRequirements:
				dt.ConditionRequirements, err = library.toConditionRequirements(reader, d, s, dt)
			case matter.SectionRevisionHistory:
				dt.Revisions, err = readRevisionHistory(reader, d, s, dt)
			default:
			}
			if err != nil {
				return
			}
		}
	}
	for _, c := range deviceTypes {
		library.addEntity(s, c)
	}
	return
}

func readDeviceTypeIDs(reader asciidoc.Reader, doc *asciidoc.Document, s *asciidoc.Section) ([]*matter.DeviceType, error) {
	ti, err := parseFirstTable(reader, doc, s)
	if err != nil {
		return nil, newGenericParseError(s, "failed reading device type ID: %w", err)
	}
	var deviceTypes []*matter.DeviceType
	for row := range ti.ContentRows() {
		c := matter.NewDeviceType(row)
		c.ID, err = ti.ReadID(reader, row, matter.TableColumnDeviceID, matter.TableColumnID)
		if err != nil {
			return nil, err
		}
		c.Name, err = ti.ReadString(reader, row, matter.TableColumnDeviceName, matter.TableColumnDeviceName)
		if err != nil {
			return nil, err
		}
		c.SupersetOf, err = ti.ReadString(reader, row, matter.TableColumnSupersetOf)
		if err != nil {
			return nil, err
		}
		c.Class, err = ti.ReadString(reader, row, matter.TableColumnClass)
		if err != nil {
			return nil, err
		}
		c.Scope, err = ti.ReadString(reader, row, matter.TableColumnScope)
		if err != nil {
			return nil, err
		}
		deviceTypes = append(deviceTypes, c)
	}

	return deviceTypes, nil
}

func (library *Library) toBaseDeviceType(reader asciidoc.Reader, section *asciidoc.Section) (baseDeviceType *matter.DeviceType, err error) {
	doc := section.Document()
	var baseClusterRequirements, elementRequirements *asciidoc.Section
	parse.Search(doc, reader, section, reader.Children(section), func(doc *asciidoc.Document, sec *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		switch library.SectionType(sec) {
		case matter.SectionClusterRequirements:
			baseClusterRequirements = sec
		case matter.SectionElementRequirements:
			elementRequirements = sec
		}
		return parse.SearchShouldContinue
	})
	if baseClusterRequirements == nil && elementRequirements == nil {
		err = fmt.Errorf("unable to find base device type cluster requirements and element requirements")
		return
	}

	baseDeviceType = matter.NewDeviceType(section)
	baseDeviceType.Name = "Base Device Type"
	baseDeviceType.ID = matter.InvalidID
	if baseClusterRequirements != nil {
		baseDeviceType.ClusterRequirements, err = library.toClusterRequirements(reader, doc, baseClusterRequirements, baseDeviceType)
		if err != nil {
			return
		}
	}
	if elementRequirements != nil {
		var extraClusterRequirements []*matter.ClusterRequirement
		baseDeviceType.ElementRequirements, extraClusterRequirements, err = library.toElementRequirements(reader, doc, elementRequirements, baseDeviceType)
		baseDeviceType.ClusterRequirements = append(baseDeviceType.ClusterRequirements, extraClusterRequirements...)
		if err != nil {
			return
		}
	}
	parse.Search(doc, reader, section, reader.Children(section), func(doc *asciidoc.Document, sec *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		switch library.SectionType(sec) {

		case matter.SectionConditions:
			var conditions []*matter.Condition
			conditions, err = library.toBaseDeviceTypeConditions(reader, doc, sec, baseDeviceType)

			baseDeviceType.Conditions = append(baseDeviceType.Conditions, conditions...)
		case matter.SectionRevisionHistory:
			baseDeviceType.Revisions, err = readRevisionHistory(reader, doc, sec, baseDeviceType)
		}
		if err != nil {
			return parse.SearchShouldStop
		}
		return parse.SearchShouldContinue
	})
	return
}

func (spec *Specification) associateDeviceTypeRequirements() (err error) {
	entityFinder := newSpecEntityFinder(spec, nil, nil)

	if spec.BaseDeviceType != nil {
		err = spec.associateDeviceTypeRequirement(spec.BaseDeviceType, entityFinder)
		if err != nil {
			return
		}
	}
	for _, dt := range spec.DeviceTypes {
		err = spec.associateDeviceTypeRequirement(dt, entityFinder)
		if err != nil {
			return
		}
	}
	for _, dt := range spec.DeviceTypes {
		err = spec.associateComposedDeviceTypeRequirement(dt, entityFinder)
		if err != nil {
			return
		}
	}
	return
}

func (spec *Specification) associateDeviceTypeRequirement(dt *matter.DeviceType, entityFinder entityFinder) (err error) {
	switch dt.SupersetOf {
	case "":
		if dt != spec.BaseDeviceType {
			dt.SubsetDeviceType = spec.BaseDeviceType
		}
	default:
		superset, ok := spec.DeviceTypesByName[dt.SupersetOf]
		if !ok {
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
		cr.Cluster = findDeviceTypeRequirementCluster(spec, cr.ClusterRef, cr, entityFinder)
	}
	for _, er := range dt.ElementRequirements {
		if er.Cluster != nil {
			continue
		}
		er.Cluster = findDeviceTypeRequirementCluster(spec, er.ClusterRef, er, entityFinder)
	}
	referencedClusters := make(map[*matter.Cluster]struct{})
	buildReferencedClusters(dt, referencedClusters)
	for _, er := range dt.ElementRequirements {
		err = associateElementRequirement(spec, dt, er, referencedClusters, entityFinder)
		if err != nil {
			return
		}
	}

	return
}

func (spec *Specification) associateComposedDeviceTypeRequirement(dt *matter.DeviceType, entityFinder entityFinder) (err error) {
	deviceTypes := make(map[*matter.DeviceType]*matter.DeviceTypeRequirement)
	for _, dr := range dt.DeviceTypeRequirements {
		if dr.DeviceType == nil {
			dr.DeviceType = findDeviceTypeRequirementDeviceType(spec, dr.DeviceTypeRef, dr, entityFinder)
		}
		if dr.DeviceType != nil {
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
	for _, cr := range dt.ConditionRequirements {
		if cr.DeviceType == nil {
			cr.DeviceType = findDeviceTypeRequirementDeviceType(spec, cr.DeviceTypeRef, cr, entityFinder)
		}
		if cr.DeviceType == nil {
			continue
		}
		if cr.Condition == nil {
			cr.Condition = findDeviceTypeRequirementCondition(spec, cr.ConditionRef, cr.DeviceType, entityFinder)
		}
	}
	for _, tr := range dt.TagRequirements {
		tr.Namespace = findTagRequirementNamespace(spec, tr.NamespaceRef, tr, entityFinder)
		if tr.Namespace != nil && tr.SemanticTag == nil && (tr.SemanticTagRef.ID.Valid() || tr.SemanticTagRef.Name != "") {
			tr.SemanticTag = findTagRequirementTag(spec, tr.Namespace, tr.SemanticTagRef, tr, entityFinder)
		}
	}
	for _, cr := range dt.ComposedDeviceTypeClusterRequirements {
		if cr.ClusterRequirement.Cluster == nil {
			cr.ClusterRequirement.Cluster = findDeviceTypeRequirementCluster(spec, cr.ClusterRequirement.ClusterRef, cr.ClusterRequirement, entityFinder)
			if cr.ClusterRequirement.Cluster == nil {
				slog.Error("unknown cluster ID for cluster requirement on composing device type",
					slog.String("clusterId", cr.ClusterRequirement.ClusterRef.ID.HexString()),
					slog.String("clusterName", cr.ClusterRequirement.ClusterRef.Name),
					slog.String("deviceType", dt.Name),
					log.Path("source", cr.ClusterRequirement))
				spec.addError(&UnknownEntityError[*matter.ClusterRequirement]{EntityType: types.EntityTypeCluster, EntityReference: &cr.ClusterRequirement.ClusterRef, Source: cr.ClusterRequirement})
			}
		}
		if cr.DeviceType == nil {
			referencedDeviceType := findDeviceTypeRequirementDeviceType(spec, cr.DeviceTypeRef, cr.ClusterRequirement, entityFinder)
			if referencedDeviceType == nil {
			} else {
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok && referencedDeviceType != spec.RootNodeDeviceType && referencedDeviceType != spec.BaseDeviceType {
					slog.Error("Cluster requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", cr.DeviceTypeRef.ID.HexString()),
						slog.String("deviceTypeName", cr.DeviceTypeRef.Name),
						slog.String("deviceType", dt.Name),
						log.Path("source", cr.ClusterRequirement))
					spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.ClusterRequirement]{Entity: dt, Source: cr.ClusterRequirement})
				} else {
					cr.DeviceType = referencedDeviceType
					cr.DeviceTypeRequirement = dtr
				}
			}
		}
	}
	for _, er := range dt.ComposedDeviceTypeElementRequirements {
		if er.ElementRequirement.Cluster == nil {
			er.ElementRequirement.Cluster = findDeviceTypeRequirementCluster(spec, er.ElementRequirement.ClusterRef, er.ElementRequirement, entityFinder)
			if er.ElementRequirement.Cluster == nil {
				slog.Error("unknown cluster ID for element requirement on composing device type",
					slog.String("clusterId", er.ElementRequirement.ClusterRef.ID.HexString()),
					slog.String("clusterName", er.ElementRequirement.ClusterRef.Name),
					slog.String("deviceType", dt.Name),
					log.Path("source", er.ElementRequirement))
				spec.addError(&UnknownEntityError[*matter.ElementRequirement]{EntityType: types.EntityTypeCluster, EntityReference: &er.ElementRequirement.ClusterRef, Source: er.ElementRequirement})
			}
		}
		if er.DeviceType == nil {
			referencedDeviceType := findDeviceTypeRequirementDeviceType(spec, er.DeviceTypeRef, er.ElementRequirement, entityFinder)
			if referencedDeviceType == nil {
				slog.Error("unknown device type ID for cluster requirement on composing device type",
					slog.String("deviceTypeId", er.DeviceTypeRef.ID.HexString()),
					slog.String("deviceTypeName", er.DeviceTypeRef.Name),
					slog.String("deviceType", dt.Name),
					log.Path("source", er.ElementRequirement))
				spec.addError(&UnknownEntityError[*matter.ElementRequirement]{EntityType: types.EntityTypeDeviceType, EntityReference: &er.DeviceTypeRef, Source: er.ElementRequirement})
			} else {
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok && referencedDeviceType != spec.RootNodeDeviceType && referencedDeviceType != spec.BaseDeviceType {
					slog.Error("Element requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", er.DeviceTypeRef.ID.HexString()),
						slog.String("deviceTypeName", er.DeviceTypeRef.Name),
						slog.String("deviceType", dt.Name),
						log.Path("source", er.ElementRequirement))
					spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.ElementRequirement]{Entity: dt, Source: er.ElementRequirement})
				} else {
					er.DeviceType = referencedDeviceType
					er.DeviceTypeRequirement = dtr
				}
			}
		}
	}
	for _, er := range dt.ComposedDeviceTypeElementRequirements {
		if er.DeviceType == nil {
			continue
		}
		referencedClusters := make(map[*matter.Cluster]struct{})
		buildReferencedClusters(er.DeviceType, referencedClusters)
		err = associateElementRequirement(spec, er.DeviceType, er.ElementRequirement, referencedClusters, entityFinder)
		if err != nil {
			return
		}
	}
	for _, tr := range dt.ComposedDeviceTagRequirements {

		if tr.DeviceType == nil {
			referencedDeviceType := findDeviceTypeRequirementDeviceType(spec, tr.DeviceTypeRef, dt, entityFinder)
			if referencedDeviceType == nil {
				slog.Error("unknown device type ID for cluster requirement on composing device type",
					slog.String("deviceTypeId", tr.DeviceTypeRef.ID.HexString()),
					slog.String("deviceTypeName", tr.DeviceTypeRef.Name),
					slog.String("deviceType", dt.Name),
					log.Path("source", tr))
				spec.addError(&UnknownEntityError[*matter.DeviceTypeTagRequirement]{EntityType: types.EntityTypeDeviceType, EntityReference: &tr.DeviceTypeRef, Source: tr})
			} else {
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok && referencedDeviceType != spec.RootNodeDeviceType && referencedDeviceType != spec.BaseDeviceType {
					slog.Error("Element requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", tr.DeviceTypeRef.ID.HexString()),
						slog.String("deviceTypeName", tr.DeviceTypeRef.Name),
						slog.String("deviceType", dt.Name),
						log.Path("source", tr))
					spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.DeviceTypeTagRequirement]{Entity: dt, Source: tr})
				} else {
					tr.DeviceType = referencedDeviceType
					tr.DeviceTypeRequirement = dtr
				}
			}
		}
		tr.TagRequirement.Namespace = findTagRequirementNamespace(spec, tr.TagRequirement.NamespaceRef, tr.TagRequirement, entityFinder)
		if tr.TagRequirement.Namespace == nil {
			slog.Error("unknown namespace for tag requirement on composing device type",
				slog.String("namespaceId", tr.TagRequirement.NamespaceRef.ID.HexString()),
				slog.String("namespaceName", tr.TagRequirement.NamespaceRef.Name),
				log.Path("source", tr))
			spec.addError(&UnknownEntityError[*matter.DeviceTypeTagRequirement]{EntityType: types.EntityTypeNamespace, EntityReference: &tr.TagRequirement.NamespaceRef, Source: tr})
		} else if tr.TagRequirement.SemanticTag == nil && (tr.TagRequirement.SemanticTagRef.ID.Valid() || tr.TagRequirement.SemanticTagRef.Name != "") {
			tr.TagRequirement.SemanticTag = findTagRequirementTag(spec, tr.TagRequirement.Namespace, tr.TagRequirement.SemanticTagRef, tr.TagRequirement, entityFinder)
			if tr.TagRequirement.SemanticTag == nil {
				slog.Error("unknown semantic tag for tag requirement on composing device type",
					slog.String("semanticTagId", tr.TagRequirement.SemanticTagRef.ID.HexString()),
					slog.String("semanticTagName", tr.TagRequirement.SemanticTagRef.Name),
					log.Path("source", tr))
				spec.addError(&UnknownEntityError[*matter.DeviceTypeTagRequirement]{EntityType: types.EntityTypeSemanticTag, EntityReference: &tr.TagRequirement.SemanticTagRef, Source: tr})
			}
		}
	}
	return
}

func findDeviceTypeRequirementCluster[SourceT types.Entity](spec *Specification, ref matter.EntityReference, entity SourceT, entityFinder entityFinder) (cluster *matter.Cluster) {
	if ref.XRef != nil {
		cluster, _ = referenceEntity[*matter.Cluster](spec, ref.XRef, entityFinder)
	}
	if cluster == nil && ref.ID.Valid() {
		cluster = spec.ClustersByID[ref.ID.Value()]
	}
	if cluster == nil && ref.Name != "" {
		var ok bool
		if cluster, ok = spec.ClustersByName[ref.Name]; ok {
			slog.Warn("linking cluster requirement by name on device type since cluster ID was not recognized",
				slog.String("clusterId", ref.ID.HexString()),
				slog.String("clusterName", ref.Name),
				matter.LogEntity("deviceTypeRequirement", entity),
				log.Path("source", entity))
			return
		}
	}
	return
}

func findDeviceTypeRequirementCondition(spec *Specification, ref matter.EntityReference, deviceType *matter.DeviceType, entityFinder entityFinder) (condition *matter.Condition) {
	if ref.XRef != nil {
		condition, _ = referenceEntity[*matter.Condition](spec, ref.XRef, entityFinder)
	}

	if condition == nil && ref.Name != "" {
		for _, c := range deviceType.Conditions {
			if c.Feature == ref.Name {
				condition = c
				break
			}
		}
	}

	if condition != nil {
		if ref.Name != "" && ref.Name != condition.Feature {
			slog.Error("Mismatch between condition requirement name and condition name", slog.String("conditionName", condition.Feature), slog.String("requirementName", ref.Name), log.Path("source", deviceType))
		}
		if !deviceType.Contains(condition) {
			slog.Error("Condition requirement on device type refers to unknown condition", slog.String("condition", condition.Feature), slog.String("deviceType", deviceType.Name), log.Path("source", deviceType))
			condition = nil
		}
	}
	return
}

func findDeviceTypeRequirementDeviceType[SourceT types.Entity](spec *Specification, ref matter.EntityReference, entity SourceT, entityFinder entityFinder) (deviceType *matter.DeviceType) {
	if strings.EqualFold(ref.Name, "Base") {
		deviceType = spec.BaseDeviceType
		return
	}
	if ref.XRef != nil {
		deviceType, _ = referenceEntity[*matter.DeviceType](spec, ref.XRef, entityFinder)
	}
	if deviceType == nil && ref.ID.Valid() {
		deviceType = spec.DeviceTypesByID[ref.ID.Value()]
	}
	if deviceType == nil && ref.Name != "" {
		var ok bool
		if deviceType, ok = spec.DeviceTypesByName[ref.Name]; ok {
			slog.Warn("linking device type requirement by name since device type ID was not recognized",
				slog.String("deviceTypeId", ref.ID.HexString()),
				slog.String("deviceTypeName", ref.Name),
				matter.LogEntity("deviceTypeRequirement", entity),
				log.Path("source", entity))
		}
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

func associateElementRequirement(spec *Specification, dt *matter.DeviceType, er *matter.ElementRequirement, referencedClusters map[*matter.Cluster]struct{}, entityFinder entityFinder) (err error) {
	if er.Cluster == nil {
		return
	}
	_, ok := referencedClusters[er.Cluster]
	if !ok {
		slog.Error("Element Requirement references non-required cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), log.Path("source", er))
		return
	}
	cluster := er.Cluster
	for cluster != nil {
		er.Entity, err = associateElementRequirementFromCluster(spec, er, dt, cluster, entityFinder)
		if err != nil {
			return
		}
		if er.Entity != nil {
			break
		}
		cluster = cluster.ParentCluster
	}
	return
}

func associateElementRequirementFromCluster(spec *Specification, er *matter.ElementRequirement, dt *matter.DeviceType, cluster *matter.Cluster, entityFinder entityFinder) (entity types.Entity, err error) {
	if er.ElementRef.XRef != nil {
		var ok bool
		entity, ok = referenceEntity[types.Entity](spec, er.ElementRef.XRef, entityFinder)
		if ok {
			return
		}
	}
	switch er.Element {
	case types.EntityTypeAttribute:
		for _, a := range cluster.Attributes {
			if strings.EqualFold(a.Name, er.ElementRef.Name) {
				entity = a
				return
			}
		}
	case types.EntityTypeFeature:
		if cluster.Features == nil {
			slog.Error("Element Requirement references missing features", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), log.Path("source", er))
			return
		}
		for _, fb := range cluster.Features.Bits {
			f, ok := fb.(*matter.Feature)
			if !ok {
				continue
			}
			if f.Code == er.ElementRef.Name || strings.EqualFold(f.Name(), er.ElementRef.Name) {
				entity = f
				return
			}
		}
	case types.EntityTypeCommand:
		for _, cmd := range cluster.Commands {
			if strings.EqualFold(cmd.Name, er.ElementRef.Name) {
				entity = cmd
				return
			}
		}
	case types.EntityTypeCommandField:
		var command *matter.Command
		for _, cmd := range cluster.Commands {
			if strings.EqualFold(cmd.Name, er.ElementRef.Name) {
				command = cmd
				break
			}
		}
		if command == nil {
			break
		}
		for _, f := range command.Fields {
			if strings.EqualFold(f.Name, er.ElementRef.Field) {
				entity = f
				return
			}
		}
	case types.EntityTypeEvent:
		for _, e := range cluster.Events {
			if strings.EqualFold(e.Name, er.ElementRef.Name) {
				entity = e
				return
			}
		}

	default:
		slog.Error("Unexpected entity type", slog.String("entityType", er.Element.String()), log.Path("source", er))
		err = newGenericParseError(er, "unexpected element type: %s", er.Element.String())
	}
	return
}
