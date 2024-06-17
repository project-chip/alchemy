package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/internal/log"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toClusters(d *Doc, entityMap map[asciidoc.Attributable][]mattertypes.Entity) (entities []mattertypes.Entity, err error) {
	var clusters []*matter.Cluster
	var description string
	p := parse.FindFirst[*asciidoc.Paragraph](s.Elements())
	if p != nil {
		se := parse.FindFirst[*asciidoc.String](p.Elements())
		if se != nil {
			description = strings.ReplaceAll(se.Value, "\n", " ")
		}
	}

	elements := parse.Skim[*Section](s.Elements())

	for _, s := range elements {
		switch s.SecType {
		case matter.SectionClusterID:
			clusters, err = readClusterIDs(d, s)
		}
		if err != nil {
			return nil, err
		}
	}

	var features *matter.Features
	var bitmaps matter.BitmapSet
	var enums matter.EnumSet
	var structs matter.StructSet
	for _, s := range elements {
		switch s.SecType {
		case matter.SectionDataTypes:
			var bs matter.BitmapSet
			var es matter.EnumSet
			var ss matter.StructSet
			bs, es, ss, err = s.toDataTypes(d, entityMap)
			if err == nil {
				bitmaps = append(bitmaps, bs...)
				enums = append(enums, es...)
				structs = append(structs, ss...)
			}
		case matter.SectionFeatures:
			features, err = s.toFeatures(d, entityMap)
		}
		if err != nil {
			return
		}
	}

	for _, c := range clusters {
		c.Description = description
		c.Bitmaps = append(c.Bitmaps, bitmaps...)
		c.Enums = append(c.Enums, enums...)
		c.Structs = append(c.Structs, structs...)
		c.Features = features

		for _, s := range elements {
			switch s.SecType {
			case matter.SectionClassification:
				err = readClusterClassification(d, c, s)

			}
			if err != nil {
				return nil, fmt.Errorf("error reading section in %s: %w", d.Path, err)
			}
		}
		for _, s := range elements {
			switch s.SecType {
			case matter.SectionAttributes:
				var attr []*matter.Field
				attr, err = s.toAttributes(d, c, entityMap)
				if err == nil {
					c.Attributes = append(c.Attributes, attr...)
				}
			case matter.SectionClassification:
				err = readClusterClassification(d, c, s)
			case matter.SectionEvents:
				c.Events, err = s.toEvents(d, c, entityMap)
			case matter.SectionCommands:
				c.Commands, err = s.toCommands(d, c, entityMap)
			case matter.SectionRevisionHistory:
				c.Revisions, err = readRevisionHistory(d, s)
			case matter.SectionDerivedClusterNamespace:
				err = parseDerivedCluster(d, s, c)
			case matter.SectionClusterID:
			case matter.SectionDataTypes, matter.SectionFeatures: // Handled above
			default:
				var looseEntities []mattertypes.Entity
				looseEntities, err = findLooseEntities(d, s, entityMap)
				if err != nil {
					return nil, fmt.Errorf("error reading section %s: %w", s.Name, err)
				}
				if len(looseEntities) > 0 {
					for _, le := range looseEntities {
						switch le := le.(type) {
						case *matter.Bitmap:
							c.Bitmaps = append(c.Bitmaps, le)
						case *matter.Enum:
							c.Enums = append(c.Enums, le)
						default:
							slog.Warn("unexpected loose entity", log.Element("path", d.Path, s.Base), "entity", le)
						}
					}
				}
			}
			if err != nil {
				return nil, fmt.Errorf("error reading section in %s: %w", d.Path, err)
			}
		}
		assignCustomDataTypes(c)
	}

	if len(clusters) == 1 {
		entities = append(entities, clusters[0])
	} else if len(clusters) > 1 {
		entities = append(entities, &matter.ClusterGroup{Clusters: clusters})
	}
	entityMap[s.Base] = append(entityMap[s.Base], entities...)
	return entities, nil
}

