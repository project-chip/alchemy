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

	Conditions []*Condition `json:"conditions,omitempty"`

	ClusterRequirements []*ClusterRequirement `json:"clusterRequirements,omitempty"`
	ElementRequirements []*ElementRequirement `json:"elementRequirements,omitempty"`
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
	ID          *Number         `json:"id,omitempty"`
	ClusterName string          `json:"clusterName,omitempty"`
	Quality     Quality         `json:"quality,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Interface   Interface       `json:"interface,omitempty"`

	Cluster *Cluster `json:"cluster,omitempty"`
}

type ElementRequirement struct {
	ID          *Number          `json:"id,omitempty"`
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
