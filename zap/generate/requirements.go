package generate

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/find"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

type clusterRequirements struct {
	id                             *matter.Number
	name                           string
	requirementsFromDeviceType     []*matter.ClusterRequirement
	requirementsFromBaseDeviceType []*matter.ClusterRequirement
	elementRequirements            []*matter.ElementRequirement
}

func alternateDeviceTypeName(deviceType *matter.DeviceType) string {
	name := matter.Case(deviceType.Name)
	return "MA-" + strings.ToLower(name)
}

func (p DeviceTypesPatcher) applyDeviceTypeToElement(spec *spec.Specification, deviceType *matter.DeviceType, dte *etree.Element) (err error) {
	setDeviceTypeName(spec, dte, deviceType)
	xml.SetOrCreateSimpleElement(dte, "domain", "CHIP", "name")
	xml.SetOrCreateSimpleElement(dte, "typeName", deviceType.Name, "name", "domain")
	xml.SetOrCreateSimpleElement(dte, "profileId", "0x0103", "name", "domain", "typeName").CreateAttr("editable", "false")
	xml.SetOrCreateSimpleElement(dte, "deviceId", deviceType.ID.HexString(), "name", "domain", "typeName", "profileId").CreateAttr("editable", "false")
	xml.SetOrCreateSimpleElement(dte, "class", deviceType.Class, "name", "domain", "typeName", "profileId", "deviceId")
	xml.SetOrCreateSimpleElement(dte, "scope", deviceType.Scope, "name", "domain", "typeName", "profileId", "deviceId", "class")
	var hasClient, hasServer bool

	for _, cr := range deviceType.ClusterRequirements {
		if !cr.ClusterID.Valid() {
			continue
		}
		if conformance.IsMandatory(cr.Conformance) {
			switch cr.Interface {
			case matter.InterfaceClient:
				hasClient = true
			case matter.InterfaceServer:
				hasServer = true
			}
		}
		if hasClient && hasServer {
			break
		}
	}
	cxt := conformance.Context{
		Values: map[string]any{
			"Matter":         true,
			deviceType.Class: true,
			"Client":         hasClient,
			"Server":         hasServer,
		},
	}

	if p.fullEndpointComposition {
		p.setEndpointCompositionElement(spec, cxt, deviceType, dte)
	}
	clusterRequirementsByID := p.buildClusterRequirements(spec, cxt, deviceType.ClusterRequirements, deviceType.ElementRequirements)
	p.setClustersElement(spec, cxt, deviceType, clusterRequirementsByID, dte)
	return
}

func setDeviceTypeName(spec *spec.Specification, dte *etree.Element, deviceType *matter.DeviceType) {
	deviceTypeName := zap.DeviceTypeName(deviceType)
	doc, ok := spec.DocRefs[deviceType]
	if ok {
		errata := doc.Errata()
		if errata != nil {
			if errata.ZAP.DeviceTypeNames != nil {
				nameOverride, ok := errata.ZAP.DeviceTypeNames[deviceType.Name]
				if ok {
					deviceTypeName = nameOverride
					slog.Warn("device type source override doc", "path", deviceTypeName)
				}
			}
		}
	}
	existingName, ok := xml.ReadSimpleElement(dte, "name")
	if ok {
		if existingName == deviceTypeName {
			return
		}
		alternateName := alternateDeviceTypeName(deviceType)
		if existingName == alternateName {
			return
		}
	}
	xml.SetOrCreateSimpleElement(dte, "name", deviceTypeName)

}

