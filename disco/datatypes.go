package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type DataTypeEntry struct {
	doc              *asciidoc.Document
	name             string
	ref              string
	dataType         string
	dataTypeCategory matter.DataTypeCategory
	section          *asciidoc.Section
	typeCell         *asciidoc.TableCell
	definitionTable  *asciidoc.Table
	indexColumn      matter.TableColumn
	existing         bool
}

func getExistingDataTypes(cxt *discoContext) {
	if cxt.parsed.dataTypes == nil {
		return
	}

	for ss := range parse.FindAll[*asciidoc.Section](cxt.doc, asciidoc.RawReader, cxt.parsed.dataTypes.section) {
		name := matter.StripDataTypeSuffixes(cxt.library.SectionName(ss))
		nameKey := strings.ToLower(name)
		dataType, _ := spec.GetDataType(cxt.library, cxt.library, cxt.doc, ss)
		if dataType == nil {
			continue
		}
		dataTypeCategory := getDataTypeCategory(dataType.Name)
		cxt.potentialDataTypes[nameKey] = append(cxt.potentialDataTypes[nameKey], &DataTypeEntry{
			doc:              cxt.doc,
			name:             name,
			ref:              name,
			section:          ss,
			dataType:         dataType.Name,
			dataTypeCategory: dataTypeCategory,
			existing:         true,
			indexColumn:      getIndexColumnType(dataTypeCategory),
		})
	}
}

func (b *Baller) getPotentialDataTypes(dc *discoContext) (err error) {
	var subSections []*subSection
	subSections = append(subSections, dc.parsed.attributes...)
	subSections = append(subSections, dc.parsed.structs...)
	subSections = append(subSections, dc.parsed.commands...)
	subSections = append(subSections, dc.parsed.events...)

	for _, ss := range subSections {
		err = b.getPotentialDataTypesForSection(dc, ss)
		if err != nil {
			return
		}
	}
	return
}

