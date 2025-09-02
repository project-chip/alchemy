package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type Condition struct {
	entity
	ID          *Number
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

type ConditionRequirement struct {
	entity
	DeviceTypeID   *Number `json:"deviceTypeId,omitempty"`
	DeviceTypeName string  `json:"deviceTypeName,omitempty"`
	ConditionName  string  `json:"conditionName,omitempty"`

	DeviceType *DeviceType                   `json:"deviceType,omitempty"`
	Location   DeviceTypeRequirementLocation `json:"location,omitempty"`

	Condition   *Condition      `json:"condition,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
}

func (c *ConditionRequirement) EntityType() types.EntityType {
	return types.EntityTypeConditionRequirement
}

func NewConditionRequirement(parent types.Entity, source asciidoc.Element) *ConditionRequirement {
	return &ConditionRequirement{entity: entity{parent: parent, source: source}}
}
