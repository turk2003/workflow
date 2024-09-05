package main

import "fmt"

var (
	name   string
	weight float64
)

func main() {
	setup()
	const age = 20
	fmt.Println("ชื่อ", name)
	fmt.Println("อายุ", age, "ปี")
	fmt.Println("น้ำหนัก", weight)

}
func setup() {
	name = "turk"
	weight = 80
}
