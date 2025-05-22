package algorithm

import (
	"math/rand"
	"maze-solver/internal/maze"
)

func AldousBroder(width, height int) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// notVisited is a map to keep track of the unvisited cells for end condition
	notVisited := make(map[maze.Pos]bool)
	for row := 1; row < m.Height-1; row += 2 {
		for col := 1; col < m.Width-1; col += 2 {
			notVisited[maze.Pos{row, col}] = true
		}
	}

	// Randomly select a starting cell from the isolated empty cells.
	startRow := rand.Intn((m.Height-1)/2)*2 + 1
	startCol := rand.Intn((m.Width-1)/2)*2 + 1
	m.Cells[startRow][startCol] = maze.Visited
	delete(notVisited, maze.Pos{startRow, startCol})

	for len(notVisited) > 0 {
		for _, dir := range randomDirections() {
			neighRow := startRow + dir[0]*2
			neighCol := startCol + dir[1]*2
			if neighRow >= 1 && neighRow < m.Height-1 && neighCol >= 1 && neighCol < m.Width-1 {
				if m.Cells[neighRow][neighCol] != maze.Visited {
					// remove the wall between the current cell and the neighbor
					m.Cells[startRow+dir[0]][startCol+dir[1]] = maze.Visited
					m.Cells[neighRow][neighCol] = maze.Visited
					delete(notVisited, maze.Pos{neighRow, neighCol})
				}
				startRow, startCol = neighRow, neighCol
				break
			}
		}
	}

	return m, nil
}
