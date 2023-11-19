package todo // Define a package called todo

import (
	"errors"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item // Creates a static public list of list of item(s)

// Appends a todo item
func (t *Todos) Add(task string) {

	todo := item{ // init a new todo
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

// Completes a todo item
func (t *Todos) Complete(index int) error {

	list := *t
	if index < 0 || index > len(list) {
		return errors.New("Invalid index: out of range of Todo list")
	}

	list[index-1].CompletedAt = time.Now()
	list[index-1].Done = true

	return nil
}
