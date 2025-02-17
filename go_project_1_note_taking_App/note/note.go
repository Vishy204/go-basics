package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (userNote Note) Display() {
	fmt.Printf("Title is %v and the content is: \n\n%v\n\n", userNote.title, userNote.content)

	fmt.Printf(" This note was created at: %v\n", userNote.createdAt)

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input empty string")

	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil

}