func (b *Baller) getPotentialDataTypesForSection(cxt *discoContext, ss *subSection) error {
	if ss.table == nil || ss.table.Element == nil {
		slog.Debug("section has no table; skipping attempt to find data type", "sectionName", cxt.library.SectionName(ss.section))
		return nil
	}
	if cxt.errata.IgnoreSection(cxt.library.SectionName(ss.section), errata.DiscoPurposeDataTypePromoteInline) {
		return nil
	}
	sectionDataMap, err := b.getDataTypes(cxt, ss.table.ColumnMap, ss.table.Rows, ss.section)
	if err != nil {
		return err
	}
	for name, dataType := range sectionDataMap {
		if dataType.section != nil {
			cxt.potentialDataTypes[name] = append(cxt.potentialDataTypes[name], dataType)
		}
	}
	for _, child := range ss.children {
		err = b.getPotentialDataTypesForSection(cxt, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Baller) getDataTypes(cxt *discoContext, columnMap spec.ColumnIndex, rows []*asciidoc.TableRow, section *asciidoc.Section) (map[string]*DataTypeEntry, error) {
	sectionDataMap := make(map[string]*DataTypeEntry)
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil, nil
	}
	typeIndex, ok := columnMap[matter.TableColumnType]
	if !ok {
		return nil, nil
	}
	for _, row := range rows {
		cv, err := spec.RenderTableCell(cxt.library, row.Cell(nameIndex))
		if err != nil {
			continue
		}
		dtv, err := spec.RenderTableCell(cxt.library, row.Cell(typeIndex))
		if err != nil {
			continue
		}
		name := strings.TrimSpace(cv)
		nameKey := strings.ToLower(name)

		dataType := strings.TrimSpace(dtv)
		dataTypeCategory := getDataTypeCategory(dataType)

		if dataTypeCategory == matter.DataTypeCategoryUnknown {
			continue
		}

		if _, ok := sectionDataMap[nameKey]; !ok {
			sectionDataMap[nameKey] = &DataTypeEntry{
				doc:              cxt.doc,
				name:             name,
				ref:              name,
				dataType:         dataType,
				dataTypeCategory: dataTypeCategory,
				typeCell:         row.Cell(typeIndex),
			}
		}
	}
	for _, el := range section.Children() {
		if s, ok := el.(*asciidoc.Section); ok {
			name := strings.TrimSpace(matter.StripReferenceSuffixes(cxt.library.SectionName(s)))

			dataType, ok := sectionDataMap[strings.ToLower(name)]
			if !ok {
				continue
			}
			table := spec.FindFirstTable(asciidoc.RawReader, s)
			if table == nil {
				continue
			}
			ti, err := spec.ReadTable(cxt.doc, asciidoc.RawReader, table)
			if err != nil {
				return nil, fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", cxt.library.SectionName(s), err)
			}
			dataType.indexColumn = getIndexColumnType(dataType.dataTypeCategory)

			if valueIndex, ok := ti.ColumnMap[dataType.indexColumn]; !ok || valueIndex > 0 {
				continue
			}
			dataType.section = s
			dataType.definitionTable = table

		}
	}
	return sectionDataMap, nil
}

func (b *Baller) promoteDataTypes(cxt *discoContext, top *asciidoc.Section) (promoted bool, err error) {
	if !b.options.PromoteDataTypes {
		return
	}

	fields := make(map[matter.DataTypeCategory]map[string]*DataTypeEntry)
	for _, infos := range cxt.potentialDataTypes {
		if len(infos) > 1 {
			err = disambiguateDataTypes(infos)
			if err != nil {
				return
			}
		}
		for _, info := range infos {
			fieldMap, ok := fields[info.dataTypeCategory]
			if !ok {
				fieldMap = make(map[string]*DataTypeEntry)
				fields[info.dataTypeCategory] = fieldMap
			}
			fieldMap[info.name] = info
		}
	}

	if len(fields) > 0 {
		for _, dtc := range matter.DataTypeOrder {
			f, ok := fields[dtc]
			if !ok {
				continue
			}
			suffix := matter.DataTypeSuffixes[dtc]
			idColumn := matter.DataTypeIdentityColumn[dtc]
			var didPromotion bool
			didPromotion, err = b.promoteDataType(cxt, top, suffix, f, idColumn, dtc)
			if err != nil {
				return
			}
			promoted = didPromotion || promoted
		}
	}
	return
}

func getIndexColumnType(dataTypeCategory matter.DataTypeCategory) matter.TableColumn {
	switch dataTypeCategory {
	case matter.DataTypeCategoryEnum:
		return matter.TableColumnValue
	case matter.DataTypeCategoryBitmap:
		return matter.TableColumnBit
	}
	return matter.TableColumnUnknown
}

func getDataTypeCategory(dataType string) matter.DataTypeCategory {
	switch dataType {
	case "enum8", "enum16", "enum32":
		return matter.DataTypeCategoryEnum
	case "map8", "map16", "map32":
		return matter.DataTypeCategoryBitmap
	}
	return matter.DataTypeCategoryUnknown
}

func (b *Baller) promoteDataType(cxt *discoContext, top *asciidoc.Section, suffix string, dataTypeFields map[string]*DataTypeEntry, firstColumnType matter.TableColumn, dtc matter.DataTypeCategory) (promoted bool, err error) {
	if dataTypeFields == nil {
		return
	}
	var dataTypesSection *asciidoc.Section
	var entityType types.EntityType
	switch dtc {
	case matter.DataTypeCategoryBitmap:
		entityType = types.EntityTypeBitmapValue
	case matter.DataTypeCategoryEnum:
		entityType = types.EntityTypeEnumValue
	case matter.DataTypeCategoryStruct:
		entityType = types.EntityTypeStructField
	}
	for _, dt := range dataTypeFields {
		if dt.existing {
			continue
		}

		if dt.section == nil {
			continue
		}
		table := spec.FindFirstTable(asciidoc.RawReader, dt.section)
		if table == nil {
			continue
		}
		var ti *spec.TableInfo
		ti, err = spec.ReadTable(cxt.doc, asciidoc.RawReader, table)
		if err != nil {
			err = fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", cxt.library.SectionName(dt.section), err)
			return
		}
		if valueIndex, ok := ti.ColumnMap[firstColumnType]; !ok || valueIndex > 0 {
			continue
		}

		summaryIndex, hasSummaryColumn := ti.ColumnMap[matter.TableColumnSummary]
		if !hasSummaryColumn {
			descriptionIndex, hasDescriptionColumn := ti.ColumnMap[matter.TableColumnDescription]
			if hasDescriptionColumn {
				// Use the description column as the summary
				delete(ti.ColumnMap, matter.TableColumnDescription)
				ti.ColumnMap[matter.TableColumnSummary] = descriptionIndex
				summaryIndex = descriptionIndex
				err = b.renameTableHeaderCells(cxt, dt.section, ti, nil)
				if err != nil {
					return
				}
			} else if len(ti.ExtraColumns) > 0 {
				// Hrm, no summary or description on this promoted data type table
				// Take the first extra column and rename it
				summaryIndex = ti.ExtraColumns[0].Offset
				ti.ColumnMap[matter.TableColumnSummary] = summaryIndex
				err = b.renameTableHeaderCells(cxt, dt.section, ti, nil)
				if err != nil {
					return
				}
			} else {
				summaryIndex, err = b.appendColumn(cxt, ti, matter.TableColumnSummary, entityType)
				if err != nil {
					return
				}
			}
		}
		_, hasNameColumn := ti.ColumnMap[matter.TableColumnName]
		if !hasNameColumn {
			var nameIndex int
			nameIndex, err = b.appendColumn(cxt, ti, matter.TableColumnName, entityType)
			if err != nil {
				return
			}
			err = copyCells(cxt, ti.Rows, ti.HeaderRowIndex, summaryIndex, nameIndex, matter.Case)
			if err != nil {
				return
			}
		}

		dataTypeName := spec.CanonicalName(dt.name + suffix)

		title := asciidoc.NewString(dataTypeName + " Type")

		if dataTypesSection == nil {
			dataTypesSection, err = ensureDataTypesSection(cxt, cxt.doc, top)
			if err != nil {
				return
			}
		}

		var removedTable bool
		parse.Filter(dt.section, func(parent asciidoc.Parent, el asciidoc.Element) (remove bool, replace asciidoc.Elements, shortCircuit bool) {
			if t, ok := el.(*asciidoc.Table); ok && table == t {
				removedTable = true
				remove = true
				shortCircuit = true
				return
			}
			return
		})

		if !removedTable {
			err = fmt.Errorf("unable to relocate enum value table")
			return
		}

		dataTypeSection := asciidoc.NewSection(asciidoc.Elements{title}, dataTypesSection.Level+1)

		se := asciidoc.NewString(fmt.Sprintf("This data type is derived from %s", dt.dataType))
		p := asciidoc.NewParagraph()
		p.SetChildren(asciidoc.Elements{se})
		dataTypeSection.Append(p)
		bl := asciidoc.NewEmptyLine("")
		dataTypeSection.Append(bl)
		dataTypeSection.Append(table)
		dataTypeSection.Append(bl)

		//newAttr := make(asciidoc.AttributeList)
		tableIDAttribute := table.GetAttributeByName(asciidoc.AttributeNameID)
		var newID string
		if tableIDAttribute != nil {
			// Reuse the table's ID if it has one, so existing links get updated
			newID = tableIDAttribute.Raw()
		} else {
			newID = "ref_" + dt.ref + suffix + ", " + dt.name + suffix
		}

		dataTypeSection.AppendAttribute(asciidoc.NewNamedAttribute(string(asciidoc.AttributeNameID), asciidoc.Elements{asciidoc.NewString(newID)}, asciidoc.AttributeQuoteTypeDouble))

		dataTypesSection.Append(dataTypeSection)
		switch dt.dataTypeCategory {
		case matter.DataTypeCategoryBitmap:
			cxt.library.SetSectionType(dataTypeSection, matter.SectionDataTypeBitmap)
		case matter.DataTypeCategoryEnum:
			cxt.library.SetSectionType(dataTypeSection, matter.SectionDataTypeEnum)
		}

		table.DeleteAttribute(asciidoc.AttributeNameID)
		table.DeleteAttribute(asciidoc.AttributeNameTitle)

		icr := asciidoc.NewCrossReference(asciidoc.NewStringElements(newID), asciidoc.CrossReferenceFormatNatural)
		dt.typeCell.SetChildren(asciidoc.Elements{icr})
		promoted = true
	}
	return
}

func ensureDataTypesSection(cxt *discoContext, doc *asciidoc.Document, top *asciidoc.Section) (*asciidoc.Section, error) {
	dataTypesSection := cxt.library.FindSectionByType(cxt.library, doc, top, matter.SectionDataTypes)
	if dataTypesSection != nil {
		return dataTypesSection, nil
	}
	title := asciidoc.NewString(matter.CanonicalSectionTypeName(matter.SectionDataTypes))

	dataTypesSection = asciidoc.NewSection(asciidoc.Elements{title}, top.Level+1)
	dataTypesSection.Append(asciidoc.NewEmptyLine(""))
	top.Append(dataTypesSection)
	cxt.library.SetSectionType(dataTypesSection, matter.SectionDataTypes)
	return dataTypesSection, nil
}

func disambiguateDataTypes(infos []*DataTypeEntry) error {
	parents := make([]any, len(infos))
	dataTypeNames := make([]string, len(infos))
	dataTypeRefs := make([]string, len(infos))
	for i, info := range infos {
		parents[i] = info.section
		dataTypeNames[i] = info.name
		dataTypeRefs[i] = info.ref
	}
	parentSections := make([]*asciidoc.Section, len(infos))
	for {
		for i := range infos {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate reference: %s in %T with invalid parent", dataTypeNames[i], parents[i])
			}
			parentSections[i] = parentSection
			refParentID := strings.TrimSpace(matter.StripReferenceSuffixes(spec.ReferenceName(asciidoc.RawReader, parentSection)))
			dataTypeNames[i] = refParentID + dataTypeNames[i]
			dataTypeRefs[i] = refParentID + dataTypeNames[i]
		}
		ids := make(map[string]struct{})
		var duplicateIds bool
		for _, refID := range dataTypeRefs {
			if _, ok := ids[refID]; ok {
				duplicateIds = true
			}
			ids[refID] = struct{}{}
		}
		if duplicateIds {
			for i, info := range infos {
				parents[i] = parentSections[i].Parent
				dataTypeNames[i] = info.name
				dataTypeRefs[i] = info.ref
			}
		} else {
			break
		}
	}
	for i, info := range infos {
		info.name = dataTypeNames[i]
		info.ref = dataTypeRefs[i]
	}
	return nil
}

