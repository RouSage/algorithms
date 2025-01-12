package structures

type sNode[T any] struct {
	value T
	prev  *sNode[T]
}

type Stack[T any] struct {
	head *sNode[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		size: 0,
	}
}

func (s *Stack[T]) Push(item T) {
	value := &sNode[T]{
		value: item,
	}

	if s.IsEmpty() {
		s.head = value
	} else {
		value.prev = s.head
		s.head = value
	}

	s.size++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	s.size--

	head := s.head
	s.head = head.prev
	head.prev = nil

	return head.value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.head.value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) Size() int {
	return s.size
}
