package utils

import (
	"crypto/md5"
	"fmt"
	"os"
)

type DataRegister struct {
	firstName string
	lastName  string
	email    string
	password string
}

func (x DataRegister) getFullName()string {
	result := fmt.Sprintf("%s %s", x.firstName, x.lastName)
	return result
}

var registerHome = `
REGISTER PAGE
------------------------`

var userRegist []DataRegister

func handleRegister() []DataRegister {
	fmt.Print(registerHome)
	fmt.Print("\nEnter your FirstName : ")
	var inputFirstName string
	fmt.Scanln(&inputFirstName)

	fmt.Print("\nEnter your LastName : ")
	var inputLastName string
	fmt.Scanln(&inputLastName)

	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter your password : ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	fmt.Print("\nConfirm your password : ")
	var confirmPassword string
	fmt.Scanln(&confirmPassword)

	if confirmPassword != inputPassword {
		fmt.Println("confirm passowrd must match!")
		os.Exit(1)
	}
	password := md5Encode(inputPassword)

	
	userRegist = append(userRegist, DataRegister{firstName: inputFirstName, lastName: inputLastName, email: inputEmail, password: password})

	if userRegist != nil {
		clear()
		defer handleLogin()
	}

	return userRegist
}

func md5Encode(input string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(input))

	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}
