package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"slices"

	"github.com/beevik/etree"
)

func FormatXML(x string) (string, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(x)
	if err != nil {
		return "", err
	}
	indent := etree.NewIndentSettings()
	indent.Spaces = 2
	indent.PreserveLeafWhitespace = true
	doc.IndentWithSettings(indent)
	return doc.WriteToString()
}

type XmlEncoder interface {
	EncodeToken(t xml.Token) error
	Close() error
}

type XmlTokenSet struct {
	tokens []xml.Token
	index  int
}

func NewXmlTokenSet(tokens []xml.Token) *XmlTokenSet {
	return &XmlTokenSet{tokens: tokens}
}

func (ts *XmlTokenSet) Token() (xml.Token, error) {
	if ts.index >= len(ts.tokens) {
		return nil, io.EOF
	}
	t := ts.tokens[ts.index]
	ts.index++
	return t, nil
}

func (ts *XmlTokenSet) ReadElement(name string) (val string, err error) {
	for {
		var tok xml.Token
		tok, err = ts.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of %s", name)
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return
			default:
				err = fmt.Errorf("unexpected %s end element: %s", name, t.Name.Local)
			}
		case xml.CharData:
			val = string(t)
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

func (ts *XmlTokenSet) Write(e XmlEncoder) (err error) {
	for _, tok := range ts.tokens {
		err = e.EncodeToken(tok)
		if err != nil {
			return
		}
	}
	return
}

func (ts *XmlTokenSet) WriteElement(e XmlEncoder, el xml.StartElement) (err error) {
	name := el.Name.Local
	err = e.EncodeToken(el)
	if err != nil {
		return
	}
	var skipNextCharData bool
	for {
		var tok xml.Token
		tok, err = ts.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of %s", name)
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			err = e.EncodeToken(tok)
			switch t.Name.Local {
			case name:
				return nil
			default:
				skipNextCharData = true
			}
		case xml.CharData:
			if skipNextCharData {
				skipNextCharData = false
				continue
			}
			err = e.EncodeToken(tok)
		default:
			err = e.EncodeToken(tok)
		}
		if err != nil {
			return
		}
	}
}

func (ts *XmlTokenSet) Ignore(name string) (err error) {
	for {
		var tok xml.Token
		tok, err = ts.Token()
		if tok == nil || err == io.EOF {
			panic(fmt.Errorf("EOF before end of %s", name))
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return nil
			default:
			}
		default:
		}
		if err != nil {
			return
		}
	}
}

func (ts *XmlTokenSet) Reset() {
	ts.index = 0
}

func XmlWriteSimpleElement(e XmlEncoder, name string, value string) (err error) {
	elName := xml.Name{Local: name}
	xfs := xml.StartElement{Name: elName}
	err = e.EncodeToken(xfs)
	if err != nil {
		return
	}
	err = e.EncodeToken(xml.CharData(value))
	if err != nil {
		return
	}
	xfe := xml.EndElement{Name: elName}
	err = e.EncodeToken(xfe)
	return
}

type XmlDecoder interface {
	Token() (xml.Token, error)
}

func XmlExtract(d XmlDecoder, el xml.StartElement) (ts *XmlTokenSet, err error) {
	var tokens []xml.Token
	tokens = append(tokens, xml.CopyToken(el))
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of %s", el.Name.Local)
		}
		if err != nil {
			return
		}
		tokens = append(tokens, xml.CopyToken(tok))
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case el.Name.Local:
				ts = NewXmlTokenSet(tokens)
				return
			default:
			}
		default:
		}
		if err != nil {
			return
		}
	}
}

func XmlAttributeGet(attrs []xml.Attr, name string) string {
	for _, a := range attrs {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

func XmlAttributeSet(attrs []xml.Attr, name string, value string) []xml.Attr {
	for i, a := range attrs {
		if a.Name.Local == name {
			attrs[i] = xml.Attr{Name: a.Name, Value: value}
			return attrs
		}
	}
	return append(attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
}

func XmlAttributeRemove(attrs []xml.Attr, name string) []xml.Attr {
	for i, a := range attrs {
		if a.Name.Local == name {
			return slices.Delete(attrs, i, i+1)
		}
	}
	return attrs
}
