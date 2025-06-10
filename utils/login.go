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

func handleLogin() DataLogin {
	data := userRegist[0]

	println(loginHome)
	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter your password : ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	if inputEmail != data.email {
		fmt.Printf("\nYour Email is Wrong!\n")
		fmt.Printf(`
		Your Password is wrong!
		1. Try Again
		2. Exit
		`)
		var inputPasswordAgain string
		fmt.Scanln(&inputPasswordAgain)
		if (inputPasswordAgain == "1"){
			handleLogin()
		}
		if (inputPasswordAgain == "2"){
			os.Exit(0)
		}
	}

	if md5Encode(inputPassword) != data.password {
		fmt.Printf(`
		Your Password is wrong!
		1. Try Again
		2. Forgot Password
		3. Exit
		`)
		var inputPasswordAgain string
		fmt.Scanln(&inputPasswordAgain)
		if (inputPasswordAgain == "1"){
			handleLogin()
		}
		if (inputPasswordAgain == "2"){
			forgotPassword()
		}
		if (inputPasswordAgain == "3"){
			os.Exit(0)
		}
	}

	defer fmt.Printf("Login Succesfully!")

	fullname := data.getFullName()

	user := DataLogin{
		fullname: fullname,
		email:    data.email,
		password: data.password,
	}

	return user
}
