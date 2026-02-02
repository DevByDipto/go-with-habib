package main

import "fmt"

func a() {
	i := 0

	fmt.Println("first", i) // আউটপুট: 0

	defer fmt.Println("second", i) // i এর বর্তমান মান 0 এখানে জমা থাকবে

	i = i + 1 // বা i++

	defer fmt.Println("fourth", i) // i এর বর্তমান মান 1 এখানে জমা থাকবে

	fmt.Println("third", i) // আউটপুট: 1

	return
}

func main() {
	a()
}
// task : simulate and output kii kii hobe ?
// --------------------------------------------------------------------------