func (p *DeviceTypesPatcher) buildClusterRequirements(spec *spec.Specification, conformanceContext conformance.Context, clusterReqs []*matter.ClusterRequirement, elementReqs []*matter.ElementRequirement) (clusterRequirementsByID map[uint64]*clusterRequirements) {
	clusterRequirementsByID = make(map[uint64]*clusterRequirements)
	for _, cr := range clusterReqs {
		if !cr.ClusterID.Valid() {
			continue
		}
		crr, ok := clusterRequirementsByID[cr.ClusterID.Value()]
		if !ok {
			crr = &clusterRequirements{id: cr.ClusterID, name: cr.ClusterName}
			clusterRequirementsByID[cr.ClusterID.Value()] = crr
		}

		crr.requirementsFromDeviceType = append(crr.requirementsFromDeviceType, cr)
	}
	for _, er := range elementReqs {
		if !er.ClusterID.Valid() {
			continue
		}
		crr, ok := clusterRequirementsByID[er.ClusterID.Value()]
		if !ok {
			// The spec has an element requirement for a cluster that wasn't required; probably a mistake, so
			// let's pretend the cluster was in the requirements, but optional
			crr = &clusterRequirements{id: er.ClusterID, name: er.ClusterName}
			clusterRequirementsByID[er.ClusterID.Value()] = crr
		}
		crr.elementRequirements = append(crr.elementRequirements, er)
	}
	for _, er := range spec.BaseDeviceType.ElementRequirements {
		if !er.ClusterID.Valid() {
			continue
		}
		crr, ok := clusterRequirementsByID[er.ClusterID.Value()]
		if ok {
			// If the Base Device Type has an element requirement for a cluster required by the Device Type, then include its element requirements too
			crr.elementRequirements = append(crr.elementRequirements, er)
		}
	}

	for _, cr := range spec.BaseDeviceType.ClusterRequirements {
		if !cr.ClusterID.Valid() {
			continue
		}
		crr, ok := clusterRequirementsByID[cr.ClusterID.Value()]
		if ok {
			// If a Base Device Type requirement specifies a cluster also specified by the device type, then include the relevant Base Device type requirement
			crr.requirementsFromBaseDeviceType = append(crr.requirementsFromBaseDeviceType, cr)
			slog.Debug("adding base device type cluster requirement", slog.String("cluster", cr.ClusterName))
		} else if !conformance.IsMandatory(cr.Conformance) {
			conf, confErr := cr.Conformance.Eval(conformanceContext)
			if confErr != nil {
				slog.Warn("Error evaluating conformance of cluster requirement", slog.String("clusterName", cr.ClusterName), slog.Any("error", confErr))
			} else if conf == conformance.StateMandatory {
				// If the Base Device Type has a requirement that is not plain Mandatory ("M"), but it returns Mandatory when evaulated, then include it
				crr = &clusterRequirements{id: cr.ClusterID, name: cr.ClusterName}
				crr.requirementsFromBaseDeviceType = append(crr.requirementsFromBaseDeviceType, cr)
				clusterRequirementsByID[cr.ClusterID.Value()] = crr
			}
		}
	}
	return
}

func (p DeviceTypesPatcher) setClustersElement(spec *spec.Specification, cxt conformance.Context, deviceType *matter.DeviceType, clusterRequirementsByID map[uint64]*clusterRequirements, parent *etree.Element) {
	clustersElement := parent.SelectElement("clusters")
	if len(deviceType.ClusterRequirements) == 0 {
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
			slog.Warn("missing cluster attribute on include", slog.String("deviceTypeId", deviceType.ID.HexString()))
			clustersElement.RemoveChild(include)
			continue
		}
		var cr *clusterRequirements
		var clusterId uint64
		for id, crs := range clusterRequirementsByID {
			if strings.EqualFold(ca.Value, crs.name) {
				cr = crs
				clusterId = id
			}
		}
		if cr == nil {
			slog.Debug("unknown cluster attribute on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", ca.Value))
			clustersElement.RemoveChild(include)
			continue
		}
		p.setIncludeAttributes(clustersElement, include, spec, deviceType, cr, cxt)
		delete(clusterRequirementsByID, clusterId)
	}
	for _, crs := range [][]*matter.ClusterRequirement{spec.BaseDeviceType.ClusterRequirements, deviceType.ClusterRequirements} {
		for _, cr := range crs {
			if !cr.ClusterID.Valid() {
				continue
			}
			crr, ok := clusterRequirementsByID[cr.ClusterID.Value()]
			if ok {
				p.setIncludeAttributes(clustersElement, nil, spec, deviceType, crr, cxt)
				delete(clusterRequirementsByID, cr.ClusterID.Value())
			}
		}
	}
}

