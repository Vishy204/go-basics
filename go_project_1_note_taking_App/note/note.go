package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (userNote Note) Display() {
	fmt.Printf("Title is %v and the content is: \n\n%v\n\n", userNote.Title, userNote.Content)
	fmt.Printf(" This note was created at: %v\n", userNote.CreatedAt)

}

func (userNote Note) Save() error {
	fileName := strings.ReplaceAll(userNote.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	data, err := json.Marshal(userNote)
	if err != nil {
		return err
	}
	fileName = fileName + ".json"
	return os.WriteFile(fileName, data, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input empty string")

	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}
