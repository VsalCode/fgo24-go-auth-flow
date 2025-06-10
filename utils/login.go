package utils

import "fmt"

func showFullName(p FullName) string {
	return p.getFullName()
}

type DataLogin struct {
	fullname string
	email    string
	password string
}

var getCurrentUser []DataRegister
var currentUser *DataLogin

var loginHome = `
LOGIN PAGE 
------------------------`

func handleLogin() DataLogin {
	defer Home()

	println(loginHome)
	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter your password : ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	getCurrentUser = []DataRegister{}

	for x := range userRegist {
		if inputEmail == userRegist[x].email && md5Encode(inputPassword) == userRegist[x].password {
			getCurrentUser = append(getCurrentUser, userRegist[x])
			break
		}
	}

	if len(getCurrentUser) == 0 {
		fmt.Printf("\nEmail or Password invalid!\n")
		fmt.Printf(`
        1. Try Again
        2. Forgot Password
        0. Back To Menu
        `)
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			return handleLogin()
		case "2":
			forgotPassword()
			return DataLogin{}
		case "0":
			Menu()
			return DataLogin{}
		default:
			fmt.Println("Invalid choice!")
			return handleLogin()
		}
	}

	fullname := showFullName(getCurrentUser[0])
	currentUser = &DataLogin{
		fullname: fullname,
		email:    getCurrentUser[0].email,
		password: getCurrentUser[0].password,
	}

	fmt.Println("Login Successful!")
	clear()

	return *currentUser
}
