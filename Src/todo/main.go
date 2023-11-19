package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/lostdanielfound/CLI-Todo-Go-App"
)

// Program const variables
const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(-1)
		}
	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(1)
	}
}
