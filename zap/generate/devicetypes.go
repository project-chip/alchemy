package generate

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

var utilityDevicesMask uint64 = 0xFF000000

type DeviceTypesPatcher struct {
	sdkRoot string
	spec    *spec.Specification
}

func NewDeviceTypesPatcher(sdkRoot string, spec *spec.Specification) *DeviceTypesPatcher {
	return &DeviceTypesPatcher{sdkRoot: sdkRoot, spec: spec}
}

func (p DeviceTypesPatcher) Name() string {
	return "Patching device types"
}

func (p DeviceTypesPatcher) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (p DeviceTypesPatcher) Process(cxt context.Context, inputs []*pipeline.Data[[]*matter.DeviceType]) (outputs []*pipeline.Data[[]byte], err error) {

	deviceTypesByID := make(map[uint64]*matter.DeviceType)
	deviceTypesByName := make(map[string]*matter.DeviceType)
	for _, input := range inputs {
		for _, dt := range input.Content {
			if dt.ID.Valid() {
				deviceTypesByID[dt.ID.Value()] = dt
			} else {
				deviceTypesByName[matterDeviceTypeName(dt)] = dt
			}
		}
	}

	deviceTypesXMLPath := filepath.Join(p.sdkRoot, "/src/app/zap-templates/zcl/data-model/chip/matter-devices.xml")

	var deviceTypesXML []byte
	deviceTypesXML, err = os.ReadFile(deviceTypesXMLPath)
	if err != nil {
		return
	}

	xml := etree.NewDocument()
	err = xml.ReadFromBytes(deviceTypesXML)
	if err != nil {
		return
	}

	configurator := xml.SelectElement("configurator")
	if configurator == nil {
		err = fmt.Errorf("missing configurator element in %s", deviceTypesXMLPath)
		return
	}

	deviceTypeElements := configurator.SelectElements("deviceType")
	for _, deviceTypeElement := range deviceTypeElements {
		var deviceType *matter.DeviceType
		deviceIDElement := deviceTypeElement.SelectElement("deviceId")
		if deviceIDElement != nil {
			deviceTypeIDText := deviceIDElement.Text()
			deviceTypeID := matter.ParseNumber(deviceTypeIDText)
			if deviceTypeID.Valid() {
				if (deviceTypeID.Value() & utilityDevicesMask) == utilityDevicesMask {
					// Exception for the all clusters app, etc
					continue
				}
				deviceType = deviceTypesByID[deviceTypeID.Value()]
				if deviceType != nil {
					delete(deviceTypesByID, deviceTypeID.Value())
				}
			} else {
				slog.Warn("invalid deviceId", "text", deviceTypeID.Text())
			}

		}
		if deviceType == nil {
			deviceTypeElement := deviceTypeElement.SelectElement("typeName")
			if deviceTypeElement == nil {
				slog.Warn("missing deviceId and typeName elements")
				continue
			}
			deviceTypeIDText := deviceTypeElement.Text()
			deviceType = deviceTypesByName[deviceTypeIDText]
			if deviceType != nil {
				delete(deviceTypesByName, deviceTypeIDText)
			}
		}
		if deviceType != nil {
			applyDeviceTypeToElement(p.spec, deviceType, deviceTypeElement)
		} else {
			configurator.RemoveChild(deviceTypeElement)
		}
	}

	for _, dt := range deviceTypesByID {
		slog.Info("missing device type", slog.String("name", dt.Name))
		applyDeviceTypeToElement(p.spec, dt, configurator.CreateElement("deviceType"))
	}

	for _, dt := range deviceTypesByName {
		slog.Info("missing device type", slog.String("name", dt.Name))
		applyDeviceTypeToElement(p.spec, dt, configurator.CreateElement("deviceType"))
	}

	var out []byte
	xml.Indent(4)
	xml.WriteSettings.CanonicalEndTags = true
	out, err = xml.WriteToBytes()
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData[[]byte](deviceTypesXMLPath, out))
	return
}

type clusterRequirements struct {
	name                    string
	clusterRequirements     []*matter.ClusterRequirement
	baseClusterRequirements []*matter.ClusterRequirement
	elementRequirements     []*matter.ElementRequirement
}

