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
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

var parentheticalExpressionPattern = regexp.MustCompile(`\s*\([^\)]+\)$`)

type commandFactory struct{}

func (cf *commandFactory) New(d *Doc, s *Section, ti *TableInfo, row *asciidoc.TableRow, name string, parent types.Entity) (*matter.Command, error) {
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

func (cf *commandFactory) Details(d *Doc, s *Section, pc *parseContext, c *matter.Command) (err error) {
	c.Description = getDescription(d, c, s.Elements())

	if !d.errata.Spec.IgnoreSection(s.Name, errata.SpecPurposeCommandArguments) {
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
		c.Fields, err = d.readFields(ti, types.EntityTypeCommandField, c)
		if err != nil {
			return
		}
		fieldMap := make(map[string]*matter.Field, len(c.Fields))
		for _, f := range c.Fields {
			fieldMap[f.Name] = f
		}
		err = s.mapFields(fieldMap, pc)
		if err != nil {
			return
		}
	}

	c.Name = CanonicalName(c.Name)
	return
}

func (cf *commandFactory) EntityName(s *Section) string {
	name := strings.ToLower(text.TrimCaseInsensitiveSuffix(s.Name, " Command"))
	return parentheticalExpressionPattern.ReplaceAllString(name, "")
}

func (cf *commandFactory) Children(d *Doc, s *Section) iter.Seq[*Section] {
	return func(yield func(*Section) bool) {
		parse.SkimFunc(s.Elements(), func(s *Section) bool {
			if s.SecType != matter.SectionCommand {
				return false
			}
			return !yield(s)
		})
	}
}

func (s *Section) toCommands(d *Doc, pc *parseContext, parent types.Entity) (commands matter.CommandSet, err error) {

	t := FindFirstTable(s)
	if t == nil {
		return nil, nil
	}

	var cf commandFactory
	commands, err = buildList(d, s, t, pc, commands, &cf, parent)

	for _, cmd := range commands {
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
