// ---------------------------------------------------------------------prase-2

// package mathlib

// import "fmt"

// func addFunc(num1 int, num2 int){
// 	sum := num1 + num2
// 	fmt.Println("sum from addFunc", sum)
// }

// func another(){
// 	addFunc(5,6)
// }

// ---------------------------------------------------------------------prase-3

package mathlib

import "fmt"

func AddFunc(num1 int, num2 int){
	sum := num1 + num2
	fmt.Println("sum from AddFunc", sum)
}

func addFunc(num1 int, num2 int){
	sum := num1 + num2
	fmt.Println("sum from addFunc", sum)
}

func another(){
	addFunc(5,6)
}