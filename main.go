package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interfacePrac/note"
	"interfacePrac/todo"
)

func main() {
	title, content := getNoteData()
	todoTitle := getUserInput("Todo title:")

	userNote, err := note.New(title, content)
	userTodo, todoErr := todo.New(todoTitle)

	if err != nil {
		fmt.Println(err)
		return
	}

	if todoErr != nil {
		fmt.Println(todoErr)
		return
	}

	userTodo.Display()
	userNote.Display()
	err = userNote.Save()
	todoErr = userTodo.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Saving the note succeeded!")

	if todoErr != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Saving the todo succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}



func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}