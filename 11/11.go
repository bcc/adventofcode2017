package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// cat input | go run 11.go

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")

		x := 0
		ne := 0
		max := 0

		// I had a crack at this with an approximation of a cube based coordinate system
		// but the maths is hard and I suck. In the end I sought inspiration from
		// https://www.redblobgames.com/grids/hexagons/ - and am now using axial coordinates
		for dir := range s {
			switch s[dir] {
			case "n":
				ne++
			case "s":
				ne--
			case "se":
				x++
			case "sw":
				x--
				ne--
			case "ne":
				x++
				ne++
			case "nw":
				x--
			}

			// Inefficient, but I'm running out of lunchtime.
			cd := getDistance(x, ne)
			if cd > max {
				max = cd
			}

		}

		fmt.Println(getDistance(x, ne), max)

	}
}

func getDistance(x int, ne int) int {

	steps := 0
	for !(ne == 0 || x == 0) {

		switch {
		case ne > 0 && x > 0:
			steps++
			ne--
			x--

		case ne < 0 && x > 0:
			steps++
			ne++
		case ne > 0 && x < 0:
			steps++
			x++

		case ne < 0 && x < 0:
			steps++
			ne++
			x++

		}
	}
	return (steps + int(math.Abs(float64(x))) + int(math.Abs(float64(ne))))
}
