package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coords are y,x as first grid array index = y coord.
var lookpos = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

// cat input | go run 19.go

func main() {

	var grid []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())

	}

	cx, cy := findStart(grid)
	dx, dy := 0, 1

	steps := 0
	for cx != -1 && cy != -1 {

		switch grid[cy][cx] {
		case '|':
		case '-':
		case '+':
		default:
			fmt.Print(string(grid[cy][cx]))
		}

		cx, cy, dx, dy = findNext(grid, cx, cy, dx, dy)
		steps++
	}
	fmt.Println(steps - 1)

}

func findStart(grid []string) (int, int) {
	for x := range grid[0] {
		if grid[0][x] != ' ' {
			return x, 0
		}
	}
	return -1, -1
}

func findNext(grid []string, cx int, cy int, dx int, dy int) (int, int, int, int) {

	nextX := cx + dx
	nextY := cy + dy
	prevX := cx - dx
	prevY := cy - dy

	current := grid[cy][cx]

	if current == '+' {
		// Check for direction change
		for check := range lookpos {
			xp := cx + lookpos[check][1]
			yp := cy + lookpos[check][0]
			//fmt.Println("pos:", xp, yp)

			// Don't go backwards
			if xp == prevX && yp == prevY {
				continue
			}
			// Bounds checking
			if xp < 0 || xp >= len(grid[0]) {
				continue
			}
			if yp < 0 || yp >= len(grid) {
				continue
			}

			// Skip blanks.
			if grid[yp][xp] == ' ' {
				continue
			}
			// Set new direction
			dx = xp - cx
			dy = yp - cy
			//fmt.Println("dir:", dx, dy)
			return xp, yp, dx, dy
		}
	}

	// No direction change
	return nextX, nextY, dx, dy

}
