package utils

import (
	"fmt"
	"os"
)

type login struct {
	fullname string
	email    string
	password string
}

var loginHome = `
LOGIN PAGE
------------------------
`

func HandleLogin(){
	dataRegist := handleRegister()
	data := dataRegist[0]

	println(loginHome)
	fmt.Print("\nEnter your Email : ")
	var inputEmail string
	fmt.Scanln(&inputEmail)

	fmt.Print("\nEnter your password : ")
	var inputPassword string
	fmt.Scanln(&inputPassword)

	if (inputEmail != data.email ){
		os.Exit(1)
	}
	
	if (md5Encode(inputPassword) != data.password ){
		os.Exit(1)
	}
}

