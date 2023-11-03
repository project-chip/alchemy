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
	"ConcentrationMeasurement.adoc": {
		topOrder:                    []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
	},
	"FlowMeasurement.adoc": {
		topOrder:            defaultErrata.topOrder,
		clusterOrder:        defaultErrata.clusterOrder,
		dataTypeOrder:       defaultErrata.dataTypeOrder,
		clusterDefinePrefix: "FLOW_",
	},
	"IlluminanceMeasurement.adoc": {
		topOrder:            defaultErrata.topOrder,
		clusterOrder:        defaultErrata.clusterOrder,
		dataTypeOrder:       defaultErrata.dataTypeOrder,
		clusterDefinePrefix: "ILLUM_",
	},
	"KeypadInput.adoc": {
		topOrder:            []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:        defaultErrata.clusterOrder,
		dataTypeOrder:       defaultErrata.dataTypeOrder,
		clusterDefinePrefix: "ILLUM_",
	},
	"LevelControl.adoc": {
		topOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		defineOverrides:             map[string]string{"REMAINING_TIME": "LEVEL_CONTROL_REMAINING_TIME"},
		suppressClusterDefinePrefix: true,
	},
	"MediaInput.adoc": {
		topOrder:        defaultErrata.topOrder,
		clusterOrder:    defaultErrata.clusterOrder,
		dataTypeOrder:   defaultErrata.dataTypeOrder,
		defineOverrides: map[string]string{"MEDIA_INPUT_INPUT_LIST": "MEDIA_INPUT_LIST"},
	},
	"MicrowaveOvenControl.adoc": {
		topOrder:                    defaultErrata.topOrder,
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
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
	"WaterControls.adoc": {
		topOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
	},
	"DiagnosticsThread.adoc": {
		topOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
	},
	"ThermostatUserInterfaceConfiguration.adoc": {
		topOrder:                    defaultErrata.topOrder,
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               defaultErrata.dataTypeOrder,
		suppressClusterDefinePrefix: true,
	},
	"EVSE.adoc": {
		topOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
		clusterOrder:                defaultErrata.clusterOrder,
		dataTypeOrder:               []matter.Section{matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap, matter.SectionDataTypeStruct},
		suppressClusterDefinePrefix: true,
	},
}
