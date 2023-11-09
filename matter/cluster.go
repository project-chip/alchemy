package matter

type Cluster struct {
	ID          *ID    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	Hierarchy string `json:"hierarchy,omitempty"`
	Role      string `json:"role,omitempty"`
	Scope     string `json:"scope,omitempty"`
	PICS      string `json:"pics,omitempty"`

	Features   []*Feature `json:"features,omitempty"`
	Bitmaps    []*Bitmap  `json:"bitmaps,omitempty"`
	Enums      []*Enum    `json:"enums,omitempty"`
	Structs    []*Struct  `json:"structs,omitempty"`
	Attributes FieldSet   `json:"attributes,omitempty"`
	Events     []*Event   `json:"events,omitempty"`
	Commands   []*Command `json:"commands,omitempty"`
}
