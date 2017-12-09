package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 7.go

type node struct {
	NodeName   string
	NodeWeight int
	Children   []string
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

		fmt.Println(node)
		allNodes[node.NodeName] = node

	}

	fmt.Println(allNodes)

	// I expect this is going to bite us in part 2, but let's just repeatedly iterate
	// over the map until all the children are removed. I think this should be quicker
	// than traversing the tree repeatedly, given we don't know the root node.
	removed := make(map[string]bool)

	for len(allNodes) > 0 {
		for k, n := range allNodes {
			fmt.Println("node:", k)
			if len(n.Children) == 0 {
				fmt.Println("removing node", k)
				removed[k] = true
				delete(allNodes, k)

			} else {

				if len(allNodes[k].Children) > 0 {
					c := n.Children[0]
					if removed[c] == true {
						fmt.Println("len:", len(allNodes[k].Children))
						var new []string
						fmt.Println("nchildren:", n.Children)
						for j := 0; j < len(n.Children)-1; j++ {
							fmt.Println("j:", j)
							new = append(new, n.Children[j+1])
						}
						n.Children = new
						allNodes[k] = n

					}
				}

			}
		}
	}
	fmt.Println(allNodes)

}
