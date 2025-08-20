package asciidoc

import (
	"fmt"
	"strconv"
	"strings"
)

type ConditionalContext interface {
	IsSet(name string) bool
	Get(name string) any
	Set(name string, value any)
	Unset(name string)
}

type IfDef struct {
	position
	raw

	Attributes AttributeNames
	Union      ConditionalUnion

	Inline bool
}

func NewIfDef(attributes []AttributeName, union ConditionalUnion) *IfDef {
	return &IfDef{Attributes: attributes, Union: union}
}

func (IfDef) Type() ElementType {
	return ElementTypeBlock
}

func (a *IfDef) Equals(o Element) bool {
	oa, ok := o.(*IfDef)
	if !ok {
		return false
	}
	return a.Attributes.Equals(oa.Attributes)
}

func (a *IfDef) Clone() Element {
	return &IfDef{position: a.position, raw: a.raw, Attributes: a.Attributes.Clone(), Union: a.Union}
}

func (a *IfDef) Eval(cc ConditionalContext) bool {
	return ifIsTrue(cc, a.Attributes, a.Union)
}

type IfNDef struct {
	position
	raw

	Attributes AttributeNames
	Union      ConditionalUnion

	Inline bool
}

func NewIfNDef(attributes []AttributeName, union ConditionalUnion) *IfNDef {
	return &IfNDef{Attributes: attributes, Union: union}
}

func (IfNDef) Type() ElementType {
	return ElementTypeBlock
}

func (id *IfNDef) Equals(o Element) bool {
	oa, ok := o.(*IfNDef)
	if !ok {
		return false
	}
	return id.Attributes.Equals(oa.Attributes)
}

func (id *IfNDef) Clone() Element {
	return &IfNDef{position: id.position, raw: id.raw, Attributes: id.Attributes.Clone(), Union: id.Union}
}

func (id *IfNDef) Eval(cc ConditionalContext) bool {
	return !ifIsTrue(cc, id.Attributes, id.Union)
}

type InlineIfDef struct {
	position
	raw

	Elements

	Attributes AttributeNames
	Union      ConditionalUnion
}

func NewInlineIfDef(attributes []AttributeName, union ConditionalUnion) *InlineIfDef {
	return &InlineIfDef{Attributes: attributes}
}

func (InlineIfDef) Type() ElementType {
	return ElementTypeInline
}

func (iid *InlineIfDef) Equals(o Element) bool {
	oa, ok := o.(*InlineIfDef)
	if !ok {
		return false
	}

	if !iid.Attributes.Equals(oa.Attributes) {
		return false
	}
	return iid.Elements.Equals(oa.Elements)
}

func (iid *InlineIfDef) Clone() Element {
	return &InlineIfDef{position: iid.position, raw: iid.raw, Attributes: iid.Attributes.Clone(), Elements: iid.Elements.Clone(), Union: iid.Union}
}

func (iid *InlineIfDef) Eval(cc ConditionalContext) bool {
	return ifIsTrue(cc, iid.Attributes, iid.Union)
}

type InlineIfNDef struct {
	position
	raw

	Elements

	Attributes AttributeNames
	Union      ConditionalUnion
}

func NewInlineIfNDef(attributes []AttributeName, union ConditionalUnion) *InlineIfNDef {
	return &InlineIfNDef{Attributes: attributes}
}

func (InlineIfNDef) Type() ElementType {
	return ElementTypeInline
}

func (iind *InlineIfNDef) Equals(o Element) bool {
	oa, ok := o.(*InlineIfNDef)
	if !ok {
		return false
	}

	if !iind.Attributes.Equals(oa.Attributes) {
		return false
	}
	return iind.Elements.Equals(oa.Elements)
}

func (iind *InlineIfNDef) Clone() Element {
	return &InlineIfNDef{position: iind.position, raw: iind.raw, Attributes: iind.Attributes.Clone(), Elements: iind.Elements.Clone(), Union: iind.Union}
}

func (iind *InlineIfNDef) Eval(cc ConditionalContext) bool {
	return !ifIsTrue(cc, iind.Attributes, iind.Union)
}

type EndIf struct {
	position
	raw

	Attributes AttributeNames
	Union      ConditionalUnion

	Open Element
}

func NewEndIf(attributes []AttributeName, union ConditionalUnion) *EndIf {
	return &EndIf{Attributes: attributes, Union: union}
}

func (EndIf) Type() ElementType {
	return ElementTypeBlock
}

func (ei *EndIf) Equals(o Element) bool {
	oa, ok := o.(*EndIf)
	if !ok {
		return false
	}
	return ei.Attributes.Equals(oa.Attributes)
}

func (ei *EndIf) Clone() Element {
	return &EndIf{position: ei.position, raw: ei.raw, Attributes: ei.Attributes.Clone(), Union: ei.Union, Open: ei.Open}

}

