package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func addGlobalEntities(spec *Specification, doc *Doc) error {
	globalEntities, err := doc.GlobalObjects()
	if err != nil {
		return err
	}
	for _, m := range globalEntities {
		spec.DocRefs[m] = doc
		switch m := m.(type) {
		case *matter.Bitmap:
			slog.Debug("Found global bitmap", "name", m.Name, "path", doc.Path)
			_, ok := spec.bitmapIndex[m.Name]
			if ok {
				slog.Warn("multiple global bitmaps with same name", "name", m.Name)
			} else {
				spec.bitmapIndex[m.Name] = m
			}
			spec.addEntity(m.Name, m, nil)
		case *matter.Enum:
			slog.Debug("Found global enum", "name", m.Name, "path", doc.Path)
			_, ok := spec.enumIndex[m.Name]
			if ok {
				slog.Warn("multiple global enums with same name", "name", m.Name)
			} else {
				spec.enumIndex[m.Name] = m
			}
			spec.addEntity(m.Name, m, nil)
		case *matter.Struct:
			slog.Debug("Found global struct", "name", m.Name, "path", doc.Path)
			_, ok := spec.structIndex[m.Name]
			if ok {
				slog.Warn("multiple global structs with same name", "name", m.Name)
			} else {
				spec.structIndex[m.Name] = m
			}
			spec.addEntity(m.Name, m, nil)
		case *matter.Command:
			_, ok := spec.commandIndex[m.Name]
			if ok {
				slog.Warn("multiple global commands with same name", "name", m.Name)
			} else {
				spec.commandIndex[m.Name] = m
			}
			spec.addEntity(m.Name, m, nil)
		case *matter.Event:
			_, ok := spec.eventIndex[m.Name]
			if ok {
				slog.Warn("multiple global events with same name", "name", m.Name)
			} else {
				spec.eventIndex[m.Name] = m
			}
			spec.addEntity(m.Name, m, nil)
		}
		spec.GlobalObjects[m] = struct{}{}

	}
	return nil
}

func (s *Section) toGlobalElements(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (entities []types.Entity, err error) {
	var commandsTable *asciidoc.Table
	parse.SkimFunc(s.Elements(), func(t *asciidoc.Table) bool {
		for _, a := range t.AttributeList.Attributes() {
			switch a := a.(type) {
			case *asciidoc.TitleAttribute:
				if a.AsciiDocString() == "Global Commands" {
					commandsTable = t
					return true
				}
			}
		}
		return false
	})
	if commandsTable == nil {
		return
	}
	var commands matter.CommandSet
	var commandMap map[string]*matter.Command
	commands, _, err = s.buildCommands(d, commandsTable)
	parse.Traverse(d, d.Elements(), func(sec *Section, parent parse.HasElements, index int) parse.SearchShould {
		switch s.SecType {
		case matter.SectionCommand:
			var c *matter.Command
			c, err := s.toCommand(d, commandMap, entityMap)
			if err != nil {
				return parse.SearchShouldContinue
			}
			if c != nil {
				entityMap[s.Base] = append(entityMap[s.Base], c)
			}
		}
		return parse.SearchShouldContinue
	})
	for _, c := range commands {
		entities = append(entities, c)
	}

	return
}
