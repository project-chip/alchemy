package matter

import (
	"fmt"

	"github.com/project-chip/alchemy/matter/conformance"
)

type DeviceTypeComposition struct {
	DeviceType             *DeviceType
	DeviceTypeRequirements []*DeviceTypeRequirement

	ClusterRequirements []*DeviceTypeClusterRequirement
	ElementRequirements []*DeviceTypeElementRequirement

	ComposedDeviceTypes map[DeviceTypeRequirementLocation][]*DeviceTypeComposition
}

type RequirementOrigin uint8

const ( // The order here is important, as we use it as priority for overriding later
	RequirementOriginUnknown RequirementOrigin = iota
	RequirementOriginBaseDeviceType
	RequirementOriginSubsetDeviceType
	RequirementOriginComposedDeviceType
	RequirementOriginDeviceType
)

var (
	requirementOrigins = map[RequirementOrigin]string{
		RequirementOriginUnknown:            "unknown",
		RequirementOriginDeviceType:         "deviceType",
		RequirementOriginBaseDeviceType:     "baseDeviceType",
		RequirementOriginSubsetDeviceType:   "subsetDeviceType",
		RequirementOriginComposedDeviceType: "composedDeviceType",
	}
)

func (s RequirementOrigin) Compare(oro RequirementOrigin) int {
	if s > oro {
		return 1
	} else if s < oro {
		return -1
	}
	return 0
}

func (s RequirementOrigin) String() string {
	str, ok := requirementOrigins[s]
	if ok {
		return str
	}
	return fmt.Sprintf("RequirementOrigin(%d)", s)
}

type DeviceTypeClusterRequirement struct {
	DeviceTypeID   *Number `json:"deviceTypeId,omitempty"`
	DeviceTypeName string  `json:"deviceTypeName,omitempty"`

	ClusterRequirement *ClusterRequirement
	Origin             RequirementOrigin

	DeviceType            *DeviceType
	DeviceTypeRequirement *DeviceTypeRequirement
}

func (dtcr *DeviceTypeClusterRequirement) Clone() *DeviceTypeClusterRequirement {
	return &DeviceTypeClusterRequirement{
		DeviceTypeID:          dtcr.DeviceTypeID,
		DeviceTypeName:        dtcr.DeviceTypeName,
		ClusterRequirement:    dtcr.ClusterRequirement,
		Origin:                dtcr.Origin,
		DeviceType:            dtcr.DeviceType,
		DeviceTypeRequirement: dtcr.DeviceTypeRequirement,
	}
}

type DeviceTypeElementRequirement struct {
	DeviceTypeID   *Number `json:"deviceTypeId,omitempty"`
	DeviceTypeName string  `json:"deviceTypeName,omitempty"`

	ElementRequirement *ElementRequirement
	Origin             RequirementOrigin

	DeviceType            *DeviceType
	DeviceTypeRequirement *DeviceTypeRequirement
}

func (dter *DeviceTypeElementRequirement) Clone() *DeviceTypeElementRequirement {
	return &DeviceTypeElementRequirement{
		DeviceTypeID:          dter.DeviceTypeID,
		DeviceTypeName:        dter.DeviceTypeName,
		ElementRequirement:    dter.ElementRequirement,
		Origin:                dter.Origin,
		DeviceType:            dter.DeviceType,
		DeviceTypeRequirement: dter.DeviceTypeRequirement,
	}
}

type ClusterComposition struct {
	Cluster *Cluster
	Server  conformance.State
	Client  conformance.State

	Elements []*ElementComposition
}

type ElementComposition struct {
	ElementRequirement *ElementRequirement
	State              conformance.ConformanceState
}

func (dc *DeviceTypeComposition) Clone() *DeviceTypeComposition {
	clone := &DeviceTypeComposition{
		DeviceType:             dc.DeviceType,
		DeviceTypeRequirements: make([]*DeviceTypeRequirement, len(dc.DeviceTypeRequirements)),
		ClusterRequirements:    make([]*DeviceTypeClusterRequirement, len(dc.ClusterRequirements)),
		ElementRequirements:    make([]*DeviceTypeElementRequirement, len(dc.ElementRequirements)),
		ComposedDeviceTypes:    make(map[DeviceTypeRequirementLocation][]*DeviceTypeComposition, len(dc.ComposedDeviceTypes)),
	}
	copy(clone.DeviceTypeRequirements, dc.DeviceTypeRequirements)
	copy(clone.ClusterRequirements, dc.ClusterRequirements)
	copy(clone.ElementRequirements, dc.ElementRequirements)
	for location, composedDeviceTypes := range dc.ComposedDeviceTypes {
		ccdt := make([]*DeviceTypeComposition, len(composedDeviceTypes))
		copy(ccdt, composedDeviceTypes)
		clone.ComposedDeviceTypes[location] = ccdt
	}
	return clone
}
