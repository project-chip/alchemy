package render

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/internal/find"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (p DeviceTypesPatcher) applyDeviceTypeToElement(spec *spec.Specification, deviceType *matter.DeviceType, dte *etree.Element, errata *errata.SDK) (err error) {
	setDeviceTypeName(spec, dte, deviceType, errata)
	xml.SetOrCreateSimpleElement(dte, "domain", "CHIP", "name")
	xml.SetOrCreateSimpleElement(dte, "typeName", errata.OverrideType(deviceType, deviceType.Name), "name", "domain")
	xml.SetOrCreateSimpleElement(dte, "profileId", "0x0103", "name", "domain", "typeName").CreateAttr("editable", "false")
	xml.SetOrCreateSimpleElement(dte, "deviceId", deviceType.ID.HexString(), "name", "domain", "typeName", "profileId").CreateAttr("editable", "false")
	xml.SetOrCreateSimpleElement(dte, "class", deviceType.Class, "name", "domain", "typeName", "profileId", "deviceId")
	xml.SetOrCreateSimpleElement(dte, "scope", deviceType.Scope, "name", "domain", "typeName", "profileId", "deviceId", "class")

	var composition *matter.DeviceTypeComposition
	composition, err = spec.ComposeDeviceType(deviceType)
	if err != nil {
		return
	}
	err = p.renderClusterIncludes(spec, dte, composition)
	if err != nil {
		return
	}
	if p.options.EndpointCompositionXML {
		err = p.setEndpointCompositionElement(spec, composition, dte)
		if err != nil {
			return
		}
	}
	return
}

func (p *DeviceTypesPatcher) renderClusterIncludes(spec *spec.Specification, parent *etree.Element, composition *matter.DeviceTypeComposition) (err error) {
	var composedClusters map[*matter.Cluster]*matter.ClusterComposition
	composedClusters, err = Compose(composition)
	if err != nil {
		return
	}
	clustersElement := parent.SelectElement("clusters")

	if len(composedClusters) == 0 {
		if clustersElement != nil {
			parent.RemoveChild(clustersElement)
		}
		return
	}
	if clustersElement == nil {
		clustersElement = parent.CreateElement("clusters")
	}
	for _, include := range clustersElement.SelectElements("include") {
		ca := include.SelectAttr("cluster")
		if ca == nil {
			slog.Warn("missing cluster attribute on include", slog.String("deviceTypeId", composition.DeviceType.ID.HexString()))
			clustersElement.RemoveChild(include)
			continue
		}
		var clusterComposition *matter.ClusterComposition
		for c, cc := range composedClusters {
			if strings.EqualFold(ca.Value, c.Name) {
				clusterComposition = cc
				delete(composedClusters, c)
				break
			}
		}
		if clusterComposition == nil {
			slog.Debug("unknown cluster attribute on include", slog.String("deviceTypeId", composition.DeviceType.ID.HexString()), slog.String("clusterName", ca.Value))
			clustersElement.RemoveChild(include)
			continue
		}
		err = p.renderClusterInclude(spec, clustersElement, include, composition.DeviceType, clusterComposition)
		if err != nil {
			return
		}
	}
	for _, clusterComposition := range composedClusters {
		err = p.renderClusterInclude(spec, clustersElement, nil, composition.DeviceType, clusterComposition)
	}
	if len(clustersElement.Child) == 0 {
		parent.RemoveChild(clustersElement)
	}
	return
}

