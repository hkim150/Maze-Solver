package cmd

import (
	"fmt"
	"maze-solver/internal/generator"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().IntP("width", "w", 25, "Width of the maze")
	rootCmd.Flags().IntP("height", "l", 25, "Height of the maze")
	rootCmd.Flags().StringP("algorithm", "a", "dfs", "Algorithm to use for maze generation (wilson, dfs, kruskal, prim, aldous-broder, recursive-division, fractal, eller, hunt-and-kill, sidewinder, binary-tree)")
}

var rootCmd = &cobra.Command{
	// Use:   "gen-and-solve",
	Short: "Generate and solve the maze",
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")
		algorithm, _ := cmd.Flags().GetString("algorithm")

		maze, err := generator.Generate(width, height, algorithm)
		if err != nil {
			fmt.Println("Error generating maze:", err)
			return
		}

		maze.Print()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
