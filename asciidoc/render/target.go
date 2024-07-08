package render

type Target interface {
	EnsureNewLine()
	WriteRune(r rune)
	WriteString(s string)
	String() string
	EnableWrap()
	DisableWrap()
	FlushWrap()
	StartBlock()
	EndBlock()
	Subtarget() Target
}
