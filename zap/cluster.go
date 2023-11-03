package zap

import (
	"encoding/xml"
	"fmt"

	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

type XMLCluster struct {
	XMLName          xml.Name  `xml:"cluster"`
	Domain           string    `xml:"domain"`
	Name             string    `xml:"name"`
	Code             string    `xml:"code"`
	Define           string    `xml:"define"`
	Description      string    `xml:"description"`
	Client           XMLClient `xml:"client"`
	Server           XMLServer `xml:"server"`
	IntroducedIn     string    `xml:"introducedIn,attr"`
	Singleton        bool      `xml:"singleton,attr"`
	ManufacturerCode string    `xml:"manufacturerCode,attr"`

	GlobalAttributes    []XMLGlobalAttribute `xml:"globalAttribute"`
	GenerateCmdHandlers []bool               `xml:"generateCmdHandlers`
	Tags                []XMLTag             `xml:"tag"`

	Attributes []XMLAttribute `xml:"attribute"`
	Commands   []XMLCommand   `xml:"command"`
	Events     []XMLEvent     `xml:"event"`
}

type XMLClusterExtension struct {
	XMLName      xml.Name `xml:"clusterExtension"`
	Code         string   `xml:"code"`
	IntroducedIn string   `xml:"introducedIn,attr"`

	Attributes []XMLAttribute `xml:"attribute"`
	Commands   []XMLCommand   `xml:"command"`
}

func (c *XMLCluster) ToModel() (mc *matter.Cluster, err error) {
	mc = &matter.Cluster{
		Name: c.Name,
	}
	var code uint64
	code, err = parse.HexOrDec(c.Code)
	if err == nil {
		mc.ID = fmt.Sprintf("%#04x", code)
	} else {
		mc.ID = c.Code
	}

	return
}
