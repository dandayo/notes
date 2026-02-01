package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func check(e error) {
    if e != nil {
    	fmt.Println("Mistake!")
    }
}

func CreateFile(name string) {
    file, err := os.Create(name + ".txt")
    check(err)
    defer file.Close()
    fmt.Println("File created successfully")
}

func AddNote(name string, note string) {
	name = name + ".txt"
    file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY ,0755)
    defer file.Close()

    data := "\n" + note
    _, err = file.WriteString(data)
    check(err)
}


func DeleteFile(name string) {
    err := os.Remove(name + ".txt")
    check(err)
    fmt.Println("File deleted")
}

func RemoveNote(fileName string, removeIndex int) {
	notes, err := os.Open(fileName)  //open the file
	check(err)
    defer notes.Close()

    var lines []string
    scanner := bufio.NewScanner(notes)
    var index int = 0

    for scanner.Scan() {
         text := scanner.Text()
         if index == removeIndex - 1{
         	index++
          	continue
         }
         lines = append(lines, text)
         index++
    }

    check(scanner.Err())

    update := strings.Join(lines, "\n")
    os.WriteFile(fileName, []byte(update), 0755)
}

func main() {
	AddNote("sello", "hello world")
}
