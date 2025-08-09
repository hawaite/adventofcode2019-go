package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type intcodeVM struct {
	initmemory []int
	memory     []int
	ip         int
}

func (vm *intcodeVM) reset() {
	vm.memory = make([]int, len(vm.initmemory))
	copy(vm.memory, vm.initmemory)
	vm.ip = 0
}

func (vm *intcodeVM) peek(addr int) (int, error) {
	if addr >= len(vm.memory) {
		return 0, fmt.Errorf("attempted acess to addr %d on memory of length %d", addr, len(vm.memory))
	}
	return vm.memory[addr], nil
}

func (vm *intcodeVM) peekIp() (int, error) {
	if vm.ip >= len(vm.memory) {
		return 0, fmt.Errorf("attempted acess to addr %d on memory of length %d", vm.ip, len(vm.memory))
	}
	return vm.memory[vm.ip], nil
}

func (vm *intcodeVM) poke(addr int, value int) error {
	if addr >= len(vm.memory) {
		return fmt.Errorf("attempted acess to addr %d on memory of length %d", addr, len(vm.memory))
	}

	vm.memory[addr] = value
	return nil
}

func (vm *intcodeVM) executeSingleOp() (opWasHalt bool, err error) {
	val, err := vm.peekIp()
	if err != nil {
		return
	}
	switch val {
	case 99:
		opWasHalt = true
	case 1:
		vm.performAddition()
	case 2:
		vm.performMultiplication()
	}
	return
}

func (vm *intcodeVM) executeUntilHalt() error {
	for {
		opWasHalt, err := vm.executeSingleOp()
		if err != nil {
			return err
		}
		if opWasHalt {
			return nil
		}
	}
}

func (vm *intcodeVM) performAddition() {
	sourceOneAddr := vm.memory[vm.ip+1]
	sourceTwoAddr := vm.memory[vm.ip+2]
	destAddr := vm.memory[vm.ip+3]

	sourceOneVal := vm.memory[sourceOneAddr]
	sourceTwoVal := vm.memory[sourceTwoAddr]

	vm.memory[destAddr] = sourceOneVal + sourceTwoVal

	vm.ip += 4
}

func (vm *intcodeVM) performMultiplication() {
	sourceOneAddr := vm.memory[vm.ip+1]
	sourceTwoAddr := vm.memory[vm.ip+2]
	destAddr := vm.memory[vm.ip+3]

	sourceOneVal := vm.memory[sourceOneAddr]
	sourceTwoVal := vm.memory[sourceTwoAddr]

	vm.memory[destAddr] = sourceOneVal * sourceTwoVal

	vm.ip += 4
}

func (vm *intcodeVM) init(initialMemory string) error {
	vm.ip = 0
	var ram []int
	for val := range strings.SplitSeq(initialMemory, ",") {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("value %s could not be converted to int", val)
		}
		ram = append(ram, intVal)
		vm.initmemory = ram
	}

	vm.reset()
	return nil
}
