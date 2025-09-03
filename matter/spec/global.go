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

func (cf *globalCommandFactory) Children(library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section) iter.Seq[*asciidoc.Section] {
	return func(yield func(*asciidoc.Section) bool) {
		parse.Search(d, reader, s, reader.Children(s), func(doc *asciidoc.Document, sec *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
			if library.SectionType(sec) != matter.SectionCommand {
				return parse.SearchShouldContinue
			}
			if !yield(sec) {
				return parse.SearchShouldStop
			}
			return parse.SearchShouldContinue
		})
	}
}

func (library *Library) toGlobalElements(spec *Specification, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (entities []types.Entity, err error) {
	var commandsTable *asciidoc.Table
	parse.SkimFunc(reader, s, reader.Children(s), func(t *asciidoc.Table) bool {
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
	commands, err = buildList(spec, library, reader, d, s, commandsTable, commands, &cf, parent)

	commandMap := make(map[string]*matter.Command)
	for _, c := range commands {
		entities = append(entities, c)
		commandMap[c.Name] = c
	}

	// The definition of global commands is frequently elsewhere, so let's scan the doc for other commmand sections
	parse.Search(d, reader, d, reader.Children(d), func(doc *asciidoc.Document, sec *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
		switch library.SectionType(sec) {
		case matter.SectionCommand:
			commandName := text.TrimCaseInsensitiveSuffix(library.SectionName(sec), " Command")
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
			err = readCommand(spec, library, reader, d, sec, command)
			if err != nil {
				return parse.SearchShouldStop
			}
		}
		return parse.SearchShouldContinue
	})

	return
}
