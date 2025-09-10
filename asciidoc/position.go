package asciidoc

type position struct {
	document *Document
	line     int
	column   int
	offset   int
}

func (p position) Document() *Document {
	return p.document
}

func (p *position) SetDocument(document *Document) {
	p.document = document
}

func (p position) Origin() (path string, line int) {
	if p.document != nil {
		return p.document.Path.Absolute, p.line
	}
	return "", p.line
}

func (p position) Position() (line int, column int, offset int) {
	line = p.line
	column = p.column
	offset = p.offset
	return
}

func (p *position) SetPosition(line int, column int, offset int) {
	p.line = line
	p.column = column
	p.offset = offset
}

type HasPosition interface {
	Document() *Document
	SetDocument(document *Document)
	Position() (line int, column int, offset int)
	SetPosition(line int, column int, offset int)
}
