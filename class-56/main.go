package main

import (
	"ecommerce/util"
	"fmt"
	// 	"ecommerce/config"
	// 	"fmt"
)




func main() {
	// cmd.Serve()
	// Example usage:
	p := util.Payload{
		Sub:         "1234567890",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		IsShopOwner: true,
	}

	token, err := util.CreateJwt("your-256-bit-secret", p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(token)

}






