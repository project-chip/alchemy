package render

import (
	"fmt"
	"log/slog"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

type AttributeFilter uint32

const (
	AttributeFilterNone AttributeFilter = 0
	AttributeFilterAll  AttributeFilter = math.MaxUint32
)

const (
	AttributeFilterID AttributeFilter = 1 << iota
	AttributeFilterTitle
	AttributeFilterStyle
	AttributeFilterCols
	AttributeFilterText
	AttributeFilterAlt
	AttributeFilterHeight
	AttributeFilterWidth
	AttributeFilterPDFWidth
)

func shouldRenderAttributeType(at AttributeFilter, include AttributeFilter, exclude AttributeFilter) bool {
	return ((at & include) == at) && ((at & exclude) != at)
}

func renderAttributes(cxt *output.Context, el interface{}, attributes types.Attributes) error {
	return renderSelectAttributes(cxt, el, attributes, AttributeFilterAll, AttributeFilterCols)
}

/*
AttrDocType = "doctype"
	// AttrDocType the "description" attribute
	AttrDescription = "description"
	// AttrSyntaxHighlighter the attribute to define the syntax highlighter on code source blocks
	AttrSyntaxHighlighter = "source-highlighter"
	// AttrChromaClassPrefix the class prefix used by Chroma when rendering source code (default: `tok-`)
	AttrChromaClassPrefix = "chroma-class-prefix"
	// AttrID the key to retrieve the ID
	AttrID = "id"
	// AttrIDPrefix the key to retrieve the ID Prefix
	AttrIDPrefix = "idprefix"
	// DefaultIDPrefix the default ID Prefix
	DefaultIDPrefix = "_"
	// AttrIDSeparator the key to retrieve the ID Separator
	AttrIDSeparator = "idseparator"
	// DefaultIDSeparator the default ID Separator
	DefaultIDSeparator = "_"
	// AttrNumbered the `numbered` attribute to trigger section numbering at renderding time
	AttrNumbered = "numbered"
	// AttrSectionNumbers the `sectnums` attribute to trigger section numbering at renderding time (an alias for `numbered`)
	AttrSectionNumbering = "sectnums"
	// AttrTableOfContents the `toc` attribute at document level
	AttrTableOfContents = "toc"
	// AttrTableOfContentsLevels the document attribute which specifies the number of levels to display in the ToC
	AttrTableOfContentsLevels = "toclevels"
	// AttrTableOfContentsTitle the document attribute which specifies the title of the table of contents
	AttrTableOfContentsTitle = "toc-title"
	// AttrNoHeader attribute to disable the rendering of document footer
	AttrNoHeader = "noheader"
	// AttrNoFooter attribute to disable the rendering of document footer
	AttrNoFooter = "nofooter"
	// AttrCustomID the key to retrieve the flag that indicates if the element ID is custom or generated
	// AttrCustomID = "@customID"
	// AttrTitle the key to retrieve the title
	AttrTitle = "title"
	// AttrAuthors the key to the authors declared after the section level 0 (at the beginning of the doc)
	AttrAuthors = "authors"
	// AttrAuthor the key to the author's full name declared as a standalone attribute
	AttrAuthor = "author"
	// AttrAuthor the key to the author's email address declared as a standalone attribute
	AttrEmail = "email"
	// AttrRevision the key to the revision declared after the section level 0 (at the beginning of the doc)
	// or as a standalone attribute
	AttrRevision = "revision"
	// AttrRole the key for a single role attribute
	AttrRole = "role"
	// AttrRoles the key to retrieve the roles attribute
	AttrRoles = "roles"
	// AttrOption the key for a single option attribute
	AttrOption = "option"
	// AttrOptions the key to retrieve the options attribute
	AttrOptions = "options"
	// AttrOpts alias for AttrOptions
	AttrOpts = "opts"
	// AttrInlineLink the key to retrieve the link
	AttrInlineLink = "link"
	// AttrQuoteAuthor attribute for the author of a verse
	AttrQuoteAuthor = "quoteAuthor"
	// AttrQuoteTitle attribute for the title of a verse
	AttrQuoteTitle = "quoteTitle"
	// AttrSource the `source` attribute for a source block or a source paragraph (this is a placeholder, ie, it does not expect any value for this attribute)
	AttrSource = "source"
	// AttrLanguage the `language` attribute for a source block or a source paragraph
	AttrLanguage = "language"
	// AttrLineNums the `linenums` attribute for a source block or a source paragraph
	AttrLineNums = "linenums"
	// AttrCheckStyle the attribute to mark the first element of an unordered list item as a checked or not
	AttrCheckStyle = "checkstyle"
	// AttrInteractive the attribute to mark the first element of an unordered list item as n interactive checkbox or not
	// (paired with `AttrCheckStyle`)
	AttrInteractive = "interactive"
	// AttrStart the `start` attribute in an ordered list
	AttrStart = "start"
	// AttrLevelOffset the `leveloffset` attribute used in file inclusions
	AttrLevelOffset = "leveloffset"
	// AttrLineRanges the `lines` attribute used in file inclusions
	AttrLineRanges = "lines"
	// AttrTagRanges the `tag`/`tags` attribute used in file inclusions
	AttrTagRanges = "tags"
	// AttrLastUpdated the "last updated" data in the document, i.e., the output/generation time
	AttrLastUpdated = "LastUpdated"
	// AttrImageAlt the image `alt` attribute
	AttrImageAlt = "alt"
	// AttrHeight the image `height` attribute
	AttrHeight = "height"
	// AttrImageWindow the `window` attribute, which becomes the target for the link
	AttrImageWindow = "window"
	// AttrImageAlign is for image alignment
	AttrImageAlign = "align"
	// AttrIconSize the icon `size`, and can be one of 1x, 2x, 3x, 4x, 5x, lg, fw
	AttrIconSize = "size"
	// AttrIconRotate the icon `rotate` attribute, and can be one of 90, 180, or 270
	AttrIconRotate = "rotate"
	// AttrIconFlip the icon `flip` attribute, and if set can be "horizontal" or "vertical"
	AttrIconFlip = "flip"
	// AttrUnicode local libasciidoc attribute to encode output as UTF-8 instead of ASCII.
	AttrUnicode = "unicode"
	// AttrCaption is the caption for block images, tables, and so forth
	AttrCaption = "caption"
	// AttrStyle paragraph, block or list style
	AttrStyle = "style"
	// AttrInlineLinkText the text attribute (first positional) of links
	AttrInlineLinkText = "text"
	// AttrInlineLinkTarget the 'window' attribute
	AttrInlineLinkTarget = "window"
	// AttrWidth the `width` attribute used ior images, tables, and so forth
	AttrWidth = "width"
	// AttrFrame the frame used mostly for tables (all, topbot, sides, none)
	AttrFrame = "frame"
	// AttrGrid the grid (none, all, cols, rows) in tables
	AttrGrid = "grid"
	// AttrStripes controls table row background (even, odd, all, none, hover)
	AttrStripes = "stripes"
	// AttrFloat is for image or table float (text flows around)
	AttrFloat = "float"
	// AttrCols the table columns attribute
	AttrCols = "cols"
	// AttrAutoWidth the `autowidth` attribute on a table
	AttrAutoWidth = "autowidth"
	// AttrPositionalIndex positional parameter index
	AttrPositionalIndex = "@positional-"
	// AttrPositional1 positional parameter 1
	AttrPositional1 = "@positional-1"
	// AttrPositional2 positional parameter 2
	AttrPositional2 = "@positional-2"
	// AttrPositional3 positional parameter 3
	AttrPositional3 = "@positional-3"
	// AttrVersionLabel labels the version number in the document
	AttrVersionLabel = "version-label"
	// AttrExampleCaption is the example caption
	AttrExampleCaption = "example-caption"
	// AttrFigureCaption is the figure (image) caption
	AttrFigureCaption = "figure-caption"
	// AttrTableCaption is the table caption
	AttrTableCaption = "table-caption"
	// AttrCautionCaption is the CAUTION caption
	AttrCautionCaption = "caution-caption"
	// AttrImportantCaption is the IMPORTANT caption
	AttrImportantCaption = "important-caption"
	// AttrNoteCaption is the NOTE caption
	AttrNoteCaption = "note-caption"
	// AttrTipCaption is the TIP caption
	AttrTipCaption = "tip-caption"
	// AttrWarningCaption is the TIP caption
	AttrWarningCaption = "warning-caption"
	// AttrSubstitutions the "subs" attribute to configure substitutions on delimited blocks and paragraphs
	AttrSubstitutions = "subs"
	// AttrImagesDir the `imagesdir` attribute
	AttrImagesDir = "imagesdir"
	// AttrXRefLabel the label of a cross reference
	AttrXRefLabel = "xrefLabel"
	// AttrExperimental a flag to enable experiment macros (for UI)
	AttrExperimental = "experimental"
	// AttrButtonLabel the label of a button
	AttrButtonLabel = "label"
	// AttrHardBreaks the attribute to set on a paragraph to render with hard breaks on each line
	AttrHardBreaks = "hardbreaks"
*/

func renderSelectAttributes(cxt *output.Context, el interface{}, attributes types.Attributes, include AttributeFilter, exclude AttributeFilter) (err error) {
	if len(attributes) == 0 {
		return
	}

	var id string
	var title string
	var style string
	var keys []string
	for key, val := range attributes {
		switch key {
		case types.AttrID:
			id = val.(string)
		case types.AttrStyle:
			style = val.(string)
		case types.AttrTitle:
			switch v := val.(type) {
			case string:
				title = v
			case []interface{}:
				renderContext := output.NewContext(cxt, cxt.Doc)
				RenderElements(renderContext, "", v)
				title = renderContext.String()
			default:
				err = fmt.Errorf("unknown title type: %T", v)
				return
			}
		default:
			keys = append(keys, key)
		}
	}

	if len(style) > 0 && shouldRenderAttributeType(AttributeFilterStyle, include, exclude) {
		switch style {
		case types.Tip, types.Note, types.Important, types.Warning, types.Caution:
			switch el.(type) {
			case *types.Paragraph:
				cxt.WriteString(fmt.Sprintf("%s: ", style))
			default:
				cxt.WriteString(fmt.Sprintf("[%s]\n", style))
			}
		case "none":
			cxt.WriteString("[none]\n")
		case types.UpperRoman, types.LowerRoman, types.Arabic, types.UpperAlpha, types.LowerAlpha:
			cxt.WriteRune('[')
			cxt.WriteString(style)
			cxt.WriteString("]\n")
		case "a2s", "actdiag", "plantuml", "qrcode", "blockdiag", "d2", "lilypond":
			renderDiagramAttributes(cxt, style, id, keys, attributes)
			return
		case "literal_paragraph":
		default:
			err = fmt.Errorf("unknown style: %s", style)
			return
		}
	}
	if len(title) > 0 && shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
		cxt.WriteNewline()
		cxt.WriteRune('.')
		cxt.WriteString(title)
		cxt.WriteNewline()
	}
	if len(id) > 0 && id[0] != '_' && shouldRenderAttributeType(AttributeFilterID, include, exclude) {
		cxt.WriteNewline()
		cxt.WriteString("[[")
		cxt.WriteString(id)
		cxt.WriteString("]]")
		cxt.WriteRune('\n')
	}
	if len(keys) > 0 {
		sort.Strings(keys)
		switch el.(type) {
		case *types.ImageBlock, *types.InlineLink, *types.InlineImage:
		default:
			cxt.WriteNewline()
		}

		count := 0
		for _, key := range keys {
			var attributeType AttributeFilter
			var skipKey = false
			switch key {
			case types.AttrCols:
				attributeType = AttributeFilterCols
			case types.AttrInlineLinkText:
				attributeType = AttributeFilterText
			case types.AttrImageAlt:
				attributeType = AttributeFilterAlt
				skipKey = true
			case types.AttrHeight:
				attributeType = AttributeFilterHeight
			case types.AttrWidth:
				attributeType = AttributeFilterWidth
			case "pdfwidth":
				attributeType = AttributeFilterPDFWidth
			}
			if !shouldRenderAttributeType(AttributeFilterAlt, include, exclude) {
				continue
			}
			val := attributes[key]
			var keyVal string

			switch attributeType {
			case AttributeFilterText:
				if s, ok := val.(string); ok {
					keyVal = s
					skipKey = true
				}
			case AttributeFilterAlt:
				if s, ok := val.(string); ok {
					keyVal = s
				}
			default:
				switch v := val.(type) {
				case string:
					keyVal = v

				case types.Options:
					for _, o := range v {
						switch opt := o.(type) {
						case string:
							keyVal = opt
						default:
							slog.Debug("unknown attribute option", "type", o)
						}
					}
				case []interface{}:

					var columns []string
					for i, e := range v {
						switch tc := e.(type) {
						case *types.TableColumn:
							var val strings.Builder
							if tc.Multiplier > 1 {
								val.WriteString(strconv.Itoa(tc.Multiplier))
								val.WriteRune('*')
							}
							if tc.HAlign != types.HAlignDefault {
								val.WriteString(string(tc.HAlign))
							}
							if tc.VAlign != types.VAlignDefault {
								val.WriteString(string(tc.VAlign))
							}
							if tc.Autowidth {
								val.WriteRune('~')
							} else if tc.Weight > 1 {
								val.WriteString(strconv.Itoa(tc.Weight))
							}
							if len(tc.Style) > 0 {
								val.WriteString(string(tc.Style))
							}
							columns = append(columns, val.String())
							if i == len(v)-1 && val.Len() == 0 {
								// The parser looks for tokens ending with commas, but these values
								// are actually joined with commas; if the last value is blank,
								// the parser will report one fewer column def, so we add it back
								columns = append(columns, "")
							}
						default:
							err = fmt.Errorf("unknown attribute: %T", e)
							return
						}
					}
					keyVal = strings.Join(columns, ",")
				default:
					err = fmt.Errorf("unknown attribute type: %T", val)
					return
				}
			}

			if len(keyVal) != 0 {
				if count == 0 {
					cxt.WriteString("[")
				} else {
					cxt.WriteRune(',')
				}
				if skipKey {
					cxt.WriteString(keyVal)
				} else {
					cxt.WriteString(key)
					cxt.WriteRune('=')
					if _, err := strconv.Atoi(strings.TrimSuffix(keyVal, "%")); err == nil {
						cxt.WriteString(keyVal)
					} else {
						cxt.WriteRune('"')
						cxt.WriteString(keyVal)
						cxt.WriteRune('"')
					}
				}

				count++
			}

		}
		if count > 0 {
			cxt.WriteRune(']')
			cxt.WriteRune('\n')
		}
	}
	return
}

