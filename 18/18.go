package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 18.go

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

	pc := 0
	lastsnd := -1

loop:
	for pc < len(instructions) && pc >= 0 {
		fmt.Println("PC: ", pc, instructions[pc])
		fmt.Println(registers, lastsnd)

		instruction := instructions[pc].instr
		register := instructions[pc].register

		_, present := registers[register]
		if !present {
			registers[register] = 0
		}

		switch instruction {
		case "set":
			pc = instrSet(&registers, instructions[pc], pc)
		case "add":
			pc = instrAdd(&registers, instructions[pc], pc)
		case "mul":
			pc = instrMul(&registers, instructions[pc], pc)
		case "mod":
			pc = instrMod(&registers, instructions[pc], pc)
		case "snd":
			pc = instrSnd(&registers, instructions[pc], pc, &lastsnd)
		case "rcv":
			pc = instrRcv(&registers, instructions[pc], pc, &lastsnd)
			break loop
		case "jgz":
			pc = instrJgz(&registers, instructions[pc], pc)
		default:
			fmt.Println("Instruction invalid!")
		}
		fmt.Println(registers, lastsnd)
		pc++

	}
	fmt.Println(registers, lastsnd)
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

func instrAdd(registers *map[string]int, i instruction, pc int) int {
	r := *registers
	if i.argIsInt {
		r[i.register] += i.argInt
	} else {
		r[i.register] += r[i.argReg]
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

func instrMod(registers *map[string]int, i instruction, pc int) int {
	r := *registers
	if i.argIsInt {
		r[i.register] %= i.argInt
	} else {
		r[i.register] %= r[i.argReg]
	}
	return pc
}

func instrSnd(registers *map[string]int, i instruction, pc int, snd *int) int {
	r := *registers
	*snd = r[i.register]
	fmt.Println("*snd:", *snd)
	return pc
}

func instrRcv(registers *map[string]int, i instruction, pc int, snd *int) int {
	r := *registers
	if i.argIsInt && i.argInt != 0 {
		r[i.register] = *snd
	}
	return pc
}

func instrJgz(registers *map[string]int, i instruction, pc int) int {
	r := *registers
	if i.argIsInt {
		if r[i.register] > 0 {
			pc = (pc + i.argInt) - 1
		}
	} else {
		if r[i.register] > 0 {
			pc = (pc + r[i.argReg]) - 1
		}
	}
	return pc
}
