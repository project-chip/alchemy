package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (s *Section) toClusters(d *Doc) (models []interface{}, err error) {
	var clusters []*matter.Cluster
	var description string
	p := parse.FindFirst[*types.Paragraph](s.Elements)
	if p != nil {
		se := parse.FindFirst[*types.StringElement](p.Elements)
		if se != nil {
			description = strings.ReplaceAll(se.Content, "\n", " ")
		}
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionClusterID:

			clusters, err = readClusterIDs(s)
		case matter.SectionRevisionHistory:
		}
		if err != nil {
			return nil, err
		}
	}

	for _, c := range clusters {
		c.Description = description

		elements := parse.Skim[*Section](s.Elements)
		for _, s := range elements {
			switch s.SecType {
			case matter.SectionDataTypes:
				err = s.toDataTypes(d, c)
			default:
			}
			if err != nil {
				return nil, err
			}
		}
		for _, s := range elements {
			switch s.SecType {
			case matter.SectionAttributes:
				var attr []*matter.Field
				attr, err = s.toAttributes(d)
				if err == nil {
					c.Attributes = append(c.Attributes, attr...)
				}
			case matter.SectionClassification:
				err = readClusterClassification(c, s)
			case matter.SectionFeatures:
				c.Features, err = s.toFeatures(d)
			case matter.SectionEvents:
				c.Events, err = s.toEvents(d)
			case matter.SectionCommands:
				c.Commands, err = s.toCommands(d)
			case matter.SectionRevisionHistory:
				readRevisionHistory(c, s)
			default:
			}
			if err != nil {
				return nil, err
			}
		}
		for _, a := range c.Attributes {
			assignCustomModel(c, a.Type)
		}
		for _, s := range c.Structs {
			for _, f := range s.Fields {
				assignCustomModel(c, f.Type)
			}
		}
		for _, e := range c.Events {
			for _, f := range e.Fields {
				assignCustomModel(c, f.Type)
			}
		}
		for _, cmd := range c.Commands {
			for _, f := range cmd.Fields {
				assignCustomModel(c, f.Type)
			}
		}
	}
	for _, c := range clusters {
		models = append(models, c)
	}
	return models, nil
}

func assignCustomModel(c *matter.Cluster, dt *matter.DataType) {
	if dt == nil {
		return
	} else if dt.IsArray() {
		assignCustomModel(c, dt.EntryType)
		return
	} else if dt.BaseType != matter.BaseDataTypeCustom {
		return
	}
	name := dt.Name
	for _, bm := range c.Bitmaps {
		if name == bm.Name {
			dt.Model = bm
			return
		}
	}
	for _, e := range c.Enums {
		if name == e.Name {
			dt.Model = e
			return
		}
	}
	for _, s := range c.Structs {
		if name == s.Name {
			dt.Model = s
			return
		}
	}
}

func readRevisionHistory(c *matter.Cluster, s *Section) error {
	rows, headerRowIndex, columnMap, _, err := parseFirstTable(s)
	if err != nil {
		return fmt.Errorf("failed reading revision history: %w", err)
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		rev := &matter.Revision{}
		rev.Number, err = readRowValue(row, columnMap, matter.TableColumnRevision)
		if err != nil {
			return fmt.Errorf("error reading revision column on cluster %s: %w", c.Name, err)
		}
		rev.Description, err = readRowValue(row, columnMap, matter.TableColumnDescription)
		if err != nil {
			return fmt.Errorf("error reading revision description column on cluster %s: %w", c.Name, err)
		}
		c.Revisions = append(c.Revisions, rev)
	}

	return nil
}

func readClusterIDs(s *Section) ([]*matter.Cluster, error) {
	rows, headerRowIndex, columnMap, _, err := parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading cluster ID: %w", err)
	}
	var clusters []*matter.Cluster
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		c := &matter.Cluster{}
		c.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return nil, err
		}
		c.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func readClusterClassification(c *matter.Cluster, s *Section) error {
	rows, headerRowIndex, columnMap, extraColumns, err := parseFirstTable(s)
	if err != nil {
		return fmt.Errorf("failed reading classification: %w", err)
	}
	row := rows[headerRowIndex+1]
	c.Hierarchy, err = readRowValue(row, columnMap, matter.TableColumnHierarchy)
	if err != nil {
		return fmt.Errorf("error reading hierarchy column on cluster %s: %w", c.Name, err)
	}
	c.Role, err = readRowValue(row, columnMap, matter.TableColumnRole)
	if err != nil {
		return fmt.Errorf("error reading role column on cluster %s: %w", c.Name, err)
	}
	c.Scope, err = readRowValue(row, columnMap, matter.TableColumnScope)
	if err != nil {
		return fmt.Errorf("error reading scope column on cluster %s: %w", c.Name, err)
	}

	c.PICS, err = readRowValue(row, columnMap, matter.TableColumnPICS)
	if err != nil {
		return fmt.Errorf("error reading PICS column on cluster %s: %w", c.Name, err)
	}
	for _, ec := range extraColumns {
		switch ec.Name {
		case "Context":
			if len(c.Scope) == 0 {
				c.Scope, err = GetTableCellValue(row.Cells[ec.Offset])
			}
		case "Primary Transaction":
			if len(c.Scope) == 0 {
				var pt string
				pt, err = GetTableCellValue(row.Cells[ec.Offset])
				if err == nil {
					if strings.HasPrefix(pt, "Type 1") {
						c.Scope = "Endpoint"
					}
				}
			}
		}
		if err != nil {
			return fmt.Errorf("error reading extra columns on cluster %s: %w", c.Name, err)
		}
	}
	return nil
}
