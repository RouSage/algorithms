package structures

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestStack(t *testing.T) {
	t.Run("NewStack", func(t *testing.T) {
		s := NewStack[int]()

		Equal(t, s.IsEmpty(), true)
		Equal(t, s.Size(), 0)
	})

	t.Run("Push", func(t *testing.T) {
		s := NewStack[int]()

		s.Push(1)
		s.Push(2)

		Equal(t, s.IsEmpty(), false)
		Equal(t, s.Size(), 2)
	})

	t.Run("Pop", func(t *testing.T) {
		s := NewStack[int]()

		s.Push(1)
		s.Push(2)

		val, ok := s.Pop()
		Equal(t, ok, true)
		Equal(t, val, 2)

		val, ok = s.Pop()
		Equal(t, ok, true)
		Equal(t, val, 1)

		val, ok = s.Pop()
		Equal(t, ok, false)
		Equal(t, val, 0)
	})

	t.Run("Peek", func(t *testing.T) {
		s := NewStack[int]()

		s.Push(1)
		s.Push(2)

		val, ok := s.Peek()
		Equal(t, ok, true)
		Equal(t, val, 2)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		s := NewStack[int]()

		Equal(t, s.IsEmpty(), true)

		s.Push(10)
		Equal(t, s.IsEmpty(), false)
	})
}
