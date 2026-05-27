package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

type docParse struct {
	doc     *asciidoc.Document
	library *spec.Library
	docType matter.DocType

	clusters map[*asciidoc.Section]*clusterInfo

	classification []*subSection
	clusterIDs     []*subSection
	attributes     []*subSection
	features       []*subSection

	dataTypes *subSection
	bitmaps   []*subSection
	enums     []*subSection
	structs   []*subSection

	commands []*subSection
	events   []*subSection

	deviceIDs                   []*subSection
	clusterRequirements         []*subSection
	elementRequirements         []*subSection
	deviceRequirements          []*subSection
	composedElementRequirements []*subSection
	composedClusterRequirements []*subSection
	conditionRequirements       []*subSection

	tableCache       map[*asciidoc.Table]*spec.TableInfo
	conformanceCache map[asciidoc.Element]conformance.Set
}

type subSection struct {
	section     *asciidoc.Section
	table       *spec.TableInfo
	parent      *subSection
	parentIndex int
	children    []*subSection
}

type clusterInfo struct {
	classification *subSection
}

type subSectionChildPattern struct {
	suffix       string
	indexColumns []matter.TableColumn
}

func newSubSectionChildPattern(suffix string, indexColumns ...matter.TableColumn) subSectionChildPattern {
	return subSectionChildPattern{suffix: suffix, indexColumns: indexColumns}
}

func (b *Baller) parseDoc(library *spec.Library, doc *asciidoc.Document, docType matter.DocType, topLevelSection *asciidoc.Section) (dp *docParse, err error) {
	dp = &docParse{
		doc:              doc,
		library:          library,
		docType:          docType,
		clusters:         make(map[*asciidoc.Section]*clusterInfo),
		conformanceCache: make(map[asciidoc.Element]conformance.Set),
		tableCache:       make(map[*asciidoc.Table]*spec.TableInfo),
	}
	for section := range parse.FindAll[*asciidoc.Section](doc, asciidoc.RawReader, topLevelSection) {
		sectionType := library.SectionType(section)
		switch sectionType {
		case matter.SectionCluster:
			dp.clusters[section] = &clusterInfo{}
		case matter.SectionAttributes:
			switch docType {
			case matter.DocTypeCluster:
				var attributes *subSection
				attributes, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Attribute", matter.TableColumnName))
				if err == nil {
					dp.attributes = append(dp.attributes, attributes)
				}
			default:
				slog.Warn("attributes section in non-cluster doc", log.Element("source", doc.Path, section))
			}
		case matter.SectionFeatures:
			switch docType {
			case matter.DocTypeCluster:
				var features *subSection
				features, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Feature", matter.TableColumnName))
				if err == nil {
					dp.features = append(dp.features, features)
				}
			default:
				slog.Warn("features section in non-cluster doc", log.Element("source", doc.Path, section))
			}
		case matter.SectionCommands:
			var commands *subSection
			commands, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Command", matter.TableColumnName), newSubSectionChildPattern(" Field", matter.TableColumnName))
			if err == nil {
				dp.commands = append(dp.commands, commands)
			}
		case matter.SectionClassification:
			switch docType {
			case matter.DocTypeDeviceType, matter.DocTypeBaseDeviceType:
				var deviceIDs *subSection
				deviceIDs, err = newSubSection(dp, section)
				if err == nil {
					dp.deviceIDs = append(dp.deviceIDs, deviceIDs)
				}
			default:
				var classification *subSection
				classification, err = newSubSection(dp, section)
				if err == nil {
					dp.classification = append(dp.classification, classification)
					ci := getSubsectionCluster(dp, section)
					if ci != nil {
						ci.classification = classification
					}
				}
			}
		case matter.SectionClusterID:
			var clusterIDs *subSection
			clusterIDs, err = newSubSection(dp, section)
			if err == nil {
				dp.clusterIDs = append(dp.clusterIDs, clusterIDs)
			}
		case matter.SectionDataTypes:
			dp.dataTypes, err = newSubSection(dp, section)
		case matter.SectionDataTypeBitmap:
			var bm *subSection
			bm, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Bitmap", matter.TableColumnBit, matter.TableColumnID))
			if err != nil {
				break
			}
			dp.bitmaps = append(dp.bitmaps, bm)
		case matter.SectionDataTypeEnum:
			var e *subSection
			e, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Enum", matter.TableColumnName))
			if err != nil {
				break
			}
			dp.enums = append(dp.enums, e)
		case matter.SectionDataTypeStruct:
			var e *subSection
			e, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Field", matter.TableColumnName))
			if err != nil {
				break
			}
			dp.structs = append(dp.structs, e)
		case matter.SectionEvents:
			var events *subSection
			events, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Event", matter.TableColumnName), newSubSectionChildPattern(" Field", matter.TableColumnName))
			if err == nil {
				dp.events = append(dp.events, events)
			}
		case matter.SectionConditionRequirements:
			var conditionRequirements *subSection
			conditionRequirements, err = newSubSection(dp, section)
			if err == nil {
				dp.conditionRequirements = append(dp.conditionRequirements, conditionRequirements)
			}
		case matter.SectionClusterRequirements:
			var clusterRequirements *subSection
			clusterRequirements, err = newSubSection(dp, section)
			if err == nil {
				dp.clusterRequirements = append(dp.clusterRequirements, clusterRequirements)
			}
		case matter.SectionElementRequirements:
			var elementRequirements *subSection
			elementRequirements, err = newSubSection(dp, section)
			if err == nil {
				dp.elementRequirements = append(dp.elementRequirements, elementRequirements)
			}
		case matter.SectionComposedDeviceTypeClusterRequirements:
			var composedClusterRequirements *subSection
			composedClusterRequirements, err = newSubSection(dp, section)
			if err == nil {
				dp.composedClusterRequirements = append(dp.composedClusterRequirements, composedClusterRequirements)
			}
		case matter.SectionComposedDeviceTypeElementRequirements:
			var composedElementRequirements *subSection
			composedElementRequirements, err = newSubSection(dp, section)
			if err == nil {
				dp.composedElementRequirements = append(dp.composedElementRequirements, composedElementRequirements)
			}
		}
		if err != nil {
			err = fmt.Errorf("error organizing subsections of section %s in %s: %w", library.SectionName(section), doc.Path, err)
			return
		}

	}

	return
}

