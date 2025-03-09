package structures

import (
	"math"
	"slices"
)

type weightedAdjacencyMatrix = [][]int

func GraphBreadthFirstSearch(graph weightedAdjacencyMatrix, source, target int) []int {
	queue := NewQueue[int]()
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))

	for i := range prev {
		prev[i] = -1
	}
	seen[source] = true
	queue.Enqueue(source)

	for !queue.IsEmpty() {
		curr, ok := queue.Dequeue()
		if !ok {
			// Should never happen since the IsEmpty() check is done before the Dequeue()
			panic("Queue is empty")
		}

		if curr == target {
			break
		}

		adjs := graph[curr]
		for i, v := range adjs {
			if v == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			queue.Enqueue(i)
		}
	}

	curr := target
	out := make([]int, 0)

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) != 0 {
		slices.Reverse(out)
		return slices.Concat([]int{source}, out)
	}

	return nil
}

type graphEdge struct {
	to, weight int
}

type weightedAdjacencyList = [][]graphEdge

func walkGraph(graph weightedAdjacencyList, curr, target int, seen []bool, path *[]int) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true
	*path = append(*path, curr)

	if curr == target {
		return true
	}

	list := graph[curr]
	for _, edge := range list {
		if walkGraph(graph, edge.to, target, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func GraphDepthFirstSearch(graph weightedAdjacencyList, source, target int) []int {
	seen := make([]bool, len(graph))
	path := make([]int, 0)

	walkGraph(graph, source, target, seen, &path)

	if len(path) != 0 {
		return path
	}

	return nil
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i, s := range seen {
		if !s && dists[i] < math.MaxInt64 {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDist := math.MaxInt64

	for i, s := range seen {
		if s {
			continue
		}

		if lowestDist > dists[i] {
			lowestDist = dists[i]
			idx = i
		}
	}

	return idx
}

func DijkstraList(graph weightedAdjacencyList, source, target int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	dists := make([]int, len(graph))

	for i := range dists {
		dists[i] = math.MaxInt64
		prev[i] = -1
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := graph[curr]
		for _, edge := range adjs {
			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + edge.weight
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	curr := target
	out := make([]int, 0)

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)

	return out
}
