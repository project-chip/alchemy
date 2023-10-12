package db

import (
	"context"
	"log/slog"
	"strings"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func (h *Host) indexCluster(cxt context.Context, ds *sectionInfo, top *ascii.Section) error {
	ci := &sectionInfo{id: h.nextId(clusterTable), parent: ds, values: &dbRow{}}
	for _, s := range ascii.Skim[*ascii.Section](top.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionClusterID:
			appendSectionToRow(cxt, s, ci.values)
		case matter.SectionClassification:
			appendSectionToRow(cxt, s, ci.values)
		case matter.SectionFeatures:
			h.readTableSection(cxt, ci, s, featureTable)
		case matter.SectionDataTypes:
			h.indexDataTypes(cxt, ci, s)
		case matter.SectionEvents:
			h.indexEvents(cxt, ci, s)
		case matter.SectionCommands:
			h.indexCommands(cxt, ci, s)
		}
		if err != nil {
			return err
		}
	}
	for _, s := range ascii.Skim[*ascii.Section](top.Elements) {
		var err error
		switch s.SecType {
		case matter.SectionAttributes:
			err = h.readTableSection(cxt, ci, s, attributeTable)
		}
		if err != nil {
			return err
		}
	}
	ds.children[clusterTable] = append(ds.children[clusterTable], ci)
	return nil
}

func (h *Host) indexDataTypes(cxt context.Context, ds *sectionInfo, dts *ascii.Section) error {
	if ds.children == nil {
		ds.children = make(map[string][]*sectionInfo)
	}
	for _, s := range ascii.Skim[*ascii.Section](dts.Elements) {
		switch s.SecType {
		case matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct:
			var t string
			switch s.SecType {
			case matter.SectionDataTypeBitmap:
				t = "bitmap"
			case matter.SectionDataTypeEnum:
				t = "enum"
			case matter.SectionDataTypeStruct:
				t = "struct"
			}
			name := strings.TrimSuffix(s.Name, " Type")
			name = matter.StripDataTypeSuffixes(name)
			ci := &sectionInfo{
				id:     h.nextId(dataTypeTable),
				parent: ds,
				values: &dbRow{
					values: map[matter.TableColumn]interface{}{
						matter.TableColumnType: t,
						matter.TableColumnName: name,
					},
				},
				children: make(map[string][]*sectionInfo),
			}
			ds.children[dataTypeTable] = append(ds.children[dataTypeTable], ci)
			switch s.SecType {
			case matter.SectionDataTypeBitmap:
				h.readTableSection(cxt, ci, s, bitmapValue)
			case matter.SectionDataTypeEnum:
				h.readTableSection(cxt, ci, s, enumValue)
			case matter.SectionDataTypeStruct:
				h.readTableSection(cxt, ci, s, structField)
			}
		}
	}
	return nil
}

func (h *Host) indexEvents(cxt context.Context, ci *sectionInfo, es *ascii.Section) error {
	if ci.children == nil {
		ci.children = make(map[string][]*sectionInfo)
	}
	err := h.readTableSection(cxt, ci, es, eventTable)
	if err != nil {
		return err
	}
	events := ci.children[eventTable]
	if len(events) == 0 {
		return nil
	}
	em := make(map[string]*sectionInfo)
	for _, si := range ci.children[eventTable] {
		name, ok := si.values.values[matter.TableColumnName]
		if ok {
			if ns, ok := name.(string); ok {
				em[ns] = si
			}
		}
	}
	for _, s := range ascii.Skim[*ascii.Section](es.Elements) {
		switch s.SecType {
		case matter.SectionEvent:
			name := strings.TrimSuffix(s.Name, " Event")
			p, ok := em[name]
			if !ok {
				slog.Error("no matching event", "name", s.Name)
				continue
			}
			if p.children == nil {
				p.children = make(map[string][]*sectionInfo)
			}
			err = h.readTableSection(cxt, p, s, eventFieldTable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *Host) indexCommands(cxt context.Context, ci *sectionInfo, es *ascii.Section) error {
	if ci.children == nil {
		ci.children = make(map[string][]*sectionInfo)
	}
	err := h.readTableSection(cxt, ci, es, commandTable)
	if err != nil {
		return err
	}
	commands := ci.children[commandTable]
	if len(commands) == 0 {
		return nil
	}
	em := make(map[string]*sectionInfo)
	for _, si := range commands {
		name, ok := si.values.values[matter.TableColumnName]
		if ok {
			if ns, ok := name.(string); ok {
				em[ns] = si
			}
		}
	}
	for _, s := range ascii.Skim[*ascii.Section](es.Elements) {
		switch s.SecType {
		case matter.SectionCommand:
			name := strings.TrimSuffix(s.Name, " Command")
			p, ok := em[name]
			if !ok {
				slog.Error("no matching command", "name", s.Name)
				continue
			}
			if p.children == nil {
				p.children = make(map[string][]*sectionInfo)
			}
			err = h.readTableSection(cxt, p, s, commandFieldTable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