func (p *DeviceTypesPatcher) renderClusterInclude(spec *spec.Specification,
	clustersElement *etree.Element,
	includeElement *etree.Element,
	deviceType *matter.DeviceType,
	clusterComposition *matter.ClusterComposition) (err error) {

	clusterDoc, ok := spec.DocRefs[clusterComposition.Cluster]
	if !ok {
		slog.Warn("unknown doc path on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", clusterComposition.Cluster.Name))
		return
	}

	if (clusterComposition.Server == conformance.StateUnknown || clusterComposition.Server == conformance.StateDisallowed) && (clusterComposition.Client == conformance.StateUnknown || clusterComposition.Client == conformance.StateDisallowed) {
		if includeElement != nil {
			clustersElement.RemoveChild(includeElement)
		}
		return
	}

	var server, client, clientLocked, serverLocked bool
	switch clusterComposition.Server {
	case conformance.StateMandatory, conformance.StateProvisional:
		server = true
		serverLocked = true
	case conformance.StateOptional, conformance.StateDeprecated:
		server = false
		serverLocked = false
	case conformance.StateDisallowed, conformance.StateUnknown:
		server = false
		serverLocked = true
		server = false
		serverLocked = true
	default:
		err = fmt.Errorf("unexpected conformance state %s", clusterComposition.Server.String())
		return
	}

	switch clusterComposition.Client {
	case conformance.StateMandatory, conformance.StateProvisional:
		client = true
		clientLocked = true
	case conformance.StateOptional, conformance.StateDeprecated:
		client = false
		clientLocked = false
	case conformance.StateDisallowed, conformance.StateUnknown:
		client = false
		clientLocked = true
	default:
		err = fmt.Errorf("unexpected conformance state %s", clusterComposition.Server.String())
		return
	}

	if serverLocked && !server && clientLocked && !client {
		// This is just completely disallowed; remove it if here
		if includeElement != nil {
			clustersElement.RemoveChild(includeElement)
			return
		}
	}

	errata := errata.GetSDK(clusterDoc.Path.Relative)

	if includeElement == nil {
		includeElement = etree.NewElement("include")
		includeElement.CreateAttr("cluster", clusterComposition.Cluster.Name)
		xml.InsertElementByAttribute(clustersElement, includeElement, "cluster")
	} else {
		includeElement.CreateAttr("cluster", clusterComposition.Cluster.Name)
	}

	if clientLocked { // Locked, so override whatever is there
		includeElement.CreateAttr("client", strconv.FormatBool(client))
	} else {
		xml.SetNonexistentAttr(includeElement, "client", strconv.FormatBool(client))
	}
	if serverLocked { // Locked, so override whatever is there
		includeElement.CreateAttr("server", strconv.FormatBool(server))
	} else {
		xml.SetNonexistentAttr(includeElement, "server", strconv.FormatBool(server))
	}
	includeElement.CreateAttr("clientLocked", strconv.FormatBool(clientLocked))
	includeElement.CreateAttr("serverLocked", strconv.FormatBool(serverLocked))

	err = p.renderClusterElementRequirements2(spec, includeElement, deviceType, clusterComposition, errata)
	return
}

func (p *DeviceTypesPatcher) renderClusterElementRequirements2(spec *spec.Specification,
	includeElement *etree.Element,
	deviceType *matter.DeviceType,
	clusterComposition *matter.ClusterComposition,
	errata *errata.SDK) (err error) {

	p.renderFeatureRequirements(spec, deviceType, includeElement, clusterComposition)
	p.renderAttributeRequirements(spec, deviceType, includeElement, clusterComposition, errata)
	p.renderCommandRequirements(spec, deviceType, includeElement, clusterComposition)
	p.renderEventRequirements(spec, deviceType, includeElement, clusterComposition)
	return
}

func (p *DeviceTypesPatcher) renderFeatureRequirements(spec *spec.Specification, deviceType *matter.DeviceType,
	root *etree.Element, clusterComposition *matter.ClusterComposition) {

	features := internal.ToMap(find.Filter(clusterComposition.Elements, func(ec *matter.ElementComposition) bool {
		_, ok := ec.ElementRequirement.Entity.(*matter.Feature)
		return ok
	}), func(ec *matter.ElementComposition) *matter.Feature {
		return ec.ElementRequirement.Entity.(*matter.Feature)
	})

	fse := root.SelectElement("features")
	if len(features) == 0 {
		if fse != nil {
			root.RemoveChild(fse)
		}
		return
	}
	if fse == nil {
		fse = root.CreateElement("features")
	}
	fes := fse.SelectElements("feature")
	for _, fe := range fes {
		featureCode := fe.SelectAttr("code")
		featureName := fe.SelectAttr("name")
		if featureCode == nil && featureName == nil {
			slog.Warn("feature element missing code and name attributes", slog.String("deviceType", deviceType.Name), slog.String("clusterName", clusterComposition.Cluster.Name))
			fse.RemoveChild(fe)
			continue
		}
		var feature *matter.Feature
		for f := range features {
			if featureCode != nil && strings.EqualFold(featureCode.Value, f.Code) {
				feature = f
				break
			}
			if featureName != nil && strings.EqualFold(featureName.Value, f.Name()) {
				feature = f
				break
			}
		}

		if feature == nil {
			slog.Warn("Unrecognized feature", slog.String("deviceType", deviceType.Name), slog.String("clusterName", clusterComposition.Cluster.Name), slog.Any("featureCode", featureCode), slog.Any("featureName", featureName))
			fse.RemoveChild(fe)
			continue
		}

		er := features[feature]
		fe.CreateAttr("code", feature.Code)
		fe.CreateAttr("name", feature.Name())
		if p.options.FeatureXML {
			renderConformance(spec, deviceType, er.ElementRequirement.Conformance, fe)
		} else {
			removeConformance(fe)
		}
		delete(features, feature)
	}
	for feature, er := range features {
		fe := etree.NewElement("feature")
		fe.CreateAttr("code", feature.Code)
		fe.CreateAttr("name", feature.Name())
		if p.options.FeatureXML {
			renderConformance(spec, deviceType, er.ElementRequirement.Conformance, fe)
		} else {
			removeConformance(fe)
		}
		xml.InsertElementByAttribute(fse, fe, "code")
	}
}

