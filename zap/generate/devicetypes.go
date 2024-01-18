package generate

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func renderDeviceTypes(cxt context.Context, spec *matter.Spec, docs []*ascii.Doc, zclRoot string, filesOptions files.Options) (err error) {

	deviceTypes := newConcurrentMap[uint64, *matter.DeviceType]()

	files.ProcessDocs(cxt, docs, func(cxt context.Context, doc *ascii.Doc, index, total int) error {

		entities, err := doc.Entities()
		if err != nil {
			return err
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

	deviceTypesXmlPath := filepath.Join(zclRoot, "/src/app/zap-templates/zcl/data-model/chip/matter-devices.xml")

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
			slog.Warn("unknown deviceId", "deviceId", deviceTypeId.HexString())
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
		out, err = xml.WriteToString()
		if err != nil {
			return
		}
		//out, _ = parse.FormatXML(out)
		err = os.WriteFile(deviceTypesXmlPath, []byte(out), os.ModeAppend|0644)
	}
	return
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
	clusterRequirementsByName := make(map[string][]*matter.ClusterRequirement)
	for _, cr := range deviceType.ClusterRequirements {
		name := strings.ToLower(cr.ClusterName)
		clusterRequirementsByName[name] = append(clusterRequirementsByName[name], cr)
	}
	for _, cr := range spec.BaseDeviceType.ClusterRequirements {
		name := strings.ToLower(cr.ClusterName)
		clusterRequirementsByName[name] = append(clusterRequirementsByName[name], cr)
	}
	elementRequirementsByCluster := make(map[string][]*matter.ElementRequirement)
	for _, er := range deviceType.ElementRequirements {
		name := strings.ToLower(er.ClusterName)
		elementRequirementsByCluster[name] = append(elementRequirementsByCluster[name], er)
	}
	for _, er := range spec.BaseDeviceType.ElementRequirements {
		name := strings.ToLower(er.ClusterName)
		elementRequirementsByCluster[name] = append(elementRequirementsByCluster[name], er)
	}
	for _, include := range clustersElement.SelectElements("include") {
		ca := include.SelectAttr("cluster")
		if ca == nil {
			slog.Warn("missing cluster attribute on include", slog.String("deviceTypeId", deviceType.ID.HexString()))
			continue
		}
		crs, ok := clusterRequirementsByName[strings.ToLower(ca.Value)]
		if !ok {
			slog.Warn("unknown cluster attribute on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", ca.Value))
			continue
		}
		cluster, ok := spec.ClustersByName[ca.Value]
		if !ok {
			slog.Warn("unknown cluster on include", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", ca.Value))
			continue
		}
		elementRequirements := elementRequirementsByCluster[strings.ToLower(ca.Value)]

		setIncludeAttributes(include, spec, deviceType, cluster, crs, elementRequirements)
		delete(clusterRequirementsByName, strings.ToLower(cluster.Name))
	}
	return
}

func setIncludeAttributes(include *etree.Element, spec *matter.Spec, deviceType *matter.DeviceType, cluster *matter.Cluster, crs []*matter.ClusterRequirement, elementRequirements []*matter.ElementRequirement) {
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
	for _, cr := range crs {
		conf, err := cr.Conformance.Eval(cxt)
		if err != nil {
			slog.Warn("Error evaluating conformance of cluster requirement", slog.String("deviceTypeId", deviceType.ID.HexString()), slog.String("clusterName", cluster.Name), slog.Any("error", err))
			continue
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
	include.CreateAttr("client", strconv.FormatBool(client))
	include.CreateAttr("server", strconv.FormatBool(server))
	include.CreateAttr("clientLocked", strconv.FormatBool(clientLocked))
	include.CreateAttr("serverLocked", strconv.FormatBool(serverLocked))

	requiredAttributes := make(map[string]struct{})
	requiredAttributeDefines := make(map[string]struct{})
	requiredCommands := make(map[string]struct{})

	for _, er := range elementRequirements {
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

func processDeviceTypeConfigurator(xd *xml.Decoder, xe parse.XmlEncoder, deviceTypes *concurrentMap[uint64, *matter.DeviceType]) (err error) {
	for {
		var tok xml.Token
		tok, err = xd.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
				return
			}
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "deviceType":
				var deviceTypeTokens *parse.XmlTokenSet
				deviceTypeTokens, err = parse.XmlExtract(xd, tok)
				if err != nil {
					return
				}
				err = processDeviceType(xe, deviceTypeTokens, deviceTypes)
				if err != nil {
					slog.Warn("error processing deviceType", slog.Any("error", err))
					deviceTypeTokens.Reset()
					err = deviceTypeTokens.Write(xe)
					if err != nil {
						return
					}
					continue
				}
			default:
				err = xe.EncodeToken(tok)
			}
		case xml.EndElement:
			switch tok.Name.Local {
			case "configurator":
				err = xe.EncodeToken(tok)
				return
			}
		default:
			err = xe.EncodeToken(tok)
		}
		if err != nil {
			return
		}
	}
}

func processDeviceType(xe parse.XmlEncoder, ts *parse.XmlTokenSet, deviceTypes *concurrentMap[uint64, *matter.DeviceType]) (err error) {
	var deviceId string

	deviceId, err = getDeviceId(ts)
	if len(deviceId) == 0 {
		err = fmt.Errorf("unable to locate deviceId")
		return
	}
	ts.Reset()

	id := matter.ParseNumber(deviceId)
	if !id.Valid() {
		err = fmt.Errorf("invalid deviceId: %s", id.Text())
		return
	}
	deviceType, ok := deviceTypes.Map[id.Value()]
	if !ok {
		err = fmt.Errorf("unknown deviceId: %s", id.HexString())
		return
	}
	slog.Info("matched device type", slog.Any("deviceType", deviceType))

	var tok xml.Token
	for {
		tok, err = ts.Token()
		if err != nil {
			return
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "name":
				err = parse.XmlWriteSimpleElement(xe, tok.Name.Local, zap.ZAPDeviceTypeName(deviceType))
				if err != nil {
					return
				}
				err = ts.Ignore(tok.Name.Local)
			case "class":
				err = parse.XmlWriteSimpleElement(xe, tok.Name.Local, deviceType.Class)
				if err != nil {
					return
				}
				err = ts.Ignore(tok.Name.Local)
			case "scope":
				err = parse.XmlWriteSimpleElement(xe, tok.Name.Local, deviceType.Scope)
				if err != nil {
					return
				}
				err = ts.Ignore(tok.Name.Local)
			case "typeName":
				err = parse.XmlWriteSimpleElement(xe, tok.Name.Local, "Matter "+deviceType.Name)
				if err != nil {
					return
				}
				err = ts.Ignore(tok.Name.Local)
			case "clusters":
				var clustersTokens *parse.XmlTokenSet
				clustersTokens, err = parse.XmlExtract(ts, tok)
				if err != nil {
					return
				}
				err = processDeviceTypeClusters(xe, clustersTokens, deviceType)
				if err != nil {
					slog.Warn("error processing deviceType clusters", slog.Any("error", err))
					clustersTokens.Reset()
					err = clustersTokens.Write(xe)
					if err != nil {
						return
					}
					continue
				}
			default:
				err = xe.EncodeToken(tok)
			}
		case xml.EndElement:
			switch tok.Name.Local {
			case "deviceType":
				err = xe.EncodeToken(tok)
				return
			default:
				err = xe.EncodeToken(tok)
			}
		default:
			err = xe.EncodeToken(tok)
		}
		if err != nil {
			return
		}
	}
}

func getDeviceId(ts *parse.XmlTokenSet) (deviceId string, err error) {
	for {
		var tok xml.Token
		tok, err = ts.Token()
		if err != nil {
			return
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "deviceId":
				deviceId, err = ts.ReadElement(tok.Name.Local)
				return
			}
		}
	}
}

func processDeviceTypeClusters(xe parse.XmlEncoder, ts *parse.XmlTokenSet, deviceType *matter.DeviceType) (err error) {
	clusterRequirements := make(map[string]*matter.ClusterRequirement)
	for _, cr := range deviceType.ClusterRequirements {
		clusterRequirements[strings.ToLower(cr.ClusterName)] = cr
	}
	var tok xml.Token
	for {
		tok, err = ts.Token()
		if err != nil {
			return
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "include":
				ts.Ignore(tok.Name.Local)
				clusterName := parse.XmlAttributeGet(tok.Attr, "cluster")
				if len(clusterName) > 0 {
					_, ok := clusterRequirements[strings.ToLower(clusterName)]
					if !ok {
						break
					}

				}
			default:
				err = xe.EncodeToken(tok)
			}
		case xml.EndElement:
			switch tok.Name.Local {
			case "clusters":
				err = xe.EncodeToken(tok)
				return
			default:
				err = xe.EncodeToken(tok)
			}
		default:
			err = xe.EncodeToken(tok)
		}
		if err != nil {
			return
		}
	}
}

func processDeviceTypeClusterInclude(xe parse.XmlEncoder, ts *parse.XmlTokenSet, deviceType *matter.DeviceType) (err error) {
	var tok xml.Token
	for {
		tok, err = ts.Token()
		if err != nil {
			return
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "include":
				ts.Ignore(tok.Name.Local)
			default:
				err = xe.EncodeToken(tok)
			}
		case xml.EndElement:
			switch tok.Name.Local {
			case "clusters":
				err = xe.EncodeToken(tok)
				return
			default:
				err = xe.EncodeToken(tok)
			}
		default:
			err = xe.EncodeToken(tok)
		}
		if err != nil {
			return
		}
	}
}
