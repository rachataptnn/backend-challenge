package main

import (
	"fmt"
	"strconv"
)

func main() {
	// input := "LLRR=" // 210122
	// input := "==RLL" // 000210
	input := "=LLRR" // 221012
	// input := "RRL=R" // 012001

	min := zaiKwa(input)
	fmt.Println(min)
}

func zaiKwa(encoded string) string {
	runeEncoded := []rune(encoded)

	lastNumL := convertLToNums(runeEncoded)
	convertRToNums(runeEncoded)
	remapLRNums(runeEncoded, encoded)
	convertEqToNums(runeEncoded)
	result := mapFirstCharToNum(runeEncoded, encoded, lastNumL)

	return result
}

// init L 2 1 0
func convertLToNums(runeEncoded []rune) int {
	lastNumL := 0
	for i := len(runeEncoded) - 1; i >= 0; i-- {
		if runeEncoded[i] == 'L' {
			runeEncoded[i] = rune('0' + lastNumL)
			lastNumL++
		} else {
			lastNumL = 0
		}
	}

	return lastNumL
}

// init R 0 1 2
func convertRToNums(runeEncoded []rune) int {
	numR := 0
	for i := 0; i < len(runeEncoded); i++ {
		if runeEncoded[i] == 'R' {
			runeEncoded[i] = rune('0' + numR)
			numR++
		} else {
			numR = 0
		}
	}

	return numR
}

func remapLRNums(runeEncoded []rune, encoded string) {
	for i := 0; i < len(runeEncoded); i++ {
		if encoded[i] == 'R' && i != 0 {
			prevVal := runeEncoded[i-1]
			if prevVal == '=' {
				runeEncoded[i] = '1'
			} else {
				runeEncoded[i] = prevVal + 1
			}
		}

		if encoded[i] == 'L' && i != 0 {
			prevVal := runeEncoded[i-1]
			currVal := runeEncoded[i]
			if prevVal <= currVal || prevVal == '=' {
				runeEncoded[i-1] = currVal + 1
			}
		}
	}
}

func convertEqToNums(runeEncoded []rune) {
	for i, v := range runeEncoded {
		if v == '=' && i == 0 {
			runeEncoded[i] = '0'
		} else if v == '=' {
			runeEncoded[i] = runeEncoded[i-1]
		}
	}
}

func mapFirstCharToNum(runeEncoded []rune, encoded string, lastNumL int) string {
	modifiedInput := string(runeEncoded)

	switch encoded[0] {
	case 'L':
		modifiedInput = strconv.Itoa(lastNumL) + modifiedInput
	case 'R':
		lastR := 0
		foundR := false
		for i, v := range encoded {
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
		for i, v := range encoded {
			if v == 'R' {
				foundR = true
				firstPart += strconv.Itoa(i + 1)
			}
			if v != 'R' && foundR {
				break
			}
		}

		modifiedInput = firstPart + secondPart
	case '=':
		modifiedInput = string(modifiedInput[0]) + modifiedInput
	default:
		modifiedInput = "0" + modifiedInput
	}

	return modifiedInput
}
