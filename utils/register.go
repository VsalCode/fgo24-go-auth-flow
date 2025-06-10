package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

type FullName interface {
	getFullName() string
}

type DataRegister struct {
	firstName string
	lastName  string
	email     string
	password  string
}

func (x DataRegister) getFullName() string {
	result := fmt.Sprintf("%s %s", x.firstName, x.lastName)
	return result
}

var registerHome = `
REGISTER PAGE
------------------------`

var userRegist []DataRegister

func handleRegister() []DataRegister {
	defer handleLogin()

	fmt.Print(registerHome)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nEnter your FirstName : ")
	scanner.Scan()
	inputFirstName := scanner.Text()

	fmt.Print("\nEnter your LastName : ")
	scanner.Scan()
	inputLastName := scanner.Text()

	fmt.Print("\nEnter your Email : ")
	scanner.Scan()
	inputEmail := scanner.Text()

	fmt.Print("\nEnter your password : ")
	scanner.Scan()
	inputPassword := scanner.Text()

	fmt.Print("\nConfirm your password : ")
	scanner.Scan()
	confirmPassword := scanner.Text()

	if confirmPassword != inputPassword {
		fmt.Println("Confirm password must match!")
		os.Exit(1)
	}
	password := md5Encode(inputPassword)

	userRegist = append(userRegist, DataRegister{
		firstName: inputFirstName,
		lastName:  inputLastName,
		email:     inputEmail,
		password:  password,
	})

	fmt.Println("Registration successful!")
	clear()

	return userRegist
}

func md5Encode(input string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(input))
	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}
