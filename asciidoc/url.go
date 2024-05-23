package asciidoc

type URL struct {
	position
	raw

	Scheme string
	Path   any
}

func NewURL(scheme string, path any) URL {
	return URL{Scheme: scheme, Path: path}
}

func (URL) Type() ElementType {
	return ElementTypeInline
}

func (a URL) Equals(o Element) bool {
	oa, ok := o.(URL)
	if !ok {
		return false
	}
	return a.Scheme == oa.Scheme && a.Path == oa.Path
}
