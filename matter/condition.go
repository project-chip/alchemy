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

func (c *Condition) Equals(e types.Entity) bool {
	oc, ok := e.(*Condition)
	if !ok {
		return false
	}
	if c.ID.Valid() && oc.ID.Valid() {
		return c.ID.Equals(oc.ID)
	}
	return c.Feature == oc.Feature
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

func (cr *ConditionRequirement) EntityType() types.EntityType {
	return types.EntityTypeConditionRequirement
}

func (cr *ConditionRequirement) Equals(e types.Entity) bool {
	ocr, ok := e.(*ConditionRequirement)
	if !ok {
		return false
	}
	if cr.DeviceTypeID.Valid() && ocr.DeviceTypeID.Valid() {
		if !cr.DeviceTypeID.Equals(ocr.DeviceTypeID) {
			return false
		}
	} else if cr.DeviceTypeName != ocr.DeviceTypeName {
		return false
	}
	return cr.ConditionName == ocr.ConditionName
}

func NewConditionRequirement(parent types.Entity, source asciidoc.Element) *ConditionRequirement {
	return &ConditionRequirement{entity: entity{parent: parent, source: source}}
}
