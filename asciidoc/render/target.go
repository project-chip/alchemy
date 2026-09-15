package render

type TextBuffer interface {
	WriteRune(r rune) (int, error)
	WriteString(s string) (int, error)
	String() string
}

type Target interface {
	TextBuffer

	EnsureNewLine()
	EnableWrap()
	DisableWrap()
	FlushWrap()
	StartBlock()
	EndBlock()
	Subtarget() Target
}
