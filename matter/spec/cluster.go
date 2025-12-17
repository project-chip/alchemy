package spec

import (
	"iter"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toClusters(spec *Specification, reader asciidoc.Reader, d *asciidoc.Document, section *asciidoc.Section, domain string) (entity types.Entity, err error) {
	var clusters []*matter.Cluster

	sections := parse.SkimList[*asciidoc.Section](reader, section, reader.Children(section))

	// Find the cluster ID section and read the IDs from the table
	for _, s := range sections {
		switch library.SectionType(s) {
		case matter.SectionClusterID:
			clusters, err = readClusterIDs(reader, d, s, domain)
		}
		if err != nil {
			return
		}
	}

	sectionName := library.SectionName(section)

	var parentEntity types.Entity
	var clusterGroup *matter.ClusterGroup
	switch len(clusters) {
	case 0:
		return
	case 1:
		// There's just the one cluster ID
		cluster := clusters[0]
		parentEntity = cluster
		sectionClusterName := toClusterName(sectionName)
		if cluster.Name != sectionClusterName {
			if clusterNamesEquivalent(cluster.Name, sectionName) {
				cluster.Name = sectionClusterName
			} else {
				clusterGroup = matter.NewClusterGroup(sectionName, section, clusters)
				parentEntity = clusterGroup
			}
		}
	default:
		// There's more than one cluster ID, so this is a group of similar clusters
		clusterGroup = matter.NewClusterGroup(sectionName, section, clusters)
		parentEntity = clusterGroup
	}

	var features *matter.Features
	var dataTypes []types.Entity
	for _, s := range sections {
		switch library.SectionType(s) {
		case matter.SectionDataTypes, matter.SectionStatusCodes:
			var dts []types.Entity
			dts, err = library.toDataTypes(spec, reader, d, s, parentEntity)
			if err == nil {
				dataTypes = append(dataTypes, dts...)
			}
		case matter.SectionFeatures:
			features, err = library.toFeatures(reader, d, s)
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}

	if clusterGroup != nil {
		entity = clusterGroup

		clusterGroup.AddDataTypes(dataTypes...)

		for _, section := range sections {
			switch library.SectionType(section) {
			case matter.SectionClassification:
				err = readClusterClassification(reader, d, clusterGroup.Name, &clusterGroup.ClusterClassification, section)
			}
			if err != nil {
				return
			}
		}
	} else {
		entity = clusters[0]
	}

	var description = library.getDescription(reader, d, clusters[0], section, reader.Children(section))

	for _, c := range clusters {
		c.Description = description
		c.AddDataTypes(dataTypes...)
		if features != nil {
			c.Features = features.CloneTo(c)
		}

		for _, s := range sections {
			switch library.SectionType(s) {
			case matter.SectionClassification:
				err = readClusterClassification(reader, d, c.Name, &c.ClusterClassification, s)
			}
			if err != nil {
				return
			}
		}
		for _, s := range sections {
			switch library.SectionType(s) {
			case matter.SectionAttributes:
				var attr []*matter.Field
				attr, err = library.toAttributes(spec, reader, d, s, c)
				if err == nil {
					c.Attributes = append(c.Attributes, attr...)
				}
			case matter.SectionEvents:
				c.Events, err = library.toEvents(spec, reader, d, s, parentEntity)
			case matter.SectionCommands:
				c.Commands, err = library.toCommands(spec, reader, d, s, parentEntity)
			case matter.SectionRevisionHistory:
				c.Revisions, err = readRevisionHistory(reader, d, s, c)
			case matter.SectionDerivedClusterNamespace:
				err = library.parseDerivedCluster(reader, d, s, c)
			case matter.SectionClusterID:
			case matter.SectionDataTypes, matter.SectionFeatures, matter.SectionStatusCodes: // Handled above
			default:
				var looseEntities []types.Entity
				looseEntities, err = library.findLooseEntities(spec, reader, d, s, parentEntity)
				if err != nil {
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
							slog.Warn("unexpected loose entity", log.Element("source", d.Path, s), "entity", le)
						}
					}
				}
			}
			if err != nil {
				return
			}
		}
	}

	return
}

