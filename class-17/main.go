package main

import "fmt"

var (
	a =5
	b=7
)

func addFunc(num1 int, num2 int){
	sum := num1 + num2
printFunc(sum)
}

func main(){
	addFunc(5,6)
}

func printFunc(sum int){
	fmt.Println("sum from addFunc", sum)
}

// ai code ta ram e kivabe exicute hocche