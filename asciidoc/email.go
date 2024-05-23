package asciidoc

type Email struct {
	position
	raw

	Address string
}

func NewEmail(address string) Email {
	return Email{Address: address}
}

func (Email) Type() ElementType {
	return ElementTypeInline
}

func (a Email) Equals(o Element) bool {
	oa, ok := o.(Email)
	if !ok {
		return false
	}
	return a.Address == oa.Address
}
