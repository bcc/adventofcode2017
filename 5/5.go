package main

import "fmt"
import "bufio"
import "os"
import "strconv"

// cat input | go run 5.go

func main() {

	var instructions []int
	i := 0
	count := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		instructions = append(instructions, i)
	}

	for {

		if i >= len(instructions) {
			fmt.Println(count)
			os.Exit(0)
		}

		count++
		current := instructions[i]

		if instructions[i] < 3 {
			instructions[i]++ // Just this one for the first solution
		} else {
			instructions[i]--
		}
		i += current

	}

}
