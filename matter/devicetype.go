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
	ID          *Number   `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Revisions   Revisions `json:"revisions,omitempty"`

	SupersetOf string `json:"supersetOf,omitempty"`
	Class      string `json:"class,omitempty"`
	Scope      string `json:"scope,omitempty"`

	ProfileID string `json:"profileId,omitempty"`

	SubsetDeviceType *DeviceType `json:"-"`

	Conditions []*Condition `json:"conditions,omitempty"`

	ClusterRequirements   []*ClusterRequirement   `json:"clusterRequirements,omitempty"`
	ElementRequirements   []*ElementRequirement   `json:"elementRequirements,omitempty"`
	ConditionRequirements []*ConditionRequirement `json:"conditionRequirements,omitempty"`
	TagRequirements       []*TagRequirement       `json:"tagRequirements,omitempty"`

	DeviceTypeRequirements                []*DeviceTypeRequirement        `json:"deviceTypeRequirements,omitempty"`
	ComposedDeviceTypeClusterRequirements []*DeviceTypeClusterRequirement `json:"composedDeviceTypeClusterRequirements,omitempty"`
	ComposedDeviceTypeElementRequirements []*DeviceTypeElementRequirement `json:"composedDeviceTypeElementRequirements,omitempty"`

	ComposedDeviceTagRequirements []*DeviceTypeTagRequirement `json:"composedDeviceSemanticTagRequirements,omitempty"`
}

func NewDeviceType(source asciidoc.Element) *DeviceType {
	return &DeviceType{entity: entity{source: source}, ProfileID: "0x0103"}
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

func (dt *DeviceType) Equals(e types.Entity) bool {
	odt, ok := e.(*DeviceType)
	if !ok {
		return false
	}
	if dt.ID.Valid() && odt.ID.Valid() {
		return dt.ID.Equals(odt.ID)
	}
	return dt.Name == odt.Name
}

func NewClusterRequirement(parent *DeviceType, source asciidoc.Element) *ClusterRequirement {
	return &ClusterRequirement{entity: entity{parent: parent, source: source}}
}

type ClusterRequirement struct {
	entity
	ClusterRef EntityReference `json:"clusterRef,omitempty"`

	Quality     Quality         `json:"quality,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Interface   Interface       `json:"interface,omitempty"`

	Cluster *Cluster `json:"cluster,omitempty"`
}

func (cr *ClusterRequirement) EntityType() types.EntityType {
	return types.EntityTypeClusterRequirement
}

func (cr *ClusterRequirement) Equals(e types.Entity) bool {
	ocr, ok := e.(*ClusterRequirement)
	if !ok {
		return false
	}
	return cr.ClusterRef.Equals(ocr.ClusterRef)
}

func (cr *ClusterRequirement) Clone() *ClusterRequirement {
	cer := &ClusterRequirement{
		entity:     entity{source: cr.source},
		ClusterRef: cr.ClusterRef.Clone(),
		Interface:  cr.Interface,
		Quality:    cr.Quality,
		Cluster:    cr.Cluster,
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
	ClusterRef EntityReference  `json:"clusterRef,omitempty"`
	Element    types.EntityType `json:"element,omitempty"`
	ElementRef ElementReference `json:"elementRef,omitempty"`

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

func (er *ElementRequirement) Equals(e types.Entity) bool {
	oer, ok := e.(*ElementRequirement)
	if !ok {
		return false
	}
	if !er.ClusterRef.Equals(oer.ClusterRef) {
		return false
	}
	if !er.ElementRef.Equals(oer.ElementRef) {
		return false
	}
	if er.Entity != nil {
		if oer.Entity != nil {
			return er.Entity.Equals(oer.Entity)
		} else {
			return false
		}
	} else if oer.Entity != nil {
		return false
	}
	return true
}

func (er *ElementRequirement) Clone() *ElementRequirement {
	cer := &ElementRequirement{
		entity:     entity{source: er.source},
		ClusterRef: er.ClusterRef.Clone(),
		ElementRef: er.ElementRef,
		Quality:    er.Quality,
		Access:     er.Access,
		Cluster:    er.Cluster,
	}
	if er.Constraint != nil {
		cer.Constraint = er.Constraint.Clone()
	}
	if len(er.Conformance) > 0 {
		cer.Conformance = er.Conformance.CloneSet()
	}
	return cer
}

type TagRequirement struct {
	entity
	Constraint  constraint.Constraint `json:"constraint,omitempty"`
	Conformance conformance.Set       `json:"conformance,omitempty"`

	NamespaceRef   EntityReference `json:"namespaceRef,omitempty"`
	SemanticTagRef EntityReference `json:"semanticTagRef,omitempty"`

	Namespace   *Namespace   `json:"namespace,omitempty"`
	SemanticTag *SemanticTag `json:"semanticTag,omitempty"`
}

func (dtcr *TagRequirement) Clone() *TagRequirement {
	return &TagRequirement{
		NamespaceRef:   dtcr.NamespaceRef.Clone(),
		SemanticTagRef: dtcr.SemanticTagRef.Clone(),
		Namespace:      dtcr.Namespace,
		SemanticTag:    dtcr.SemanticTag,
	}
}

func NewTagRequirement(parent *DeviceType, source asciidoc.Element) *TagRequirement {
	return &TagRequirement{entity: entity{parent: parent, source: source}}
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
	DeviceTypeRef  EntityReference               `json:"deviceTypeRef,omitempty"`
	Constraint     constraint.Constraint         `json:"constraint,omitempty"`
	Conformance    conformance.Set               `json:"conformance,omitempty"`
	AllowsSuperset bool                          `json:"allowsSuperset,omitempty"`
	Location       DeviceTypeRequirementLocation `json:"location,omitempty"`

	DeviceType *DeviceType `json:"deviceType,omitempty"`
}

func (dtr *DeviceTypeRequirement) EntityType() types.EntityType {
	return types.EntityTypeDeviceTypeRequirement
}

func (dtr *DeviceTypeRequirement) Equals(e types.Entity) bool {
	odtr, ok := e.(*DeviceTypeRequirement)
	if !ok {
		return false
	}
	return dtr.DeviceTypeRef.Equals(odtr.DeviceTypeRef)
}

func (dtr *DeviceTypeRequirement) Clone() *DeviceTypeRequirement {
	cdtr := &DeviceTypeRequirement{
		entity:         entity{source: dtr.source},
		DeviceTypeRef:  dtr.DeviceTypeRef.Clone(),
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

func NewDeviceTypeRequirement(parent *DeviceType, source asciidoc.Element) *DeviceTypeRequirement {
	return &DeviceTypeRequirement{entity: entity{parent: parent, source: source}}
}
