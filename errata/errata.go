package errata

import "github.com/project-chip/alchemy/matter"

type Errata struct {
	Spec Spec
	ZAP  ZAP
}

var DefaultErrata = &Errata{}

var Erratas = map[string]*Errata{
	"ACL-Cluster.adoc": {
		ZAP: ZAP{TemplatePath: "access-control-cluster"},
	},
	"AdminCommissioningCluster.adoc": {
		ZAP: ZAP{TemplatePath: "administrator-commissioning-cluster"},
	},
	"AirQuality.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"ApplicationBasic.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "APPLICATION_",
			DefineOverrides: map[string]string{"APPLICATION_APPLICATION": "APPLICATION_APP"}},
	},
	"ApplicationLauncher.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "APPLICATION_LAUNCHER_",
			DefineOverrides: map[string]string{
				"APPLICATION_LAUNCHER_CATALOG_LIST": "APPLICATION_LAUNCHER_LIST",
			}},
	},
	"AudioOutput.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "AUDIO_OUTPUT_",
			DefineOverrides: map[string]string{
				"AUDIO_OUTPUT_OUTPUT_LIST": "AUDIO_OUTPUT_LIST",
			}},
	},
	"BallastConfiguration.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"BasicInformationCluster.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"BooleanState.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"bridge-clusters.adoc": {
		ZAP: ZAP{ClusterSplit: map[string]string{
			"0x0025": "actions-cluster",
			"0x0039": "bridged-device-basic-information",
		}},
	},

	"Channel.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "CHANNEL_"},
	},
	"ColorControl.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "COLOR_CONTROL_"},
	},
	"ConcentrationMeasurement.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"ContentLauncher.adoc": {
		ZAP: ZAP{TemplatePath: "content-launch-cluster",
			ClusterDefinePrefix: "CONTENT_LAUNCHER_"},
	},
	"DemandResponseLoadControl.adoc": {
		ZAP: ZAP{TemplatePath: "drlc-cluster",
			DefineOverrides: map[string]string{
				"EVENTS":        "LOAD_CONTROL_EVENTS",
				"ACTIVE_EVENTS": "LOAD_CONTROL_ACTIVE_EVENTS",
			}},
	},
	"DiagnosticsGeneral.adoc": {
		ZAP: ZAP{TemplatePath: "general-diagnostics-cluster"},
	},
	"DiagnosticsEthernet.adoc": {
		ZAP: ZAP{TemplatePath: "ethernet-network-diagnostics-cluster"},
	},
	"DiagnosticLogsCluster.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"DiagnosticsSoftware.adoc": {
		ZAP: ZAP{TemplatePath: "software-diagnostics-cluster"},
	},
	"DiagnosticsThread.adoc": {
		ZAP: ZAP{TemplatePath: "thread-network-diagnostics-cluster",
			SuppressClusterDefinePrefix: true},
	},
	"DiagnosticsWiFi.adoc": {
		ZAP: ZAP{TemplatePath: "wifi-network-diagnostics-cluster"},
	},
	"DoorLock.adoc": {
		ZAP: ZAP{DefineOverrides: map[string]string{
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
		}},
	},
	"EVSE.adoc": {
		ZAP: ZAP{TemplatePath: "energy-evse-cluster",
			SuppressClusterDefinePrefix: true},
	},
	"FanControl.adoc": {
		ZAP: ZAP{SuppressAttributePermissions: true,
			SuppressClusterDefinePrefix: true},
	},
	"FlowMeasurement.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "FLOW_"},
	},
	"Group-Key-Management-Cluster.adoc": {
		ZAP: ZAP{TemplatePath: "group-key-mgmt-cluster"},
	},
	"Groups.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "GROUP_"},
	},
	"IlluminanceMeasurement.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "ILLUM_"},
	},
	"KeypadInput.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "ILLUM_"},
	},
	"Label-Cluster.adoc": {
		ZAP: ZAP{TemplatePath: "user-label-cluster",
			ClusterSplit: map[string]string{
				"0x0040": "fixed-label-cluster",
				"0x0041": "user-label-cluster",
			}},
	},
	"LaundryWasherControls.adoc": {
		ZAP: ZAP{TemplatePath: "washer-controls-cluster"},
	},
	"LevelControl.adoc": {
		ZAP: ZAP{DefineOverrides: map[string]string{"REMAINING_TIME": "LEVEL_CONTROL_REMAINING_TIME"},
			SuppressClusterDefinePrefix: true},
	},
	"LocalizationTimeFormat.adoc": {
		ZAP: ZAP{TemplatePath: "time-format-localization-cluster"},
	},
	"LocalizationUnit.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP,
			TemplatePath: "unit-localization-cluster"},
	},
	"meas_and_sense.adoc": {
		ZAP: ZAP{TemplatePath: "measurement-and-sensing"},
	},
	"MediaInput.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "MEDIA_INPUT_",
			DefineOverrides: map[string]string{"MEDIA_INPUT_INPUT_LIST": "MEDIA_INPUT_LIST"}},
	},
	"MediaPlayback.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "MEDIA_PLAYBACK_",
			DefineOverrides: map[string]string{
				"MEDIA_PLAYBACK_CURRENT_STATE":    "MEDIA_PLAYBACK_STATE",
				"MEDIA_PLAYBACK_SAMPLED_POSITION": "MEDIA_PLAYBACK_PLAYBACK_POSITION",
				"MEDIA_PLAYBACK_SEEK_RANGE_END":   "MEDIA_PLAYBACK_PLAYBACK_SEEK_RANGE_END",
				"MEDIA_PLAYBACK_SEEK_RANGE_START": "MEDIA_PLAYBACK_PLAYBACK_SEEK_RANGE_START",
			}},
	},
	"MicrowaveOvenControl.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"ModeSelect.adoc": {
		ZAP: ZAP{DefineOverrides: map[string]string{"DESCRIPTION": "MODE_DESCRIPTION"}},
	},
	"Mode_Dishwasher.adoc": {
		ZAP: ZAP{TemplatePath: "dishwasher-mode-cluster"},
	},
	"Mode_LaundryWasher.adoc": {
		ZAP: ZAP{TemplatePath: "laundry-washer-mode-cluster"},
	},
	"Mode_MicrowaveOven.adoc": {
		ZAP: ZAP{TemplatePath: "microwave-oven-mode-cluster"},
	},
	"Mode_Oven.adoc": {
		ZAP: ZAP{TemplatePath: "oven-mode-cluster"},
	},
	"Mode_Refrigerator.adoc": {
		ZAP: ZAP{TemplatePath: "refrigerator-and-temperature-controlled-cabinet-mode-cluster"},
	},
	"Mode_RVCClean.adoc": {
		ZAP: ZAP{TemplatePath: "rvc-clean-mode-cluster"},
	},
	"Mode_RVCRun.adoc": {
		ZAP: ZAP{TemplatePath: "rvc-run-mode-cluster"},
	},
	"OnOff.adoc": {
		ZAP: ZAP{TemplatePath: "onoff-cluster"},
	},
	"OperationalCredentialCluster.adoc": {
		ZAP: ZAP{TemplatePath: "operational-credentials-cluster"},
	},
	"OperationalState_RVC": {
		ZAP: ZAP{TemplatePath: "operational-state-rvc-cluster"},
	},
	"PowerSourceConfigurationCluster.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"PowerSourceCluster.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"PressureMeasurement.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "PRESSURE_"},
	},
	"PumpConfigurationControl.adoc": {
		ZAP: ZAP{TemplatePath: "pump-configuration-and-control-cluster"},
	},
	"RefrigeratorAlarm.adoc": {
		ZAP: ZAP{TemplatePath: "refrigerator-alarm"},
	},
	"ResourceMonitoring.adoc": {
		ZAP: ZAP{SeparateStructs: map[string]struct{}{"ReplacementProductStruct": {}}},
	},
	"Scenes.adoc": {
		Spec: Spec{
			IgnoreSections: map[string]struct{}{"Logical Scene Table": {}},
		},
		ZAP: ZAP{TemplatePath: "scene"},
	},
	"SmokeCOAlarm.adoc": {
		ZAP: ZAP{DefineOverrides: map[string]string{
			"HARDWARE_FAULT_ALERT":    "HARDWARE_FAULTALERT",
			"END_OF_SERVICE_ALERT":    "END_OF_SERVICEALERT",
			"SMOKE_SENSITIVITY_LEVEL": "SENSITIVITY_LEVEL",
		}},
	},
	"Switch.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP}, // wth?
	},
	"TargetNavigator.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "TARGET_NAVIGATOR_",
			DefineOverrides: map[string]string{
				"TARGET_NAVIGATOR_TARGET_LIST": "TARGET_NAVIGATOR_LIST",
			}},
	},
	"TemperatureControl.adoc": {
		ZAP: ZAP{DefineOverrides: map[string]string{
			"TEMPERATURE_SETPOINT":         "TEMP_SETPOINT",
			"MIN_TEMPERATURE":              "MIN_TEMP",
			"MAX_TEMPERATURE":              "MAX_TEMP",
			"SELECTED_TEMPERATURE_LEVEL":   "SELECTED_TEMP_LEVEL",
			"SUPPORTED_TEMPERATURE_LEVELS": "SUPPORTED_TEMP_LEVELS",
		}},
	},
	"TemperatureMeasurement.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "TEMP_"},
	},
	"Thermostat.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true,
			DefineOverrides: map[string]string{
				"OCCUPANCY": "THERMOSTAT_OCCUPANCY",
			}},
	},
	"ThermostatUserInterfaceConfiguration.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"TimeSync.adoc": {
		ZAP: ZAP{TemplatePath: "time-synchronization-cluster"},
	},
	"ValidProxies-Cluster.adoc": {
		ZAP: ZAP{TemplatePath: "proxy-valid-cluster"},
	},
	"ValveConfigurationControl.adoc": {
		ZAP: ZAP{TemplatePath: "valve-configuration-and-control-cluster"},
	},
	"WakeOnLAN.adoc": {
		ZAP: ZAP{DefineOverrides: map[string]string{
			"MAC_ADDRESS": "WAKE_ON_LAN_MAC_ADDRESS",
		}},
	},
	"WaterContentMeasurement.adoc": {
		ZAP: ZAP{TemplatePath: "relative-humidity-measurement-cluster"},
	},
	"WaterControls.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"WindowCovering.adoc": {
		ZAP: ZAP{TemplatePath: "window-covering",
			ClusterDefinePrefix: "WC_",
			DefineOverrides: map[string]string{
				"WC_TARGET_POSITION_LIFT_PERCENT_100_THS":  "WC_TARGET_POSITION_LIFT_PERCENT100THS",
				"WC_TARGET_POSITION_TILT_PERCENT_100_THS":  "WC_TARGET_POSITION_TILT_PERCENT100THS",
				"WC_CURRENT_POSITION_LIFT_PERCENT_100_THS": "WC_CURRENT_POSITION_LIFT_PERCENT100THS",
				"WC_CURRENT_POSITION_TILT_PERCENT_100_THS": "WC_CURRENT_POSITION_TILT_PERCENT100THS",
			}},
	},
}
