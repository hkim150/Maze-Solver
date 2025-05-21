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
	rootCmd.Flags().StringP("algorithm", "a", "dfs", "Algorithm to use for maze generation (wilson, dfs, kruskal, prim, aldous-broder, recursive-division, fractal, eller, hunt-and-kill, sidewinder)")
}

var rootCmd = &cobra.Command{
	// Use:   "gen-and-solve",
	Short: "Generate and solve the maze",
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")
		algorithm, _ := cmd.Flags().GetString("algorithm")

		var maze generator.Maze
		var err error

		switch algorithm {
		case "dfs":
			maze, err = generator.RandomizedDFS(width, height)
		case "kruskal":
			maze, err = generator.RandomizedKruskal(width, height)
		case "prim":
			maze, err = generator.RandomizedPrim(width, height)
		case "wilson":
			maze, err = generator.WilsonsAlgorithm(width, height)
		case "aldous-broder":
			maze, err = generator.AldousBroder(width, height)
		case "recursive-division":
			maze, err = generator.RecursiveDivision(width, height)
		case "fractal":
			maze, err = generator.FractalTessellation(width, height)
		case "eller":
			maze, err = generator.EllersAlgorithm(width, height)
		case "hunt-and-kill":
			maze, err = generator.HuntAndKill(width, height)
		case "sidewinder":
			maze, err = generator.SidewinderAlgorithm(width, height)
		default:
			fmt.Println("Unknown algorithm:", algorithm)
			fmt.Println("Available algorithms: dfs, kruskal, prim, wilson, aldous-broder, recursive-division, fractal, eller, hunt-and-kill, sidewinder")
			return
		}

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
