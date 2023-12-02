package matter

type DeviceType struct {
	ID          *Number `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`

	Superset string `json:"superset,omitempty"`
	Class    string `json:"class,omitempty"`
	Scope    string `json:"scope,omitempty"`

	ClusterRequirements []*ClusterRequirement `json:"clusterRequirements,omitempty"`
	ElementRequirements []*ElementRequirement `json:"elementRequirements,omitempty"`
}

func (c *DeviceType) ModelType() Entity {
	return EntityDeviceType
}

type ClusterRequirement struct {
	ID          *Number   `json:"id,omitempty"`
	Cluster     string    `json:"cluster,omitempty"`
	Quality     Quality   `json:"quality,omitempty"`
	Conformance string    `json:"conformance,omitempty"`
	Interface   Interface `json:"interface,omitempty"`
}

type ElementRequirement struct {
	ID      *Number `json:"id,omitempty"`
	Cluster string  `json:"cluster,omitempty"`
	Element Entity  `json:"element,omitempty"`
	Name    string  `json:"name,omitempty"`

	Constraint  Constraint `json:"constraint,omitempty"`
	Access      Access     `json:"access,omitempty"`
	Conformance string     `json:"conformance,omitempty"`
}
