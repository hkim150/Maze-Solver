package internal

import (
	"fmt"
	"math/rand"
)

type MazeGenerator interface {
	Generate(width, height int) Maze
}

type mazeGeneratorFunc func(width, height int) Maze

func (f mazeGeneratorFunc) Generate(width, height int) Maze {
	return f(width, height)
}

func RandomizedDFS(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	startRow := rand.Intn((maze.Height-1)/2)*2 + 1
	startCol := rand.Intn((maze.Width-1)/2)*2 + 1

	maze.Cells[startRow][startCol] = Visited
	stack := [][2]int{{startRow, startCol}}

	for len(stack) > 0 {
		row, col := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		directions := [][2]int{{-2, 0}, {0, -2}, {2, 0}, {0, 2}}
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})
		for _, dir := range directions {
			neighRow := row + dir[0]
			neighCol := col + dir[1]
			if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && maze.Cells[neighRow][neighCol] != Visited {
				maze.Cells[row+dir[0]/2][col+dir[1]/2] = Visited
				maze.Cells[neighRow][neighCol] = Visited
				stack = append(stack, [2]int{row, col})
				stack = append(stack, [2]int{neighRow, neighCol})
				break
			}
		}
	}

	maze.Cells[1][1] = Start
	maze.Cells[maze.Height-2][maze.Width-2] = End

	return maze, nil
}

// baseMaze generates a maze with grid like structure
// where all cells are walls except for the odd row and odd column cells are empty
// Because the outer edges are walls, the width and height must be at least 5
// Also the width and height will be round down to the nearest odd number
func baseMaze(width, height int) (Maze, error) {
	if width < 5 || height < 5 {
		return Maze{}, fmt.Errorf("width and height must be at least 5")
	}

	if width%2 == 0 {
		width--
	}
	if height%2 == 0 {
		height--
	}

	cells := make([][]CellType, height)
	for i := range cells {
		cells[i] = make([]CellType, width)
		for j := range cells[i] {
			if i%2 == 1 && j%2 == 1 {
				cells[i][j] = Empty
			} else {
				cells[i][j] = Wall
			}
		}
	}

	maze := Maze{
		Width:  width,
		Height: height,
		Cells:  cells,
	}

	return maze, nil
}
