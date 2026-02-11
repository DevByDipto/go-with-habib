package database

import (
	"fmt"
)

// Product struct definition
type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"` // স্ক্রিনশটে ইন্টিজার হিসেবে দেখা যাচ্ছে
	ImgUrl      string `json:"imgUrl"`
}

// Global slice to store products
var productList []Product

// Store adds a new product to the list
func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

// List returns all products
func List() []Product {
	return productList
}

// Get returns a pointer to a product by its ID
func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product 
		}
	}
	return nil
}

// Update modifies an existing product in the list
func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

// Delete removes a product from the list by ID
func Delete(productID int) {
	var tempList []Product
	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

// init function to populate initial data
func init() {
	prd4 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is boring. I feel bored eating banana.",
		Price:       50,
		ImgUrl:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biN",
	}

	prd5 := Product{
		ID:          4, // ID ইউনিক হওয়া উচিত, স্ক্রিনশটে ৩ ছিল
		Title:       "Angur Fol",
		Description: "Banana is boring. I feel bored eating banana.",
		Price:       140,
		ImgUrl:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biN",
	}

	Store(prd4)
	Store(prd5)
	
	fmt.Println("Database initialized with", len(productList), "products.")
}