package main

import "fmt"

func main() {
	for  {
		fmt.Print("Enter a number: ")
		var input int
		count, err := fmt.Scan(&input)
		if err != nil || count != 1 {
			fmt.Println("Invalid input, please enter ONE number.")
		}else{
			fmt.Printf("You entered: %d", input)
			break
		}
	}
}
