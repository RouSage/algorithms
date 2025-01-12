package structures

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		size: 0,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	value := &node[T]{
		value: item,
	}

	if q.IsEmpty() {
		q.tail, q.head = value, value
	} else {
		q.tail.next = value
		q.tail = value
	}

	q.size++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	q.size--

	head := q.head
	q.head = head.next
	head.next = nil

	if q.IsEmpty() {
		q.tail = nil
	}

	return head.value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.head.value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}
