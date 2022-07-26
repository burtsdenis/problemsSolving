package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	result := 0
	romanSymbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		str := string(s[i])
		if i + 1 < len(s) && romanSymbols[str] < romanSymbols[string(s[i + 1])] {
			result -= romanSymbols[str]
		} else {
			result += romanSymbols[str]
		}
	}

	return result
}
