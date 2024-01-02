package zap

import (
	"encoding/xml"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

type XMLEnumItem struct {
	XMLName xml.Name `xml:"item"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type XMLEnum struct {
	XMLName xml.Name       `xml:"enum"`
	Name    string         `xml:"name,attr"`
	Type    string         `xml:"type,attr"`
	Cluster XMLClusterCode `xml:"cluster"`
	Items   []XMLEnumItem  `xml:"item"`
}

func (e *XMLEnum) Enum() (me *matter.Enum, err error) {
	me = &matter.Enum{Name: e.Name, Type: types.NewDataType(ConvertZapToDataTypeName(e.Type), false)}
	for _, ei := range e.Items {
		me.Values = append(me.Values, &matter.EnumValue{
			Name:  ei.Name,
			Value: matter.ParseNumber(ei.Value),
		})
	}
	return
}
