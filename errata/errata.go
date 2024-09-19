package errata

import (
	"github.com/project-chip/alchemy/matter"
)

type Errata struct {
	Disco    Disco    `yaml:"disco,omitempty"`
	Spec     Spec     `yaml:"spec,omitempty"`
	TestPlan TestPlan `yaml:"test-plan,omitempty"`
	ZAP      ZAP      `yaml:"zap,omitempty"`
}

var DefaultErrata = &Errata{}

func GetErrata(path string) *Errata {
	errata, ok := Erratas[path]
	if ok {
		return errata
	}
	return DefaultErrata
}

var Erratas = map[string]*Errata{
	"src/matter-defines.adoc": {
		Spec: Spec{UtilityInclude: true},
	},
	"src/app_clusters/AirQuality.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/BallastConfiguration.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/BooleanState.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/booleanstate.adoc",
		},
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/ColorControl.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/colorcontrol.adoc",
		},
		ZAP: ZAP{ClusterDefinePrefix: "COLOR_CONTROL_"},
	},
	"src/app_clusters/ConcentrationMeasurement.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/DemandResponseLoadControl.adoc": {
		ZAP: ZAP{
			TemplatePath: "drlc-cluster",
			ClusterAliases: map[string][]string{
				"Demand Response Load Control": {"Demand Response and Load Control"},
			},
			DefineOverrides: map[string]string{
				"EVENTS":        "LOAD_CONTROL_EVENTS",
				"ACTIVE_EVENTS": "LOAD_CONTROL_ACTIVE_EVENTS",
			}},
	},
	"src/app_clusters/DoorLock.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/door_lock_Cluster.adoc",
		},
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
	"src/app_clusters/EnergyEVSE.adoc": {
		ZAP: ZAP{TemplatePath: "energy-evse-cluster",
			SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/FanControl.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/FanControl.adoc",
		},
		ZAP: ZAP{SuppressAttributePermissions: true,
			SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/FlowMeasurement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/flowmeasurement.adoc",
		},
		ZAP: ZAP{ClusterDefinePrefix: "FLOW_"},
	},
	"src/app_clusters/Groups.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "GROUP_"},
	},
	"src/app_clusters/IlluminanceMeasurement.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "ILLUM_"},
	},
	"src/app_clusters/LaundryDryerControls.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/LaundryDryerControls.adoc",
		},
	},
	"src/app_clusters/LaundryWasherControls.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/LaundryWasherControls.adoc",
		},
		ZAP: ZAP{TemplatePath: "washer-controls-cluster"},
	},
	"src/app_clusters/LevelControl.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/levelcontrol.adoc",
		},
		ZAP: ZAP{DefineOverrides: map[string]string{"REMAINING_TIME": "LEVEL_CONTROL_REMAINING_TIME"},
			SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/meas_and_sense.adoc": {
		ZAP: ZAP{TemplatePath: "measurement-and-sensing"},
	},
	"src/app_clusters/MicrowaveOvenControl.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/MicrowaveOvenControl.adoc",
		},
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/ModeBase.adoc": {
		Spec: Spec{
			Sections: map[string]SpecSection{
				"Mode Base Status CommonCodes Range": {Skip: SpecPurposeDataTypesEnum},
			},
		},
	},
	"src/app_clusters/ModeBase_ModeTag_BaseValues.adoc": {
		Spec: Spec{UtilityInclude: true},
	},
	"src/app_clusters/ModeSelect.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/modeselect.adoc",
		},
		ZAP: ZAP{DefineOverrides: map[string]string{"DESCRIPTION": "MODE_DESCRIPTION"}},
	},
	"src/app_clusters/Mode_DeviceEnergyManagement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_device_energy_management.adoc",
		},
	},
	"src/app_clusters/Mode_Dishwasher.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_dishwasher.adoc",
		},
		ZAP: ZAP{TemplatePath: "dishwasher-mode-cluster"},
	},
	"src/app_clusters/Mode_EVSE.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_energy_EVSE.adoc",
		},
	},
	"src/app_clusters/Mode_LaundryWasher.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_laundry_washer.adoc",
		},
		ZAP: ZAP{TemplatePath: "laundry-washer-mode-cluster"},
	},
	"src/app_clusters/Mode_MicrowaveOven.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_MicrowaveOven.adoc",
		},
		ZAP: ZAP{TemplatePath: "microwave-oven-mode-cluster"},
	},
	"src/app_clusters/Mode_Oven.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_oven.adoc",
		},
		ZAP: ZAP{TemplatePath: "oven-mode-cluster"},
	},
	"src/app_clusters/Mode_Refrigerator.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_ref_tcc.adoc",
		},
		ZAP: ZAP{
			TemplatePath: "refrigerator-and-temperature-controlled-cabinet-mode-cluster",
			ClusterAliases: map[string][]string{
				"Refrigerator And Temperature Controlled Cabinet Mode": {"Refrigerator and Temperature Controlled Cabinet Mode"},
			},
		},
	},
	"src/app_clusters/Mode_RVCClean.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_rvc_clean.adoc",
		},
		ZAP: ZAP{TemplatePath: "rvc-clean-mode-cluster"},
	},
	"src/app_clusters/Mode_RVCRun.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_rvc_run.adoc",
		},
		ZAP: ZAP{TemplatePath: "rvc-run-mode-cluster"},
	},
	"src/app_clusters/Mode_WaterHeater.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/mode_WaterHeater.adoc",
		},
	},
	"src/app_clusters/OccupancySensing.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/occupancysensing.adoc",
		},
	},
	"src/app_clusters/OnOff.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/onoff.adoc",
		},
		ZAP: ZAP{
			TemplatePath: "onoff-cluster",
			ClusterAliases: map[string][]string{
				"On/Off": {"OnOff"},
			},
		},
	},
	"src/app_clusters/OperationalState.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/operationalstate.adoc",
		},
		Spec: Spec{
			Sections: map[string]SpecSection{
				"ErrorStateEnum GeneralErrors Range": {Skip: SpecPurposeDataTypesEnum},
				"Resume Command":                     {Skip: SpecPurposeCommandArguments},
				"Pause Command":                      {Skip: SpecPurposeCommandArguments},
			},
		},
	},
	"src/app_clusters/OperationalState_ErrorStateEnum_BaseValues.adoc": {
		Spec: Spec{UtilityInclude: true},
	},
	"src/app_clusters/OperationalState_OperationalStateEnum_BaseValues.adoc": {
		Spec: Spec{UtilityInclude: true},
	},
	"src/app_clusters/OperationalState_Oven.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/ovenoperationalstate.adoc",
		},
	},
	"src/app_clusters/OperationalState_RVC.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/rvcoperationalstate.adoc",
		},
		ZAP: ZAP{TemplatePath: "operational-state-rvc-cluster"},
	},
	"src/app_clusters/PressureMeasurement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/pressuremeasurement.adoc",
		},
		ZAP: ZAP{ClusterDefinePrefix: "PRESSURE_"},
	},
	"src/app_clusters/PumpConfigurationControl.adoc": {
		ZAP: ZAP{TemplatePath: "pump-configuration-and-control-cluster"},
	},
	"src/app_clusters/RefrigeratorAlarm.adoc": {
		ZAP: ZAP{TemplatePath: "refrigerator-alarm"},
	},
	"src/app_clusters/ResourceMonitoring.adoc": {
		ZAP: ZAP{SeparateStructs: map[string]struct{}{"ReplacementProductStruct": {}}},
	},
	"src/app_clusters/Scenes.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/scenes.adoc",
		},
		Spec: Spec{
			Sections: map[string]SpecSection{
				"Form of ExtensionFieldSetStruct": {Skip: SpecPurposeDataTypes},
				"Logical Scene Table":             {Skip: SpecPurposeDataTypes},
			},
		},
		ZAP: ZAP{TemplatePath: "scene"},
	},
	"src/app_clusters/SmokeCOAlarm.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/smco.adoc",
		},
		ZAP: ZAP{DefineOverrides: map[string]string{
			"HARDWARE_FAULT_ALERT":    "HARDWARE_FAULTALERT",
			"END_OF_SERVICE_ALERT":    "END_OF_SERVICEALERT",
			"SMOKE_SENSITIVITY_LEVEL": "SENSITIVITY_LEVEL",
		}},
	},
	"src/app_clusters/Switch.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP}, // wth?
	},
	"src/app_clusters/TemperatureControl.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/temperaturecontrol.adoc",
		},
		ZAP: ZAP{DefineOverrides: map[string]string{
			"TEMPERATURE_SETPOINT":         "TEMP_SETPOINT",
			"MIN_TEMPERATURE":              "MIN_TEMP",
			"MAX_TEMPERATURE":              "MAX_TEMP",
			"SELECTED_TEMPERATURE_LEVEL":   "SELECTED_TEMP_LEVEL",
			"SUPPORTED_TEMPERATURE_LEVELS": "SUPPORTED_TEMP_LEVELS",
		}},
	},
	"src/app_clusters/TemperatureMeasurement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/temperaturemeasurement.adoc",
		},
		ZAP: ZAP{ClusterDefinePrefix: "TEMP_"},
	},
	"src/app_clusters/Thermostat.adoc": {
		ZAP: ZAP{SuppressClusterDefinePrefix: true,
			DefineOverrides: map[string]string{
				"OCCUPANCY": "THERMOSTAT_OCCUPANCY",
			}},
	},
	"src/app_clusters/ThermostatUserInterfaceConfiguration.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/thermostatuserconfiguration.adoc",
		},
		ZAP: ZAP{SuppressClusterDefinePrefix: true},
	},
	"src/app_clusters/ValveConfigurationControl.adoc": {
		ZAP: ZAP{TemplatePath: "valve-configuration-and-control-cluster"},
	},
	"src/app_clusters/WaterContentMeasurement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/relativehumiditymeasurement.adoc",
		},
		ZAP: ZAP{TemplatePath: "relative-humidity-measurement-cluster"},
	},
	"src/app_clusters/WaterHeaterManagement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/WaterHeaterManagement.adoc",
		},
	},
	"src/app_clusters/WiFiNetworkManagement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/wifi_network_management.adoc",
		},
	},
	"src/app_clusters/WindowCovering.adoc": {
		ZAP: ZAP{TemplatePath: "window-covering",
			ClusterDefinePrefix: "WC_",
			DefineOverrides: map[string]string{
				"WC_TARGET_POSITION_LIFT_PERCENT_100_THS":  "WC_TARGET_POSITION_LIFT_PERCENT100THS",
				"WC_TARGET_POSITION_TILT_PERCENT_100_THS":  "WC_TARGET_POSITION_TILT_PERCENT100THS",
				"WC_CURRENT_POSITION_LIFT_PERCENT_100_THS": "WC_CURRENT_POSITION_LIFT_PERCENT100THS",
				"WC_CURRENT_POSITION_TILT_PERCENT_100_THS": "WC_CURRENT_POSITION_TILT_PERCENT100THS",
			}},
	},
	"src/app_clusters/media/ApplicationBasic.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "APPLICATION_",
			DefineOverrides: map[string]string{"APPLICATION_APPLICATION": "APPLICATION_APP"}},
	},
	"src/app_clusters/media/ApplicationLauncher.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "APPLICATION_LAUNCHER_",
			DefineOverrides: map[string]string{
				"APPLICATION_LAUNCHER_CATALOG_LIST": "APPLICATION_LAUNCHER_LIST",
			}},
	},
	"src/app_clusters/media/AudioOutput.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "AUDIO_OUTPUT_",
			DefineOverrides: map[string]string{
				"AUDIO_OUTPUT_OUTPUT_LIST": "AUDIO_OUTPUT_LIST",
			}},
	},
	"src/app_clusters/media/Channel.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "CHANNEL_"},
	},
	"src/app_clusters/media/ContentLauncher.adoc": {
		ZAP: ZAP{TemplatePath: "content-launch-cluster",
			ClusterDefinePrefix: "CONTENT_LAUNCHER_"},
	},
	"src/app_clusters/media/KeypadInput.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "ILLUM_"},
	},
	"src/app_clusters/media/MediaInput.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "MEDIA_INPUT_",
			DefineOverrides: map[string]string{"MEDIA_INPUT_INPUT_LIST": "MEDIA_INPUT_LIST"}},
	},
	"src/app_clusters/media/MediaPlayback.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "MEDIA_PLAYBACK_",
			DefineOverrides: map[string]string{
				"MEDIA_PLAYBACK_CURRENT_STATE":    "MEDIA_PLAYBACK_STATE",
				"MEDIA_PLAYBACK_SAMPLED_POSITION": "MEDIA_PLAYBACK_PLAYBACK_POSITION",
				"MEDIA_PLAYBACK_SEEK_RANGE_END":   "MEDIA_PLAYBACK_PLAYBACK_SEEK_RANGE_END",
				"MEDIA_PLAYBACK_SEEK_RANGE_START": "MEDIA_PLAYBACK_PLAYBACK_SEEK_RANGE_START",
			}},
	},
	"src/app_clusters/media/TargetNavigator.adoc": {
		ZAP: ZAP{ClusterDefinePrefix: "TARGET_NAVIGATOR_",
			DefineOverrides: map[string]string{
				"TARGET_NAVIGATOR_TARGET_LIST": "TARGET_NAVIGATOR_LIST",
			}},
	},
	"src/app_clusters/media/WakeOnLAN.adoc": {
		ZAP: ZAP{
			DefineOverrides: map[string]string{
				"MAC_ADDRESS": "WAKE_ON_LAN_MAC_ADDRESS",
			},
			ClusterAliases: map[string][]string{
				"Wake on LAN": {"WakeOnLAN"},
			},
		},
	},
	"src/data_model/ACL-Cluster.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/AccessControl.adoc",
		},
		ZAP: ZAP{
			TemplatePath: "access-control-cluster",
			ClusterAliases: map[string][]string{
				"AccessControl": {"Access Control"},
			},
		},
	},
	"src/data_model/bridge-clusters.adoc": {
		ZAP: ZAP{ClusterSplit: map[string]string{
			"0x0025": "actions-cluster",
			"0x0039": "bridged-device-basic-information",
		}},
	},
	"src/data_model/Data-Model.adoc": {
		Spec: Spec{
			Sections: map[string]SpecSection{
				"Struct Type":          {Skip: SpecPurposeDataTypesStruct},
				"Quality Conformance":  {Skip: SpecPurposeDataTypesStruct},
				"Choice Conformance":   {Skip: SpecPurposeDataTypesStruct},
				"Fabric-Scoped Struct": {Skip: SpecPurposeDataTypesStruct},
			},
		},
	},
	"src/data_model/Encoding-Specification.adoc": {
		Spec: Spec{
			Sections: map[string]SpecSection{
				"Discrete - Bitmap":   {Skip: SpecPurposeDataTypes},
				"Collection - Struct": {Skip: SpecPurposeDataTypes},
			},
		},
	},
	"src/data_model/Group-Key-Management-Cluster.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/group_communication.adoc",
		},
		ZAP: ZAP{
			TemplatePath: "group-key-mgmt-cluster",
			ClusterAliases: map[string][]string{
				"GroupKeyManagement": {"Group Key Management"},
			},
		},
	},
	"src/data_model/ICDManagement.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/icdmanagement.adoc",
		},
		ZAP: ZAP{
			ClusterAliases: map[string][]string{
				"ICDManagement": {"ICD Management"},
			},
		},
	},
	"src/data_model/Label-Cluster.adoc": {
		ZAP: ZAP{TemplatePath: "user-label-cluster",
			ClusterSplit: map[string]string{
				"0x0040": "fixed-label-cluster",
				"0x0041": "user-label-cluster",
			}},
	},
	"src/data_model/ValidProxies-Cluster.adoc": {
		ZAP: ZAP{TemplatePath: "proxy-valid-cluster"},
	},
	"src/device_types/OtaRequestor.adoc": {
		Disco: Disco{
			Sections: map[string]DiscoSection{
				"ProviderLocation Type": {
					Skip: DiscoPurposeDataTypeRename,
				},
			},
		},
	},
	"src/secure_channel/Discovery.adoc": {
		Spec: Spec{
			Sections: map[string]SpecSection{
				"Common TXT Key/Value Pairs":      {Skip: SpecPurposeDataTypes},
				"TXT key for pairing hint (`PH`)": {Skip: SpecPurposeDataTypes},
			},
		},
	},
	"src/service_device_management/AdminAssistedCommissioningFlows.adoc": {
		Disco: Disco{
			Sections: map[string]DiscoSection{
				"Basic Commissioning Method (BCM)":           {Skip: DiscoPurposeNormalizeAnchor},
				"Enhanced Commissioning Method (ECM)":        {Skip: DiscoPurposeNormalizeAnchor},
				"Presentation of Onboarding Payload for ECM": {Skip: DiscoPurposeNormalizeAnchor},
				"Open Commissioning Window":                  {Skip: DiscoPurposeNormalizeAnchor},
			},
		},
	},
	"src/service_device_management/AdminCommissioningCluster.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/multiplefabrics.adoc",
		},
		ZAP: ZAP{TemplatePath: "administrator-commissioning-cluster"},
	},
	"src/service_device_management/BasicInformationCluster.adoc": {
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"src/service_device_management/DeviceCommissioningFlows.adoc": {
		Spec: Spec{
			Sections: map[string]SpecSection{"Enhanced Setup Flow (ESF)": {Skip: SpecPurposeDataTypes}},
		},
	},
	"src/service_device_management/DiagnosticsGeneral.adoc": {
		Disco: Disco{
			Sections: map[string]DiscoSection{"NetworkInterface Type": {Skip: DiscoPurposeDataTypeRename}},
		},
		ZAP: ZAP{TemplatePath: "general-diagnostics-cluster"},
	},
	"src/service_device_management/DiagnosticsEthernet.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/ethernet_diagnostics.adoc",
		},
		ZAP: ZAP{TemplatePath: "ethernet-network-diagnostics-cluster"},
	},
	"src/service_device_management/DiagnosticLogsCluster.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/logs_diagnostics.adoc",
		},
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"src/service_device_management/DiagnosticsSoftware.adoc": {
		ZAP: ZAP{TemplatePath: "software-diagnostics-cluster"},
	},
	"src/service_device_management/DiagnosticsThread.adoc": {
		Disco: Disco{
			Sections: map[string]DiscoSection{
				"SecurityPolicy Type": {
					Skip: DiscoPurposeDataTypeRename,
				},
				"OperationalDatasetComponents Type": {
					Skip: DiscoPurposeDataTypeRename,
				},
			},
		},
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/thread_nw_diagnostics.adoc",
		},
		ZAP: ZAP{TemplatePath: "thread-network-diagnostics-cluster",
			SuppressClusterDefinePrefix: true},
	},
	"src/service_device_management/DiagnosticsWiFi.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/wifi_diagnostics.adoc",
		},
		ZAP: ZAP{TemplatePath: "wifi-network-diagnostics-cluster"},
	},
	"src/service_device_management/GeneralCommissioningCluster.adoc": {
		Disco: Disco{
			Sections: map[string]DiscoSection{"BasicCommissioningInfo Type": {Skip: DiscoPurposeDataTypeRename}},
		},
	},
	"src/service_device_management/LocalizationConfiguration.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/Localization.adoc",
		},
	},
	"src/service_device_management/LocalizationTimeFormat.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/timeformatlocalization.adoc",
		},
		ZAP: ZAP{TemplatePath: "time-format-localization-cluster"},
	},
	"src/service_device_management/LocalizationUnit.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/unitlocalization.adoc",
		},
		ZAP: ZAP{Domain: matter.DomainCHIP,
			TemplatePath: "unit-localization-cluster"},
	},
	"src/service_device_management/OperationalCredentialCluster.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/node_operational_credentials.adoc",
		},
		ZAP: ZAP{
			TemplatePath: "operational-credentials-cluster",
			ClusterAliases: map[string][]string{
				"Operational Credentials": {"Node Operational Credentials"},
			},
		},
	},
	"src/service_device_management/PowerSourceConfigurationCluster.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/powersourceconfiguration.adoc",
		},
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"src/service_device_management/PowerSourceCluster.adoc": {
		TestPlan: TestPlan{
			TestPlanPath: "src/cluster/powersource.adoc",
		},
		ZAP: ZAP{Domain: matter.DomainCHIP},
	},
	"src/service_device_management/TimeSync.adoc": {
		ZAP: ZAP{TemplatePath: "time-synchronization-cluster"},
	},
}
