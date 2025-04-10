package render

import (
	"log/slog"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (p DeviceTypesPatcher) setEndpointCompositionElement(spec *spec.Specification, cxt conformance.Context, deviceType *matter.DeviceType, parent *etree.Element) {

	composedDeviceTypes, composedDeviceTypeRequirements, composedDeviceTypeElementRequirements := p.buildComposedDeviceRequirements(deviceType, spec)

	endpointCompositionElement := parent.SelectElement("endpointComposition")
	if len(composedDeviceTypes) == 0 {
		if endpointCompositionElement != nil {
			parent.RemoveChild(endpointCompositionElement)
		}
		return
	}
	if endpointCompositionElement == nil {
		endpointCompositionElement = parent.CreateElement("endpointComposition")
	}
	xml.SetOrCreateSimpleElement(endpointCompositionElement, "compositionType", "tree")
	endpointElement := xml.SetOrCreateSimpleElement(endpointCompositionElement, "endpoint", "")
	endpointElement.RemoveAttr("conformance")
	endpointElement.CreateAttr("constraint", "min 1")
	xml.RemoveElements(endpointElement, "deviceType")
	for _, dt := range composedDeviceTypes {
		dte := endpointElement.CreateElement("deviceType")
		req := composedDeviceTypeRequirements[dt]
		dte.CreateAttr("id", dt.ID.HexString())
		dte.CreateAttr("name", dt.Name)
		renderConformance(spec, dt, deviceType, req.Conformance, dte)
		clusterRequirements := make([]*matter.ClusterRequirement, 0, len(dt.ClusterRequirements))
		for _, cr := range dt.ClusterRequirements {
			clusterRequirements = append(clusterRequirements, cr.Clone())
		}
		elementRequirements := make([]*matter.ElementRequirement, 0, len(dt.ElementRequirements))
		for _, dtr := range dt.ElementRequirements {
			elementRequirements = append(elementRequirements, dtr.Clone())
		}
		creq := composedDeviceTypeElementRequirements[dt]
		for _, cdtr := range creq {
			if cdtr.Element == types.EntityTypeUnknown { // Element Requirements with no feature are changing the qualities of the cluster requirement
				var matched bool
				for _, cr := range clusterRequirements {
					if cdtr.ClusterID.Valid() && !cdtr.ClusterID.Equals(cr.ClusterID) {
						continue
					}
					if !strings.EqualFold(cdtr.ClusterName, cr.ClusterName) {
						continue
					}
					cdtr.Quality.Inherit(cr.Quality)
					cr.Quality = cdtr.Quality
					if len(cdtr.Conformance) > 0 {
						cr.Conformance = cdtr.Conformance.CloneSet()
					}
					matched = true
					break
				}
				if !matched {
					slog.Warn("Composed device type requirement references unknown cluster",
						slog.String("deviceTypeId", deviceType.ID.HexString()),
						slog.String("deviceTypeName", deviceType.Name),
						slog.String("composedDeviceTypeId", dt.ID.HexString()),
						slog.String("composedDeviceTypeName", dt.Name),
						slog.String("clusterId", cdtr.ClusterID.HexString()),
						slog.String("clusterName", cdtr.ClusterName),
					)
				}
			} else {
				var matched bool
				for i, dtr := range elementRequirements {
					if cdtr.ClusterID.Valid() && !cdtr.ClusterID.Equals(dtr.ClusterID) {
						continue
					}
					if !strings.EqualFold(cdtr.ClusterName, dtr.ClusterName) {
						continue
					}
					if cdtr.Element != dtr.Element {
						continue
					}
					if !strings.EqualFold(cdtr.Name, dtr.Name) {
						continue
					}
					if !strings.EqualFold(cdtr.Field, dtr.Field) {
						continue
					}
					cdter := cdtr.ElementRequirement.Clone()
					if cdtr.Constraint == nil && dtr.Constraint != nil {
						cdter.Constraint = dtr.Constraint.Clone()
					}
					if len(cdtr.Conformance) == 0 && len(dtr.Conformance) > 0 {
						cdter.Conformance = dtr.Conformance.CloneSet()
					}
					cdter.Access.Inherit(dtr.Access)
					cdter.Quality.Inherit(dtr.Quality)
					elementRequirements[i] = cdter
					matched = true
					break
				}
				if !matched {
					elementRequirements = append(elementRequirements, &cdtr.ElementRequirement)
				}
			}

		}
		clusterRequirementsByID := p.buildClusterRequirements(spec, cxt, clusterRequirements, elementRequirements)
		p.setClustersElement(spec, cxt, dt, clusterRequirementsByID, dte)
	}
}

func (DeviceTypesPatcher) buildComposedDeviceRequirements(deviceType *matter.DeviceType, spec *spec.Specification) ([]*matter.DeviceType, map[*matter.DeviceType]*matter.DeviceTypeRequirement, map[*matter.DeviceType][]*matter.ComposedDeviceTypeRequirement) {
	var composedDeviceTypes []*matter.DeviceType
	composedDeviceTypeRequirements := make(map[*matter.DeviceType]*matter.DeviceTypeRequirement)
	for _, dtr := range deviceType.DeviceTypeRequirements {
		var dt *matter.DeviceType
		var ok bool
		if dtr.DeviceTypeID.Valid() {
			dt, ok = spec.DeviceTypesByID[dtr.DeviceTypeID.Value()]
			if !ok {
				slog.Warn("unknown composed device type ID", slog.String("deviceTypeId", dtr.DeviceTypeID.HexString()))
				continue
			}
		} else {
			dt, ok = spec.DeviceTypesByName[dtr.DeviceTypeName]
			if !ok {
				slog.Warn("unknown composed device type name", slog.String("deviceTypeName", dtr.DeviceTypeName))
				continue
			}
		}
		composedDeviceTypes = append(composedDeviceTypes, dt)
		_, ok = composedDeviceTypeRequirements[dt]
		if ok {
			slog.Warn("Duplicate composed device type requirement, ignoring...", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("deviceTypeName", deviceType.Name), slog.String("composedDeviceTypeId", dt.ID.HexString()), slog.String("composedDeviceTypeName", dt.Name))
			continue
		}
		composedDeviceTypeRequirements[dt] = dtr
	}
	composedDeviceTypeElementRequirements := make(map[*matter.DeviceType][]*matter.ComposedDeviceTypeRequirement)
	for _, cdtr := range deviceType.ComposedDeviceTypeRequirements {
		var dt *matter.DeviceType
		var ok bool
		if cdtr.DeviceTypeID.Valid() {
			dt, ok = spec.DeviceTypesByID[cdtr.DeviceTypeID.Value()]
			if !ok {
				slog.Warn("unknown composed device type ID", slog.String("deviceTypeId", cdtr.DeviceTypeID.HexString()))
				continue
			}
		} else {
			dt, ok = spec.DeviceTypesByName[cdtr.DeviceTypeName]
			if !ok {
				slog.Warn("unknown composed device type name", slog.String("deviceTypeName", cdtr.DeviceTypeName))
				continue
			}
		}
		_, ok = composedDeviceTypeRequirements[dt]
		if !ok {
			// Hunh; there's an element requirement for a device type that wasn't in the list of device types; we'll just pretend it was there and optional
			composedDeviceTypes = append(composedDeviceTypes, dt)
			composedDeviceTypeRequirements[dt] = &matter.DeviceTypeRequirement{DeviceTypeID: dt.ID.Clone(), DeviceTypeName: dt.Name, Conformance: conformance.Set{&conformance.Optional{}}}
		}
		composedDeviceTypeElementRequirements[dt] = append(composedDeviceTypeElementRequirements[dt], cdtr)
	}
	return composedDeviceTypes, composedDeviceTypeRequirements, composedDeviceTypeElementRequirements
}
