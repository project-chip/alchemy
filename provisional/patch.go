package provisional

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func patchProvisional(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, s *spec.Specification, violations map[string][]spec.Violation, writer files.Writer[string]) (err error) {
	docs := make(map[*asciidoc.Document][]spec.Violation)
	violationsByEntity := make(map[types.Entity]spec.Violation)
	for _, vs := range violations {
		for _, v := range vs {
			if v.Entity != nil {
				if _, ok := violationsByEntity[v.Entity]; ok {
					slog.Warn("Multiple provisional violations found for entity", matter.LogEntity("entity", v.Entity))
				}
				violationsByEntity[v.Entity] = v
				if doc, ok := s.DocRefs[v.Entity]; ok {
					docs[doc] = append(docs[doc], v)
				} else {
					slog.Error("entity with missing doc", matter.LogEntity("entity", v.Entity))
				}
			} else {
				slog.Error("Violation with nil entity!")
			}
		}
	}
	for doc, vs := range docs {
		err = patchViolations(s, doc, vs)
		if err != nil {
			return
		}
	}

	renderDocs := pipeline.NewConcurrentMapPresized[string, *pipeline.Data[*asciidoc.Document]](len(docs))
	for doc := range docs {
		renderDocs.Store(doc.Path.Relative, pipeline.NewData(doc.Path.Relative, doc))
	}

	var includeDocs []*asciidoc.Document
	includeDocs, err = patchIncludes(s, violations)
	if err != nil {
		return
	}
	for _, doc := range includeDocs {
		renderDocs.Store(doc.Path.Relative, pipeline.NewData(doc.Path.Relative, doc))
	}

	renderer := render.NewRenderer()
	var renders pipeline.StringSet
	renders, err = pipeline.Parallel(cxt, pipelineOptions, renderer, renderDocs)
	if err != nil {
		return err
	}

	err = writer.Write(cxt, renders, pipelineOptions)
	return
}

type tableViolation struct {
	row       *asciidoc.TableRow
	violation spec.Violation
}

var inProgressAttributes []asciidoc.AttributeName = []asciidoc.AttributeName{"in-progress", "env-github"}

func patchViolations(s *spec.Specification, doc *asciidoc.Document, violations []spec.Violation) (err error) {
	library, ok := s.LibraryForDocument(doc)
	if !ok {
		err = fmt.Errorf("unable to find library for doc %s", doc.Path.Relative)
		return
	}
	wrappedEntities := make(map[types.Entity]struct{})
	tableViolationsByEntityType := make(map[types.EntityType]map[*asciidoc.Table][]tableViolation)
	for _, v := range violations {
		switch v.Entity.EntityType() {
		case types.EntityTypeAttribute,
			types.EntityTypeBitmapValue,
			types.EntityTypeCluster,
			types.EntityTypeEnumValue,
			types.EntityTypeEvent,
			types.EntityTypeEventField,
			types.EntityTypeCommand,
			types.EntityTypeCommandField,
			types.EntityTypeStructField:
			table, row, err := tableForViolation(library, doc, v)
			if err != nil {
				return err
			}
			if table == nil {
				slog.Info("missing table when attempting to fix provisional violation", log.Path("source", v.Entity))
				continue
			}
			if row == nil {
				slog.Info("missing table row for violation", log.Path("source", v.Entity))
				continue

			}
			violationsByTable, ok := tableViolationsByEntityType[v.Entity.EntityType()]
			if !ok {
				violationsByTable = make(map[*asciidoc.Table][]tableViolation)
				tableViolationsByEntityType[v.Entity.EntityType()] = violationsByTable
			}
			violationsByTable[table] = append(violationsByTable[table], tableViolation{row: row, violation: v})
		case types.EntityTypeEnum, types.EntityTypeBitmap, types.EntityTypeStruct:
			switch source := v.Entity.Source().(type) {
			case *asciidoc.Section:
				wrapIfDef(doc, source)
				wrappedEntities[v.Entity] = struct{}{}
			case nil:
				return fmt.Errorf("nil entity source: %s", matter.EntityName(v.Entity))
			default:
				return fmt.Errorf("unexpected entity source type %T", v.Entity.Source())
			}
		default:
			return fmt.Errorf("unknown entity type %s", v.Entity.EntityType())
		}
	}
	for entityType, tables := range tableViolationsByEntityType {
		idColumns := idColumnsForEntityType(entityType)
		if len(idColumns) == 0 {
			slog.Warn("unknown id columns for entity type", "entityType", entityType)
			continue
		}
		for table, vs := range tables {
			var ifDefIndexes []int
			children := table.Children()
			for i, child := range children {
				switch child := child.(type) {
				case *asciidoc.TableRow:
					for _, v := range vs {
						if v.row == child {
							if v.violation.Type.Has(spec.ViolationTypeNotIfDefd) {
								ifDefIndexes = append(ifDefIndexes, i)
							}
							if v.violation.Type.Has(spec.ViolationTypeNonProvisional) {
								addProvisionalConformance(library, doc, v.violation.Entity, child)
							}
						}
					}
				}
			}
			if len(ifDefIndexes) > 0 {
				ifDefdTableElements := make(asciidoc.Elements, 0, len(children)+len(ifDefIndexes))
				lastInsertedIndex := -2

				for i, child := range children {
					if len(ifDefIndexes) > 0 && i == ifDefIndexes[0] {
						if lastInsertedIndex != i-1 {
							ifDefdTableElements = append(ifDefdTableElements, asciidoc.NewIfDef(inProgressAttributes, asciidoc.ConditionalUnionAny))
						}
						lastInsertedIndex = i
						ifDefIndexes = ifDefIndexes[1:]
					} else if lastInsertedIndex >= 0 {
						if ifdef, ok := child.(*asciidoc.IfDef); ok {
							// If we were going to insert an end if, but it's actually the start of the same ifdef, then combine
							if ifdef.Attributes.Equals(inProgressAttributes) {
								lastInsertedIndex = -2
								continue
							}
							ifDefdTableElements = append(ifDefdTableElements, asciidoc.NewEndIf(nil, asciidoc.ConditionalUnionAny))
							lastInsertedIndex = -2
						}
					}
					ifDefdTableElements = append(ifDefdTableElements, child)
				}
				if lastInsertedIndex >= 0 {
					ifDefdTableElements = append(ifDefdTableElements, asciidoc.NewEndIf(nil, asciidoc.ConditionalUnionAny))
				}
				table.SetChildren(ifDefdTableElements)
			}
		}

	}
	return
}

