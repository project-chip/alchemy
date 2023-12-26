package matter

type ClusterRefs map[Model]map[*Cluster]struct{}

type Spec struct {
	ClustersByID   map[uint64]*Cluster
	ClustersByName map[string]*Cluster
	DeviceTypes    map[uint64]*DeviceType

	ClusterRefs ClusterRefs
	DocRefs     map[Model]string

	Bitmaps map[string]*Bitmap
	Enums   map[string]*Enum
	Structs map[string]*Struct
}

func (cr ClusterRefs) Add(c *Cluster, m Model) {
	cm, ok := cr[m]
	if !ok {
		cm = make(map[*Cluster]struct{})
		cr[m] = cm
	}
	cm[c] = struct{}{}
}
