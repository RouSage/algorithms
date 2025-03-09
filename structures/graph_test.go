package structures

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestGraphMatrixBFS(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	graph1 := GraphMatrix{
		{0, 3, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}
	// (0)-->(3)
	//  ^\     ^
	//  | \    |
	//  v  \   v
	// (1)  ->(2)
	graph2 := GraphMatrix{
		{0, 20, 19, 18},
		{17, 0, 16, 0},
		{0, 0, 0, 15},
		{0, 0, 14, 0},
	}

	tests := []struct {
		graph    GraphMatrix
		source   int
		target   int
		expected []int
	}{
		{
			graph:    graph1,
			source:   0,
			target:   6,
			expected: []int{0, 1, 4, 5, 6},
		},
		{
			graph:    graph1,
			source:   6,
			target:   0,
			expected: nil,
		},
		{
			graph:    graph1,
			source:   3,
			target:   1,
			expected: nil,
		},
		{
			graph: graph2, source: 0,
			target:   3,
			expected: []int{0, 3},
		},
		{
			graph:    graph2,
			source:   3,
			target:   0,
			expected: nil,
		},
	}

	for _, tt := range tests {
		actual := tt.graph.BreadthFirstSearch(tt.source, tt.target)

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

func TestGraphListDFS(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	graph := GraphList{
		{
			graphEdge{to: 1, weight: 3},
			graphEdge{to: 2, weight: 1},
		},
		{
			graphEdge{to: 4, weight: 1},
		},
		{
			graphEdge{to: 3, weight: 7},
		},
		{},
		{
			graphEdge{to: 1, weight: 1},
			graphEdge{to: 3, weight: 5},
			graphEdge{to: 5, weight: 2},
		},
		{
			graphEdge{to: 2, weight: 18},
			graphEdge{to: 6, weight: 1},
		},
		{
			graphEdge{to: 3, weight: 1},
		},
	}

	tests := []struct {
		graph    GraphList
		source   int
		target   int
		expected []int
	}{
		{
			graph:    graph,
			source:   0,
			target:   6,
			expected: []int{0, 1, 4, 5, 6},
		},
		{
			graph:    graph,
			source:   6,
			target:   0,
			expected: nil,
		},
		{
			graph:    graph,
			source:   3,
			target:   2,
			expected: nil,
		},
	}

	for _, tt := range tests {
		actual := tt.graph.DepthFirstSearch(tt.source, tt.target)

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

func TestDijkstraList(t *testing.T) {
	//      (1) --- (4) ---- (5)
	//    /  |       |       /|
	// (0)   | ------|------- |
	//    \  |/      |        |
	//      (2) --- (3) ---- (6)
	graph := GraphList{
		{
			graphEdge{to: 1, weight: 3},
			graphEdge{to: 2, weight: 1},
		},
		{
			graphEdge{to: 0, weight: 3},
			graphEdge{to: 2, weight: 4},
			graphEdge{to: 4, weight: 1},
		},
		{
			graphEdge{to: 1, weight: 4},
			graphEdge{to: 3, weight: 7},
			graphEdge{to: 0, weight: 1},
		},
		{
			graphEdge{to: 2, weight: 7},
			graphEdge{to: 4, weight: 5},
			graphEdge{to: 6, weight: 1},
		},
		{
			graphEdge{to: 1, weight: 1},
			graphEdge{to: 3, weight: 5},
			graphEdge{to: 5, weight: 2},
		},
		{
			graphEdge{to: 6, weight: 1},
			graphEdge{to: 4, weight: 2},
			graphEdge{to: 2, weight: 18},
		},
		{
			graphEdge{to: 3, weight: 1},
			graphEdge{to: 5, weight: 1},
		},
	}

	tests := []struct {
		graph    GraphList
		source   int
		target   int
		expected []int
	}{
		{
			graph:    graph,
			source:   0,
			target:   6,
			expected: []int{0, 1, 4, 5, 6},
		},
		{
			graph:    graph,
			source:   4,
			target:   1,
			expected: []int{4, 1},
		},
		{
			graph:    graph,
			source:   5,
			target:   0,
			expected: []int{5, 4, 1, 0},
		},
	}

	for _, tt := range tests {
		actual := tt.graph.Dijkstra(tt.source, tt.target)

		Equal(t, len(tt.expected), len(actual))

		for i, v := range tt.expected {
			Equal(t, v, actual[i])
		}
	}
}
