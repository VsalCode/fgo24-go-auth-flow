package utils

import (
	"fmt"
)

var forgotUi = `
FORGOT PASSWORD
------------------------`

func forgotPassword() {
	fmt.Print(forgotUi)
	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter new Password : ")
	var newPassword string
	fmt.Scanln(&newPassword)

	for i := range userRegist {
		if userRegist[i].email == inputEmail {
			userRegist[i].password = md5Encode(newPassword)
			fmt.Println("Password updated successfully!")
			clear()
			handleLogin()
			return
		}
	}

	fmt.Println("Email not found!")
	fmt.Printf(`
    1. Try Again
    0. Back To Menu
    `)
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		forgotPassword()
	case "0":
		Menu()
	default:
		fmt.Println("Invalid choice!")
		forgotPassword()
	}
}
