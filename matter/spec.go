package matter

type Spec struct {
	Clusters    map[uint64]*Cluster
	DeviceTypes map[uint64]*DeviceType

	ClusterRefs map[Model]map[*Cluster]struct{}
	DocRefs     map[Model]string

	Bitmaps map[string]*Bitmap
	Enums   map[string]*Enum
	Structs map[string]*Struct
}
