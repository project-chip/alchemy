package dm

import (
	"bytes"
	"fmt"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func getDeviceTypePath(dmRoot string, path asciidoc.Path, deviceTypeName string) string {
	p := path.Base()
	file := strings.TrimSuffix(p, path.Ext())
	if len(deviceTypeName) > 0 {
		file += "-" + deviceTypeName
	}
	return filepath.Join(dmRoot, fmt.Sprintf("/device_types/%s.xml", file))
}

func renderDeviceType(deviceType *matter.DeviceType) (output string, err error) {
	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())
	c := x.CreateElement("deviceType")
	c.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	c.CreateAttr("xsi:schemaLocation", "types types.xsd devicetype devicetype.xsd")
	if deviceType.ID != nil && deviceType.ID.Valid() {
		c.CreateAttr("id", deviceType.ID.HexString())
	}
	c.CreateAttr("name", deviceType.Name)
	mostRecentRevision := deviceType.Revisions.MostRecent()
	if mostRecentRevision != nil {
		c.CreateAttr("revision", mostRecentRevision.Number.IntString())
	}

	revs := c.CreateElement("revisionHistory")
	for _, r := range deviceType.Revisions {
		if r.Number.Valid() {
			rev := revs.CreateElement("revision")
			rev.CreateAttr("revision", r.Number.IntString())
			rev.CreateAttr("summary", scrubDescription(r.Description))
		}
	}
	if deviceType.Class != "" || deviceType.Scope != "" {
		class := c.CreateElement("classification")
		if deviceType.SupersetOf != "" {
			class.CreateAttr("superset", deviceType.SupersetOf)
		}
		if deviceType.Class != "" {
			class.CreateAttr("class", strings.ToLower(deviceType.Class))
		}
		if deviceType.Scope != "" {
			class.CreateAttr("scope", strings.ToLower(deviceType.Scope))
		}
	}

	if len(deviceType.Conditions) > 0 {
		conditions := c.CreateElement("conditions")
		for _, condition := range deviceType.Conditions {
			cx := conditions.CreateElement("condition")
			cx.CreateAttr("name", condition.Feature)
			cx.CreateAttr("summary", scrubDescription(condition.Description))
		}
	}

	if len(deviceType.ConditionRequirements) > 0 {
		reqs := make(map[*matter.DeviceType][]*matter.ConditionRequirement)
		for _, cr := range deviceType.ConditionRequirements {
			if cr.Condition != nil && cr.DeviceType != nil {
				reqs[cr.DeviceType] = append(reqs[cr.DeviceType], cr)
			}
		}
		if len(reqs) > 0 {
			cre := c.CreateElement("conditionRequirements")
			for dt := range internal.IterateMapAlphabetically(reqs, func(dt *matter.DeviceType) string {
				return dt.Name
			}) {
				dte := cre.CreateElement("deviceType")
				if dt.ID.Valid() {
					dte.CreateAttr("id", dt.ID.HexString())
				}
				dte.CreateAttr("name", dt.Name)
				for _, cr := range reqs[dt] {
					cre := dte.CreateElement("conditionRequirement")
					cre.CreateAttr("name", cr.Condition.Feature)
					err = renderConformanceElement(cr.Conformance, cre, nil)
					if err != nil {
						return
					}
				}
			}
		}
	}

	if len(deviceType.ClusterRequirements) > 0 {
		cx := c.CreateElement("clusters")
		reqs := make([]*matter.ClusterRequirement, len(deviceType.ClusterRequirements))
		copy(reqs, deviceType.ClusterRequirements)
		slices.SortStableFunc(reqs, sortClusterRequirements)
		for _, cr := range reqs {
			clx := cx.CreateElement("cluster")
			clx.CreateAttr("id", cr.ClusterID.HexString())
			clx.CreateAttr("name", cr.ClusterName)
			switch cr.Interface {
			case matter.InterfaceClient:
				clx.CreateAttr("side", "client")
			case matter.InterfaceServer:
				clx.CreateAttr("side", "server")
			}
			renderQuality(clx, cr.Quality)
			err = renderConformanceElement(cr.Conformance, clx, nil)
			if err != nil {
				return
			}
			err = renderElementRequirements(deviceType, cr, deviceType.ElementRequirements, clx)
			if err != nil {
				return
			}

		}
	}

	type deviceTypeInstance struct {
		DeviceTypeName string
		Label          string
	}

	type deviceTypeRequirements struct {
		DeviceType          *matter.DeviceType
		deviceRequirements  []*matter.DeviceTypeRequirement
		clusterRequirements []*matter.DeviceTypeClusterRequirement
		elementRequirements map[string][]*matter.ElementRequirement
	}

	dtrs := make(map[deviceTypeInstance]*deviceTypeRequirements)

	for _, dr := range deviceType.DeviceTypeRequirements {
		if dr.DeviceType == nil {
			continue
		}

		key := deviceTypeInstance{DeviceTypeName: dr.DeviceType.Name, Label: ""}
		dtr, ok := dtrs[key]
		if !ok {
			dtr = &deviceTypeRequirements{DeviceType: dr.DeviceType}
			dtrs[key] = dtr
		}
		dtr.deviceRequirements = append(dtr.deviceRequirements, dr)
	}

	for _, cr := range deviceType.ComposedDeviceTypeClusterRequirements {
		if cr.DeviceType == nil {
			continue
		}
		if cr.ClusterRequirement.Cluster == nil {
			continue
		}

		label := cr.InstanceLabel
		if label == cr.DeviceType.Name {
			label = ""
		}
		key := deviceTypeInstance{DeviceTypeName: cr.DeviceType.Name, Label: label}
		dtr, ok := dtrs[key]
		if !ok {
			dtr = &deviceTypeRequirements{DeviceType: cr.DeviceType}
			dtrs[key] = dtr
		}
		dtr.clusterRequirements = append(dtr.clusterRequirements, cr)
	}

	for _, er := range deviceType.ComposedDeviceTypeElementRequirements {
		if er.DeviceType == nil {
			continue
		}
		if er.ElementRequirement.Cluster == nil {
			continue
		}
		label := er.InstanceLabel
		if label == er.DeviceType.Name {
			label = ""
		}
		key := deviceTypeInstance{DeviceTypeName: er.DeviceType.Name, Label: label}
		dtr, ok := dtrs[key]
		if !ok {
			dtr = &deviceTypeRequirements{DeviceType: er.DeviceType}
			dtrs[key] = dtr
		}
		
		foundCluster := false
		for _, cr := range dtr.clusterRequirements {
			if cr.ClusterRequirement.ClusterID.Equals(er.ElementRequirement.ClusterID) {
				foundCluster = true
				break
			}
		}
		if !foundCluster {
			cr := matter.NewClusterRequirement(er.DeviceType, er.Source())
			cr.ClusterID = er.ElementRequirement.ClusterID
			cr.ClusterName = er.ElementRequirement.ClusterName
			cr.Interface = matter.InterfaceServer // Default to server

			
			dtcr := matter.NewDeviceTypeClusterRequirement(er.DeviceType, cr, er.Source())
			dtcr.DeviceTypeID = er.DeviceTypeID
			dtcr.DeviceTypeName = er.DeviceTypeName
			dtcr.InstanceLabel = er.InstanceLabel
			
			dtr.clusterRequirements = append(dtr.clusterRequirements, dtcr)
		}

		if dtr.elementRequirements == nil {
			dtr.elementRequirements = make(map[string][]*matter.ElementRequirement)
		}
		cidStr := er.ElementRequirement.ClusterID.HexString()
		dtr.elementRequirements[cidStr] = append(dtr.elementRequirements[cidStr], er.ElementRequirement)

	}

	if len(dtrs) > 0 {
		cx := c.CreateElement("composedDeviceTypes")
		instances := make([]deviceTypeInstance, 0, len(dtrs))
		for key := range dtrs {
			instances = append(instances, key)
		}
		slices.SortStableFunc(instances, func(a, b deviceTypeInstance) int {
			cmp := strings.Compare(a.DeviceTypeName, b.DeviceTypeName)
			if cmp != 0 {
				return cmp
			}
			return strings.Compare(a.Label, b.Label)
		})
		for _, inst := range instances {
			dtr := dtrs[inst]
			dt := dtr.DeviceType
			dte := cx.CreateElement("deviceType")
			if dt.ID.Valid() {
				dte.CreateAttr("deviceTypeId", dt.ID.HexString())
			}
			dte.CreateAttr("deviceTypeName", dt.Name)
			
			var baseConformance conformance.Set
			for _, dr := range deviceType.DeviceTypeRequirements {
				if dr.DeviceType == dt {
					baseConformance = dr.Conformance
					break
				}
			}

			if len(dtr.deviceRequirements) > 0 {
				err = renderConformanceElement(dtr.deviceRequirements[0].Conformance, dte, nil)
				if err != nil {
					return
				}
				err = renderConstraintElement(dtr.deviceRequirements[0].Constraint, nil, dte, nil)
				if err != nil {
					return
				}
			} else if inst.Label != "" {
				if baseConformance != nil {
					err = renderConformanceElement(baseConformance, dte, nil)
					if err != nil {
						return
					}
				}

			}
			
			if len(dtr.clusterRequirements) > 0 {
				crx := dte.CreateElement("clusterRequirements")
				reqs := make([]*matter.DeviceTypeClusterRequirement, len(dtr.clusterRequirements))
				copy(reqs, dtr.clusterRequirements)
				slices.SortStableFunc(reqs, func(a, b *matter.DeviceTypeClusterRequirement) int {
					cmp := a.ClusterRequirement.ClusterID.Compare(b.ClusterRequirement.ClusterID)
					if cmp != 0 {
						return cmp
					}
					return a.ClusterRequirement.Interface.Compare(b.ClusterRequirement.Interface)
				})
				for _, cr := range reqs {
					clx := crx.CreateElement("cluster")
					clx.CreateAttr("id", cr.ClusterRequirement.ClusterID.HexString())
					clx.CreateAttr("name", cr.ClusterRequirement.ClusterName)
					if len(cr.ClusterRequirement.Conformance) > 0 {
						err = renderConformanceElement(cr.ClusterRequirement.Conformance, clx, nil)
						if err != nil {
							return
						}
					}
					renderQuality(clx, cr.ClusterRequirement.Quality)
					renderElementRequirements(dt, cr.ClusterRequirement, dtr.elementRequirements[cr.ClusterRequirement.ClusterID.HexString()], clx)
				}
			}
		}

	}

	x.Indent(2)

	var b bytes.Buffer
	_, err = x.WriteTo(&b)
	output = b.String()
	return
}

