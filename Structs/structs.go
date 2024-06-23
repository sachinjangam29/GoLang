package main

import (
	"Structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your DOB (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	adminUser := user.NewAdmin("test@gmail.com", "test123")

	adminUser.User.OututUser()
	adminUser.User.ClearUser()
	adminUser.User.OututUser()

	/* We can also use below code replacing the above code but only when we use only User instead of User User inside the Admin struct
	adminUser.User.OututUser()
	adminUser.User.ClearUser()
	adminUser.User.OututUser()
	*/
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OututUser()
	appUser.ClearUser()
	appUser.OututUser()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
