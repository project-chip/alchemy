package asciidoc

type raw string

func (r raw) Raw() string {
	return string(r)
}

func (r *raw) SetRaw(s string) {
	*r = raw(s)
}

type HasRaw interface {
	Raw() string
	SetRaw(s string)
}
