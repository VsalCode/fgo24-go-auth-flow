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
	fmt.Println(menuInteractive)
	var choice string
	fmt.Scanln(&choice)
	clear()
	switch choice {
	case "1":
		handleRegister()
	case "2":
		handleLogin()
	case "0":
		os.Exit(0)
	default:
		fmt.Println("Invalid choice!")
		Menu()
	}
}
