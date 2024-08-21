package spec

import (
	"fmt"
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

func (s *Section) toCommands(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (commands matter.CommandSet, err error) {

	t := FindFirstTable(s)
	if t == nil {

		return
	}

	var commandMap map[string]*matter.Command
	commands, commandMap, err = s.buildCommands(d, t)
	if err != nil {
		err = fmt.Errorf("error reading commands table: %w", err)
		return
	}

	for _, s := range parse.Skim[*Section](s.Elements()) {
		switch s.SecType {
		case matter.SectionCommand:
			var c *matter.Command
			c, err = s.toCommand(d, commandMap, entityMap)
			if err != nil {
				return
			}
			if c != nil {
				entityMap[s.Base] = append(entityMap[s.Base], c)
			}
		}
	}
	return
}

func (s *Section) buildCommands(d *Doc, t *asciidoc.Table) (commands matter.CommandSet, commandMap map[string]*matter.Command, err error) {
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseTable(d, s, t)
	if err != nil {
		return
	}
	commandMap = make(map[string]*matter.Command)
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		cmd := &matter.Command{}
		cmd.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		cmd.Name, err = ReadRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		cmd.Name = text.TrimCaseInsensitiveSuffix(cmd.Name, " Command")
		var dir string
		dir, err = readRowASCIIDocString(row, columnMap, matter.TableColumnDirection)
		if err != nil {
			return
		}
		cmd.Direction = ParseCommandDirection(dir)
		cmd.Response, err = readRowASCIIDocString(row, columnMap, matter.TableColumnResponse)
		if err != nil {
			return
		}
		cmd.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		var a string
		a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cmd.Access, _ = ParseAccess(a, types.EntityTypeCommand)
		commands = append(commands, cmd)
		commandMap[strings.ToLower(cmd.Name)] = cmd
	}

	for _, cmd := range commands {
		if cmd.Response != "" {
			if responseCommand, ok := commandMap[strings.ToLower(cmd.Response)]; ok && responseCommand.Access.Invoke == matter.PrivilegeUnknown {
				responseCommand.Access.Invoke = cmd.Access.Invoke
			}
		}
	}
	return
}

func (s *Section) toCommand(d *Doc, commandMap map[string]*matter.Command, entityMap map[asciidoc.Attributable][]types.Entity) (*matter.Command, error) {
	name := strings.ToLower(text.TrimCaseInsensitiveSuffix(s.Name, " Command"))
	c, ok := commandMap[name]
	if !ok {
		// Command sometimes have an parenthetical abbreviation after their name
		name = parentheticalExpressionPattern.ReplaceAllString(name, "")
		c, ok = commandMap[name]
		if !ok {
			slog.Warn("unknown command", log.Element("path", d.Path, s.Base), "command", s.Name)
			return nil, nil
		}
	}

	c.Description = getDescription(d, s.Elements())

	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	var err error
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			slog.Warn("No valid command parameter table found", log.Element("path", d.Path, s.Base), "command", name)
			err = nil
		}
		return nil, nil
	}
	c.Fields, err = d.readFields(headerRowIndex, rows, columnMap, types.EntityTypeCommand)
	if err != nil {
		return nil, err
	}
	fieldMap := make(map[string]*matter.Field, len(c.Fields))
	for _, f := range c.Fields {
		fieldMap[f.Name] = f
	}
	err = s.mapFields(fieldMap, entityMap)
	if err != nil {
		return nil, err
	}
	c.Name = specName(c.Name)
	return c, nil
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
