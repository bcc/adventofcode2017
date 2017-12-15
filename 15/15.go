package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run 15.go inputA inputB

func main() {

	genA,_ := strconv.Atoi(os.Args[1])
	genB,_ := strconv.Atoi(os.Args[2])

	rounds1 := 40000000
	rounds2 := 5000000

	fmt.Println(Calc(genA, genB, 1,1,rounds1))
	fmt.Println(Calc(genA, genB, 4,8,rounds2))
}


func Calc (genA int, genB int, modA int, modB int, rounds int) int {
	genAFactor := 16807
	genBFactor := 48271
	div := 2147483647
	judge := 0

	for i := 0; i < rounds; i++ {

		for {
			genA = (genA * genAFactor) % div
			if genA % modA == 0 {
				break
			}
		}

		for {
			genB = (genB * genBFactor) % div
			if genB % modB == 0 {
				break
			}
		}

		tA := genA
		tA = tA % 65536

		tB := genB
		tB = tB % 65536

		if tA == tB {
			judge++
		}
	}
	return judge
}
