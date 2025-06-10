package utils

import (
	"fmt"
)

var forgotUi = `
REGISTER PAGE
------------------------`

func forgotPassword(){
	fmt.Print(forgotUi)
	fmt.Print("\nEnter new Password : ")
	var newPassword string
	fmt.Scanln(&newPassword)	

	userRegist[0].password = md5Encode(newPassword)

	defer handleLogin()
}