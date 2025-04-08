package types

func MinExtreme(left DataTypeExtreme, right DataTypeExtreme) (min DataTypeExtreme) {
	if !left.IsNumeric() || !right.IsNumeric() {
		return
	}
	if left.Big().Cmp(right.Big()) < 0 {
		min = left
	} else {
		min = right
	}
	return
}

func MaxExtreme(left DataTypeExtreme, right DataTypeExtreme) (max DataTypeExtreme) {
	if !left.IsNumeric() || !right.IsNumeric() {
		return
	}
	if left.Big().Cmp(right.Big()) > 0 {
		max = left
	} else {
		max = right
	}
	return
}
