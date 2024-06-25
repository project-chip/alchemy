package zap

import "github.com/project-chip/alchemy/matter"

func StringToDomain(name string) matter.Domain {
	switch name {
	case "Home Appliances":
		name = "Appliances"
	case "Measurement and Sensing":
		name = "Measurement & Sensing"
	}
	var domain matter.Domain
	for d, dn := range matter.DomainNames {
		if dn == name {
			domain = d
			break
		}
	}
	return domain
}
