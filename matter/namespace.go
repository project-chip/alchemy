package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type Namespace struct {
	entity
	ID           *Number        `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	SemanticTags []*SemanticTag `json:"semanticTags,omitempty"`
}

func NewNamespace(source asciidoc.Element) *Namespace {
	return &Namespace{
		entity: entity{source: source},
	}
}

func (*Namespace) EntityType() types.EntityType {
	return types.EntityTypeNamespace
}

func (ns *Namespace) Equals(e types.Entity) bool {
	ons, ok := e.(*Namespace)
	if !ok {
		return false
	}
	if ns.ID.Valid() && ons.ID.Valid() {
		return ns.ID.Equals(ons.ID)
	}
	return ns.Name == ons.Name
}

func (ns *Namespace) Clone() *Namespace {
	nc := &Namespace{entity: entity{source: ns.source, parent: ns.parent}, ID: ns.ID.Clone(), Name: ns.Name}
	nc.SemanticTags = make([]*SemanticTag, 0, len(ns.SemanticTags))
	for _, t := range ns.SemanticTags {
		tc := t.Clone()
		tc.parent = nc
		nc.SemanticTags = append(nc.SemanticTags, tc)
	}
	return nc
}

type SemanticTag struct {
	entity

	ID          *Number `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
}

func (*SemanticTag) EntityType() types.EntityType {
	return types.EntityTypeSemanticTag
}

func (st *SemanticTag) Equals(e types.Entity) bool {
	ost, ok := e.(*SemanticTag)
	if !ok {
		return false
	}
	namespace, ok := st.parent.(*Namespace)
	if !ok {
		return false
	}
	otherNamespace, ok := ost.parent.(*Namespace)
	if !ok {
		return false
	}
	if !namespace.Equals(otherNamespace) {
		return false
	}
	if st.ID.Valid() && ost.ID.Valid() {
		return st.ID.Equals(ost.ID)
	}
	return st.Name == ost.Name
}

func NewSemanticTag(namespace *Namespace, source asciidoc.Element) *SemanticTag {
	return &SemanticTag{
		entity: entity{parent: namespace, source: source},
	}
}

func (ns *SemanticTag) Clone() *SemanticTag {
	nc := &SemanticTag{entity: entity{source: ns.source, parent: ns.parent}, ID: ns.ID.Clone(), Name: ns.Name, Description: ns.Description}
	return nc
}
