package main

import "fmt"

func main() {
	fmt.Printf("result: %v\n", sum(1, 2, 3))
	fmt.Printf("type sum: %T\n", sum)
	fmt.Printf("type x: %T\n", x)
	// fmt.Printf("result: %v\n", sum(1, 2,))
}
func sum(nums ...int) int {
	return nums[0] + nums[1] + nums[2]
}
func x(a, b, c int) []int {
	return []int{1, 2, 3, 4}
}
