package main

import "fmt"

func main() {
	var m map[string]int = make(map[string]int)
	// Insert
	m["Alice"] = 2
	m["Bob"] = 3
	// GET, check key
	value, ok := m["Bob"]
	if ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key does not exists")
	}
	fmt.Println(value)
	// Delete
	delete(m, "Alice")
	fmt.Println(m)
	for key, value := range m {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
