package structures

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestGraphBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		graph    weightedAdjacencyMatrix
		source   int
		target   int
		expected []int
	}{
		{
			//     >(1)<--->(4) ---->(5)
			//    /          |       /|
			// (0)     ------|------- |
			//    \   v      v        v
			//     >(2) --> (3) <----(6)
			graph: weightedAdjacencyMatrix{
				{0, 3, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 7, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 5, 0, 2, 0},
				{0, 0, 18, 0, 0, 0, 1},
				{0, 0, 0, 1, 0, 0, 1},
			},
			source:   0,
			target:   6,
			expected: []int{0, 1, 4, 5, 6},
		},
		{
			graph: weightedAdjacencyMatrix{
				{0, 3, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 7, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 5, 0, 2, 0},
				{0, 0, 18, 0, 0, 0, 1},
				{0, 0, 0, 1, 0, 0, 1},
			},
			source:   6,
			target:   0,
			expected: nil,
		},
		{
			graph: weightedAdjacencyMatrix{
				{0, 3, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 7, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 5, 0, 2, 0},
				{0, 0, 18, 0, 0, 0, 1},
				{0, 0, 0, 1, 0, 0, 1},
			},
			source:   3,
			target:   1,
			expected: nil,
		},
		{
			// (0)-->(3)
			//  ^\     ^
			//  | \    |
			//  v  \   v
			// (1)  ->(2)
			graph: weightedAdjacencyMatrix{
				{0, 20, 19, 18},
				{17, 0, 16, 0},
				{0, 0, 0, 15},
				{0, 0, 14, 0},
			},
			source:   0,
			target:   3,
			expected: []int{0, 3},
		},
		{
			graph: weightedAdjacencyMatrix{
				{0, 20, 19, 18},
				{17, 0, 16, 0},
				{0, 0, 0, 15},
				{0, 0, 14, 0},
			},
			source:   3,
			target:   0,
			expected: nil,
		},
	}

	for _, tt := range tests {
		actual := GraphBreadthFirstSearch(tt.graph, tt.source, tt.target)

		if tt.expected == nil {
			Equal(t, actual, nil)
		} else {
			Equal(t, len(tt.expected), len(actual))

			for i, v := range tt.expected {
				Equal(t, v, actual[i])
			}
		}
	}
}
