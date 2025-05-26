package algorithm

import (
	"fmt"
	"math/rand"
	"maze-solver/internal/maze"
)

// gridMaze returns a maze with grid like structure, where each empty cell is isolated and surrounded by walls
// The width and height must be at least 5 as the outer edge is a wall and a maze needs at least two cells for start and end
// width and height are round down to the nearest odd number to have 1 size padding on all four edges
func gridMaze(width, height int) (*maze.Maze, error) {
	m, err := blankMaze(width, height)
	if err != nil {
		return m, err
	}

	for r := 1; r < m.Height-1; r++ {
		for c := 1; c < m.Width-1; c++ {
			if r%2 != 1 || c%2 != 1 {
				m.Cells[r][c] = maze.Wall
			}
		}
	}

	return m, nil
}

// blankMaze generates a maze that only has the outer walls on the edge
// The width and height must be at least 5 as the outer edge is a wall and a maze needs at least two cells for start and end
// width and height are round down to the nearest odd number to have 1 size padding on all four edges
func blankMaze(width, height int) (*maze.Maze, error) {
	if width < 5 || height < 5 {
		return &maze.Maze{}, fmt.Errorf("width and height must be at least 5")
	}

	if width%2 == 0 {
		width--
	}

	if height%2 == 0 {
		height--
	}

	m := maze.NewMaze(width, height)

	for r := 0; r < m.Height; r++ {
		m.Cells[r][0] = maze.Wall
		m.Cells[r][m.Width-1] = maze.Wall
	}

	for c := 1; c < m.Width-1; c++ {
		m.Cells[0][c] = maze.Wall
		m.Cells[m.Height-1][c] = maze.Wall
	}

	return m, nil
}

// randomDirections returns a random order of directions in the maze (up, left, down, right)
func randomDirections() []maze.Pos {
	directions := []maze.Pos{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})
	return directions
}
