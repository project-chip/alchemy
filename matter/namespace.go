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
