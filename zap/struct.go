package zap

import (
	"encoding/xml"

	"github.com/hasty/alchemy/matter"
)

type XMLStructField struct {
	XMLName      xml.Name `xml:"field"`
	FieldID      int      `xml:"fieldId,attr"`
	Name         string   `xml:"name,attr"`
	Type         string   `xml:"type,attr"`
	IntroducedIn string   `xml:"introducedIn,attr"`

	IsFabricSensitive bool `xml:"isFabricSensitive,attr"`
	IsNullable        bool `xml:"isNullable,attr"`
	IsArray           bool `xml:"array,attr"`
	IsEnum            bool `xml:"enum,attr"`
	Length            int  `xml:"length,attr"`
}

type XMLStruct struct {
	XMLName        xml.Name         `xml:"struct"`
	Name           string           `xml:"name,attr"`
	IsFabricScoped string           `xml:"isFabricScoped,attr"`
	Cluster        XMLClusterCode   `xml:"cluster"`
	Fields         []XMLStructField `xml:"field"`
}

func (s *XMLStruct) ToModel() (ms *matter.Struct, err error) {
	ms = &matter.Struct{Name: s.Name}
	for _, sf := range s.Fields {
		f := matter.NewField()
		f.ID = matter.NewNumber(uint64(sf.FieldID))
		f.Name = sf.Name
		f.Type = matter.NewDataType(ConvertZapToDataTypeName(sf.Type), sf.IsArray)
		var q matter.Quality
		if sf.IsNullable {
			q |= matter.QualityNullable
		}
		f.Quality = q
		ms.Fields = append(ms.Fields, f)
	}
	return
}
