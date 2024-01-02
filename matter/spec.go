package matter

import "github.com/hasty/alchemy/matter/types"

type ClusterRefs map[types.Entity]map[*Cluster]struct{}

type Spec struct {
	ClustersByID   map[uint64]*Cluster
	ClustersByName map[string]*Cluster
	DeviceTypes    map[uint64]*DeviceType

	ClusterRefs ClusterRefs
	DocRefs     map[types.Entity]string

	Bitmaps map[string]*Bitmap
	Enums   map[string]*Enum
	Structs map[string]*Struct
}

func (cr ClusterRefs) Add(c *Cluster, m types.Entity) {
	cm, ok := cr[m]
	if !ok {
		cm = make(map[*Cluster]struct{})
		cr[m] = cm
	}
	cm[c] = struct{}{}
}
