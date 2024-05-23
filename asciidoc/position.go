package asciidoc

type position struct {
	line   int
	column int
	offset int
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
	Position() (line int, column int, offset int)
	SetPosition(line int, column int, offset int)
}
