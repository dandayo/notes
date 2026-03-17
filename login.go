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

	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	password, err := ui.Ask("Enter your password:", &input.Options{
		Required: true,
		Mask:     true, // hide  input
		Loop:     true,
	})

	check(err)

	return password
}

func encryptPassword(string) []byte {
	var encrypted []byte

	return encrypted
}

func decryptPassword([]byte) string {
	var decrypted string

	return decrypted
}

func CreatePassword() { //check do we have any notes or we need to create a password for notes
	if !CheckFolder() {
		err := os.Mkdir("secret", 0755) //create a folder to save the password
		check(err)

		fmt.Println("Create a password for your notes, here is a hidden input, be careful")
		encryptPassword := encryptPassword(passwordCheckWord)

		os.WriteFile("secret/check.dat", encryptPassword, 0644)
	} else {
		pass := hiddenInput()
		CheckPassword(pass)
		if CheckPassword(pass) {
			fmt.Println("\033[32mPassword is correct! Welcome to your notes\033[0m")
		} else {
			fmt.Println("\033[31mWrong password! Try again\033[0m")
		}
	}
}

func CheckPassword(string) bool { //check if the password is correct or not

	return false
}
