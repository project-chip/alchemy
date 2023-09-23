package matter

type DocType uint8

const (
	DocTypeUnknown DocType = iota
	DocTypeAppCluster
	DocTypeAppClusterIndex
	DocTypeDeviceType
	DocTypeDeviceTypeIndex
	DocTypeCommonProtocol
	DocTypeDataModel
	DocTypeDeviceAttestation
	DocTypeServiceDeviceManagement
)
