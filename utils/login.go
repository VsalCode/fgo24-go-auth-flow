package utils

import "fmt"

type DataLogin struct {
	fullname string
	email    string
	password string
}

var getCurrentUser []DataRegister
var currentUser *DataRegister

var loginHome = `
LOGIN PAGE 
------------------------`

func handleLogin() DataLogin {

	println(loginHome)
	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter your password : ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	for x := range userRegist {

		if inputEmail != userRegist[x].email {
			fmt.Printf("\nYour Email is Wrong!\n")
			fmt.Printf(`
		Your Password is wrong!
		1. Try Again
		0. Back To Home
		`)
			var inputPasswordAgain string
			fmt.Scanln(&inputPasswordAgain)
			if inputPasswordAgain == "1" {
				handleLogin()
			}
			if inputPasswordAgain == "0" {
				Menu()
			}
		}

		if md5Encode(inputPassword) != userRegist[x].password {
			fmt.Printf(`
		Your Password is wrong!
		1. Try Again
		2. Forgot Password
		0. Back To Home
		`)
			var inputPasswordAgain string
			fmt.Scanln(&inputPasswordAgain)
			if inputPasswordAgain == "1" {
				handleLogin()
			}
			if inputPasswordAgain == "2" {
				forgotPassword()
			}
			if inputPasswordAgain == "0" {
				Menu()
			}
		}

		if inputEmail == userRegist[x].email && md5Encode(inputPassword) == userRegist[x].password {
			getCurrentUser = append(getCurrentUser, userRegist[x]) 
		}
	}

	fullname := getCurrentUser[0].getFullName()

	currentUser := DataLogin{
		fullname: fullname, 
		email: getCurrentUser[0].email, 
		password: getCurrentUser[0].password,
	}

	fmt.Println("Login Successfuly!")
	clear()
	defer Home()

	return currentUser
}
