package cmd

import (
	"fmt"
	"maze-solver/internal/generator"
	"maze-solver/internal/maze"
	"maze-solver/internal/solver"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().IntP("width", "w", 25, "Width of the maze")
	rootCmd.Flags().IntP("height", "l", 25, "Height of the maze")
	rootCmd.Flags().StringP("generator", "g", "dfs", "maze generator algorithm - choose between [aldous-broder, binary-tree, dfs, eller, fractal-tessellation, hunt-and-kill, kruskal, prim, recursive-division, sidewinder, wilson]")
	rootCmd.Flags().StringP("solver", "s", "dfs", "maze solver algorithm - choose between [a-star, bfs, dead-end-filling, dfs, hand-on-wall, lee, pledge, random-mouse, recursive, tremaux]")
	rootCmd.Flags().BoolP("animate", "a", false, "Show generating and solving animation")
}

var rootCmd = &cobra.Command{
	Short: "Generate and solve the maze",
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")
		genAlgo, _ := cmd.Flags().GetString("generator")
		solverAlgo, _ := cmd.Flags().GetString("solver")
		animate, _ := cmd.Flags().GetBool("animate")

		m, err := generator.Generate(width, height, genAlgo, animate)
		if err != nil {
			fmt.Println("Error generating maze: ", err)
			return
		}

		// pause for a moment after showing the generated maze
		m.PrintForAnimation(1000 * time.Millisecond)
		
		err = solver.Solve(m, solverAlgo, animate)
		if err != nil {
			fmt.Println("Error solving maze: ", err)
			return
		}

		m.PrintForAnimation(0)
		
		// Ensure cursor is shown when animation finishes
		if animate {
			maze.ShowCursor()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
