package main

import "os"
import "fmt"
import "strings"
import "strconv"

// go run 1.go <input>

func main() {

	s := strings.Split(os.Args[1], "")
	checksum := 0
	checksum2 := 0

	for index, element := range s {
		pre := index - 1
		pre2 := (index + (len(s) / 2)) % len(s)
		if pre == -1 {
			pre = len(s) - 1
		}

		fmt.Println(index, element, pre, s[pre])

		if element == s[pre] {
			i, _ := strconv.Atoi(element)
			checksum += i
		}
		if element == s[pre2] {
			i, _ := strconv.Atoi(element)
			checksum2 += i
		}

	}

	fmt.Println(checksum, checksum2)
}
