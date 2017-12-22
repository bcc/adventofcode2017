package main

import (
	"bufio"
	"fmt"
	"os"
)

// cat input | go run 22.go

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
				grid[coords] = 1
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

	for i:=0; i<10000; i++ {
		current := fmt.Sprint(posX, posY)
		node := grid[current]
		// Turn
		dir = Turn(dir, node)
		// Infect/clean
		if node == 1 {
			grid[current] = 0
		} else {
			grid[current] = 1
			infections++
		}
		// Move
		posX, posY = Move(posX, posY, dir)
	}
	fmt.Println(infections)

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

func Turn (dir string, infected int) string {
	if infected == 1 {
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
	return ""
}
