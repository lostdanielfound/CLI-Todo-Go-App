package todo // Define a package called todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

// Completes a todo item given an index between 1-len(list)
func (t *Todos) Complete(index int) error {

	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("Invalid index: out of range of Todo list")
	}

	list[index-1].CompletedAt = time.Now()
	list[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {

	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("Invalid index: out of range of Todo list")
	}

	// Appending the Todos list up to [index - 1] and appending it to the list starting from [index] to the end.
	// [1 2 3] + [5 6 7]
	*t = append(list[:index-1], list[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	// Read in file
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	// Loading the JSON in a form of Todos list into t
	err = json.Unmarshal(filecontent, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {

	// Taking list of Todos t and converting it to an JSON stream
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {

	for i, item := range *t {
		i++
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}
