package main

import "os"
import "fmt"
import "strings"
import "strconv"
import "bufio"

// cat input | go run 2.go

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	checksum := 0
	for scanner.Scan() {

		min := 9999999999999 // messy, but whatever.
		max := 0
		s := strings.Split(scanner.Text(), "\t")

		for index, element := range s {
			fmt.Println(index, element)
			i, _ := strconv.Atoi(element)
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}

		diff := max - min
		checksum += diff
		fmt.Println("min", min, "max", max, "diff", diff)
		fmt.Println("---")

	}

	fmt.Println(checksum)

}
