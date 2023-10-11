package db

import (
	"context"
	"path/filepath"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

/*type docState struct {
	doc *ascii.Doc

	clusters []*sectionInfo

	id int32
}*/

func (h *Host) indexDoc(ctx context.Context, doc *ascii.Doc) (*sectionInfo, error) {
	ds := &sectionInfo{id: h.nextId(documentTable), values: &dbRow{}, children: make(map[string][]*sectionInfo)}
	dt, _ := doc.DocType()
	ds.values.values = map[matter.TableColumn]interface{}{matter.TableColumnName: filepath.Base(doc.Path), matter.TableColumnType: dt}
	ds.values.extras = map[string]interface{}{"path": doc.Path}
	for _, top := range ascii.Skim[*ascii.Section](doc.Elements) {
		parse.AssignSectionTypes(top)
		for _, s := range ascii.Skim[*ascii.Section](top.Elements) {
			var err error
			switch s.SecType {
			case matter.SectionClusterID:
				err = h.indexCluster(ctx, ds, top)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return ds, nil
}
