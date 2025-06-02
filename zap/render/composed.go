package render

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type ClusterComposition struct {
	Cluster *matter.Cluster
	Server  conformance.State
	Client  conformance.State

	Elements []*ElementComposition
}

type ElementComposition struct {
	ElementRequirement *matter.ElementRequirement
	State              conformance.State
}

func Compose(dc *matter.DeviceTypeComposition) (composedClusters map[*matter.Cluster]*matter.ClusterComposition, err error) {
	composedClusters = make(map[*matter.Cluster]*matter.ClusterComposition)

	servers := make(map[*matter.Cluster][]*matter.DeviceTypeClusterRequirement)
	clients := make(map[*matter.Cluster][]*matter.DeviceTypeClusterRequirement)
	elementRequirements := make(map[*matter.Cluster][]*matter.DeviceTypeElementRequirement)

	for _, cr := range dc.ClusterRequirements {
		switch cr.ClusterRequirement.Interface {
		case matter.InterfaceServer:
			servers[cr.ClusterRequirement.Cluster] = append(servers[cr.ClusterRequirement.Cluster], cr)
		case matter.InterfaceClient:
			clients[cr.ClusterRequirement.Cluster] = append(clients[cr.ClusterRequirement.Cluster], cr)
		}
	}
	for _, er := range dc.ElementRequirements {
		if _, ok := servers[er.ElementRequirement.Cluster]; !ok {
			continue
		}
		elementRequirements[er.ElementRequirement.Cluster] = append(elementRequirements[er.ElementRequirement.Cluster], er)
	}

	firstContext := conformance.Context{
		Values: map[string]any{
			"Matter":            true,
			dc.DeviceType.Class: true,
		},
	}

	var hasClient, hasServer bool

	slog.Info("composing device type", "name", dc.DeviceType.Name)
	for cluster, requirements := range servers {
		slices.SortFunc(requirements, func(a *matter.DeviceTypeClusterRequirement, b *matter.DeviceTypeClusterRequirement) int {
			return a.Origin.Compare(b.Origin)
		})
		slog.Info("sorted")
		for _, req := range requirements {
			slog.Info("sorted", "cluster", cluster.Name, "origin", req.Origin.String())
		}
		var firstPassState conformance.State
		firstPassState, err = getConformanceState(firstContext, requirements, elementRequirements)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			return
		}
		switch firstPassState {
		case conformance.StateMandatory, conformance.StateProvisional, conformance.StateOptional:
			hasServer = true
		}
	}
	for cluster, requirements := range clients {
		slices.SortFunc(requirements, func(a *matter.DeviceTypeClusterRequirement, b *matter.DeviceTypeClusterRequirement) int {
			return a.Origin.Compare(b.Origin)
		})
		if hasClient {
			continue
		}
		var firstPassState conformance.State
		firstPassState, err = getConformanceState(firstContext, requirements, elementRequirements)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			return
		}
		switch firstPassState {
		case conformance.StateMandatory, conformance.StateProvisional, conformance.StateOptional:
			hasClient = true
		}
	}

	cxt := conformance.Context{
		Values: map[string]any{
			"Matter":            true,
			dc.DeviceType.Class: true,
			"Client":            hasClient,
			"Server":            hasServer,
		},
	}
	for cluster, clusterRequirements := range servers {

		cc, ok := composedClusters[cluster]
		if !ok {
			cc = &matter.ClusterComposition{Cluster: cluster}
			composedClusters[cluster] = cc
		}

		cc.Server, err = getConformanceState(cxt, clusterRequirements, elementRequirements)

		if err != nil {
			return
		}

		ers := elementRequirements[cluster]
		if len(ers) == 0 {
			continue
		}
		elements := make(map[types.Entity][]*matter.DeviceTypeElementRequirement)
		for _, er := range ers {
			if er.ElementRequirement.Entity == nil {
				continue
			}
			elements[er.ElementRequirement.Entity] = append(elements[er.ElementRequirement.Entity], er)
		}

		for _, ers := range elements {
			slices.SortFunc(ers, func(a *matter.DeviceTypeElementRequirement, b *matter.DeviceTypeElementRequirement) int {
				return a.Origin.Compare(b.Origin)
			})
			req := ers[len(ers)-1].ElementRequirement
			if conformance.IsZigbee(req.Conformance) {
				continue
			}
			var conf conformance.State
			conf, err = req.Conformance.Eval(cxt)
			if err != nil {
				slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
				return
			}
			cc.Elements = append(cc.Elements, &matter.ElementComposition{ElementRequirement: req, State: conf})
		}
	}
	for cluster, clusterRequirements := range clients {

		cc, ok := composedClusters[cluster]
		if !ok {
			cc = &matter.ClusterComposition{Cluster: cluster}
			composedClusters[cluster] = cc
		}

		cc.Client, err = getConformanceState(cxt, clusterRequirements, elementRequirements)

		if err != nil {
			return
		}
	}
	return
}

