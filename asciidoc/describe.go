package asciidoc

import (
	"fmt"
	"strings"
)

func Describe(el Element) string {
	switch el := el.(type) {
	case *Section:
		return fmt.Sprintf("Section %s (%d)", DescribeSet(el.Title), el.Level)
	case *String:
		return fmt.Sprintf("String: \"%s\" ", el.Value)
	case *BlockAttributes:
		return fmt.Sprintf("block attributes (%d)", len(el.AttributeList))
	case *AttributeEntry:
		return fmt.Sprintf("Attribute %s", el.Name)
	case *AttributeReset:
		return fmt.Sprintf("attribute reset %s", el.Name)
	case *UserAttributeReference:
		return fmt.Sprintf("Attribute reference %s", el.Name())
	case *CharacterReplacementReference:
		return fmt.Sprintf("character replacement %s", el.Name())
	case *Bold:
		return fmt.Sprintf("bold (%d)", len(el.Elements))
	case *DoubleBold:
		return fmt.Sprintf("double bold (%d)", len(el.Elements))
	case *ThematicBreak:
		return "thematic break"
	case *PageBreak:
		return "page break"
	case *SingleLineComment:
		return fmt.Sprintf("Comment %s ", el.Value)
	case *MultiLineComment:
		return fmt.Sprintf("MultiLine Comment %s ", strings.Join(el.Lines(), "\n"))
	case Email:
		return fmt.Sprintf("email %s ", el.Address)
	case *ExampleBlock:
		return "example"
	//case *Footnote:
	//		return fmt.Sprintf(" Footnote %s (%v)", ae.ID, ae.Value)
	case *Icon:
		return fmt.Sprintf(" icon %s ", el.Path)
	case *BlockImage:
		return fmt.Sprintf(" block image %s ", DescribeSet(el.ImagePath))
	case *InlineImage:
		return fmt.Sprintf(" inline image %s ", DescribeSet(el.ImagePath))
	case *FileInclude:
		return "include"
	case *Italic:
		return fmt.Sprintf("Italic (%d)", len(el.Elements))
	case *DoubleItalic:
		return fmt.Sprintf("double italic (%d)", len(el.Elements))
	case *EmptyLine:
		return "empty line"
	case *Link:
		return fmt.Sprintf("link %s ", Describe(el.URL))
	case *LinkMacro:
		return fmt.Sprintf("link macro %s ", Describe(el.URL))
	case *OrderedListItem:
		return fmt.Sprintf("orderedListItem (%s)", el.Marker)
	case *UnorderedListItem:
		return fmt.Sprintf("unorderedListItem (%s)", el.Marker)
	case *Listing:
		return fmt.Sprintf("listing (%d)", el.Delimiter.Length)
	case *LiteralBlock:
		return "literal"
	case *Marked:
		return fmt.Sprintf("marked (%d)", len(el.Elements))
	case *DoubleMarked:
		return fmt.Sprintf("double marked (%d)", len(el.Elements))
	case *Monospace:
		return fmt.Sprintf("monospace (%d)", len(el.Elements))
	case *DoubleMonospace:
		return fmt.Sprintf("double monospace (%d)", len(el.Elements))
	case *Subscript:
		return fmt.Sprintf("subscript (%d)", len(el.Elements))
	case *Superscript:
		return fmt.Sprintf("superscript (%d)", len(el.Elements))

	case *OpenBlock:
		return "open"
	case *Admonition:
		return fmt.Sprintf("admonition (%s)", Describe(el.AdmonitionType))
	case AdmonitionType:
		switch el {
		case AdmonitionTypeNone:
			return ""
		case AdmonitionTypeNote:
			return "note"
		case AdmonitionTypeTip:
			return "tip"
		case AdmonitionTypeImportant:
			return "important"
		case AdmonitionTypeCaution:
			return "caution"
		case AdmonitionTypeWarning:
			return "warning"
		}
		return ""
	case *Paragraph:
		s := "paragraph "
		if el.Admonition != AdmonitionTypeNone {
			s += "(" + Describe(el.Admonition) + ") "
		}
		return s + fmt.Sprintf("(%d elements)", len(el.Elements))
	case *StemBlock:
		return "stem"
	case *QuoteBlock:
		return "quote"
	case *SidebarBlock:
		return "sidebar"

	case *CrossReference:
		return fmt.Sprintf("xref %s", el.ID)
	case *DocumentCrossReference:
		return fmt.Sprintf("doc-xref %v", el.ReferencePath)
	case *TableCell:
		return fmt.Sprintf("table cell (%d)", len(el.Elements))
	case *TableRow:
		return fmt.Sprintf("table row (%d)", len(el.TableCells()))
	case *Table:
		return fmt.Sprintf("table (%d)", len(el.Elements))
	case SpecialCharacter:
		return fmt.Sprintf("SpecialCharacter %s", el.Character)
	case URL:
		return fmt.Sprintf("URL %s %v", el.Scheme, el.Path)
	case LineContinuation:
		return "line continuation"
	case *NewLine:
		return "new line"
	case *IfDef:
		return fmt.Sprintf("ifdef %s (%d)", el.Attributes.Join(), len(el.Attributes))
	case *IfDefBlock:
		return fmt.Sprintf("ifdef block %s (%d)", el.Attributes.Join(), len(el.Attributes))
	case *IfNDef:
		return fmt.Sprintf("ifndef %s (%d)", el.Attributes.Join(), len(el.Attributes))
	case *IfNDefBlock:
		return fmt.Sprintf("ifndef block %s (%d)", el.Attributes.Join(), len(el.Attributes))
	case *InlineIfDef:
		return fmt.Sprintf("inline ifdef %s", el.Attributes.Join())
	case *InlineIfNDef:
		return fmt.Sprintf("inline ifdef %s", el.Attributes.Join())
	case *EndIf:
		return fmt.Sprintf("endif %s", el.Attributes.Join())
	case *IfEval:
		return fmt.Sprintf("ifeval %v %v %v", el.Left, el.Operator, el.Right)
	case *IfEvalBlock:
		return fmt.Sprintf("ifeval block %v %v %v", el.Left, el.Operator, el.Right)
	case *DescriptionListItem:
		return fmt.Sprintf("description list item (%s - %v)", el.Marker, el.Term)
	case *LineContinuation:
		return "line continuation"
	case *LineBreak:
		return "line break"
	case *ListContinuation:
		return "list continuation"
	case *InlinePassthrough:
		return "inline passthrough"
	case *InlineDoublePassthrough:
		return "inline double passthrough"
	case *Counter:
		s := fmt.Sprintf("counter \"%s\"", el.Name)
		if el.InitialValue != "" {
			s += fmt.Sprintf(" (initial value: %s)", el.InitialValue)
		}
		if el.Display.Visible() {
			s += " (display)"
		}
		return s
	case *Anchor:
		if len(el.Elements) > 0 {
			return fmt.Sprintf("anchor [%s] (%v)", el.ID, el.Elements)
		}
		return fmt.Sprintf("anchor [%s]", el.ID)
	default:
		return fmt.Sprintf("unknown: %T", el)
	}
}

func DescribeSet(el Elements) string {
	var sb strings.Builder
	for _, e := range el {
		sb.WriteString(Describe(e))
	}
	return sb.String()
}
