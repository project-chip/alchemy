package matter

type Cluster struct {
	ID          string
	Name        string
	Description string

	Hierarchy string
	Role      string
	Scope     string
	PICS      string

	Features   []*Feature
	DataTypes  []interface{}
	Attributes []*Attribute
	Events     []*Event
	Commands   []*Command
}
