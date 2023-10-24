package matter

type Attribute struct {
	ID   int
	Name string
	Type string

	Constraint  string
	Quality     Quality
	Access      map[AccessCategory]string
	Default     string
	Conformance string
}
