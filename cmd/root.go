package cmd

import (
	"fmt"
	"maze-solver/internal/generator"
	"maze-solver/internal/solver"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().IntP("width", "w", 25, "Width of the maze")
	rootCmd.Flags().IntP("height", "l", 25, "Height of the maze")
	rootCmd.Flags().StringP("algorithm", "a", "dfs", "Algorithm to use for maze generation (wilson, dfs, kruskal, prim, aldous-broder, recursive-division, fractal, eller, hunt-and-kill, sidewinder, binary-tree)")
	rootCmd.Flags().BoolP("animate", "s", false, "Show generation animation")
}

var rootCmd = &cobra.Command{
	Short: "Generate and solve the maze",
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")
		algorithm, _ := cmd.Flags().GetString("algorithm")
		animate, _ := cmd.Flags().GetBool("animate")

		m, err := generator.Generate(width, height, "dfs", false)
		// m, err := generator.Generate(width, height, algorithm, animate)
		if err != nil {
			fmt.Println("Error generating maze:", err)
			return
		}

		if animate {
			m.PrintForAnimation(2 * time.Second)
		}

		err = solver.Solve(m, algorithm, animate)
		if err != nil {
			fmt.Println("Error solving maze", err)
			return
		}

		if animate {
			m.PrintForAnimation(0 * time.Second)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
