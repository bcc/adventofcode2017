package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type firewall struct {
	pos     int
	size    int
	current int
	up      bool
}

// cat input | go run 13.go

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var fwdata = make(map[int]firewall)

	for scanner.Scan() {

		s := strings.Split(scanner.Text(), ": ")
		pos, _ := strconv.Atoi(s[0])
		size, _ := strconv.Atoi(s[1])
		fwdata[pos] = firewall{pos, size - 1, 0, true}
	}

	fmt.Println(RunFirewall(fwdata, 0))
	// We only need to do this once with RunFirewallFast
	fwdata = ResetHash(fwdata)

	start := time.Now()

	fmt.Println(FindDelay(fwdata))

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}

func ResetHash(fw map[int]firewall) map[int]firewall {
	for k, v := range fw {
		v.current = 0
		fw[k] = v
	}
	return fw
}

func FindDelay(fw map[int]firewall) int {

	// Work out the final position.
	var keys []int
	for k := range fw {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	endpos := keys[len(keys)-1]

	// Using an array instead of hashmap drops from 300ms to 125 on my machine for my solution (3.96 million delay)
	var t []int
	for x := 0; x <= endpos; x++ {
		t = append(t, (fw[x].size * 2))
	}

Loop:
	for i := 0; ; i++ {
		// You can see earlier, slower (~300ms) version of this inner loop as a function below. This brings it down to ~100ms.
		for j := i; j <= i+endpos; j++ {
			p := t[j-i]
			if p != 0 && j%p == 0 {
				continue Loop
			}
		}
		return i
	}

	// slower hashmap version
	/*
		for i := 0; ; i++ {
			_, caught := RunFirewallFast(fw, i, endpos)
			if !caught {
				fmt.Println(i)
				break
			}
		}
	*/
}

// Well, the obvious solution is incredibly slow, and we don't need
// to resolve the full 'caught' number, so we can make this dramatically faster.
// It took less time to rewrite this than to run the result using RunFirewall...
func RunFirewallFast(fw map[int]firewall, delay int, endpos int) (int, bool) {
	for j := delay; j <= delay+endpos; j++ {
		a, exists := fw[j-delay]
		if exists {
			interval := a.size * 2
			if j%interval == 0 {
				return 0, true
			}
		}
	}

	return 0, false
}

func RunFirewall(fw map[int]firewall, delay int) (int, bool) {
	caught := 0
	caughtBool := false

	var keys []int
	for k := range fw {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	endpos := keys[len(keys)-1]
	mypos := -1

	for i := 0; ; i++ {

		// Start moving after the delay
		if i >= delay {
			mypos++
		}

		// Have we been caught?
		_, exists := fw[mypos]
		if exists && mypos >= 0 && fw[mypos].current == 0 {
			caught = caught + (fw[mypos].pos * (fw[mypos].size + 1))
			fmt.Println("Caught", mypos, caught, i, delay)
			caughtBool = true
			// Skip further rounds if we're just looking for the delay
			if delay > 0 {
				return 0, true
			}
		}

		// End!
		if mypos == endpos {
			break
		}

		// Move scanners for this round
		for _, k := range keys {
			c := fw[k]

			if c.current == c.size {
				c.up = false
			}
			if c.current == 0 {
				c.up = true
			}

			if c.up {
				c.current++
			} else {
				c.current--
			}

			fw[k] = c
		}

	}

	return caught, caughtBool

}
