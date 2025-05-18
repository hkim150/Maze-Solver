package cmd

import (
	"fmt"
	"maze-solver/internal/generator"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().IntP("width", "w", 20, "Width of the maze")
	rootCmd.Flags().IntP("height", "l", 20, "Height of the maze")
	rootCmd.Flags().StringP("algorithm", "a", "wilson", "Algorithm to use for maze generation (wilson, dfs, kruskal, prim, aldous-broder)")
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
		default:
			fmt.Println("Unknown algorithm:", algorithm)
			fmt.Println("Available algorithms: dfs, kruskal, prim, wilson, aldous-broder")
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