func readRevisionHistory(reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, parent types.Entity) (revisions []*matter.Revision, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, doc, section)
	if err != nil {
		err = newGenericParseError(section, "failed reading revision history: %w", err)
		return
	}
	for row := range ti.ContentRows() {
		rev := matter.NewRevision(parent, row)
		var number string
		number, err = ti.ReadString(reader, row, matter.TableColumnRevision)
		if err != nil {
			err = newGenericParseError(row, "error reading revision column: %w", err)
			return
		}
		rev.Number = matter.ParseNumber(number)
		if !rev.Number.Valid() {
			err = newGenericParseError(row, "error reading revision column: %w", err)
			return
		}
		rev.Description, err = ti.ReadValue(reader, row, matter.TableColumnDescription)
		if err != nil {
			err = newGenericParseError(row, "error reading revision description: %w", err)
			return
		}
		revisions = append(revisions, rev)
	}

	return
}

func readClusterIDs(reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, domain string) ([]*matter.Cluster, error) {
	ti, err := parseFirstTable(reader, doc, section)
	if err != nil {
		return nil, newGenericParseError(section, "failed reading cluster ID: %w", err)
	}
	var clusters []*matter.Cluster
	for row := range ti.ContentRows() {
		c := matter.NewCluster(section)
		c.Domain = domain
		c.ID, err = ti.ReadID(reader, row, matter.IDColumns.Cluster...)
		if err != nil {
			return nil, err
		}
		var name string
		name, err = ti.ReadValue(reader, row, matter.TableColumnClusterName, matter.TableColumnName)
		if err != nil {
			return nil, err
		}
		c.Name = toClusterName(name)
		c.PICS, err = ti.ReadString(reader, row, matter.TableColumnPICS, matter.TableColumnPICSCode)
		if err != nil {
			return nil, err
		}

		c.Conformance = ti.ReadConformance(reader, row, matter.TableColumnConformance)
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

func isBaseOrDerivedCluster(e types.Entity) bool {
	switch e := e.(type) {
	case *matter.Cluster:
		if text.HasCaseInsensitiveSuffix(e.Name, " Base") {
			return true
		}
		if e.Hierarchy != "Base" {
			return true
		}
	}
	return false
}

func readClusterClassification(reader asciidoc.Reader, doc *asciidoc.Document, name string, classification *matter.ClusterClassification, s *asciidoc.Section) error {
	ti, err := parseFirstTable(reader, doc, s)
	if err != nil {
		return newGenericParseError(s, "failed reading classification: %w", err)
	}
	for row := range ti.ContentRows() {
		classification.Hierarchy, err = ti.ReadString(reader, row, matter.TableColumnHierarchy)
		if err != nil {
			return newGenericParseError(row, "error reading hierarchy column on cluster %s: %w", name, err)
		}
		classification.Role, err = ti.ReadString(reader, row, matter.TableColumnRole)
		if err != nil {
			return newGenericParseError(row, "error reading role column on cluster %s: %w", name, err)
		}
		classification.Scope, err = ti.ReadString(reader, row, matter.TableColumnScope, matter.TableColumnContext)
		if err != nil {
			newGenericParseError(row, "error reading scope column on cluster %s: %w", name, err)
		}
		if len(classification.PICS) == 0 {
			classification.PICS, err = ti.ReadString(reader, row, matter.TableColumnPICS, matter.TableColumnPICSCode)
			if err != nil {
				return newGenericParseError(row, "error reading PICS column on cluster %s: %w", name, err)
			}
		}
		classification.Quality, err = ti.ReadQuality(reader, row, types.EntityTypeCluster, matter.TableColumnQuality)
		if err != nil {
			return newGenericParseError(row, "error reading Quality column on cluster %s: %w", name, err)
		}
		tableCells := row.TableCells()
		for _, ec := range ti.ExtraColumns {
			switch ec.Name {
			case "Context":
				if len(classification.Scope) == 0 {
					classification.Scope, err = RenderTableCell(reader, tableCells[ec.Offset])
				}
			case "Primary Transaction":
				if len(classification.Scope) == 0 {
					var pt string
					pt, err = RenderTableCell(reader, tableCells[ec.Offset])
					if err == nil {
						if strings.HasPrefix(pt, "Type 1") {
							classification.Scope = "Endpoint"
						}
					}
				}
			}
			if err != nil {
				return newGenericParseError(row, "error reading extra columns on cluster %s: %w", name, err)
			}
		}
		return nil
	}
	return nil
}

func (library *Library) parseDerivedCluster(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, c *matter.Cluster) error {
	elements := parse.Skim[*asciidoc.Section](reader, s, reader.Children(s))
	for s := range elements {
		switch library.SectionType(s) {
		case matter.SectionModeTags:
			en, err := library.toModeTags(reader, d, s, c)
			if err != nil {
				return err
			}
			c.Enums = append(c.Enums, en)
		case matter.SectionStatusCodes:
			en, err := library.toStatusCodes(reader, d, s, c)
			if err != nil {
				return err
			}
			c.Enums = append(c.Enums, en)
		}
	}
	return nil
}

type clusterEntityFinder struct {
	entityFinderCommon

	cluster *matter.Cluster
}

func newClusterEntityFinder(cluster *matter.Cluster, inner entityFinder) *clusterEntityFinder {
	return &clusterEntityFinder{entityFinderCommon: entityFinderCommon{inner: inner}, cluster: cluster}
}

func (cf *clusterEntityFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {

	e := cf.findEntityInCluster(cf.cluster, identifier)
	if e != nil {
		return e
	}
	if cf.cluster.ParentCluster != nil {
		e = cf.findEntityInCluster(cf.cluster.ParentCluster, identifier)
		if e != nil {
			return e
		}
	}
	if cf.inner != nil {
		return cf.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (cf *clusterEntityFinder) findEntityInCluster(cluster *matter.Cluster, identifier string) types.Entity {
	if cluster.Features != nil {
		for f := range cluster.Features.FeatureBits() {
			if f.Code == identifier && f != cf.identity {
				return f
			}
		}
	}
	for _, bm := range cluster.Bitmaps {
		if bm.Name == identifier && bm != cf.identity {
			return bm
		}
	}
	for _, en := range cluster.Enums {
		if en.Name == identifier && en != cf.identity {
			return en
		}
	}
	for _, s := range cluster.Structs {
		if s.Name == identifier && s != cf.identity {
			return s
		}
	}
	return nil
}

func (cef *clusterEntityFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, cef.suggestEntityInCluster(cef.cluster))
	if cef.cluster.ParentCluster != nil {
		suggest.PossibleEntities(identifier, suggestions, cef.suggestEntityInCluster(cef.cluster.ParentCluster))
	}
	if cef.inner != nil {
		cef.inner.suggestIdentifiers(identifier, suggestions)
	}
}

func (cf *clusterEntityFinder) suggestEntityInCluster(cluster *matter.Cluster) iter.Seq2[string, types.Entity] {
	return func(yield func(string, types.Entity) bool) {
		if cluster.Features != nil {
			for f := range cluster.Features.FeatureBits() {
				if f != cf.identity && !yield(f.Code, f) {
					return
				}
			}
		}
		for _, bm := range cluster.Bitmaps {
			if bm != cf.identity && !yield(bm.Name, bm) {
				return
			}
		}
		for _, en := range cluster.Enums {
			if en != cf.identity && !yield(en.Name, en) {
				return
			}
		}
		for _, s := range cluster.Structs {
			if s != cf.identity && !yield(s.Name, s) {
				return
			}
		}
	}

}

func (cef *clusterEntityFinder) findEntityByReference(reference string, label string, source log.Source) (e types.Entity) {
	if cef.inner != nil {
		e = cef.inner.findEntityByReference(reference, label, source)
	}
	if cef.cluster.Hierarchy != "Base" {
		switch en := e.(type) {
		case *matter.Bitmap:
			local := cef.findEntityInCluster(cef.cluster, en.Name)
			if local != nil {
				e = local
				return
			}
		case *matter.Enum:
			local := cef.findEntityInCluster(cef.cluster, en.Name)
			if local != nil {
				e = local
				return
			}
		case *matter.Struct:
			local := cef.findEntityInCluster(cef.cluster, en.Name)
			if local != nil {
				e = local
				return
			}
		}
	}

	return
}