func (b *Baller) canonicalizeDataTypeSectionName(cxt *discoContext, s *asciidoc.Section, dataTypeName string) {
	name := cxt.library.SectionName(s)

	if cxt.errata.IgnoreSection(name, errata.DiscoPurposeDataTypeRename) {
		return
	}
	if text.HasCaseInsensitiveSuffix(name, dataTypeName+" type") {
		return
	}
	var newName string
	if text.HasCaseInsensitiveSuffix(name, dataTypeName) {
		newName = spec.CanonicalName(name) + " Type"
	} else if text.HasCaseInsensitiveSuffix(name, " type") {
		newName = spec.CanonicalName(name[:len(name)-len(" type")])
		newName += dataTypeName + " Type"
	} else {
		newName = spec.CanonicalName(name) + dataTypeName + " Type"
	}
	if name == newName {
		return
	}
	setSectionTitle(cxt, cxt.doc, s, newName)
	oldName := text.TrimCaseInsensitiveSuffix(name, " type")
	newName = text.TrimCaseInsensitiveSuffix(newName, " type")
	if oldName == newName {
		return
	}
	renameDataType(cxt, cxt.parsed.attributes, oldName, newName)
	renameDataType(cxt, cxt.parsed.commands, oldName, newName)
	renameDataType(cxt, cxt.parsed.events, oldName, newName)
}

