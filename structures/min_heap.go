package structures

type minHeap[T int | float64] struct {
	Length int
	data   []T
}

func NewMinHeap[T int | float64]() *minHeap[T] {
	return &minHeap[T]{
		Length: 0,
		data:   make([]T, 0),
	}
}

func (h *minHeap[T]) Insert(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(h.Length)
	h.Length++
}

func (h *minHeap[T]) Delete() (T, bool) {
	if h.Length == 0 {
		var zero T
		return zero, false
	}

	out := h.data[0]
	h.Length--

	if h.Length == 0 {
		return out, true
	}

	h.data[0] = h.data[h.Length]
	h.heapifyDown(0)

	return out, true
}

func (h *minHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := parent(idx)
	parentVal := h.data[p]
	val := h.data[idx]

	if parentVal > val {
		h.data[p], h.data[idx] = val, parentVal
		h.heapifyUp(p)
	}
}

func (h *minHeap[T]) heapifyDown(idx int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if idx >= h.Length || lIdx >= h.Length {
		return
	}

	lVal := h.data[lIdx]
	rVal := h.data[rIdx]
	val := h.data[idx]

	if lVal > rVal && val > rVal {
		h.data[idx], h.data[rIdx] = rVal, val
		h.heapifyDown(rIdx)
	} else if rVal > lVal && val > lVal {
		h.data[idx], h.data[lIdx] = lVal, val
		h.heapifyDown(lIdx)
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}
