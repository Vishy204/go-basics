package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
	"example.com/go-project/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	text := getTodoData()
	userNote, err := note.New(title, content)
	toDo, errr := todo.New(text)

	if errr != nil {
		fmt.Println(errr)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	toDo.Display()
	userNote.Display()
	err = saveData(toDo)
	if err != nil {
		return
	}

	err = saveData(userNote)
	if err != nil {
		return
	}

}
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Failed to save note successfully")
		return err
	}

	fmt.Println("Saving the note successfully")
	return nil
}
func getTodoData() string {
	text := getUserInput("Give todo content")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note content: ")

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
