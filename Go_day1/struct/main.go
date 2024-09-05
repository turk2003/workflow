package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Greeting() {
	fmt.Printf("Hello, %v\n", u.Name)
}
func main() {
	u := User{
		Name: "Bas",
		Age:  18,
	}
	fmt.Println(u)
	u.Age = 19
	fmt.Println(u)
	fmt.Println(u.Name)
	// new
	a := &User{} // ใช้บ่อย
	fmt.Println(a)
	b := new(User)
	fmt.Println(b)
	c := newUser() // ใช้บ่อย ⭐️
	fmt.Println(c)
	u.Greeting()
}
func newUser() *User {
	return &User{}
}
