package day01

import (
	_ "embed"
	"fmt"

	"github.com/hawaite/aoc19/day01/part1"
	"github.com/hawaite/aoc19/day01/part2"
	"github.com/spf13/cobra"
)

//go:embed inputs/input.txt
var inputText string

//go:embed inputs/example.txt
var exampleInputText string

func NewDay01Cmd() *cobra.Command {
	day01Cmd := &cobra.Command{
		Use:     "day1",
		Aliases: []string{"1"},
		Short:   "Day 01",
	}

	day01Cmd.AddCommand(newPart1Cmd(inputText, exampleInputText))
	day01Cmd.AddCommand(newPart2Cmd(inputText, exampleInputText))

	day01Cmd.PersistentFlags().BoolP("useExampleInput", "e", false, "Use the example input instead of the real input.")
	return day01Cmd
}

func newPart1Cmd(inputText string, exampleInputText string) *cobra.Command {
	part1Cmd := &cobra.Command{
		Use:     "part1",
		Aliases: []string{"1"},
		Short:   "Day 01 Part 1",
		Run: func(cmd *cobra.Command, args []string) {
			useExampleInput, _ := cmd.Flags().GetBool("useExampleInput")
			fmt.Printf("useExampleInput: %t\n", useExampleInput)
			if useExampleInput {
				part1.Run(exampleInputText)
			} else {
				part1.Run(inputText)
			}
		},
	}

	return part1Cmd
}

func newPart2Cmd(inputText string, exampleInputText string) *cobra.Command {
	part2Cmd := &cobra.Command{
		Use:     "part2",
		Aliases: []string{"2"},
		Short:   "Day 01 Part 2",
		Run: func(cmd *cobra.Command, args []string) {
			useExampleInput, _ := cmd.Flags().GetBool("useExampleInput")
			if useExampleInput {
				part2.Run(exampleInputText)
			} else {
				part2.Run(inputText)
			}
		},
	}

	return part2Cmd
}
