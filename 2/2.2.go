package main

import "os"
import "fmt"
import "strings"
import "strconv"
import "bufio"

// cat input | go run 2.2.go

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	checksum := 0
	for scanner.Scan() {

		s := strings.Split(scanner.Text(), "\t")
		fmt.Println(s)

		for index, element := range s {
			fmt.Println(index, element)
			i, _ := strconv.Atoi(element)

			for index2, e2 := range s {

				i2, _ := strconv.Atoi(e2)

				if index == index2 {
					continue
				}
				if i2 > i {
					continue
				}

				result := float64(i) / float64(i2)
				ir := float64(int(result))

				if result == ir {
					fmt.Println("yes!", i, i2, result)
					checksum += int(result)
				}

			}
		}

		fmt.Println("---")

	}

	fmt.Println(checksum)
}
