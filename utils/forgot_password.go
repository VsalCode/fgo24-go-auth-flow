package utils

import (
	"fmt"
)

var forgotUi = `
FORGOT PASSWOD
------------------------`

func forgotPassword(){
	fmt.Print(forgotUi)
	fmt.Print("\nEnter new Password : ")
	var newPassword string
	fmt.Scanln(&newPassword)	

	currentUser.password = md5Encode(newPassword)

	clear()
	defer handleLogin()
}