func (p *DeviceTypesPatcher) renderCommandRequirements(spec *spec.Specification, deviceType *matter.DeviceType,
	root *etree.Element, clusterComposition *matter.ClusterComposition) {

	commands := internal.ToMap(find.Filter(clusterComposition.Elements, func(ec *matter.ElementComposition) bool {
		_, ok := ec.ElementRequirement.Entity.(*matter.Command)

		return ok
	}), func(ec *matter.ElementComposition) *matter.Command {
		return ec.ElementRequirement.Entity.(*matter.Command)
	})

	rcs := root.SelectElements("requireCommand")
	for _, rc := range rcs {
		rct := rc.Text()
		command, ec, required := find.FirstPairFunc(commands, func(a *matter.Command) bool { return strings.EqualFold(rct, a.Name) })

		if !required || ec.State.State != conformance.StateMandatory {
			root.RemoveChild(rc)
		} else {
			delete(commands, command)
		}
	}
	for cmd, ec := range commands {
		if ec.State.State != conformance.StateMandatory {
			continue
		}
		rae := etree.NewElement("requireCommand")
		rae.SetText(cmd.Name)
		xml.InsertElementByName(root, rae, "requireAttribute")
	}
}

func (p *DeviceTypesPatcher) renderEventRequirements(spec *spec.Specification, deviceType *matter.DeviceType,
	root *etree.Element, clusterComposition *matter.ClusterComposition) {

	events := internal.ToMap(find.Filter(clusterComposition.Elements, func(ec *matter.ElementComposition) bool {
		_, ok := ec.ElementRequirement.Entity.(*matter.Event)

		return ok
	}), func(ec *matter.ElementComposition) *matter.Event {
		return ec.ElementRequirement.Entity.(*matter.Event)
	})

	rcs := root.SelectElements("requireEvent")
	for _, rc := range rcs {
		rct := rc.Text()
		event, ec, required := find.FirstPairFunc(events, func(a *matter.Event) bool { return strings.EqualFold(rct, a.Name) })

		if !required || ec.State.State != conformance.StateMandatory {
			root.RemoveChild(rc)
		} else {
			delete(events, event)
		}
	}
	for cmd, ec := range events {
		if ec.State.State != conformance.StateMandatory {
			continue
		}
		rae := etree.NewElement("requireEvent")
		rae.SetText(cmd.Name)
		xml.InsertElementByName(root, rae, "requireAttribute", "requireCommand")
	}
}

func (p *DeviceTypesPatcher) renderAttributeRequirements(spec *spec.Specification, deviceType *matter.DeviceType,
	root *etree.Element, clusterComposition *matter.ClusterComposition, errata *errata.SDK) {

	attributes := internal.ToMap(find.Filter(clusterComposition.Elements, func(ec *matter.ElementComposition) bool {
		f, ok := ec.ElementRequirement.Entity.(*matter.Field)
		if ok {
			return f.EntityType() == types.EntityTypeAttribute
		}
		return ok
	}), func(ec *matter.ElementComposition) *matter.Field {
		return ec.ElementRequirement.Entity.(*matter.Field)
	})

	defines := make(map[*matter.Field]string, len(attributes))
	for a := range attributes {
		defines[a] = getDefine(a.Name, errata.ClusterDefinePrefix, errata)
	}

	ras := root.SelectElements("requireAttribute")
	for _, ra := range ras {
		rat := ra.Text()
		attribute, _, required := find.FirstPairFunc(defines, func(a *matter.Field) bool { return strings.EqualFold(rat, defines[a]) })
		ec, ok := attributes[attribute]
		if !required || (ok && ec.State.State != conformance.StateMandatory) {
			root.RemoveChild(ra)
		} else {
			delete(defines, attribute)
		}
	}
	for _, ra := range defines {
		rae := etree.NewElement("requireAttribute")
		rae.SetText(ra)
		xml.InsertElementByName(root, rae, "")
	}
}

func setDeviceTypeName(spec *spec.Specification, dte *etree.Element, deviceType *matter.DeviceType, errata *errata.SDK) {
	deviceTypeName := zap.DeviceTypeName(deviceType, errata)

	existingName, ok := xml.ReadSimpleElement(dte, "name")
	if ok && existingName == deviceTypeName {
		return
	}
	xml.SetOrCreateSimpleElement(dte, "name", deviceTypeName)
}
