package main

import "fmt"

// User struct definition
type User struct {
	Name string
	Age  int
}

// printUserDetails is a regular function that takes a User parameter
func printUserDetails(usr User) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

// printDetails is a receiver function (method) on User
func (usr User) printDetails() {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

// call is a method that takes an int parameter
func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	// Create first user instance
	var user1 User

	user1 = User{ // instantiate
		Name: "Habib",
		Age:  30,
	}

	// Call methods on user1
	// printUserDetails(user1)  // commented out in original
	user1.printDetails()
	user1.call(10)

	// Create second user instance
	user2 := User{ // Instance or object
		Name: "Roki",
		Age:  16,
	}

	printUserDetails(user2)
}

// task simulate the code