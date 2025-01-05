package search

func LinearSearch[V int64 | float64](arr []V, target V) int {
	idx := -1

	for i, v := range arr {
		if v == target {
			idx = i
			break
		}
	}

	return idx
}
