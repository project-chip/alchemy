package spec

import (
	"iter"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

var parentheticalExpressionPattern = regexp.MustCompile(`\s*\([^\)]+\)$`)

type commandFactory struct{}

func (cf *commandFactory) New(spec *Specification, d *Doc, s *Section, ti *TableInfo, row *asciidoc.TableRow, name string, parent types.Entity) (*matter.Command, error) {
	cmd := matter.NewCommand(s.Base, parent)
	var err error
	cmd.ID, err = ti.ReadID(row, matter.TableColumnID)
	if err != nil {
		return nil, err
	}

	cmd.Name = text.TrimCaseInsensitiveSuffix(name, " Command")
	var dir string
	dir, err = ti.ReadString(row, matter.TableColumnDirection)
	if err != nil {
		return nil, err
	}
	cmd.Direction = ParseCommandDirection(dir)
	cmd.Response, _ = ti.ReadDataType(row, matter.TableColumnResponse)
	if cmd.Response != nil {
		cmd.Response.Name = text.TrimCaseInsensitiveSuffix(cmd.Response.Name, " Command")
	}
	cmd.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
	var a string
	a, err = ti.ReadString(row, matter.TableColumnAccess)
	if err != nil {
		return nil, err
	}
	cmd.Quality, err = ti.ReadQuality(row, types.EntityTypeCommand, matter.TableColumnQuality)
	if err != nil {
		return nil, err
	}
	cmd.Access, _ = ParseAccess(a, types.EntityTypeCommand)
	return cmd, nil
}

func (cf *commandFactory) Details(spec *Specification, d *Doc, s *Section, pc *parseContext, c *matter.Command) (err error) {
	return readCommand(pc, spec, d, s, c)
}

func readCommand(pc *parseContext, spec *Specification, d *Doc, s *Section, c *matter.Command) (err error) {
	c.Description = getDescription(d, c, s.Children())

	c.Name = CanonicalName(c.Name)

	if d.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeCommandArguments) {
		return
	}
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			slog.Warn("No valid command parameter table found", log.Element("source", d.Path, s.Base), "command", c.Name)
			err = nil
		}
		return
	}
	var fieldMap map[string]*matter.Field
	c.Fields, fieldMap, err = d.readFields(spec, ti, types.EntityTypeCommandField, c)
	if err != nil {
		return
	}
	err = s.mapFields(fieldMap, pc)
	if err != nil {
		return
	}
	return
}

func (cf *commandFactory) EntityName(s *Section) string {
	name := strings.ToLower(text.TrimCaseInsensitiveSuffix(s.Name, " Command"))
	return parentheticalExpressionPattern.ReplaceAllString(name, "")
}

func (cf *commandFactory) Children(d *Doc, s *Section) iter.Seq[*Section] {
	return func(yield func(*Section) bool) {
		parse.SkimFunc(s.Children(), func(s *Section) bool {
			if s.SecType != matter.SectionCommand {
				return false
			}
			return !yield(s)
		})
	}
}

func (s *Section) toCommands(spec *Specification, d *Doc, pc *parseContext, parent types.Entity) (commands matter.CommandSet, err error) {

	t := FindFirstTable(s)
	if t == nil {
		return nil, nil
	}

	var cf commandFactory
	commands, err = buildList(spec, d, s, t, pc, commands, &cf, parent)

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
	for c := range spec.Clusters {
		for _, s := range c.Commands {
			validateFields(spec, s, s.Fields)
		}
	}
	for obj := range spec.GlobalObjects {
		switch obj := obj.(type) {
		case *matter.Command:
			validateFields(spec, obj, obj.Fields)
		}
	}
}
