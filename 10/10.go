package main

import (
	//"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

//  go run 10.go "n,n,n,n"

func main() {

	size := 256

	s := strings.Split(os.Args[1], ",")
	list := make([]int, size)

	for i := 0; i < size; i++ {
		list[i] = i
	}

	pos := 0
	skip := 0

	for x := range s {

		i,_ := strconv.Atoi(s[x])

		// Reverse the section
		r := i/2
		fmt.Println("i:", i, "range:", r)
		
		for y := 0; y < r; y++ {
			bottom := (pos+y) % size
			top := (pos+i-y-1) % size

			fmt.Println("pos:", pos, "i:", i, "y:", y, "bottom:", bottom, "top:", top)

			tmp := list[bottom]
			list[bottom] = list[top]
			list[top] = tmp
		}

		pos = (pos+i+skip) % size
		fmt.Println(pos, i, list)
		skip++
	}
	fmt.Println(list[0]*list[1])

}
