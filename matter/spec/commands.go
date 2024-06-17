package spec

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/internal/log"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

var parentheticalExpressionPattern = regexp.MustCompile(`\s*\([^\)]+\)$`)

func (s *Section) toCommands(d *Doc, cluster *matter.Cluster, entityMap map[asciidoc.Attributable][]mattertypes.Entity) (commands matter.CommandSet, err error) {
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading commands table: %w", err)
		}
		return
	}
	commandMap := make(map[string]*matter.Command)
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
		cmd.Name = strings.TrimSuffix(cmd.Name, " Command")
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
		if err != nil {
			return
		}

		var a string
		a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cmd.Access, _ = ParseAccess(a, mattertypes.EntityTypeCommand)
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

	for _, s := range parse.Skim[*Section](s.Elements()) {
		switch s.SecType {
		case matter.SectionCommand:

			name := strings.TrimSuffix(strings.ToLower(s.Name), " command")
			c, ok := commandMap[name]
			if !ok {
				// Command sometimes have an parenthetical abbreviation after their name
				name = parentheticalExpressionPattern.ReplaceAllString(name, "")
				c, ok = commandMap[name]
				if !ok {
					slog.Warn("unknown command", log.Element("path", d.Path, s.Base), "command", s.Name)
					continue
				}
			}

			var desc = parse.FindFirst[*asciidoc.String](s.Elements())
			if desc != nil {
				c.Description = strings.ReplaceAll(desc.Value, "\n", " ")
			}

			var rows []*asciidoc.TableRow
			var headerRowIndex int
			var columnMap ColumnIndex
			rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
			if err != nil {
				if err == ErrNoTableFound {
					err = nil
				} else {
					slog.Warn("No valid command parameter table found", log.Element("path", d.Path, s.Base), "command", name)
					err = nil
				}
				continue
			}
			c.Fields, err = d.readFields(headerRowIndex, rows, columnMap, mattertypes.EntityTypeCommand)
			if err != nil {
				return
			}
			entityMap[s.Base] = append(entityMap[s.Base], c)
			fieldMap := make(map[string]*matter.Field, len(c.Fields))
			for _, f := range c.Fields {
				fieldMap[f.Name] = f
			}
			err = s.mapFields(fieldMap, entityMap)
			if err != nil {
				return
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
