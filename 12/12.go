package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 12.go

type pipedata struct {
	pipe     int
	linkedTo []int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var pd = []pipedata{}

	for scanner.Scan() {

		s := strings.Split(scanner.Text(), " <-> ")
		pipe, _ := strconv.Atoi(s[0])

		var linkedTo = []int{}
		l := strings.Split(s[1], ", ")
		for t := range l {
			i, _ := strconv.Atoi(l[t])
			linkedTo = append(linkedTo, i)
		}

		pd = append(pd, pipedata{pipe, linkedTo})
	}

	cp := CountPipes(pd)
	seenGroups := make(map[int]int)
	groups, _ := CountGroups(seenGroups, pd)

	fmt.Println(cp, groups)
}

func CountPipes(pipedata []pipedata) int {
	seen := make(map[int]int)
	seen = getChildren(seen, pipedata, 0)
	fmt.Println(seen)

	count := 0
	for range seen {
		count++
	}

	return count

}

func CountGroups(seen map[int]int, pipedata []pipedata) (int, map[int]int) {
	count := 0
	for i := range pipedata {
		if seen[i] > 0 {
			continue
		} else {
			seen = getChildren(seen, pipedata, i)
			count++
		}
	}
	return count, seen
}

func getChildren(seen map[int]int, pipedata []pipedata, pipe int) map[int]int {

	p := pipedata[pipe]
	for l := range p.linkedTo {
		child := p.linkedTo[l]
		if seen[child] == 0 {
			seen[child]++
			seen = getChildren(seen, pipedata, child)
		}
	}

	return seen
}
