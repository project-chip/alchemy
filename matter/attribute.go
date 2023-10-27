package matter

type Attribute struct {
	ID   string
	Name string
	Type *DataType

	Constraint  string
	Quality     Quality
	Access      map[AccessCategory]string
	Default     string
	Conformance string
}
