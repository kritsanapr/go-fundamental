package variables

import (
	"fmt"
	"strconv"
)

// Declare variables at package level
var fullname string
var email string
var c, python, java bool = true, false, true

// var age int = 25
var typeb bool = true
var typef float32 = 3.14 // Floating point number
var types string = "Hi!" // String

// Constant
const Pi = 3.14
const vat = 7

func Learn() {
	// Assign variables
	fullname = "Kritsanapoom"
	email = "kritsana@gmail.com"

	// Print line
	fmt.Println("Learn variables")
	fmt.Println("Fullname: ", fullname)
	fmt.Println("Email: ", email)
	fmt.Println("C: ", c)
	fmt.Println("Python: ", python)
	fmt.Println("Java: ", java)
	// Print format
	fmt.Printf("Fullname: %s, Email: %s", fullname, email)
	// Print type
	fmt.Printf("Type of fullname: %T, Type of email: %T", fullname, email)

	// Short variable declarations
	firstname := "Kritsanapoom"
	lastname := "L"

	// Convert string to int
	convertage := strconv.Itoa(vat)
	fmt.Println("convert string to int", convertage)
	fmt.Println("Firstname: ", firstname)
	fmt.Println("Lastname: ", lastname)

	// Print type
	fmt.Printf("Type of firstname: %T, Type of lastname: %T", firstname, lastname)

}
