package main

import "os"
import "fmt"
import "strconv"

// go run 6.go N N N N

func main() {

	var banks []int
	seen := make(map[string]int)

	s := os.Args[1:]
	for _, element := range s {
		i, _ := strconv.Atoi(element)
		banks = append(banks, i)
	}

	fmt.Println(banks)

	count := 0
	first := 0
	for {
		count++

		// find largest (lowest indexed wins)
		largest := 0
		largestSize := -1
		for pos, size := range banks {
			if largestSize < size {
				largestSize = size
				largest = pos
			}
		}

		fmt.Println(largest, largestSize)

		// redistribute
		banks[largest] = 0
		for i := 0; i < largestSize; i++ {
			largest = (largest + 1) % len(banks)
			banks[largest]++
		}

		// laziness..
		keep := fmt.Sprint(banks)

		fmt.Println(count, keep)

		if seen[keep] == 1 && first == 0 {
			first = count
		}
		if seen[keep] == 2 {
			fmt.Println("total:", count, "first:", first, "second:", (count - first))
			os.Exit(0)
		}

		seen[keep]++

	}
}
