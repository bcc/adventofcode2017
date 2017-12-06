package main

import "os"
import "fmt"
import "strings"
import "bufio"
import "sort"

// cat input | go run 4.go

func main() {

	failed := 0
	count := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		count++
		s := strings.Split(scanner.Text(), " ")
		fmt.Println(s)

		sort.Strings(s)
		for i := 1; i < len(s); i++ {
			fmt.Println(i, s[i])
			if s[i] == s[i-1] {
				failed++
				fmt.Println("Failed!")
				break
			}
		}

		fmt.Println("---")
	}
	fmt.Println(failed, count, count-failed)
}
