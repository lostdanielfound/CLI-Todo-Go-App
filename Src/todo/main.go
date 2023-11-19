package main

import (
	"flag"
	"fmt"
)

// Program const variables
const (
	todoFile = ".todos.json"
)

func main() {
	fmt.Println("Hello")
	add := flag.Bool("add", false, "add a new todo")

	todos := &todo.Todos{}

}
