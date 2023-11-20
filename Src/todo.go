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

/**
 * Name: Add
 * Desc: appends to the current list of items
 *
 * return: nothing
 */
func (t *Todos) Add(task string) {

	todo := item{ // init a new todo
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

/**
 * Name: Complete
 * Desc: completes a item within the list, index (int) that is passed
 * must be an available index within the list from [1 - N] where N
 * is the total number of items within the Todos list
 *
 * return: error on failure / nil on success
 */
func (t *Todos) Complete(index int) error {

	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index: out of range of Todo list")
	}

	list[index-1].CompletedAt = time.Now()
	list[index-1].Done = true

	return nil
}

/**
 * Name: Uncomplete
 * Desc: sets the done flag from true to false
 *
 * return: error on failure / nil on success
 */

/**
 * Name: Delete
 * Desc: removes a item within the list, index (int) that is passed
 * must be an available index within the list from [1 - N] where N
 * is the total number of items within the Todos list
 *
 * return: error on failure / nil on success
 */
func (t *Todos) Delete(index int) error {

	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index: out of range of Todo list")
	}

	// Appending the Todos list up to [index - 1] and appending it to the list starting from [index] to the end.
	// [1 2 3] + [5 6 7]
	*t = append(list[:index-1], list[index:]...)

	return nil
}

/**
 * Name: Load
 * Desc: loads the current state of the .todos.json file and unmarshalizes
 * the json into a Todos list object.
 *
 * return: error on failure / nil on success
 */
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

/**
 * Name: Store
 * Desc: Stores the current state of the static public Todos list and
 * Marshalizes into a .todos.json file within the local directory.
 *
 * return: error on failure / nil on success
 */
func (t *Todos) Store(filename string) error {

	// Taking list of Todos t and converting it to an JSON stream
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

/*
*
  - Name: Print
  - Desc: Prints the Todos list onto the CLI using the Simpletable library
  - the following is a preview of the view within the CLI.
    ╔═════╤═════════════════════╤═══════╤═════════════════════╤═════════════════════╗
    ║ #id │        Task         │ Done? │     Created-At      │    Completed-At     ║
    ╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
    ║ 1   │ Perform Tests       │ false │ 20 Nov 23 15:13 PST │ 01 Jan 01 00:00 UTC ║
    ║ 2   │ Complete assignment │ false │ 20 Nov 23 15:14 PST │ 01 Jan 01 00:00 UTC ║
    ║ 3   │ Go for a walk       │ false │ 20 Nov 23 15:14 PST │ 01 Jan 01 00:00 UTC ║
    ╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
    ║                         You have 3 uncompleted tasks                          ║
    ╚═════╧═════════════════════╧═══════╧═════════════════════╧═════════════════════╝
    *
  - return: nothing
*/
func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "Created-At"},
			{Align: simpletable.AlignCenter, Text: "Completed-At"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++

		task := blue(item.Task)
		if item.Done {
			task = green(item.Task)
		}

		completeText := item.CompletedAt.Format(time.RFC822)
		if item.CompletedAt.Compare(time.Time{}) == 0 {
			completeText = ""
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: completeText},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	uncompletedCount := t.CountPending()
	if uncompletedCount == 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: green("No remaining tasks! ✅")},
		}}
	} else {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d uncompleted tasks", t.CountPending()))},
		}}
	}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

/**
 * Name: CountPending
 * Desc: returns the number of pending tasks that still need to
 * be complete
 *
 * return: number of pending tasks
 */
func (t *Todos) CountPending() int {

	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
