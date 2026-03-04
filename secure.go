package main

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/blowfish" //blowfish import to encrypt a note in file
)

// key to encrypt / decrypt
var key = []byte("password123")

// here we encrypt file
func encryptNote(note string) string {

	block, err := blowfish.NewCipher(key)
	check(err)

	toEncrypt := []byte(note)

	iv := make([]byte, block.BlockSize()) // initialization vector to make a
	rand.Read(iv)                         //random IV

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(toEncrypt, toEncrypt)

	encNote := append(iv, toEncrypt...)
	return base64.StdEncoding.EncodeToString(encNote)
}

// here we decrypt file
func decryptNote(enc string) string {
	raw, _ := base64.StdEncoding.DecodeString(enc) // get encrypted note

	if len(raw) < 8 { // Blowfish block size
		return "[invalid note]"
	}

	block, _ := blowfish.NewCipher(key)

	iv := raw[:block.BlockSize()]
	data := raw[block.BlockSize():] // read the note

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(data, data)

	return string(data)
}
