package matter

type Field struct {
	ID   int
	Name string
	Type string

	Constraint  string
	Quality     string
	Access      map[AccessCategory]string
	Default     string
	Conformance string
}
