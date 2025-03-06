package main

import (
	"fmt"
	"strconv"
)

func main() {
	// input := "LLRR=" // 210122
	// input := "==RLL" // 000210
	// input := "=LLRR" // 221012
	input := "RRL=R" // 012001

	min := zaiKwa(input)
	fmt.Println(min)
}

func zaiKwa(input string) string {
	runeInput := []rune(input)

	// init L 2 1 0
	numL := 0
	for i := len(runeInput) - 1; i >= 0; i-- {
		if runeInput[i] == 'L' {
			runeInput[i] = rune('0' + numL)
			numL++
		} else {
			numL = 0
		}
	}

	// init R 0 1 2
	numR := 0
	for i, v := range input {
		if v == 'R' {
			runeInput[i] = rune('0' + numR)
			numR++
		} else {
			numR = 0
		}
	}

	// check if 'R' is not more the left side
	modifiedInput := string(runeInput)
	for i := range modifiedInput {
		if input[i] == 'R' && i != 0 {
			prevVal, _ := strconv.Atoi(string(modifiedInput[i-1]))
			runeInput[i] = rune('0' + prevVal + 1)
			modifiedInput = string(runeInput)
		}

		if input[i] == 'L' && i != 0 {
			prevVal, _ := strconv.Atoi(string(modifiedInput[i-1]))
			currVal, _ := strconv.Atoi(string(modifiedInput[i]))

			if prevVal <= currVal {
				runeInput[i-1] = rune('0' + currVal + 1)
				modifiedInput = string(runeInput)
			}
		}
	}

	for i, v := range runeInput {
		if v == '=' && i == 0 {
			runeInput[i] = '0'
		} else if v == '=' {
			runeInput[i] = runeInput[i-1]
		}
	}

	modifiedInput = string(runeInput)
	if input[0] == 'L' {
		modifiedInput = strconv.Itoa(numL) + modifiedInput
	} else if input[0] == 'R' {
		lastR := 0
		foundR := false
		for i, v := range input {
			if v == 'R' {
				foundR = true
			}
			if v != 'R' && foundR {
				lastR = i - 1
			}
		}

		secondPart := modifiedInput[lastR:]

		firstPart := "0"
		foundR = false
		for i, v := range input {
			if v == 'R' {
				foundR = true
				firstPart += strconv.Itoa(i + 1)
			}
			if v != 'R' && foundR {
				break
			}
		}

		modifiedInput = firstPart + secondPart

		// modifiedInput = "0" + cumputed

	} else if input[0] == '=' {
		modifiedInput = string(modifiedInput[0]) + modifiedInput
	} else {
		modifiedInput = "0" + modifiedInput
	}

	return modifiedInput
}
