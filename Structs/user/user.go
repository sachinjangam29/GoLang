package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// * is used with the user type with a function to refer to the same data, if we do not use
// * then this function will refer to a new copy of the struct data which will give new data and it will not
// refer to the same data.
func (u *User) OututUser() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUser() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	// Validation
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First Name, Last Name, Birth Date cannot be blank")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}