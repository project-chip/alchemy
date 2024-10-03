package dm

import (
	"fmt"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func renderEnums(doc *spec.Doc, cluster *matter.Cluster, dt *etree.Element) (err error) {
	enums := make([]*matter.Enum, len(cluster.Enums))
	copy(enums, cluster.Enums)
	slices.SortStableFunc(enums, func(a, b *matter.Enum) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, e := range enums {
		en := dt.CreateElement("enum")
		en.CreateAttr("name", e.Name)
		for index, v := range e.Values {
			err = renderEnumValue(doc, cluster, en, index, v)
			if err != nil {
				return
			}
		}
	}
	return
}

func renderEnumValue(doc *spec.Doc, cluster *matter.Cluster, en *etree.Element, index int, v *matter.EnumValue) (err error) {
	var val, from, to *matter.Number
	var valFormat, fromFormat, toFormat types.NumberFormat
	if v.Value.Valid() {
		val = v.Value
		valFormat = v.Value.Format()
	} else {
		from, fromFormat, to, toFormat = matter.ParseFormattedIDRange(v.Value.Text())
		if !from.Valid() {
			return
		}
	}

	i := en.CreateElement("item")
	if val.Valid() {
		switch valFormat {
		case types.NumberFormatAuto, types.NumberFormatInt:
			i.CreateAttr("value", val.IntString())
		case types.NumberFormatHex:
			i.CreateAttr("value", val.ShortHexString())
		}
	} else {
		switch fromFormat {
		case types.NumberFormatAuto, types.NumberFormatInt:
			i.CreateAttr("from", from.IntString())
		case types.NumberFormatHex:
			i.CreateAttr("from", from.ShortHexString())
		}
		switch toFormat {
		case types.NumberFormatAuto, types.NumberFormatInt:
			i.CreateAttr("to", to.IntString())
		case types.NumberFormatHex:
			i.CreateAttr("to", to.ShortHexString())
		}
	}
	name := v.Name
	if len(name) == 0 {
		name = fmt.Sprintf("Item%d", index)
	}
	i.CreateAttr("name", name)
	if len(v.Summary) > 0 {
		i.CreateAttr("summary", scrubDescription(v.Summary))
	}
	err = RenderConformanceElement(doc, cluster, v.Conformance, i)
	return
}
