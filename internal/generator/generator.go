package generator

import (
	"fmt"
	"maze-solver/internal/generator/algorithm"
	"maze-solver/internal/maze"
	"strings"
)

var generators = map[string]generatorFunc{
	"dfs":                algorithm.DFS,
	"kruskal":            algorithm.Kruskal,
	"prim":               algorithm.Prim,
	"wilson":             algorithm.Wilson,
	"aldous-broder":      algorithm.AldousBroder,
	"recursive-division": algorithm.RecursiveDivision,
	"fractal":            algorithm.FractalTessellation,
	"eller":              algorithm.Ellers,
	"hunt-and-kill":      algorithm.HuntAndKill,
	"sidewinder":         algorithm.Sidewinder,
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

		return &maze.Maze{}, fmt.Errorf("Unknown generator algorithm; Choose one from: %v\n", algorithms)
	}

	m, err := genFunc.Generate(width, height, animate)
	if err != nil {
		return m, err
	}

	m.CleanUp()
	return m, nil
}
