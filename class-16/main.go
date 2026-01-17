/*
akta folder e jodi duita file thake tahole duita file e akoi package name use korte hobe
/*
question: 
1. package kii ?
2. package scop kii ?
*/


// package main

// func main(){
// 	a := 2
// 	b :=4
// 		 add(a,b)
// }

// comand line calaw main.go -> bujar chesta koro j keno run hocche nah
// comand line calaw main.go add.go -> bujar chesta koro j keno run hocche 
// ai jinista ram e kivabe exicute hocche explain koro

// ---------------------------------------------------------------------prase-2

// package main

// func main(){
// 	a := 2
// 	b :=4
// 	addFunc(a,b) // jodi kono function akoi package er vitor hoi tahole seta import nah kore o kaj kore . aikhane jehutu difrent function er vitor tai age import korte hobe ar j function ta different package theke asbe setar variable baa function name capital letter hote hobe like AddFunc
// 	// runcomand go run main.go 
// 		 add(a,b)
// }

// my question: niyom ta amn keno j different package theke onno variable baa function keno import nah kore use kora jai nah abong variable baa function tar first letter capital letter hote hoi keno ar kintu same package hole import kora lage nah keno ?

// ---------------------------------------------------------------------prase-3


package main

import (
	"example.com/mathlib"
)

func main(){
	a := 2
	b :=4
	mathlib.AddFunc(a,b) // AddFunc call korar age add function er package ke import kore nite hobe.

	// my question: mathlib.AddFunc e mathlib. keno use hoise mathlib ki akta object ?
		 
}