package user

import (
	"fmt"
	"lern-go/pkg/todo"
)

func Greeting() {
	fmt.Println("Hello", name)
	fmt.Println(todo.Task)

}
