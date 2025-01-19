package sort

func SelectionSort[T int64 | float64](arr []T) {
	for i := 0; i < len(arr)-1; i++ {
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
