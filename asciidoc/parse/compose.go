package parse

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
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

func setElements[T asciidoc.ParentElement](els any) composeOption[T] {
	return func(t T) error {
		if els == nil {
			return nil
		}
		as, ok := els.(asciidoc.Elements)
		if !ok {
			return fmt.Errorf("non-element list passed to setElements: %T", els)
		}
		t.SetChildren(as)
		return nil
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

func populatePosition[T asciidoc.HasPosition](c *current, el T) T {
	line, col, offset := c.currentPosition()
	switch any(el).(type) {
	case *asciidoc.EmptyLine:
		// Newlines end on the following line, so for logging purposes we shift the line back one
		line -= 1
	}
	el.SetPath(c.parser.filename)
	el.SetPosition(line, col, offset)
	return el
}
