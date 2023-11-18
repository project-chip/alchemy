package render

import "github.com/hasty/alchemy/matter"

var defaultOrder = []matter.Section{matter.SectionAttributes, matter.SectionCommands, matter.SectionEvents}

type Errata struct {
	topOrder      []matter.Section
	clusterOrder  []matter.Section
	dataTypeOrder []matter.Section

	SuppressAttributePermissions bool
	ClusterDefinePrefix          string
	SuppressClusterDefinePrefix  bool
	DefineOverrides              map[string]string

	WriteRoleAsPrivilege bool
	SeparateStructs      map[string]struct{}
}

var DefaultErrata = &Errata{
	topOrder:      []matter.Section{matter.SectionFeatures, matter.SectionCluster, matter.SectionDataTypes},
	clusterOrder:  []matter.Section{matter.SectionAttributes, matter.SectionCommands, matter.SectionEvents},
	dataTypeOrder: []matter.Section{matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct},
}

var Erratas = map[string]*Errata{
	"ApplicationBasic.adoc": {
		topOrder:            DefaultErrata.topOrder,
		clusterOrder:        DefaultErrata.clusterOrder,
		dataTypeOrder:       DefaultErrata.dataTypeOrder,
		ClusterDefinePrefix: "APPLICATION_",
		DefineOverrides:     map[string]string{"APPLICATION_APPLICATION": "APPLICATION_APP"},
	},
	"FanControl.adoc": {
		topOrder:                     []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                 DefaultErrata.clusterOrder,
		dataTypeOrder:                []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
		SuppressAttributePermissions: true,
		SuppressClusterDefinePrefix:  true,
	},
	"ConcentrationMeasurement.adoc": {
		topOrder:                    []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
	},
	"FlowMeasurement.adoc": {
		topOrder:            DefaultErrata.topOrder,
		clusterOrder:        DefaultErrata.clusterOrder,
		dataTypeOrder:       DefaultErrata.dataTypeOrder,
		ClusterDefinePrefix: "FLOW_",
	},
	"IlluminanceMeasurement.adoc": {
		topOrder:            DefaultErrata.topOrder,
		clusterOrder:        DefaultErrata.clusterOrder,
		dataTypeOrder:       DefaultErrata.dataTypeOrder,
		ClusterDefinePrefix: "ILLUM_",
	},
	"KeypadInput.adoc": {
		topOrder:            []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:        DefaultErrata.clusterOrder,
		dataTypeOrder:       DefaultErrata.dataTypeOrder,
		ClusterDefinePrefix: "ILLUM_",
	},
	"LevelControl.adoc": {
		topOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		DefineOverrides:             map[string]string{"REMAINING_TIME": "LEVEL_CONTROL_REMAINING_TIME"},
		SuppressClusterDefinePrefix: true,
	},
	"MediaInput.adoc": {
		topOrder:        DefaultErrata.topOrder,
		clusterOrder:    DefaultErrata.clusterOrder,
		dataTypeOrder:   DefaultErrata.dataTypeOrder,
		DefineOverrides: map[string]string{"MEDIA_INPUT_INPUT_LIST": "MEDIA_INPUT_LIST"},
	},
	"MicrowaveOvenControl.adoc": {
		topOrder:                    DefaultErrata.topOrder,
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
	},
	"DemandResponseLoadControl.adoc": {
		topOrder:             []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:         DefaultErrata.clusterOrder,
		dataTypeOrder:        DefaultErrata.dataTypeOrder,
		WriteRoleAsPrivilege: true,
	},
	"Thermostat.adoc": {
		topOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
		WriteRoleAsPrivilege:        true,
	},
	"AudioOutput.adoc": {
		topOrder:      DefaultErrata.topOrder,
		clusterOrder:  DefaultErrata.clusterOrder,
		dataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"ApplicationLauncher.adoc": {
		topOrder:      []matter.Section{matter.SectionCluster, matter.SectionDataTypes, matter.SectionFeatures},
		clusterOrder:  DefaultErrata.clusterOrder,
		dataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"AirQuality.adoc": {
		topOrder:                    []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
		SuppressClusterDefinePrefix: true,
	},
	"Groups.adoc": {
		topOrder:      []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:  DefaultErrata.clusterOrder,
		dataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"BooleanState.adoc": {
		topOrder:                    DefaultErrata.topOrder,
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
	},
	"BallastConfiguration.adoc": {
		topOrder:                    DefaultErrata.topOrder,
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
	},
	"WaterControls.adoc": {
		topOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
	},
	"DiagnosticsThread.adoc": {
		topOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
	},
	"ThermostatUserInterfaceConfiguration.adoc": {
		topOrder:                    DefaultErrata.topOrder,
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               DefaultErrata.dataTypeOrder,
		SuppressClusterDefinePrefix: true,
	},
	"EVSE.adoc": {
		topOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
		clusterOrder:                DefaultErrata.clusterOrder,
		dataTypeOrder:               []matter.Section{matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap, matter.SectionDataTypeStruct},
		SuppressClusterDefinePrefix: true,
	},
	"ResourceMonitoring.adoc": {
		topOrder:        DefaultErrata.topOrder,
		clusterOrder:    DefaultErrata.clusterOrder,
		dataTypeOrder:   DefaultErrata.dataTypeOrder,
		SeparateStructs: map[string]struct{}{"ReplacementProductStruct": {}},
	},
}
