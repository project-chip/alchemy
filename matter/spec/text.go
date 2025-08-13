package spec

import (
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

var endOfSentencePattern = regexp.MustCompile(`(?m)(\.( |$)|\n\n)`)

func getDescription(doc *Doc, entity types.Entity, parent asciidoc.Parent, els asciidoc.Elements) string {
	var sb strings.Builder
	readDescription(doc, parent, els, &sb)
	description := sb.String()
	endOfSentences := endOfSentencePattern.FindAllStringIndex(description, -1)
	for _, endOfSentence := range endOfSentences {
		endOfSentenceIndex := endOfSentence[0]
		if description[endOfSentenceIndex] == '.' {
			endOfSentenceIndex++
		}
		possible := description[:endOfSentenceIndex]
		if text.HasCaseInsensitiveSuffix(possible, "i.e.") || text.HasCaseInsensitiveSuffix(possible, "e.g.") {
			continue
		}
		description = possible
		break
	}
	if description == "" {
		slog.Warn("Missing description for entity", matter.LogEntity("entity", entity), log.Elements("source", doc.Path, els))
	} else if !strings.HasSuffix(description, ".") {
		slog.Warn("Description for entity is not sentence", matter.LogEntity("entity", entity), slog.String("description", description), log.Elements("source", doc.Path, els))
	}
	return description
}

func readDescription(doc *Doc, parent asciidoc.Parent, els asciidoc.Elements, value *strings.Builder) (err error) {
	var foundNonBlock bool
	for el := range doc.Iterator().Iterate(parent, els) {

		switch el.Type() {
		case asciidoc.ElementTypeBlock, asciidoc.ElementTypeDocument:
			if foundNonBlock {
				return
			}
			continue
		case asciidoc.ElementTypeAttribute, asciidoc.ElementTypeAttributes:
			continue
		}
		foundNonBlock = true
		switch el := el.(type) {
		case *asciidoc.String:
			value.WriteString(el.Value)
		case asciidoc.FormattedTextElement:
			err = readDescription(doc, el, el.Children(), value)
		case *asciidoc.CrossReference:
			if len(el.Elements) > 0 {
				var label strings.Builder
				readDescription(doc, el, el.Children(), &label)
				value.WriteString(strings.TrimSpace(label.String()))
			} else {
				var val string
				anchor := doc.FindAnchorByID(el.ID, el, el)
				if anchor != nil {
					val = matter.StripTypeSuffixes(ReferenceName(anchor.Document, anchor.Element))
				} else {
					val = doc.anchorId(doc.Iterator(), el, el, el.ID)
					val = strings.TrimPrefix(val, "_")
					val = strings.TrimPrefix(val, "ref_") // Trim, and hope someone else has it defined
				}
				value.WriteString(val)
			}
		case *asciidoc.Link:
			var textAttribute, titleAttribute, altTextAttribute asciidoc.Attribute
			for _, a := range el.AttributeList.Attributes() {
				switch a.AttributeType() {
				case asciidoc.AttributeTypeTitle:
					titleAttribute = a
				case asciidoc.AttributeTypeAlternateText:
					altTextAttribute = a
				}
			}
			if altTextAttribute != nil {
				textAttribute = altTextAttribute
			} else if titleAttribute != nil {
				textAttribute = titleAttribute
			}
			if textAttribute != nil {
				switch val := textAttribute.Value().(type) {
				case asciidoc.Elements:
					readDescription(doc, &val, val, value)
				default:
					slog.Warn("Unexpected value type when reading entity description", log.Type("valueType", val), log.Path("source", el))
				}
			} else {
				value.WriteString(el.URL.Scheme)
				readDescription(doc, &el.URL.Path, el.URL.Path, value)
			}
		case *asciidoc.LinkMacro:
			value.WriteString(el.URL.Scheme)
			readDescription(doc, &el.URL.Path, el.URL.Path, value)
		case *asciidoc.Superscript:
			// In the special case of superscript elements, we do checks to make sure it's not an asterisk or a footnote, which should be ignored
			var quotedText strings.Builder
			err = readDescription(doc, el, el.Children(), &quotedText)
			if err != nil {
				return
			}
			qt := quotedText.String()
			if qt == "*" { //
				continue
			}
			_, parseErr := strconv.Atoi(qt)
			if parseErr == nil {
				// This is probably a footnote
				// The similar buildConstraintValue method does not do this, as there are exponential values in contraints
				continue
			}
			value.WriteString(qt)
		case *asciidoc.SpecialCharacter:
			value.WriteString(el.Character)
		case *asciidoc.InlinePassthrough:
			value.WriteString("+")
			err = readDescription(doc, el, el.Children(), value)
		case *asciidoc.InlineDoublePassthrough:
			value.WriteString("++")
			err = readDescription(doc, el, el.Children(), value)
		case *asciidoc.ThematicBreak:
		case *asciidoc.EmptyLine:
		case *asciidoc.NewLine:
			value.WriteString(" ")
		case asciidoc.ParentElement:
			err = readDescription(doc, el, el.Children(), value)
		case *asciidoc.LineBreak:
			value.WriteString(" ")
		default:
			return fmt.Errorf("unexpected type in description: %T", el)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func CanonicalName(name string) string {
	if !strings.ContainsRune(name, ' ') {
		return name
	}
	casedName := matter.Case(name)
	if casedName != name {
		slog.Debug("Canonlicalizing name", slog.String("from", name), slog.String("to", casedName))
		return casedName
	}
	return name
}
