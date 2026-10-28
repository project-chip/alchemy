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
		cr.Cluster, _ = resolveEntityReference(spec, &cr.ClusterRef, entityFinder, spec.findClusterByID, spec.findClusterByName, cr)
	}
	for _, er := range dt.ElementRequirements {
		if er.Cluster != nil {
			continue
		}
		er.Cluster, _ = resolveEntityReference(spec, &er.ClusterRef, entityFinder, spec.findClusterByID, spec.findClusterByName, er)
	}
	for _, er := range dt.ElementRequirements {
		err = associateElementRequirement(spec, dt, er, entityFinder)
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
			cr.Condition, _ = resolveEntityReference(spec, &cr.ConditionRef, entityFinder, nil, findConditionByName(cr.DeviceType), cr)
			if cr.Condition != nil && !cr.DeviceType.Contains(cr.Condition) {
				slog.Error("Condition requirement on device type refers to unknown condition", slog.String("condition", cr.Condition.Feature), slog.String("deviceType", cr.DeviceType.Name), log.Path("source", cr))
				cr.Condition = nil
			}
		}
	}
	for _, tr := range dt.TagRequirements {
		tr.Namespace, _ = resolveEntityReference(spec, &tr.NamespaceRef, entityFinder, spec.findNamespaceByID, spec.findNamespaceByName, tr)
		if tr.Namespace != nil && tr.SemanticTag == nil {
			tr.SemanticTag, _ = resolveEntityReference(spec, &tr.SemanticTagRef, entityFinder, findTagByID(tr.Namespace), findTagByName(tr.Namespace), tr)
		}
	}
	for _, cr := range dt.ComposedDeviceTypeClusterRequirements {
		if cr.ClusterRequirement.Cluster == nil {
			var ok bool
			cr.ClusterRequirement.Cluster, ok = resolveEntityReference(spec, &cr.ClusterRequirement.ClusterRef, entityFinder, spec.findClusterByID, spec.findClusterByName, cr)
			if !ok {
				slog.Error("unknown cluster ID for cluster requirement on composing device type",
					slog.String("clusterId", cr.ClusterRequirement.ClusterRef.ID.HexString()),
					slog.String("clusterName", cr.ClusterRequirement.ClusterRef.Name),
					slog.String("deviceType", dt.Name),
					log.Path("source", cr.ClusterRequirement))
			}
		}
		if cr.DeviceType == nil {
			cr.DeviceType = findDeviceTypeRequirementDeviceType(spec, cr.DeviceTypeRef, cr.ClusterRequirement, entityFinder)
			if cr.DeviceType != nil {
				if dtr, ok := deviceTypes[cr.DeviceType]; !ok && cr.DeviceType != spec.RootNodeDeviceType && cr.DeviceType != spec.BaseDeviceType {
					slog.Error("Cluster requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", cr.DeviceTypeRef.ID.HexString()),
						slog.String("deviceTypeName", cr.DeviceTypeRef.Name),
						slog.String("deviceType", dt.Name),
						log.Path("source", cr.ClusterRequirement))
				} else {
					cr.DeviceType = cr.DeviceType
					cr.DeviceTypeRequirement = dtr
				}
			}
		}
	}
	for _, er := range dt.ComposedDeviceTypeElementRequirements {
		if er.ElementRequirement.Cluster == nil {
			er.ElementRequirement.Cluster, _ = resolveEntityReference(spec, &er.ElementRequirement.ClusterRef, entityFinder, spec.findClusterByID, spec.findClusterByName, er.ElementRequirement)
			if er.ElementRequirement.Cluster == nil {
				slog.Error("unknown cluster ID for element requirement on composing device type",
					slog.String("clusterId", er.ElementRequirement.ClusterRef.ID.HexString()),
					slog.String("clusterName", er.ElementRequirement.ClusterRef.Name),
					slog.String("deviceType", dt.Name),
					log.Path("source", er.ElementRequirement))
			}
		}
		if er.DeviceType == nil {
			er.DeviceType = findDeviceTypeRequirementDeviceType(spec, er.DeviceTypeRef, er.ElementRequirement, entityFinder)
			if er.DeviceType != nil {
				if dtr, ok := deviceTypes[er.DeviceType]; !ok && er.DeviceType != spec.RootNodeDeviceType && er.DeviceType != spec.BaseDeviceType {
					slog.Error("Element requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", er.DeviceTypeRef.ID.HexString()),
						slog.String("deviceTypeName", er.DeviceTypeRef.Name),
						slog.String("deviceType", dt.Name),
						log.Path("source", er.ElementRequirement))
				} else {
					er.DeviceTypeRequirement = dtr
				}
			}
		}
	}
	for _, er := range dt.ComposedDeviceTypeElementRequirements {
		if er.DeviceType == nil {
			continue
		}
		err = associateElementRequirement(spec, er.DeviceType, er.ElementRequirement, entityFinder)
		if err != nil {
			return
		}
	}
	for _, tr := range dt.ComposedDeviceTagRequirements {
		if tr.DeviceType == nil {
			referencedDeviceType := findDeviceTypeRequirementDeviceType(spec, tr.DeviceTypeRef, tr, entityFinder)
			if referencedDeviceType == nil {
				slog.Error("unknown device type ID for cluster requirement on composing device type",
					slog.String("deviceTypeId", tr.DeviceTypeRef.ID.HexString()),
					slog.String("deviceTypeName", tr.DeviceTypeRef.Name),
					slog.String("deviceType", dt.Name),
					log.Path("source", tr))
			} else {
				if dtr, ok := deviceTypes[referencedDeviceType]; !ok && referencedDeviceType != spec.RootNodeDeviceType && referencedDeviceType != spec.BaseDeviceType {
					slog.Error("Tag requirement on composing device type refers to unincluded device type",
						slog.String("deviceTypeId", tr.DeviceTypeRef.ID.HexString()),
						slog.String("deviceTypeName", tr.DeviceTypeRef.Name),
						slog.String("deviceType", dt.Name),
						log.Path("source", tr))
				} else {
					tr.DeviceType = referencedDeviceType
					tr.DeviceTypeRequirement = dtr
				}
			}
		}
		tr.TagRequirement.Namespace, _ = resolveEntityReference(spec, &tr.TagRequirement.NamespaceRef, entityFinder, spec.findNamespaceByID, spec.findNamespaceByName, tr.TagRequirement)
		if tr.TagRequirement.Namespace == nil {
			slog.Error("unknown namespace for tag requirement on composing device type",
				slog.String("namespaceId", tr.TagRequirement.NamespaceRef.ID.HexString()),
				slog.String("namespaceName", tr.TagRequirement.NamespaceRef.Name),
				log.Path("source", tr))
		} else if tr.TagRequirement.SemanticTag == nil {
			tr.TagRequirement.SemanticTag, _ = resolveEntityReference(spec, &tr.TagRequirement.SemanticTagRef, entityFinder, findTagByID(tr.TagRequirement.Namespace), findTagByName(tr.TagRequirement.Namespace), tr.TagRequirement)
			if tr.TagRequirement.SemanticTag == nil && !tr.TagRequirement.SemanticTagRef.Empty() {
				slog.Error("unknown semantic tag for tag requirement on composing device type",
					slog.String("semanticTagId", tr.TagRequirement.SemanticTagRef.ID.HexString()),
					slog.String("semanticTagName", tr.TagRequirement.SemanticTagRef.Name),
					log.Path("source", tr))
			}
		}
	}
	return
}

func findDeviceTypeRequirementDeviceType[SourceT types.Entity](spec *Specification, ref matter.EntityReference, entity SourceT, entityFinder entityFinder) (deviceType *matter.DeviceType) {
	if strings.EqualFold(ref.Name, "Base") {
		deviceType = spec.BaseDeviceType
		return
	}
	deviceType, _ = resolveEntityReference(spec, &ref, entityFinder, spec.findDeviceTypeByID, spec.findDeviceTypeByName, entity)
	return
}

func associateElementRequirement(spec *Specification, dt *matter.DeviceType, er *matter.ElementRequirement, entityFinder entityFinder) (err error) {
	if er.Cluster == nil {
		return
	}

	cluster := er.Cluster
	for cluster != nil {
		err = resolveElementReference(spec, er, entityFinder, cluster, dt)
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
