package main

import "fmt"

func main() {
	var age int
	fmt.Print("your age : ")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("ไม่บรรลุ")
	} else if age >= 18 && age <= 60 {
		fmt.Println("วัยทำงาน")
	} else {
		fmt.Println("วัยเกษียร")
	}
}
