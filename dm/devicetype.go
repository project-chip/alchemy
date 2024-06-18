package dm

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/matter/types"
)

func getDeviceTypePath(sdkRoot string, path string) string {
	path = filepath.Base(path)
	return filepath.Join(sdkRoot, fmt.Sprintf("/data_model/device_types/%s.xml", strings.TrimSuffix(path, filepath.Ext(path))))
}

func renderDeviceType(cxt context.Context, doc *spec.Doc, deviceTypes []*matter.DeviceType) (output string, err error) {
	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())
	for _, deviceType := range deviceTypes {
		c := x.CreateElement("deviceType")
		c.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
		c.CreateAttr("xsi:schemaLocation", "types types.xsd devicetype devicetype.xsd")
		c.CreateAttr("id", deviceType.ID.HexString())
		c.CreateAttr("name", deviceType.Name)

		revs := c.CreateElement("revisionHistory")
		var latestRev uint64 = 0
		for _, r := range deviceType.Revisions {
			id := matter.ParseNumber(r.Number)
			if id.Valid() {
				rev := revs.CreateElement("revision")
				rev.CreateAttr("revision", id.IntString())
				rev.CreateAttr("summary", r.Description)
				latestRev = max(id.Value(), latestRev)
			}
		}
		c.CreateAttr("revision", strconv.FormatUint(latestRev, 10))
		class := c.CreateElement("classification")
		if deviceType.Superset != "" {
			class.CreateAttr("superset", deviceType.Superset)
		}
		class.CreateAttr("class", strings.ToLower(deviceType.Class))
		class.CreateAttr("scope", strings.ToLower(deviceType.Scope))

		if len(deviceType.Conditions) > 0 {
			conditions := c.CreateElement("conditions")
			for _, condition := range deviceType.Conditions {
				cx := conditions.CreateElement("condition")
				cx.CreateAttr("name", condition.Feature)
				cx.CreateAttr("summary", condition.Description)
			}
		}

		if len(deviceType.ClusterRequirements) > 0 {
			cx := c.CreateElement("clusters")
			reqs := make([]*matter.ClusterRequirement, len(deviceType.ClusterRequirements))
			copy(reqs, deviceType.ClusterRequirements)
			slices.SortFunc(reqs, func(a, b *matter.ClusterRequirement) int {
				cmp := a.ID.Compare(b.ID)
				if cmp != 0 {
					return cmp
				}
				return a.Interface.Compare(b.Interface)
			})
			for _, cr := range reqs {
				clx := cx.CreateElement("cluster")
				clx.CreateAttr("id", cr.ID.HexString())
				clx.CreateAttr("name", cr.ClusterName)
				switch cr.Interface {
				case matter.InterfaceClient:
					clx.CreateAttr("side", "client")
				case matter.InterfaceServer:
					clx.CreateAttr("side", "server")
				}
				renderQuality(clx, cr.Quality, matter.QualityAll)
				err = renderConformanceString(doc, deviceType, cr.Conformance, clx)
				if err != nil {
					return
				}
				err = renderElementRequirements(doc, deviceType, cr, clx)
				if err != nil {
					return
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

func renderElementRequirements(doc *spec.Doc, deviceType *matter.DeviceType, cr *matter.ClusterRequirement, clx *etree.Element) (err error) {
	erMap := make(map[types.EntityType][]*matter.ElementRequirement)
	for _, er := range deviceType.ElementRequirements {
		if er.ID.Equals(cr.ID) {
			erMap[er.Element] = append(erMap[er.Element], er)
		}
	}
	var featureRequirements []*matter.ElementRequirement
	var attributeRequirements []*matter.ElementRequirement
	var commandRequirements map[string][]*matter.ElementRequirement
	var eventRequirements []*matter.ElementRequirement
	for _, er := range deviceType.ElementRequirements {
		if er.ID.Equals(cr.ID) {
			switch er.Element {
			case types.EntityTypeFeature:
				featureRequirements = append(featureRequirements, er)
			case types.EntityTypeAttribute:
				attributeRequirements = append(attributeRequirements, er)
			case types.EntityTypeEvent:
				eventRequirements = append(eventRequirements, er)
			case types.EntityTypeCommand, types.EntityTypeCommandField:
				if commandRequirements == nil {
					commandRequirements = make(map[string][]*matter.ElementRequirement)
				}
				commandRequirements[er.Name] = append(commandRequirements[er.Name], er)
			}
		}
	}
	if len(featureRequirements) > 0 {
		erx := clx.CreateElement("features")
		for _, fr := range featureRequirements {
			ex := erx.CreateElement("feature")
			var code string

			featureCode := fr.Name
			if strings.ContainsRune(featureCode, ' ') {
				featureCode = matter.Case(featureCode)
			}

			if fr.Cluster != nil && fr.Cluster.Features != nil {
				for _, b := range fr.Cluster.Features.Bits {
					f := b.(*matter.Feature)
					if f.Code == featureCode {
						code = f.Code
						break
					}
				}
			}
			ex.CreateAttr("code", code)
			ex.CreateAttr("name", fr.Name)
			err = renderConformanceString(doc, deviceType, fr.Conformance, ex)
			if err != nil {
				return
			}
		}
	}
	if len(attributeRequirements) > 0 {
		erx := clx.CreateElement("attributes")
		for _, ar := range attributeRequirements {
			err = renderAttributeRequirement(doc, deviceType, ar, erx)
			if err != nil {
				return
			}
		}
	}
	if len(commandRequirements) > 0 {
		erx := clx.CreateElement("commands")

		for name, crs := range commandRequirements {
			var commandRequirement *matter.ElementRequirement
			var fieldRequirements []*matter.ElementRequirement
			for _, cr := range crs {
				switch cr.Element {
				case types.EntityTypeCommand:
					commandRequirement = cr
				case types.EntityTypeCommandField:
					fieldRequirements = append(fieldRequirements, cr)
				}
			}
			ex := erx.CreateElement("command")
			var code string
			if cr.Cluster != nil {
				for _, cmd := range cr.Cluster.Commands {
					if cmd.Name == name {
						code = cmd.ID.HexString()
						break
					}
				}
			}

			if code != "" {
				ex.CreateAttr("id", code)
			}
			ex.CreateAttr("name", name)
			if commandRequirement != nil {
				err = renderConformanceString(doc, deviceType, commandRequirement.Conformance, ex)
				if err != nil {
					return
				}
			}
			for _, fr := range fieldRequirements {
				fx := ex.CreateElement("field")
				fx.CreateAttr("name", fr.Field)
				err = renderConformanceString(doc, deviceType, fr.Conformance, fx)
				if err != nil {
					return
				}
			}
		}

	}
	if len(eventRequirements) > 0 {
		erx := clx.CreateElement("events")
		for _, er := range eventRequirements {
			ex := erx.CreateElement("event")
			var code string
			if er.Cluster != nil {
				for _, ev := range er.Cluster.Events {
					if ev.Name == er.Name {
						code = ev.ID.HexString()
						break
					}
				}
			}
			if code != "" {
				ex.CreateAttr("id", code)
			}
			ex.CreateAttr("name", er.Name)
			err = renderConformanceString(doc, deviceType, er.Conformance, ex)
			if err != nil {
				return
			}
		}
	}

	return
}

func renderAttributeRequirement(doc *spec.Doc, deviceType *matter.DeviceType, er *matter.ElementRequirement, parent *etree.Element) (err error) {
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
	renderQuality(ex, er.Quality, matter.QualityAll^matter.QualitySingleton)
	err = renderConformanceString(doc, deviceType, er.Conformance, ex)
	if err != nil {
		return
	}
	err = renderConstraint(er.Constraint, dataType, ex)
	return
}
