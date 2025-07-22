package dm

import (
	"bytes"
	"fmt"
	"log/slog"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
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
	if deviceType.ID != nil {
		c.CreateAttr("id", deviceType.ID.HexString())
	}
	c.CreateAttr("name", deviceType.Name)

	revs := c.CreateElement("revisionHistory")
	var latestRev uint64 = 0
	for _, r := range deviceType.Revisions {
		id := matter.ParseNumber(r.Number)
		if id.Valid() {
			rev := revs.CreateElement("revision")
			rev.CreateAttr("revision", id.IntString())
			rev.CreateAttr("summary", scrubDescription(r.Description))
			latestRev = max(id.Value(), latestRev)
		}
	}
	c.CreateAttr("revision", strconv.FormatUint(latestRev, 10))
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

	if len(deviceType.ClusterRequirements) > 0 {
		cx := c.CreateElement("clusters")
		reqs := make([]*matter.ClusterRequirement, len(deviceType.ClusterRequirements))
		copy(reqs, deviceType.ClusterRequirements)
		slices.SortStableFunc(reqs, func(a, b *matter.ClusterRequirement) int {
			cmp := a.ClusterID.Compare(b.ClusterID)
			if cmp != 0 {
				return cmp
			}
			return a.Interface.Compare(b.Interface)
		})
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
			err = renderElementRequirements(deviceType, cr, clx)
			if err != nil {
				return
			}

		}
	}

	x.Indent(2)

	var b bytes.Buffer
	_, err = x.WriteTo(&b)
	output = b.String()
	return
}

type commandRequirement struct {
	command     *matter.Command
	requirement *matter.ElementRequirement
	fields      []*matter.ElementRequirement
}

func renderElementRequirements(deviceType *matter.DeviceType, cr *matter.ClusterRequirement, clx *etree.Element) (err error) {
	erMap := make(map[types.EntityType][]*matter.ElementRequirement)
	for _, er := range deviceType.ElementRequirements {
		if er.ClusterID.Equals(cr.ClusterID) {
			erMap[er.Element] = append(erMap[er.Element], er)
		}
	}
	var featureRequirements []*matter.ElementRequirement
	var attributeRequirements []*matter.ElementRequirement
	var commandRequirements []*commandRequirement
	var eventRequirements []*matter.ElementRequirement
	for _, er := range deviceType.ElementRequirements {
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
