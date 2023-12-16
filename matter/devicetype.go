package matter

type DeviceType struct {
	ID          *Number     `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Revisions   []*Revision `json:"revisions,omitempty"`

	Superset string `json:"superset,omitempty"`
	Class    string `json:"class,omitempty"`
	Scope    string `json:"scope,omitempty"`

	Conditions []*Condition `json:"conditions,omitempty"`

	ClusterRequirements []*ClusterRequirement `json:"clusterRequirements,omitempty"`
	ElementRequirements []*ElementRequirement `json:"elementRequirements,omitempty"`
}

func (c *DeviceType) Entity() Entity {
	return EntityDeviceType
}

func (dt *DeviceType) ConformanceReference(name string) HasConformance {
	for _, c := range dt.Conditions {
		if c.Feature == name {
			return c
		}
	}
	return nil
}

type ClusterRequirement struct {
	ID          *Number     `json:"id,omitempty"`
	ClusterName string      `json:"clusterName,omitempty"`
	Quality     Quality     `json:"quality,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`
	Interface   Interface   `json:"interface,omitempty"`

	Cluster *Cluster `json:"cluster,omitempty"`
}

type ElementRequirement struct {
	ID          *Number `json:"id,omitempty"`
	ClusterName string  `json:"clusterName,omitempty"`
	Element     Entity  `json:"element,omitempty"`
	Name        string  `json:"name,omitempty"`

	Constraint  Constraint  `json:"constraint,omitempty"`
	Quality     Quality     `json:"quality,omitempty"`
	Access      Access      `json:"access,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`

	Cluster *Cluster `json:"cluster,omitempty"`
}

type Condition struct {
	Feature     string
	Description string
}

func (c *Condition) Entity() Entity {
	return EntityCondition
}

func (c *Condition) GetConformance() Conformance {
	return nil
}
