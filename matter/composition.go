package matter

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
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
	entity
	DeviceTypeID   *Number `json:"deviceTypeId,omitempty"`
	DeviceTypeName string  `json:"deviceTypeName,omitempty"`

	ClusterRequirement *ClusterRequirement
	Origin             RequirementOrigin

	DeviceType            *DeviceType
	DeviceTypeRequirement *DeviceTypeRequirement
}

func NewDeviceTypeClusterRequirement(parent types.Entity, clusterRequirement *ClusterRequirement, source asciidoc.Element) *DeviceTypeClusterRequirement {
	return &DeviceTypeClusterRequirement{entity: entity{parent: parent, source: source}, ClusterRequirement: clusterRequirement}
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
	entity
	DeviceTypeID   *Number `json:"deviceTypeId,omitempty"`
	DeviceTypeName string  `json:"deviceTypeName,omitempty"`

	ElementRequirement *ElementRequirement
	Origin             RequirementOrigin

	DeviceType            *DeviceType
	DeviceTypeRequirement *DeviceTypeRequirement
}

func NewDeviceTypeElementRequirement(parent types.Entity, elementRequirement *ElementRequirement, source asciidoc.Element) *DeviceTypeElementRequirement {
	return &DeviceTypeElementRequirement{entity: entity{parent: parent, source: source}, ElementRequirement: elementRequirement}
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
	Server  conformance.ConformanceState
	Client  conformance.ConformanceState

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

type DeviceTypeTagRequirement struct {
	entity
	DeviceTypeID   *Number                       `json:"deviceTypeId,omitempty"`
	DeviceTypeName string                        `json:"deviceTypeName,omitempty"`
	Constraint     constraint.Constraint         `json:"constraint,omitempty"`
	Conformance    conformance.Set               `json:"conformance,omitempty"`
	Location       DeviceTypeRequirementLocation `json:"location,omitempty"`

	NamespaceID   *Number `json:"namespaceId,omitempty"`
	NamespaceName string  `json:"namespaceName,omitempty"`

	SemanticTagID   *Number `json:"semanticTagId,omitempty"`
	SemanticTagName string  `json:"semanticTagName,omitempty"`

	DeviceType            *DeviceType `json:"deviceType,omitempty"`
	DeviceTypeRequirement *DeviceTypeRequirement
	Namespace             *Namespace   `json:"namespace,omitempty"`
	SemanticTag           *SemanticTag `json:"semanticTag,omitempty"`
}

func (dtcr *DeviceTypeTagRequirement) Clone() *DeviceTypeTagRequirement {
	return &DeviceTypeTagRequirement{
		DeviceTypeID:    dtcr.DeviceTypeID,
		DeviceTypeName:  dtcr.DeviceTypeName,
		DeviceType:      dtcr.DeviceType,
		Location:        dtcr.Location,
		NamespaceID:     dtcr.NamespaceID,
		NamespaceName:   dtcr.NamespaceName,
		SemanticTagID:   dtcr.SemanticTagID,
		SemanticTagName: dtcr.SemanticTagName,
		Namespace:       dtcr.Namespace,
		SemanticTag:     dtcr.SemanticTag,
	}
}

func NewDeviceTypeTagRequirement(parent *DeviceType, source asciidoc.Element) *DeviceTypeTagRequirement {
	return &DeviceTypeTagRequirement{entity: entity{parent: parent, source: source}}
}
