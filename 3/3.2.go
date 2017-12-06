package main

import "os"
import "fmt"
import "math"
import "strconv"

// go run 3.2.go <input>

// Need to brute force part two anyway, should have done that for part one...

func main() {

	target, _ := strconv.Atoi(os.Args[1])

	// This is way overkill for the real solution, but needed it for testing initial grid size with value = i
	squareSize := int(math.Ceil(math.Sqrt(float64(target)))) + 2
	grid := make([][]int, squareSize)

	for i := 0; i < squareSize; i++ {
		y := make([]int, squareSize)
		grid[i] = y
	}

	dirX := 1
	dirY := 0
	startVal := 1
	posX := (squareSize - 1) / 2
	posY := (squareSize - 1) / 2
	grid[posX][posY] = 1

	fmt.Println("sq", squareSize, "xy", posX, posY, "dir", dirX, dirY)

	for i := startVal + 1; i <= target; i++ {

		posX += dirX
		posY += dirY

		adj := grid[posX-1][posY-1] + grid[posX-1][posY] + grid[posX-1][posY+1] +
			grid[posX][posY-1] + grid[posX][posY+1] +
			grid[posX+1][posY-1] + grid[posX+1][posY] + grid[posX+1][posY+1]

		//adj = i

		grid[posX][posY] = adj
		if adj > target {
			fmt.Println("Result:", adj)
			os.Exit(0)
		}

		if dirY == -1 && grid[posX+1][posY] == 0 {
			dirY = 0
			dirX = 1
		}
		if dirY == 1 && grid[posX-1][posY] == 0 {
			dirY = 0
			dirX = -1
		}
		if dirX == -1 && grid[posX][posY-1] == 0 {
			dirY = -1
			dirX = 0
		}
		if dirX == 1 && grid[posX][posY+1] == 0 {
			dirY = 1
			dirX = 0
		}

	}

	//fmt.Println(grid)

}
