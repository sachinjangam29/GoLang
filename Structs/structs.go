package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// * is used with the user type with a function to refer to the same data, if we do not use
// * then this function will refer to a new copy of the struct data which will give new data and it will not
// refer to the same data.
func (u *user) oututUser() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *user) clearUser() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your DOB (MM/DD/YYYY): ")

	var appUser *user

	appUser = newUser(userFirstName, userLastName, userBirthDate)

	appUser.oututUser()
	appUser.clearUser()
	appUser.oututUser()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
