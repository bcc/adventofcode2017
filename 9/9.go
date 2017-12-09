package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cat input | go run 9.go

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		var keep []string

		// Remove garbage
		garbage := false
		garbagecount := 0
		for i := 0; i < len(s); i++ {

			switch s[i] {
			case "!":
				i++
				continue
			case "<":
				if garbage == true {
					fmt.Println("garbage:", s[i])
					garbagecount++
				}
				garbage = true
				continue
			case ">":
				garbage = false
				continue
			default:
				if !garbage {
					keep = append(keep, s[i])
				} else {
					fmt.Println("garbage:", s[i])
					garbagecount++
				}
			}
		}
		fmt.Println(keep)

		// Calculate depth
		depth := 0
		total := 0
		for i := 0; i < len(keep); i++ {
			switch keep[i] {
			case "{":
				depth++
				total += depth
			case "}":
				depth--
			}
		}
		fmt.Println(total, garbagecount)
	}	


}
