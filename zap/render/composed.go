package render

import (
	"fmt"
	"log/slog"
	"slices"

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

type hasEndpoint struct {
	hasEndpoint bool
	confidence  conformance.Confidence
}

func (he *hasEndpoint) Confidence() conformance.Confidence {
	return he.confidence
}

func (he *hasEndpoint) Value() any {
	return he.hasEndpoint
}

func (he *hasEndpoint) IsTrue() bool {
	return he.hasEndpoint
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

	var hasClient, hasServer hasEndpoint
	hasClient.confidence = conformance.ConfidenceDefinite
	hasServer.confidence = conformance.ConfidenceDefinite

	for cluster, requirements := range servers {
		if hasServer.confidence == conformance.ConfidenceDefinite && hasServer.hasEndpoint {
			break
		}
		slices.SortFunc(requirements, func(a *matter.DeviceTypeClusterRequirement, b *matter.DeviceTypeClusterRequirement) int {
			return a.Origin.Compare(b.Origin)
		})
		var firstPassState conformance.State
		firstPassState, err = getConformanceState(firstContext, requirements, elementRequirements)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			return
		}
		switch firstPassState {
		case conformance.StateMandatory:
			hasServer.confidence = conformance.ConfidenceDefinite
			hasServer.hasEndpoint = true
		case conformance.StateProvisional, conformance.StateOptional:
			hasServer.confidence = conformance.ConfidencePossible
			hasServer.hasEndpoint = true
		}
	}
	for cluster, requirements := range clients {
		if hasClient.confidence == conformance.ConfidenceDefinite && hasClient.hasEndpoint {
			break
		}
		slices.SortFunc(requirements, func(a *matter.DeviceTypeClusterRequirement, b *matter.DeviceTypeClusterRequirement) int {
			return a.Origin.Compare(b.Origin)
		})
		var firstPassState conformance.State
		firstPassState, err = getConformanceState(firstContext, requirements, elementRequirements)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			return
		}
		switch firstPassState {
		case conformance.StateMandatory:
			hasClient.confidence = conformance.ConfidenceDefinite
			hasClient.hasEndpoint = true
		case conformance.StateProvisional, conformance.StateOptional:
			hasClient.confidence = conformance.ConfidencePossible
			hasClient.hasEndpoint = true
		}
	}

	cxt := conformance.Context{
		Values: map[string]any{
			"Matter":            true,
			dc.DeviceType.Class: true,
			"Client":            &hasClient,
			"Server":            &hasServer,
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
			var conf conformance.ConformanceState
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
		var evalState conformance.ConformanceState
		evalState, err = req.ClusterRequirement.Conformance.Eval(cxt)
		if err != nil {
			err = fmt.Errorf("error evaluating conformance of cluster requirement %s: %w", req.ClusterRequirement.ClusterName, err)
			return
		}
		// ...unless the base device requirement is not mandatory, but it evaluates as mandatory in this context
		if state != conformance.StateMandatory && evalState.State == conformance.StateMandatory {
			switch evalState.Confidence {
			case conformance.ConfidenceDefinite:
				state = conformance.StateMandatory
			case conformance.ConfidencePossible:
				state = conformance.StateOptional
			}
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

func (p DeviceTypesPatcher) setEndpointCompositionElement(spec *spec.Specification, composition *matter.DeviceTypeComposition, parent *etree.Element) error {

	childDeviceTypes := composition.ComposedDeviceTypes[matter.DeviceTypeRequirementLocationChildEndpoint]
	endpointCompositionElement := parent.SelectElement("endpointComposition")
	if len(childDeviceTypes) == 0 {
		if endpointCompositionElement != nil {
			parent.RemoveChild(endpointCompositionElement)
		}
		return nil
	}
	if endpointCompositionElement == nil {
		endpointCompositionElement = etree.NewElement("endpointComposition")
		xml.AppendElement(parent, endpointCompositionElement, "scope")
	}
	endpointCompositionElement.Child = nil
	xml.SetOrCreateSimpleElement(endpointCompositionElement, "compositionType", "tree")
	slices.SortFunc(childDeviceTypes, func(a *matter.DeviceTypeComposition, b *matter.DeviceTypeComposition) int {
		return a.DeviceType.ID.Compare(b.DeviceType.ID)
	})
	for _, childDeviceType := range childDeviceTypes {
		var req *matter.DeviceTypeRequirement
		for _, r := range composition.DeviceTypeRequirements {
			if childDeviceType.DeviceType == r.DeviceType {
				req = r
			}
		}
		if req == nil {
			continue
		}
		endpoint := endpointCompositionElement.CreateElement("endpoint")
		conf := req.Conformance.ASCIIDocString()
		if conf != "" {
			endpoint.CreateAttr("conformance", conf)
		}
		cons := req.Constraint.ASCIIDocString(nil)
		if cons != "" {
			endpoint.CreateAttr("constraint", cons)
		}
		endpoint.CreateElement("deviceType").SetText(childDeviceType.DeviceType.ID.HexString())
		//endpoint.CreateElement("typeName").SetText(childDeviceType.DeviceType.Name)
		p.renderClusterIncludes(spec, endpoint, childDeviceType)
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
