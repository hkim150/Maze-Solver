package cmd

import (
	"fmt"
	"maze-solver/internal"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen-and-solve",
	Short: "Generate and solve the maze",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Maze Generator and Solver!")
		maze, err := internal.RandomizedDFS(20, 20)
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
