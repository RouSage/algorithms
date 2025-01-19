package sort

func InsertionSort[T int64 | float64](arr []T) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
