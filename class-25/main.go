
// real simulation aikhan theke aromvo

package main

import "fmt"

const a = 10 // constant

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y 
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}

func main() {
	call()

	fmt.Println(a)
}

func init() {
}





// anonymas add function tar reference stack e rakha hoi tokhon  jokhon call() function call hoi 

// note: const value gulo normal value hole compile time e jekhane a er man dorkar se khane direct a er value bose jai jmn  add(p,a) aita compile time e add(p,10) hoi jabe . abar const variable er value complex hole like a= string or a=34.534534 amn kichu tokhon const value er jonno readonly data segment toiri hoye tate value rakha hoi.