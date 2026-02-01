package main

import (
    "fmt"
    "os"

)

func main() {
	EmptyTerminal()
	fmt.Println("\033[32m\n===Welcome to the notes tool!===\033[0m")
	fileName := os.Args[1]
	OpenFile(fileName)
	Menu(fileName)
}
