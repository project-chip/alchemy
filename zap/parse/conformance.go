package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"

	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/shopspring/decimal"
)

func isConformanceElement(e xml.StartElement) bool {
	switch e.Name.Local {
	case "optionalConform",
		"mandatoryConform",
		"disallowConform",
		"provisionalConform",
		"deprecateConform",
		"describedConform",
		"otherwiseConform":
		return true
	default:
		return false
	}
}

func parseConformance(d *xml.Decoder, e xml.StartElement) (c conformance.Conformance, err error) {

	switch e.Name.Local {
	case "optionalConform":
		c, err = parseOptionalConformance(d, e)
	case "mandatoryConform":
		c, err = parseMandatoryConformance(d, e)
	case "disallowConform":
		err = parseStandalone(d, e, e.Name.Local)
		c = &conformance.Disallowed{}
	case "provisionalConform":
		err = parseStandalone(d, e, e.Name.Local)
		c = &conformance.Disallowed{}
	case "deprecateConform":
		err = parseStandalone(d, e, e.Name.Local)
		c = &conformance.Deprecated{}
	case "describedConform":
		err = parseStandalone(d, e, e.Name.Local)
		c = &conformance.Described{}
	case "otherwiseConform":
		c, err = parseOtherwiseConform(d, e)
	default:
		err = fmt.Errorf("unexpected conformance element: %s", e.Name.Local)
	}
	if err != nil {
		return
	}

	return
}

