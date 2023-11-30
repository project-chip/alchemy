package zap

import "github.com/hasty/alchemy/matter"

var defaultOrder = []matter.Section{matter.SectionAttributes, matter.SectionCommands, matter.SectionEvents}

type Errata struct {
	TopOrder      []matter.Section
	ClusterOrder  []matter.Section
	DataTypeOrder []matter.Section

	SuppressAttributePermissions bool
	ClusterDefinePrefix          string
	SuppressClusterDefinePrefix  bool
	DefineOverrides              map[string]string

	WriteRoleAsPrivilege bool
	SeparateStructs      map[string]struct{}
}

var DefaultErrata = &Errata{
	TopOrder:      []matter.Section{matter.SectionFeatures, matter.SectionCluster, matter.SectionDataTypes},
	ClusterOrder:  []matter.Section{matter.SectionAttributes, matter.SectionCommands, matter.SectionEvents},
	DataTypeOrder: []matter.Section{matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct},
}

var Erratas = map[string]*Errata{
	"AirQuality.adoc": {
		TopOrder:                    []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
		DataTypeOrder:               []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
		SuppressClusterDefinePrefix: true,
	},
	"ApplicationBasic.adoc": {
		ClusterDefinePrefix: "APPLICATION_",

		DefineOverrides: map[string]string{"APPLICATION_APPLICATION": "APPLICATION_APP"},
	},
	"ApplicationLauncher.adoc": {
		ClusterDefinePrefix: "APPLICATION_LAUNCHER_",
		DefineOverrides: map[string]string{
			"APPLICATION_LAUNCHER_CATALOG_LIST": "APPLICATION_LAUNCHER_LIST",
		},
		TopOrder:      []matter.Section{matter.SectionCluster, matter.SectionDataTypes, matter.SectionFeatures},
		DataTypeOrder: []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
	},
	"AudioOutput.adoc": {
		DataTypeOrder:       []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
		ClusterDefinePrefix: "AUDIO_OUTPUT_",
		DefineOverrides: map[string]string{
			"AUDIO_OUTPUT_OUTPUT_LIST": "AUDIO_OUTPUT_LIST",
		},
	},
	"BallastConfiguration.adoc": {
		SuppressClusterDefinePrefix: true,
	},
	"BooleanState.adoc": {
		SuppressClusterDefinePrefix: true,
	},
	"Channel.adoc": {
		ClusterDefinePrefix: "CHANNEL_",
	},
	"ColorControl.adoc": {
		ClusterDefinePrefix: "COLOR_CONTROL_",
	},
	"ConcentrationMeasurement.adoc": {
		TopOrder:                    []matter.Section{matter.SectionCluster, matter.SectionFeatures, matter.SectionDataTypes},
		SuppressClusterDefinePrefix: true,
	},
	"ContentLauncher.adoc": {
		ClusterDefinePrefix: "CONTENT_LAUNCHER_",
	},
	"DemandResponseLoadControl.adoc": {
		TopOrder:             []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		ClusterOrder:         DefaultErrata.ClusterOrder,
		DataTypeOrder:        DefaultErrata.DataTypeOrder,
		WriteRoleAsPrivilege: true,
		DefineOverrides: map[string]string{
			"EVENTS":        "LOAD_CONTROL_EVENTS",
			"ACTIVE_EVENTS": "LOAD_CONTROL_ACTIVE_EVENTS",
		},
	},
	"DiagnosticsThread.adoc": {
		TopOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
		SuppressClusterDefinePrefix: true,
	},
	"DoorLock.adoc": {
		DefineOverrides: map[string]string{
			"NUMBER_OF_TOTAL_USERS_SUPPORTED":                 "NUM_TOTAL_USERS_SUPPORTED",
			"NUMBER_OF_PIN_USERS_SUPPORTED":                   "NUM_PIN_USERS_SUPPORTED",
			"NUMBER_OF_RFID_USERS_SUPPORTED":                  "NUM_RFID_USERS_SUPPORTED",
			"NUMBER_OF_WEEK_DAY_SCHEDULES_SUPPORTED_PER_USER": "NUM_WEEKDAY_SCHEDULES_SUPPORTED_PER_USER",
			"NUMBER_OF_YEAR_DAY_SCHEDULES_SUPPORTED_PER_USER": "NUM_YEARDAY_SCHEDULES_SUPPORTED_PER_USER",
			"NUMBER_OF_HOLIDAY_SCHEDULES_SUPPORTED":           "NUM_HOLIDAY_SCHEDULES_SUPPORTED",
			"MAX_PIN_CODE_LENGTH":                             "MAX_PIN_LENGTH",
			"MIN_PIN_CODE_LENGTH":                             "MIN_PIN_LENGTH",
			"NUMBER_OF_CREDENTIALS_SUPPORTED_PER_USER":        "NUM_CREDENTIALS_SUPPORTED_PER_USER",
			"REQUIRE_PI_NFOR_REMOTE_OPERATION":                "REQUIRE_PIN_FOR_REMOTE_OPERATION",
		},
	},
	"EVSE.adoc": {
		TopOrder:                    []matter.Section{matter.SectionDataTypes, matter.SectionCluster, matter.SectionFeatures},
		DataTypeOrder:               []matter.Section{matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap, matter.SectionDataTypeStruct},
		SuppressClusterDefinePrefix: true,
	},
	"FanControl.adoc": {
		TopOrder:                     []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		DataTypeOrder:                []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
		SuppressAttributePermissions: true,
		SuppressClusterDefinePrefix:  true,
	},
	"FlowMeasurement.adoc": {
		ClusterDefinePrefix: "FLOW_",
	},
	"Groups.adoc": {
		TopOrder:            []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		DataTypeOrder:       []matter.Section{matter.SectionDataTypeStruct, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap},
		ClusterDefinePrefix: "GROUP_",
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
		ClusterDefinePrefix: "MEDIA_INPUT_",
		DefineOverrides:     map[string]string{"MEDIA_INPUT_INPUT_LIST": "MEDIA_INPUT_LIST"},
	},
	"MediaPlayback.adoc": {
		ClusterDefinePrefix: "MEDIA_PLAYBACK_",
		DefineOverrides: map[string]string{
			"MEDIA_PLAYBACK_CURRENT_STATE":    "MEDIA_PLAYBACK_STATE",
			"MEDIA_PLAYBACK_SAMPLED_POSITION": "MEDIA_PLAYBACK_PLAYBACK_POSITION",
			"MEDIA_PLAYBACK_SEEK_RANGE_END":   "MEDIA_PLAYBACK_PLAYBACK_SEEK_RANGE_END",
			"MEDIA_PLAYBACK_SEEK_RANGE_START": "MEDIA_PLAYBACK_PLAYBACK_SEEK_RANGE_START",
		},
	},
	"MicrowaveOvenControl.adoc": {
		SuppressClusterDefinePrefix: true,
	},
	"ModeSelect.adoc": {
		DefineOverrides: map[string]string{"DESCRIPTION": "MODE_DESCRIPTION"},
	},
	"PressureMeasurement.adoc": {
		ClusterDefinePrefix: "PRESSURE_",
	},
	"ResourceMonitoring.adoc": {
		SeparateStructs: map[string]struct{}{"ReplacementProductStruct": {}},
	},
	"SmokeCOAlarm.adoc": {
		DefineOverrides: map[string]string{
			"HARDWARE_FAULT_ALERT":    "HARDWARE_FAULTALERT",
			"END_OF_SERVICE_ALERT":    "END_OF_SERVICEALERT",
			"SMOKE_SENSITIVITY_LEVEL": "SENSITIVITY_LEVEL",
		},
	},
	"TargetNavigator.adoc": {
		ClusterDefinePrefix: "TARGET_NAVIGATOR_",
		DefineOverrides: map[string]string{
			"TARGET_NAVIGATOR_TARGET_LIST": "TARGET_NAVIGATOR_LIST",
		},
	},
	"TemperatureControl.adoc": {
		DefineOverrides: map[string]string{
			"TEMPERATURE_SETPOINT":         "TEMP_SETPOINT",
			"MIN_TEMPERATURE":              "MIN_TEMP",
			"MAX_TEMPERATURE":              "MAX_TEMP",
			"SELECTED_TEMPERATURE_LEVEL":   "SELECTED_TEMP_LEVEL",
			"SUPPORTED_TEMPERATURE_LEVELS": "SUPPORTED_TEMP_LEVELS",
		},
	},
	"TemperatureMeasurement.adoc": {
		ClusterDefinePrefix: "TEMP_",
	},
	"Thermostat.adoc": {
		TopOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		SuppressClusterDefinePrefix: true,
		WriteRoleAsPrivilege:        true,
		DefineOverrides: map[string]string{
			"OCCUPANCY": "THERMOSTAT_OCCUPANCY",
		},
	},
	"ThermostatUserInterfaceConfiguration.adoc": {
		SuppressClusterDefinePrefix: true,
	},
	"WakeOnLAN.adoc": {
		DefineOverrides: map[string]string{
			"MAC_ADDRESS": "WAKE_ON_LAN_MAC_ADDRESS",
		},
	},
	"WaterControls.adoc": {
		TopOrder:                    []matter.Section{matter.SectionFeatures, matter.SectionDataTypes, matter.SectionCluster},
		SuppressClusterDefinePrefix: true,
	},
	"WindowCovering.adoc": {
		ClusterDefinePrefix: "WC_",
		DefineOverrides: map[string]string{
			"WC_TARGET_POSITION_LIFT_PERCENT_100_THS":  "WC_TARGET_POSITION_LIFT_PERCENT100THS",
			"WC_TARGET_POSITION_TILT_PERCENT_100_THS":  "WC_TARGET_POSITION_TILT_PERCENT100THS",
			"WC_CURRENT_POSITION_LIFT_PERCENT_100_THS": "WC_CURRENT_POSITION_LIFT_PERCENT100THS",
			"WC_CURRENT_POSITION_TILT_PERCENT_100_THS": "WC_CURRENT_POSITION_TILT_PERCENT100THS",
		},
	},
}
