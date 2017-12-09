package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 7.2.go

type node struct {
	NodeName    string
	NodeWeight  int
	Children    []string
	ChildWeight int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	allNodes := make(map[string]node)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " -> ")
		var children []string

		if len(s) == 2 {
			// Node has children
			children = strings.Split(s[1], ", ")
		}

		// Name and Weight
		tNode := strings.Split(s[0], " ")
		node := node{NodeName: tNode[0]}
		node.NodeWeight, _ = strconv.Atoi(strings.TrimRight(strings.TrimLeft(tNode[1], "("), ")"))
		node.Children = children
		allNodes[node.NodeName] = node
	}

	for _, n := range allNodes {
		n.ChildWeight = weighChildren(allNodes, n.NodeName)

	}

	fmt.Println(allNodes)

}

func weighChildren(allNodes map[string]node, child string) int {

	if len(allNodes[child].Children) == 0 {
		return allNodes[child].NodeWeight
	}

	sum := allNodes[child].NodeWeight
	check := 0
	uneven := false

	for _, e := range allNodes[child].Children {

		r := weighChildren(allNodes, e)
		sum += r

		if check == 0 {
			check = r
		}

		if check != r {
			uneven = true
		}
	}

	f := allNodes[child]
	f.ChildWeight = sum
	allNodes[child] = f

	if uneven {
		for _, e := range allNodes[child].Children {
			fmt.Println(e, allNodes[e].ChildWeight, allNodes[e].NodeWeight)
		}
		os.Exit(0)
	}
	return sum
}
