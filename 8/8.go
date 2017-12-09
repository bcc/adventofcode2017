package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cat input | go run 8.go

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	registers := make(map[string]int)
	maxEverVal := 0

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		register := s[0]
		instruction := s[1]
		data, _ := strconv.Atoi(s[2])
		condition := s[3]
		conditionRegister := s[4]
		conditionTest := s[5]
		conditionData, _ := strconv.Atoi(s[6])

		_, present := registers[register]
		if !present {
			registers[register] = 0
		}

		_, present = registers[conditionRegister]
		if !present {
			registers[conditionRegister] = 0
		}

		fmt.Println(s)
		fmt.Println(registers)

		pass := false
		if condition == "if" {
			switch conditionTest {
			case "==":
				pass = registers[conditionRegister] == conditionData
			case ">":
				pass = registers[conditionRegister] > conditionData
			case "<":
				pass = registers[conditionRegister] < conditionData
			case "<=":
				pass = registers[conditionRegister] <= conditionData
			case ">=":
				pass = registers[conditionRegister] >= conditionData
			case "!=":
				pass = registers[conditionRegister] != conditionData
			default:
				fmt.Println("Condition failed!", s)
			}

			if pass {
				switch instruction {
				case "inc":
					registers[register] += data
				case "dec":
					registers[register] -= data
				default:
					fmt.Println("Invalid instruction", s)
				}

				if registers[register] > maxEverVal {
					maxEverVal = registers[register]
				}
			}

		} else {
			fmt.Println("Invalid condition", s)
		}
	}

	maxVal := 0
	maxReg := ""
	for k, v := range registers {
		if v > maxVal {
			maxVal = v
			maxReg = k
		}
	}
	fmt.Println(registers)
	fmt.Println(maxVal, maxReg, maxEverVal)

}
