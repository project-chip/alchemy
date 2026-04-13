package db

import (
	"context"
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (h *Host) indexDoc(ctx context.Context, spec *spec.Specification, doc *asciidoc.Document) (*sectionInfo, error) {
	library, ok := spec.LibraryForDocument(doc)
	if !ok {
		return nil, fmt.Errorf("unable to find library for document %s", doc.Path.Relative)
	}
	ds := h.newSectionInfo(documentTable, nil, &dbRow{}, nil)
	dt, _ := library.DocType(doc)
	dts := matter.DocTypeNames[dt]
	ds.values.values = map[matter.TableColumn]any{matter.TableColumnName: doc.Path.Base(), matter.TableColumnType: dts}
	ds.values.extras = map[string]any{"path": doc.Path.Absolute}

	entities := library.Spec.EntitiesForDocument(doc)
	for _, m := range entities {
		var err error
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
		case *matter.Namespace:
			err = h.indexNamepsace(ctx, ds, v)
		}
		if err != nil {
			return nil, err
		}
	}

	return ds, nil
}
