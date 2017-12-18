package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 18.2.go

type instruction struct {
	instr    string
	register string
	argIsInt bool
	argInt   int
	argReg   string
}

var sends int

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	registers0 := make(map[string]int)
	registers1 := make(map[string]int)
	var instructions = []instruction{}
	var q0 = []int{}
	var q1 = []int{}

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		i := instruction{s[0], s[1], false, -1, ""}

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

	pc0 := 0
	pc1 := 0
	registers0["p"] = 0
	registers1["p"] = 1

	for pc0 < len(instructions) && pc0 >= 0 && pc1 < len(instructions) && pc1 >= 0 {
		pcpre0, pcpre1 := pc0, pc1
		pc0 = runInstr(0, &registers0, instructions[pc0], pc0, &q0, &q1)
		pc1 = runInstr(1, &registers1, instructions[pc1], pc1, &q1, &q0)
		if pc0 == pcpre0 && pc1 == pcpre1 {
			fmt.Println("deadlock!", sends)
			break
		}
	}
	fmt.Println(sends)
}

func runInstr(p int, registers *map[string]int, i instruction, pc int, qin *[]int, qout *[]int) int {
	//fmt.Println(p, "PC:", pc, i, *registers, qin, qout)
	register := i.register

	// Skip register if this is an int, because 'jgz 1 3'
	_, err := strconv.Atoi(register)
	if err != nil {
		r := *registers
		_, present := r[register]
		if !present {
			r[register] = 0
		}
	}

	switch i.instr {
	case "set":
		pc = instrSet(registers, i, pc)
	case "add":
		pc = instrAdd(registers, i, pc)
	case "mul":
		pc = instrMul(registers, i, pc)
	case "mod":
		pc = instrMod(registers, i, pc)
	case "snd":
		pc = instrSnd(registers, i, pc, qout)
		if p == 1 {
			sends++
		}
	case "rcv":
		pc = instrRcv(registers, i, pc, qin)
	case "jgz":
		pc = instrJgz(registers, i, pc)
	default:
		fmt.Println("Instruction invalid!")
	}
	pc++
	//fmt.Println(p, "PC:", pc, i, *registers, qin, qout)
	//fmt.Println()
	return pc

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

func instrSnd(registers *map[string]int, i instruction, pc int, snd *[]int) int {
	r := *registers
	n, err := strconv.Atoi(i.register)
	if err == nil {
		*snd = append(*snd, n)
	} else {
		*snd = append(*snd, r[i.register])
	}
	return pc
}

func instrRcv(registers *map[string]int, i instruction, pc int, snd *[]int) int {
	r := *registers
	s := *snd
	if len(s) == 0 {
		pc--
	} else {
		x, t := s[0], s[1:]
		r[i.register] = x
		*snd = t
	}
	return pc
}

func instrJgz(registers *map[string]int, i instruction, pc int) int {
	r := *registers

	cmp := r[i.register]
	n, err := strconv.Atoi(i.register)
	if err == nil {
		cmp = n
	}

	if cmp > 0 {
		if i.argIsInt {
			pc = (pc + i.argInt) - 1
		} else {
			pc = (pc + r[i.argReg]) - 1
		}
	}

	return pc
}
