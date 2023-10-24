package matter

type Command struct {
	ID          string
	Name        string
	Description string
	Direction   string
	Response    string
	Conformance string
	Access      map[AccessCategory]string

	Fields []*Field
}
