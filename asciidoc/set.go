package asciidoc

import "unicode"

type Set []Element

func (s Set) Elements() Set {
	return s
}

func (s *Set) Append(e ...Element) {
	*s = append(*s, e...)
}

func (s *Set) SetElements(els Set) {
	*s = els
}

func (s Set) Equals(o Set) bool {
	if len(s) != len(o) {
		return false
	}
	for i, e := range s {
		oe := o[i]
		if !e.Equals(oe) {
			return false
		}
	}
	return true
}

func (s Set) IsWhitespace() bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		switch v := v.(type) {
		case *String:
			for _, r := range v.Value {
				if !unicode.IsSpace(r) {
					return false
				}
			}
		default:
			return false
		}
	}
	return true
}
