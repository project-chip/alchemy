package matter

type Event struct {
	ID              string
	Name            string
	Description     string
	Priority        string
	FabricSensitive bool
	Conformance     string
	Access          map[AccessCategory]string

	Fields []*Field
}
