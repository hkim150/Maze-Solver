package generator

import (
	"fmt"
	"maze-solver/internal/generator/algorithm"
	"maze-solver/internal/maze"
	"strings"
)

var generators = map[string]generatorFunc{
	"dfs":                algorithm.DFS,
	// "kruskal":            algorithm.Kruskal,
	// "prim":               algorithm.Prim,
	// "wilson":             algorithm.Wilson,
	"aldous-broder":      algorithm.AldousBroder,
	// "recursive-division": algorithm.RecursiveDivision,
	"fractal":            algorithm.FractalTessellation,
	"eller":              algorithm.Ellers,
	// "hunt-and-kill":      algorithm.HuntAndKill,
	// "sidewinder":         algorithm.Sidewinder,
	"binary-tree":        algorithm.BinaryTree,
}

type generatorFunc func(width, height int, animate bool) (*maze.Maze, error)

func (f generatorFunc) Generate(width, height int, animate bool) (*maze.Maze, error) {
	return f(width, height, animate)
}

func Generate(width, height int, algorithm string, animate bool) (*maze.Maze, error) {
	genFunc, ok := generators[algorithm]
	if !ok {
		keys := make([]string, 0, len(generators))
		for key := range generators {
			keys = append(keys, key)
		}
		algorithms := strings.Join(keys, ", ")

		return &maze.Maze{}, fmt.Errorf("Unknown algorithm; Choose one from: %v\n", algorithms)
	}

	m, err := genFunc.Generate(width, height, animate)
	if err != nil {
		return m, err
	}

	prepareForSolving(m)
	return m, nil
}

func prepareForSolving(m *maze.Maze) {
	for r := 1; r < m.Height-1; r++ {
		for c := 1; c < m.Width-1; c++ {
			if m.Cells[r][c] != maze.Wall {
				m.Cells[r][c] = maze.Empty
			}
		}
	}

	m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Start
	m.Cells[m.EndPos[0]][m.EndPos[1]] = maze.End
}
