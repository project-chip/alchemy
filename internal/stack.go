package internal

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(v T) {
	s.stack = append(s.stack, v)
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if len(s.stack) == 0 {
		return
	}

	ok = true
	v = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return
}

func (s *Stack[T]) Peek() (v T) {
	if len(s.stack) == 0 {
		return
	}
	v = s.stack[len(s.stack)-1]
	return
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}
