package main

import "fmt"

func main() {
	fmt.Printf("result: %T\n", sum)
	fmt.Printf("compute result: %v\n", compute(sum))
	fmt.Printf("compute result: %v\n", compute(sub))
}
func sum(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func compute(operation func(int, int) int) int {
	return operation(10, 20)
}