func wrapIfDef(doc *asciidoc.Document, source *asciidoc.Section) {
	elements := doc.Children()
	var whitespaceBuffer asciidoc.Elements
	newElements := make(asciidoc.Elements, 0, len(elements)+2)
	insertEndIf := false
	for _, element := range elements {
		switch element := element.(type) {
		case *asciidoc.Section:
			if element == source {
				newElements = append(newElements, whitespaceBuffer...)
				whitespaceBuffer = nil
				newElements = append(newElements, asciidoc.NewIfDef(inProgressAttributes, asciidoc.ConditionalUnionAny))
				insertEndIf = true
			} else if insertEndIf {
				newElements = append(newElements, asciidoc.NewEndIf(nil, asciidoc.ConditionalUnionAny))
				newElements = append(newElements, whitespaceBuffer...)
				whitespaceBuffer = nil
				insertEndIf = false
			} else {
				newElements = append(newElements, whitespaceBuffer...)
				whitespaceBuffer = nil
			}
		case *asciidoc.IfDef:
			if insertEndIf {
				newElements = append(newElements, asciidoc.NewEndIf(nil, asciidoc.ConditionalUnionAny))
				insertEndIf = false
			}
			newElements = append(newElements, whitespaceBuffer...)
			whitespaceBuffer = nil
		case *asciidoc.EmptyLine:
			whitespaceBuffer = append(whitespaceBuffer, element)
			continue
		case asciidoc.Element:
			newElements = append(newElements, whitespaceBuffer...)
			whitespaceBuffer = nil
		}
		newElements = append(newElements, element)
	}
	if insertEndIf {
		newElements = append(newElements, asciidoc.NewEndIf(nil, asciidoc.ConditionalUnionAny))
	}
	newElements = append(newElements, whitespaceBuffer...)
	doc.SetChildren(newElements)
}

func addProvisionalConformance(library *spec.Library, doc *asciidoc.Document, e types.Entity, source asciidoc.Element) (err error) {
	switch source := source.(type) {
	case *asciidoc.TableRow:
		var table *spec.TableInfo
		table, err = spec.ReadTable(doc, asciidoc.RawReader, source.Parent)
		if err != nil {
			return
		}
		var conf conformance.Set
		conformanceIndex, ok := table.ColumnMap[matter.TableColumnConformance]
		if !ok {
			conformanceIndex, err = table.AppendColumn(library, matter.TableColumnConformance, e.EntityType())
			if err != nil {
				return
			}
		} else {
			conf = table.ReadConformance(asciidoc.RawReader, source, matter.TableColumnConformance)
			if conformance.IsProvisional(conf) {
				slog.Warn("Attempting to add provisionality to an entity which is already provisional", matter.LogEntity("entity", e), log.Path("source", source))
				return
			}
		}
		conf = append(conformance.Set{&conformance.Provisional{}}, conf...)
		setCellString(source.TableCells()[conformanceIndex], conf.ASCIIDocString())
	default:
		slog.Error("Unexpected provisional conformance source", matter.LogEntity("entity", e), log.Type("sourceType", source))
	}
	return
}

func setCellString(cell *asciidoc.TableCell, v string) {
	se := asciidoc.NewString(v)
	cell.SetChildren(asciidoc.Elements{se})
}

