package spec

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type globalCommandFactory struct {
	commandFactory
}

func (cf *globalCommandFactory) Children(d *Doc, s *Section) iter.Seq[*Section] {
	return func(yield func(*Section) bool) {
		parse.Traverse(d, d.Elements(), func(sec *Section, parent parse.HasElements, index int) parse.SearchShould {
			if s.SecType != matter.SectionCommand {
				return parse.SearchShouldContinue
			}
			if !yield(s) {
				return parse.SearchShouldStop
			}
			return parse.SearchShouldContinue
		})
	}
}

func (s *Section) toGlobalElements(spec *Specification, d *Doc, pc *parseContext, parent types.Entity) (entities []types.Entity, err error) {
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

	var cf globalCommandFactory
	var commands matter.CommandSet
	commands, err = buildList(spec, d, s, commandsTable, pc, commands, &cf, parent)

	for _, c := range commands {
		entities = append(entities, c)
	}

	return
}
