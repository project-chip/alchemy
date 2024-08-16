package spec

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type ClusterRefs map[types.Entity]map[*matter.Cluster]struct{}

type Specification struct {
	ClustersByID   map[uint64]*matter.Cluster
	ClustersByName map[string]*matter.Cluster
	DeviceTypes    []*matter.DeviceType
	BaseDeviceType *matter.DeviceType

	ClusterRefs ClusterRefs
	DocRefs     map[types.Entity]string

	bitmapIndex  map[string]*matter.Bitmap
	enumIndex    map[string]*matter.Enum
	structIndex  map[string]*matter.Struct
	commandIndex map[string]*matter.Command
	eventIndex   map[string]*matter.Event

	GlobalObjects map[types.Entity]struct{}

	entities map[string]map[types.Entity]*matter.Cluster

	DocGroups []*DocGroup
}

func newSpec() *Specification {
	return &Specification{

		ClustersByID:   make(map[uint64]*matter.Cluster),
		ClustersByName: make(map[string]*matter.Cluster),
		ClusterRefs:    make(map[types.Entity]map[*matter.Cluster]struct{}),
		DocRefs:        make(map[types.Entity]string),

		bitmapIndex:  make(map[string]*matter.Bitmap),
		enumIndex:    make(map[string]*matter.Enum),
		structIndex:  make(map[string]*matter.Struct),
		commandIndex: make(map[string]*matter.Command),
		eventIndex:   make(map[string]*matter.Event),

		GlobalObjects: make(map[types.Entity]struct{}),

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
