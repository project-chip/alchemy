package matter

import "fmt"

type Field struct {
	ID   string    `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Type *DataType `json:"type,omitempty"`

	Constraint  Constraint `json:"constraint,omitempty"`
	Quality     Quality    `json:"quality,omitempty"`
	Access      Access     `json:"access,omitempty"`
	Default     string     `json:"default,omitempty"`
	Conformance string     `json:"conformance,omitempty"`
}

func (f *Field) Compare(of *Field) {
	if f.Name != of.Name {
		fmt.Printf("field %s name different in ZAP: %s vs %s\n", f.ID, f.Name, of.Name)
	}
	if f.Type.Name != of.Type.Name {
		fmt.Printf("field %s type different in ZAP: %s vs %s\n", f.ID, f.Type.Name, of.Type.Name)
	}
	if f.Type.IsArray != of.Type.IsArray {
		fmt.Printf("field %s array different in ZAP: %v vs %v\n", f.ID, f.Type.IsArray, of.Type.IsArray)
	}
	if !f.Constraint.Equal(of.Constraint) {
		fmt.Printf("field %s constraint different in ZAP: %s vs %s\n", f.ID, f.Constraint.AsciiDocString(), of.Constraint.AsciiDocString())
	}
	if f.Quality != of.Quality {
		fmt.Printf("field %s quality different in ZAP: %v vs %v\n", f.ID, f.Quality, of.Quality)
	}
}

type Fields []*Field

func (fs Fields) compare(ofs Fields) {
	fieldMap := make(map[string]*Field)
	for _, f := range fs {
		fieldMap[f.ID] = f
	}

	oFieldMap := make(map[string]*Field)
	for _, f := range ofs {
		oFieldMap[f.ID] = f
	}

	for code, of := range oFieldMap {
		f, ok := fieldMap[of.ID]
		if !ok {
			continue
		}
		delete(oFieldMap, code)
		delete(fieldMap, code)
		f.Compare(of)
	}
	for _, f := range fieldMap {
		fmt.Printf("field %s not present in ZAP\n", f.Name)
	}
	for _, f := range oFieldMap {
		fmt.Printf("field %s not present in spec\n", f.Name)
	}

}
