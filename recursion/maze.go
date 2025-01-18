package recursion

type point struct {
	x, y int
}

var directions = []point{
	{x: -1, y: 0},
	{x: 1, y: 0},
	{x: 0, y: -1},
	{x: 0, y: 1},
}

func walkMaze(maze []string, wall byte, end, curr point, seen [][]bool, path *[]point) bool {
	// off the maze
	if curr.x < 0 || curr.x >= len(maze[0]) ||
		curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// hit the wall
	if maze[curr.y][curr.x] == wall {
		return false
	}

	// hit the end
	if curr == end {
		*path = append(*path, end)
		return true
	}

	// already seen
	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for _, dir := range directions {
		next := point{
			x: curr.x + dir.x,
			y: curr.y + dir.y,
		}

		if walkMaze(maze, wall, end, next, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func SolveMaze(maze []string, wall byte, start, end point) []point {
	seen := make([][]bool, len(maze))
	path := make([]point, 0)

	for i := range seen {
		seen[i] = make([]bool, len(maze[i]))
		for j := range seen[i] {
			seen[i][j] = false
		}
	}

	walkMaze(maze, wall, end, start, seen, &path)

	return path
}
