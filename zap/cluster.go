package zap

import (
	"encoding/xml"
	"path/filepath"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/iancoleman/strcase"
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

func (c *XMLCluster) Cluster() (mc *matter.Cluster, err error) {
	mc = &matter.Cluster{
		Name: c.Name,
		ID:   matter.ParseNumber(c.Code),
	}
	return
}

func ZAPClusterName(path string, errata *Errata, entities []types.Entity) string {

	if errata.TemplatePath != "" {
		return errata.TemplatePath
	}

	path = filepath.Base(path)
	name := strings.TrimSuffix(path, filepath.Ext(path))

	var suffix string
	for _, m := range entities {
		switch m.(type) {
		case *matter.Cluster:
			suffix = "Cluster"
		}
	}
	if !strings.HasSuffix(name, suffix) {
		name += " " + suffix
	}
	return strcase.ToKebab(name)
}
