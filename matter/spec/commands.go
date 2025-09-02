package spec

import (
	"iter"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

var parentheticalExpressionPattern = regexp.MustCompile(`\s*\([^\)]+\)$`)

type commandFactory struct {
}

func (cf *commandFactory) New(spec *Specification, library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, ti *TableInfo, row *asciidoc.TableRow, name string, parent types.Entity) (*matter.Command, error) {
	cmd := matter.NewCommand(s, parent)
	var err error
	cmd.ID, err = ti.ReadID(reader, row, matter.TableColumnCommandID, matter.TableColumnID)
	if err != nil {
		return nil, err
	}

	cmd.Name = text.TrimCaseInsensitiveSuffix(name, " Command")
	var dir string
	dir, err = ti.ReadString(reader, row, matter.TableColumnDirection)
	if err != nil {
		return nil, err
	}
	cmd.Direction = ParseCommandDirection(dir)
	cmd.Response, _ = ti.ReadDataType(library, reader, row, matter.TableColumnResponse)
	if cmd.Response != nil {
		cmd.Response.Name = text.TrimCaseInsensitiveSuffix(cmd.Response.Name, " Command")
	}
	cmd.Conformance = ti.ReadConformance(library, row, matter.TableColumnConformance)
	var a string
	a, err = ti.ReadString(reader, row, matter.TableColumnAccess)
	if err != nil {
		return nil, err
	}
	cmd.Quality, err = ti.ReadQuality(reader, row, types.EntityTypeCommand, matter.TableColumnQuality)
	if err != nil {
		return nil, err
	}
	cmd.Access, _ = ParseAccess(a, types.EntityTypeCommand)
	return cmd, nil
}

func (cf *commandFactory) Details(spec *Specification, library *Library, reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, command *matter.Command) (err error) {
	return readCommand(spec, library, reader, doc, section, command)
}

func readCommand(spec *Specification, library *Library, reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, command *matter.Command) (err error) {
	command.Description = library.getDescription(reader, doc, command, section, reader.Children(section))

	command.Name = CanonicalName(command.Name)

	de := errata.GetSpec(doc.Path.Relative)

	if de.IgnoreSection(library.SectionName(section), errata.SpecPurposeCommandArguments) {
		return
	}
	var ti *TableInfo
	ti, err = parseFirstTable(reader, doc, section)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			slog.Warn("No valid command parameter table found", log.Element("source", doc.Path, section), "command", command.Name)
			err = nil
		}
		return
	}
	var fieldMap map[string]*matter.Field
	command.Fields, fieldMap, err = library.readFields(spec, reader, ti, types.EntityTypeCommandField, command)
	if err != nil {
		return
	}
	err = library.mapFields(reader, doc, section, fieldMap)
	if err != nil {
		return
	}
	return
}

func (cf *commandFactory) EntityName(library *Library, reader asciidoc.Reader, doc *asciidoc.Document, s *asciidoc.Section) string {
	name := strings.ToLower(text.TrimCaseInsensitiveSuffix(library.SectionName(s), " Command"))
	return parentheticalExpressionPattern.ReplaceAllString(name, "")
}

func (cf *commandFactory) Children(library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section) iter.Seq[*asciidoc.Section] {
	return func(yield func(*asciidoc.Section) bool) {
		parse.SkimFunc(reader, s, reader.Children(s), func(s *asciidoc.Section) bool {
			if library.SectionType(s) != matter.SectionCommand {
				return false
			}
			return !yield(s)
		})
	}
}

func (library *Library) toCommands(spec *Specification, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (commands matter.CommandSet, err error) {

	t := FindFirstTable(reader, s)
	if t == nil {
		return nil, nil
	}

	cf := commandFactory{}
	commands, err = buildList(spec, library, reader, d, s, t, commands, &cf, parent)

	cmdMap := make(map[matter.Interface]map[uint64]*matter.Command)

	for _, cmd := range commands {
		if cmdMap[cmd.Direction] == nil {
			cmdMap[cmd.Direction] = make(map[uint64]*matter.Command)
		}
		existing, ok := cmdMap[cmd.Direction][cmd.ID.Value()]
		if ok {
			slog.Error("Duplicate Command ID",
				slog.String("clusterName",
					matter.EntityName(parent)),
				slog.String("commandName", cmd.Name),
				slog.String("commandID", cmd.ID.IntString()),
				slog.String("direction",
					cmd.Direction.String()),
				slog.String("previousCommand", existing.Name),
			)
			spec.addError(&DuplicateEntityIDError{Entity: cmd, Previous: existing})
		}

		cmdMap[cmd.Direction][cmd.ID.Value()] = cmd

		if cmd.Response != nil {
			for _, rc := range commands {
				if strings.EqualFold(cmd.Response.Name, rc.Name) {
					rc.Access.Invoke = cmd.Access.Invoke
					break
				}
			}
		}
	}

	return
}

func ParseCommandDirection(s string) matter.Interface {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "client => server", "server <= client":
		return matter.InterfaceServer
	case "server => client", "client <= server":
		return matter.InterfaceClient
	default:
		return matter.InterfaceUnknown
	}
}

type commandFinder struct {
	entityFinderCommon

	commands []*matter.Command
}

func newCommandFinder(commands []*matter.Command, inner entityFinder) *commandFinder {
	return &commandFinder{entityFinderCommon: entityFinderCommon{inner: inner}, commands: commands}
}

func (cf *commandFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	for _, c := range cf.commands {
		if c.Name == identifier {
			return c
		}
	}
	if cf.inner != nil {
		return cf.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (cf *commandFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for _, f := range cf.commands {

			if f == cf.identity {
				continue
			}
			if !yield(f.Name, f) {
				return
			}

		}
	})
	if cf.inner != nil {
		cf.inner.suggestIdentifiers(identifier, suggestions)
	}

}

func validateCommands(spec *Specification) {
	for cluster := range spec.Clusters {
		idsu := make(idUniqueness[*matter.Command])
		idcu := make(idUniqueness[*matter.Command])
		nu := make(nameUniqueness[*matter.Command])
		cv := make(conformanceValidation)

		for _, command := range cluster.Commands {
			switch command.Direction {
			case matter.InterfaceServer:
				idsu.check(spec, command.ID, command)
			case matter.InterfaceClient:
				idcu.check(spec, command.ID, command)
			}
			nu.check(spec, command)
			cv.add(command, command.Conformance)
			validateFields(spec, command, command.Fields)
		}

		cv.check(spec)
	}
	idsu := make(idUniqueness[*matter.Command])
	idcu := make(idUniqueness[*matter.Command])
	nu := make(nameUniqueness[*matter.Command])
	for obj := range spec.GlobalObjects {
		switch command := obj.(type) {
		case *matter.Command:
			switch command.Direction {
			case matter.InterfaceServer:
				idsu.check(spec, command.ID, command)
			case matter.InterfaceClient:
				idcu.check(spec, command.ID, command)
			}
			nu.check(spec, command)
			validateFields(spec, command, command.Fields)
		}
	}
}
