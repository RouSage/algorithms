package structures

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestLinkedList(t *testing.T) {
	t.Run("NewLinkedList", func(t *testing.T) {
		list := NewLinkedList[int]()
		Equal(t, list.Size(), 0)
	})

	t.Run("InsertAt", func(t *testing.T) {
		list := NewLinkedList[int]()

		err := list.InsertAt(1, 0)
		Equal(t, list.Size(), 1)
		Equal(t, err, nil)

		err = list.InsertAt(2, 0)
		Equal(t, list.Size(), 2)
		Equal(t, err, nil)

		err = list.InsertAt(3, 4)
		Equal(t, list.Size(), 2)
		Equal(t, err, ErrIndexOutOfRange)

		err = list.InsertAt(3, -1)
		Equal(t, list.Size(), 2)
		Equal(t, err, ErrIndexOutOfRange)
	})

	t.Run("Remove", func(t *testing.T) {
		list := NewLinkedList[int]()

		list.InsertAt(1, 0)
		list.InsertAt(2, 1)
		list.InsertAt(3, 2)
		list.InsertAt(4, 3)

		removed, ok := list.Remove(2)
		Equal(t, list.Size(), 3)
		Equal(t, ok, true)
		Equal(t, removed, 2)

		removed, ok = list.Remove(5)
		Equal(t, list.Size(), 3)
		Equal(t, ok, false)
	})

	t.Run("RemoveAt", func(t *testing.T) {
		list := NewLinkedList[int]()

		list.InsertAt(1, 0)
		list.InsertAt(2, 1)
		list.InsertAt(3, 2)
		list.InsertAt(4, 3)

		removed, ok := list.RemoveAt(0)
		Equal(t, list.Size(), 3)
		Equal(t, ok, true)
		Equal(t, removed, 1)

		removed, ok = list.RemoveAt(4)
		Equal(t, list.Size(), 3)
		Equal(t, ok, false)
	})

	t.Run("Append", func(t *testing.T) {
		list := NewLinkedList[int]()

		list.Append(1)
		Equal(t, list.Size(), 1)

		list.Append(2)
		Equal(t, list.Size(), 2)

		val, ok := list.Get(0)
		Equal(t, ok, true)
		Equal(t, val, 1)

	})

	t.Run("Prepend", func(t *testing.T) {
		list := NewLinkedList[int]()

		list.Prepend(1)
		Equal(t, list.Size(), 1)

		list.Prepend(2)
		Equal(t, list.Size(), 2)

		val, ok := list.Get(0)
		Equal(t, ok, true)
		Equal(t, val, 2)
	})

	t.Run("Get", func(t *testing.T) {
		list := NewLinkedList[int]()

		list.Append(1)
		list.Append(2)
		list.Append(3)
		list.Append(4)

		item, ok := list.Get(2)
		Equal(t, ok, true)
		Equal(t, item, 3)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		list := NewLinkedList[int]()
		Equal(t, list.IsEmpty(), true)

		list.Append(1)
		Equal(t, list.IsEmpty(), false)
	})

	t.Run("Size", func(t *testing.T) {
		list := NewLinkedList[int]()

		Equal(t, list.Size(), 0)

		list.Append(1)
		Equal(t, list.Size(), 1)

		list.Append(2)
		Equal(t, list.Size(), 2)

		list.Remove(2)
		Equal(t, list.Size(), 1)
	})
}
