package main

import "fmt"
import "time"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	for {
		_, err := fmt.Scan(&number)
		if err == nil {
			break
		}
		fmt.Print("Invalid input. Please enter a number: ")
		
	}
	
	for range 10 {
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("Hello, World! %d\n", number)
	}
}
