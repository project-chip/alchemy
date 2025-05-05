package spec

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toAttributes(spec *Specification, d *Doc, cluster *matter.Cluster, pc *parseContext) (attributes matter.FieldSet, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		}
		return
	}
	var attributeMap map[string]*matter.Field
	attributes, attributeMap, err = d.readFields(spec, ti, types.EntityTypeAttribute, cluster)
	if err != nil {
		return
	}
	err = s.mapFields(attributeMap, pc)
	return
}
