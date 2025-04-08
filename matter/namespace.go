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

type SemanticTag struct {
	entity

	ID          *Number `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
}

func (*SemanticTag) EntityType() types.EntityType {
	return types.EntityTypeSemanticTag
}

func NewSemanticTag(namespace *Namespace, source asciidoc.Element) *SemanticTag {
	return &SemanticTag{
		entity: entity{parent: namespace, source: source},
	}
}
