package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
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

func CreatePassword() { //check do we have any notes or we need to create a password for notes
	if !CheckFolder() {
		err := os.Mkdir("secret", 0755) //create a folder to save the password
		check(err)

		f, err := os.Create("secret/" + "check" + ".txt")
		check(err)

		fmt.Println("Create a password for your notes, here is a hidden input, be careful")
		password := input()

		block, err := aes.NewCipher(key)
		check(err)

		aesgcm, err := cipher.NewGCM(block)
		check(err)

		nonce := make([]byte, aesgcm.NonceSize())

		if _, err := io.ReadFull(rand.Reader, nonce){
			check(err)
		}

	} else {
		CheckPassword()
	}
}

func CheckPassword() bool { //check if the password is correct or not

	return false
}
