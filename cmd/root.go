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
}

var rootCmd = &cobra.Command{
	Use:   "gen-and-solve",
	Short: "Generate and solve the maze",
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")

		maze, err := generator.WilsonsAlgorithm(width, height)
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
