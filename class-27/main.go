package main

import "fmt"


type User struct {
	Name string // member variable /field / property
	Age  int    
}

func main() {

	
	var user1 User


	// Step 2: user1 এ value বসানো (instantiate করা)
	user1 = User{
		Name: "Habib",
		Age:  30,
	}

	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)

	user2 := User{ // instance or object
		Name: "Roki",
		Age:  16,
	}

	
	fmt.Println("Name:", user2.Name)
	fmt.Println("Age:", user2.Age)
}

// task- samulate korte hobe