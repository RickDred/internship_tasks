package main

import (
	"fmt"

	"github.com/RickDred/internship_tasks/tree/second_task/module_example/internal/pakexample"
)

// place where program will start
func main() {
	fmt.Println("Please write your name:")
	var name string
	fmt.Scan(&name)
	pakexample.SayHello(name)
}