type ConditionalOperator uint8

const (
	ConditionalOperatorNone ConditionalOperator = iota
	ConditionalOperatorEqual
	ConditionalOperatorNotEqual
	ConditionalOperatorLessThan
	ConditionalOperatorLessThanOrEqual
	ConditionalOperatorGreaterThan
	ConditionalOperatorGreaterThanOrEqual
)

func (co ConditionalOperator) String() string {
	switch co {
	case ConditionalOperatorEqual:
		return "=="
	case ConditionalOperatorNotEqual:
		return "!="
	case ConditionalOperatorLessThan:
		return "<"
	case ConditionalOperatorLessThanOrEqual:
		return "<="
	case ConditionalOperatorGreaterThan:
		return ">"
	case ConditionalOperatorGreaterThanOrEqual:
		return ">="
	default:
		return ""
	}
}

type ConditionalUnion uint8

const (
	ConditionalUnionAny ConditionalUnion = iota
	ConditionalUnionAll
)

type IfEval struct {
	position
	raw

	Left     IfEvalValue
	Operator ConditionalOperator
	Right    IfEvalValue

	Inline bool
}

func NewIfEval(left IfEvalValue, operator ConditionalOperator, right IfEvalValue) *IfEval {
	return &IfEval{Left: left, Operator: operator, Right: right}
}

func (IfEval) Type() ElementType {
	return ElementTypeBlock
}

func (ie *IfEval) Equals(o Element) bool {
	oa, ok := o.(*IfEval)
	if !ok {
		return false
	}
	if !ie.Left.Equals(oa.Left) {
		return false
	}
	if ie.Operator != oa.Operator {
		return false
	}
	return ie.Right.Equals(oa.Right)
}

func (ie *IfEval) Clone() Element {
	return &IfEval{position: ie.position, raw: ie.raw, Left: ie.Left.Clone(), Operator: ie.Operator, Right: ie.Right.Clone()}
}

func (ie *IfEval) Eval(cc ConditionalContext) (bool, error) {
	return eval(cc, ie.Left, ie.Operator, ie.Right)
}

type IfDefBlock struct {
	position
	raw

	Attributes AttributeNames
	Elements
	Union ConditionalUnion
}

func NewIfDefBlock(attributes []AttributeName, union ConditionalUnion) *IfDefBlock {
	return &IfDefBlock{Attributes: attributes, Union: union}
}

func (IfDefBlock) Type() ElementType {
	return ElementTypeBlock
}

func (idb *IfDefBlock) Equals(o Element) bool {
	oa, ok := o.(*IfDefBlock)
	if !ok {
		return false
	}
	if !idb.Attributes.Equals(oa.Attributes) {
		return false
	}
	return idb.Elements.Equals(oa.Elements)
}

func (idb *IfDefBlock) Clone() Element {
	return &IfDefBlock{position: idb.position, raw: idb.raw, Attributes: idb.Attributes.Clone(), Elements: idb.Elements.Clone(), Union: idb.Union}
}

func (idb *IfDefBlock) Eval(cc ConditionalContext) bool {
	return ifIsTrue(cc, idb.Attributes, idb.Union)
}

func ifIsTrue(cc ConditionalContext, a AttributeNames, union ConditionalUnion) bool {
	if len(a) == 0 {
		return false
	}
	for _, name := range a {
		isSet := cc.IsSet(string(name))
		switch union {
		case ConditionalUnionAny:
			if isSet {
				return true
			}
		case ConditionalUnionAll:
			if !isSet {
				return false
			}
		}
	}
	switch union {
	case ConditionalUnionAny:
		// If we got here, they were all unset
		return false
	default: //ConditionalUnionAll
		// If we got here, they were all set
		return true
	}

}

type IfNDefBlock struct {
	position
	raw

	Attributes AttributeNames
	Elements
	Union ConditionalUnion
}

func NewIfNDefBlock(attributes []AttributeName, union ConditionalUnion) *IfNDefBlock {
	return &IfNDefBlock{Attributes: attributes, Union: union}
}

func (IfNDefBlock) Type() ElementType {
	return ElementTypeBlock
}

func (indb *IfNDefBlock) Equals(o Element) bool {
	oa, ok := o.(*IfNDefBlock)
	if !ok {
		return false
	}
	if !indb.Attributes.Equals(oa.Attributes) {
		return false
	}
	return indb.Elements.Equals(oa.Elements)
}

func (indb *IfNDefBlock) Clone() Element {
	return &IfNDefBlock{position: indb.position, raw: indb.raw, Attributes: indb.Attributes.Clone(), Elements: indb.Elements.Clone(), Union: indb.Union}
}

func (indb *IfNDefBlock) Eval(cc ConditionalContext) bool {
	return !ifIsTrue(cc, indb.Attributes, indb.Union)
}

