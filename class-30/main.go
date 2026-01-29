// package main

// import "fmt"

// func main(){
// 	// pointer-> address of memory

// 	x:=20

// 	fmt.Print("x= ", x)

// 	p := &x // &->ampersand , &x-> address of x, & diye amra kono variable er adress ber korte pari

// 	*p=30 // *->address er value , * diye amra kono adress er value ber korte pari

// 		fmt.Print("x= ", x)
// 		fmt.Print("adress ", p) // p is the address of x
// 		fmt.Print("value at the address ", *p) // value of address p
// }

//--------------------------------------------------

// package main

// import "fmt"

// func print (numbers [3]int){
// 	fmt.print(numbers)
// }

// func main(){
	
// 	arr := [3]int{1,2,3} 

// 	print(arr)
// }

// task simulate
// aikhane first e compile hobe then exicution hobe then main call hobe then arr er jonno main e 3 ta sell er akta nam hobe arr tate 123 value gulo sell e bosbe. tarpor print call hobe arekta stack frame alocate hobe print function er jonno tate abar o 3 ta sell number name book hobe ar akhon numbers cell e ak ak kore age 1 tarpor 2 tarpor 3 aivabe arr er man gulo copy hoye numbers e bosbe.
// ate akta problem ase ak ak kore value copy hoye bosle onek time consuming hobe tai pointer aikhane helpfull.

// ----------------------------------------------------------------

package main

import "fmt"

func print (numbers *[3]int){
	fmt.print(numbers)
}

func main(){
	
	arr := [3]int{1,2,3} 

	print(&arr) // aikhane arr er starting address with koita sell e info ta jai.
}
// task aikhane print(&arr) call hole arr er memory adress print function e jabe then numbers e arr er address refarence bosbe . pore number print hole ta address ke follow kore value ane dekhai dibe

/*
accha * simbol tar mane to kono adress er value ta dibe tai to ? tahole aikhane 

```go
(numbers *[3]int) numbers er value keno bosche nah ?
```
1️⃣ Type Declaration এ * (আপনার case):
gofunc print(numbers *[3]int) {
    //         ↑
    //    এটা TYPE বলছে!
}
মানে: "numbers হলো একটা pointer যা [3]int এর address রাখবে"

2️⃣ Variable এর সাথে * (Dereference):
gofunc print(numbers *[3]int) {
    value := *numbers
    //       ↑
    //   এটা VALUE নিচ্ছে!
}
মানে: "numbers এর address এ যাও এবং সেখানকার value নিয়ে আসো"
*/


// arthath ai class e duita bishoi shiklam 1. pass by value 2. pass by reference qer gurotto