func sortClusterRequirements(a, b *matter.ClusterRequirement) int {
	cmp := a.ClusterID.Compare(b.ClusterID)
	if cmp != 0 {
		return cmp
	}
	return a.Interface.Compare(b.Interface)
}

type commandRequirement struct {
	command     *matter.Command
	requirement *matter.ElementRequirement
	fields      []*matter.ElementRequirement
}

func renderElementRequirements(deviceType *matter.DeviceType, cr *matter.ClusterRequirement, ers []*matter.ElementRequirement, clx *etree.Element) (err error) {
	erMap := make(map[types.EntityType][]*matter.ElementRequirement)
	for _, er := range ers {
		if er.ClusterID.Equals(cr.ClusterID) {
			erMap[er.Element] = append(erMap[er.Element], er)
		}
	}
	var featureRequirements []*matter.ElementRequirement
	var attributeRequirements []*matter.ElementRequirement
	var commandRequirements []*commandRequirement
	var eventRequirements []*matter.ElementRequirement
	for _, er := range ers {
		if er.ClusterID.Equals(cr.ClusterID) {
			switch er.Element {
			case types.EntityTypeFeature:
				featureRequirements = append(featureRequirements, er)
			case types.EntityTypeAttribute:
				attributeRequirements = append(attributeRequirements, er)
			case types.EntityTypeEvent:
				eventRequirements = append(eventRequirements, er)
			case types.EntityTypeCommand, types.EntityTypeCommandField:
				var cmd *matter.Command
				if er.Entity == nil {
				}
				switch entity := er.Entity.(type) {
				case *matter.Command:
					cmd = entity
				case *matter.Field:
					parent := entity.Parent()
					var isCommand bool
					cmd, isCommand = parent.(*matter.Command)
					if !isCommand {
						slog.Warn("Missing parent command on element requirement", slog.String("deviceType", deviceType.Name), slog.String("commandName", er.Name), slog.String("clusterName", cr.ClusterName))
					}
				case nil:
				default:
					err = fmt.Errorf("unexpected entity type on command or command field requirement: %T", entity)
				}
				if cmd == nil {
					slog.Warn("Unknown command on element requirement", slog.String("deviceType", deviceType.Name), slog.String("commandName", er.Name), slog.String("clusterName", cr.ClusterName))
					break
				}
				var cr *commandRequirement
				for _, c := range commandRequirements {
					if c.command == cmd {
						cr = c
						break
					}
				}
				if cr == nil {
					cr = &commandRequirement{command: cmd}
					commandRequirements = append(commandRequirements, cr)
				}
				switch er.Element {
				case types.EntityTypeCommand:
					cr.requirement = er
				case types.EntityTypeCommandField:
					cr.fields = append(cr.fields, er)
				}
			}
		}
	}
	if len(featureRequirements) > 0 {
		erx := clx.CreateElement("features")
		for _, fr := range featureRequirements {
			ex := erx.CreateElement("feature")
			switch feature := fr.Entity.(type) {
			case *matter.Feature:
				ex.CreateAttr("code", feature.Code)
			case nil:
				slog.Warn("Unknown feature on element requirement", slog.String("deviceType", deviceType.Name), slog.String("featureName", fr.Name), slog.String("clusterName", cr.ClusterName))
				continue
			}
			err = renderConformanceElement(fr.Conformance, ex, nil)
			if err != nil {
				return
			}
		}
	}
	if len(attributeRequirements) > 0 {
		erx := clx.CreateElement("attributes")
		for _, ar := range attributeRequirements {
			err = renderAttributeRequirement(deviceType, ar, erx)
			if err != nil {
				return
			}
		}
	}
	if len(commandRequirements) > 0 {
		erx := clx.CreateElement("commands")

		slices.SortStableFunc(commandRequirements, func(a, b *commandRequirement) int {
			cmp := a.command.Direction.Compare(b.command.Direction)
			if cmp != 0 {
				return cmp
			}
			return a.command.ID.Compare(b.command.ID)
		})
		for _, cr := range commandRequirements {

			ex := erx.CreateElement("command")

			ex.CreateAttr("id", cr.command.ID.HexString())

			ex.CreateAttr("name", cr.command.Name)
			if cr.command != nil && cr.requirement != nil {
				err = renderConformanceElement(cr.requirement.Conformance, ex, nil)
				if err != nil {
					return
				}
			}
			for _, fr := range cr.fields {
				fx := ex.CreateElement("field")
				fx.CreateAttr("name", fr.Field)
				err = renderConformanceElement(fr.Conformance, fx, nil)
				if err != nil {
					return
				}
			}
		}

	}
	if len(eventRequirements) > 0 {
		erx := clx.CreateElement("events")
		for _, er := range eventRequirements {
			switch entity := er.Entity.(type) {
			case *matter.Event:
				ex := erx.CreateElement("event")
				ex.CreateAttr("id", entity.ID.HexString())
				ex.CreateAttr("name", entity.Name)
				err = renderConformanceElement(er.Conformance, ex, nil)
				if err != nil {
					return
				}
			case nil:
				slog.Warn("Unknown event on element requirement", slog.String("deviceType", deviceType.Name), slog.String("eventName", er.Name), slog.String("clusterName", cr.ClusterName))
			}
		}
	}

	return
}

func renderAttributeRequirement(deviceType *matter.DeviceType, er *matter.ElementRequirement, parent *etree.Element) (err error) {
	var code string
	var attribute *matter.Field
	var dataType *types.DataType
	if er.Cluster != nil {
		for _, a := range er.Cluster.Attributes {
			if a.Name == er.Name {
				attribute = a
				dataType = a.Type
				break
			}
		}
	}
	if attribute != nil {
		code = attribute.ID.HexString()
	}
	ex := parent.CreateElement("attribute")
	ex.CreateAttr("code", code)
	ex.CreateAttr("name", er.Name)

	renderAttributeAccess(ex, er.Access)
	renderQuality(ex, er.Quality)
	err = renderConformanceElement(er.Conformance, ex, er.Cluster)
	if err != nil {
		return
	}
	err = renderConstraint(er.Constraint, dataType, ex, er.Cluster)
	return
}
