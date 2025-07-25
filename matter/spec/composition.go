package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
)

func (spec *Specification) ComposeDeviceType(deviceType *matter.DeviceType) (composition *matter.DeviceTypeComposition, err error) {

	if existing, ok := spec.deviceTypeCompositionCache[deviceType]; ok {
		return existing, nil
	}

	composition = &matter.DeviceTypeComposition{DeviceType: deviceType}
	compositionRequirements := make(map[*matter.DeviceTypeComposition]*matter.DeviceTypeRequirement)
	composedDevices := make(map[*matter.DeviceType]*matter.DeviceTypeComposition)

	if deviceType.SubsetDeviceType != nil {

		var subsetDeviceTypeComposition *matter.DeviceTypeComposition
		subsetDeviceTypeComposition, err = spec.ComposeDeviceType(deviceType.SubsetDeviceType)
		if err != nil {
			return
		}

		for _, cr := range subsetDeviceTypeComposition.ClusterRequirements {

			if cr.ClusterRequirement.Cluster == nil {
				slog.Warn("cluster requirement missing cluster", slog.String("deviceType", deviceType.Name))
				continue
			}
			var subsetOrigin = cr.Origin
			if subsetOrigin != matter.RequirementOriginBaseDeviceType {
				subsetOrigin = matter.RequirementOriginSubsetDeviceType
			}

			composition.ClusterRequirements = append(composition.ClusterRequirements, &matter.DeviceTypeClusterRequirement{ClusterRequirement: cr.ClusterRequirement, Origin: subsetOrigin})
		}

		for _, er := range subsetDeviceTypeComposition.ElementRequirements {
			var subsetOrigin = er.Origin
			if subsetOrigin != matter.RequirementOriginBaseDeviceType {
				subsetOrigin = matter.RequirementOriginSubsetDeviceType
			}
			composition.ElementRequirements = append(composition.ElementRequirements, &matter.DeviceTypeElementRequirement{ElementRequirement: er.ElementRequirement, Origin: subsetOrigin})
		}
	}

	for _, dtr := range deviceType.DeviceTypeRequirements {
		if dtr.DeviceType == nil {
			continue
		}
		var cdt *matter.DeviceTypeComposition
		cdt, err = spec.ComposeDeviceType(dtr.DeviceType)
		if err != nil {
			return
		}
		if composition.ComposedDeviceTypes == nil {
			composition.ComposedDeviceTypes = make(map[matter.DeviceTypeRequirementLocation][]*matter.DeviceTypeComposition)
		}
		cdt = cdt.Clone()
		composition.ComposedDeviceTypes[dtr.Location] = append(composition.ComposedDeviceTypes[dtr.Location], cdt)
		switch dtr.Location {
		case matter.DeviceTypeRequirementLocationDeviceEndpoint:
			compositionRequirements[cdt] = dtr
		}
		composedDevices[dtr.DeviceType] = cdt
		composition.DeviceTypeRequirements = append(composition.DeviceTypeRequirements, dtr)
	}

	for location, drs := range composition.ComposedDeviceTypes {
		switch location {
		case matter.DeviceTypeRequirementLocationDeviceEndpoint:
			for _, dr := range drs {
				req := compositionRequirements[dr]
				for _, cr := range dr.ClusterRequirements {
					origin := matter.RequirementOriginComposedDeviceType
					if cr.Origin == matter.RequirementOriginBaseDeviceType {
						origin = matter.RequirementOriginBaseDeviceType
					}
					composition.ClusterRequirements = append(composition.ClusterRequirements, &matter.DeviceTypeClusterRequirement{ClusterRequirement: cr.ClusterRequirement, Origin: origin, DeviceTypeRequirement: req})
				}
				for _, er := range dr.ElementRequirements {
					origin := matter.RequirementOriginComposedDeviceType
					if er.Origin == matter.RequirementOriginBaseDeviceType {
						origin = matter.RequirementOriginBaseDeviceType
					}
					composition.ElementRequirements = append(composition.ElementRequirements, &matter.DeviceTypeElementRequirement{ElementRequirement: er.ElementRequirement, Origin: origin, DeviceTypeRequirement: req})
				}
			}
		}
	}

	for _, cr := range deviceType.ComposedDeviceTypeClusterRequirements {
		if cr.ClusterRequirement.Cluster == nil {
			slog.Warn("Cluster requirement on composed device type missing cluster", log.Path("source", cr.ClusterRequirement))
			continue
		}
		if cr.DeviceType == nil {
			slog.Warn("Cluster requirement on composed device type missing device type", log.Path("source", cr.ClusterRequirement))
			continue
		}
		var comp *matter.DeviceTypeComposition
		switch cr.DeviceTypeRequirement.Location {
		case matter.DeviceTypeRequirementLocationDeviceEndpoint:
			comp = composition
		default:
			var ok bool
			comp, ok = composedDevices[cr.DeviceType]
			if !ok {
				slog.Warn("Cluster requirement on composed device type not found", log.Path("source", cr.ClusterRequirement))
				continue
			}
		}
		cr = cr.Clone()
		cr.Origin = matter.RequirementOriginComposedDeviceType
		comp.ClusterRequirements = append(comp.ClusterRequirements, cr)
	}

	for _, er := range deviceType.ComposedDeviceTypeElementRequirements {
		if er.ElementRequirement.Cluster == nil {
			slog.Warn("Element requirement on composed device type missing cluster", log.Path("source", er.ElementRequirement))
			continue
		}
		if er.DeviceType == nil {
			slog.Warn("Element requirement on composed device type missing device type", log.Path("source", er.ElementRequirement))
			continue
		}
		var comp *matter.DeviceTypeComposition
		switch er.DeviceTypeRequirement.Location {
		case matter.DeviceTypeRequirementLocationDeviceEndpoint:
			comp = composition
		default:
			var ok bool
			comp, ok = composedDevices[er.DeviceType]
			if !ok {
				slog.Warn("Element requirement on composed device type not found", log.Path("source", er.ElementRequirement))
				continue
			}
		}
		er = er.Clone()
		er.Origin = matter.RequirementOriginComposedDeviceType
		comp.ElementRequirements = append(comp.ElementRequirements, er)
	}

	var origin = matter.RequirementOriginDeviceType
	if deviceType == spec.BaseDeviceType {
		origin = matter.RequirementOriginBaseDeviceType
	}

	for _, cr := range deviceType.ClusterRequirements {
		if cr.Cluster == nil {
			continue
		}
		composition.ClusterRequirements = append(composition.ClusterRequirements, &matter.DeviceTypeClusterRequirement{ClusterRequirement: cr, Origin: origin})
	}

	for _, er := range deviceType.ElementRequirements {
		if er.Cluster == nil {
			continue
		}
		if er.Entity == nil {
			continue
		}
		composition.ElementRequirements = append(composition.ElementRequirements, &matter.DeviceTypeElementRequirement{ElementRequirement: er, Origin: origin})
	}

	spec.deviceTypeCompositionCache[deviceType] = composition
	return
}