func applyDeviceTypeToElement(spec *spec.Specification, deviceType *matter.DeviceType, dte *etree.Element) (err error) {
	xml.SetOrCreateSimpleElement(dte, "name", zap.DeviceTypeName(deviceType))
	xml.SetOrCreateSimpleElement(dte, "domain", "CHIP")
	xml.SetOrCreateSimpleElement(dte, "typeName", matterDeviceTypeName(deviceType))
	xml.SetOrCreateSimpleElement(dte, "profileId", "0x0103").CreateAttr("editable", "false")
	xml.SetOrCreateSimpleElement(dte, "deviceId", deviceType.ID.HexString()).CreateAttr("editable", "false")
	xml.SetOrCreateSimpleElement(dte, "class", deviceType.Class)
	xml.SetOrCreateSimpleElement(dte, "scope", deviceType.Scope)
	clustersElement := dte.SelectElement("clusters")
	if len(deviceType.ClusterRequirements) == 0 {
		if clustersElement != nil {
			dte.RemoveChild(clustersElement)
		}
		return
	}
	if clustersElement == nil {
		clustersElement = dte.CreateElement("clusters")
	}
	clusterRequirementsByName := make(map[string]*clusterRequirements)
	for _, cr := range deviceType.ClusterRequirements {
		name := strings.ToLower(cr.ClusterName)
		crr, ok := clusterRequirementsByName[name]
		if !ok {
			crr = &clusterRequirements{name: cr.ClusterName}
			clusterRequirementsByName[name] = crr
		}
		crr.clusterRequirements = append(crr.clusterRequirements, cr)
	}
	for _, cr := range spec.BaseDeviceType.ClusterRequirements {
		name := strings.ToLower(cr.ClusterName)
		crr, ok := clusterRequirementsByName[name]
		if !ok {
			crr = &clusterRequirements{name: cr.ClusterName}
			clusterRequirementsByName[name] = crr
		}
		crr.baseClusterRequirements = append(crr.baseClusterRequirements, cr)
		slog.Debug("adding base device type cluster requirement", slog.String("cluster", cr.ClusterName))
	}
	for _, ers := range [][]*matter.ElementRequirement{deviceType.ElementRequirements, spec.BaseDeviceType.ElementRequirements} {
		for _, er := range ers {
			name := strings.ToLower(er.ClusterName)
			crr, ok := clusterRequirementsByName[name]
			if !ok {
				slog.Warn("element requirement with missing cluster requirement", log.Path("source", deviceType.Source), slog.String("deviceType", deviceType.Name), slog.String("cluster", er.ClusterName))
				continue
			}
			crr.elementRequirements = append(crr.elementRequirements, er)
		}

	}
	for _, include := range clustersElement.SelectElements("include") {
		ca := include.SelectAttr("cluster")
		if ca == nil {
			slog.Warn("missing cluster attribute on include", slog.String("deviceTypeId", deviceType.ID.HexString()))
			clustersElement.RemoveChild(include)
			continue
		}
		crs, ok := clusterRequirementsByName[strings.ToLower(ca.Value)]
		if !ok {
			slog.Debug("unknown cluster attribute on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", ca.Value))
			clustersElement.RemoveChild(include)
			continue
		}
		setIncludeAttributes(clustersElement, include, spec, deviceType, crs)
		delete(clusterRequirementsByName, strings.ToLower(ca.Value))
	}
	for _, crs := range [][]*matter.ClusterRequirement{spec.BaseDeviceType.ClusterRequirements, deviceType.ClusterRequirements} {
		for _, cr := range crs {
			crr, ok := clusterRequirementsByName[strings.ToLower(cr.ClusterName)]
			if ok {
				setIncludeAttributes(clustersElement, nil, spec, deviceType, crr)
			}
		}
	}
	return
}

