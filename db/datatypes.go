package db

import (
	"context"
	"strings"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

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