func assignCustomDataTypes(c *matter.Cluster) {
	for _, s := range c.Structs {
		for _, f := range s.Fields {
			assignCustomDataType(c, f.Type)
		}
	}
	for _, e := range c.Events {
		for _, f := range e.Fields {
			assignCustomDataType(c, f.Type)
		}
	}
	for _, cmd := range c.Commands {
		for _, f := range cmd.Fields {
			assignCustomDataType(c, f.Type)
		}
	}
	for _, a := range c.Attributes {
		assignCustomDataType(c, a.Type)
	}
}

func assignCustomDataType(c *matter.Cluster, dt *mattertypes.DataType) {
	if dt == nil {
		return
	} else if dt.IsArray() {
		assignCustomDataType(c, dt.EntryType)
		return
	} else if dt.BaseType != mattertypes.BaseDataTypeCustom {
		return
	}
	if dt.Entity != nil {
		return
	}
	name := dt.Name
	for _, bm := range c.Bitmaps {
		if name == bm.Name {
			dt.Entity = bm
			return
		}
	}
	for _, e := range c.Enums {
		if name == e.Name {
			dt.Entity = e
			return
		}
	}
	for _, s := range c.Structs {
		if name == s.Name {
			dt.Entity = s
			return
		}
	}
}

func readRevisionHistory(doc *Doc, s *Section) (revisions []*matter.Revision, err error) {
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(doc, s)
	if err != nil {
		err = fmt.Errorf("failed reading revision history: %w", err)
		return
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		rev := &matter.Revision{}
		rev.Number, err = readRowASCIIDocString(row, columnMap, matter.TableColumnRevision)
		if err != nil {
			err = fmt.Errorf("error reading revision column: %w", err)
			return
		}
		rev.Description, err = ReadRowValue(doc, row, columnMap, matter.TableColumnDescription)
		if err != nil {
			err = fmt.Errorf("error reading revision description: %w", err)
			return
		}
		revisions = append(revisions, rev)
	}

	return
}

func readClusterIDs(doc *Doc, s *Section) ([]*matter.Cluster, error) {
	rows, headerRowIndex, columnMap, _, err := parseFirstTable(doc, s)
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
		c.Name, err = ReadRowValue(doc, row, columnMap, matter.TableColumnName)
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func readClusterClassification(doc *Doc, c *matter.Cluster, s *Section) error {
	rows, headerRowIndex, columnMap, extraColumns, err := parseFirstTable(doc, s)
	if err != nil {
		return fmt.Errorf("failed reading classification: %w", err)
	}
	row := rows[headerRowIndex+1]
	c.Hierarchy, err = readRowASCIIDocString(row, columnMap, matter.TableColumnHierarchy)
	if err != nil {
		return fmt.Errorf("error reading hierarchy column on cluster %s: %w", c.Name, err)
	}
	c.Role, err = readRowASCIIDocString(row, columnMap, matter.TableColumnRole)
	if err != nil {
		return fmt.Errorf("error reading role column on cluster %s: %w", c.Name, err)
	}
	c.Scope, err = readRowASCIIDocString(row, columnMap, matter.TableColumnScope, matter.TableColumnContext)
	if err != nil {
		return fmt.Errorf("error reading scope column on cluster %s: %w", c.Name, err)
	}

	c.PICS, err = readRowASCIIDocString(row, columnMap, matter.TableColumnPICS)
	if err != nil {
		return fmt.Errorf("error reading PICS column on cluster %s: %w", c.Name, err)
	}
	tableCells := row.TableCells()
	for _, ec := range extraColumns {
		switch ec.Name {
		case "Context":
			if len(c.Scope) == 0 {
				c.Scope, err = RenderTableCell(tableCells[ec.Offset])
			}
		case "Primary Transaction":
			if len(c.Scope) == 0 {
				var pt string
				pt, err = RenderTableCell(tableCells[ec.Offset])
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

func parseDerivedCluster(d *Doc, s *Section, c *matter.Cluster) error {
	elements := parse.Skim[*Section](s.Elements())
	for _, s := range elements {
		switch s.SecType {
		case matter.SectionModeTags:
			en, err := s.toModeTags(d)
			if err != nil {
				return err
			}
			c.Enums = append(c.Enums, en)

		}
	}
	return nil
}
