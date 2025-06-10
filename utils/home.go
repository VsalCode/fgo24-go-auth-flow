package utils

import (
	"fmt"
	"os"
)

var homeUi = `
=====================
|   Login Success!  |
=====================
|      WELCOME      |
=====================
1. Back To Menu
0. Exit This Program
=====================
`

func Home() {
	fmt.Println(homeUi)
	if currentUser != nil {
		fmt.Printf("Welcome, %s! (Email: %s)\n", currentUser.fullname, currentUser.email)
	} else {
		fmt.Println("No user logged in!")
	}
	fmt.Println("Registered users:", userRegist)

	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		clear()
		Menu()
	case "0":
		os.Exit(0)
	default:
		fmt.Println("Invalid choice!")
		clear()
		Home()
	}
}
