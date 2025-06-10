package utils

import (
	"fmt"
	"os"
)
type DataLogin struct {
	fullname string
	email    string
	password string
}

var loginHome = `
LOGIN PAGE
------------------------`

func HandleLogin() []DataLogin {
	dataRegist := handleRegister()
	data := dataRegist[0]

	println(loginHome)
	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter your password : ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	if inputEmail != data.email {
		fmt.Printf("\nYour Email is Wrong!\n")
		os.Exit(1)
	}

	if md5Encode(inputPassword) != data.password {
		fmt.Printf("\nYour Email is true!\n")
		os.Exit(1)
	}
	
	defer fmt.Printf("Login Succesfully!")
	
	fullname := data.getFullName()

	var userLogin []DataLogin
	userLogin = append(userLogin, DataLogin{fullname: fullname, email: data.email, password: data.password })

	return userLogin
}
