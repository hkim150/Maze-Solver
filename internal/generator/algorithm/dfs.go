package algorithm

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
	"maze-solver/internal/maze"
)

func DFS(width, height int) (*maze.Maze, error) {
	// Create a base m filled with walls and isolated empty cells at odd rows and columns.
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// Randomly select a starting cell from the isolated empty cells.
	startRow := rand.Intn((m.Height-1)/2)*2 + 1
	startCol := rand.Intn((m.Width-1)/2)*2 + 1

	m.Cells[startRow][startCol] = maze.Visited
	stack := dataStructure.NewStack[maze.Pos]()
	stack.Push(maze.Pos{startRow, startCol})

	for !stack.IsEmpty() {
		cell, _ := stack.Pop()
		row, col := cell[0], cell[1]

		for _, dir := range randomDirections() {
			neighRow := row + dir[0]*2
			neighCol := col + dir[1]*2
			if neighRow >= 1 && neighRow < m.Height-1 && neighCol >= 1 && neighCol < m.Width-1 && m.Cells[neighRow][neighCol] != maze.Visited {
				m.Cells[row+dir[0]][col+dir[1]] = maze.Visited
				m.Cells[neighRow][neighCol] = maze.Visited
				stack.Push(maze.Pos{row, col})
				stack.Push(maze.Pos{neighRow, neighCol})
				break
			}
		}
	}

	return m, nil
}
