package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type particle struct {
	px int
	py int
	pz int
	vx int
	vy int
	vz int
	ax int
	ay int
	az int
}

// cat input | go run 20.go

func main() {

	var particles []particle

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t1 := strings.Split(scanner.Text(), ">")
		p1 := strings.Split(t1[0], "<")
		v1 := strings.Split(t1[1], "<")
		a1 := strings.Split(t1[2], "<")

		p := strings.Split(p1[1], ",")
		v := strings.Split(v1[1], ",")
		a := strings.Split(a1[1], ",")

		p = append(p, v...)
		p = append(p, a...)

		var keep []int
		for i := range p {
			x, _ := strconv.Atoi(strings.Trim(p[i], " "))
			keep = append(keep, x)
		}

		part := particle{keep[0], keep[1], keep[2], keep[3], keep[4], keep[5], keep[6], keep[7], keep[8]}
		particles = append(particles, part)
	}

	for {

		closest := 999999999999999999
		closestParticle := -1

		collide := make(map[string]int)

		for i := range particles {
			current := particles[i]
			current.vx += current.ax
			current.vy += current.ay
			current.vz += current.az
			current.px += current.vx
			current.py += current.vy
			current.pz += current.vz

			test := fmt.Sprint(current.px, current.py, current.pz)
			collide[test]++

			dist := int(math.Abs(float64(current.px)) + math.Abs(float64(current.py)) + math.Abs(float64(current.pz)))
			if dist < closest {
				closest = dist
				closestParticle = i
			}
			particles[i] = current
		}

		// Make new array of non-collision particles
		var new []particle
		for r := range particles {
			current := particles[r]
			test := fmt.Sprint(current.px, current.py, current.pz)
			t, _ := collide[test]
			if t < 2 {
				new = append(new, particles[r])
			}
		}
		particles = new

		fmt.Println(closestParticle, len(particles))
	}

}