func getConformanceState(cxt conformance.Context, clusterRequirements []*matter.DeviceTypeClusterRequirement, elementRequirements map[*matter.Cluster][]*matter.DeviceTypeElementRequirement) (state conformance.State, err error) {
	req := clusterRequirements[len(clusterRequirements)-1]

	if conformance.IsMandatory(req.ClusterRequirement.Conformance) {
		state = conformance.StateMandatory
	} else if conformance.IsDisallowed(req.ClusterRequirement.Conformance) {
		state = conformance.StateDisallowed
	} else {
		state = conformance.StateOptional
	}

	switch req.Origin {
	case matter.RequirementOriginBaseDeviceType, matter.RequirementOriginSubsetDeviceType:
		// Normally, we do not include clusters from the Base Device Type...
		var evalState conformance.State
		evalState, err = req.ClusterRequirement.Conformance.Eval(cxt)
		if err != nil {
			err = fmt.Errorf("error evaluating conformance of cluster requirement %s: %w", req.ClusterRequirement.ClusterName, err)
			return
		}
		// ...unless the base device requirement is not mandatory, but it evaluates as mandatory in this context
		if state != conformance.StateMandatory && evalState == conformance.StateMandatory {
			state = conformance.StateMandatory
			break
		}
		ers := elementRequirements[req.ClusterRequirement.Cluster]
		// ... or unless there are element requirement for this cluster that are NOT from the base device type
		var hasElements bool
		for _, er := range ers {
			switch er.Origin {
			case matter.RequirementOriginBaseDeviceType, matter.RequirementOriginSubsetDeviceType:
				continue
			}
			hasElements = true
			break
		}
		if hasElements {
			break
		}

		state = conformance.StateUnknown
		return
	}

	if req.DeviceTypeRequirement != nil {

		var deviceState conformance.State
		if conformance.IsMandatory(req.DeviceTypeRequirement.Conformance) {
			deviceState = conformance.StateMandatory
		} else if conformance.IsDisallowed(req.DeviceTypeRequirement.Conformance) {
			deviceState = conformance.StateDisallowed
		} else {
			deviceState = conformance.StateOptional
		}
		switch deviceState {
		case conformance.StateMandatory:
		case conformance.StateOptional, conformance.StateProvisional:
			// If the composed device type isn't mandatory, then the cluster is optional
			state = conformance.StateOptional
		case conformance.StateDeprecated:
			state = conformance.StateOptional
		case conformance.StateDisallowed:
			state = conformance.StateDisallowed
		}
	}

	return
}

func (p DeviceTypesPatcher) setEndpointCompositionElement(spec *spec.Specification, cxt conformance.Context, deviceType *matter.DeviceType, parent *etree.Element) error {

	composedDeviceTypes, composedDeviceTypeRequirements, composedDeviceTypeClusterRequirements, composedDeviceTypeElementRequirements, err := p.buildComposedDeviceRequirements(deviceType, spec)
	if err != nil {
		return err
	}
	endpointCompositionElement := parent.SelectElement("endpointComposition")
	if len(composedDeviceTypes) == 0 {
		if endpointCompositionElement != nil {
			parent.RemoveChild(endpointCompositionElement)
		}
		return nil
	}
	if endpointCompositionElement == nil {
		endpointCompositionElement = etree.NewElement("endpointComposition")
		xml.AppendElement(parent, endpointCompositionElement, "clusters")
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
		renderConformance(spec, dt, req.Conformance, dte)
		clusterRequirements := make([]*matter.ClusterRequirement, 0, len(dt.ClusterRequirements))
		for _, cr := range dt.ClusterRequirements {
			clusterRequirements = append(clusterRequirements, cr.Clone())
		}
		elementRequirements := make([]*matter.ElementRequirement, 0, len(dt.ElementRequirements))
		for _, dtr := range dt.ElementRequirements {
			elementRequirements = append(elementRequirements, dtr.Clone())
		}
		deviceTypeRequirements := make([]*matter.DeviceTypeRequirement, 0, len(dt.DeviceTypeRequirements))
		for _, dtr := range dt.DeviceTypeRequirements {
			deviceTypeRequirements = append(deviceTypeRequirements, dtr.Clone())
		}
		crcq := composedDeviceTypeClusterRequirements[dt]
		for _, cdtr := range crcq {
			var matched bool
			for _, cr := range clusterRequirements {
				if cdtr.ClusterRequirement.ClusterID.Valid() && !cdtr.ClusterRequirement.ClusterID.Equals(cr.ClusterID) {
					continue
				}
				if !strings.EqualFold(cdtr.ClusterRequirement.ClusterName, cr.ClusterName) {
					continue
				}
				if cdtr.ClusterRequirement.Interface != cr.Interface {
					continue
				}
				slog.Info("inherited cluster requirement on composed device type", matter.LogEntity("entity", cdtr.ClusterRequirement))
				cdtr.ClusterRequirement.Quality.Inherit(cr.Quality)
				cr.Quality = cdtr.ClusterRequirement.Quality
				if len(cdtr.ClusterRequirement.Conformance) > 0 {
					cr.Conformance = cdtr.ClusterRequirement.Conformance.CloneSet()
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
					slog.String("clusterId", cdtr.ClusterRequirement.ClusterID.HexString()),
					slog.String("clusterName", cdtr.ClusterRequirement.ClusterName),
				)
				for _, cr := range clusterRequirements {
					slog.Info("cluster requirement", matter.LogEntity("entity", cr))
				}
			}
		}
		creq := composedDeviceTypeElementRequirements[dt]
		for _, cdtr := range creq {
			var matched bool
			for i, dtr := range elementRequirements {
				if cdtr.ElementRequirement.ClusterID.Valid() && !cdtr.ElementRequirement.ClusterID.Equals(dtr.ClusterID) {
					continue
				}
				if !strings.EqualFold(cdtr.ElementRequirement.ClusterName, dtr.ClusterName) {
					continue
				}
				if cdtr.ElementRequirement.Element != dtr.Element {
					continue
				}
				if !strings.EqualFold(cdtr.ElementRequirement.Name, dtr.Name) {
					continue
				}
				if !strings.EqualFold(cdtr.ElementRequirement.Field, dtr.Field) {
					continue
				}
				cdter := cdtr.ElementRequirement.Clone()
				if cdtr.ElementRequirement.Constraint == nil && dtr.Constraint != nil {
					cdter.Constraint = dtr.Constraint.Clone()
				}
				if len(cdtr.ElementRequirement.Conformance) == 0 && len(dtr.Conformance) > 0 {
					cdter.Conformance = dtr.Conformance.CloneSet()
				}
				cdter.Access.Inherit(dtr.Access)
				cdter.Quality.Inherit(dtr.Quality)
				elementRequirements[i] = cdter
				matched = true
				break
			}
			if !matched {
				elementRequirements = append(elementRequirements, cdtr.ElementRequirement)
			}
		}
		/*clusterRequirementsByID := p.buildClusterRequirements(spec, cxt, dt, clusterRequirements, elementRequirements)
		p.setClustersElement(spec, cxt, dt, clusterRequirementsByID, dte)*/
	}
	return nil
}

