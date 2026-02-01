package main
import ("bufio"
		"fmt"
		"os"
	//	"strconv"
	    "strings"
		)


func input() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()  


	return strings.TrimSpace(scanner.Text())

}


func show() {

	
}


func add() {


}


func delete() {


}

func menu() {

	msg := ""

	for msg != "4" {

		fmt.Println("\n===Welcome to the notes tool!===\n")
		fmt.Println("Select operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		msg = input()

		switch msg {

			case "1":
				show()

			case "2": 
				add()
			
			case "3":
				delete()
			
			case "4":
				fmt.Println("===Goodbye!===\n")
				os.Exit(0)
			
			default:
				fmt.Println("\n===You have entered an invalid input!===\n")
				continue
			
		}

		
	}


}		


func main() {

	menu()
}