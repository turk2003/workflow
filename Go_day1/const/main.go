package main

import "fmt"

const name = "Turk"
const (
	male = iota
	female
)

func main() {

	fmt.Printf("Hello %v\n", name)
	fmt.Printf("male %v\n", male)
	fmt.Printf("female %v\n", female)
}
