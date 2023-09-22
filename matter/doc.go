package matter

type Doc uint8

const (
	DocUnknown Doc = iota
	DocAppCluster
	DocAppClusterIndex
	DocDeviceType
	DocDeviceTypeIndex
	DocCommonProtocol
	DocDataModel
	DocDeviceAttestation
	DocServiceDeviceManagement
)
