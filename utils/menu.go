package utils

import (
	"fmt"
	"os"
)

var menuInteractive = `
=====================
|      WELCOME      |
=====================
1. Register
2. Login
0. Exit
=====================		
`

func Menu() {
	fmt.Println()
	fmt.Printf(menuInteractive)
	var inputPasswordAgain string
	fmt.Scanln(&inputPasswordAgain)
	clear()
	if inputPasswordAgain == "1" {
		handleRegister()
	}
	if inputPasswordAgain == "2" {
		handleLogin()
	}
	if inputPasswordAgain == "3" {
		os.Exit(0)
	}
}

