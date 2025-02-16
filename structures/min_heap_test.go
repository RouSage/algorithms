package structures

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap[int]()

	t.Run("Insert", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			h.Insert(i)
		}

		Equal(t, 10, h.Length)
	})

	t.Run("Delete", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			out, ok := h.Delete()
			Equal(t, true, ok)
			Equal(t, i, out)
		}

		Equal(t, 0, h.Length)

		_, ok := h.Delete()
		Equal(t, false, ok)
	})
}
