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
	DocTypeMultiAdmin
	DocTypeNamespaces
	DocTypeQRCode
	DocTypeRendezvous
	DocTypeSecureChannel
	DocTypeServiceDeviceManagement
	DocTypeSoftAP
)

var DocTypeNames = map[DocType]string{
	DocTypeUnknown:                 "Unknown",
	DocTypeAppCluster:              "AppCluster",
	DocTypeAppClusterIndex:         "AppClusterIndex",
	DocTypeDeviceType:              "DeviceType",
	DocTypeDeviceTypeIndex:         "DeviceTypeIndex",
	DocTypeCommonProtocol:          "CommonProtocol",
	DocTypeDataModel:               "DataModel",
	DocTypeDeviceAttestation:       "DeviceAttestation",
	DocTypeMultiAdmin:              "MultiAdmin",
	DocTypeNamespaces:              "Namespaces",
	DocTypeQRCode:                  "QRCode",
	DocTypeRendezvous:              "Rendezvous",
	DocTypeSecureChannel:           "SecureChannel",
	DocTypeServiceDeviceManagement: "ServiceDeviceManagement",
	DocTypeSoftAP:                  "SoftAP",
}
