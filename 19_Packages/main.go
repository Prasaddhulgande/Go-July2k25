package main

import (
	"GO-JULY2K25/auth"
	"GO-JULY2K25/user"
	"fmt"

	// "image/color"
	"github.com/fatih/color"
)

func main() {

	auth.LoginWithCredentiaks("Prasad@12345", "asd@123")

	session := auth.GetSession()

	fmt.Println("Session", session)

	usser := user.User{

		Email: "user@gmail.com",
		Name:  "Rober downy",
		Phone: 7507766435,
	}

	//fmt.Println(usser.Email, usser.Name, usser.Phone)

	// Use color package

	color.Yellow(usser.Email)
}
