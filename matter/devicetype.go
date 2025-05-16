package matter

import (
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

	Superset string `json:"superset,omitempty"`
	Class    string `json:"class,omitempty"`
	Scope    string `json:"scope,omitempty"`

	Conditions             []*Condition             `json:"conditions,omitempty"`
	DeviceTypeRequirements []*DeviceTypeRequirement `json:"deviceTypeRequirements,omitempty"`

	ClusterRequirements            []*ClusterRequirement            `json:"clusterRequirements,omitempty"`
	ElementRequirements            []*ElementRequirement            `json:"elementRequirements,omitempty"`
	ComposedDeviceTypeRequirements []*ComposedDeviceTypeRequirement `json:"composedDeviceTypeRequirements,omitempty"`
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

func NewClusterRequirement(source asciidoc.Element) *ClusterRequirement {
	return &ClusterRequirement{entity: entity{source: source}}
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

func NewElementRequirement(source asciidoc.Element) ElementRequirement {
	return ElementRequirement{entity: entity{source: source}}
}

type ElementRequirement struct {
	entity
	ClusterID   *Number          `json:"clusterId,omitempty"`
	ClusterName string           `json:"clusterName,omitempty"`
	Element     types.EntityType `json:"element,omitempty"`
	Name        string           `json:"name,omitempty"`
	Field       string           `json:"field,omitempty"`

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

func NewDeviceTypeRequirement(source asciidoc.Element) *DeviceTypeRequirement {
	return &DeviceTypeRequirement{entity: entity{source: source}}
}

type DeviceTypeRequirement struct {
	entity
	DeviceTypeID   *Number               `json:"deviceTypeId,omitempty"`
	DeviceTypeName string                `json:"deviceTypeName,omitempty"`
	Constraint     constraint.Constraint `json:"constraint,omitempty"`
	Conformance    conformance.Set       `json:"conformance,omitempty"`
	AllowsSuperset bool                  `json:"allowsSuperset,omitempty"`
}

func (dtr *DeviceTypeRequirement) EntityType() types.EntityType {
	return types.EntityTypeDeviceTypeRequirement
}

type ComposedDeviceTypeRequirement struct {
	DeviceTypeID   *Number `json:"deviceTypeId,omitempty"`
	DeviceTypeName string  `json:"deviceTypeName,omitempty"`
	ElementRequirement
}

type Condition struct {
	entity
	Feature     string
	Description string
}

func NewCondition(source asciidoc.Element, dt *DeviceType) *Condition {
	return &Condition{entity: entity{source: source, parent: dt}}
}

func (c *Condition) EntityType() types.EntityType {
	return types.EntityTypeCondition
}

func (c *Condition) GetConformance() conformance.Set {
	return nil
}
