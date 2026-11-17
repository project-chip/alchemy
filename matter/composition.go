package matter

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type DeviceTypeComposition struct {
	DeviceType             *DeviceType
	DeviceTypeRequirements []*DeviceTypeRequirement

	ClusterRequirements []*DeviceTypeClusterRequirement
	ElementRequirements []*DeviceTypeElementRequirement
	TagRequirements     []*DeviceTypeTagRequirement

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
	DeviceTypeRef EntityReference `json:"deviceTypeRef,omitempty"`

	ClusterRef  EntityReference `json:"clusterRef,omitempty"`
	Quality     Quality         `json:"quality,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Interface   Interface       `json:"interface,omitempty"`

	RequirementOrigin RequirementOrigin

	DeviceType            *DeviceType
	DeviceTypeRequirement *DeviceTypeRequirement

	Cluster            *Cluster
	ClusterRequirement *ClusterRequirement
}

func NewDeviceTypeClusterRequirement(parent types.Entity, source asciidoc.Element) *DeviceTypeClusterRequirement {
	return &DeviceTypeClusterRequirement{parent: parent, source: source}
}

func (dtcr *DeviceTypeClusterRequirement) EntityType() types.EntityType {
	return types.EntityTypeDeviceTypeClusterRequirement
}

func (dtcr *DeviceTypeClusterRequirement) Equals(other types.Entity) bool {
	odtcr, ok := other.(*DeviceTypeClusterRequirement)
	if !ok {
		return false
	}
	return dtcr.DeviceTypeRef.Equals(odtcr.DeviceTypeRef) && dtcr.ClusterRef.Equals(odtcr.ClusterRef)
}

func (dtcr *DeviceTypeClusterRequirement) CloneTo(parent types.Entity) *DeviceTypeClusterRequirement {
	return &DeviceTypeClusterRequirement{
		parent:                parent,
		source:                dtcr.source,
		DeviceTypeRef:         dtcr.DeviceTypeRef,
		ClusterRef:            dtcr.ClusterRef,
		Cluster:               dtcr.Cluster,
		ClusterRequirement:    dtcr.ClusterRequirement,
		Interface:             dtcr.Interface,
		Quality:               dtcr.Quality,
		Conformance:           dtcr.Conformance.CloneSet(),
		RequirementOrigin:     dtcr.RequirementOrigin,
		DeviceType:            dtcr.DeviceType,
		DeviceTypeRequirement: dtcr.DeviceTypeRequirement,
	}
}

type DeviceTypeElementRequirement struct {
	entity
	DeviceTypeRef EntityReference `json:"deviceTypeRef,omitempty"`

	ElementRequirement *ElementRequirement
	RequirementOrigin  RequirementOrigin

	DeviceType            *DeviceType
	DeviceTypeRequirement *DeviceTypeRequirement
}

func NewDeviceTypeElementRequirement(parent types.Entity, elementRequirement *ElementRequirement, source asciidoc.Element) *DeviceTypeElementRequirement {
	return &DeviceTypeElementRequirement{entity: entity{parent: parent, source: source}, ElementRequirement: elementRequirement}
}

func (dter *DeviceTypeElementRequirement) EntityType() types.EntityType {
	return types.EntityTypeDeviceTypeElementRequirement
}

func (dter *DeviceTypeElementRequirement) Equals(other types.Entity) bool {
	odter, ok := other.(*DeviceTypeElementRequirement)
	if !ok {
		return false
	}
	return dter.DeviceTypeRef.Equals(odter.DeviceTypeRef) && dter.ElementRequirement.Equals(odter.ElementRequirement)
}

func (dter *DeviceTypeElementRequirement) CloneTo(parent types.Entity) *DeviceTypeElementRequirement {
	return &DeviceTypeElementRequirement{
		parent:                parent,
		source:                dter.source,
		DeviceTypeRef:         dter.DeviceTypeRef,
		ElementRequirement:    dter.ElementRequirement,
		RequirementOrigin:     dter.RequirementOrigin,
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
		TagRequirements:        make([]*DeviceTypeTagRequirement, len(dc.TagRequirements)),
		ComposedDeviceTypes:    make(map[DeviceTypeRequirementLocation][]*DeviceTypeComposition, len(dc.ComposedDeviceTypes)),
	}
	copy(clone.DeviceTypeRequirements, dc.DeviceTypeRequirements)
	copy(clone.ClusterRequirements, dc.ClusterRequirements)
	copy(clone.ElementRequirements, dc.ElementRequirements)
	copy(clone.TagRequirements, dc.TagRequirements)
	for location, composedDeviceTypes := range dc.ComposedDeviceTypes {
		ccdt := make([]*DeviceTypeComposition, len(composedDeviceTypes))
		copy(ccdt, composedDeviceTypes)
		clone.ComposedDeviceTypes[location] = ccdt
	}
	return clone
}

type DeviceTypeTagRequirement struct {
	entity
	TagRequirement *TagRequirement

	DeviceTypeRef EntityReference `json:"deviceTypeRef,omitempty"`

	DeviceType            *DeviceType `json:"deviceType,omitempty"`
	DeviceTypeRequirement *DeviceTypeRequirement
}

func (dttr *DeviceTypeTagRequirement) EntityType() types.EntityType {
	return types.EntityTypeDeviceTypeTagRequirement
}

func (dttr *DeviceTypeTagRequirement) Equals(other types.Entity) bool {
	dttr2, ok := other.(*DeviceTypeTagRequirement)
	if !ok {
		return false
	}
	return dttr.TagRequirement.Equals(dttr2.TagRequirement) && dttr.DeviceTypeRef.Equals(dttr2.DeviceTypeRef)
}

func (dttr *DeviceTypeTagRequirement) Clone() *DeviceTypeTagRequirement {
	return &DeviceTypeTagRequirement{
		entity:                entity{parent: dttr.parent, source: dttr.source},
		DeviceTypeRef:         dttr.DeviceTypeRef,
		DeviceType:            dttr.DeviceType,
		DeviceTypeRequirement: dttr.DeviceTypeRequirement,
		TagRequirement:        dttr.TagRequirement,
	}
}

func NewDeviceTypeTagRequirement(parent *DeviceType, source asciidoc.Element) *DeviceTypeTagRequirement {
	return &DeviceTypeTagRequirement{entity: entity{parent: parent, source: source}}
}