func newSubSection(dp *docParse, section *asciidoc.Section) (ss *subSection, err error) {
	return newParentSubSection(dp, section)
}

func newParentSubSection(dp *docParse, section *asciidoc.Section, childPatterns ...subSectionChildPattern) (ss *subSection, err error) {
	ss = &subSection{section: section}
	ss.table, err = firstTableInfo(dp, section)
	if err != nil {
		return
	}
	if ss.table == nil || ss.table.Element == nil {
		return
	}
	if len(childPatterns) > 0 {
		ss.children, err = findSubsections(dp, ss, childPatterns...)
	}
	return
}

func firstTableInfo(dp *docParse, section *asciidoc.Section) (ti *spec.TableInfo, err error) {

	table := spec.FindFirstTable(asciidoc.RawReader, section)
	if table != nil {
		ti, err = spec.ReadTable(dp.doc, asciidoc.RawReader, table)
		if err != nil {
			return
		}
		dp.tableCache[ti.Element] = ti
	}
	return
}

func findSubsections(dp *docParse, parent *subSection, childPatterns ...subSectionChildPattern) (subSections []*subSection, err error) {
	if parent.table == nil || parent.table.Element == nil {
		return
	}
	var index int
	var ok bool
	var childPattern subSectionChildPattern
	childPattern, childPatterns = childPatterns[0], childPatterns[1:]
	for i, ic := range childPattern.indexColumns {
		index, ok = parent.table.ColumnMap[ic]
		if ok {
			if i > 0 {
				// The first element of indexColumns should be the expected column name, so we'll switch out
				delete(parent.table.ColumnMap, matter.TableColumnID)
				parent.table.ColumnMap[childPattern.indexColumns[0]] = index
			}
			break
		}
	}
	if !ok {
		return
	}
	subSectionNames := make(map[string]int, len(parent.table.Rows))
	for i, row := range parent.table.Rows {
		subSectionName, err := spec.RenderTableCell(dp.library, row.Cell(index))
		if err != nil {
			slog.Debug("could not get cell value for entity index", "err", err)
			continue
		}
		subSectionNames[subSectionName] = i
	}
	var i int
	for ss := range parse.Skim[*asciidoc.Section](asciidoc.RawReader, parent.section, parent.section.Children()) {
		name := text.TrimCaseInsensitiveSuffix(dp.library.SectionName(ss), childPattern.suffix)
		var ok bool
		if _, ok = subSectionNames[name]; !ok {
			i++
			continue
		}
		var subs *subSection
		subs, err = newParentSubSection(dp, ss, childPatterns...)
		if err != nil {
			return
		}
		subs.table, err = firstTableInfo(dp, ss)
		if err != nil {
			return
		}
		subs.parent = parent
		subs.parentIndex = i
		i++
		subSections = append(subSections, subs)
	}
	return
}

func getSubsectionCluster(docParse *docParse, section *asciidoc.Section) *clusterInfo {
	parent, ok := section.Parent().(*asciidoc.Section)
	if ok {
		for parent != nil {
			if docParse.library.SectionType(parent) == matter.SectionCluster {
				ci, ok := docParse.clusters[parent]
				if !ok {
					ci = &clusterInfo{}
					docParse.clusters[parent] = ci
				}
				return ci
			}
			if parent, ok = parent.Parent().(*asciidoc.Section); !ok {
				break
			}
		}
	}
	return nil
}

func getSubsectionDeviceType(docParse *docParse, section *asciidoc.Section) *clusterInfo {
	parent, ok := section.Parent().(*asciidoc.Section)
	if ok {
		for parent != nil {
			if docParse.library.SectionType(parent) == matter.SectionDeviceType {
				ci, ok := docParse.clusters[parent]
				if !ok {
					ci = &clusterInfo{}
					docParse.clusters[parent] = ci
				}
				return ci
			}
			if parent, ok = parent.Parent().(*asciidoc.Section); !ok {
				break
			}
		}
	}
	return nil
}
