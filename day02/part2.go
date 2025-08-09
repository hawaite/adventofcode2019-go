package day02

import "fmt"

func part2(input string) error {
	target := 19690720
	vm := &intcodeVM{}

	err := vm.init(input)
	if err != nil {
		return err
	}

	for noun := range 100 {
		for verb := range 100 {
			vm.reset()
			vm.poke(1, noun)
			vm.poke(2, verb)
			vm.executeUntilHalt()

			result, err := vm.peek(0)
			if err != nil {
				return err
			}

			if result == target {
				fmt.Printf("noun: %d, verb: %d\n", noun, verb)
				fmt.Printf("result: %d\n", (100*noun)+verb)
				return nil
			}
		}
	}

	return fmt.Errorf("no combination found")
}
