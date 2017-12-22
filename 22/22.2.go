package main

import (
	"bufio"
	"fmt"
	"os"
)

// cat input | go run 22.2.go

func main() {

	grid := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)

	y := 0
	for scanner.Scan() {
		text := scanner.Text()
		for x := range text {
			coords := fmt.Sprint(x,y)
			fmt.Println(coords)
			if text[x] == '#' {
				grid[coords] = 2
			}
		}
		y++

	}

	posX := y/2
	posY := y/2
	dir := "u"
	infections := 0

	fmt.Println(grid)
	fmt.Println(posX, posY, dir)

	for i:=0; i<10000000; i++ {
		current := fmt.Sprint(posX, posY)
		node := grid[current]
		// Turn
		dir = Turn(dir, node)
		// Infect/clean
		result := ProcessNode(node)
		grid[current] = result
		if result == 2 {
			infections++
		}
		// Move
		posX, posY = Move(posX, posY, dir)
	}
	fmt.Println(infections)

}

func ProcessNode(node int) int {
	// 0 clean
	// 1 weakened
	// 2 infected
	// 3 flagged
	if node == 3 {
		node = 0
	} else {
		node++
	}
	return node
}

func Move (posX int, posY int, dir string) (int,int) {
	switch dir {
	case "u":
		posY--
	case "d":
		posY++
	case "l":
		posX--
	case "r":
		posX++
	}
	return posX, posY
}

func Turn (dir string, state int) string {

	// flagged
	if state == 3 {
		switch dir {
		case "u":
			return "d"
		case "d":
			return "u"
		case "l":
			return "r"
		case "r":
			return "l"
		}
	} 

	// infected
	if state == 2 {
		switch dir {
		case "u":
			return "r"
		case "d":
			return "l"
		case "l":
			return "u"
		case "r":
			return "d"
		}
	} 
 
	// clean (state 0)
	if state == 0 {
		switch dir {
		case "u":
			return "l"
		case "d":
			return "r"
		case "l":
			return "d"
		case "r":
			return "u"
		}
	}
	
	// weakened (state 1)
	return dir
}
