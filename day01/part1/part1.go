package part1

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateFuel(moduleMass int) int {
	fuel := (moduleMass / 3) - 2
	return max(fuel, 0)
}

func Run(inputText string) {
	lines := strings.Split(inputText, "\n")

	total := 0
	for i := range lines {
		value, _ := strconv.Atoi(lines[i])
		total += calculateFuel(value)
	}

	fmt.Printf("total: %d\n", total)
}
