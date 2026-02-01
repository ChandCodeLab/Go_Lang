package main

import "fmt"

//constant
const LoginToken string = "makdhe" //public varibale  denoted by first uppercase

// jwtToken :=30000 not allowed outside
func main() {
	// var username string = "Manoj"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.354339393838
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var stringVariable string
	fmt.Println(stringVariable)
	fmt.Printf("Variable is of type: %T \n", stringVariable)

	//implicit type

	var website = "Learncodeline.in"
	fmt.Println(website)

	//
	numberOfUser := 30000.00
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
