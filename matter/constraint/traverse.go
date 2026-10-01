package constraint

import "iter"

func yieldConstraint(constraint Constraint, yield func(Constraint) bool) bool {
	if !yield(constraint) {
		return false
	}
	switch cc := constraint.(type) {
	case *ListConstraint:
		if !yieldConstraint(cc.Constraint, yield) {
			return false
		}
		if !yieldConstraint(cc.EntryConstraint, yield) {
			return false
		}
	case *LogicalConstraint:
		if !yieldConstraint(cc.Left, yield) {
			return false
		}
		for _, rl := range cc.Right {
			if !yieldConstraint(rl, yield) {
				return false
			}
		}
	case Set:
		for _, constr := range cc {
			if !yieldConstraint(constr, yield) {
				return false
			}
		}
	}
	return true
}

func Traverse(c Constraint) iter.Seq[Constraint] {
	return func(yield func(Constraint) bool) {
		yieldConstraint(c, yield)
	}
}

func yieldLimit(limit Limit, yield func(Limit) bool) bool {
	if limit != nil {
		if !yield(limit) {
			return false
		}
		switch ll := limit.(type) {
		case *CharacterLimit:
			if ll.ByteCount != nil && !yieldLimit(ll.ByteCount, yield) {
				return false
			}
			if ll.CodepointCount != nil && !yieldLimit(ll.CodepointCount, yield) {
				return false
			}
		case *IdentifierLimit:
			if ll.Field != nil && !yieldLimit(ll.Field, yield) {
				return false
			}
		case *LengthLimit:
			if ll.Reference != nil && !yieldLimit(ll.Reference, yield) {
				return false
			}
		case *LogicalLimit:
			if !yieldLimit(ll.Left, yield) {
				return false
			}
			for _, rl := range ll.Right {
				if !yieldLimit(rl, yield) {
					return false
				}
			}
		case *MathExpressionLimit:
			if !yieldLimit(ll.Left, yield) {
				return false
			}
			if !yieldLimit(ll.Right, yield) {
				return false
			}
		case *MaxOfLimit:
			for _, m := range ll.Maximums {
				if !yieldLimit(m, yield) {
					return false
				}
			}
		case *MinOfLimit:
			for _, m := range ll.Minimums {
				if !yieldLimit(m, yield) {
					return false
				}
			}
		case *ReferenceLimit:
			if !yieldLimit(ll.Field, yield) {
				return false
			}
		}
	}
	return true
}

func yieldConstraintLimits(c Constraint, l Limit, yield func(Constraint, Limit) bool) bool {

	yieldConstraintLimit := func(l Limit) bool {
		return yield(c, l)
	}
	return yieldLimit(l, yieldConstraintLimit)
}

func traverseConstraintLimits(c Constraint, yield func(Constraint, Limit) bool) bool {
	switch cc := c.(type) {
	case *ExactConstraint:
		if cc.Value != nil && !yieldConstraintLimits(c, cc.Value, yield) {
			return false
		}
	case *ListConstraint:
		if !traverseConstraintLimits(cc.Constraint, yield) {
			return false
		}
		if !traverseConstraintLimits(cc.EntryConstraint, yield) {
			return false
		}
	case *MaxConstraint:
		if !yieldConstraintLimits(cc, cc.Maximum, yield) {
			return false
		}
	case *MinConstraint:
		if !yieldConstraintLimits(cc, cc.Minimum, yield) {
			return false
		}
	case *RangeConstraint:
		if cc.Minimum != nil && !yieldConstraintLimits(c, cc.Minimum, yield) {
			return false
		}
		if cc.Maximum != nil && !yieldConstraintLimits(c, cc.Maximum, yield) {
			return false
		}
	case *TagListConstraint:
		if cc.Tags != nil && !yieldConstraintLimits(c, cc.Tags, yield) {
			return false
		}
	}
	return true
}

func TraverseConstraintLimits(c Constraint) iter.Seq2[Constraint, Limit] {
	return func(yield func(Constraint, Limit) bool) {
		for c := range Traverse(c) {
			if !traverseConstraintLimits(c, yield) {
				return
			}
		}
	}
}

func TraverseLimit(l Limit) iter.Seq[Limit] {
	return func(yield func(Limit) bool) {
		yieldLimit(l, yield)
	}
}
