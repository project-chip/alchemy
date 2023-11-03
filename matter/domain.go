package matter

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
}
