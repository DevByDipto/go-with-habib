package main

import "fmt"

const a = 10 // constant

var p = 100 // run time

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}

	return show
}

func call() {
	incr1 := outer() // show
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

// task - simulate how it work