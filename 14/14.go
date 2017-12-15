package main

import (
	"fmt"
	"os"
	//"strconv"
)

// go run 14.go "input"

func main() {

	s := os.Args[1]
	ones := 0

	squareSize := 128
	grid := make([][]int, squareSize)

	for x := 0; x < squareSize; x++ {
		lineSeed := fmt.Sprint(s, "-", x)
		hash := MakeHash(lineSeed)
		ones += CountOnes(hash)
		grid[x] = BuildGrid(hash)
	}

	regionCount := 1
	for y := 0; y < squareSize; y++ {
		for x := 0; x < squareSize; x++ {
			grid, regionCount = CountRegions(grid, x, y, regionCount)
		}
	}

	fmt.Println(ones)
	fmt.Println(regionCount-1)

}

func CountRegions (grid [][]int, x int, y int, regionCount int) ([][]int, int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid) {
		return grid,regionCount
	}

	if grid[x][y] == -1 {
		return grid, regionCount
	}

	if grid[x][y] == 0 {
		grid[x][y] = regionCount
		grid,_ = CountRegions(grid, x+1, y, regionCount)
		grid,_ = CountRegions(grid, x-1, y, regionCount)
		grid,_ = CountRegions(grid, x, y+1, regionCount)
		grid,_ = CountRegions(grid, x, y-1, regionCount)
		regionCount++
	}
	
	return grid, regionCount
}

func BuildGrid (s string) []int {
	y := make([]int, len(s))

	for i := range s {
		if s[i] == 49 {
			y[i] = 0
		} else {
			y[i] = -1
		}
	}

	return y
}

func CountOnes (s string) int {
	count := 0
	for x := range s {
		if s[x] == 49 {
			count++
		}
	}
	return count
}


func MakeHash (s string) string {

	size := 256
	appendList := []byte{17, 31, 73, 47, 23}

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
			
			for y := 0; y < r; y++ {
				bottom := (pos+y) % size
				top := (pos+i-y-1) % size
				tmp := list[bottom]
				list[bottom] = list[top]
				list[top] = tmp
			}

			pos = (pos+i+skip) % size
			skip++
		}
	}

	ret := ""
	for a:=0; a<16; a++ {
		r := list[a*16]
		for b:=1; b<16; b++ {
			xorpos := (a*16)+b
			r = r ^ list[xorpos]
		}
		ret = ret + fmt.Sprintf("%08b", r) //+ " "
	}
	return ret

}
