package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter"
)

func ParseDomain(name string) matter.Domain {
	switch name {
	case "Home Appliances":
		name = "Appliances"
	case "Measurement and Sensing":
		name = "Measurement & Sensing"
	case "WebRTC Transport":
		name = "Cameras"
	case "Entry Control":
		name = "Closures"
	}
	var domain matter.Domain
	for d, dn := range matter.DomainNames {
		if dn == name {
			domain = d
			break
		}
	}
	if domain == matter.DomainUnknown {
		slog.Info("Unrecognized domain", "name", name)
		domain = matter.DomainGeneral
	}
	return domain
}
