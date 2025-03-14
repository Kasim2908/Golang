package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	// comma ok || err ok
	fmt.Println("Enter the rating for our pizza: ")
	// cmma ok || err ok
	input, _ := reader.ReadString('\n') // try or catch part
	// comma ok || err ok
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