func setIncludeAttributes(clustersElement *etree.Element, include *etree.Element, spec *spec.Specification, deviceType *matter.DeviceType, cr *clusterRequirements) {
	cluster, ok := spec.ClustersByName[cr.name]
	if !ok {
		slog.Debug("unknown cluster on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cr.name))
		return
	}

	path, ok := spec.DocRefs[cluster]
	if !ok {
		slog.Warn("unknown doc path on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name))

	}
	errata, ok := zap.Erratas[filepath.Base(path)]
	if !ok {
		errata = zap.DefaultErrata
	}
	cxt := conformance.Context{
		Values: map[string]any{"Matter": true},
	}

	var server, client, clientLocked, serverLocked bool
	clientLocked = true
	serverLocked = true
	for i, crs := range [][]*matter.ClusterRequirement{cr.baseClusterRequirements, cr.clusterRequirements} {
		for _, cr := range crs {
			conf, err := cr.Conformance.Eval(cxt)
			if err != nil {
				slog.Warn("Error evaluating conformance of cluster requirement", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
				continue
			}

			switch conf {
			case conformance.StateOptional, conformance.StateProvisional:
				if i == 0 { // Base cluster requirements only get written when mandatory, apparently?
					return
				}
			case conformance.StateMandatory:
			default:
				return
			}

			switch cr.Interface {
			case matter.InterfaceServer:
				server = true
				serverLocked = conf == conformance.StateMandatory || conf == conformance.StateProvisional
			case matter.InterfaceClient:
				client = true
				clientLocked = conf == conformance.StateMandatory || conf == conformance.StateProvisional
			}
		}
	}

	if include == nil {
		include = clustersElement.CreateElement("include")
	}
	include.CreateAttr("cluster", cr.name)

	xml.SetNonexistentAttr(include, "client", strconv.FormatBool(client))
	xml.SetNonexistentAttr(include, "server", strconv.FormatBool(server))
	xml.SetNonexistentAttr(include, "clientLocked", strconv.FormatBool(clientLocked))
	xml.SetNonexistentAttr(include, "serverLocked", strconv.FormatBool(serverLocked))

	requiredAttributes := make(map[string]struct{})
	requiredAttributeDefines := make(map[string]struct{})
	requiredCommands := make(map[string]struct{})
	requiredCommandFields := make(map[string]map[string]struct{})
	requiredEvents := make(map[string]struct{})

	for _, er := range cr.elementRequirements {
		conf, err := er.Conformance.Eval(cxt)
		if err != nil {
			slog.Warn("Error evaluating conformance of element requirement", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			continue
		}
		if conf == conformance.StateMandatory || conf == conformance.StateProvisional {
			switch er.Element {
			case types.EntityTypeFeature:
				cxt.Values[er.Name] = true
			case types.EntityTypeAttribute:
				requiredAttributes[er.Name] = struct{}{}
				cxt.Values[er.Name] = true
			case types.EntityTypeCommand:
				requiredCommands[er.Name] = struct{}{}
				cxt.Values[er.Name] = true
			case types.EntityTypeEvent:
				requiredEvents[er.Name] = struct{}{}
				cxt.Values[er.Name] = true
			case types.EntityTypeCommandField:
				cf, ok := requiredCommandFields[er.Name]
				if !ok {
					cf = make(map[string]struct{})
					requiredCommandFields[er.Name] = cf
					cxt.Values[er.Name] = true
				}
				cf[er.Field] = struct{}{}
			default:
				slog.Warn("Element requirement with unrecognized element type", slog.String("deviceType", deviceType.Name), slog.String("entityType", er.Element.String()))
			}
		}
	}

	for _, attr := range cluster.Attributes {
		_, required := requiredAttributes[attr.Name]
		if !required {
			conf, err := attr.Conformance.Eval(cxt)
			if err != nil {
				slog.Warn("Error evaluating conformance of attribute", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
				continue

			}
			if conf == conformance.StateMandatory || conf == conformance.StateProvisional {
				required = true
			}

		}
		if required {
			requiredAttributeDefines[getDefine(attr.Name, errata.ClusterDefinePrefix, errata)] = struct{}{}
		}
	}
	for _, cmd := range cluster.Commands {
		conf, err := cmd.Conformance.Eval(cxt)
		if err != nil {
			slog.Warn("Error evaluating conformance of command", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			continue

		}
		if conf == conformance.StateMandatory || conf == conformance.StateProvisional {
			requiredCommands[cmd.Name] = struct{}{}
		}
	}
	for _, ev := range cluster.Events {
		conf, err := ev.Conformance.Eval(cxt)
		if err != nil {
			slog.Warn("Error evaluating conformance of event", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			continue

		}
		if conf == conformance.StateMandatory || conf == conformance.StateProvisional {
			requiredEvents[ev.Name] = struct{}{}
		}
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
		_, required := requiredCommands[rct]
		if required {
			delete(requiredCommands, rct)
		} else {
			include.RemoveChild(rc)
		}
	}
	for ra := range requiredCommands {
		rae := etree.NewElement("requireCommand")
		rae.SetText(ra)
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
		fields, required := requiredCommandFields[rct]
		if required {
			delete(requiredCommandFields, rct)
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
		rcfce.SetText(rcfc)
		for rcff := range rcffs {
			rcfe.CreateElement("field").SetText(rcff)
		}
		xml.InsertElementByName(include, rcfe, "requireAttribute", "requireCommand")
	}
	res := include.SelectElements("requireEvent")
	for _, re := range res {
		ret := re.Text()
		_, required := requiredEvents[ret]
		if required {
			delete(requiredEvents, ret)
		} else {
			include.RemoveChild(re)
		}
	}
	for ra := range requiredEvents {
		rae := etree.NewElement("requireEvent")
		rae.SetText(ra)
		xml.InsertElementByName(include, rae, "requireAttribute", "requireCommand", "requireCommandField")
	}
}

func matterDeviceTypeName(deviceType *matter.DeviceType) string {
	return fmt.Sprintf("Matter %s", deviceType.Name)
}
