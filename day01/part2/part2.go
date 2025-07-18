package part2

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateFuel(moduleMass int) int {
	fuel := (moduleMass / 3) - 2
	return max(fuel, 0)
}

func recursiveCalculateFuel(mass int) int {
	if calculateFuel(mass) != 0 {
		return calculateFuel(mass) + recursiveCalculateFuel(calculateFuel(mass))
	}
	return 0
}

func Run(inputText string) {
	lines := strings.Split(inputText, "\n")

	total := 0
	for i := range lines {
		value, _ := strconv.Atoi(lines[i])
		total += recursiveCalculateFuel(value)
	}

	fmt.Printf("total: %d\n", total)
}
