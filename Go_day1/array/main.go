package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	a = append(a, 2, 3)
	fmt.Println(a)
	a = a[:len(a)-1]
	fmt.Println(a)
	name := []string{}
	name = append(name, "Turk", "Tan")
	for i, value := range name {
		fmt.Println("id", i, "name", value)

	}
}
