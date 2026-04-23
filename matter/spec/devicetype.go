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
	/*for _, c := range deviceTypes {
		library.addEntity(s, c)
	}*/
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
	for _, cr := range dt.ConditionRequirements {
		if cr.DeviceType == nil {
			cr.DeviceType = findDeviceTypeRequirementDeviceType(spec, cr.DeviceTypeID, cr.DeviceTypeName, cr)
		}
		if cr.DeviceType == nil {
			slog.Error("unknown device type ID for condition requirement on device type",
				slog.String("deviceTypeId", cr.DeviceTypeID.HexString()),
				slog.String("deviceTypeName", cr.DeviceTypeName),
				slog.String("deviceType", dt.Name),
				log.Path("source", cr))
			spec.addError(&UnknownConditionRequirementDeviceTypeError{Requirement: cr})
			continue
		} else {
			if _, ok := deviceTypes[cr.DeviceType]; !ok && cr.DeviceType != spec.RootNodeDeviceType && cr.DeviceType != spec.BaseDeviceType {
				slog.Error("Condition requirement on device type refers to unincluded device type",
					slog.String("deviceTypeId", cr.DeviceTypeID.HexString()),
					slog.String("deviceTypeName", cr.DeviceTypeName),
					slog.String("deviceType", dt.Name),
					log.Path("source", cr))
				spec.addError(&UnreferencedConditionRequirementDeviceTypeError{Requirement: cr})
				continue
			}
		}
		if cr.Condition == nil {
			for _, condition := range cr.DeviceType.Conditions {
				if condition.Feature == cr.ConditionName {
					cr.Condition = condition
					break
				}
			}
			if cr.Condition == nil {
				slog.Error("unknown condition for condition requirement on device type",
					slog.String("deviceTypeId", cr.DeviceTypeID.HexString()),
					slog.String("deviceTypeName", cr.DeviceTypeName),
					slog.String("deviceType", dt.Name),
					slog.String("condition", cr.ConditionName),
					log.Path("source", cr))
				spec.addError(&UnknownConditionRequirementConditionError{Requirement: cr})
				continue
			}
		}
	}
	for _, tr := range dt.TagRequirements {
		tr.Namespace = findTagRequirementNamespace(spec, tr.NamespaceID, tr.NamespaceName, tr)
		if tr.Namespace == nil {
			slog.Error("unknown namespace for tag requirement on device type",
				slog.String("namespaceId", tr.NamespaceID.HexString()),
				slog.String("namespaceName", tr.NamespaceName),
				log.Path("source", tr))
			spec.addError(&UnknownNamespaceTagRequirementError{Requirement: tr})
		} else if tr.SemanticTag == nil && (tr.SemanticTagID.Valid() || tr.SemanticTagName != "") {
			tr.SemanticTag = findTagRequirementTag(spec, tr.Namespace, tr.SemanticTagID, tr.SemanticTagName, tr)
			if tr.SemanticTag == nil {
				slog.Error("unknown semantic tag for tag requirement on device type",
					slog.String("semanticTagId", tr.SemanticTagID.HexString()),
					slog.String("semanticTagName", tr.SemanticTagName),
					log.Path("source", tr))
				spec.addError(&UnknownTagRequirementError{Requirement: tr})
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
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok && referencedDeviceType != spec.RootNodeDeviceType && referencedDeviceType != spec.BaseDeviceType {
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
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok && referencedDeviceType != spec.RootNodeDeviceType && referencedDeviceType != spec.BaseDeviceType {
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
		if er.DeviceType == nil {
			continue
		}
		referencedClusters := make(map[*matter.Cluster]struct{})
		buildReferencedClusters(er.DeviceType, referencedClusters)
		err = associateElementRequirement(spec, er.DeviceType, er.ElementRequirement, referencedClusters)
		if err != nil {
			return
		}
	}
	for _, tr := range dt.ComposedDeviceTagRequirements {

		if tr.DeviceType == nil {
			referencedDeviceType := findDeviceTypeRequirementDeviceType(spec, tr.DeviceTypeID, tr.DeviceTypeName, dt)
			if referencedDeviceType == nil {
				slog.Error("unknown device type ID for cluster requirement on composing device type",
					slog.String("deviceTypeId", tr.DeviceTypeID.HexString()),
					slog.String("deviceTypeName", tr.DeviceTypeName),
					slog.String("deviceType", dt.Name),
					log.Path("source", tr))
				spec.addError(&UnknownComposingDeviceTypeTagRequirementDeviceTypeError{Requirement: tr})
			} else {
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok && referencedDeviceType != spec.RootNodeDeviceType && referencedDeviceType != spec.BaseDeviceType {
					slog.Error("Element requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", tr.DeviceTypeID.HexString()),
						slog.String("deviceTypeName", tr.DeviceTypeName),
						slog.String("deviceType", dt.Name),
						log.Path("source", tr))
					spec.addError(&UnreferencedTagRequirementDeviceTypeError{Requirement: tr})
				} else {
					tr.DeviceType = referencedDeviceType
					tr.DeviceTypeRequirement = dtr
				}
			}
		}
		tr.TagRequirement.Namespace = findTagRequirementNamespace(spec, tr.TagRequirement.NamespaceID, tr.TagRequirement.NamespaceName, tr.TagRequirement)
		if tr.TagRequirement.Namespace == nil {
			slog.Error("unknown namespace for tag requirement on composing device type",
				slog.String("namespaceId", tr.TagRequirement.NamespaceID.HexString()),
				slog.String("namespaceName", tr.TagRequirement.NamespaceName),
				log.Path("source", tr))
			spec.addError(&UnknownNamespaceTagRequirementError{Requirement: tr.TagRequirement})
		} else if tr.TagRequirement.SemanticTag == nil && (tr.TagRequirement.SemanticTagID.Valid() || tr.TagRequirement.SemanticTagName != "") {
			tr.TagRequirement.SemanticTag = findTagRequirementTag(spec, tr.TagRequirement.Namespace, tr.TagRequirement.SemanticTagID, tr.TagRequirement.SemanticTagName, tr.TagRequirement)
			if tr.TagRequirement.SemanticTag == nil {
				slog.Error("unknown semantic tag for tag requirement on composing device type",
					slog.String("semanticTagId", tr.TagRequirement.SemanticTagID.HexString()),
					slog.String("semanticTagName", tr.TagRequirement.SemanticTagName),
					log.Path("source", tr))
				spec.addError(&UnknownTagRequirementError{Requirement: tr.TagRequirement})
			}
		}
	}
	return
}

func findDeviceTypeRequirementCluster(spec *Specification, id *matter.Number, name string, entity types.Entity) (cluster *matter.Cluster) {
	var ok bool
	if cluster, ok = spec.ClustersByID[id.Value()]; ok {
		if name != cluster.Name {
			slog.Error("Mismatch between cluster requirement ID and cluster name", slog.String("clusterId", id.HexString()), slog.String("clusterName", cluster.Name), slog.String("requirementName", name), log.Path("source", entity))
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
	if strings.EqualFold(name, "Base") {
		deviceType = spec.BaseDeviceType
		return
	}
	var ok bool
	if deviceType, ok = spec.DeviceTypesByID[id.Value()]; ok {
		if name != deviceType.Name {
			slog.Error("Mismatch between device type ID and device type name", slog.String("deviceTypeId", id.HexString()), slog.String("deviceTypeName", deviceType.Name), slog.String("requirementName", name), log.Path("source", entity))
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
	iu := make(idUniqueness[*matter.DeviceType])
	nu := make(nameUniqueness[*matter.DeviceType])
	for _, dt := range spec.DeviceTypes {
		iu.check(spec, dt.ID, dt)
		nu.check(spec, dt)
		referencedClusters := make(map[*matter.Cluster]struct{})
		buildReferencedClusters(spec.BaseDeviceType, referencedClusters)
		buildReferencedClusters(dt, referencedClusters)
		crcv := make(conformanceValidation)
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
			crcv.add(cr, cr.Conformance)
		}
		crcv.check(spec)
		ercv := make(conformanceValidation)
		for _, er := range dt.ElementRequirements {
			validateAccess(spec, er, er.Access)
			validateElementRequirement(spec, dt, er, referencedClusters)
			ercv.add(er, er.Conformance)
		}
		ercv.check(spec)
		cdercv := make(conformanceValidation)
		for _, der := range dt.ComposedDeviceTypeElementRequirements {
			if der.DeviceType == nil {
				continue
			}
			validateAccess(spec, der.ElementRequirement, der.ElementRequirement.Access)
			validateElementRequirement(spec, der.DeviceType, der.ElementRequirement, referencedClusters)
			cdercv.add(der.ElementRequirement, der.ElementRequirement.Conformance)
		}
		cdercv.check(spec)
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
	cluster := er.Cluster
	for cluster != nil {
		er.Entity, err = associateElementRequirementFromCluster(er, dt, cluster)
		if err != nil {
			return
		}
		if er.Entity != nil {
			break
		}
		if cluster.ParentCluster == nil {
			if cluster.ClusterClassification.Hierarchy != "Base" {
				cluster = spec.ClustersByName[cluster.ClusterClassification.Hierarchy]
			} else {
				break
			}
		} else {
			cluster = cluster.ParentCluster
		}
	}
	return
}

func associateElementRequirementFromCluster(er *matter.ElementRequirement, dt *matter.DeviceType, cluster *matter.Cluster) (entity types.Entity, err error) {
	switch er.Element {
	case types.EntityTypeAttribute:

		for _, a := range cluster.Attributes {
			if strings.EqualFold(a.Name, er.Name) {
				entity = a
				return
			}
		}
	case types.EntityTypeFeature:
		if cluster.Features == nil {
			slog.Error("Element Requirement references missing features", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), log.Path("source", er))
			return
		}
		for _, fb := range cluster.Features.Bits {
			f, ok := fb.(*matter.Feature)
			if !ok {
				continue
			}
			if f.Code == er.Name || strings.EqualFold(f.Name(), er.Name) {
				entity = f
				return
			}
		}
	case types.EntityTypeCommand:
		for _, cmd := range cluster.Commands {
			if strings.EqualFold(cmd.Name, er.Name) {
				entity = cmd
				return
			}
		}
	case types.EntityTypeCommandField:
		var command *matter.Command
		for _, cmd := range cluster.Commands {
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
				entity = f
				return
			}
		}
	case types.EntityTypeEvent:
		for _, e := range cluster.Events {
			if strings.EqualFold(e.Name, er.Name) {
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
