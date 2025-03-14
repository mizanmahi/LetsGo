package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

func New(title string) (*Todo, error) {

	if title == "" {
		return &Todo{}, errors.New("title cannot be empty")
	}

	return &Todo{Title: title, Done: false}, nil
}

func (t Todo) Display() {
	fmt.Printf("Title: %s\nDone: %t\n", t.Title, t.Done)
}

func (t Todo) Save() error {
	todoJson, err := json.Marshal(t)

	if err != nil {
		fmt.Println(err)
		return errors.New("failed to save the todo")
	}

	os.WriteFile("todo.json", todoJson, 0644)
	return nil
}