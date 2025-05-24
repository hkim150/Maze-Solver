package algorithm

import (
	"math/rand"
	"maze-solver/internal/maze"
	"time"
)

func RecursiveDivision(width, height int, animate bool) (*maze.Maze, error) {
	m, err := blankMaze(width, height)
	if err != nil {
		return m, err
	}
	var delay time.Duration
	if animate {
		delay = 80 * time.Millisecond
		m.PrintForAnimation(delay)
	}

	var division func(minRow, maxRow, minCol, maxCol int)
	division = func(minRow, maxRow, minCol, maxCol int) {
		if maxRow-minRow < 4 || maxCol-minCol < 4 {
			return
		}

		// randomly choose row and col to place walls that divide the area into four sections
		divRow := minRow + (1+rand.Intn((maxRow-minRow)/2-1))*2
		divCol := minCol + (1+rand.Intn((maxCol-minCol)/2-1))*2

		// create a horizontal wall
		for col := minCol + 1; col < maxCol; col++ {
			m.Cells[divRow][col] = maze.Wall
		}

		// create a vertical wall
		for row := minRow + 1; row < maxRow; row++ {
			m.Cells[row][divCol] = maze.Wall
		}

		if animate {
			m.PrintForAnimation(delay)
		}

		// randomly choose four locations on each of the wall to make a hole
		holeRow1 := minRow + rand.Intn((divRow-minRow)/2)*2 + 1
		holeRow2 := divRow + rand.Intn((maxRow-divRow)/2)*2 + 1
		holeCol1 := minCol + rand.Intn((divCol-minCol)/2)*2 + 1
		holeCol2 := divCol + rand.Intn((maxCol-divCol)/2)*2 + 1
		holeCells := [4]maze.Pos{{holeRow1, divCol}, {holeRow2, divCol}, {divRow, holeCol1}, {divRow, holeCol2}}

		// randomly choose three out of four hole locations and create holes
		rand.Shuffle(4, func(i, j int) {
			holeCells[i], holeCells[j] = holeCells[j], holeCells[i]
		})

		for i := 0; i < 3; i++ {
			holeRow, holeCol := holeCells[i][0], holeCells[i][1]
			if animate {
				m.Cells[holeRow][holeCol] = maze.Highlight
				m.PrintForAnimation(delay)
			}
			m.Cells[holeRow][holeCol] = maze.Empty
		}

		// recursively repeat the process on each of the four sub-sections
		division(minRow, divRow, minCol, divCol)
		division(minRow, divRow, divCol, maxCol)
		division(divRow, maxRow, minCol, divCol)
		division(divRow, maxRow, divCol, maxCol)
	}

	division(0, m.Height-1, 0, m.Width-1)

	return m, nil
}
