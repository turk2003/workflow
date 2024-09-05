package main

import (
	"fmt"
)

func main() {
	// ans := RemoveChar("turk")
	name := "Turk"
	name = "" + name[1:]

	fmt.Println(name)
}

// func RemoveChar(word string) string {
//   word[0] = ""
//   word[len(str)-1] = ""
//   return word
// }
