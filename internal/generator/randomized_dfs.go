package generator

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
)

func RandomizedDFS(width, height int) (Maze, error) {
	// Create a base maze filled with walls and isolated empty cells at odd rows and columns.
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// Randomly select a starting cell from the isolated empty cells.
	startRow := rand.Intn((maze.Height-1)/2)*2 + 1
	startCol := rand.Intn((maze.Width-1)/2)*2 + 1

	maze.Cells[startRow][startCol] = Visited
	stack := dataStructure.NewStack[[2]int]()
	stack.Push([2]int{startRow, startCol})

	for !stack.IsEmpty() {
		cell, _ := stack.Pop()
		row, col := cell[0], cell[1]

		for _, dir := range randomDirections() {
			neighRow := row + dir[0]*2
			neighCol := col + dir[1]*2
			if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && maze.Cells[neighRow][neighCol] != Visited {
				maze.Cells[row+dir[0]][col+dir[1]] = Visited
				maze.Cells[neighRow][neighCol] = Visited
				stack.Push([2]int{row, col})
				stack.Push([2]int{neighRow, neighCol})
				break
			}
		}
	}

	maze.Cells[maze.StartCell[0]][maze.StartCell[1]] = Start
	maze.Cells[maze.EndCell[0]][maze.EndCell[1]] = End

	return maze, nil
}
