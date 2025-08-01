package matter

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type DeviceType struct {
	entity
	ID          *Number     `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Revisions   []*Revision `json:"revisions,omitempty"`

	SupersetOf string `json:"supersetOf,omitempty"`
	Class      string `json:"class,omitempty"`
	Scope      string `json:"scope,omitempty"`

	SubsetDeviceType *DeviceType `json:"-"`

	Conditions []*Condition `json:"conditions,omitempty"`

	ClusterRequirements   []*ClusterRequirement   `json:"clusterRequirements,omitempty"`
	ElementRequirements   []*ElementRequirement   `json:"elementRequirements,omitempty"`
	ConditionRequirements []*ConditionRequirement `json:"conditionRequirements,omitempty"`

	DeviceTypeRequirements                []*DeviceTypeRequirement        `json:"deviceTypeRequirements,omitempty"`
	ComposedDeviceTypeClusterRequirements []*DeviceTypeClusterRequirement `json:"composedDeviceTypeClusterRequirements,omitempty"`
	ComposedDeviceTypeElementRequirements []*DeviceTypeElementRequirement `json:"composedDeviceTypeElementRequirements,omitempty"`
}

func NewDeviceType(source asciidoc.Element) *DeviceType {
	return &DeviceType{entity: entity{source: source}}
}

func (dt *DeviceType) EntityType() types.EntityType {
	return types.EntityTypeDeviceType
}

func (dt *DeviceType) Identifier(name string) (types.Entity, bool) {
	for _, c := range dt.Conditions {
		if c.Feature == name {
			return c, true
		}
	}
	return nil, false
}

func NewClusterRequirement(parent *DeviceType, source asciidoc.Element) *ClusterRequirement {
	return &ClusterRequirement{entity: entity{parent: parent, source: source}}
}

type ClusterRequirement struct {
	entity
	ClusterID   *Number         `json:"clusterId,omitempty"`
	ClusterName string          `json:"clusterName,omitempty"`
	Quality     Quality         `json:"quality,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Interface   Interface       `json:"interface,omitempty"`

	Cluster *Cluster `json:"cluster,omitempty"`
}

func (cr *ClusterRequirement) EntityType() types.EntityType {
	return types.EntityTypeClusterRequirement
}

func (cr *ClusterRequirement) Clone() *ClusterRequirement {
	cer := &ClusterRequirement{
		entity:      entity{source: cr.source},
		ClusterID:   cr.ClusterID.Clone(),
		ClusterName: cr.ClusterName,
		Interface:   cr.Interface,
		Quality:     cr.Quality,
		Cluster:     cr.Cluster,
	}
	if len(cr.Conformance) > 0 {
		cer.Conformance = cr.Conformance.CloneSet()
	}
	return cer
}

func NewElementRequirement(parent types.Entity, source asciidoc.Element) ElementRequirement {
	return ElementRequirement{entity: entity{parent: parent, source: source}}
}

type ElementRequirement struct {
	entity
	ClusterID   *Number          `json:"clusterId,omitempty"`
	ClusterName string           `json:"clusterName,omitempty"`
	Element     types.EntityType `json:"element,omitempty"`
	Name        string           `json:"name,omitempty"`
	Field       string           `json:"field,omitempty"`

	Entity types.Entity `json:"entity,omitempty"`

	Constraint  constraint.Constraint `json:"constraint,omitempty"`
	Quality     Quality               `json:"quality,omitempty"`
	Access      Access                `json:"access,omitempty"`
	Conformance conformance.Set       `json:"conformance,omitempty"`

	Cluster *Cluster `json:"cluster,omitempty"`
}

func (er *ElementRequirement) EntityType() types.EntityType {
	return types.EntityTypeElementRequirement
}

func (er *ElementRequirement) Clone() *ElementRequirement {
	cer := &ElementRequirement{
		entity:      entity{source: er.source},
		ClusterID:   er.ClusterID.Clone(),
		ClusterName: er.ClusterName,
		Element:     er.Element,
		Name:        er.Name,
		Field:       er.Field,
		Quality:     er.Quality,
		Access:      er.Access,
		Cluster:     er.Cluster,
	}
	if er.Constraint != nil {
		cer.Constraint = er.Constraint.Clone()
	}
	if len(er.Conformance) > 0 {
		cer.Conformance = er.Conformance.CloneSet()
	}
	return cer
}

func NewDeviceTypeRequirement(parent *DeviceType, source asciidoc.Element) *DeviceTypeRequirement {
	return &DeviceTypeRequirement{entity: entity{parent: parent, source: source}}
}

type DeviceTypeRequirementLocation uint8

const (
	DeviceTypeRequirementLocationUnknown DeviceTypeRequirementLocation = iota
	DeviceTypeRequirementLocationDeviceEndpoint
	DeviceTypeRequirementLocationChildEndpoint
	DeviceTypeRequirementLocationRootEndpoint
	DeviceTypeRequirementLocationDescendantEndpoint
)

var (
	deviceTypeRequirementRelations = map[DeviceTypeRequirementLocation]string{
		DeviceTypeRequirementLocationUnknown:            "unknown",
		DeviceTypeRequirementLocationDeviceEndpoint:     "deviceEndpoint",
		DeviceTypeRequirementLocationChildEndpoint:      "childEndpoint",
		DeviceTypeRequirementLocationRootEndpoint:       "rootEndpoint",
		DeviceTypeRequirementLocationDescendantEndpoint: "descendantEndpoint",
	}
)

func (s DeviceTypeRequirementLocation) String() string {
	str, ok := deviceTypeRequirementRelations[s]
	if ok {
		return str
	}
	return fmt.Sprintf("DeviceTypeRequirementRelation(%d)", s)
}

type DeviceTypeRequirement struct {
	entity
	DeviceTypeID   *Number                       `json:"deviceTypeId,omitempty"`
	DeviceTypeName string                        `json:"deviceTypeName,omitempty"`
	Constraint     constraint.Constraint         `json:"constraint,omitempty"`
	Conformance    conformance.Set               `json:"conformance,omitempty"`
	AllowsSuperset bool                          `json:"allowsSuperset,omitempty"`
	Location       DeviceTypeRequirementLocation `json:"location,omitempty"`

	DeviceType *DeviceType `json:"deviceType,omitempty"`
}

func (dtr *DeviceTypeRequirement) EntityType() types.EntityType {
	return types.EntityTypeDeviceTypeRequirement
}

func (dtr *DeviceTypeRequirement) Clone() *DeviceTypeRequirement {
	cdtr := &DeviceTypeRequirement{
		entity:         entity{source: dtr.source},
		DeviceTypeID:   dtr.DeviceTypeID.Clone(),
		DeviceTypeName: dtr.DeviceTypeName,
		AllowsSuperset: dtr.AllowsSuperset,
		DeviceType:     dtr.DeviceType,
	}
	if dtr.Constraint != nil {
		cdtr.Constraint = dtr.Constraint.Clone()
	}
	if len(dtr.Conformance) > 0 {
		cdtr.Conformance = dtr.Conformance.CloneSet()
	}
	return cdtr
}
