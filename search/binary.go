package search

func BinarySearch[V int64 | float64](arr []V, target V) int {
	idx := -1
	low, high := 0, len(arr)

	for low < high {
		mid := (low + high) / 2
		val := arr[mid]

		if val == target {
			idx = mid
			break
		} else if val < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return idx
}
