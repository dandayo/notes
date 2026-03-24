package main

import (
	"crypto/sha256" // hash function to encrypt the password
	"fmt"
	"os"

	"golang.org/x/term"
)

func CheckFolder() bool { //check is it first usage or not
	files, err := os.ReadDir("secret/")
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

func CreatePassword() string { //check do we have any notes or we need to create a password for notes
	if !CheckFolder() {
		err := os.MkdirAll("secret", 0755) //create a folder to save the password
		check(err)
		fmt.Println("\033[32mCreate a password for your notes, here is a hidden input, be careful\033[0m")
		pass := hiddenInput()
		hash := sha256.Sum256([]byte(pass))
		err = os.WriteFile("secret/check.dat", hash[:], 0600)
		check(err)
		return pass
	} else {
		pass := hiddenInput()
		ok := CheckPassword(pass)
		if ok {
			fmt.Println("\033[32mPassword is correct! Welcome to your notes\033[0m\n")
			return pass
		} else {
			fmt.Println("\033[31mWrong password!\033[0m")
			return CreatePassword()
		}
	}
}

func CheckPassword(password string) bool { //check if the password is correct or not
	hash := sha256.Sum256([]byte(password))
	checkHash, err := os.ReadFile("secret/check.dat")
	check(err)

	if string(hash[:]) == string(checkHash) {
		return true
	}

	return false
}

func updatePassword() { //update the password if the user wants to change it
	err := os.Remove("secret/check.dat")
	os.RemoveAll("notes/")
	check(err)
	fmt.Println("\n\033[97;41mPassword deleted :(\033[0m")
	fmt.Printf("\033[42mGoodbye!\033[0m\n")
}