func renderDiagramAttributes(cxt *output.Context, style string, id string, keys []string, attributes types.Attributes) {
	cxt.WriteString("[")
	cxt.WriteString(style)
	if len(id) > 0 {
		cxt.WriteString(", id=\"")
		cxt.WriteString(id)
		cxt.WriteRune('"')
	}
	for _, k := range keys {
		v, ok := attributes[k]
		if !ok {
			continue
		}
		cxt.WriteString(", ")
		cxt.WriteString(k)
		s, ok := v.(string)
		if ok && len(s) > 0 {
			cxt.WriteString(`="`)
			cxt.WriteString(s)
			cxt.WriteRune('"')
		}
	}
	cxt.WriteRune(']')
	cxt.WriteRune('\n')
}

func renderAttributeDeclaration(cxt *output.Context, ad *types.AttributeDeclaration) (err error) {
	switch ad.Name {
	case "authors":
		if authors, ok := ad.Value.(types.DocumentAuthors); ok {
			for _, author := range authors {
				if len(author.Email) > 0 {
					cxt.WriteString(author.Email)
					cxt.WriteString(" ")
				}
				if author.DocumentAuthorFullName != nil {
					cxt.WriteString(author.DocumentAuthorFullName.FullName())
				}
				cxt.WriteRune('\n')
			}
		}
	default:
		cxt.WriteRune(':')
		cxt.WriteString(ad.Name)
		cxt.WriteString(":")
		switch val := ad.Value.(type) {
		case string:
			cxt.WriteRune(' ')
			cxt.WriteString(val)
		case *types.Paragraph:
			var previous interface{}
			err = renderParagraph(cxt, val, &previous)
		case nil:
		default:
			err = fmt.Errorf("unknown attribute declaration value type: %T", ad.Value)
		}
		cxt.WriteRune('\n')
	}
	return
}
