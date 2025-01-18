package recursion

import (
	"testing"
)

func TestSolveMaze(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	wall := byte('x')
	start := point{x: 10, y: 0}
	end := point{x: 1, y: 5}

	want := []point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	path := SolveMaze(maze, wall, start, end)

	for i := 0; i < len(path); i++ {
		if path[i] != want[i] {
			t.Errorf("SolveMaze() = %v, want %v", path, want)
		}
	}
}