func renameDataType(cxt *discoContext, subSections []*subSection, oldName string, newName string) {
	for _, ss := range subSections {
		renameDataType(cxt, ss.children, oldName, newName)
		if ss.table == nil || ss.table.Element == nil {
			continue
		}
		typeIndex, ok := ss.table.ColumnMap[matter.TableColumnType]
		if !ok {
			continue
		}
		for i, row := range ss.table.Rows {
			if i == ss.table.HeaderRowIndex {
				continue
			}
			typeCell := row.Cell(typeIndex)
			vc, e := spec.RenderTableCell(cxt.library, typeCell)
			if e != nil {
				continue
			}
			if strings.EqualFold(oldName, strings.TrimSpace(vc)) {
				slog.Debug("renaming data type", "oldName", oldName, "newName", newName)
				setCellString(typeCell, newName)
			}
		}
	}
}

func (b *Baller) removeMandatoryFallbacks(cxt *discoContext, ti *spec.TableInfo) {
	if !b.options.RemoveMandatoryFallbacks {
		return
	}
	fallbackIndex, hasFallback := ti.ColumnMap[matter.TableColumnFallback]
	_, hasConformance := ti.ColumnMap[matter.TableColumnConformance]
	if !hasFallback || !hasConformance {
		return
	}
	for row := range ti.Body() {
		conf := ti.ReadConformance(cxt.library, row, matter.TableColumnConformance)
		if !conformance.IsRequired(conf) {
			continue
		}
		fallbackCell := row.Cell(fallbackIndex)
		def, err := cxt.library.GetHeaderCellString(asciidoc.RawReader, fallbackCell)
		if err != nil {
			continue
		}
		def = strings.TrimSpace(def)
		if def != "" {
			setCellString(fallbackCell, "")
		}
	}
}
