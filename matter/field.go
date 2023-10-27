package matter

type Field struct {
	ID   string
	Name string
	Type *DataType

	MinLength int
	MaxLength int

	Constraint  string
	Quality     string
	Access      map[AccessCategory]string
	Default     string
	Conformance string
}
