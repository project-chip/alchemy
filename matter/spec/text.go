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

func getDescription(doc *Doc, entity types.Entity, els asciidoc.Set) string {
	var sb strings.Builder
	readDescription(doc, els, &sb)
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

func readDescription(doc *Doc, els asciidoc.Set, value *strings.Builder) (err error) {
	var foundNonBlock bool
	for _, el := range els {
		var e asciidoc.Element
		switch el := el.(type) {
		case *Element:
			e = el.Base
		case asciidoc.Element:
			e = el
		default:
			return
		}

		switch e.Type() {
		case asciidoc.ElementTypeBlock, asciidoc.ElementTypeDocument:
			if foundNonBlock {
				return
			}
			continue
		case asciidoc.ElementTypeAttribute, asciidoc.ElementTypeAttributes:
			continue
		}
		foundNonBlock = true
		switch el := e.(type) {
		case *asciidoc.String:
			value.WriteString(el.Value)
		case asciidoc.FormattedTextElement:
			err = readDescription(doc, el.Elements(), value)
		case *asciidoc.CrossReference:
			if len(el.Set) > 0 {
				var label strings.Builder
				readDescription(doc, el.Set, &label)
				value.WriteString(strings.TrimSpace(label.String()))
			} else {
				var val string
				anchor := doc.FindAnchor(el.ID, el)
				if anchor != nil {
					val = matter.StripTypeSuffixes(ReferenceName(anchor.Element))
				} else {
					val = strings.TrimPrefix(el.ID, "_")
					val = strings.TrimPrefix(val, "ref_") // Trim, and hope someone else has it defined
				}
				value.WriteString(val)
			}
		case *asciidoc.Link:
			value.WriteString(el.URL.Scheme)
			readDescription(doc, el.URL.Path, value)
		case *asciidoc.LinkMacro:
			value.WriteString(el.URL.Scheme)
			readDescription(doc, el.URL.Path, value)
		case *asciidoc.Superscript:
			// In the special case of superscript elements, we do checks to make sure it's not an asterisk or a footnote, which should be ignored
			var quotedText strings.Builder
			err = readDescription(doc, el.Elements(), &quotedText)
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
			err = readDescription(doc, el.Elements(), value)
		case *asciidoc.InlineDoublePassthrough:
			value.WriteString("++")
			err = readDescription(doc, el.Elements(), value)
		case *asciidoc.ThematicBreak:
		case asciidoc.EmptyLine:
		case *asciidoc.NewLine:
			value.WriteString(" ")
		case asciidoc.HasElements:
			err = readDescription(doc, el.Elements(), value)
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
