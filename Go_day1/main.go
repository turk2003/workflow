package main

import (
	"fmt"
	"learn-go/pkg/user"
)

// Main function
func main() {
	fmt.Println("Hello Go")
	user.Greeting()
}

// Multiple init() function
func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello! init() function")
}
