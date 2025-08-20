package asciidoc

type URL struct {
	position
	raw

	Scheme string
	Path   Elements
}

func NewURL(scheme string, path Elements) URL {
	return URL{Scheme: scheme, Path: path}
}

func (URL) Type() ElementType {
	return ElementTypeInline
}

func (u URL) Equals(o Element) bool {
	oa, ok := o.(URL)
	if !ok {
		return false
	}
	return u.Scheme == oa.Scheme && u.Path.Equals(oa.Path)
}

func (u URL) Clone() Element {
	return &URL{position: u.position, raw: u.raw, Scheme: u.Scheme, Path: u.Path.Clone()}
}
