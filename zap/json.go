package zap

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter"
)

type File struct {
	FileFormat    int            `json:"fileFormat"`
	FeatureLevel  int            `json:"featureLevel"`
	Creator       string         `json:"creator"`
	KeyValuePairs []KeyValuePair `json:"keyValuePairs"`
	Package       []Package      `json:"package"`
	EndpointTypes []EndpointType `json:"endpointTypes"`
	Endpoints     []Endpoint     `json:"endpoints"`
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Package struct {
	PathRelativity string       `json:"pathRelativity"`
	Path           string       `json:"path"`
	Type           string       `json:"type"`
	Category       string       `json:"category"`
	Version        NumberString `json:"version"`
	Description    string       `json:"description"`
}

type EndpointType struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	DeviceTypeRef     DeviceTypeRef   `json:"deviceTypeRef"`
	DeviceTypes       []DeviceTypeRef `json:"deviceTypes"`
	DeviceVersions    []int           `json:"deviceVersions"`
	DeviceIdentifiers []int           `json:"deviceIdentifiers"`

	DeviceTypeCode      int    `json:"deviceTypeCode"`
	DeviceTypeProfileId int    `json:"deviceTypeProfileId"`
	DeviceTypeName      string `json:"deviceTypeName"`

	Clusters []ClusterRef `json:"clusters"`

	DeviceType *matter.DeviceType `json:"-"`
}

type DeviceTypeRef struct {
	Code            int    `json:"code"`
	ProfileId       int    `json:"profileId"`
	Label           string `json:"label"`
	Name            string `json:"name"`
	DeviceTypeOrder int    `json:"deviceTypeOrder"`
}

type ClusterRef struct {
	Name        string `json:"name"`
	Code        int    `json:"code"`
	MfgCode     *int   `json:"mfgCode"`
	Define      string `json:"define"`
	Side        string `json:"side"`
	ApiMaturity string `json:"apiMaturity"`
	Enabled     int    `json:"enabled"`

	Cluster *matter.Cluster `json:"-"`

	Attributes []AttributeRef `json:"attributes"`
	Commands   []CommandRef   `json:"commands"`
	Events     []EventRef     `json:"events"`
}

type CommandRef struct {
	Name       string `json:"name"`
	Code       int    `json:"code"`
	MfgCode    *int   `json:"mfgCode"`
	Source     string `json:"source"`
	IsIncoming int    `json:"isIncoming"`
	IsEnabled  int    `json:"isEnabled"`

	Command *matter.Command `json:"-"`
}

type AttributeRef struct {
	Name             string `json:"name"`
	Code             int    `json:"code"`
	MfgCode          *int   `json:"mfgCode"`
	Side             string `json:"side"`
	Type             string `json:"type"`
	Included         int    `json:"included"`
	StorageOption    string `json:"storageOption"`
	Singleton        int    `json:"singleton"`
	Bounded          int    `json:"bounded"`
	DefaultValue     string `json:"defaultValue"`
	Reportable       int    `json:"reportable"`
	MinInterval      int    `json:"minInterval"`
	MaxInterval      int    `json:"maxInterval"`
	ReportableChange int    `json:"reportableChange"`
}

type EventRef struct {
	Name     string `json:"name"`
	Code     int    `json:"code"`
	MfgCode  *int   `json:"mfgCode"`
	Side     string `json:"side"`
	Included int    `json:"included"`
}

type Endpoint struct {
	EndpointTypeName         string `json:"endpointTypeName"`
	EndpointTypeIndex        int    `json:"endpointTypeIndex"`
	ParentEndpointIdentifier *int   `json:"parentEndpointIdentifier"`
	ProfileId                int    `json:"profileId"`
	EndpointId               int    `json:"endpointId"`
	NetworkId                int    `json:"networkId"`
}

type NumberString struct {
	Number *int
	String *string
}

func (f *NumberString) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '"' {
		return json.Unmarshal(b, &f.String)

	}
	return json.Unmarshal(b, &f.Number)
}