type IfEvalValue struct {
	Quote AttributeQuoteType
	Value Elements
}

func (iev IfEvalValue) Equals(oiev IfEvalValue) bool {
	return iev.Quote == oiev.Quote && iev.Value.Equals(oiev.Value)
}

func (iev IfEvalValue) Clone() IfEvalValue {
	return IfEvalValue{Quote: iev.Quote, Value: iev.Value.Clone()}
}

func (iev IfEvalValue) Eval(cc ConditionalContext) (any, error) {
	switch iev.Quote {
	case AttributeQuoteTypeNone:
		val, err := iev.StringValue(cc)
		if err != nil {
			return nil, err
		}
		b, err := strconv.ParseBool(val)
		if err == nil {
			return b, nil
		}
		n, err := strconv.Atoi(val)
		if err == nil {
			return n, nil
		}
		if iev.Value.IsWhitespace() {
			return " ", nil
		}
		return val, nil
	default:
		return iev.StringValue(cc)
	}
}

func (iev IfEvalValue) StringValue(cc ConditionalContext) (string, error) {
	var sb strings.Builder
	for _, el := range iev.Value {
		switch el := el.(type) {
		case *String:
			sb.WriteString(el.Value)
		case *CharacterReplacementReference:
			sb.WriteString(el.ReplacementValue())
		case *UserAttributeReference:
			val := cc.Get(el.Name())
			switch v := val.(type) {
			case nil:
			case *String:
				sb.WriteString(v.Value)
			case string:
				sb.WriteString(v)
			default:
				return "", fmt.Errorf("unexpected ifeval user attribute reference value type: %T", val)
			}
		default:
			return "", fmt.Errorf("unexpected ifeval value type: %T", el)
		}
	}
	return sb.String(), nil
}

type IfEvalBlock struct {
	position
	raw

	Attributes AttributeNames
	Elements

	Left     IfEvalValue
	Operator ConditionalOperator
	Right    IfEvalValue
}

func NewIfEvalBlock(left IfEvalValue, operator ConditionalOperator, right IfEvalValue) *IfEvalBlock {
	return &IfEvalBlock{Left: left, Operator: operator, Right: right}
}

func (IfEvalBlock) Type() ElementType {
	return ElementTypeBlock
}

func (ieb *IfEvalBlock) Equals(o Element) bool {
	oa, ok := o.(*IfEvalBlock)
	if !ok {
		return false
	}
	if !ieb.Left.Equals(oa.Left) {
		return false
	}
	if ieb.Operator != oa.Operator {
		return false
	}
	if !ieb.Right.Equals(oa.Right) {
		return false
	}
	return ieb.Elements.Equals(oa.Elements)
}

func (ieb *IfEvalBlock) Clone() Element {
	return &IfEvalBlock{position: ieb.position, raw: ieb.raw, Left: ieb.Left.Clone(), Operator: ieb.Operator, Right: ieb.Right.Clone(), Elements: ieb.Elements.Clone()}
}

func (ieb *IfEvalBlock) Eval(cc ConditionalContext) (bool, error) {
	return eval(cc, ieb.Left, ieb.Operator, ieb.Right)
}

func eval(cc ConditionalContext, leftValue IfEvalValue, operator ConditionalOperator, rightValue IfEvalValue) (bool, error) {
	left, err := leftValue.Eval(cc)
	if err != nil {
		return false, err
	}
	right, err := rightValue.Eval(cc)
	if err != nil {
		return false, err
	}
	switch left := left.(type) {
	case bool:
		r, ok := right.(bool)
		if !ok {
			return false, nil
		}
		switch operator {
		case ConditionalOperatorEqual:
			return left == r, nil
		default:
			return left != r, nil
		}
	case string:
		r, ok := right.(string)
		if !ok {
			return false, nil
		}
		c := strings.Compare(left, r)
		switch operator {
		case ConditionalOperatorEqual:
			return c == 0, nil
		case ConditionalOperatorNotEqual:
			return c != 0, nil
		case ConditionalOperatorLessThan:
			return c == -1, nil
		case ConditionalOperatorLessThanOrEqual:
			return c <= 0, nil
		case ConditionalOperatorGreaterThan:
			return c == 1, nil
		case ConditionalOperatorGreaterThanOrEqual:
			return c >= 0, nil
		}
	case int:
		r, ok := right.(int)
		if !ok {
			return false, nil
		}
		switch operator {
		case ConditionalOperatorEqual:
			return left == r, nil
		case ConditionalOperatorNotEqual:
			return left != r, nil
		case ConditionalOperatorLessThan:
			return left < r, nil
		case ConditionalOperatorLessThanOrEqual:
			return left <= r, nil
		case ConditionalOperatorGreaterThan:
			return left > r, nil
		case ConditionalOperatorGreaterThanOrEqual:
			return left >= r, nil
		}
	}
	return false, nil
}
