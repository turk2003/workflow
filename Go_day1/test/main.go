package main

import (
	"fmt"
	"strings"
)

func countLetters(input string) map[rune]int {
	counts := make(map[rune]int)
	input = strings.ToUpper(input)

	for _, char := range input {
		counts[char]++
	}

	return counts
}

func main() {
	var input string
	fmt.Print("Input: ")
	fmt.Scanln(&input)

	counts := countLetters(input)

	for char, count := range counts {
		fmt.Printf("%c: %d\n", char, count)
	}
}
