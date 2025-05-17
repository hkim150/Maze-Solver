package dataStructure

type Stack[T any] struct {
	stack []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.stack) == 0 {
		var zero T
		return zero, false
	}

	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.stack) == 0 {
		var zero T
		return zero, false
	}

	return s.stack[len(s.stack)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) Clear() {
	s.stack = []T{}
}
