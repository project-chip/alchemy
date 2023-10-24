package ascii

import (
	"log/slog"

	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (s *Section) toCluster() (*matter.Cluster, error) {
	c := &matter.Cluster{}
	for _, s := range parse.Skim[*Section](s.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionClusterID:
			err = readClusterID(c, s)
		case matter.SectionAttributes:
			c.Attributes, err = s.toAttributes()
		case matter.SectionClassification:
			err = readClusterClassification(c, s)
		case matter.SectionFeatures:
			c.Features, err = s.toFeatures()
		case matter.SectionDataTypes:
			slog.Info("parsing data types")
			c.DataTypes, err = s.toDataTypes()
			slog.Info("parsed data types", "count", len(c.DataTypes))
		case matter.SectionEvents:
			c.Events, err = s.toEvents()
		case matter.SectionCommands:
			c.Commands, err = s.toCommands()
		default:
			slog.Info("unknown section type", "name", s.Name, "type", s.SecType)
		}
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func readClusterID(c *matter.Cluster, s *Section) error {
	rows, headerRowIndex, columnMap, _, err := parseFirstTable(s)
	if err != nil {
		return err
	}
	row := rows[headerRowIndex+1]
	c.ID, err = readRowValue(row, columnMap, matter.TableColumnID)
	if err != nil {
		return err
	}
	c.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
	if err != nil {
		return err
	}
	return nil
}

func readClusterClassification(c *matter.Cluster, s *Section) error {
	rows, headerRowIndex, columnMap, _, err := parseFirstTable(s)
	if err != nil {
		return err
	}
	row := rows[headerRowIndex+1]
	c.Hierarchy, err = readRowValue(row, columnMap, matter.TableColumnHierarchy)
	if err != nil {
		return err
	}
	c.Role, err = readRowValue(row, columnMap, matter.TableColumnRole)
	if err != nil {
		return err
	}
	c.Scope, err = readRowValue(row, columnMap, matter.TableColumnScope)
	if err != nil {
		return err
	}
	c.PICS, err = readRowValue(row, columnMap, matter.TableColumnPICS)
	if err != nil {
		return err
	}
	return nil
}
