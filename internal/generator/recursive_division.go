package generator

import (
	"math/rand"
)

func RecursiveDivision(width, height int) (Maze, error) {
	maze, err := emptyMaze(width, height)
	if err != nil {
		return maze, err
	}

	var division func(minRow, maxRow, minCol, maxCol int)
	division = func(minRow, maxRow, minCol, maxCol int) {
		if maxRow-minRow < 4 || maxCol-minCol < 4 {
			return
		}

		divRow := minRow + (1+rand.Intn((maxRow-minRow)/2-1))*2
		divCol := minCol + (1+rand.Intn((maxCol-minCol)/2-1))*2

		for col := minCol + 1; col < maxCol; col++ {
			maze.Cells[divRow][col] = Wall
		}

		for row := minRow + 1; row < maxRow; row++ {
			maze.Cells[row][divCol] = Wall
		}

		holeRow1 := minRow + rand.Intn((divRow-minRow)/2)*2 + 1
		holeRow2 := divRow + rand.Intn((maxRow-divRow)/2)*2 + 1
		holeCol1 := minCol + rand.Intn((divCol-minCol)/2)*2 + 1
		holeCol2 := divCol + rand.Intn((maxCol-divCol)/2)*2 + 1
		holeCells := [4][2]int{{holeRow1, divCol}, {holeRow2, divCol}, {divRow, holeCol1}, {divRow, holeCol2}}

		rand.Shuffle(4, func(i, j int) {
			holeCells[i], holeCells[j] = holeCells[j], holeCells[i]
		})

		for i := 0; i < 3; i++ {
			holeRow, holeCol := holeCells[i][0], holeCells[i][1]
			maze.Cells[holeRow][holeCol] = Empty
		}

		division(minRow, divRow, minCol, divCol)
		division(minRow, divRow, divCol, maxCol)
		division(divRow, maxRow, minCol, divCol)
		division(divRow, maxRow, divCol, maxCol)
	}

	division(0, maze.Height-1, 0, maze.Width-1)

	return maze, nil
}
