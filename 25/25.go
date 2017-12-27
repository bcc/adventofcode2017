package main

import (
	"fmt"
)

// go run 25.go

type instruction struct {
	state     string
	condition int
	write     int
	move      int
	newState  string
}

var production = []instruction{
	{"a", 0, 1, 1, "b"},
	{"a", 1, 0, -1, "d"},
	{"b", 0, 1, 1, "c"},
	{"b", 1, 0, 1, "f"},
	{"c", 0, 1, -1, "c"},
	{"c", 1, 1, -1, "a"},
	{"d", 0, 0, -1, "e"},
	{"d", 1, 1, 1, "a"},
	{"e", 0, 1, -1, "a"},
	{"e", 1, 0, 1, "b"},
	{"f", 0, 0, 1, "c"},
	{"f", 1, 0, 1, "e"},
}

func main() {
	fmt.Println(ProcessInstructions(production, 12317297))
}

func ProcessInstructions(instructions []instruction, count int) int {
	pos := 0
	state := instructions[0].state
	tape := make(map[int]int)

	for i := 0; i < count; i++ {

		for r := range instructions {
			inst := instructions[r]
			if inst.state == state && inst.condition == tape[pos] {
				state = inst.newState
				tape[pos] = inst.write
				pos += inst.move
				break
			}
		}
	}

	cksum := 0
	for _, v := range tape {
		if v == 1 {
			cksum++
		}
	}

	return cksum
}
