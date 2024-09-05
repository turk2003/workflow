package user

import (
	"fmt"
	"learn-go/pkg/todo"
)

func Greeting() {
	fmt.Println("Hello", name)
	fmt.Println(todo.Task)
	fmt.Println(Sum(1, 2))
}
