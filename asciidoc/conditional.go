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

func (a *IfNDef) Equals(o Element) bool {
	oa, ok := o.(*IfNDef)
	if !ok {
		return false
	}
	return a.Attributes.Equals(oa.Attributes)
}

func (a *IfNDef) Eval(cc ConditionalContext) bool {
	return !ifIsTrue(cc, a.Attributes, a.Union)
}

type InlineIfDef struct {
	position
	raw

	Set

	Attributes AttributeNames
	Union      ConditionalUnion
}

func NewInlineIfDef(attributes []AttributeName, union ConditionalUnion) *InlineIfDef {
	return &InlineIfDef{Attributes: attributes}
}

func (InlineIfDef) Type() ElementType {
	return ElementTypeInline
}

func (a *InlineIfDef) Equals(o Element) bool {
	oa, ok := o.(*InlineIfDef)
	if !ok {
		return false
	}

	if !a.Attributes.Equals(oa.Attributes) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

func (a *InlineIfDef) Eval(cc ConditionalContext) bool {
	return ifIsTrue(cc, a.Attributes, a.Union)
}

type InlineIfNDef struct {
	position
	raw

	Set

	Attributes AttributeNames
	Union      ConditionalUnion
}

func NewInlineIfNDef(attributes []AttributeName, union ConditionalUnion) *InlineIfNDef {
	return &InlineIfNDef{Attributes: attributes}
}

func (InlineIfNDef) Type() ElementType {
	return ElementTypeInline
}

func (a *InlineIfNDef) Equals(o Element) bool {
	oa, ok := o.(*InlineIfNDef)
	if !ok {
		return false
	}

	if !a.Attributes.Equals(oa.Attributes) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

func (a *InlineIfNDef) Eval(cc ConditionalContext) bool {
	return !ifIsTrue(cc, a.Attributes, a.Union)
}

type EndIf struct {
	position
	raw

	Attributes AttributeNames
	Union      ConditionalUnion
}

func NewEndIf(attributes []AttributeName, union ConditionalUnion) *EndIf {
	return &EndIf{Attributes: attributes, Union: union}
}

func (EndIf) Type() ElementType {
	return ElementTypeBlock
}

func (a *EndIf) Equals(o Element) bool {
	oa, ok := o.(*EndIf)
	if !ok {
		return false
	}
	return a.Attributes.Equals(oa.Attributes)
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

func (a *IfEval) Equals(o Element) bool {
	oa, ok := o.(*IfEval)
	if !ok {
		return false
	}
	if a.Left.Equals(oa.Left) {
		return false
	}
	if a.Operator != oa.Operator {
		return false
	}
	return a.Right.Equals(oa.Right)
}

func (a *IfEval) Eval(cc ConditionalContext) (bool, error) {
	return eval(cc, a.Left, a.Operator, a.Right)
}

type IfDefBlock struct {
	position
	raw

	Attributes AttributeNames
	Set
	Union ConditionalUnion
}

func NewIfDefBlock(attributes []AttributeName, union ConditionalUnion) *IfDefBlock {
	return &IfDefBlock{Attributes: attributes, Union: union}
}

func (IfDefBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *IfDefBlock) Equals(o Element) bool {
	oa, ok := o.(*IfDefBlock)
	if !ok {
		return false
	}
	if !a.Attributes.Equals(oa.Attributes) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

func (a *IfDefBlock) Eval(cc ConditionalContext) bool {
	return ifIsTrue(cc, a.Attributes, a.Union)
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
	Set
	Union ConditionalUnion
}

func NewIfNDefBlock(attributes []AttributeName, union ConditionalUnion) *IfNDefBlock {
	return &IfNDefBlock{Attributes: attributes, Union: union}
}

func (IfNDefBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *IfNDefBlock) Equals(o Element) bool {
	oa, ok := o.(*IfNDefBlock)
	if !ok {
		return false
	}
	if !a.Attributes.Equals(oa.Attributes) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

func (a *IfNDefBlock) Eval(cc ConditionalContext) bool {
	return !ifIsTrue(cc, a.Attributes, a.Union)
}

type IfEvalValue struct {
	Quote AttributeQuoteType
	Value Set
}

func (iev IfEvalValue) Equals(oiev IfEvalValue) bool {
	return iev.Quote == oiev.Quote && iev.Value.Equals(oiev.Value)
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
	Set

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

func (a *IfEvalBlock) Equals(o Element) bool {
	oa, ok := o.(*IfEvalBlock)
	if !ok {
		return false
	}
	if !a.Left.Equals(oa.Left) {
		return false
	}
	if a.Operator != oa.Operator {
		return false
	}
	if !a.Right.Equals(oa.Right) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

func (a *IfEvalBlock) Eval(cc ConditionalContext) (bool, error) {
	return eval(cc, a.Left, a.Operator, a.Right)
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
