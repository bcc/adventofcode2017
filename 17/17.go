package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run 17.go input

func main() {

	steps,_ := strconv.Atoi(os.Args[1])
	iterations1 := 2017 
	iterations2 := 50000000

	fmt.Println(Part1(steps, iterations1))
	fmt.Println(Part2(steps, iterations2))
}

func Part1(steps int, iterations int) int {
	buf := make([]int, iterations+1)
	buf[0] = 0
	pos := 0

	for i:=1; i<=iterations; i++ {
		pos = (pos + steps + 1) % i

		// Move items from top down, insert value into array
		for j:=i; j > pos; j-- {
			buf[j] = buf[j-1]
		}
		buf[pos+1] = i

	}
	return buf[pos+2]
}

func Part2(steps int, iterations int) int {
	buf := make([]int, iterations+1)
	buf[0] = 0
	pos := 0
	valAfterZero := -1

	for i:=1; i<=iterations; i++ {
		pos = (pos + steps + 1) % i

		// And zero is always first position in the list, so no need to fiddle with the array...
		if (pos == 0) {
			valAfterZero = i
		}

	}
	return valAfterZero
}
