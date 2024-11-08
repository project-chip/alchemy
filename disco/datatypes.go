package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type DataTypeEntry struct {
	name             string
	ref              string
	dataType         string
	dataTypeCategory matter.DataTypeCategory
	section          *spec.Section
	typeCell         *asciidoc.TableCell
	definitionTable  *asciidoc.Table
	indexColumn      matter.TableColumn
	existing         bool
}

func getExistingDataTypes(cxt *discoContext, dp *docParse) {
	if dp.dataTypes == nil {
		return
	}

	for _, ss := range parse.FindAll[*spec.Section](dp.dataTypes.section.Elements()) {
		name := matter.StripDataTypeSuffixes(ss.Name)
		nameKey := strings.ToLower(name)
		dataType := ss.GetDataType()
		if dataType == nil {
			continue
		}
		dataTypeCategory := getDataTypeCategory(dataType.Name)
		cxt.potentialDataTypes[nameKey] = append(cxt.potentialDataTypes[nameKey], &DataTypeEntry{
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

func (b *Ball) getPotentialDataTypes(dc *discoContext, dp *docParse) (err error) {
	var subSections []*subSection
	subSections = append(subSections, dp.attributes...)
	subSections = append(subSections, dp.structs...)
	subSections = append(subSections, dp.commands...)
	subSections = append(subSections, dp.events...)

	for _, ss := range subSections {
		err = b.getPotentialDataTypesForSection(dc, dp, ss)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) getPotentialDataTypesForSection(cxt *discoContext, dp *docParse, ss *subSection) error {
	if ss.table == nil || ss.table.Element == nil {
		slog.Debug("section has no table; skipping attempt to find data type", "sectionName", ss.section.Name)
		return nil
	}
	if b.errata.IgnoreSection(ss.section.Name, errata.DiscoPurposeDataTypePromoteInline) {
		return nil
	}
	sectionDataMap, err := b.getDataTypes(ss.table.ColumnMap, ss.table.Rows, ss.section)
	if err != nil {
		return err
	}
	for name, dataType := range sectionDataMap {
		if dataType.section != nil {
			cxt.potentialDataTypes[name] = append(cxt.potentialDataTypes[name], dataType)
		}
	}
	for _, child := range ss.children {
		err = b.getPotentialDataTypesForSection(cxt, dp, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Ball) getDataTypes(columnMap spec.ColumnIndex, rows []*asciidoc.TableRow, section *spec.Section) (map[string]*DataTypeEntry, error) {
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
		cv, err := spec.RenderTableCell(row.Cell(nameIndex))
		if err != nil {
			continue
		}
		dtv, err := spec.RenderTableCell(row.Cell(typeIndex))
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
				name:             name,
				ref:              name,
				dataType:         dataType,
				dataTypeCategory: dataTypeCategory,
				typeCell:         row.Cell(typeIndex),
			}
		}
	}
	for _, el := range section.Elements() {
		if s, ok := el.(*spec.Section); ok {
			name := strings.TrimSpace(matter.StripReferenceSuffixes(s.Name))

			dataType, ok := sectionDataMap[strings.ToLower(name)]
			if !ok {
				continue
			}
			table := spec.FindFirstTable(s)
			if table == nil {
				continue
			}
			ti, err := spec.ReadTable(b.doc, table)
			if err != nil {
				return nil, fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", s.Name, err)
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

func (b *Ball) promoteDataTypes(cxt *discoContext, top *spec.Section) (promoted bool, err error) {
	if !b.options.promoteDataTypes {
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
			didPromotion, err = b.promoteDataType(top, suffix, f, idColumn, dtc)
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

func (b *Ball) promoteDataType(top *spec.Section, suffix string, dataTypeFields map[string]*DataTypeEntry, firstColumnType matter.TableColumn, dtc matter.DataTypeCategory) (promoted bool, err error) {
	if dataTypeFields == nil {
		return
	}
	var dataTypesSection *spec.Section
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
		table := spec.FindFirstTable(dt.section)
		if table == nil {
			continue
		}
		var ti *spec.TableInfo
		ti, err = spec.ReadTable(b.doc, table)
		if err != nil {
			err = fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", dt.section.Name, err)
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
				err = b.renameTableHeaderCells(top.Doc, dt.section, ti, nil)
				if err != nil {
					return
				}
			} else if len(ti.ExtraColumns) > 0 {
				// Hrm, no summary or description on this promoted data type table
				// Take the first extra column and rename it
				summaryIndex = ti.ExtraColumns[0].Offset
				ti.ColumnMap[matter.TableColumnSummary] = summaryIndex
				err = b.renameTableHeaderCells(top.Doc, dt.section, ti, nil)
				if err != nil {
					return
				}
			} else {
				summaryIndex, err = b.appendColumn(ti, matter.TableColumnSummary, entityType)
				if err != nil {
					return
				}
			}
		}
		_, hasNameColumn := ti.ColumnMap[matter.TableColumnName]
		if !hasNameColumn {
			var nameIndex int
			nameIndex, err = b.appendColumn(ti, matter.TableColumnName, entityType)
			if err != nil {
				return
			}
			err = copyCells(ti.Rows, ti.HeaderRowIndex, summaryIndex, nameIndex, matter.Case)
			if err != nil {
				return
			}
		}

		dataTypeName := spec.CanonicalName(dt.name + suffix)

		title := asciidoc.NewString(dataTypeName + " Type")

		if dataTypesSection == nil {
			dataTypesSection, err = ensureDataTypesSection(top)
			if err != nil {
				return
			}
		}

		var removedTable bool
		parse.Filter(dt.section, func(i any) (remove bool, shortCircuit bool) {
			if t, ok := i.(*asciidoc.Table); ok && table == t {
				removedTable = true
				return true, true
			}
			return false, false
		})

		if !removedTable {
			err = fmt.Errorf("unable to relocate enum value table")
			return
		}

		dataTypeSection := asciidoc.NewSection(asciidoc.Set{title}, dataTypesSection.Base.Level+1)

		se := asciidoc.NewString(fmt.Sprintf("This data type is derived from %s", dt.dataType))
		p := asciidoc.NewParagraph()
		p.SetElements(asciidoc.Set{se})
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

		dataTypeSection.AppendAttribute(asciidoc.NewNamedAttribute(string(asciidoc.AttributeNameID), asciidoc.Set{asciidoc.NewString(newID)}, asciidoc.AttributeQuoteTypeDouble))

		var s *spec.Section
		s, err = spec.NewSection(top.Doc, dataTypesSection, dataTypeSection)

		if err != nil {
			return
		}
		switch dt.dataTypeCategory {
		case matter.DataTypeCategoryBitmap:
			s.SecType = matter.SectionDataTypeBitmap
		case matter.DataTypeCategoryEnum:
			s.SecType = matter.SectionDataTypeEnum
		}

		dataTypesSection.AppendSection(s)

		table.DeleteAttribute(asciidoc.AttributeNameID)
		table.DeleteAttribute(asciidoc.AttributeNameTitle)

		icr := asciidoc.NewCrossReference(newID)
		dt.typeCell.SetElements(asciidoc.Set{icr})
		promoted = true
	}
	return
}

func ensureDataTypesSection(top *spec.Section) (*spec.Section, error) {
	dataTypesSection := spec.FindSectionByType(top, matter.SectionDataTypes)
	if dataTypesSection != nil {
		return dataTypesSection, nil
	}
	title := asciidoc.NewString(matter.SectionTypeName(matter.SectionDataTypes))

	ts := asciidoc.NewSection(asciidoc.Set{title}, top.Base.Level+1)
	ts.Append(asciidoc.NewEmptyLine(""))
	dataTypesSection, err := spec.NewSection(top.Doc, top, ts)
	if err != nil {
		return nil, err
	}
	dataTypesSection.SecType = matter.SectionDataTypes
	top.AppendSection(dataTypesSection)
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
	parentSections := make([]*spec.Section, len(infos))
	for {
		for i := range infos {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate reference: %s in %T with invalid parent", dataTypeNames[i], parents[i])
			}
			parentSections[i] = parentSection
			refParentID := strings.TrimSpace(matter.StripReferenceSuffixes(spec.ReferenceName(parentSection.Base)))
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

func (b *Ball) canonicalizeDataTypeSectionName(dp *docParse, s *spec.Section, dataTypeName string) {
	if b.errata.IgnoreSection(s.Name, errata.DiscoPurposeDataTypeRename) {
		return
	}
	name := s.Name
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
	setSectionTitle(s, newName)
	oldName := text.TrimCaseInsensitiveSuffix(name, " type")
	newName = text.TrimCaseInsensitiveSuffix(newName, " type")
	if oldName == newName {
		return
	}
	renameDataType(dp.attributes, oldName, newName)
	renameDataType(dp.commands, oldName, newName)
	renameDataType(dp.events, oldName, newName)
}

func renameDataType(subSections []*subSection, oldName string, newName string) {
	for _, ss := range subSections {
		renameDataType(ss.children, oldName, newName)
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
			vc, e := spec.RenderTableCell(typeCell)
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

func (b *Ball) removeMandatoryFallbacks(ti *spec.TableInfo) {
	if !b.options.removeMandatoryFallbacks {
		return
	}
	fallbackIndex, hasFallback := ti.ColumnMap[matter.TableColumnFallback]
	_, hasConformance := ti.ColumnMap[matter.TableColumnConformance]
	if !hasFallback || !hasConformance {
		return
	}
	for row := range ti.Body() {
		conf := ti.ReadConformance(row, matter.TableColumnConformance)
		if !conformance.IsMandatory(conf) {
			continue
		}
		fallbackCell := row.Cell(fallbackIndex)
		def, err := ti.Doc.GetHeaderCellString(fallbackCell)
		if err != nil {
			continue
		}
		def = strings.TrimSpace(def)
		if def != "" {
			setCellString(fallbackCell, "")
		}
	}
}
