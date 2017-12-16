package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	instr    byte
	arg1    int
	arg2    int
	n1	byte
	n2	byte
}

// cat input | go run 16.go

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		instr := ParseInstr(strings.Split(scanner.Text(), ","))

		input := "abcdefghijklmnop"
		seen := make(map[string]string)

		fmt.Println(string(Run16(input, instr)))
		rounds := 1000000000

		for i := 1; i <= rounds; i++ {
			if i % (rounds/10) == 0 {
				fmt.Print("/")
			} else if i % (rounds/100) == 0 {
				fmt.Print(".")
			}
			
			// Cache known results. Suspect we could cache repeating patterns, but this is fast enough.
			cache, exists := seen[input]
			if exists {
				input = cache
			} else {
				pre := input
				input = Run16(input,instr)
				seen[pre] = input
			}
		}

		fmt.Println()
		fmt.Println(input)
	}
}

func Run16(i string, instr []instruction) string {
	input := []byte(i)
	//fmt.Println(i)
	for i := range instr {
		if instr[i].instr == 's' {
			input = Spin(input, instr[i].arg1)
		}
		if instr[i].instr == 'x' {
			input = SwapPos(input, instr[i].arg1, instr[i].arg2)
		}
		if instr[i].instr == 'p' {
			input = SwapName(input, instr[i].n1, instr[i].n2)
		}
	}
	return string(input)
}

func Spin(input []byte, pos int) []byte {
	//fmt.Println("Spin", input, pos)
	t := make([]byte, len(input))

	p := 0
	for i := len(input)-pos; i < len(input); i++ {
		t[p] = input[i]
		p++
	}
	for i := 0; i < len(input)-pos; i++ {
		t[p] = input[i]
		p++
	}

	return t
}

func SwapPos(ia []byte, a int, b int) []byte {
	//fmt.Println("SwapPos", input, args)
	ia[a],ia[b] = ia[b],ia[a]
	return ia
}

func SwapName(ia []byte, s1 byte, s2 byte) []byte {
	//fmt.Println("SwapName", ia, s1, s2)
	a:= -1
	b:= -1

	for i := 0; i < len(ia); i++ {
		if ia[i] == s1 {
			a = i
		}
		if ia[i] == s2 {
			b = i
		}
		if a != -1 && b != -1 {
			break
		}
	}

	ia[a],ia[b] = ia[b],ia[a]
	return ia
}

func ParseInstr(l []string) []instruction {
	var instructions []instruction
	for i := range l {
		var res = instruction{0, 0, 0, 0, 0}

		res.instr = l[i][0]

		if l[i][0] == 's' {
			res.arg1,_ = strconv.Atoi( l[i][1:len(l[i])] )
		}
		if l[i][0] == 'x' {
			s := strings.Split(l[i][1:len(l[i])], "/")
			res.arg1,_ = strconv.Atoi(s[0])
			res.arg2,_ = strconv.Atoi(s[1])
		}
		if l[i][0] == 'p' {
			s := strings.Split(l[i][1:len(l[i])], "/")
			res.n1 = s[0][0]
			res.n2 = s[1][0]
		}
		instructions = append(instructions, res)
	}
	return instructions
}