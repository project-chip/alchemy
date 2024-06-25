package db

import (
	"context"
	"path/filepath"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (h *Host) indexDoc(ctx context.Context, doc *spec.Doc, raw bool) (*sectionInfo, error) {
	ds := &sectionInfo{id: h.nextID(documentTable), values: &dbRow{}, children: make(map[string][]*sectionInfo)}
	dt, _ := doc.DocType()
	dts := matter.DocTypeNames[dt]
	ds.values.values = map[matter.TableColumn]any{matter.TableColumnName: filepath.Base(doc.Path), matter.TableColumnType: dts}
	ds.values.extras = map[string]any{"path": doc.Path}
	if raw {
		for _, top := range parse.Skim[*spec.Section](doc.Elements()) {
			err := spec.AssignSectionTypes(doc, top)
			if err != nil {
				return nil, err
			}
			for _, s := range parse.Skim[*spec.Section](top.Elements()) {
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
			case *matter.ClusterGroup:
				for _, c := range v.Clusters {
					err = h.indexClusterModel(ctx, ds, c)
					if err != nil {
						break
					}
				}
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
