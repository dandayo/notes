package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

const passwordCheckWord = "password-created-and-checked-user-can-login-in-it"

func CheckFolder() bool { //check is it first usage or not
	files, err := os.ReadDir("notes/")
	if err != nil {
		return false //if we have an error that means we don't have the folder and it's first usage
	}

	return len(files) > 0
}

func hiddenInput() string { // hidden input for password that you don't see ater tap a key
	for {
		fmt.Println("\033[33mEnter your password:\033[0m")
		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		check(err)

		fmt.Println()

		if len(bytePassword) == 0 {
			fmt.Println("Password cannot be empty")
			continue
		}

		return string(bytePassword)
	}
}

func encryptPassword(string) []byte {
	var encrypted []byte

	return encrypted
}

func CreatePassword() { //check do we have any notes or we need to create a password for notes
	if !CheckFolder() {
		err := os.MkdirAll("secret", 0755) //create a folder to save the password
		check(err)

		fmt.Println("Create a password for your notes, here is a hidden input, be careful")
		pass := hiddenInput()
		encPassword := encryptPassword(pass)

		err = os.WriteFile("secret/check.dat", encPassword, 0644)
		check(err)
	} else {
		pass := hiddenInput()

		ok := CheckPassword(pass)
		if ok {
			fmt.Println("\033[32mPassword is correct! Welcome to your notes\033[0m")
		} else {
			fmt.Println("\033[31mWrong password! Try again\033[0m")
		}
	}
}

func CheckPassword(string) bool { //check if the password is correct or not

	return false
}
