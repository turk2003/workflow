package main

import "fmt"

func main() {
	// var p *int
	// v := 42
	// p = &v
	// fmt.Printf("Memory Address %v\n", p)
	// fmt.Printf("Value %v\n", *p)
	// *p = 41
	// fmt.Printf("Value of p %v\n", *p)
	// fmt.Printf("Value of v %v\n", v)
	var a int = 1
	increase(&a)
	increase(&a)
	increase(&a)
	fmt.Printf("a: %v\n", a)
}
func increase(a *int) { // clone
	*a++
}
