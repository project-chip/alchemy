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

type contextBoolean struct {
	value      bool
	confidence conformance.Confidence
}

func (he *contextBoolean) Confidence() conformance.Confidence {
	return he.confidence
}

func (he *contextBoolean) Result() any {
	return he.value
}

func (he *contextBoolean) IsTrue() bool {
	return he.value
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
			"Matter":            &contextBoolean{value: true, confidence: conformance.ConfidenceDefinite},
			"Zigbee":            &contextBoolean{value: false, confidence: conformance.ConfidenceImpossible},
			dc.DeviceType.Class: &contextBoolean{value: true, confidence: conformance.ConfidenceDefinite},
		},
	}

	var hasClient, hasServer contextBoolean
	hasClient.confidence = conformance.ConfidenceDefinite
	hasServer.confidence = conformance.ConfidenceDefinite

	for cluster, requirements := range servers {
		if hasServer.confidence == conformance.ConfidenceDefinite && hasServer.value {
			break
		}
		slices.SortFunc(requirements, func(a *matter.DeviceTypeClusterRequirement, b *matter.DeviceTypeClusterRequirement) int {
			return a.Origin.Compare(b.Origin)
		})
		var firstPassState conformance.ConformanceState
		firstPassState, err = getConformanceState(firstContext, requirements, elementRequirements)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			return
		}
		switch firstPassState.State {
		case conformance.StateMandatory:
			hasServer.confidence = firstPassState.Confidence
			hasServer.value = true
		case conformance.StateProvisional, conformance.StateOptional:
			hasServer.confidence = firstPassState.Confidence
			hasServer.value = true
		}
	}
	for cluster, requirements := range clients {
		if hasClient.confidence == conformance.ConfidenceDefinite && hasClient.value {
			break
		}
		slices.SortFunc(requirements, func(a *matter.DeviceTypeClusterRequirement, b *matter.DeviceTypeClusterRequirement) int {
			return a.Origin.Compare(b.Origin)
		})
		var firstPassState conformance.ConformanceState
		firstPassState, err = getConformanceState(firstContext, requirements, elementRequirements)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			return
		}
		switch firstPassState.State {
		case conformance.StateMandatory:
			hasClient.confidence = firstPassState.Confidence
			hasClient.value = true
		case conformance.StateProvisional, conformance.StateOptional:
			hasClient.confidence = conformance.ConfidencePossible
			hasClient.value = true
		}
	}

	cxt := conformance.Context{
		Values: map[string]any{
			"Matter":            &contextBoolean{value: true, confidence: conformance.ConfidenceDefinite},
			"Zigbee":            &contextBoolean{value: false, confidence: conformance.ConfidenceImpossible},
			dc.DeviceType.Class: &contextBoolean{value: true, confidence: conformance.ConfidenceDefinite},
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
			var conf conformance.ConformanceState
			conf, err = req.Conformance.Eval(cxt)
			if err != nil {
				slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", dc.DeviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
				return
			}
			if conf.Confidence == conformance.ConfidenceImpossible {
				continue
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

func getConformanceState(cxt conformance.Context, clusterRequirements []*matter.DeviceTypeClusterRequirement, elementRequirements map[*matter.Cluster][]*matter.DeviceTypeElementRequirement) (state conformance.ConformanceState, err error) {
	req := clusterRequirements[len(clusterRequirements)-1]

	state, err = req.ClusterRequirement.Conformance.Eval(cxt)
	if err != nil {
		err = fmt.Errorf("error evaluating conformance of cluster requirement %s: %w", req.ClusterRequirement.ClusterName, err)
		return

	}

	switch req.Origin {
	case matter.RequirementOriginBaseDeviceType, matter.RequirementOriginSubsetDeviceType:
		// Normally, we do not include clusters from the Base Device Type or subset device types...

		var hasElements bool
		ers := elementRequirements[req.ClusterRequirement.Cluster]

		for _, er := range ers {
			switch er.Origin {
			case matter.RequirementOriginBaseDeviceType, matter.RequirementOriginSubsetDeviceType:
				continue
			}
			hasElements = true
			break
		}
		// ... unless there are element requirements for this cluster that are NOT from the base device type or subset device type
		if hasElements {
			break
		}
		// ...or unless the base device/subset requirement is not mandatory, but it evaluates as mandatory in this context
		if !conformance.IsMandatory(req.ClusterRequirement.Conformance) && state.State == conformance.StateMandatory {
			break
		}
		state.State = conformance.StateUnknown
		state.Confidence = conformance.ConfidenceDefinite
		return
	}

	if req.DeviceTypeRequirement != nil {
		var deviceState conformance.ConformanceState
		deviceState, err = req.DeviceTypeRequirement.Conformance.Eval(cxt)
		if err != nil {
			err = fmt.Errorf("error evaluating conformance of device requirement %s: %w", req.ClusterRequirement.ClusterName, err)
			return
		}
		switch deviceState.State {
		case conformance.StateMandatory:
			switch deviceState.Confidence {
			case conformance.ConfidencePossible, conformance.ConfidenceImpossible:
				state.Confidence = deviceState.Confidence
			}
		case conformance.StateOptional, conformance.StateProvisional:
			// If the composed device type isn't mandatory, then the cluster is optional
			state.State = deviceState.State
			switch deviceState.Confidence {
			case conformance.ConfidencePossible, conformance.ConfidenceImpossible:
				state.Confidence = deviceState.Confidence
			}
		case conformance.StateDeprecated:
			state.State = conformance.StateOptional
		case conformance.StateDisallowed:
			state.State = conformance.StateDisallowed
			state.Confidence = deviceState.Confidence
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
		dte := endpoint.CreateElement("deviceType")
		dte.CreateAttr("id", childDeviceType.DeviceType.ID.HexString())
		dte.CreateAttr("name", childDeviceType.DeviceType.Name)
		p.renderClusterIncludes(spec, dte, childDeviceType)
	}
	return nil
}
