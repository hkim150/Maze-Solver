package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gen-and-solve",
	Short: "Generate and solve the maze",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Maze Generator and Solver!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
