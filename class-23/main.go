// package main

// import "fmt"

// func processOperation(a int, b int, op function(p int, q int)){
// 	op(a,b)
// }

// func add(x int, y int){
// 	z := x+y
// 	fmt.Println(z)
// }

// func main(){
// 	processOperation(2,5,add)
// }
// // task : ram e code ta kivabe exicute hobe ?
//----------------------------------------------------------------

package main

import "fmt"

func call() func(x int, y int){
	return add
}

func add(x int, y int){
	z := x+y
	fmt.Println(z)
}

func main(){
	sum := call()

	sum(4,3)
	
}
// // task : ram e code ta kivabe exicute hobe ?