package main

import (
	"fmt"
)

func main() {
	var name string
	var m map[string]int = make(map[string]int)

	fmt.Scanln(&name)

	for i := 0; i < len(name); i++ {

		char := string(name[i])
		m[char]++
	}
	for char, count := range m {
		fmt.Printf("%s: %d\n", char, count)
	}

}
