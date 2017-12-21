package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type rule struct {
	input  []string
	output []string
}

// cat input | go run 21.go

func main() {

	iterations := 18
	var rules []rule
	grid := []string{".#.", "..#", "###"}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " => ")
		i := strings.Split(t[0], "/")
		o := strings.Split(t[1], "/")

		r := rule{}
		r.output = o
		r.input = i
		rules = append(rules, r)
		r.input = FlipH(i)
		rules = append(rules, r)
		r.input = FlipV(i)
		rules = append(rules, r)

		for j := 1; j <= 3; j++ {
			r := rule{}
			i = Rotate90(i)
			r.output = o
			r.input = i
			rules = append(rules, r)
			r.input = FlipH(i)
			rules = append(rules, r)
			r.input = FlipV(i)
			rules = append(rules, r)
		}

	}

	size := 0
	for i := 0; i < iterations; i++ {
		fmt.Println("============ iter:", i)
		fmt.Println("grid:", grid)
		grid, size = Embiggen(grid)
		fmt.Println("embiggened:", grid, size)
		grid = Cromulate(grid, size, rules)
		fmt.Println("cromulated:", grid, size)
	}

	countstr := strings.Join(grid, "")
	count := 0
	for i := range countstr {
		if countstr[i] == '#' {
			count++
		}
	}
	fmt.Println(count)

}

func Cromulate(grid []string, osize int, rules []rule) []string {
	size := osize + 1
	//steps := (len(grid) / size) - 1

	//fmt.Println("cromulating:", osize, len(grid), steps)

	for yo := 0; yo < len(grid); yo += size {
		for xo := 0; xo < len(grid); xo += size {
			//fmt.Println("crom:", xo, yo)

			var match []string
			for i := 0; i < osize; i++ {
				m := grid[yo+i][xo : xo+osize]
				match = append(match, m)
			}
			result := FindMatch(match, rules)
			//fmt.Println("match:", match, result)

			for y := 0; y < len(result); y++ {
				yarr := []byte(grid[yo+y])

				for x := 0; x < len(result); x++ {
					yarr[xo+x] = byte(result[y][x])
				}
				//fmt.Println("swap:", grid[yo+y], string(yarr))
				grid[yo+y] = string(yarr)
			}

		}
	}

	return grid
}

func FindMatch(match []string, rules []rule) []string {
	m := strings.Join(match, "")
	for r := range rules {
		c := strings.Join(rules[r].input, "")
		if m == c {
			return rules[r].output
		}
	}
	return match
}

func Embiggen(grid []string) ([]string, int) {
	var out []string

	size := 3
	if len(grid)%2 == 0 {
		size = 2
	}

	for y := 0; y < len(grid); y++ {
		s := ""
		for x := 0; x < len(grid[y]); x++ {
			b := grid[y][x]
			s += string(b)
			if x%size == size-1 {
				s += "X"
			}
		}
		out = append(out, s)

		if y%size == size-1 {
			tmpstr := "X"
			for ts := 0; ts < len(out[0])-1; ts++ {
				tmpstr += "X"
			}
			out = append(out, tmpstr)
		}
	}

	return out, size
}

func Rotate90(in []string) []string {
	var out []string
	for y := 0; y < len(in); y++ {
		s := ""

		for x := len(in[y]) - 1; x >= 0; x-- {
			b := in[x][y]
			s += string(b)

		}

		out = append(out, s)
	}
	return out
}

func FlipH(in []string) []string {
	var out []string
	for y := 0; y < len(in); y++ {
		s := ""

		for x := len(in[y]) - 1; x >= 0; x-- {
			b := in[y][x]
			s += string(b)

		}

		out = append(out, s)
	}
	return out
}

func FlipV(in []string) []string {
	var out []string

	for y := len(in) - 1; y >= 0; y-- {
		s := ""

		for x := 0; x < len(in); x++ {
			b := in[y][x]
			s += string(b)

		}

		out = append(out, s)
	}
	return out
}
