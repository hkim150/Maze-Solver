package algorithm

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
	"maze-solver/internal/maze"
	"time"
)

func Prim(width, height int, animate bool) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// directionsToUnvisitedNeighbors is a helper function that returns the directions to the unvisited neighbor
	directionsToUnvisitedNeighbors := func(row, col int) []maze.Pos {
		directions := []maze.Pos{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		unvisitedDir := make([]maze.Pos, 0)
		for _, dir := range directions {
			neighRow := row + dir[0]*2
			neighCol := col + dir[1]*2
			if neighRow >= 1 && neighRow < m.Height-1 && neighCol >= 1 && neighCol < m.Width-1 && m.Cells[neighRow][neighCol] != maze.Visited {
				m.Cells[neighRow][neighCol] = maze.Visited
				unvisitedDir = append(unvisitedDir, dir)
			}
		}

		return unvisitedDir
	}

	row := rand.Intn(m.Height/2)*2 + 1
	col := rand.Intn(m.Width/2)*2 + 1
	m.Cells[row][col] = maze.Visited

	// initialize a randomized set to store frontier walls.
	// each wall is represented as [row, col, rowDir, colDir], where rowDir and colDir
	// indicate the direction to the neighboring cell.
	rs := dataStructure.NewRandomizedSet[[4]int]()
	dirs := directionsToUnvisitedNeighbors(row, col)
	for _, dir := range dirs {
		rs.Add([4]int{row, col, dir[0], dir[1]})
	}

	var delay time.Duration
	if animate {
		delay = 40 * time.Millisecond
		m.PrintForAnimation(delay)
	}

	for !rs.IsEmpty() {
		// Randomly select a frontier wall from the pool
		elem, _ := rs.GetRandom()
		rs.Remove(elem)

		// Remove the wall
		row, col, rowDir, colDir := elem[0], elem[1], elem[2], elem[3]
		if animate {
			m.Cells[row+rowDir][col+colDir] = maze.Highlight
			m.PrintForAnimation(delay)
		}
		m.Cells[row+rowDir][col+colDir] = maze.Visited

		// mark the neighboring cell as visited and add the walls in between to the pool
		dirs := directionsToUnvisitedNeighbors(row+rowDir*2, col+colDir*2)
		for _, dir := range dirs {
			rs.Add([4]int{row + rowDir*2, col + colDir*2, dir[0], dir[1]})
		}
	}

	return m, nil
}
