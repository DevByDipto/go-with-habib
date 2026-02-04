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
	/*
	outpot:
	first 0
	third 1
	fourth 1
	second 0
	*/  
}
// task : simulate and output kii kii hobe ?
// --------------------------------------------------------------------------
package main

import "fmt"

func sum(a int, b int) (s int) { // named retuns values
	s = a + b
	return 
}

func main() {
	res := sum(3, 4)

	fmt.Println(res)
}
// --------------------------------------------------------------------------------
/*
named return values

1. all codes execute
2. defer function store kora hobe defer box e
3. return -> all defer functions execute korbe
4. return korbe named variables gular values

just return types

1. all codes execute
2. defer function store kora hobe magic box e
3. return values are evaluated at this time (store the return value)
4. all defer functions execute korbe
*/

package main

import "fmt"

func calculate() (result int) {
	fmt.Println("first", result) // first 0

	show := func() {
		result = result + 10
		fmt.Println("defer", result) // defer 15
	}

	defer show()

	result = 5
	fmt.Println("second", result) // second 5

	return
}

func calc() int {
	result := 0
	fmt.Println("first", result) // first 0

	show := func() {
		result = result + 10
		fmt.Println("defer", result) // defer 10
	}

	defer show() 

	result = 5
	fmt.Println("second", result) // second 5

	return result
}

func main() {
	a := calculate()
	fmt.Println("main first", a) // 15

	b := calc()
	fmt.Println("main second", b) // 5
}
// task : how the code work

/*
calculate
calculateAnonymus1
calc
calcAnonymous1
main
*/
// --------------------------------------------------------------------------------
package main

import "fmt"

func calculate() (result int) {
	// First frame shows the initial print
	fmt.Println("first", result)

	// Define the show function as a closure
	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	// First deferred call
	defer show()

	result = 5

	// Define function p that takes an int
	p := func(a int) {
		fmt.Println("ami", a)
	}

	// Second deferred call: result is evaluated NOW (5)
	defer p(result)

	// Third deferred call: result is evaluated NOW (5)
	defer fmt.Println(result)

	fmt.Println("second", result)

	// Fourth deferred call: constant value
	defer fmt.Println(5)

	return
}

func main() {
	a := calculate()
	fmt.Println("main first", a)
}

// task - simulate with how defer box exectly works
// defer kon deta structure maintain kore ?
// ans: defer link list data structure maintain kore with stack behabior 
// defer actually kivabe kaj kore ?