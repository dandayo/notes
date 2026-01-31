package main

import (
    "fmt"
    "os"
    "bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CreateFile() {
    file, err := os.Create("new.txt")
    check(err)
    defer file.Close()
    fmt.Println("File created successfully")
}

func DeleteFile() {
    err := os.Remove("new.txt")
    check(err)
    fmt.Println("File deleted")
}

func WriteInFile() {
    file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    check(err)
    defer file.Close()

    data := "new new info \n"
    _, err = file.WriteString(data)
    check(err)

}

func ReadFile() {
	notes, err := os.Open("new.txt")  //open the file
	check(err)
    defer notes.Close()

    scanner := bufio.NewScanner(notes)  //scan the contents of a file and print line by line
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }
    check(err)
}


func RemoveElement(array []string, n int)[]string{
		if (len (array)< 1|| n-1 > len (array)){
			return array
		}
		return append (array[:n], array[n+1:]...)
}

func main() {
	read := ReadFile()
	return read[0]
}
