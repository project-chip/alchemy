package zap

import (
	"encoding/xml"

	"github.com/hasty/alchemy/matter"
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

func (c *XMLCluster) ToModel() (mc *matter.Cluster, err error) {
	mc = &matter.Cluster{
		Name: c.Name,
		ID:   matter.ParseID(c.Code),
	}
	return
}

func ZAPName(name string) string {
	switch name {
	case "OnOff":
		name = "onoff"
	case "Mode_Laundry":
		name = "laundry washer mode"
	case "LaundryWasherControls":
		name = "Washer Controls"
	case "Scenes":
		return "scene"
	case "WindowCovering":
		return "window-covering"
	case "RefrigeratorAlarm":
		return "refrigerator-alarm"
	case "OperationalState_RVC":
		name = "Operational State RVC"
	case "PumpConfigurationControl":
		name = "PumpConfigurationAndControl"
	case "ContentLauncher":
		name = "Content Launch"
	case "Mode_RVCClean":
		name = "RVC Clean Mode"
	case "Mode_RVCRun":
		name = "RVC Run Mode"
	case "Mode_Dishwasher":
		name = "Dishwasher Mode"
	}
	return strcase.ToKebab(name + " Cluster")
}
