package todo // Define a package called todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#id"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "Created-At"},
			{Align: simpletable.AlignCenter, Text: "Completed-At"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: item.Task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "Your todos are here"},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
