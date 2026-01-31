package main

import (
    "fmt"
    "os"
)

func CreateFile() {
    file, err := os.Create("new.txt")
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
    file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
    	fmt.Println("Mistake, try again")
    }
    defer file.Close()

    data := "new new info \n"
    _, err = file.WriteString(data)
    if err != nil {
    	fmt.Println("Mistake, try again")
    }

}

func main() {
	WriteInFile()
}
