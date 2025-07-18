package main

import (
	"os"

	"github.com/hawaite/aoc19/day01"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc19",
	Short: "Runner for Advent of Code 2019",
}

func main() {
	rootCmd.AddCommand(day01.NewDay01Cmd())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
