package ascii

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
)

var parentheticalExpressionPattern = regexp.MustCompile(`\s*\([^\)]+\)$`)

func (s *Section) toCommands(d *Doc, entityMap map[types.WithAttributes][]mattertypes.Entity) (commands matter.CommandSet, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == NoTableFound {
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
		cmd.Name, err = readRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		cmd.Name = strings.TrimSuffix(cmd.Name, " Command")
		var dir string
		dir, err = readRowAsciiDocString(row, columnMap, matter.TableColumnDirection)
		if err != nil {
			return
		}
		cmd.Direction = parseCommandDirection(dir)
		cmd.Response, err = readRowAsciiDocString(row, columnMap, matter.TableColumnResponse)
		if err != nil {
			return
		}
		cmd.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}

		var a string
		a, err = readRowAsciiDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cmd.Access = ParseAccess(a, mattertypes.EntityTypeCommand)
		if cmd.Access.Invoke == matter.PrivilegeUnknown && cmd.Direction == matter.InterfaceClient {
			// Response commands sometimes leave out the privilege, so we're assuming it's operate
			cmd.Access.Invoke = matter.PrivilegeOperate
		}
		commands = append(commands, cmd)
		commandMap[strings.ToLower(cmd.Name)] = cmd
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionCommand:

			name := strings.TrimSuffix(strings.ToLower(s.Name), " command")
			c, ok := commandMap[name]
			if !ok {
				// Command sometimes have an parenthetical abbreviation after their name
				name = parentheticalExpressionPattern.ReplaceAllString(name, "")
				c, ok = commandMap[name]
				if !ok {
					slog.Warn("unknown command", "path", d.Path, "command", s.Name)
					continue
				}
			}
			p := parse.FindFirst[*types.Paragraph](s.Elements)
			if p != nil {
				se := parse.FindFirst[*types.StringElement](p.Elements)
				if se != nil {
					c.Description = strings.ReplaceAll(se.Content, "\n", " ")
				}
			}

			var rows []*types.TableRow
			var headerRowIndex int
			var columnMap ColumnIndex
			rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
			if err != nil {
				if err == NoTableFound {
					err = nil
				} else {
					slog.Warn("No valid command parameter table found", "command", name, "path", d.Path)
					err = nil
				}
				continue
			}
			c.Fields, err = d.readFields(headerRowIndex, rows, columnMap, mattertypes.EntityTypeCommand)
			if err != nil {
				return
			}
			entityMap[s.Base] = append(entityMap[s.Base], c)
		}
	}
	return
}

func parseCommandDirection(s string) matter.Interface {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "client => server", "server <= client":
		return matter.InterfaceServer
	case "server => client", "client <= server":
		return matter.InterfaceClient
	default:
		return matter.InterfaceUnknown
	}
}
