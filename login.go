package main

import (
	"fmt"
	"os"
)

const passwordCheckWord = "password-created-and-checked-user-can-login-in-it"

func CheckFolder() bool { //check is it first usage or not
	files, _ := os.ReadDir("notes/")
	if len(files) >= 1 {
		return true
	}
	return false
}

func hiddenInput() { //input without showing the input in the terminal

}

func CreatePassword() string { //check do we have any notes or we need to create a password for notes
	if !CheckFolder() {
		err := os.Mkdir("secret", 0755) //create a folder to save the password
		check(err)

		file, err := os.Create("secret/" + "check" + ".txt")
		check(err)

		fmt.Println("Create a password for your notes, here is a hidden input, be careful")
		password := input()

		return password
	} else {
		password := input()
		CheckPassword(password)
		return password
	}
}

func CheckPassword(string) bool { //check if the password is correct or not

	return false
}