func tableForViolation(library *spec.Library, doc *asciidoc.Document, v spec.Violation) (table *asciidoc.Table, row *asciidoc.TableRow, err error) {
	source := v.Entity.Source()
	switch source := source.(type) {
	case *asciidoc.TableRow:
		table = source.Parent
		row = source
	case *asciidoc.Section:
		table, err = findTableForSection(library, doc, source, v.Entity)
		if err != nil {
			return
		}
		if table == nil {
			return
		}
		var ti *spec.TableInfo
		ti, err = spec.ReadTable(doc, asciidoc.RawReader, table)
		if err != nil {
			return nil, nil, err
		}
		row, err = findRowForEntity(ti, v.Entity)
		if row != nil || err != nil {
			return
		}
	default:
		slog.Error("Unexpected provisional table source", matter.LogEntity("entity", v.Entity), log.Type("sourceType", source))
	}
	return
}

func findRowForEntity(ti *spec.TableInfo, entity types.Entity) (row *asciidoc.TableRow, err error) {
	idColumns := idColumnsForEntityType(entity.EntityType())
	if len(idColumns) == 0 {
		return nil, fmt.Errorf("unknown ID columns for entity type %s", entity.EntityType())
	}

	id := matter.EntityID(entity)
	if id.Valid() {
		for contentRow := range ti.ContentRows() {
			var rowId *matter.Number
			rowId, err = ti.ReadID(asciidoc.RawReader, contentRow, idColumns...)
			if err != nil {
				continue
			}
			if id.Equals(rowId) {
				row = contentRow
				return
			}
		}
	} else {
		name := matter.EntityName(entity)
		if name != "" {
			nameColumns := nameColumnsForEntityType(entity.EntityType())
			if len(nameColumns) == 0 {
				return nil, fmt.Errorf("unknown name columns for entity type %s", entity.EntityType())
			}
			for contentRow := range ti.ContentRows() {
				var rowName string
				rowName, _, err = ti.ReadName(asciidoc.RawReader, contentRow, nameColumns...)
				if err != nil {
					continue
				}
				if strings.EqualFold(rowName, name) {
					row = contentRow
					return
				}
			}
		}
	}
	return
}

func findTableForSection(library *spec.Library, doc *asciidoc.Document, section *asciidoc.Section, entity types.Entity) (table *asciidoc.Table, err error) {
	neededSectionType := parentSectionTypeForEntity(entity)
	if neededSectionType != matter.SectionUnknown {
		var parent asciidoc.Element = section
		for parent != nil {
			switch parent := parent.(type) {
			case *asciidoc.Section:
				if library.SectionType(parent) == neededSectionType {
					table = spec.FindFirstTable(library, parent)
					return
				}
			}
			if p, ok := parent.(asciidoc.ChildElement); ok {
				parent = library.Parent(p)
			} else {
				break
			}
		}
	} else {
		neededSectionType = childSectionTypeForEntity(entity)
		if neededSectionType == matter.SectionUnknown {
			return nil, fmt.Errorf("unknown provisional section type for entity: %T", entity)
		}
		if library.SectionType(section) == neededSectionType {
			table = spec.FindFirstTable(library, section)
			return
		}
		parse.Search(doc, library, section, section.Elements, func(doc *asciidoc.Document, child *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
			if library.SectionType(child) == neededSectionType {
				table = spec.FindFirstTable(library, child)
				return parse.SearchShouldStop

			}
			return parse.SearchShouldContinue
		})
		return
	}

	return
}

func childSectionTypeForEntity(e types.Entity) matter.Section {
	switch e.(type) {
	case *matter.Cluster:
		return matter.SectionClusterID
	default:
		slog.Warn("unexpected entity type for child section", log.Type("entityType", e))
		return matter.SectionUnknown
	}
}

func parentSectionTypeForEntity(e types.Entity) matter.Section {
	switch e.(type) {
	case *matter.Event:
		return matter.SectionEvents
	case *matter.Command:
		return matter.SectionCommand

	default:
		slog.Debug("unexpected entity type for parent section", log.Type("entityType", e))
		return matter.SectionUnknown
	}
}

func idColumnsForEntityType(entityType types.EntityType) []matter.TableColumn {
	switch entityType {
	case types.EntityTypeAttribute:
		return matter.IDColumns.Attribute
	case types.EntityTypeEvent:
		return matter.IDColumns.Event
	case types.EntityTypeCommand:
		return matter.IDColumns.Command
	case types.EntityTypeBitmapValue:
		return matter.IDColumns.BitmapBit
	case types.EntityTypeEnumValue:
		return matter.IDColumns.EnumValue
	case types.EntityTypeCluster:
		return matter.IDColumns.Cluster
	case types.EntityTypeStructField, types.EntityTypeCommandField, types.EntityTypeEventField:
		return matter.IDColumns.Field
	default:
		return []matter.TableColumn{}
	}
}

func nameColumnsForEntityType(entityType types.EntityType) []matter.TableColumn {
	switch entityType {
	case types.EntityTypeCluster:
		return []matter.TableColumn{matter.TableColumnClusterName, matter.TableColumnName}

	default:
		return []matter.TableColumn{}
	}
}
