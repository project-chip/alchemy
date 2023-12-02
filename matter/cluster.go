package matter

type Cluster struct {
	ID          *Number     `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Revisions   []*Revision `json:"revisions,omitempty"`

	Hierarchy string `json:"hierarchy,omitempty"`
	Role      string `json:"role,omitempty"`
	Scope     string `json:"scope,omitempty"`
	PICS      string `json:"pics,omitempty"`

	Features   FeatureSet `json:"features,omitempty"`
	Bitmaps    []*Bitmap  `json:"bitmaps,omitempty"`
	Enums      []*Enum    `json:"enums,omitempty"`
	Structs    []*Struct  `json:"structs,omitempty"`
	Attributes FieldSet   `json:"attributes,omitempty"`
	Events     EventSet   `json:"events,omitempty"`
	Commands   CommandSet `json:"commands,omitempty"`
}

func (c *Cluster) ModelType() Entity {
	return EntityCluster
}
