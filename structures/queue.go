package structures

type qNode[T any] struct {
	value T
	next  *qNode[T]
}

type Queue[T any] struct {
	head *qNode[T]
	tail *qNode[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		size: 0,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	value := &qNode[T]{
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
