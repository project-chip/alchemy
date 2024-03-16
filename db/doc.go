package db

import (
	"context"
	"path/filepath"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
)

func (h *Host) indexDoc(ctx context.Context, doc *ascii.Doc, raw bool) (*sectionInfo, error) {
	ds := &sectionInfo{id: h.nextId(documentTable), values: &dbRow{}, children: make(map[string][]*sectionInfo)}
	dt, _ := doc.DocType()
	dts := matter.DocTypeNames[dt]
	ds.values.values = map[matter.TableColumn]interface{}{matter.TableColumnName: filepath.Base(doc.Path), matter.TableColumnType: dts}
	ds.values.extras = map[string]interface{}{"path": doc.Path}
	if raw {
		for _, top := range parse.Skim[*ascii.Section](doc.Elements) {
			err := ascii.AssignSectionTypes(doc, top)
			if err != nil {
				return nil, err
			}
			for _, s := range parse.Skim[*ascii.Section](top.Elements) {
				var err error
				switch s.SecType {
				case matter.SectionClusterID:
					err = h.indexCluster(ctx, doc, ds, top)
				}
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		entities, err := doc.Entities()
		if err != nil {
			return nil, err
		}
		for _, m := range entities {
			switch v := m.(type) {
			case *matter.Cluster:
				err = h.indexClusterModel(ctx, ds, v)
			case *matter.DeviceType:
				err = h.indexDeviceTypeModel(ctx, ds, v)
			}
			if err != nil {
				return nil, err
			}
		}

	}
	return ds, nil
}
