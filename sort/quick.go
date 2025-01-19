package sort

func qs[T int64 | float64](arr []T, low, high int) {
	if low >= high {
		return
	}

	pivotIdx := partition(arr, low, high)

	qs(arr, low, pivotIdx-1)
	qs(arr, pivotIdx+1, high)
}

func partition[T int64 | float64](arr []T, low, high int) int {
	pivot := arr[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[high], arr[idx] = arr[idx], pivot

	return idx
}

func QuickSort[T int64 | float64](arr []T) {
	qs(arr, 0, len(arr)-1)
}
