package matter

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

type Domain uint8

const (
	DomainUnknown Domain = iota
	DomainGeneral
	DomainCHIP
	DomainMedia
	DomainLighting
	DomainAppliances
	DomainClosures
	DomainHVAC
	DomainMeasurementAndSensing
	DomainRobots
	DomainHomeAutomation
	DomainEnergyManagement
	DomainCameras
	DomainNetworkInfrastructure
)

var DomainNames = map[Domain]string{
	DomainUnknown:               "Unknown",
	DomainGeneral:               "General",
	DomainCHIP:                  "CHIP",
	DomainMedia:                 "Media",
	DomainLighting:              "Lighting",
	DomainAppliances:            "Appliances",
	DomainClosures:              "Closures",
	DomainHVAC:                  "HVAC",
	DomainMeasurementAndSensing: "Measurement & Sensing",
	DomainRobots:                "Robots",
	DomainHomeAutomation:        "Home Automation",
	DomainEnergyManagement:      "Energy Management",
	DomainNetworkInfrastructure: "Network Infrastructure",
	DomainCameras:               "Cameras",
}

func DomainFromDocType(docType DocType) Domain {
	switch docType {
	case DocTypeServiceDeviceManagement:
		return DomainGeneral
	default:
		return DomainUnknown
	}
}

func (i Domain) MarshalYAML() ([]byte, error) {
	name, ok := DomainNames[i]
	if ok {
		return yaml.Marshal(name)
	}
	return nil, fmt.Errorf("missing name for domain: %d", i)
}

func (i *Domain) UnmarshalYAML(b []byte) error {

	var v string
	if err := yaml.Unmarshal(b, &v); err != nil {
		return err
	}
	for d, n := range DomainNames {
		if n == v {
			*i = d
			return nil
		}
	}
	return fmt.Errorf("unknown domain name: %s", v)

}
