package generate

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

var utilityDevicesMask uint64 = 0xFF000000

type DeviceTypesPatcher struct {
	sdkRoot        string
	spec           *spec.Specification
	clusterAliases map[string]string

	generateFeatureXML bool
}

type DeviceTypePatcherOption func(dtp *DeviceTypesPatcher)

func DeviceTypePatcherGenerateFeatureXML(generate bool) DeviceTypePatcherOption {
	return func(dtp *DeviceTypesPatcher) {
		dtp.generateFeatureXML = generate
	}
}

func NewDeviceTypesPatcher(sdkRoot string, spec *spec.Specification, clusterAliases pipeline.Map[string, []string], options ...DeviceTypePatcherOption) *DeviceTypesPatcher {
	dtp := &DeviceTypesPatcher{sdkRoot: sdkRoot, spec: spec, clusterAliases: make(map[string]string)}
	clusterAliases.Range(func(cluster string, aliases []string) bool {
		for _, alias := range aliases {
			dtp.clusterAliases[alias] = cluster
		}
		return true
	})
	for _, opt := range options {
		opt(dtp)
	}
	return dtp
}

func (p DeviceTypesPatcher) Name() string {
	return "Patching device types"
}

func (p DeviceTypesPatcher) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (p DeviceTypesPatcher) Process(cxt context.Context, inputs []*pipeline.Data[[]*matter.DeviceType]) (outputs []*pipeline.Data[[]byte], err error) {

	deviceTypesToUpdateByID := make(map[uint64]*matter.DeviceType)
	deviceTypesToUpdateByName := make(map[string]*matter.DeviceType)
	for _, input := range inputs {
		for _, dt := range input.Content {
			if dt.ID.Valid() {
				deviceTypesToUpdateByID[dt.ID.Value()] = dt
			} else {
				deviceTypesToUpdateByName[dt.Name] = dt
			}
		}
	}

	allDeviceTypesByID := make(map[uint64]*matter.DeviceType)
	allDeviceTypesByName := make(map[string]*matter.DeviceType)

	for _, deviceType := range p.spec.DeviceTypes {
		if deviceType.ID.Valid() {
			allDeviceTypesByID[deviceType.ID.Value()] = deviceType
		} else {
			allDeviceTypesByName[deviceType.Name] = deviceType
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
		var deviceTypeToUpdate *matter.DeviceType
		deviceIDElement := deviceTypeElement.SelectElement("deviceId")
		if deviceIDElement != nil {
			deviceTypeIDText := deviceIDElement.Text()
			deviceTypeID := matter.ParseNumber(deviceTypeIDText)
			if deviceTypeID.Valid() {
				if (deviceTypeID.Value() & utilityDevicesMask) == utilityDevicesMask {
					// Exception for the all clusters app, etc
					continue
				}
				deviceTypeToUpdate = deviceTypesToUpdateByID[deviceTypeID.Value()]
				if deviceTypeToUpdate != nil {
					delete(deviceTypesToUpdateByID, deviceTypeID.Value())
				} else {
					deviceType = allDeviceTypesByID[deviceTypeID.Value()]
				}
			} else if deviceTypeIDText != "ID-TBD" {
				slog.Warn("invalid deviceId", "text", deviceTypeID.Text())
			}

		}
		if deviceTypeToUpdate == nil {
			deviceTypeElement := deviceTypeElement.SelectElement("typeName")
			if deviceTypeElement == nil {
				slog.Warn("missing deviceId and typeName elements")
				continue
			}
			deviceTypeNameText := deviceTypeElement.Text()
			deviceTypeToUpdate = deviceTypesToUpdateByName[deviceTypeNameText]
			if deviceTypeToUpdate != nil {
				delete(deviceTypesToUpdateByName, deviceTypeNameText)
			} else if deviceType == nil {
				deviceType = allDeviceTypesByName[deviceTypeNameText]
			}
		}
		if deviceTypeToUpdate != nil {
			if !matter.NonGlobalIDInvalidForEntity(deviceTypeToUpdate.ID, types.EntityTypeDeviceType) {
				p.applyDeviceTypeToElement(p.spec, deviceTypeToUpdate, deviceTypeElement)
			} else {
				configurator.RemoveChild(deviceTypeElement)
			}
		} else if deviceType == nil {
			configurator.RemoveChild(deviceTypeElement)
		}
	}

	for _, dt := range deviceTypesToUpdateByID {
		slog.Info("Adding new device type", slog.String("name", dt.Name))
		if matter.NonGlobalIDInvalidForEntity(dt.ID, types.EntityTypeDeviceType) {
			continue
		}
		p.applyDeviceTypeToElement(p.spec, dt, configurator.CreateElement("deviceType"))
	}

	for _, dt := range deviceTypesToUpdateByName {
		slog.Info("Adding new device type", slog.String("name", dt.Name))
		p.applyDeviceTypeToElement(p.spec, dt, configurator.CreateElement("deviceType"))
	}

	var out string
	xml.Indent(4)
	xml.WriteSettings.CanonicalEndTags = true
	out, err = xml.WriteToString()
	if err != nil {
		return
	}
	out = postProcessTemplate(out)
	outputs = append(outputs, pipeline.NewData[[]byte](deviceTypesXMLPath, []byte(out)))
	return
}
