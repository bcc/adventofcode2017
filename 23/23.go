package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 23.go

type instruction struct {
	instr    string
	register string
	argIsInt bool
	argInt   int
	argReg   string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	registers := make(map[string]int)
	var instructions = []instruction{}

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		i := instruction{s[0], s[1], false, 0, ""}

		if len(s) == 3 {
			int, err := strconv.Atoi(s[2])
			if err != nil {
				i.argIsInt = false
				i.argReg = s[2]
			} else {
				i.argIsInt = true
				i.argInt = int
			}
		}

		instructions = append(instructions, i)
	}

	reg := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for r := range reg {
		registers[reg[r]] = 0
	}

	pc := 0
	mulcount := 0

	for pc < len(instructions) && pc >= 0 {

		instruction := instructions[pc]

		/*
			fmt.Printf("%v: %8v %8v %8v %8v %8v %8v %8v %8v\n",pc, registers["a"],
				registers["b"],registers["c"],
				registers["d"],registers["e"],
				registers["f"],registers["g"],registers["h"])
		*/

		switch instruction.instr {
		case "set":
			pc = instrSet(&registers, instruction, pc)
		case "sub":
			pc = instrSub(&registers, instruction, pc)
		case "mul":
			pc = instrMul(&registers, instruction, pc)
			mulcount++
		case "jnz":
			pc = instrJnz(&registers, instruction, pc)
		default:
			fmt.Println("Instruction invalid!")
		}
		pc++
	}
	fmt.Println(pc, registers, mulcount, registers["h"])
}

func instrSet(registers *map[string]int, i instruction, pc int) int {
	r := *registers
	if i.argIsInt {
		r[i.register] = i.argInt
	} else {
		r[i.register] = r[i.argReg]
	}
	return pc
}

func instrSub(registers *map[string]int, i instruction, pc int) int {
	r := *registers
	if i.argIsInt {
		r[i.register] -= i.argInt
	} else {
		r[i.register] -= r[i.argReg]
	}
	return pc
}

func instrMul(registers *map[string]int, i instruction, pc int) int {
	r := *registers
	if i.argIsInt {
		r[i.register] *= i.argInt
	} else {
		r[i.register] *= r[i.argReg]
	}
	return pc
}

func instrJnz(registers *map[string]int, i instruction, pc int) int {
	r := *registers

	cmp := r[i.register]
	n, err := strconv.Atoi(i.register)
	if err == nil {
		cmp = n
	}

	if cmp != 0 {
		if i.argIsInt {
			pc = (pc + i.argInt) - 1
		} else {
			pc = (pc + r[i.argReg]) - 1
		}
	}

	return pc
}
