package spec

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type globalCommandFactory struct {
	commandFactory
}

func (cf *globalCommandFactory) Children(d *Doc, s *asciidoc.Section) iter.Seq[*asciidoc.Section] {
	return func(yield func(*asciidoc.Section) bool) {
		parse.Search(d.Iterator(), d, d.Children(), func(sec *asciidoc.Section, parent asciidoc.Parent, index int) parse.SearchShould {
			if d.SectionType(sec) != matter.SectionCommand {
				return parse.SearchShouldContinue
			}
			if !yield(s) {
				return parse.SearchShouldStop
			}
			return parse.SearchShouldContinue
		})
	}
}

func toGlobalElements(spec *Specification, d *Doc, s *asciidoc.Section, pc *parseContext, parent types.Entity) (entities []types.Entity, err error) {
	var commandsTable *asciidoc.Table
	parse.SkimFunc(d.Iterator(), s, s.Children(), func(t *asciidoc.Table) bool {
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

	commandMap := make(map[string]*matter.Command)
	for _, c := range commands {
		entities = append(entities, c)
		commandMap[c.Name] = c
	}

	// The definnition of global commands is frequently elsewhere, so let's scan the doc for other commmand sections
	parse.Search(d.Iterator(), d, d.Children(), func(sec *asciidoc.Section, parent asciidoc.Parent, index int) parse.SearchShould {
		switch d.SectionType(sec) {
		case matter.SectionCommand:
			commandName := text.TrimCaseInsensitiveSuffix(d.SectionName(sec), " Command")
			command, ok := commandMap[commandName]
			if !ok {
				return parse.SearchShouldContinue
			}
			if len(command.Fields) > 0 { // We've already found fields for this command, so skip
				return parse.SearchShouldContinue
			}
			if command.Source() == sec {
				return parse.SearchShouldContinue
			}
			err = readCommand(pc, spec, d, sec, command)
			if err != nil {
				return parse.SearchShouldStop
			}
		}
		return parse.SearchShouldContinue
	})

	return
}
