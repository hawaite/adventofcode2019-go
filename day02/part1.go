package day02

import (
	"fmt"
)

func part1(input string) error {
	vm := &intcodeVM{}

	err := vm.init(input)
	if err != nil {
		return err
	}

	vm.poke(1, 12)
	vm.poke(2, 2)
	vm.executeUntilHalt()

	result, err := vm.peek(0)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}
