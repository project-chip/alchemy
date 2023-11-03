package matter

type Cluster struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	Hierarchy string `json:"hierarchy,omitempty"`
	Role      string `json:"role,omitempty"`
	Scope     string `json:"scope,omitempty"`
	PICS      string `json:"pics,omitempty"`

	Features Features `json:"features,omitempty"`
	//DataTypes  []any      `json:"dataTypes,omitempty"`
	Bitmaps    []*Bitmap  `json:"bitmaps,omitempty"`
	Enums      []*Enum    `json:"enums,omitempty"`
	Structs    []*Struct  `json:"structs,omitempty"`
	Attributes Fields     `json:"attributes,omitempty"`
	Events     []*Event   `json:"events,omitempty"`
	Commands   []*Command `json:"commands,omitempty"`
}

func (c *Cluster) Compare(oc *Cluster) {
	if c.ID != oc.ID {

	}
	if c.Name != oc.Name {

	}

	c.Features.compare(oc.Features)
	c.Attributes.compare(oc.Attributes)
}
