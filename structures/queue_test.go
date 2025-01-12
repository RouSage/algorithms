package structures

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestQueue(t *testing.T) {
	t.Run("NewQueue", func(t *testing.T) {
		q := NewQueue[int]()
		Equal(t, q.Size(), 0)

		_, ok := q.Peek()
		Equal(t, ok, false)
	})

	t.Run("Enqueue", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		Equal(t, q.Size(), 3)

		val, ok := q.Peek()
		Equal(t, ok, true)
		Equal(t, val, 1)
	})

	t.Run("Peek", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		val, ok := q.Peek()
		Equal(t, ok, true)
		Equal(t, val, 1)
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		val, ok := q.Dequeue()
		Equal(t, ok, true)
		Equal(t, val, 1)

		val, ok = q.Dequeue()
		Equal(t, ok, true)
		Equal(t, val, 2)

		val, ok = q.Dequeue()
		Equal(t, ok, true)
		Equal(t, val, 3)

		val, ok = q.Dequeue()
		Equal(t, ok, false)
		Equal(t, val, 0)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		q := NewQueue[int]()
		Equal(t, q.IsEmpty(), true)

		q.Enqueue(1)
		Equal(t, q.IsEmpty(), false)
	})
}
