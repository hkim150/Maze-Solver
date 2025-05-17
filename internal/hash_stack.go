package internal

type HashStack[T comparable] struct {
	stack []T
	elemCount map[T]int
}

func NewHashStack[T comparable]() *HashStack[T] {
	return &HashStack[T]{
		stack: []T{},
		elemCount:   make(map[T]int),
	}
}

func (s *HashStack[T]) Push(value T) {
	s.stack = append(s.stack, value)
	s.elemCount[value]++
}

func (s *HashStack[T]) Pop() (T, bool) {
	if len(s.stack) == 0 {
		var zero T
		return zero, false
	}

	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.elemCount[value]--

	if s.elemCount[value] == 0 {
		delete(s.elemCount, value)
	}

	return value, true
}

func (s *HashStack[T]) Peek() (T, bool) {
	if len(s.stack) == 0 {
		var zero T
		return zero, false
	}

	return s.stack[len(s.stack)-1], true
}

func (s *HashStack[T]) Len () int {
	return len(s.stack)
}

func (s *HashStack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *HashStack[T]) Contains(value T) bool {
	count, _ := s.elemCount[value]
	return count > 0
}

func (s *HashStack[T]) Clear() {
	s.stack = []T{}
	s.elemCount = make(map[T]int)
}

func (s *HashStack[T]) ToSlice() []T {
	return s.stack
}
