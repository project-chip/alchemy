package zap

import (
	"encoding/xml"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

type XMLAttribute struct {
	XMLName           xml.Name `xml:"attribute"`
	Name              string   `xml:",chardata"`
	Side              string   `xml:"side,attr"`
	Code              string   `xml:"code,attr"`
	Define            string   `xml:"define,attr"`
	Type              string   `xml:"type,attr"`
	EntryType         string   `xml:"entryType,attr"`
	Default           string   `xml:"default,attr"`
	IntroducedIn      string   `xml:"introducedIn,attr"`
	Length            string   `xml:"length,attr"`
	Max               string   `xml:"max,attr"`
	Min               string   `xml:"min,attr"`
	ReportMaxInterval string   `xml:"reportMaxInterval,attr"`
	ReportMinInterval string   `xml:"reportMinInterval,attr"`
	ReportableChange  string   `xml:"reportableChange,attr"`
	Readable          bool     `xml:"readable,attr"`
	Writable          bool     `xml:"writable,attr"`
	Reportable        bool     `xml:"reportable,attr"`
	Optional          bool     `xml:"optional,attr"`
	Description       string   `xml:"description"`
	IsFabricSensitive bool     `xml:"isFabricSensitive,attr"`
	IsArray           bool     `xml:"array,attr"`
	IsNullable        bool     `xml:"isNullable,attr"`

	Access []XMLAccess `xml:"access"`
}

func (xa *XMLAttribute) ToModel() (ma *matter.Field, err error) {
	ma = matter.NewAttribute()

	ma.ID = matter.ParseNumber(xa.Code)
	ma.Name = xa.Name
	ma.Type = matter.NewDataType(ConvertZapToDataTypeName(xa.Type), xa.IsArray)
	ma.Default = xa.Default

	if xa.IsNullable {
		ma.Quality |= matter.QualityNullable
	}
	ma.Access = ToAccessModel(xa.Access)
	if xa.IsFabricSensitive {
		ma.Access.FabricSensitive = true
	}
	if !xa.Optional {
		ma.Conformance = conformance.Set{&conformance.Mandatory{}}
	}
	return
}
