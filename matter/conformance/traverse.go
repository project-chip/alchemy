package conformance

import (
	"iter"
)

func yieldConformanceExpression(c Conformance, e Expression, yield func(Conformance, Expression) bool) bool {
	if !yield(c, e) {
		return false
	}
	switch exp := e.(type) {
	case *EqualityExpression:
		if !yieldConformanceExpression(c, exp.Left, yield) {
			return false
		}
		if !yieldConformanceExpression(c, exp.Right, yield) {
			return false
		}
	case *LogicalExpression:
		if !yieldConformanceExpression(c, exp.Left, yield) {
			return false
		}
		for _, r := range exp.Right {
			if !yieldConformanceExpression(c, r, yield) {
				return false
			}
		}
	}
	return true
}

func yieldComparisonValue(v ComparisonValue, yield func(ComparisonValue) bool) bool {
	if !yield(v) {
		return false
	}
	switch v := v.(type) {
	case *MathOperation:
		if !yieldComparisonValue(v.Left, yield) {
			return false
		}
		if !yieldComparisonValue(v.Right, yield) {
			return false
		}
	case *IdentifierValue:
		if !yieldComparisonValue(v.Field, yield) {
			return false
		}
	case *ReferenceValue:
		if !yieldComparisonValue(v.Field, yield) {
			return false
		}
	}
	return true
}

func TraverseComparisonValue(v ComparisonValue) iter.Seq[ComparisonValue] {
	return func(yield func(ComparisonValue) bool) {
		yieldComparisonValue(v, yield)
	}
}

func TraverseExpressions(c Conformance) iter.Seq2[Conformance, Expression] {
	return func(yield func(Conformance, Expression) bool) {
		for c := range Traverse(c) {
			switch cc := c.(type) {
			case *Mandatory:
				if !yieldConformanceExpression(cc, cc.Expression, yield) {
					return
				}
			case *Optional:
				if !yieldConformanceExpression(cc, cc.Expression, yield) {
					return
				}
			}
		}
	}
}

func Traverse(c Conformance) iter.Seq[Conformance] {
	return func(yield func(Conformance) bool) {
		switch cc := c.(type) {
		case Set:
			for _, constr := range cc {
				if !yield(constr) {
					return
				}
			}
		default:
			yield(c)
		}
	}
}
