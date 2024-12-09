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

type ClusterRequirement struct {
	ClusterID   *Number         `json:"clusterId,omitempty"`
	ClusterName string          `json:"clusterName,omitempty"`
	Quality     Quality         `json:"quality,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Interface   Interface       `json:"interface,omitempty"`

	Cluster *Cluster `json:"cluster,omitempty"`
}

func (cr *ClusterRequirement) Clone() *ClusterRequirement {
	cer := &ClusterRequirement{
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

type ElementRequirement struct {
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

func (er *ElementRequirement) Clone() *ElementRequirement {
	cer := &ElementRequirement{
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

type DeviceTypeRequirement struct {
	DeviceTypeID   *Number               `json:"deviceTypeId,omitempty"`
	DeviceTypeName string                `json:"deviceTypeName,omitempty"`
	Constraint     constraint.Constraint `json:"constraint,omitempty"`
	Conformance    conformance.Set       `json:"conformance,omitempty"`
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

func NewCondition(source asciidoc.Element) *Condition {
	return &Condition{entity: entity{source: source}}
}

func (c *Condition) EntityType() types.EntityType {
	return types.EntityTypeCondition
}

func (c *Condition) GetConformance() conformance.Set {
	return nil
}
