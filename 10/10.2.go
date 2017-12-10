package main

import (
	"fmt"
	"os"
)

// go run 10.2.go "input"

func main() {

	size := 256
	appendList := []byte{17, 31, 73, 47, 23}

	s := os.Args[1]
	s = s + string(appendList)

	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}

	pos := 0
	skip := 0

	for j := 0; j < 64; j++ {
		for x := range s {

			i := int(s[x])

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
	}

	for a:=0; a<16; a++ {

		r := list[a*16]
		for b:=1; b<16; b++ {
			xorpos := (a*16)+b
			//fmt.Println("xor:", xorpos, r, list[xorpos])
			r = r ^ list[xorpos]
		}

		fmt.Printf("%x", r)
		//fmt.Println(" ", r)
	}
	fmt.Println("")

}
