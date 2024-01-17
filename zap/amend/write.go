package amend

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

type renderer struct {
	spec *matter.Spec
	doc  *ascii.Doc

	configurator *zap.Configurator

	errata *zap.Errata
}

func newLine(e xmlEncoder) error {
	return e.EncodeToken(xml.CharData{'\n'})
}

type loggingDecoder struct {
	d parse.XmlDecoder
}

func (le *loggingDecoder) Token() (xml.Token, error) {
	tok, err := le.d.Token()
	if err != nil {
		return tok, err
	}
	switch t := tok.(type) {
	case xml.StartElement:
		fmt.Fprintf(os.Stderr, "decoding start element %s\n ", t.Name.Local)
	case xml.EndElement:
		fmt.Fprintf(os.Stderr, "decoding end element %s\n ", t.Name.Local)
	case xml.CharData:
		fmt.Fprintf(os.Stderr, "decoding char data element %s\n ", string(t))
	case xml.Comment:
		fmt.Fprintf(os.Stderr, "decoding comment %s\n ", string(t))
	case xml.ProcInst:
		fmt.Fprintf(os.Stderr, "decoding proc inst\n")
	case xml.Directive:
		fmt.Fprintf(os.Stderr, "decoding directive\n")
	default:

	}
	return tok, err
}

type xmlEncoder interface {
	EncodeToken(t xml.Token) error
	Close() error
}

type loggingEncoder struct {
	w io.Writer
	e *xml.Encoder
}

func (le *loggingEncoder) EncodeToken(t xml.Token) error {
	switch t := t.(type) {
	case xml.StartElement:
		fmt.Fprintf(os.Stderr, "encoding start element %s\n ", t.Name.Local)
	case xml.EndElement:
		fmt.Fprintf(os.Stderr, "encoding end element %s\n ", t.Name.Local)
	case xml.CharData:
		fmt.Fprintf(os.Stderr, "encoding char data element %s\n ", string(t))
	case xml.Comment:
		fmt.Fprintf(os.Stderr, "encoding comment %s\n ", string(t))
	case xml.ProcInst:
		fmt.Fprintf(os.Stderr, "encoding proc inst\n")
	case xml.Directive:
		fmt.Fprintf(os.Stderr, "encoding directive\n")
	default:

	}
	return le.e.EncodeToken(t)
}

func (le *loggingEncoder) Flush() error {
	return le.e.Flush()
}

func (le *loggingEncoder) WriteNewline() {
	le.w.Write([]byte{'\n'})
}

func (le *loggingEncoder) Indent(level int) {
	for i := 0; i < level; i++ {
		le.w.Write([]byte{'\t'})
	}
}

func (le *loggingEncoder) Close() error {
	return le.e.Flush()
}

type newLineEncoder struct {
	inner io.Writer
}

func (w newLineEncoder) Write(data []byte) (n int, err error) {
	n = len(data)
	data = bytes.Replace(data, []byte("&#xA;"), []byte("\n"), -1)
	_, err = w.inner.Write(data)
	return
}

func Render(spec *matter.Spec, doc *ascii.Doc, r io.Reader, w io.Writer, configurator *zap.Configurator, errata *zap.Errata) (err error) {
	d := xml.NewDecoder(r)
	e := xml.NewEncoder(&newLineEncoder{inner: w})
	e.Indent("", "  ")

	//e := &loggingEncoder{w: w, e: en}

	rend := &renderer{
		spec:         spec,
		doc:          doc,
		configurator: configurator,
		errata:       errata,
	}

	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = nil
			return e.Close()
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "configurator":
				err = rend.writeConfigurator(d, e, t, configurator)
			default:
				err = e.EncodeToken(tok)
			}
		default:
			err = e.EncodeToken(tok)
		}

		if err != nil {
			if err == io.EOF {
				err = nil
				return e.Close()
			}
			return
		}
	}
}

func getAttributeValue(attrs []xml.Attr, name string) string {
	for _, a := range attrs {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

func setAttributeValue(attrs []xml.Attr, name string, value string) []xml.Attr {
	for i, a := range attrs {
		if a.Name.Local == name {
			attrs[i] = xml.Attr{Name: a.Name, Value: value}
			return attrs
		}
	}
	return append(attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
}

func removeAttribute(attrs []xml.Attr, name string) []xml.Attr {
	for i, a := range attrs {
		if a.Name.Local == name {
			return slices.Delete(attrs, i, i+1)
		}
	}
	return attrs
}
