package render

import (
	"context"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func Wrap(lineLength int) Option {
	return func(r *Renderer) {
		r.wordWrapLength = lineLength
	}
}

type wrappedTarget struct {
	context.Context

	out      []byte
	lastRune rune

	disableWrapCount    int
	lastNewline         int
	lastInsertedNewline int
	lastRemovedNewline  int
	lastSpace           int
	indented            bool

	blocks        []*strings.Builder
	currentBlock  *strings.Builder
	lastBlockRune rune

	wrapLength int
}

func NewWrappedTarget(parent context.Context, wrapLength int) Target {
	return &wrappedTarget{
		Context:             parent,
		wrapLength:          wrapLength,
		lastSpace:           -1,
		lastInsertedNewline: -1,
		lastRemovedNewline:  -1,
		lastNewline:         -1,
	}
}

func (o *wrappedTarget) WriteString(s string) {
	if len(s) == 0 {
		return
	}
	if o.disableWrapCount > 0 || o.currentBlock != nil {
		if o.disableWrapCount > 0 {
			// We're not wrapping at all right now
			o.out = append(o.out, s...)
			o.lastRune, _ = utf8.DecodeLastRune(o.out)
			return
		}
		if o.currentBlock != nil {
			o.currentBlock.WriteString(s)
			o.lastBlockRune, _ = utf8.DecodeLastRuneInString(s)
			return
		}

	}
	o.writeWrappedText(s)
}

func (o *wrappedTarget) writeWrappedText(s string) {
	for _, r := range s {
		o.writeRune(r)
	}
}

func (o *wrappedTarget) writeRune(r rune) {
	index := len(o.out)
	switch r {
	case '\r': // We just strip these out as we go
		return
	case '\n':
		if o.lastNewline == index-1 { // Double new line; don't remove
			o.lastNewline = index
			o.lastSpace = -1
			o.indented = false
		} else {
			if o.lastRemovedNewline >= 0 && o.lastRemovedNewline == index-1 {
				// Hrm; newline right after a removed newline; restore the old newline
				o.out[o.lastRemovedNewline] = '\n'
				o.lastNewline = o.lastRemovedNewline
				o.lastRemovedNewline = -1
			}
			lineLength := index - o.lastNewline
			shortLine := lineLength > 1 && lineLength < o.wrapLength
			if shortLine && !o.indented && o.lastRune != '+' { // We have a new line, but we're under our wrap length; tentatively remove it
				r = ' '
				o.lastSpace = index
				o.lastRemovedNewline = index
			} else {
				o.lastNewline = index
				o.lastSpace = -1
				o.indented = false
			}
		}
	default:
		if isWhitespace(r) {
			if o.lastRemovedNewline != -1 && o.lastRemovedNewline == index-1 {
				o.indented = true
				o.out[o.lastRemovedNewline] = '\n'
				o.lastNewline = o.lastRemovedNewline
				o.lastRemovedNewline = -1
			}
			o.lastSpace = index
		}
	}

	o.out = utf8.AppendRune(o.out, r)
	o.lastRune = r
	if o.indented {
		return
	}
	if len(o.out)-o.lastNewline <= o.wrapLength { // We're within our wrap length
		return
	}
	if o.lastSpace == -1 || o.lastSpace < o.lastNewline { // We're outside our wrap length, but there's no space to split on
		return
	}
	o.out[o.lastSpace] = '\n'
	if o.lastSpace == len(o.out)-1 {
		o.lastRune = '\n'
	}
	o.lastNewline = o.lastSpace
	o.lastInsertedNewline = o.lastSpace
	o.lastSpace = -1

}

func (o *wrappedTarget) writeBlockText(s string) {
	if o.lastNewline != -1 && (len(o.out)-o.lastNewline)+len(s) > o.wrapLength && o.lastNewline != len(o.out)-1 {
		// We're outside our wrap length, but we can't split the block, so prepend a newline
		o.insertNewLine()
	}
	o.out = append(o.out, s...)
	o.lastRune, _ = utf8.DecodeLastRune(o.out)
	o.updateLastIndexes()
}

func (o *wrappedTarget) insertNewLine() {
	o.lastInsertedNewline = len(o.out)
	o.lastNewline = len(o.out)
	o.out = append(o.out, '\n')
	o.lastRune = '\n'
}

func (o *wrappedTarget) WriteRune(r rune) {
	if o.disableWrapCount > 0 {
		switch r {
		case '\r': // We just strip these out as we go
			return
		case '\n':
			o.lastNewline = len(o.out)
			o.lastSpace = -1
		default:
			if isWhitespace(r) {
				o.lastSpace = len(o.out)
			}
		}
		o.out = utf8.AppendRune(o.out, r)
		o.lastRune = r
		return
	}
	if o.currentBlock != nil {
		o.currentBlock.WriteRune(r)
		o.lastBlockRune = r
		return
	}
	o.writeRune(r)
}

func (o *wrappedTarget) EnsureNewLine() {
	if o.disableWrapCount > 0 || o.currentBlock == nil { // We're writing straight to the stream
		if o.lastRune == '\n' {
			return
		}
		if o.lastRune == ' ' && len(o.out)-1 == o.lastRemovedNewline { // We just removed a newline due to wrapping, so put it back
			o.out[o.lastRemovedNewline] = '\n'
			o.lastNewline = o.lastRemovedNewline
			o.lastRemovedNewline = -1
			return
		}
		o.WriteRune('\n')
		return
	}
	if o.lastBlockRune == '\n' { // The current block just wrote a newline, so skip
		return
	}
	o.currentBlock.WriteRune('\n')
	o.lastBlockRune = '\n'
}

func (o *wrappedTarget) String() string {
	for i := len(o.blocks) - 1; i >= 0; i-- { // Flush any blocks that haven't been written for some reason
		o.writeBlockText(o.blocks[i].String())
	}
	o.blocks = nil
	o.currentBlock = nil
	return unsafe.String(unsafe.SliceData(o.out), len(o.out))
}

func (o *wrappedTarget) FlushWrap() {
	if o.lastRemovedNewline >= 0 && len(o.out) == o.lastRemovedNewline+1 {
		// We ended on a new line and removed it; restore the new line
		o.out[o.lastRemovedNewline] = '\n'
		o.lastRune = '\n'
		o.lastNewline = o.lastRemovedNewline
		o.lastRemovedNewline = -1
	}
}

func (o *wrappedTarget) EnableWrap() {
	o.disableWrapCount -= 1
	if o.disableWrapCount == 0 {
		o.updateLastIndexes()
	}
}

func (o *wrappedTarget) DisableWrap() {
	o.disableWrapCount += 1
}

func (o *wrappedTarget) updateLastIndexes() {
	o.lastNewline = lastIndexRune(o.out, '\n')
	o.lastSpace = lastIndexRune(o.out, ' ')
	if o.lastSpace < o.lastNewline {
		o.lastSpace = -1
	}
	o.indented = false
}

func (o *wrappedTarget) StartBlock() {
	if o.disableWrapCount > 0 { // We're not wrapping right now, so this block doesn't matter
		return
	}
	sb := &strings.Builder{}
	o.blocks = append(o.blocks, sb)
	o.currentBlock = sb
}

func (o *wrappedTarget) EndBlock() {
	if o.disableWrapCount > 0 { // We're not wrapping right now, so this block doesn't matter
		return
	}
	bt := o.currentBlock.String()
	if len(o.blocks) > 1 { // If we have a nested block, write this block to the parent block
		o.blocks = o.blocks[:len(o.blocks)-1]
		o.currentBlock = o.blocks[len(o.blocks)-1]
		if len(bt) > 0 {
			o.currentBlock.WriteString(bt)
			o.lastBlockRune, _ = utf8.DecodeLastRuneInString(bt)
		}
	} else { // Otherwise, just write to the stream
		o.blocks = nil
		o.currentBlock = nil
		o.lastBlockRune = 0
		o.writeBlockText(bt)
	}
}

func (o *wrappedTarget) Subtarget() Target {
	// If we're doing a sub-render of this target, don't bother wrapping; it'll be handled when the parent gets written to
	return NewUnwrappedTarget(o.Context)
}

func lastIndexRune(b []byte, r rune) int {
	for i := len(b); i > 0; {
		cr, size := utf8.DecodeLastRune(b[:i])
		i -= size
		if cr == r {
			return i
		}
	}
	return -1
}

func isWhitespace(r rune) bool {
	switch r {
	case '\t', ' ', 0xA0: // Tabs, spaces and nbsps
		return true
	}
	return false
}
