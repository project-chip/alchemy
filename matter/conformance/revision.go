package conformance

import "fmt"

type Revision int64

const CurrentRevsion Revision = -1

type RevisionExpression struct {
	Op    ComparisonOperator
	Left  Revision
	Right Revision
}

func (cr Revision) String() string {
	switch cr {
	case CurrentRevsion:
		return "Rev"
	default:
		return fmt.Sprintf("v%d", cr)
	}
}

func (re *RevisionExpression) ASCIIDocString() string {
	return fmt.Sprintf("%s %s %s", re.Left, re.Op.String(), re.Right)
}

func (re *RevisionExpression) Description() string {
	return fmt.Sprintf("%s %s %s", re.Left, re.Op.String(), re.Right)
}

func (re *RevisionExpression) Eval(context Context) (ExpressionResult, error) {
	revision, found := context.Value("Revision")
	if !found {
		return &expressionResult{value: true, confidence: ConfidencePossible}, nil
	}
	var currentRevision Revision
	switch rev := revision.(type) {
	case Revision:
		currentRevision = rev
	case int64:
		currentRevision = Revision(rev)
	case Confidence:
		return &expressionResult{value: true, confidence: rev}, nil
	default:
		return &expressionResult{value: false, confidence: ConfidenceImpossible}, nil
	}

	left := re.Left
	right := re.Right
	if left == CurrentRevsion {
		left = currentRevision
	} else if right == CurrentRevsion {
		right = currentRevision
	}
	switch re.Op {
	case ComparisonOperatorLessThan:
		if left < right {
			return &expressionResult{value: true, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorLessThanOrEqual:
		if left <= right {
			return &expressionResult{value: true, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorGreaterThan:
		if left > right {
			return &expressionResult{value: true, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorGreaterThanOrEqual:
		if left >= right {
			return &expressionResult{value: true, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorEqual:
		if left == right {
			return &expressionResult{value: true, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorNotEqual:
		if left != right {
			return &expressionResult{value: true, confidence: ConfidenceDefinite}, nil
		}
	}
	return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
}

func (re *RevisionExpression) Equal(e Expression) bool {
	if re == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	oee, ok := e.(*RevisionExpression)
	if !ok {
		return false
	}
	if re.Op != oee.Op {
		return false
	}
	if re.Left != oee.Left {
		return false
	}
	if re.Right != oee.Right {
		return false
	}
	return true
}

func (re *RevisionExpression) Clone() Expression {
	return &RevisionExpression{
		Op:    re.Op,
		Left:  re.Left,
		Right: re.Right,
	}
}

type RevisionRangeExpression struct {
	LeftOp  ComparisonOperator
	Left    Revision
	RightOp ComparisonOperator
	Right   Revision
}

func (rre *RevisionRangeExpression) ASCIIDocString() string {
	return fmt.Sprintf("%s %s %s %s %s", rre.Left, rre.LeftOp, CurrentRevsion, rre.RightOp, rre.Right)
}

func (rre *RevisionRangeExpression) Description() string {
	return rre.ASCIIDocString()
}

func (rre *RevisionRangeExpression) Eval(context Context) (ExpressionResult, error) {
	revision, found := context.Value("Revision")
	if !found {
		return &expressionResult{value: true, confidence: ConfidencePossible}, nil
	}
	var currentRevision Revision
	switch rev := revision.(type) {
	case Revision:
		currentRevision = rev
	case int64:
		currentRevision = Revision(rev)
	case Confidence:
		return &expressionResult{value: true, confidence: rev}, nil
	default:
		return &expressionResult{value: false, confidence: ConfidenceImpossible}, nil
	}
	switch rre.LeftOp {
	case ComparisonOperatorLessThan:
		if currentRevision <= rre.Left {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorLessThanOrEqual:
		if currentRevision < rre.Left {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorGreaterThan:
		if currentRevision >= rre.Left {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorGreaterThanOrEqual:
		if currentRevision > rre.Left {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	}
	switch rre.RightOp {
	case ComparisonOperatorLessThan:
		if currentRevision >= rre.Right {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorLessThanOrEqual:
		if currentRevision > rre.Right {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorGreaterThan:
		if currentRevision <= rre.Right {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	case ComparisonOperatorGreaterThanOrEqual:
		if currentRevision < rre.Right {
			return &expressionResult{value: false, confidence: ConfidenceDefinite}, nil
		}
	}
	return &expressionResult{value: true, confidence: ConfidenceDefinite}, nil
}

func (rre *RevisionRangeExpression) Equal(e Expression) bool {
	if rre == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	orre, ok := e.(*RevisionRangeExpression)
	if !ok {
		return false
	}
	if rre.LeftOp != orre.LeftOp {
		return false
	}
	if rre.Left != orre.Left {
		return false
	}
	if rre.RightOp != orre.RightOp {
		return false
	}
	if rre.Right != orre.Right {
		return false
	}
	return true
}

func (rre *RevisionRangeExpression) Clone() Expression {
	return &RevisionRangeExpression{
		LeftOp:  rre.LeftOp,
		Left:    rre.Left,
		RightOp: rre.RightOp,
		Right:   rre.Right,
	}
}

func checkRevision(rre *RevisionExpression, raw string) error {
	if rre.Left == 0 || rre.Right == 0 {
		return fmt.Errorf("invalid value in revision expression: %s", raw)
	}
	if rre.Left == CurrentRevsion {
		if rre.Right == CurrentRevsion {
			return fmt.Errorf("invalid comparison in revision expression: %s", raw)
		}
	} else if rre.Right != CurrentRevsion {
		return fmt.Errorf("revision expression must contain reference to current revision: %s", raw)
	}
	return nil
}

func checkRevisionRange(rre *RevisionRangeExpression, raw string) error {
	if rre.Left == 0 || rre.Right == 0 {
		return fmt.Errorf("invalid value in revision expression: %s", raw)
	}
	versionRange := rre.Right - rre.Left
	switch rre.LeftOp {
	case ComparisonOperatorLessThan:
		switch rre.RightOp {
		case ComparisonOperatorLessThan:
			if versionRange <= 1 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		case ComparisonOperatorLessThanOrEqual:
			if versionRange < 1 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		default:
			return fmt.Errorf("invalid comparison in revision range: %s", raw)
		}
	case ComparisonOperatorLessThanOrEqual:
		switch rre.RightOp {
		case ComparisonOperatorLessThan:
			if versionRange < 1 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		case ComparisonOperatorLessThanOrEqual:
			if versionRange < 0 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		default:
			return fmt.Errorf("invalid comparison in revision range: %s", raw)
		}
	case ComparisonOperatorGreaterThan:
		switch rre.RightOp {
		case ComparisonOperatorGreaterThan:
			if versionRange >= -1 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		case ComparisonOperatorGreaterThanOrEqual:
			if versionRange > -1 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		default:
			return fmt.Errorf("invalid comparison in revision range: %s", raw)
		}
	case ComparisonOperatorGreaterThanOrEqual:
		switch rre.RightOp {
		case ComparisonOperatorGreaterThan:
			if versionRange > -1 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		case ComparisonOperatorGreaterThanOrEqual:
			if versionRange > 0 {
				return fmt.Errorf("impossible comparison in revision range: %s", raw)
			}
		default:
			return fmt.Errorf("invalid comparison in revision range: %s", raw)
		}
	default:
		return fmt.Errorf("invalid comparison in revision range: %s", raw)
	}
	return nil
}

/*


2 < r < 5  3
2 < r < 4  2
2 < r < 3  1
2 < r < 2 0
2 < r < 1  -1
3 < r < 1 -2

2 < r <= 5  3
2 < r <= 4  2
2 < r <= 3  1
2 < r <= 2 0
2 < r <= 1  -1
3 < r <= 1 -2

2 <= r < 5  3
2 <= r < 4  2
2 <= r < 3  1
2 <= r <= 2 0 !
2 <= r <= 1  -1
3 <= r < 1 -2


2 <= r <= 5  3
2 <= r <= 4  2
2 <= r <= 3  1
2 <= r <= 2 0
2 <= r <= 1  -1
3 <= r <= 1 -2


2 > r > 5  3
2 > r > 4  2
2 > r > 3  1
2 > r > 2 0
2 > r > 1  -1
3 > r > 1 -2

2 > r >= 5  3
2 > r >= 4  2
2 > r >= 3  1
2 > r >= 2 0
2 > r >= 1  -1
3 > r >= 1 -2
*/
