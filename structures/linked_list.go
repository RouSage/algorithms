package structures

import "errors"

var (
	ErrIndexOutOfRange = errors.New("index out of range")
)

type node[T comparable] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

type LinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		size: 0,
	}
}

func (l *LinkedList[T]) InsertAt(item T, idx int) error {
	if idx < 0 || idx > l.size {
		return ErrIndexOutOfRange
	} else if idx == l.size {
		l.Append(item)
		return nil
	} else if idx == 0 {
		l.Prepend(item)
		return nil
	}

	l.size++

	curr := l.getAt(idx)

	value := &node[T]{
		value: item,
		prev:  curr.prev,
		next:  curr,
	}

	curr.prev = value

	if value.prev != nil {
		value.prev.next = value
	}

	return nil
}

func (l *LinkedList[T]) Remove(item T) (T, bool) {
	curr := l.head
	for i := 0; i < l.size && curr != nil; i++ {
		if curr.value == item {
			break
		}

		curr = curr.next
	}

	if curr == nil {
		var zero T
		return zero, false
	}

	return l.removeNode(curr)
}

func (l *LinkedList[T]) RemoveAt(idx int) (T, bool) {
	node := l.getAt(idx)

	if node == nil {
		var zero T
		return zero, false
	}

	return l.removeNode(node)
}

func (l *LinkedList[T]) Append(item T) {
	value := &node[T]{
		value: item,
		prev:  l.tail,
	}

	if l.IsEmpty() {
		l.size++
		l.head = value
		l.tail = value
		return
	}

	l.size++

	l.tail.next = value
	l.tail = value
}

func (l *LinkedList[T]) Prepend(item T) {
	value := &node[T]{
		value: item,
		next:  l.head,
	}

	if l.IsEmpty() {
		l.size++
		l.head = value
		l.tail = value
		return
	}

	l.size++

	l.head.prev = value
	l.head = value
}

func (l *LinkedList[T]) Get(idx int) (T, bool) {
	node := l.getAt(idx)

	if node == nil {
		var zero T
		return zero, false
	}

	return node.value, true
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) getAt(idx int) *node[T] {
	curr := l.head
	for i := 0; i < idx && curr != nil; i++ {
		curr = curr.next
	}

	return curr
}

func (l *LinkedList[T]) removeNode(node *node[T]) (T, bool) {
	l.size--

	if l.IsEmpty() {
		l.head = nil
		l.tail = nil
		return node.value, true
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == l.head {
		l.head = node.next
	}
	if node == l.tail {
		l.tail = node.prev
	}

	node.prev = nil
	node.next = nil

	return node.value, true

}