func (DeviceTypesPatcher) buildComposedDeviceRequirements(deviceType *matter.DeviceType,
	spec *spec.Specification) (composedDeviceTypes []*matter.DeviceType,
	composedDeviceTypeRequirements map[*matter.DeviceType]*matter.DeviceTypeRequirement,
	composedDeviceTypeClusterRequirements map[*matter.DeviceType][]*matter.DeviceTypeClusterRequirement,
	composedDeviceTypeElementRequirements map[*matter.DeviceType][]*matter.DeviceTypeElementRequirement,
	err error) {
	composedDeviceTypeRequirements = make(map[*matter.DeviceType]*matter.DeviceTypeRequirement)
	for _, dtr := range deviceType.DeviceTypeRequirements {
		if dtr.DeviceType == nil {
			continue
		}
		dt := dtr.DeviceType
		composedDeviceTypes = append(composedDeviceTypes, dt)
		_, ok := composedDeviceTypeRequirements[dt]
		if ok {
			slog.Warn("Duplicate composed device type requirement, ignoring...", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("deviceTypeName", deviceType.Name), slog.String("composedDeviceTypeId", dt.ID.HexString()), slog.String("composedDeviceTypeName", dt.Name))
			continue
		}
		composedDeviceTypeRequirements[dt] = dtr
	}
	composedDeviceTypeClusterRequirements = make(map[*matter.DeviceType][]*matter.DeviceTypeClusterRequirement)
	for _, cdtr := range deviceType.ComposedDeviceTypeClusterRequirements {
		if cdtr.DeviceType == nil {
			continue
		}
		dt := cdtr.DeviceType
		_, ok := composedDeviceTypeRequirements[dt]
		if !ok {
			err = fmt.Errorf("unknown composed device type requirement for device type: %s", dt.Name)
			return
		}
		composedDeviceTypeClusterRequirements[dt] = append(composedDeviceTypeClusterRequirements[dt], cdtr)
	}
	composedDeviceTypeElementRequirements = make(map[*matter.DeviceType][]*matter.DeviceTypeElementRequirement)
	for _, cdtr := range deviceType.ComposedDeviceTypeElementRequirements {
		if cdtr.DeviceType == nil {
			continue
		}
		dt := cdtr.DeviceType
		_, ok := composedDeviceTypeRequirements[cdtr.DeviceType]
		if !ok {
			// Hunh; there's an element requirement for a device type that wasn't in the list of device types; we'll just pretend it was there and optional
			composedDeviceTypes = append(composedDeviceTypes, dt)
			composedDeviceTypeRequirements[dt] = &matter.DeviceTypeRequirement{DeviceTypeID: dt.ID.Clone(), DeviceTypeName: dt.Name, Conformance: conformance.Set{&conformance.Optional{}}}
		}
		composedDeviceTypeElementRequirements[cdtr.DeviceType] = append(composedDeviceTypeElementRequirements[cdtr.DeviceType], cdtr)
	}
	return
}
