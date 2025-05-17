package generator

type MazeGenerator interface {
	Generate(width, height int) Maze
}

type mazeGeneratorFunc func(width, height int) Maze

func (f mazeGeneratorFunc) Generate(width, height int) Maze {
	return f(width, height)
}
