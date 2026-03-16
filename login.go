package main

import (
	"fmt"
	"os"

	"github.com/tcnksm/go-input"
)

const passwordCheckWord = "password-created-and-checked-user-can-login-in-it"

func CheckFolder() bool { //check is it first usage or not
	files, _ := os.ReadDir("notes/")
	if len(files) >= 1 {
		return true
	}
	return false
}

func hiddenInput() string { //input without showing the input in the terminal
	fmt.Println("Enter your password: ")

	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	password, err := ui.Ask("Enter password:", &input.Options{
		Required: true,
		Mask:     '*', // hide  input
		Loop:     true,
	})

	check(err)

	return password
}

func encryptPassword(string) string {
	var encrypted string

	return encrypted
}

func decryptPassword(string) string {
	var decrypted string

	return decrypted
}

func CreatePassword() { //check do we have any notes or we need to create a password for notes
	if !CheckFolder() {
		err := os.Mkdir("secret", 0755) //create a folder to save the password
		check(err)

		file, err := os.Create("secret/" + "check" + ".dat")
		check(err)

		fmt.Println("Create a password for your notes, here is a hidden input, be careful")
		pass := hiddenInput()

		os.WriteFile("secret/check.dat", encrypted, 0644)
	} else {
		CheckPassword(pass)
	}
}

func CheckPassword(string) bool { //check if the password is correct or not

	return false
}
