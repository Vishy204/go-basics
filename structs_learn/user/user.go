package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// adding methods to struct
func (u User) SploutputUserDetails() {
	fmt.Println(u.birthdate, u.createdAt, u.firstName, u.lastName)
}

func (u *User) DelNames() {
	u.firstName = "Sambar"
	u.lastName = "chutney"
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("Empty firstname and lastname")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil

}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "adam",
			lastName:  "maker",
			createdAt: time.Now(),
			birthdate: "----",
		},
	}
}
