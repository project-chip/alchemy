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
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

func renderDeviceTypes(cxt context.Context, spec *matter.Spec, docs []*ascii.Doc, sdkRoot string, filesOptions files.Options) (err error) {

	deviceTypes := newConcurrentMap[uint64, *matter.DeviceType]()

	err = files.ProcessDocs(cxt, docs, func(cxt context.Context, doc *ascii.Doc, index, total int) error {

		entities, err := doc.Entities()
		if err != nil {
			slog.WarnContext(cxt, "failed converting doc to device type entities", slog.String("path", doc.Path), slog.Any("error", err))
			return nil
		}
		for _, m := range entities {
			switch m := m.(type) {
			case *matter.DeviceType:
				deviceTypes.Lock()
				deviceTypes.Map[m.ID.Value()] = m
				deviceTypes.Unlock()
			}
		}
		return nil
	}, filesOptions)
	if err != nil {
		return
	}

	deviceTypesXmlPath := filepath.Join(sdkRoot, "/src/app/zap-templates/zcl/data-model/chip/matter-devices.xml")

	var deviceTypesXml []byte
	deviceTypesXml, err = os.ReadFile(deviceTypesXmlPath)
	if err != nil {
		return
	}

	xml := etree.NewDocument()
	err = xml.ReadFromBytes(deviceTypesXml)
	if err != nil {
		return
	}

	configurator := xml.SelectElement("configurator")
	if configurator == nil {
		err = fmt.Errorf("missing configurator element in %s", deviceTypesXmlPath)
		return
	}

	deviceTypeElements := configurator.SelectElements("deviceType")
	for _, deviceTypeElement := range deviceTypeElements {
		deviceIdElement := deviceTypeElement.SelectElement("deviceId")
		if deviceIdElement == nil {
			slog.Warn("missing deviceId element")
			continue
		}
		deviceTypeIdText := deviceIdElement.Text()
		deviceTypeId := matter.ParseNumber(deviceTypeIdText)
		if !deviceTypeId.Valid() {
			slog.Warn("invalid deviceId", "text", deviceTypeId.Text())
			continue
		}
		deviceType, ok := deviceTypes.Map[deviceTypeId.Value()]
		if !ok {
			continue
		}
		applyDeviceTypeToElement(spec, deviceType, deviceTypeElement)
		delete(deviceTypes.Map, deviceTypeId.Value())
	}

	for _, dt := range deviceTypes.Map {
		slog.Info("missing device type", slog.String("name", dt.Name))
		applyDeviceTypeToElement(spec, dt, configurator.CreateElement("deviceType"))
	}

	if !filesOptions.DryRun {
		var out string
		xml.Indent(4)
		xml.WriteSettings.CanonicalEndTags = true
		out, err = xml.WriteToString()
		if err != nil {
			return
		}
		//out, _ = parse.FormatXML(out)
		err = os.WriteFile(deviceTypesXmlPath, []byte(out), os.ModeAppend|0644)
	}
	return
}

type clusterRequirements struct {
	name                    string
	clusterRequirements     []*matter.ClusterRequirement
	baseClusterRequirements []*matter.ClusterRequirement
	elementRequirements     []*matter.ElementRequirement
}

func applyDeviceTypeToElement(spec *matter.Spec, deviceType *matter.DeviceType, dte *etree.Element) (err error) {
	setOrCreateSimpleElement(dte, "name", zap.ZAPDeviceTypeName(deviceType))
	setOrCreateSimpleElement(dte, "domain", "CHIP")
	setOrCreateSimpleElement(dte, "typeName", fmt.Sprintf("Matter %s", deviceType.Name))
	setOrCreateSimpleElement(dte, "profileId", "0x0103").CreateAttr("editable", "false")
	setOrCreateSimpleElement(dte, "deviceId", deviceType.ID.HexString()).CreateAttr("editable", "false")
	setOrCreateSimpleElement(dte, "class", deviceType.Class)
	setOrCreateSimpleElement(dte, "scope", deviceType.Scope)
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
		slog.Info("adding base device type cluster requirement", slog.String("cluster", cr.ClusterName))
	}
	for _, ers := range [][]*matter.ElementRequirement{deviceType.ElementRequirements, spec.BaseDeviceType.ElementRequirements} {
		for _, er := range ers {
			name := strings.ToLower(er.ClusterName)
			crr, ok := clusterRequirementsByName[name]
			if !ok {
				slog.Warn("elementr requirement with missing cluster requirement", slog.String("deviceType", deviceType.Name), slog.String("cluster", er.ClusterName))
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
			slog.Warn("unknown cluster attribute on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", ca.Value))
			clustersElement.RemoveChild(include)
			continue
		}
		setIncludeAttributes(include, spec, deviceType, crs)
		delete(clusterRequirementsByName, strings.ToLower(ca.Value))
	}
	for _, crs := range [][]*matter.ClusterRequirement{spec.BaseDeviceType.ClusterRequirements, deviceType.ClusterRequirements} {
		for _, cr := range crs {
			crr, ok := clusterRequirementsByName[strings.ToLower(cr.ClusterName)]
			if ok {
				setIncludeAttributes(clustersElement, spec, deviceType, crr)
			}
		}
	}
	return
}

func setIncludeAttributes(clustersElement *etree.Element, spec *matter.Spec, deviceType *matter.DeviceType, cr *clusterRequirements) {
	cluster, ok := spec.ClustersByName[cr.name]
	if !ok {
		slog.Warn("unknown cluster on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cr.name))
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

	include := clustersElement.CreateElement("include")
	include.CreateAttr("cluster", cr.name)

	setNonexistentAttr(include, "client", strconv.FormatBool(client))
	setNonexistentAttr(include, "server", strconv.FormatBool(server))
	setNonexistentAttr(include, "clientLocked", strconv.FormatBool(clientLocked))
	setNonexistentAttr(include, "serverLocked", strconv.FormatBool(serverLocked))

	requiredAttributes := make(map[string]struct{})
	requiredAttributeDefines := make(map[string]struct{})
	requiredCommands := make(map[string]struct{})

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
			default:
				slog.Warn("Element requirement with unrecognized element type", slog.Any("entityType", er.Element))
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
}
