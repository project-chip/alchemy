package zcl

import "github.com/hasty/matterfmt/matter"

var defaultOrder = []matter.Section{matter.SectionAttributes, matter.SectionCommands, matter.SectionEvents}

type errata struct {
	topOrder      []matter.Section
	clusterOrder  []matter.Section
	dataTypeOrder []matter.Section

	suppressAttributePermissions bool
	clusterDefinePrefix          string
	suppressClusterDefinePrefix  bool
	defineOverrides              map[string]string
}

var defaultErrata = errata{
	topOrder:      []matter.Section{matter.SectionCluster, matter.SectionDataTypes},
	clusterOrder:  []matter.Section{matter.SectionAttributes, matter.SectionCommands, matter.SectionEvents},
	dataTypeOrder: []matter.Section{matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct},
}

var erratas = map[string]*errata{
	"ApplicationBasic.adoc": {
		topOrder:            defaultErrata.topOrder,
		clusterOrder:        defaultErrata.clusterOrder,
		dataTypeOrder:       defaultErrata.dataTypeOrder,
		clusterDefinePrefix: "APPLICATION_",
		defineOverrides:     map[string]string{"APPLICATION_APPLICATION": "APPLICATION_APP"},
	},
	"FanControl.adoc": {
		topOrder:                     []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                 defaultErrata.clusterOrder,
		dataTypeOrder:                []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
		suppressAttributePermissions: true,
	},
	"FlowMeasurement.adoc": {
		topOrder:            defaultErrata.topOrder,
		clusterOrder:        defaultErrata.clusterOrder,
		dataTypeOrder:       defaultErrata.dataTypeOrder,
		clusterDefinePrefix: "FLOW_",
	},
	"Thermostat.adoc": {
		topOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
	},
	"AudioOutput.adoc": {
		topOrder:      defaultErrata.topOrder,
		clusterOrder:  defaultErrata.clusterOrder,
		dataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"ApplicationLauncher.adoc": {
		topOrder:      []matter.Section{matter.SectionCluster, matter.SectionDataTypes, matter.SectionFeatures},
		clusterOrder:  defaultErrata.clusterOrder,
		dataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"AirQuality.adoc": {
		topOrder:      []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
		clusterOrder:  defaultErrata.clusterOrder,
		dataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"Groups.adoc": {
		topOrder:      []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:  defaultErrata.clusterOrder,
		dataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"BooleanState.adoc": {
		topOrder:                    defaultErrata.topOrder,
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
	},
	"BallastConfiguration.adoc": {
		topOrder:                    defaultErrata.topOrder,
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
	},
}
