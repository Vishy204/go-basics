package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("The Todo text is %v\n\n", todo.Text)

}

func (todo Todo) Save() error {
	fileName := "todo.json"
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input empty string")

	}
	return Todo{
		Text: text,
	}, nil

}
