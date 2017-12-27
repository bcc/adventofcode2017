package main

import (
	"fmt"
)

/*
	This was my third attempt at solving this - the first time trying to work out the
	pattern of the registers from interpreting the 'assembly'. I gave up on that and
	turned my input instructions into something closer to very messy go. That didn't
	quite work either, I had to figure out the register values properly and restructure
	the loops resulting in the final version below. I found this the most challenging of
	the puzzles so far - not sure I'd have finished this without some of the hints from
	the AoC subreddit.
*/

func main() {
	f := 0
	h := 0

	for b := 106500; b <= 123500; b += 17 {
		f = 1
		for d := 2; d <= b; d++ {
			for e := 2; e <= b; e++ {
				if b%d != 0 { // thanks to AoC reddit for this hint
					break
				}
				if d*e == b {
					f = 0
				}
			}
		}
		if f == 0 {
			h += 1
		}
	}
	fmt.Println(h)
}
