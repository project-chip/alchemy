package parse

func initialPosition(line, col, offset int) Option {
	return func(p *parser) Option {
		p.offset.line = line - 1
		p.offset.col = col
		p.offset.offset = offset
		return initialPosition(0, 0, 0)
	}
}
