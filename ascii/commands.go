package ascii

import (
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (s *Section) toCommands(d *Doc) (commands []*matter.Command, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		if err == NoTableFound {
			err = nil
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
		cmd.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		var dir string
		dir, err = readRowValue(row, columnMap, matter.TableColumnDirection)
		if err != nil {
			return
		}
		cmd.Direction = matter.ParseCommandDirection(dir)
		cmd.Response, err = readRowValue(row, columnMap, matter.TableColumnResponse)
		if err != nil {
			return
		}
		cmd.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}

		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cmd.Access = ParseAccess(a)
		commands = append(commands, cmd)
		commandMap[strings.ToLower(cmd.Name)] = cmd
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionCommand:

			name := strings.TrimSuffix(strings.ToLower(s.Name), " command")
			c, ok := commandMap[name]
			if !ok {
				slog.Info("unknown command", "command", name)
				continue
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
			var columnMap map[matter.TableColumn]int
			rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
			if err != nil {
				if err == NoTableFound {
					err = nil
					continue
				}
				return
			}
			c.Fields, err = d.readFields(headerRowIndex, rows, columnMap)
		}
	}
	return
}
