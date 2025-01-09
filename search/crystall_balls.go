package search

import "math"

func TwoCrystallBalls(breaks []bool) int {
	jump := int(math.Sqrt(float64(len(breaks))))
	i := jump

	for ; i < len(breaks); i += jump {
		if breaks[i] {
			break
		}
	}

	i -= jump

	for j := 0; j <= jump && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	return -1
}
