package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 24.go

type comp struct {
	a int
	b int
}

type bridge struct {
	components []comp
	last       int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var components = []comp{}

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "/")
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[1])
		c := comp{a, b}
		components = append(components, c)
	}

	b := bridge{[]comp{}, 0}
	bridges := []bridge{b}
	oldcount := -1

	for oldcount != 0 {
		nb := []bridge{}
		oldcount = 0
		for x := range bridges {
			c, tb := BuildBridges(bridges[x], components)
			oldcount += c
			nb = append(nb, tb...)
		}
		bridges = nb
	}

	maxlen := 0
	for p := range bridges {
		if len(bridges[p].components) > maxlen {
			maxlen = len(bridges[p].components)
		}
	}

	maxtot := 0
	maxlongest := 0
	for p := range bridges {
		tot := 0
		for x := range bridges[p].components {
			tot += bridges[p].components[x].a
			tot += bridges[p].components[x].b
		}
		if len(bridges[p].components) == maxlen && tot > maxlongest {
			maxlongest = tot
		}
		if tot > maxtot {
			maxtot = tot
		}
	}
	fmt.Println(maxtot, maxlongest)

}

func BuildBridges(b bridge, components []comp) (int, []bridge) {
	var bridges = []bridge{}

	used := 0
	for c := range components {

		tc := components[c]
		if isUsable(tc, b.components) && (tc.a == b.last || tc.b == b.last) {
			used++
			nb := make([]comp, len(b.components))
			copy(nb, b.components)
			nb = append(nb, tc)
			nlast := -1
			if tc.a == b.last {
				nlast = tc.b
			} else {
				nlast = tc.a
			}
			newbridge := bridge{nb, nlast}
			bridges = append(bridges, newbridge)
		}
	}
	if used == 0 {
		bridges = append(bridges, b)
	}

	return used, bridges
}

func isUsable(n comp, h []comp) bool {
	for c := range h {
		if n.a == h[c].a && n.b == h[c].b {
			return false
		}
	}
	return true
}
