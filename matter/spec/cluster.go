package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toClusters(d *Doc, pc *parseContext) (err error) {
	var clusters []*matter.Cluster
	var description = getDescription(d, s.Elements())

	elements := parse.Skim[*Section](s.Elements())

	// Find the cluster ID section and read the IDs from the table
	for _, s := range elements {
		switch s.SecType {
		case matter.SectionClusterID:
			clusters, err = readClusterIDs(d, s)
		}
		if err != nil {
			return
		}
	}

	var parentEntity types.Entity
	var clusterGroup *matter.ClusterGroup
	switch len(clusters) {
	case 0:
		return
	case 1:
		// There's just the one cluster ID
		cluster := clusters[0]
		parentEntity = cluster
		sectionClusterName := toClusterName(s.Name)
		if cluster.Name != sectionClusterName {
			if clusterNamesEquivalent(cluster.Name, s.Name) {
				cluster.Name = sectionClusterName
			} else {
				clusterGroup = matter.NewClusterGroup(s.Name, s.Base, clusters)
				parentEntity = clusterGroup
			}
		}
	default:
		// There's more than one cluster ID, so this is a group of similar clusters
		clusterGroup = matter.NewClusterGroup(s.Name, s.Base, clusters)
		parentEntity = clusterGroup
	}

	var features *matter.Features
	var dataTypes []types.Entity
	for _, s := range elements {
		switch s.SecType {
		case matter.SectionDataTypes, matter.SectionStatusCodes:
			var dts []types.Entity
			dts, err = s.toDataTypes(d, pc, parentEntity)
			if err == nil {
				dataTypes = append(dataTypes, dts...)
			}
		case matter.SectionFeatures:
			features, err = s.toFeatures(d, pc)
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}

	if clusterGroup != nil {
		pc.addRootEntity(clusterGroup, s.Base)

		clusterGroup.AddDataTypes(dataTypes...)

		for _, s := range elements {
			switch s.SecType {
			case matter.SectionClassification:
				err = readClusterClassification(d, clusterGroup.Name, &clusterGroup.ClusterClassification, s)
			}
			if err != nil {
				err = fmt.Errorf("error reading section in %s: %w", d.Path, err)
				return
			}
		}
	} else {
		pc.addRootEntity(clusters[0], s.Base)
	}

	for _, c := range clusters {
		c.Description = description
		c.AddDataTypes(dataTypes...)
		c.Features = features

		for _, s := range elements {
			switch s.SecType {
			case matter.SectionClassification:
				err = readClusterClassification(d, c.Name, &c.ClusterClassification, s)
			}
			if err != nil {
				err = fmt.Errorf("error reading section in %s: %w", d.Path, err)
				return
			}
		}
		for _, s := range elements {
			switch s.SecType {
			case matter.SectionAttributes:
				var attr []*matter.Field
				attr, err = s.toAttributes(d, c, pc)
				if err == nil {
					c.Attributes = append(c.Attributes, attr...)
				}
			case matter.SectionEvents:
				c.Events, err = s.toEvents(d, pc, parentEntity)
			case matter.SectionCommands:
				c.Commands, err = s.toCommands(d, pc, parentEntity)
			case matter.SectionRevisionHistory:
				c.Revisions, err = readRevisionHistory(d, s)
			case matter.SectionDerivedClusterNamespace:
				err = parseDerivedCluster(d, s, c)
			case matter.SectionClusterID:
			case matter.SectionDataTypes, matter.SectionFeatures, matter.SectionStatusCodes: // Handled above
			default:
				var looseEntities []types.Entity
				looseEntities, err = findLooseEntities(d, s, pc, parentEntity)
				if err != nil {
					err = fmt.Errorf("error reading section %s: %w", s.Name, err)
					return
				}
				if len(looseEntities) > 0 {
					for _, le := range looseEntities {
						switch le := le.(type) {
						case *matter.Bitmap:
							c.AddBitmaps(le)
						case *matter.Enum:
							c.AddEnums(le)
						case *matter.Struct:
							c.AddStructs(le)
						case *matter.TypeDef:
							c.TypeDefs = append(c.TypeDefs, le)
						default:
							slog.Warn("unexpected loose entity", log.Element("source", d.Path, s.Base), "entity", le)
						}
					}
				}
			}
			if err != nil {
				err = fmt.Errorf("error reading section in %s: %w", d.Path, err)
				return
			}
		}
		assignCustomDataTypes(c)
	}

	return
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

func assignCustomDataType(c *matter.Cluster, dt *types.DataType) {
	if dt == nil {
		return
	} else if dt.IsArray() {
		assignCustomDataType(c, dt.EntryType)
		return
	} else if dt.BaseType != types.BaseDataTypeCustom {
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
	for _, t := range c.TypeDefs {
		if name == t.Name {
			dt.Entity = t
			return
		}
	}
	slog.Debug("unable to find data type for field", slog.String("dataType", name), log.Type("source", dt.Source))
}

func readRevisionHistory(doc *Doc, s *Section) (revisions []*matter.Revision, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(doc, s)
	if err != nil {
		err = fmt.Errorf("failed reading revision history: %w", err)
		return
	}
	for row := range ti.Body() {
		rev := &matter.Revision{}
		rev.Number, err = ti.ReadString(row, matter.TableColumnRevision)
		if err != nil {
			err = fmt.Errorf("error reading revision column: %w", err)
			return
		}
		rev.Description, err = ti.ReadValue(row, matter.TableColumnDescription)
		if err != nil {
			err = fmt.Errorf("error reading revision description: %w", err)
			return
		}
		revisions = append(revisions, rev)
	}

	return
}

func readClusterIDs(doc *Doc, s *Section) ([]*matter.Cluster, error) {
	ti, err := parseFirstTable(doc, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading cluster ID: %w", err)
	}
	var clusters []*matter.Cluster
	for row := range ti.Body() {
		c := matter.NewCluster(s.Base)
		c.ID, err = ti.ReadID(row, matter.TableColumnID)
		if err != nil {
			return nil, err
		}
		var name string
		name, err = ti.ReadValue(row, matter.TableColumnName)
		if err != nil {
			return nil, err
		}
		c.Name = toClusterName(name)
		c.PICS, err = ti.ReadString(row, matter.TableColumnPICS, matter.TableColumnPICSCode)
		if err != nil {
			return nil, err
		}

		c.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		clusters = append(clusters, c)
	}

	return clusters, nil
}

// The canonical name of clusters doesn't have the "cluster" suffix, but is frequently included
func toClusterName(name string) string {
	return text.TrimCaseInsensitiveSuffix(name, " Cluster")
}

// Tests whether two cluster names are equivalent after cleaning
func clusterNamesEquivalent(name1 string, name2 string) bool {
	name1 = toClusterName(name1)
	name2 = toClusterName(name2)
	name1 = strings.ReplaceAll(name1, " ", "")
	name2 = strings.ReplaceAll(name2, " ", "")
	return strings.EqualFold(name1, name2)
}

func readClusterClassification(doc *Doc, name string, classification *matter.ClusterClassification, s *Section) error {
	ti, err := parseFirstTable(doc, s)
	if err != nil {
		return fmt.Errorf("failed reading classification: %w", err)
	}
	for row := range ti.Body() {
		classification.Hierarchy, err = ti.ReadString(row, matter.TableColumnHierarchy)
		if err != nil {
			return fmt.Errorf("error reading hierarchy column on cluster %s: %w", name, err)
		}
		classification.Role, err = ti.ReadString(row, matter.TableColumnRole)
		if err != nil {
			return fmt.Errorf("error reading role column on cluster %s: %w", name, err)
		}
		classification.Scope, err = ti.ReadString(row, matter.TableColumnScope, matter.TableColumnContext)
		if err != nil {
			return fmt.Errorf("error reading scope column on cluster %s: %w", name, err)
		}
		if len(classification.PICS) == 0 {
			classification.PICS, err = ti.ReadString(row, matter.TableColumnPICS, matter.TableColumnPICSCode)
			if err != nil {
				return fmt.Errorf("error reading PICS column on cluster %s: %w", name, err)
			}
		}
		classification.Quality, err = ti.ReadQuality(row, types.EntityTypeCluster, matter.TableColumnQuality)
		if err != nil {
			return fmt.Errorf("error reading Quality column on cluster %s: %w", name, err)
		}
		tableCells := row.TableCells()
		for _, ec := range ti.ExtraColumns {
			switch ec.Name {
			case "Context":
				if len(classification.Scope) == 0 {
					classification.Scope, err = RenderTableCell(tableCells[ec.Offset])
				}
			case "Primary Transaction":
				if len(classification.Scope) == 0 {
					var pt string
					pt, err = RenderTableCell(tableCells[ec.Offset])
					if err == nil {
						if strings.HasPrefix(pt, "Type 1") {
							classification.Scope = "Endpoint"
						}
					}
				}
			}
			if err != nil {
				return fmt.Errorf("error reading extra columns on cluster %s: %w", name, err)
			}
		}
		return nil
	}
	return nil
}

func parseDerivedCluster(d *Doc, s *Section, c *matter.Cluster) error {
	elements := parse.Skim[*Section](s.Elements())
	for _, s := range elements {
		switch s.SecType {
		case matter.SectionModeTags:
			en, err := s.toModeTags(d, c)
			if err != nil {
				return err
			}
			c.Enums = append(c.Enums, en)

		}
	}
	return nil
}
