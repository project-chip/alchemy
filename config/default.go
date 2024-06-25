package config

import "github.com/project-chip/alchemy/matter"

var Default = Settings{
	Zap: ZapSettings{
		Erratas: map[string]*ZapErrata{
			"ACL-Cluster.adoc": {
				TemplatePath: "access-control-cluster",
			},
			"AdminCommissioningCluster.adoc": {
				TemplatePath: "administrator-commissioning-cluster",
			},
			"AirQuality.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"ApplicationBasic.adoc": {
				ClusterDefinePrefix: "APPLICATION_",
				DefineOverrides:     map[string]string{"APPLICATION_APPLICATION": "APPLICATION_APP"},
			},
			"ApplicationLauncher.adoc": {
				ClusterDefinePrefix: "APPLICATION_LAUNCHER_",
				DefineOverrides: map[string]string{
					"APPLICATION_LAUNCHER_CATALOG_LIST": "APPLICATION_LAUNCHER_LIST",
				},
			},
			"AudioOutput.adoc": {
				ClusterDefinePrefix: "AUDIO_OUTPUT_",
				DefineOverrides: map[string]string{
					"AUDIO_OUTPUT_OUTPUT_LIST": "AUDIO_OUTPUT_LIST",
				},
			},
			"BallastConfiguration.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"BasicInformationCluster.adoc": {
				Domain: matter.DomainCHIP,
			},
			"BooleanState.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"bridge-clusters.adoc": {
				ClusterSplit: map[string]string{
					"0x0025": "actions-cluster",
					"0x0039": "bridged-device-basic-information",
				},
			},

			"Channel.adoc": {
				ClusterDefinePrefix: "CHANNEL_",
			},
			"ColorControl.adoc": {
				ClusterDefinePrefix: "COLOR_CONTROL_",
			},
			"ConcentrationMeasurement.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"ContentLauncher.adoc": {
				TemplatePath:        "content-launch-cluster",
				ClusterDefinePrefix: "CONTENT_LAUNCHER_",
			},
			"DemandResponseLoadControl.adoc": {
				TemplatePath: "drlc-cluster",
				DefineOverrides: map[string]string{
					"EVENTS":        "LOAD_CONTROL_EVENTS",
					"ACTIVE_EVENTS": "LOAD_CONTROL_ACTIVE_EVENTS",
				},
			},
			"DiagnosticsGeneral.adoc": {
				TemplatePath: "general-diagnostics-cluster",
			},
			"DiagnosticsEthernet.adoc": {
				TemplatePath: "ethernet-network-diagnostics-cluster",
			},
			"DiagnosticLogsCluster.adoc": {
				Domain: matter.DomainCHIP,
			},
			"DiagnosticsSoftware.adoc": {
				TemplatePath: "software-diagnostics-cluster",
			},
			"DiagnosticsThread.adoc": {
				TemplatePath:                "thread-network-diagnostics-cluster",
				SuppressClusterDefinePrefix: true,
			},
			"DiagnosticsWiFi.adoc": {
				TemplatePath: "wifi-network-diagnostics-cluster",
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
				TemplatePath:                "energy-evse-cluster",
				SuppressClusterDefinePrefix: true,
			},
			"FanControl.adoc": {
				SuppressAttributePermissions: true,
				SuppressClusterDefinePrefix:  true,
			},
			"FlowMeasurement.adoc": {
				ClusterDefinePrefix: "FLOW_",
			},
			"Group-Key-Management-Cluster.adoc": {
				TemplatePath: "group-key-mgmt-cluster",
			},
			"Groups.adoc": {
				ClusterDefinePrefix: "GROUP_",
			},
			"IlluminanceMeasurement.adoc": {
				ClusterDefinePrefix: "ILLUM_",
			},
			"KeypadInput.adoc": {
				ClusterDefinePrefix: "ILLUM_",
			},
			"Label-Cluster.adoc": {
				TemplatePath: "user-label-cluster",
			},
			"LaundryWasherControls.adoc": {
				TemplatePath: "washer-controls-cluster",
			},
			"LevelControl.adoc": {
				DefineOverrides:             map[string]string{"REMAINING_TIME": "LEVEL_CONTROL_REMAINING_TIME"},
				SuppressClusterDefinePrefix: true,
			},
			"LocalizationTimeFormat.adoc": {
				TemplatePath: "time-format-localization-cluster",
			},
			"LocalizationUnit.adoc": {
				Domain:       matter.DomainCHIP,
				TemplatePath: "unit-localization-cluster",
			},
			"meas_and_sense.adoc": {
				TemplatePath: "measurement-and-sensing",
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
			"Mode_Dishwasher.adoc": {
				TemplatePath: "dishwasher-mode-cluster",
			},
			"Mode_LaundryWasher.adoc": {
				TemplatePath: "laundry-washer-mode-cluster",
			},
			"Mode_MicrowaveOven.adoc": {
				TemplatePath: "microwave-oven-mode-cluster",
			},
			"Mode_Oven.adoc": {
				TemplatePath: "oven-mode-cluster",
			},
			"Mode_Refrigerator.adoc": {
				TemplatePath: "refrigerator-and-temperature-controlled-cabinet-mode-cluster",
			},
			"Mode_RVCClean.adoc": {
				TemplatePath: "rvc-clean-mode-cluster",
			},
			"Mode_RVCRun.adoc": {
				TemplatePath: "rvc-run-mode-cluster",
			},
			"OnOff.adoc": {
				TemplatePath: "onoff-cluster",
			},
			"OperationalCredentialCluster.adoc": {
				TemplatePath: "operational-credentials-cluster",
			},
			"OperationalState_RVC": {
				TemplatePath: "operational-state-rvc-cluster",
			},
			"PowerSourceConfigurationCluster.adoc": {
				Domain: matter.DomainCHIP,
			},
			"PowerSourceCluster.adoc": {
				Domain: matter.DomainCHIP,
			},
			"PressureMeasurement.adoc": {
				ClusterDefinePrefix: "PRESSURE_",
			},
			"PumpConfigurationControl.adoc": {
				TemplatePath: "pump-configuration-and-control-cluster",
			},
			"RefrigeratorAlarm.adoc": {
				TemplatePath: "refrigerator-alarm",
			},
			"ResourceMonitoring.adoc": {
				SeparateStructs: map[string]struct{}{"ReplacementProductStruct": {}},
			},
			"Scenes.adoc": {
				TemplatePath: "scene",
			},
			"SmokeCOAlarm.adoc": {
				DefineOverrides: map[string]string{
					"HARDWARE_FAULT_ALERT":    "HARDWARE_FAULTALERT",
					"END_OF_SERVICE_ALERT":    "END_OF_SERVICEALERT",
					"SMOKE_SENSITIVITY_LEVEL": "SENSITIVITY_LEVEL",
				},
			},
			"Switch.adoc": {
				Domain: matter.DomainCHIP, // wth?
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
				SuppressClusterDefinePrefix: true,
				DefineOverrides: map[string]string{
					"OCCUPANCY": "THERMOSTAT_OCCUPANCY",
				},
			},
			"ThermostatUserInterfaceConfiguration.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"TimeSync.adoc": {
				TemplatePath: "time-synchronization-cluster",
			},
			"ValidProxies-Cluster.adoc": {
				TemplatePath: "proxy-valid-cluster",
			},
			"ValveConfigurationControl.adoc": {
				TemplatePath: "valve-configuration-and-control-cluster",
			},
			"WakeOnLAN.adoc": {
				DefineOverrides: map[string]string{
					"MAC_ADDRESS": "WAKE_ON_LAN_MAC_ADDRESS",
				},
			},
			"WaterContentMeasurement.adoc": {
				TemplatePath: "relative-humidity-measurement-cluster",
			},
			"WaterControls.adoc": {
				SuppressClusterDefinePrefix: true,
			},
			"WindowCovering.adoc": {
				TemplatePath:        "window-covering",
				ClusterDefinePrefix: "WC_",
				DefineOverrides: map[string]string{
					"WC_TARGET_POSITION_LIFT_PERCENT_100_THS":  "WC_TARGET_POSITION_LIFT_PERCENT100THS",
					"WC_TARGET_POSITION_TILT_PERCENT_100_THS":  "WC_TARGET_POSITION_TILT_PERCENT100THS",
					"WC_CURRENT_POSITION_LIFT_PERCENT_100_THS": "WC_CURRENT_POSITION_LIFT_PERCENT100THS",
					"WC_CURRENT_POSITION_TILT_PERCENT_100_THS": "WC_CURRENT_POSITION_TILT_PERCENT100THS",
				},
			},
		}},
}
