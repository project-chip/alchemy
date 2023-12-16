package matter

type DocType uint8

const (
	DocTypeUnknown DocType = iota
	DocTypeCluster
	DocTypeAppClusterIndex
	DocTypeAppClusters
	DocTypeDeviceType
	DocTypeDeviceTypes
	DocTypeDeviceTypeIndex
	DocTypeCommonProtocol
	DocTypeDataModel
	DocTypeDeviceAttestation
	DocTypeMultiAdmin
	DocTypeNamespace
	DocTypeNamespaces
	DocTypeQRCode
	DocTypeRendezvous
	DocTypeSecureChannel
	DocTypeServiceDeviceManagement
	DocTypeSoftAP
)

var DocTypeNames = map[DocType]string{
	DocTypeUnknown:                 "Unknown",
	DocTypeAppClusterIndex:         "AppClusterIndex",
	DocTypeAppClusters:             "AppClusters",
	DocTypeDeviceType:              "DeviceType",
	DocTypeDeviceTypes:             "DeviceTypes",
	DocTypeDeviceTypeIndex:         "DeviceTypeIndex",
	DocTypeCluster:                 "Cluster",
	DocTypeCommonProtocol:          "CommonProtocol",
	DocTypeDataModel:               "DataModel",
	DocTypeDeviceAttestation:       "DeviceAttestation",
	DocTypeMultiAdmin:              "MultiAdmin",
	DocTypeNamespace:               "Namespace",
	DocTypeNamespaces:              "Namespaces",
	DocTypeQRCode:                  "QRCode",
	DocTypeRendezvous:              "Rendezvous",
	DocTypeSecureChannel:           "SecureChannel",
	DocTypeServiceDeviceManagement: "ServiceDeviceManagement",
	DocTypeSoftAP:                  "SoftAP",
}