func (p *DeviceTypesPatcher) setIncludeAttributes(clustersElement *etree.Element, include *etree.Element, spec *spec.Specification, deviceType *matter.DeviceType, cr *clusterRequirements, cxt conformance.Context) {
	var cluster *matter.Cluster
	var ok bool
	cluster, ok = spec.ClustersByID[cr.id.Value()]
	if !ok {
		cluster, ok = spec.ClustersByName[cr.name]
		if !ok {
			var alias string
			alias, ok = p.clusterAliases[cr.name]
			if ok {
				cluster, ok = spec.ClustersByName[alias]
				if !ok {
					slog.Warn("unknown cluster alias on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cr.name), slog.String("alias", alias))
					return
				}
			} else {
				slog.Warn("Removing unknown cluster on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cr.name))
				if include != nil {
					clustersElement.RemoveChild(include)
				}
				return
			}
		}
	}

	clusterDoc, ok := spec.DocRefs[cluster]
	if !ok {
		slog.Warn("unknown doc path on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name))
	}

	errata := errata.GetZAP(clusterDoc.Path.Relative)

	var server, client, clientLocked, serverLocked bool
	clientLocked = true
	serverLocked = true
	// Any requirements brought over from the Base Device Type should be overridden by device type requirements, so we parse in order
	for _, crs := range [][]*matter.ClusterRequirement{cr.requirementsFromBaseDeviceType, cr.requirementsFromDeviceType} {
		for _, cr := range crs {
			if conformance.IsBlank(cr.Conformance) {
				// If the device type has blank conformance, ignore it
				continue
			}
			conf, err := getClusterRequirementConformance(cxt, cr)
			if err != nil {
				slog.Warn("Error evaluating conformance of cluster requirement", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
				continue
			}

			switch conf {
			case conformance.StateOptional, conformance.StateProvisional:
				switch cr.Interface {
				case matter.InterfaceServer:
					server = false
					serverLocked = false
				case matter.InterfaceClient:
					client = false
					clientLocked = false
				}
			case conformance.StateMandatory:
				switch cr.Interface {
				case matter.InterfaceServer:
					server = true
					serverLocked = true
				case matter.InterfaceClient:
					client = true
					clientLocked = true
				}
			case conformance.StateDisallowed:
				switch cr.Interface {
				case matter.InterfaceServer:
					server = false
					serverLocked = true
				case matter.InterfaceClient:
					client = false
					clientLocked = true
				}
			default:
				slog.Warn("Unexpected conformance", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("conformance", conf.String()))
				return
			}
		}
	}

	if include == nil {
		include = etree.NewElement("include")
		include.CreateAttr("cluster", cluster.Name)
		xml.InsertElementByAttribute(clustersElement, include, "cluster")
	} else {
		include.CreateAttr("cluster", cluster.Name)
	}

	if !client && clientLocked { // Disallowed, so override whatever is there
		include.CreateAttr("client", strconv.FormatBool(client))
	} else {
		xml.SetNonexistentAttr(include, "client", strconv.FormatBool(client))
	}
	if !server && serverLocked { // Disallowed, so override whatever is there
		include.CreateAttr("server", strconv.FormatBool(server))
	} else {
		xml.SetNonexistentAttr(include, "server", strconv.FormatBool(server))
	}
	include.CreateAttr("clientLocked", strconv.FormatBool(clientLocked))
	include.CreateAttr("serverLocked", strconv.FormatBool(serverLocked))

	requiredAttributes := make(map[*matter.Field]*matter.ElementRequirement)
	requiredAttributeDefines := make(map[string]struct{})
	requiredCommands := make(map[*matter.Command]*matter.ElementRequirement)
	requiredCommandFields := make(map[*matter.Command]map[string]*matter.ElementRequirement)
	requiredEvents := make(map[*matter.Event]*matter.ElementRequirement)
	requiredFeatures := make(map[*matter.Feature]*matter.ElementRequirement)

	for _, er := range cr.elementRequirements {
		if conformance.IsZigbee(cluster, er.Conformance) {
			continue
		}
		conf, err := er.Conformance.Eval(cxt)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			continue
		}
		if conf == conformance.StateMandatory || conf == conformance.StateProvisional {
			switch er.Element {
			case types.EntityTypeFeature:
				var feature *matter.Feature
				for f := range cluster.Features.FeatureBits() {
					if strings.EqualFold(er.Name, f.Name()) {
						feature = f
						break
					}
				}
				if feature == nil {
					slog.Warn("unknown feature in element requirement", slog.String("deviceType", deviceType.Name), slog.String("clusterName", cluster.Name), slog.String("feature", er.Name))
					continue
				}
				cxt.Values[er.Name] = true
				requiredFeatures[feature] = er
			case types.EntityTypeAttribute:
				var attribute *matter.Field
				for _, a := range cluster.Attributes {
					if strings.EqualFold(a.Name, er.Name) {
						attribute = a
						break
					}
				}
				if attribute == nil {
					slog.Warn("unknown attribute in element requirement", slog.String("deviceType", deviceType.Name), slog.String("clusterName", cluster.Name), slog.String("attribute", er.Name))
					continue
				}
				requiredAttributes[attribute] = er
				requiredAttributeDefines[getDefine(attribute.Name, errata.ClusterDefinePrefix, errata)] = struct{}{}
				cxt.Values[er.Name] = true
			case types.EntityTypeCommand:
				var command *matter.Command
				for _, cmd := range cluster.Commands {
					if strings.EqualFold(cmd.Name, er.Name) {
						command = cmd
						break
					}
				}
				if command == nil {
					slog.Warn("unknown command in element requirement", slog.String("deviceType", deviceType.Name), slog.String("clusterName", cluster.Name), slog.String("attribute", er.Name))
					continue
				}
				requiredCommands[command] = er
				cxt.Values[er.Name] = true
			case types.EntityTypeEvent:
				var event *matter.Event
				for _, ev := range cluster.Events {
					if strings.EqualFold(ev.Name, er.Name) {
						event = ev
						break
					}
				}
				if event == nil {
					slog.Warn("unknown event in element requirement", slog.String("deviceType", deviceType.Name), slog.String("clusterName", cluster.Name), slog.String("attribute", er.Name))
					continue
				}
				requiredEvents[event] = er
				cxt.Values[er.Name] = true
			case types.EntityTypeCommandField:
				var command *matter.Command
				for _, a := range cluster.Commands {
					if strings.EqualFold(a.Name, er.Name) {
						command = a
						break
					}
				}
				if command == nil {
					slog.Warn("unknown command in element requirement", slog.String("deviceType", deviceType.Name), slog.String("clusterName", cluster.Name), slog.String("attribute", er.Name))
					continue
				}
				cf, ok := requiredCommandFields[command]
				if !ok {
					cf = make(map[string]*matter.ElementRequirement)
					requiredCommandFields[command] = cf
					cxt.Values[er.Name] = true
				}
				cf[er.Field] = er
			default:
				slog.Warn("Element requirement with unrecognized element type", slog.String("deviceType", deviceType.Name), slog.String("entityType", er.Element.String()))
			}
		}
	}

	fse := include.SelectElement("features")
	if len(requiredFeatures) > 0 {
		if fse == nil {
			fse = include.CreateElement("features")
		}
		fes := fse.SelectElements("feature")
		for _, fe := range fes {
			featureCode := fe.SelectAttr("code")
			featureName := fe.SelectAttr("name")
			if featureCode == nil && featureName == nil {
				slog.Warn("feature element missing code and name attributes", slog.String("deviceType", deviceType.Name), slog.String("clusterName", cluster.Name))
				include.RemoveChild(fe)
				continue
			}
			feature, er, required := find.FirstPairFunc(requiredFeatures, func(f *matter.Feature) bool {
				if featureCode != nil && strings.EqualFold(featureCode.Value, f.Code) {
					return true
				}
				return featureName != nil && strings.EqualFold(featureName.Value, f.Name())
			})
			if !required {
				include.RemoveChild(fe)
				continue
			}

			fe.CreateAttr("code", feature.Code)
			fe.CreateAttr("name", feature.Name())
			if p.generateFeatureXml {
				renderConformance(p.spec, deviceType, cluster, er.Conformance, fe)
			} else {
				removeConformance(fe)
			}
			delete(requiredFeatures, feature)
		}
		for feature, er := range requiredFeatures {
			fe := etree.NewElement("feature")
			fe.CreateAttr("code", feature.Code)
			fe.CreateAttr("name", feature.Name())
			if p.generateFeatureXml {
				renderConformance(p.spec, deviceType, cluster, er.Conformance, fe)
			}
			xml.InsertElementByAttribute(fse, fe, "code")
		}
	} else if fse != nil {
		include.RemoveChild(fse)
	}

	ras := include.SelectElements("requireAttribute")
	for _, ra := range ras {
		rat := ra.Text()
		_, required := requiredAttributeDefines[rat]
		if required {
			delete(requiredAttributeDefines, rat)
		} else {
			include.RemoveChild(ra)
		}
	}
	for ra := range requiredAttributeDefines {
		rae := etree.NewElement("requireAttribute")
		rae.SetText(ra)
		xml.InsertElementByName(include, rae, "")
	}
	rcs := include.SelectElements("requireCommand")
	for _, rc := range rcs {
		rct := rc.Text()
		cmd, _, required := find.FirstPairFunc(requiredCommands, func(command *matter.Command) bool { return strings.EqualFold(rct, command.Name) })
		if required {
			delete(requiredCommands, cmd)
		} else {
			include.RemoveChild(rc)
		}
	}
	for ra := range requiredCommands {
		rae := etree.NewElement("requireCommand")
		rae.SetText(ra.Name)
		xml.InsertElementByName(include, rae, "requireAttribute")
	}
	rcfs := include.SelectElements("requireCommandField")
	for _, rc := range rcfs {
		rcfc := rc.SelectElement("command")
		if rcfc == nil {
			include.RemoveChild(rc)
			continue
		}
		rct := rcfc.Text()
		cmd, fields, required := find.FirstPairFunc(requiredCommandFields, func(command *matter.Command) bool { return strings.EqualFold(rct, command.Name) })
		if required {
			delete(requiredCommandFields, cmd)
		} else {
			include.RemoveChild(rc)
			continue
		}
		rcffs := rc.SelectElements("field")
		for _, rcff := range rcffs {
			rcft := rc.Text()
			_, required := fields[rcft]
			if required {
				delete(fields, rcft)
			} else {
				rc.RemoveChild(rcff)
			}
		}
		for rcft := range fields {
			rae := etree.NewElement("field")
			rae.SetText(rcft)
			xml.InsertElementByName(rc, rae, "command")
		}
	}
	for rcfc, rcffs := range requiredCommandFields {
		rcfe := etree.NewElement("requireCommandField")
		rcfce := rcfe.CreateElement("command")
		rcfce.SetText(rcfc.Name)
		for rcff := range rcffs {
			rcfe.CreateElement("field").SetText(rcff)
		}
		xml.InsertElementByName(include, rcfe, "requireAttribute", "requireCommand")
	}
	res := include.SelectElements("requireEvent")
	for _, re := range res {
		ret := re.Text()
		ev, _, required := find.FirstPairFunc(requiredEvents, func(e *matter.Event) bool { return strings.EqualFold(e.Name, ret) })
		if required {
			delete(requiredEvents, ev)
		} else {
			include.RemoveChild(re)
		}
	}
	for ra := range requiredEvents {
		rae := etree.NewElement("requireEvent")
		rae.SetText(ra.Name)
		xml.InsertElementByName(include, rae, "requireAttribute", "requireCommand", "requireCommandField")
	}
}

func getClusterRequirementConformance(cxt conformance.Context, cr *matter.ClusterRequirement) (conf conformance.State, err error) {

	conf, err = cr.Conformance.Eval(cxt)
	if err != nil {
		return
	}
	if conf == conformance.StateMandatory || conf == conformance.StateProvisional {
		// If evaluating the conformance yields a mandatory or provisional result, take that
		return
	}
	// Otherwise, conformances might be returning Disallowed due to conditions we're unaware of
	// We'll check if the conformance is non-conditionally mandatory or disallowed, where we can be sure
	if conformance.IsMandatory(cr.Conformance) {
		conf = conformance.StateMandatory
		return
	}
	if conformance.IsDisallowed(cr.Conformance) {
		conf = conformance.StateDisallowed
		return
	}
	// All other conditions will be assumed optional
	conf = conformance.StateOptional
	return
}
