package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

type docParse struct {
	doc     *spec.Doc
	docType matter.DocType

	clusters map[*spec.Section]*clusterInfo

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

	tableCache       map[*asciidoc.Table]*tableInfo
	conformanceCache map[asciidoc.Element]conformance.Set
}

type subSection struct {
	section     *spec.Section
	table       tableInfo
	parent      *subSection
	parentIndex int
	children    []*subSection
}

type clusterInfo struct {
	classification *subSection
}

type tableInfo struct {
	element      *asciidoc.Table
	rows         []*asciidoc.TableRow
	headerRow    int
	columnMap    spec.ColumnIndex
	extraColumns []spec.ExtraColumn
}

func (ti *tableInfo) getColumnIndex(columns ...matter.TableColumn) (index int, ok bool) {
	for _, column := range columns {
		index, ok = ti.columnMap[column]
		if ok {
			return
		}
	}
	return
}

type subSectionChildPattern struct {
	suffix       string
	indexColumns []matter.TableColumn
}

func newSubSectionChildPattern(suffix string, indexColumns ...matter.TableColumn) subSectionChildPattern {
	return subSectionChildPattern{suffix: suffix, indexColumns: indexColumns}
}

func (b *Ball) parseDoc(doc *spec.Doc, docType matter.DocType, topLevelSection *spec.Section) (dp *docParse, err error) {
	dp = &docParse{
		doc:              doc,
		docType:          docType,
		clusters:         make(map[*spec.Section]*clusterInfo),
		conformanceCache: make(map[asciidoc.Element]conformance.Set),
		tableCache:       make(map[*asciidoc.Table]*tableInfo),
	}
	for _, section := range parse.FindAll[*spec.Section](topLevelSection.Elements()) {
		switch section.SecType {
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
				slog.Warn("attributes section in non-cluster doc", log.Element("path", doc.Path, section.Base))
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
				slog.Warn("features section in non-cluster doc", log.Element("path", doc.Path, section.Base))
			}
		case matter.SectionCommands:
			var commands *subSection
			commands, err = newParentSubSection(dp, section, newSubSectionChildPattern(" Command", matter.TableColumnName), newSubSectionChildPattern(" Field", matter.TableColumnName))
			if err == nil {
				dp.commands = append(dp.commands, commands)
			}
		case matter.SectionClassification:
			var classification *subSection
			classification, err = newSubSection(dp, section)
			if err == nil {
				dp.classification = append(dp.classification, classification)
				ci := getSubsectionCluster(dp, section)
				if ci != nil {
					ci.classification = classification
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
		}
		if err != nil {
			err = fmt.Errorf("error organizing subsections of section %s in %s: %w", section.Name, doc.Path, err)
			return
		}
	}
	return
}

func newSubSection(dp *docParse, section *spec.Section) (ss *subSection, err error) {
	return newParentSubSection(dp, section)
}

func newParentSubSection(dp *docParse, section *spec.Section, childPatterns ...subSectionChildPattern) (ss *subSection, err error) {
	ss = &subSection{section: section}
	ss.table, err = firstTableInfo(dp, section)
	if err != nil {
		return
	}
	if ss.table.element == nil {
		return
	}
	if len(childPatterns) > 0 {
		ss.children, err = findSubsections(dp, ss, childPatterns...)
	}
	return
}

func firstTableInfo(dp *docParse, section *spec.Section) (ti tableInfo, err error) {
	ti.element = spec.FindFirstTable(section)
	if ti.element != nil {
		ti.rows = ti.element.TableRows()
		ti.headerRow, ti.columnMap, ti.extraColumns, err = spec.MapTableColumns(dp.doc, ti.rows)
		if err != nil {
			return
		}
		dp.tableCache[ti.element] = &ti
	}
	return
}

func findSubsections(dp *docParse, parent *subSection, childPatterns ...subSectionChildPattern) (subSections []*subSection, err error) {
	if parent.table.element == nil {
		return
	}
	var index int
	var ok bool
	var childPattern subSectionChildPattern
	childPattern, childPatterns = childPatterns[0], childPatterns[1:]
	for i, ic := range childPattern.indexColumns {
		index, ok = parent.table.columnMap[ic]
		if ok {
			if i > 0 {
				// The first element of indexColumns should be the expected column name, so we'll switch out
				delete(parent.table.columnMap, matter.TableColumnID)
				parent.table.columnMap[childPattern.indexColumns[0]] = index
			}
			break
		}
	}
	if !ok {
		return
	}
	subSectionNames := make(map[string]int, len(parent.table.rows))
	for i, row := range parent.table.rows {
		subSectionName, err := spec.RenderTableCell(row.Cell(index))
		if err != nil {
			slog.Debug("could not get cell value for entity index", "err", err)
			continue
		}
		subSectionNames[subSectionName] = i
	}
	for i, ss := range parse.Skim[*spec.Section](parent.section.Elements()) {
		name := strings.TrimSuffix(ss.Name, childPattern.suffix)
		var ok bool
		if _, ok = subSectionNames[name]; !ok {
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
		subSections = append(subSections, subs)
	}
	return
}

func getSubsectionCluster(ds *docParse, section *spec.Section) *clusterInfo {
	parent, ok := section.Parent.(*spec.Section)
	if ok {
		for parent != nil {
			if parent.SecType == matter.SectionCluster {
				ci, ok := ds.clusters[parent]
				if !ok {
					ci = &clusterInfo{}
					ds.clusters[parent] = ci
				}
				return ci
			}
			if parent, ok = parent.Parent.(*spec.Section); !ok {
				break
			}
		}
	}
	return nil
}
