package zap

import (
	"encoding/xml"
	"fmt"

	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

type XMLTag struct {
	XMLName     xml.Name `xml:"tag"`
	Name        string   `xml:"name,attr"`
	Description string   `xml:"description,attr"`
}

type XMLDomain struct {
	XMLName     xml.Name `xml:"domain"`
	Name        string   `xml:"name,attr"`
	DependsOn   string   `xml:"dependsOn,attr"`
	Spec        string   `xml:"spec,attr"`
	Certifiable bool     `xml:"certifiable,attr"`
}

type XMLClusterCode struct {
	XMLName xml.Name `xml:"cluster"`
	Code    string   `xml:"code,attr"`
}

type XMLClient struct {
	XMLName xml.Name `xml:"client"`
	Init    bool     `xml:"init,attr"`
	Tick    bool     `xml:"tick,attr"`
	Value   bool     `xml:",chardata"`
}

type XMLServer struct {
	XMLName       xml.Name `xml:"server"`
	Init          bool     `xml:"init,attr"`
	Tick          bool     `xml:"tick,attr"`
	TickFrequency string   `xml:"tickFrequency,attr"`
	Value         bool     `xml:",chardata"`
}

type XMLAccessControl struct {
	XMLName   xml.Name `xml:"accessControl"`
	Operation string   `xml:"operation,attr"`
	Role      string   `xml:"role,attr"`
	Modifier  string   `xml:"modifier,attr"`
}

type XMLDefaultAccess struct {
	XMLName xml.Name `xml:"defaultAccess"`
	Type    string   `xml:"type,attr"`

	Access []XMLAccess `xml:"access"`
}

type XMLFeatureBit struct {
	XMLName xml.Name `xml:"featureBit"`
	Tag     string   `xml:"tag,attr"`
	Bit     int      `xml:"bit,attr"`
}

type XMLGlobalAttribute struct {
	XMLName xml.Name `xml:"globalAttribute"`
	Name    string   `xml:",chardata"`
	Side    string   `xml:"side,attr"`
	Code    string   `xml:"code,attr"`
	Value   string   `xml:"value,attr"`

	FeatureBits []XMLFeatureBit `xml:"featureBit"`
}

type XMLCommandArg struct {
	XMLName      xml.Name `xml:"arg"`
	FieldID      string   `xml:"fieldId,attr"`
	Name         string   `xml:"name,attr"`
	Type         string   `xml:"type,attr"`
	Length       int      `xml:"length,attr"`
	IsNullable   bool     `xml:"isNullable,attr"`
	IsArray      bool     `xml:"array,attr"`
	ArrayLength  bool     `xml:"arrayLength,attr"`
	Default      string   `xml:"default,attr"`
	Description  string   `xml:"description,attr"`
	IntroducedIn string   `xml:"introducedIn,attr"`
	RemovedIn    string   `xml:"removedIn,attr"`
	PresentIf    string   `xml:"presentIf,attr"`
	Optional     bool     `xml:"optional,attr"`
	CountArg     string   `xml:"countArg,attr"`
}

type XMLCommand struct {
	XMLName                 xml.Name `xml:"command"`
	Name                    string   `xml:"name,attr"`
	Optional                bool     `xml:"optional,attr"`
	Cli                     string   `xml:"cli,attr"`
	CliFunctionName         string   `xml:"cliFunctionName,attr"`
	Code                    string   `xml:"code"`
	DisableDefaultResponse  bool     `xml:"disableDefaultResponse,attr"`
	FunctionName            string   `xml:"functionName,attr"`
	Group                   string   `xml:"group,attr"`
	IntroducedIn            string   `xml:"introducedIn,attr"`
	NoDefaultImplementation bool     `xml:"noDefaultImplementation,attr"`
	ManufacturerCode        string   `xml:"manufacturerCode,attr"`
	Source                  string   `xml:"source,attr"`
	Restriction             string   `xml:"restriction,attr"`
	Response                string   `xml:"response,attr"`

	Access []XMLAccess     `xml:"access"`
	Args   []XMLCommandArg `xml:"arg"`
}

type XMLEventField struct {
	XMLName    xml.Name `xml:"field"`
	ID         int      `xml:"id,attr"`
	Name       string   `xml:"name,attr"`
	Type       string   `xml:"type,attr"`
	IsNullable bool     `xml:"isNullable,attr"`
	IsArray    bool     `xml:"array,attr"`
}

type XMLEvent struct {
	XMLName  xml.Name `xml:"event"`
	Name     string   `xml:"name,attr"`
	Side     string   `xml:"side,attr"`
	Code     string   `xml:"code,attr"`
	Priority string   `xml:"priority,attr"`

	Access []XMLAccess     `xml:"access"`
	Fields []XMLEventField `xml:"field"`
}

type XMLGlobal struct {
	XMLName xml.Name `xml:"global"`

	Attributes []XMLAttribute `xml:"attribute"`
	Commands   []XMLCommand   `xml:"command"`
}

type XMLConfigurator struct {
	XMLName           xml.Name              `xml:"configurator"`
	Bitmaps           []XMLBitmap           `xml:"bitmap"`
	Structs           []XMLStruct           `xml:"struct"`
	Enums             []XMLEnum             `xml:"enum"`
	Clusters          []XMLCluster          `xml:"cluster"`
	Tags              []XMLTag              `xml:"tag"`
	Domains           []XMLDomain           `xml:"domain"`
	ClusterExtensions []XMLClusterExtension `xml:"clusterExtension"`
	Globals           []XMLGlobal           `xml:"global"`
	AccessControl     []XMLAccessControl    `xml:"accessControl"`
	DefaultAccess     []XMLDefaultAccess    `xml:"defaultAccess"`

	/*
			<xs:complexType>
		      <xs:sequence>
		        <xs:choice>
		          <xs:element ref="callback"/>
		          <xs:element minOccurs="0" maxOccurs="unbounded" ref="deviceType"/>
		        </xs:choice>
		        <xs:choice minOccurs="0" maxOccurs="unbounded">
		          <xs:element ref="atomic"/>
		        </xs:choice>
		      </xs:sequence>
		    </xs:complexType>
	*/
}

func (config *XMLConfigurator) ToModels() (models []interface{}, err error) {

	var bitmaps []*matter.Bitmap
	for _, b := range config.Bitmaps {
		var mb *matter.Bitmap
		mb, err = b.ToModel()
		if err != nil {
			return
		}
		bitmaps = append(bitmaps, mb)
	}
	var enums []*matter.Enum
	for _, e := range config.Enums {
		var me *matter.Enum
		me, err = e.ToModel()
		if err != nil {
			return
		}
		enums = append(enums, me)
	}
	var structs []*matter.Struct
	for _, s := range config.Structs {
		var ms *matter.Struct
		ms, err = s.ToModel()
		if err != nil {
			return
		}
		structs = append(structs, ms)
	}
	var clusters []*matter.Cluster
	for _, c := range config.Clusters {
		var mc *matter.Cluster
		mc, err = c.ToModel()
		if err != nil {
			return
		}
		for _, b := range bitmaps {
			mc.Bitmaps = append(mc.Bitmaps, b)
		}
		for _, e := range enums {
			mc.Enums = append(mc.Enums, e)
		}
		for _, s := range structs {
			mc.Structs = append(mc.Structs, s)
		}
		clusters = append(clusters, mc)
	}
	for _, c := range clusters {
		models = append(models, c)
	}
	return
}

func convertNumber(n string) string {
	num, err := parse.HexOrDec(n)
	if err == nil {
		return fmt.Sprintf("%#04x", num)
	} else {
		return n
	}
}
