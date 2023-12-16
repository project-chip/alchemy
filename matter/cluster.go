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

func (c *Cluster) Entity() Entity {
	return EntityCluster
}

func (c *Cluster) ConformanceReference(name string) HasConformance {
	if c == nil {
		return nil
	}
	var cr HasConformance
	if c.Features != nil {
		cr = c.Features.ConformanceReference(name)
		if cr != nil {
			return cr
		}

	}
	cr = c.Attributes.ConformanceReference(name)
	if cr != nil {
		return cr
	}
	for _, cmd := range c.Commands {
		if cmd.Name == name {
			return cmd
		}
	}
	for _, e := range c.Events {
		if e.Name == name {
			return e
		}
	}
	return nil
}
