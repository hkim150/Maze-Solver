package internal

type MazeGenerator interface {
    Generate(width, height int) [][]int
}

type mazeGeneratorFunc func(width, height int) [][]int

func (f mazeGeneratorFunc) Generate(width, height int) [][]int {
    return f(width, height)
}

func randomizedDFS(width, height int) [][]int {
    return [][]int{}
}

func isolatedCells(width, height int) [][]int {
    
}