func parseOptionalConformance(d *xml.Decoder, e xml.StartElement) (c *conformance.Optional, err error) {
	var exp conformance.Expression
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if isConformanceExpressionElement(t) {
				exp, err = parseConformanceExpression(d, t)
			} else if isIdentifierExpressionElement(t) {
				exp, err = parseIdentifierExpression(d, t)
			} else {
				err = fmt.Errorf("unexpected optionalConform start element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "optionalConform":
				c = &conformance.Optional{Expression: exp}
				return
			default:
				err = fmt.Errorf("unexpected optionalConform end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected optionalConform level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func parseMandatoryConformance(d *xml.Decoder, e xml.StartElement) (c *conformance.Mandatory, err error) {
	var exp conformance.Expression
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if isConformanceExpressionElement(t) {
				exp, err = parseConformanceExpression(d, t)
			} else if isIdentifierExpressionElement(t) {
				exp, err = parseIdentifierExpression(d, t)
			} else {
				err = fmt.Errorf("unexpected mandatoryConform start element: %s", t.Name.Local)
			}

		case xml.EndElement:
			switch t.Name.Local {
			case "mandatoryConform":
				c = &conformance.Mandatory{Expression: exp}
				return
			default:
				err = fmt.Errorf("unexpected mandatoryConform end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected mandatoryConform level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func parseStandalone(d *xml.Decoder, e xml.StartElement, name string) (err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:

			err = fmt.Errorf("unexpected standalone %s start element: %s", name, t.Name.Local)

		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return
			default:
				err = fmt.Errorf("unexpected %s end element: %s", name, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

func isConformanceExpressionElement(e xml.StartElement) bool {
	if isLogicalExpressionElement(e) {
		return true
	}
	if isIdentifierExpressionElement(e) {
		return true
	}
	if isNegativeExpressionElement(e) {
		return true
	}
	if isComparisonExpressionElement(e) {
		return true
	}
	return false
}

func parseOtherwiseConform(d *xml.Decoder, e xml.StartElement) (set conformance.Set, err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if isConformanceElement(t) {
				var cs conformance.Conformance
				cs, err = parseConformance(d, t)
				if err == nil {
					set = append(set, cs)
				}
			} else {
				err = fmt.Errorf("unexpected otherwiseConform start element: %s", t.Name.Local)
			}

		case xml.EndElement:
			switch t.Name.Local {
			case "otherwiseConform":
				return
			default:
				err = fmt.Errorf("unexpected otherwiseConform end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected otherwiseConform level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func parseConformanceExpression(d *xml.Decoder, e xml.StartElement) (exp conformance.Expression, err error) {
	if isLogicalExpressionElement(e) {
		exp, err = parseLogicalExpression(d, e)
	} else if isIdentifierExpressionElement(e) {
		exp, err = parseIdentifierExpression(d, e)
	} else if isNegativeExpressionElement(e) {
		exp, err = parseNotExpression(d, e)
	} else if isComparisonExpressionElement(e) {
		exp, err = parseComparisonExpression(d, e)
	} else {
		err = fmt.Errorf("unexpected conformance expression element: %s", e.Name.Local)
	}
	return
	/*if err != nil {
		return
	}

	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:


		case xml.EndElement:
			switch t.Name.Local {
			case e.Name.Local:
				return
			default:
				err = fmt.Errorf("unexpected conformance %s end element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected conformance %s type: %T", e.Name.Local, t)
		}
		if err != nil {
			return
		}
	}*/
}

func isNegativeExpressionElement(e xml.StartElement) bool {
	switch e.Name.Local {
	case "notTerm":
		return true
	default:
		return false
	}
}

func parseNotExpression(d *xml.Decoder, e xml.StartElement) (exp conformance.Expression, err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if isConformanceExpressionElement(t) {
				exp, err = parseConformanceExpression(d, t)
			} else {
				err = fmt.Errorf("unexpected notTerm %s start element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case e.Name.Local:
				switch exp := exp.(type) {
				case *conformance.LogicalExpression:
					exp.Not = true
				case *conformance.IdentifierExpression:
					exp.Not = true
				case nil:
					l, c := d.InputPos()
					slog.Info("notTerm missing expression", "name", t.Name.Local)

					err = fmt.Errorf("notTerm missing expression: %d:%d", l, c)
				default:
					err = fmt.Errorf("unexpected notTerm expression type: %T", exp)
				}
				return
			default:
				err = fmt.Errorf("unexpected notTerm %s end element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected notTerm %s level type: %T", e.Name.Local, t)
		}
		if err != nil {
			return
		}
	}
}

func parseLogicalExpression(d *xml.Decoder, e xml.StartElement) (exp *conformance.LogicalExpression, err error) {
	exp = &conformance.LogicalExpression{}
	var components []conformance.Expression
	switch e.Name.Local {
	case "andTerm":
		exp.Operand = "&"
	case "orTerm":
		exp.Operand = "|"
	case "xorTerm":
		exp.Operand = "^"
	default:
		err = fmt.Errorf("unexpected logical expression element: %s", e.Name.Local)
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if isConformanceExpressionElement(t) {
				var ie conformance.Expression
				ie, err = parseConformanceExpression(d, t)
				if err == nil {
					components = append(components, ie)
				}
			} else {
				err = fmt.Errorf("unexpected logical expression %s start element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case e.Name.Local:
				if len(components) < 2 {
					err = fmt.Errorf("not enough components for %s", e.Name.Local)
					return
				}
				exp.Left = components[0]
				exp.Right = components[1:]
				return
			default:
				err = fmt.Errorf("unexpected logical expression %s end element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected logical expression %s type: %T", e.Name.Local, t)
		}
		if err != nil {
			return
		}
	}
}

func isLogicalExpressionElement(e xml.StartElement) bool {
	switch e.Name.Local {
	case "andTerm",
		"orTerm",
		"xorTerm":
		return true
	default:
		return false
	}
}

func isIdentifierExpressionElement(e xml.StartElement) bool {
	switch e.Name.Local {
	case "feature",
		"command",
		"event",
		"condition",
		"typeDef",
		"enum",
		"bitmap",
		"value",
		"attribute":
		return true
	default:
		return false
	}
}

func parseIdentifierExpression(d *xml.Decoder, e xml.StartElement) (exp conformance.Expression, err error) {
	var name, field string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			name = a.Value
		case "field":
			field = a.Value
		default:
			err = fmt.Errorf("unexpected identifier %s attribute: %s", e.Name.Local, a.Name.Local)
			return
		}
	}
	switch e.Name.Local {
	case "feature":
		exp = &conformance.IdentifierExpression{ID: name}
	case "command":
		exp = &conformance.IdentifierExpression{ID: name}
	case "attribute":
		exp = &conformance.IdentifierExpression{ID: name}
	case "event":
		exp = &conformance.IdentifierExpression{ID: name}
	case "condition":
		exp = &conformance.IdentifierExpression{ID: name}
	case "typeDef":
		exp = &conformance.IdentifierExpression{ID: name}
	default:
		err = fmt.Errorf("unexpected identifier element: %s.%s", e.Name.Local, field)
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			err = fmt.Errorf("unexpected identifier start element: %s", t.Name.Local)
		case xml.EndElement:
			switch t.Name.Local {
			case e.Name.Local:
				return
			default:
				err = fmt.Errorf("unexpected identifier %s end element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected identifier %s level type: %T", e.Name.Local, t)
		}
		if err != nil {
			return
		}
	}
}

func isComparisonExpressionElement(e xml.StartElement) bool {
	switch e.Name.Local {
	case "equalTerm",
		"notEqualTerm",
		"greaterTerm",
		"greaterOrEqualTerm",
		"lessTerm",
		"lessOrEqualTerm":
		return true
	default:
		return false
	}
}

func parseComparisonExpression(d *xml.Decoder, e xml.StartElement) (exp *conformance.ComparisonExpression, err error) {
	exp = &conformance.ComparisonExpression{}
	switch e.Name.Local {
	case "equalTerm":
		exp.Op = conformance.ComparisonOperatorEqual
	case "notEqualTerm":
		exp.Op = conformance.ComparisonOperatorNotEqual
	case "greaterTerm":
		exp.Op = conformance.ComparisonOperatorGreaterThan
	case "greaterOrEqualTerm":
		exp.Op = conformance.ComparisonOperatorGreaterThanOrEqual
	case "lessTerm":
		exp.Op = conformance.ComparisonOperatorLessThan
	case "lessOrEqualTerm":
		exp.Op = conformance.ComparisonOperatorLessThanOrEqual
	default:
		err = fmt.Errorf("unexpected comparison expression element: %s", e.Name.Local)
		return
	}
	var comparisons []conformance.ComparisonValue
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "literal":
				var ie conformance.ComparisonValue
				ie, err = parseLiteral(d, t)
				if err == nil {
					comparisons = append(comparisons, ie)
				}
			default:
				if isIdentifierExpressionElement(t) {
					var ie conformance.ComparisonValue
					ie, err = parseIdentifierValue(d, t)
					if err == nil {
						comparisons = append(comparisons, ie)
					}

				} else {
					err = fmt.Errorf("unexpected comparison start element: %s", t.Name.Local)
				}
			}
		case xml.EndElement:
			switch t.Name.Local {
			case e.Name.Local:
				if len(comparisons) != 2 {
					err = fmt.Errorf("incorrect number of comparisons for %s", e)
				} else {
					exp.Left = comparisons[0]
					exp.Right = comparisons[1]
				}
				return
			default:
				err = fmt.Errorf("unexpected comparison %s end element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected comparison %s level type: %T", e.Name.Local, t)
		}
		if err != nil {
			return
		}
	}
}

func parseLiteral(d *xml.Decoder, e xml.StartElement) (exp conformance.ComparisonValue, err error) {
	var value string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "value":
			value = a.Value
		default:
			err = fmt.Errorf("unexpected literal value %s attribute: %s", e.Name.Local, a.Name.Local)
			return
		}
	}
	switch value {
	case "true", "false":
		exp = &conformance.BooleanValue{Boolean: value == "true"}
	default:
		var n decimal.Decimal
		n, err = decimal.NewFromString(value)
		if err == nil {
			if n.IsInteger() {
				exp = &conformance.IntValue{Int: n.IntPart()}
			} else {
				exp = &conformance.FloatValue{Float: n}
			}
		}
	}
	if err != nil {
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of literal")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:

			err = fmt.Errorf("unexpected literal start element: %s", t.Name.Local)

		case xml.EndElement:
			switch t.Name.Local {
			case e.Name.Local:

				return
			default:
				err = fmt.Errorf("unexpected literal %s end element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected literal %s level type: %T", e.Name.Local, t)
		}
		if err != nil {
			return
		}
	}
}

func parseIdentifierValue(d *xml.Decoder, e xml.StartElement) (exp *conformance.IdentifierValue, err error) {
	var name, field string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			name = a.Value
		case "field":
			field = a.Value
		default:
			err = fmt.Errorf("unexpected identifier value %s attribute: %s", e.Name.Local, a.Name.Local)
			return
		}
	}
	switch e.Name.Local {
	case "feature":
		exp = &conformance.IdentifierValue{ID: name}
	case "command":
		exp = &conformance.IdentifierValue{ID: name}
	case "attribute":
		exp = &conformance.IdentifierValue{ID: name}
	case "event":
		exp = &conformance.IdentifierValue{ID: name}
	case "condition":
		exp = &conformance.IdentifierValue{ID: name}
	case "typeDef":
		exp = &conformance.IdentifierValue{ID: name}
	default:
		err = fmt.Errorf("unexpected identifier value element: %s.%s", e.Name.Local, field)
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			err = fmt.Errorf("unexpected identifier value start element: %s", t.Name.Local)
		case xml.EndElement:
			switch t.Name.Local {
			case e.Name.Local:
				return
			default:
				err = fmt.Errorf("unexpected identifier value %s end element: %s", e.Name.Local, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected identifier value %s level type: %T", e.Name.Local, t)
		}
		if err != nil {
			return
		}
	}
}

/*

func renderConformanceEntity(parent *etree.Element, entity types.Entity, name string, field conformance.ComparisonValue, parentEntity types.Entity) {
	var el *etree.Element
	switch entity := entity.(type) {
	case *matter.Feature:
		el = parent.CreateElement("feature")
		el.CreateAttr("name", entity.Code)
	case *matter.Command:
		el = parent.CreateElement("command")
		el.CreateAttr("name", entity.Name)
	case *matter.Event:
		el = parent.CreateElement("event")
		el.CreateAttr("name", entity.Name)
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			el = parent.CreateElement("attribute")
			el.CreateAttr("name", entity.Name)
		case types.EntityTypeStructField, types.EntityTypeCommandField, types.EntityTypeEventField:
			el = writeOptionalParentElement(parent, entity, parentEntity)
		default:
			slog.Warn("Unexpected field entity type when rendering conformance", slog.String("entityType", entity.EntityType().String()))
		}
	case *matter.Condition:
		el = parent.CreateElement("condition")
		el.CreateAttr("name", entity.Feature)
	case *matter.TypeDef:
		el = parent.CreateElement("typeDef")
		el.CreateAttr("name", entity.Name)
	case *matter.Enum:
		el = parent.CreateElement("enum")
		el.CreateAttr("name", entity.Name)
	case nil:
		el = parent.CreateElement("condition")
		el.CreateAttr("name", name)
	case *matter.EnumValue:
		enumParent := entity.Parent()
		if enumParent != nil && enumParent != parentEntity {
			switch enumParent := enumParent.(type) {
			case *matter.Enum:
				el = parent.CreateElement("enum")
				el.CreateAttr("name", enumParent.Name)
				el.CreateAttr("value", entity.Name)
			default:
				slog.Warn("Unexpected enum value parent entity type", log.Type("parentEntityType", enumParent))
			}
		} else {
			el = parent.CreateElement("value")
			el.CreateAttr("name", entity.Name)
		}
	case *matter.BitmapBit:
		bitParent := entity.Parent()
		if bitParent != nil && bitParent != parentEntity {
			switch bitParent := bitParent.(type) {
			case *matter.Bitmap:
				el = parent.CreateElement("bitmap")
				el.CreateAttr("name", bitParent.Name)
				el.CreateAttr("value", entity.Name())
			default:
				slog.Warn("Unexpected bitmap bit parent entity type", log.Type("parentEntityType", bitParent))
			}
		} else {
			el = parent.CreateElement("value")
			el.CreateAttr("name", entity.Name())
		}
	default:
		slog.Warn("Unexpected conformance entity type", log.Type("type", entity))
	}

	if field != nil && el != nil {
		renderComparisonValue(field, el, entity)
	}
}


func renderConformanceExpression(doc *spec.Doc, exp conformance.Expression, parent *etree.Element, parentEntity types.Entity) error {
	if exp == nil {
		return nil
	}
	switch e := exp.(type) {
	case *conformance.IdentifierExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		renderConformanceEntity(parent, e.Entity, e.ID, e.Field, parentEntity)
	case *conformance.EqualityExpression:
		return renderConformanceEqualityExpression(doc, e, parent, parentEntity)
	case *conformance.LogicalExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		var el *etree.Element
		switch e.Operand {
		case "&":
			el = parent.CreateElement("andTerm")
		case "|":
			el = parent.CreateElement("orTerm")
		case "^":
			el = parent.CreateElement("xorTerm")
		default:
			return fmt.Errorf("unknown operand: %s", e.Operand)
		}
		err := renderConformanceExpression(doc, e.Left, el, parentEntity)
		if err != nil {
			return fmt.Errorf("error rendering conformance expression %s: %w", e.Left.ASCIIDocString(), err)
		}
		for _, r := range e.Right {
			err = renderConformanceExpression(doc, r, el, parentEntity)
			if err != nil {
				return fmt.Errorf("error rendering conformance expression %s: %w", r.ASCIIDocString(), err)
			}

		}

	case *conformance.ReferenceExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		renderConformanceEntity(parent, e.Entity, e.Reference, e.Field, parentEntity)

	case *conformance.ComparisonExpression:
		switch e.Op {
		case conformance.ComparisonOperatorEqual:
			parent = parent.CreateElement("equalTerm")
		case conformance.ComparisonOperatorNotEqual:
			parent = parent.CreateElement("notEqualTerm")
		case conformance.ComparisonOperatorGreaterThan:
			parent = parent.CreateElement("greaterTerm")
		case conformance.ComparisonOperatorGreaterThanOrEqual:
			parent = parent.CreateElement("greaterOrEqualTerm")
		case conformance.ComparisonOperatorLessThan:
			parent = parent.CreateElement("lessTerm")
		case conformance.ComparisonOperatorLessThanOrEqual:
			parent = parent.CreateElement("lessOrEqualTerm")
		default:
			return fmt.Errorf("unexpected comparison expression operator: %s", e.Op.String())
		}
		err := renderComparisonValue(e.Left, parent, parentEntity)
		if err != nil {
			return err
		}
		err = renderComparisonValue(e.Right, parent, parentEntity)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unimplemented conformance expression type: %T", exp)
	}
	return nil
}
*/
