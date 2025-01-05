package search

func BinarySearch[V int64 | float64](arr []V, target V) int {
	idx := -1
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			idx = mid
			break
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return idx
}
