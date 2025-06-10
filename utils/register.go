package utils

import (
	"crypto/md5"
	"fmt"
	"os"
)


type DataRegister struct {
	fullname        string
	email           string
	password        string
}

var registerHome = `
REGISTER PAGE
------------------------
`

func handleRegister() []DataRegister {
	fmt.Print(registerHome)
	fmt.Print("\nEnter Fullname : ")
	var inputFullname string
	fmt.Scanln(&inputFullname)

	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter your password : ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	fmt.Print("\nEnter your password : ")
	var confirmPassword string
	fmt.Scanln(&confirmPassword)

	if confirmPassword != inputPassword {
		os.Exit(1)
	}
	password := md5Encode(inputPassword)

	var userRegist []DataRegister
	userRegist = append(userRegist, DataRegister{fullname: inputFullname, email: inputEmail, password: password })

	return userRegist
}

func md5Encode(input string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(input))

	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}