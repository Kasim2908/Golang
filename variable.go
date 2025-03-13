package main

import "fmt"

const LgginToken string = "hghghhg" //public

func main() {
	var username string = "kasim"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallfloat float64 = 255.4554433
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var styled
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LgginToken)
	fmt.Printf("Variable is of type: %T \n", LgginToken)
}
