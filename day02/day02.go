package day02

import (
	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed inputs/input.txt
var inputText string

//go:embed inputs/example.txt
var exampleInputText string

func NewDay02Cmd() *cobra.Command {
	day02Cmd := &cobra.Command{
		Use:     "day02",
		Aliases: []string{"2"},
		Short:   "Day 02",
	}

	day02Cmd.AddCommand(newPart1Cmd(inputText, exampleInputText))
	day02Cmd.AddCommand(newPart2Cmd(inputText, exampleInputText))

	day02Cmd.PersistentFlags().BoolP("useExampleInput", "e", false, "Use the example input instead of the real input.")
	return day02Cmd
}

func newPart1Cmd(inputText string, exampleInputText string) *cobra.Command {
	part1Cmd := &cobra.Command{
		Use:     "part1",
		Aliases: []string{"1"},
		Short:   "Day 01 Part 1",
		RunE: func(cmd *cobra.Command, args []string) error {
			useExampleInput, _ := cmd.Flags().GetBool("useExampleInput")
			if useExampleInput {
				return part1(exampleInputText)
			} else {
				return part1(inputText)
			}
		},
	}
	return part1Cmd
}

func newPart2Cmd(inputText string, exampleInputText string) *cobra.Command {
	part2Cmd := &cobra.Command{
		Use:     "part2",
		Aliases: []string{"2"},
		Short:   "Day 02 Part 2",
		RunE: func(cmd *cobra.Command, args []string) error {
			useExampleInput, _ := cmd.Flags().GetBool("useExampleInput")
			if useExampleInput {
				return part2(exampleInputText)
			} else {
				return part2(inputText)
			}
		},
	}
	return part2Cmd
}
