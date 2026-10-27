package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type CrossReference struct {
	Document  *asciidoc.Document
	Reference *asciidoc.CrossReference
	Parent    asciidoc.Parent
	Source    matter.Source
}

func (library *Library) SyncCrossReference(cr *CrossReference, id asciidoc.Elements) {
	if !id.Equals(cr.Reference.ID) {
		library.changeCrossReference(library, cr, id)
		cr.Reference.ID = id
	}
}

func ReferenceName(reader asciidoc.Reader, element any) string {
	if element == nil {
		return ""
	}
	switch el := element.(type) {
	case *asciidoc.Anchor:
		return buildReferenceName(reader, el, el.Elements)
	case *asciidoc.Section:
		return buildReferenceName(reader, el, el.Title)
	case asciidoc.Attributable:
		return referenceNameFromAttributes(reader, el)
	case asciidoc.Element:
		return ReferenceName(reader, el)
	default:
		slog.Warn("Unknown type to get reference name", "type", fmt.Sprintf("%T", element))
	}
	return ""
}

func buildReferenceName(reader asciidoc.Reader, parent asciidoc.Parent, set asciidoc.Elements) string {
	var val strings.Builder

	for el := range reader.Iterate(parent, set) {
		switch el := el.(type) {
		case *asciidoc.String:
			val.WriteString(el.Value)
		case *asciidoc.SpecialCharacter:
			var char string
			switch el.Character {
			case "&", ">", "<":
				char = el.Character
			default:
				slog.Warn("unrecognized special character", "char", el.Character, "context", val.String())
			}
			val.WriteString(char)
		case asciidoc.Attributable:
			val.WriteString(referenceNameFromAttributes(reader, el))
		case *asciidoc.CharacterReplacementReference:
			val.WriteString(el.Value)
		case *asciidoc.Counter:
		default:
			slog.Warn("unknown reference name element", "element", el, "type", fmt.Sprintf("%T", el))
		}
	}
	return val.String()
}

func referenceNameFromAttributes(reader asciidoc.Reader, el asciidoc.Attributable) string {
	for _, a := range el.Attributes() {
		switch a := a.(type) {
		case *asciidoc.AnchorAttribute:
		case *asciidoc.TableColumnsAttribute:
		case *asciidoc.PositionalAttribute:
		case *asciidoc.ShorthandAttribute:
		case *asciidoc.TitleAttribute:
			return a.AsciiDocString()
		case *asciidoc.NamedAttribute:
			if a.Name == asciidoc.AttributeNameTitle {
				switch v := a.Value().(type) {
				case string:
					return v
				case asciidoc.Elements:
					return buildReferenceName(reader, &a.Val, v)
				default:
					slog.Warn("unexpected value of section title attribute", slog.String("type", fmt.Sprintf("%T", a.Value())))
				}
			}
		default:
			slog.Warn("Unknown attribute type to get reference name", "type", fmt.Sprintf("%T", a), "val", a.Value())
		}
	}
	return ""
}

func referenceInfo(library *Library, reference *asciidoc.CrossReference) (referenceID string, label string) {
	referenceID = library.elementIdentifier(library, reference, reference, reference.ID)
	if len(reference.Elements) > 0 {
		label = buildReferenceName(library, reference, reference.Elements)
	}
	return
}

func findEntityByReference(spec *Specification, reference *asciidoc.CrossReference, entityFinder entityFinder) (entity types.Entity, ok bool) {
	var library *Library
	library, ok = spec.libraryIndex[reference.Document()]
	if !ok {
		slog.Warn("Missing library for cross-reference", log.Path("source", reference))
		return
	}

	referenceID, label := referenceInfo(library, reference)
	entity = entityFinder.findEntityByReference(referenceID, label, reference)
	if entity == nil {
		slog.Warn("No entity found for reference", slog.String("referenceID", referenceID), slog.String("label", label), log.Path("source", reference))
	}
	return
}

func resolveEntityReference[T types.Entity](spec *Specification, ref *matter.EntityReference, entityFinder entityFinder, findByID findEntityByID[T], findByName findEntityByName[T], source types.Entity) (entity T, ok bool) {
	if ref.XRef != nil {
		ref.Resolved, ok = findEntityByReference(spec, ref.XRef, entityFinder)
		if ok {
			switch e := ref.Resolved.(type) {
			case T:
				entity = e
				ok = true
				return
			case nil:
			default:
				slog.Warn("Entity type mismatch for reference", slog.String("expectedEntityType", entity.EntityType().String()), slog.String("actualEntityType", e.EntityType().String()), log.Path("source", source))
			}
		}
	}
	if ref.Resolved == nil && findByID != nil && ref.ID.Valid() {
		entity, ok = findByID(ref.ID)
		if ok {
			ref.Resolved = entity
		}
	}
	if ref.Resolved == nil && findByName != nil && ref.Name != "" {
		entity, ok = findByName(ref.Name)
		if ok {
			slog.Warn("linking entity reference by name since ID was not recognized",
				slog.String("ID", ref.ID.HexString()),
				slog.String("name", ref.Name),
				log.Path("source", entity))
			ref.Resolved = entity
		}
	}
	return
}

func resolveElementReference(spec *Specification, er *matter.ElementRequirement, entityFinder entityFinder, cluster *matter.Cluster, source types.Entity) (err error) {
	if er.ElementRef.XRef != nil {
		e, ok := findEntityByReference(spec, er.ElementRef.XRef, entityFinder)
		if ok {
			if e.EntityType() != er.Element {
				slog.Warn("Entity type mismatch for reference", slog.String("expectedEntityType", er.Element.String()), slog.String("actualEntityType", e.EntityType().String()), log.Path("source", source))
			}
			er.Entity = e
			return
		}
	}

	switch er.Element {
	case types.EntityTypeAttribute:
		for _, a := range cluster.Attributes {
			if strings.EqualFold(a.Name, er.ElementRef.Name) {
				er.Entity = a
				return
			}
		}
	case types.EntityTypeFeature:
		if cluster.Features == nil {
			slog.Error("Element Requirement references missing features", slog.String("clusterName", cluster.Name), slog.String("clusterId", cluster.ID.HexString()), slog.String("feature", er.ElementRef.Name), log.Path("source", er))
			return
		}
		for _, fb := range cluster.Features.Bits {
			f, ok := fb.(*matter.Feature)
			if !ok {
				continue
			}
			if f.Code == er.ElementRef.Name || strings.EqualFold(f.Name(), er.ElementRef.Name) {
				er.Entity = f
				return
			}
		}
	case types.EntityTypeCommand:
		for _, cmd := range cluster.Commands {
			if strings.EqualFold(cmd.Name, er.ElementRef.Name) {
				er.Entity = cmd
				return
			}
		}
	case types.EntityTypeCommandField:
		var command *matter.Command
		for _, cmd := range cluster.Commands {
			if strings.EqualFold(cmd.Name, er.ElementRef.Name) {
				command = cmd
				break
			}
		}
		if command == nil {
			break
		}
		for _, f := range command.Fields {
			if strings.EqualFold(f.Name, er.ElementRef.Field) {
				er.Entity = f
				return
			}
		}
	case types.EntityTypeEvent:
		for _, e := range cluster.Events {
			if strings.EqualFold(e.Name, er.ElementRef.Name) {
				er.Entity = e
				return
			}
		}

	default:
		slog.Error("Unexpected entity type", slog.String("entityType", er.Element.String()), log.Path("source", er))

	}
	return
}
