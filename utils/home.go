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
`

func Home() {
	fmt.Println()
	fmt.Println(userRegist)
	fmt.Printf(homeUi)
	var inputPasswordAgain string
	fmt.Scanln(&inputPasswordAgain)
	if inputPasswordAgain == "1" {
		clear()
		currentUser = nil
		Menu()
	}
	if inputPasswordAgain == "0" {
		os.Exit(0)
	}
}
