package spec

import (
	"sync"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type ClusterRefs struct {
	sync.RWMutex
	refs map[types.Entity]map[*matter.Cluster]struct{}
}

type Specification struct {
	ClustersByID   map[uint64]*matter.Cluster
	ClustersByName map[string]*matter.Cluster
	DeviceTypes    []*matter.DeviceType
	BaseDeviceType *matter.DeviceType

	ClusterRefs ClusterRefs
	DocRefs     map[types.Entity]*Doc

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
		ClusterRefs:    ClusterRefs{refs: make(map[types.Entity]map[*matter.Cluster]struct{})},
		DocRefs:        make(map[types.Entity]*Doc),

		bitmapIndex:  make(map[string]*matter.Bitmap),
		enumIndex:    make(map[string]*matter.Enum),
		structIndex:  make(map[string]*matter.Struct),
		commandIndex: make(map[string]*matter.Command),
		eventIndex:   make(map[string]*matter.Event),

		GlobalObjects: make(map[types.Entity]struct{}),

		entities: make(map[string]map[types.Entity]*matter.Cluster),
	}
}

func (cr *ClusterRefs) Add(c *matter.Cluster, m types.Entity) {
	cr.Lock()
	cm, ok := cr.refs[m]
	if !ok {
		cm = make(map[*matter.Cluster]struct{})
		cr.refs[m] = cm
	}
	cr.Unlock()
	cm[c] = struct{}{}
}

func (cr *ClusterRefs) Get(m types.Entity) (map[*matter.Cluster]struct{}, bool) {
	cr.RLock()
	cm, ok := cr.refs[m]
	cr.RUnlock()
	return cm, ok
}
