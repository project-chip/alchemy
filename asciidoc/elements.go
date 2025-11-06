package asciidoc

import "unicode"

type Elements []Element

func (s Elements) Children() Elements {
	return s
}

func (s *Elements) Append(e ...Element) {
	*s = append(*s, e...)
}

func (s *Elements) SetChildren(els Elements) {
	*s = els
}

func (s Elements) Equals(o Elements) bool {
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

func (s Elements) Clone() Elements {
	var els Elements
	for _, e := range s {
		els = append(els, e.Clone())
	}
	return els
}

func (s Elements) IsWhitespace() bool {
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
