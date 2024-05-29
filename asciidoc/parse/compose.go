package parse

import (
	"fmt"

	"github.com/hasty/alchemy/asciidoc"
)

type composeOption[T asciidoc.Element] func(t T) error

func compose[T asciidoc.Element](c *current, el T, options ...composeOption[T]) (out T, err error) {
	if hr, ok := any(el).(asciidoc.HasRaw); ok {
		hr.SetRaw(string(c.text))
	}
	if hp, ok := any(el).(asciidoc.HasPosition); ok {
		populatePosition(c, hp)
	}
	for _, o := range options {
		err = o(el)
		if err != nil {
			return
		}
	}
	out = el
	return
}

func setAttributes[T asciidoc.AttributableElement](attributes any) composeOption[T] {
	return func(t T) error {
		if attributes == nil {
			return nil
		}
		switch as := attributes.(type) {
		case []asciidoc.Attribute:
			return t.ReadAttributes(t, as...)
		case asciidoc.Attribute:
			return t.ReadAttributes(t, as)
		default:
			return fmt.Errorf("non-attribute list passed to setAttributes: %T", attributes)
		}
	}
}

func setElements[T asciidoc.HasElements](els any) composeOption[T] {
	return func(t T) error {
		if els == nil {
			return nil
		}
		as, ok := els.(asciidoc.Set)
		if !ok {
			return fmt.Errorf("non-element list passed to setElements: %T", els)
		}
		return t.SetElements(as)
	}
}

func setLines[T asciidoc.HasLines](els any) composeOption[T] {
	return func(t T) error {
		if els == nil {
			return nil
		}
		as, ok := els.([]string)
		if !ok {
			return fmt.Errorf("non-string slice passed to setLines: %T", els)
		}
		t.SetLines(as)
		return nil
	}
}

func populatePosition(c *current, el asciidoc.HasPosition) asciidoc.HasPosition {
	el.SetPath(c.parser.filename)
	el.SetPosition(c.currentPosition())
	return el
}
