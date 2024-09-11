package spec

import (
	"iter"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

var parentheticalExpressionPattern = regexp.MustCompile(`\s*\([^\)]+\)$`)

type commandFactory struct{}

func (cf *commandFactory) New(d *Doc, s *Section, row *asciidoc.TableRow, columnMap ColumnIndex, name string) (*matter.Command, error) {
	cmd := matter.NewCommand(s.Base)
	var err error
	cmd.ID, err = readRowID(row, columnMap, matter.TableColumnID)
	if err != nil {
		return nil, err
	}

	cmd.Name = text.TrimCaseInsensitiveSuffix(name, " Command")
	var dir string
	dir, err = readRowASCIIDocString(row, columnMap, matter.TableColumnDirection)
	if err != nil {
		return nil, err
	}
	cmd.Direction = ParseCommandDirection(dir)
	cmd.Response, err = readRowASCIIDocString(row, columnMap, matter.TableColumnResponse)
	if err != nil {
		return nil, err
	}
	cmd.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
	var a string
	a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
	if err != nil {
		return nil, err
	}
	var q string
	q, err = readRowASCIIDocString(row, columnMap, matter.TableColumnQuality)
	if err != nil {
		return nil, err
	}
	cmd.Quality = parseQuality(q, types.EntityTypeCommand, d, row)
	cmd.Access, _ = ParseAccess(a, types.EntityTypeCommand)
	return cmd, nil
}

func (cf *commandFactory) Details(d *Doc, s *Section, entityMap map[asciidoc.Attributable][]types.Entity, c *matter.Command) (err error) {
	c.Description = getDescription(d, s.Elements())

	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			slog.Warn("No valid command parameter table found", log.Element("path", d.Path, s.Base), "command", c.Name)
			err = nil
		}
		return
	}
	c.Fields, err = d.readFields(headerRowIndex, rows, columnMap, types.EntityTypeCommandField)
	if err != nil {
		return
	}
	fieldMap := make(map[string]*matter.Field, len(c.Fields))
	for _, f := range c.Fields {
		fieldMap[f.Name] = f
	}
	err = s.mapFields(fieldMap, entityMap)
	if err != nil {
		return
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

func (s *Section) toCommands(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (commands matter.CommandSet, err error) {

	t := FindFirstTable(s)
	if t == nil {
		return nil, nil
	}

	var cf commandFactory
	commands, err = buildList(d, s, t, entityMap, commands, &cf)

	for _, cmd := range commands {
		if cmd.Response != "" {
			for _, rc := range commands {
				if strings.EqualFold(cmd.Response, rc.Name) {
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
