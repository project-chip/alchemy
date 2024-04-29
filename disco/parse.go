package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
)

type docParse struct {
	doc     *ascii.Doc
	docType matter.DocType

	classification []*subSection
	clusterIDs     []*subSection
	attributes     []*subSection

	dataTypes *subSection
	bitmaps   []*subSection
	enums     []*subSection
	structs   []*subSection

	commands []*subSection
	events   []*subSection
}

type subSection struct {
	section     *ascii.Section
	table       tableInfo
	parent      *subSection
	parentIndex int
	children    []*subSection
}

type tableInfo struct {
	element      *elements.Table
	rows         []*elements.TableRow
	headerRow    int
	columnMap    ascii.ColumnIndex
	extraColumns []ascii.ExtraColumn
}

type subSectionChildPattern struct {
	suffix       string
	indexColumns []matter.TableColumn
}

func newSubSectionChildPattern(suffix string, indexColumns ...matter.TableColumn) subSectionChildPattern {
	return subSectionChildPattern{suffix: suffix, indexColumns: indexColumns}
}

func (b *Ball) parseDoc(doc *ascii.Doc, docType matter.DocType, topLevelSection *ascii.Section) (ds *docParse, err error) {
	ds = &docParse{doc: doc, docType: docType}
	for _, section := range parse.FindAll[*ascii.Section](topLevelSection.Elements) {
		switch section.SecType {
		case matter.SectionCluster:

		case matter.SectionAttributes:
			switch docType {
			case matter.DocTypeCluster:
				var attributes *subSection
				attributes, err = newParentSubSection(doc, section, newSubSectionChildPattern(" Attribute", matter.TableColumnName))
				if err == nil {
					ds.attributes = append(ds.attributes, attributes)
				}
			default:
				slog.Warn("attributes section in non-cluster doc", slog.String("path", doc.Path))
			}
		case matter.SectionCommands:
			var commands *subSection
			commands, err = newParentSubSection(doc, section, newSubSectionChildPattern(" Command", matter.TableColumnName), newSubSectionChildPattern(" Field", matter.TableColumnName))
			if err == nil {
				ds.commands = append(ds.commands, commands)
			}
		case matter.SectionClassification:
			var classification *subSection
			classification, err = newSubSection(doc, section)
			if err == nil {
				ds.classification = append(ds.classification, classification)
			}
		case matter.SectionClusterID:
			var clusterIDs *subSection
			clusterIDs, err = newSubSection(doc, section)
			if err == nil {
				ds.clusterIDs = append(ds.clusterIDs, clusterIDs)
			}
		case matter.SectionDataTypes:
			ds.dataTypes, err = newSubSection(doc, section)
		case matter.SectionDataTypeBitmap:
			var bm *subSection
			bm, err = newParentSubSection(doc, section, newSubSectionChildPattern(" Bitmap", matter.TableColumnBit, matter.TableColumnID))
			if err != nil {
				break
			}
			ds.bitmaps = append(ds.bitmaps, bm)
		case matter.SectionDataTypeEnum:
			var e *subSection
			e, err = newParentSubSection(doc, section, newSubSectionChildPattern(" Enum", matter.TableColumnName))
			if err != nil {
				break
			}
			ds.enums = append(ds.enums, e)
		case matter.SectionDataTypeStruct:
			var e *subSection
			e, err = newParentSubSection(doc, section, newSubSectionChildPattern(" Field", matter.TableColumnName))
			if err != nil {
				break
			}
			ds.structs = append(ds.structs, e)
		case matter.SectionEvents:
			var events *subSection
			events, err = newParentSubSection(doc, section, newSubSectionChildPattern(" Event", matter.TableColumnName), newSubSectionChildPattern(" Field", matter.TableColumnName))
			if err == nil {
				ds.events = append(ds.events, events)
			}
		}
		if err != nil {
			err = fmt.Errorf("error organizing subsections of section %s in %s: %w", section.Name, doc.Path, err)
			return
		}
	}
	return
}

func newSubSection(doc *ascii.Doc, section *ascii.Section) (ss *subSection, err error) {
	return newParentSubSection(doc, section)
}

func newParentSubSection(doc *ascii.Doc, section *ascii.Section, childPatterns ...subSectionChildPattern) (ss *subSection, err error) {
	ss = &subSection{section: section}
	ss.table, err = firstTableInfo(doc, section)
	if err != nil {
		return
	}
	if ss.table.element == nil {
		return
	}
	if len(childPatterns) > 0 {
		ss.children, err = findSubsections(doc, ss, childPatterns...)
	}
	return
}

func firstTableInfo(doc *ascii.Doc, section *ascii.Section) (ti tableInfo, err error) {
	ti.element = ascii.FindFirstTable(section)
	if ti.element != nil {
		ti.rows = ti.element.TableRows
		ti.headerRow, ti.columnMap, ti.extraColumns, err = ascii.MapTableColumns(doc, ti.rows)
		if err != nil {
			return
		}
	}
	return
}

func findSubsections(doc *ascii.Doc, parent *subSection, childPatterns ...subSectionChildPattern) (subSections []*subSection, err error) {
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
		subSectionName, err := ascii.RenderTableCell(row.TableCells[index])
		if err != nil {
			slog.Debug("could not get cell value for entity index", "err", err)
			continue
		}
		subSectionNames[subSectionName] = i
	}
	for i, ss := range parse.Skim[*ascii.Section](parent.section.Elements) {
		name := strings.TrimSuffix(ss.Name, childPattern.suffix)
		var ok bool
		if _, ok = subSectionNames[name]; !ok {
			continue
		}
		var subs *subSection
		subs, err = newParentSubSection(doc, ss, childPatterns...)
		if err != nil {
			return
		}
		subs.table, err = firstTableInfo(doc, ss)
		if err != nil {
			return
		}
		subs.parent = parent
		subs.parentIndex = i
		subSections = append(subSections, subs)
	}
	return
}
