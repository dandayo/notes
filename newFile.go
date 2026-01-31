package main

import (
    "fmt"
    "os"
)

func CreateFile() {
	newFile := input()
    file, err := os.Create(newFile)
    if err != nil {
      fmt.Println("Mistake, try again")
    }
    defer file.Close()
    fmt.Println("File created successfully")
}

func DeleteFile() {
    err := os.Remove("new.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("File deleted")
}

func WriteInFile() {
    file, err := os.OpenFile("new.txt", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
    	fmt.Println("Mistake, try again")
    }
    defer file.Close()

    data := "i""Input()"
    _, err = file.WriteString(data)
    if err != nil {
    fmt.Println("Mistake, try again")
    }

}

func main() {

}
