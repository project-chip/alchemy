package config

import "github.com/hasty/alchemy/matter"

var Default = Settings{
	Zap: ZapSettings{
		Erratas: map[string]*ZapErrata{
			"ApplicationBasic.adoc": {
				ClusterDefinePrefix: "APPLICATION_",
				DefineOverrides:     map[string]string{"APPLICATION_APPLICATION": "APPLICATION_APP"},
			},
			"FanControl.adoc": {
				TopOrder:                     []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
				DataTypeOrder:                []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
				SuppressAttributePermissions: true,
				SuppressClusterDefinePrefix:  true,
			},
			"ConcentrationMeasurement.adoc": {
				TopOrder:                    []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
				SuppressClusterDefinePrefix: true,
			},
			"FlowMeasurement.adoc": {
				ClusterDefinePrefix: "FLOW_",
			},
			"IlluminanceMeasurement.adoc": {
				ClusterDefinePrefix: "ILLUM_",
			},
			"KeypadInput.adoc": {
				TopOrder:            []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
				ClusterDefinePrefix: "ILLUM_",
			},
			"LevelControl.adoc": {
				TopOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
				DefineOverrides:             map[string]string{"REMAINING_TIME": "LEVEL_CONTROL_REMAINING_TIME"},
				SuppressClusterDefinePrefix: true,
			},
			"MediaInput.adoc": {
				DefineOverrides: map[string]string{"MEDIA_INPUT_INPUT_LIST": "MEDIA_INPUT_LIST"},
			},
			"MicrowaveOvenControl.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"DemandResponseLoadControl.adoc": {
				TopOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
				SuppressClusterDefinePrefix: true,
			},
			"Thermostat.adoc": {
				TopOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
				SuppressClusterDefinePrefix: true,
				WriteRoleAsPrivilege:        true,
			},
			"AudioOutput.adoc": {
				DataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
			},
			"ApplicationLauncher.adoc": {
				TopOrder:      []matter.Section{matter.SectionCluster, matter.SectionDataTypes, matter.SectionFeatures},
				DataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
			},
			"AirQuality.adoc": {
				TopOrder:                    []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
				DataTypeOrder:               []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
				SuppressClusterDefinePrefix: true,
			},
			"Groups.adoc": {
				TopOrder:      []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
				DataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
			},
			"BooleanState.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"BallastConfiguration.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"WaterControls.adoc": {
				TopOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
				SuppressClusterDefinePrefix: true,
			},
			"DiagnosticsThread.adoc": {
				TopOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
				SuppressClusterDefinePrefix: true,
			},
			"ThermostatUserInterfaceConfiguration.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"EVSE.adoc": {
				TopOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
				DataTypeOrder:               []matter.Section{matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap, matter.SectionDataTypeStruct},
				SuppressClusterDefinePrefix: true,
			},
			"ResourceMonitoring.adoc": {
				SeparateStructs: []string{"ReplacementProductStruct"},
			},
		}},
}
