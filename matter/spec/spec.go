package spec

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type ClusterRefs map[types.Entity]map[*matter.Cluster]struct{}

type Specification struct {
	ClustersByID   map[uint64]*matter.Cluster
	ClustersByName map[string]*matter.Cluster
	DeviceTypes    map[uint64]*matter.DeviceType
	BaseDeviceType *matter.DeviceType

	ClusterRefs ClusterRefs
	DocRefs     map[types.Entity]string

	Bitmaps map[string]*matter.Bitmap
	Enums   map[string]*matter.Enum
	Structs map[string]*matter.Struct

	entities map[string]map[types.Entity]*matter.Cluster

	DocGroups []*DocGroup
}

func newSpec() *Specification {
	return &Specification{

		ClustersByID:   make(map[uint64]*matter.Cluster),
		ClustersByName: make(map[string]*matter.Cluster),
		DeviceTypes:    make(map[uint64]*matter.DeviceType),
		ClusterRefs:    make(map[types.Entity]map[*matter.Cluster]struct{}),
		DocRefs:        make(map[types.Entity]string),

		Bitmaps: make(map[string]*matter.Bitmap),
		Enums:   make(map[string]*matter.Enum),
		Structs: make(map[string]*matter.Struct),

		entities: make(map[string]map[types.Entity]*matter.Cluster),
	}
}

func (cr ClusterRefs) Add(c *matter.Cluster, m types.Entity) {
	cm, ok := cr[m]
	if !ok {
		cm = make(map[*matter.Cluster]struct{})
		cr[m] = cm
	}
	cm[c] = struct{}{}
}
